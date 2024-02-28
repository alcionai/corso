package teamschats

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alcionai/clues"
	kjson "github.com/microsoft/kiota-serialization-json-go"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var _ data.BackupCollection = &lazyFetchCollection[chatsItemer]{}

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4
)

// updateStatus is a utility function used to send the status update through
// the channel.
func updateStatus(
	ctx context.Context,
	statusUpdater support.StatusUpdater,
	attempted int,
	streamedItems int64,
	totalBytes int64,
	folderPath string,
	err error,
) {
	status := support.CreateStatus(
		ctx,
		support.Backup,
		1,
		support.CollectionMetrics{
			Objects:   attempted,
			Successes: int(streamedItems),
			Bytes:     totalBytes,
		},
		folderPath)

	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())

	statusUpdater(status)
}

// State of the collection is set as an observation of the current
// and previous paths.  If the curr path is nil, the state is assumed
// to be deleted.  If the prev path is nil, it is assumed newly created.
// If both are populated, then state is either moved (if they differ),
// or notMoved (if they match).
func NewCollection[I chatsItemer](
	baseCol data.BaseCollection,
	filler fillItemer[I],
	protectedResource string,
	items []I,
	contains container[I],
	statusUpdater support.StatusUpdater,
) data.BackupCollection {
	return &lazyFetchCollection[I]{
		BaseCollection:    baseCol,
		items:             items,
		contains:          contains,
		filler:            filler,
		statusUpdater:     statusUpdater,
		stream:            make(chan data.Item, collectionChannelBufferSize),
		protectedResource: protectedResource,
	}
}

// -----------------------------------------------------------------------------
// lazyFetchCollection
// -----------------------------------------------------------------------------

type lazyFetchCollection[I chatsItemer] struct {
	data.BaseCollection
	protectedResource string
	stream            chan data.Item

	contains container[I]

	items []I

	filler fillItemer[I]

	statusUpdater support.StatusUpdater
}

func (col *lazyFetchCollection[I]) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	go col.streamItems(ctx, errs)
	return col.stream
}

func (col *lazyFetchCollection[I]) streamItems(ctx context.Context, errs *fault.Bus) {
	var (
		streamedItems   int64
		wg              sync.WaitGroup
		progressMessage chan<- struct{}
		el              = errs.Local()
	)

	ctx = clues.Add(ctx, "category", col.Category().String())

	defer func() {
		close(col.stream)
		logger.Ctx(ctx).Infow(
			"finished stream backup collection items",
			"stats", col.Counter.Values())

		updateStatus(
			ctx,
			col.statusUpdater,
			len(col.items),
			streamedItems,
			0,
			col.FullPath().Folder(false),
			errs.Failure())
	}()

	if len(col.items) > 0 {
		progressMessage = observe.CollectionProgress(
			ctx,
			col.Category().HumanString(),
			col.LocationPath().Elements())
		defer close(progressMessage)
	}

	semaphoreCh := make(chan struct{}, col.Opts().Parallelism.ItemFetch)
	defer close(semaphoreCh)

	// add any new items
	for _, item := range col.items {
		if el.Failure() != nil {
			break
		}

		modTime := ptr.Val(item.GetLastUpdatedDateTime())

		wg.Add(1)
		semaphoreCh <- struct{}{}

		go func(item I, modTime time.Time) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			itemID := ptr.Val(item.GetId())
			ictx := clues.Add(ctx, "item_id", itemID)

			col.stream <- data.NewLazyItemWithInfo(
				ictx,
				&lazyItemFiller[I]{
					modTime:      modTime,
					filler:       col.filler,
					resourceID:   col.protectedResource,
					item:         item,
					containerIDs: col.FullPath().Folders(),
					contains:     col.contains,
					parentPath:   col.LocationPath().String(),
				},
				itemID,
				modTime,
				col.Counter,
				el)

			atomic.AddInt64(&streamedItems, 1)

			if progressMessage != nil {
				progressMessage <- struct{}{}
			}
		}(item, modTime)
	}

	wg.Wait()
}

type lazyItemFiller[I chatsItemer] struct {
	filler       fillItemer[I]
	resourceID   string
	item         I
	parentPath   string
	containerIDs path.Elements
	modTime      time.Time
	contains     container[I]
}

func (lig *lazyItemFiller[I]) GetData(
	ctx context.Context,
	errs *fault.Bus,
) (io.ReadCloser, *details.ItemInfo, bool, error) {
	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	item, info, err := lig.filler.fillItem(ctx, lig.item)
	if err != nil {
		// For items that were deleted in flight, add the skip label so that
		// they don't lead to recoverable failures during backup.
		if clues.HasLabel(err, graph.LabelStatus(http.StatusNotFound)) || errors.Is(err, core.ErrNotFound) {
			logger.CtxErr(ctx, err).Info("item deleted in flight. skipping")

			// Returning delInFlight as true here for correctness, although the caller is going
			// to ignore it since we are returning an error.
			return nil, nil, true, clues.Wrap(err, "deleted item").Label(graph.LabelsSkippable)
		}

		err = clues.WrapWC(ctx, err, "getting item data").Label(fault.LabelForceNoBackupCreation)
		errs.AddRecoverable(ctx, err)

		return nil, nil, false, err
	}

	if err := writer.WriteObjectValue("", item); err != nil {
		err = clues.WrapWC(ctx, err, "writing item to serializer").Label(fault.LabelForceNoBackupCreation)
		errs.AddRecoverable(ctx, err)

		return nil, nil, false, err
	}

	itemData, err := writer.GetSerializedContent()
	if err != nil {
		err = clues.WrapWC(ctx, err, "serializing item").Label(fault.LabelForceNoBackupCreation)
		errs.AddRecoverable(ctx, err)

		return nil, nil, false, err
	}

	info.ParentPath = lig.parentPath
	// Update the mod time to what we already told kopia about. This is required
	// for proper details merging.
	info.Modified = lig.modTime

	return io.NopCloser(bytes.NewReader(itemData)),
		&details.ItemInfo{TeamsChats: info},
		false,
		nil
}
