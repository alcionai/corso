package groups

import (
	"bytes"
	"context"
	"io"
	"sync"
	"sync/atomic"

	"github.com/alcionai/clues"
	kjson "github.com/microsoft/kiota-serialization-json-go"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var _ data.BackupCollection = &Collection[graph.GetIDer, groupsItemer]{}

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4
)

type Collection[C graph.GetIDer, I groupsItemer] struct {
	data.BaseCollection
	protectedResource string
	stream            chan data.Item

	contains container[C]

	// added is a list of existing item IDs that were added to a container
	added map[string]struct{}
	// removed is a list of item IDs that were deleted from, or moved out, of a container
	removed map[string]struct{}

	getAndAugment getItemAndAugmentInfoer[C, I]

	statusUpdater support.StatusUpdater
}

// NewExchangeDataCollection creates an ExchangeDataCollection.
// State of the collection is set as an observation of the current
// and previous paths.  If the curr path is nil, the state is assumed
// to be deleted.  If the prev path is nil, it is assumed newly created.
// If both are populated, then state is either moved (if they differ),
// or notMoved (if they match).
func NewCollection[C graph.GetIDer, I groupsItemer](
	baseCol data.BaseCollection,
	getAndAugment getItemAndAugmentInfoer[C, I],
	protectedResource string,
	added map[string]struct{},
	removed map[string]struct{},
	contains container[C],
	statusUpdater support.StatusUpdater,
) Collection[C, I] {
	collection := Collection[C, I]{
		BaseCollection:    baseCol,
		added:             added,
		contains:          contains,
		getAndAugment:     getAndAugment,
		removed:           removed,
		statusUpdater:     statusUpdater,
		stream:            make(chan data.Item, collectionChannelBufferSize),
		protectedResource: protectedResource,
	}

	return collection
}

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (col *Collection[C, I]) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	go col.streamItems(ctx, errs)
	return col.stream
}

// ---------------------------------------------------------------------------
// items() production
// ---------------------------------------------------------------------------

func (col *Collection[C, I]) streamItems(ctx context.Context, errs *fault.Bus) {
	var (
		streamedItems int64
		totalBytes    int64
		wg            sync.WaitGroup
		colProgress   chan<- struct{}
		el            = errs.Local()
	)

	ctx = clues.Add(ctx, "category", col.Category().String())

	defer func() {
		logger.Ctx(ctx).Infow(
			"finished stream backup collection items",
			"stats", col.Counter.Values())
		col.finishPopulation(ctx, streamedItems, totalBytes, errs.Failure())
	}()

	if len(col.added)+len(col.removed) > 0 {
		colProgress = observe.CollectionProgress(
			ctx,
			col.Category().HumanString(),
			col.LocationPath().Elements())
		defer close(colProgress)
	}

	semaphoreCh := make(chan struct{}, col.Opts().Parallelism.ItemFetch)
	defer close(semaphoreCh)

	// delete all removed items
	for id := range col.removed {
		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			col.stream <- data.NewDeletedItem(id)

			atomic.AddInt64(&streamedItems, 1)
			col.Counter.Inc(count.StreamItemsRemoved)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	// add any new items
	for id := range col.added {
		if el.Failure() != nil {
			break
		}

		wg.Add(1)
		semaphoreCh <- struct{}{}

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			writer := kjson.NewJsonSerializationWriter()
			defer writer.Close()

			item, info, err := col.getAndAugment.getItem(
				ctx,
				col.protectedResource,
				col.FullPath().Folders(),
				id)
			if err != nil {
				err = clues.Wrap(err, "getting channel message data").Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			col.getAndAugment.augmentItemInfo(info, col.contains.container)

			if err := writer.WriteObjectValue("", item); err != nil {
				err = clues.Wrap(err, "writing channel message to serializer").Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			itemData, err := writer.GetSerializedContent()
			if err != nil {
				err = clues.Wrap(err, "serializing channel message").Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			info.ParentPath = col.LocationPath().String()

			storeItem, err := data.NewPrefetchedItemWithInfo(
				io.NopCloser(bytes.NewReader(itemData)),
				id,
				details.ItemInfo{Groups: info})
			if err != nil {
				err := clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation)
				el.AddRecoverable(ctx, err)

				return
			}

			col.stream <- storeItem

			atomic.AddInt64(&streamedItems, 1)
			atomic.AddInt64(&totalBytes, info.Size)

			if col.Counter.Inc(count.StreamItemsAdded)%1000 == 0 {
				logger.Ctx(ctx).Infow("item stream progress", "stats", col.Counter.Values())
			}

			col.Counter.Add(count.StreamBytesAdded, info.Size)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	wg.Wait()
}

// finishPopulation is a utility function used to close a Collection's data channel
// and to send the status update through the channel.
func (col *Collection[C, I]) finishPopulation(
	ctx context.Context,
	streamedItems, totalBytes int64,
	err error,
) {
	close(col.stream)

	attempted := len(col.added) + len(col.removed)
	status := support.CreateStatus(
		ctx,
		support.Backup,
		1,
		support.CollectionMetrics{
			Objects:   attempted,
			Successes: int(streamedItems),
			Bytes:     totalBytes,
		},
		col.FullPath().Folder(false))

	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())

	col.statusUpdater(status)
}
