// Package exchange provides support for retrieving M365 Exchange objects
// from M365 servers using the Graph API. M365 object support centers
// on the applications: Mail, Contacts, and Calendar.
package exchange

import (
	"bytes"
	"context"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var (
	_ data.BackupCollection = &prefetchCollection{}
	_ data.BackupCollection = &lazyFetchCollection{}
)

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
	success int,
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
			Successes: success,
			Bytes:     totalBytes,
		},
		folderPath)

	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())

	statusUpdater(status)
}

func getItemAndInfo(
	ctx context.Context,
	getter itemGetterSerializer,
	userID string,
	id string,
	useImmutableIDs bool,
	parentPath string,
) ([]byte, *details.ExchangeInfo, error) {
	item, info, err := getter.GetItem(
		ctx,
		userID,
		id,
		useImmutableIDs,
		fault.New(true)) // temporary way to force a failFast error
	if err != nil {
		return nil, nil, clues.Wrap(err, "fetching item").
			WithClues(ctx).
			Label(fault.LabelForceNoBackupCreation)
	}

	itemData, err := getter.Serialize(ctx, item, userID, id)
	if err != nil {
		return nil, nil, clues.Wrap(err, "serializing item").WithClues(ctx)
	}

	// In case of mail the size of itemData is calc as- size of body content+size of attachment
	// in all other case the size is - total item's serialized size
	if info.Size <= 0 {
		info.Size = int64(len(itemData))
	}

	info.ParentPath = parentPath

	return itemData, info, nil
}

// NewExchangeDataCollection creates an ExchangeDataCollection.
// State of the collection is set as an observation of the current
// and previous paths.  If the curr path is nil, the state is assumed
// to be deleted.  If the prev path is nil, it is assumed newly created.
// If both are populated, then state is either moved (if they differ),
// or notMoved (if they match).
func NewCollection(
	bc data.BaseCollection,
	user string,
	items itemGetterSerializer,
	origAdded map[string]time.Time,
	origRemoved []string,
	validModTimes bool,
	statusUpdater support.StatusUpdater,
	counter *count.Bus,
) data.BackupCollection {
	added := maps.Clone(origAdded)
	removed := make(map[string]struct{}, len(origRemoved))

	// Remove any deleted IDs from the set of added IDs because items that are
	// deleted and then restored will have a different ID than they did
	// originally.
	//
	// TODO(ashmrtn): If we switch to immutable IDs then we'll need to handle this
	// sort of operation in the pager since this would become order-dependent
	// unless Graph started consolidating the changes into a single delta result.
	for _, r := range origRemoved {
		delete(added, r)

		removed[r] = struct{}{}
	}

	counter.Add(count.ItemsAdded, int64(len(added)))
	counter.Add(count.ItemsRemoved, int64(len(removed)))

	if !validModTimes {
		return &prefetchCollection{
			BaseCollection: bc,
			user:           user,
			added:          added,
			removed:        removed,
			getter:         items,
			statusUpdater:  statusUpdater,
		}
	}

	return &lazyFetchCollection{
		BaseCollection: bc,
		user:           user,
		added:          added,
		removed:        removed,
		getter:         items,
		statusUpdater:  statusUpdater,
	}
}

// prefetchCollection implements the interface from data.BackupCollection
// Structure holds data for an Exchange application for a single user
type prefetchCollection struct {
	data.BaseCollection

	user string

	// added is a list of existing item IDs that were added to a container
	added map[string]time.Time
	// removed is a list of item IDs that were deleted from, or moved out, of a container
	removed map[string]struct{}

	getter itemGetterSerializer

	statusUpdater support.StatusUpdater
}

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (col *prefetchCollection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	stream := make(chan data.Item, collectionChannelBufferSize)
	go col.streamItems(ctx, stream, errs)

	return stream
}

// streamItems is a utility function that uses col.collectionType to be able to serialize
// all the M365IDs defined in the added field. data channel is closed by this function
func (col *prefetchCollection) streamItems(
	ctx context.Context,
	stream chan<- data.Item,
	errs *fault.Bus,
) {
	var (
		success     int64
		totalBytes  int64
		wg          sync.WaitGroup
		colProgress chan<- struct{}

		user = col.user
		log  = logger.Ctx(ctx).With(
			"service", path.ExchangeService.String(),
			"category", col.Category().String())
	)

	defer func() {
		close(stream)
		updateStatus(
			ctx,
			col.statusUpdater,
			len(col.added)+len(col.removed),
			int(success),
			totalBytes,
			col.FullPath().Folder(false),
			errs.Failure())
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

			stream <- data.NewDeletedItem(id)

			atomic.AddInt64(&success, 1)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	var (
		parentPath = col.LocationPath().String()
		el         = errs.Local()
	)

	// add any new items
	for id := range col.added {
		if el.Failure() != nil {
			break
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			itemData, info, err := getItemAndInfo(
				ctx,
				col.getter,
				user,
				id,
				col.Opts().ToggleFeatures.ExchangeImmutableIDs,
				parentPath)
			if err != nil {
				// Don't report errors for deleted items as there's no way for us to
				// back up data that is gone. Record it as a "success", since there's
				// nothing else we can do, and not reporting it will make the status
				// investigation upset.
				if graph.IsErrDeletedInFlight(err) {
					atomic.AddInt64(&success, 1)
					log.With("err", err).Infow("item not found", clues.InErr(err).Slice()...)
				} else {
					el.AddRecoverable(ctx, clues.Wrap(err, "fetching item").Label(fault.LabelForceNoBackupCreation))
				}

				return
			}

			item, err := data.NewPrefetchedItemWithInfo(
				io.NopCloser(bytes.NewReader(itemData)),
				id,
				details.ItemInfo{Exchange: info})
			if err != nil {
				el.AddRecoverable(
					ctx,
					clues.Stack(err).
						WithClues(ctx).
						Label(fault.LabelForceNoBackupCreation))

				return
			}

			stream <- item

			atomic.AddInt64(&success, 1)
			atomic.AddInt64(&totalBytes, info.Size)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	wg.Wait()
}

// -----------------------------------------------------------------------------
// lazyFetchCollection
// -----------------------------------------------------------------------------

// lazyFetchCollection implements the interface from data.BackupCollection
// Structure holds data for an Exchange application for a single user. It lazily
// fetches the data associated with each item when kopia requests it during
// upload.
//
// When accounting for stats, items are marked as successful when the basic
// information (path and mod time) is handed to kopia. Total bytes across all
// items is not tracked.
type lazyFetchCollection struct {
	data.BaseCollection

	user string

	// added is a list of existing item IDs that were added to a container
	added map[string]time.Time
	// removed is a list of item IDs that were deleted from, or moved out, of a container
	removed map[string]struct{}

	getter itemGetterSerializer

	statusUpdater support.StatusUpdater
}

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (col *lazyFetchCollection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	stream := make(chan data.Item, collectionChannelBufferSize)
	go col.streamItems(ctx, stream, errs)

	return stream
}

// streamItems is a utility function that uses col.collectionType to be able to
// serialize all the M365IDs defined in the added field. data channel is closed
// by this function.
func (col *lazyFetchCollection) streamItems(
	ctx context.Context,
	stream chan<- data.Item,
	errs *fault.Bus,
) {
	var (
		success     int64
		colProgress chan<- struct{}

		user = col.user
	)

	defer func() {
		close(stream)
		updateStatus(
			ctx,
			col.statusUpdater,
			len(col.added)+len(col.removed),
			int(success),
			0,
			col.FullPath().Folder(false),
			errs.Failure())
	}()

	if len(col.added)+len(col.removed) > 0 {
		colProgress = observe.CollectionProgress(
			ctx,
			col.Category().HumanString(),
			col.LocationPath().Elements())
		defer close(colProgress)
	}

	// delete all removed items
	for id := range col.removed {
		stream <- data.NewDeletedItem(id)

		atomic.AddInt64(&success, 1)

		if colProgress != nil {
			colProgress <- struct{}{}
		}
	}

	parentPath := col.LocationPath().String()

	// add any new items
	for id, modTime := range col.added {
		if errs.Failure() != nil {
			break
		}

		ictx := clues.Add(
			ctx,
			"item_id", id,
			"parent_path", path.LoggableDir(parentPath),
			"service", path.ExchangeService.String(),
			"category", col.Category().String())

		stream <- data.NewLazyItemWithInfo(
			ictx,
			&lazyItemGetter{
				userID:       user,
				itemID:       id,
				getter:       col.getter,
				modTime:      modTime,
				immutableIDs: col.Opts().ToggleFeatures.ExchangeImmutableIDs,
				parentPath:   parentPath,
			},
			id,
			modTime,
			errs)

		atomic.AddInt64(&success, 1)

		if colProgress != nil {
			colProgress <- struct{}{}
		}
	}
}

type lazyItemGetter struct {
	getter       itemGetterSerializer
	userID       string
	itemID       string
	parentPath   string
	modTime      time.Time
	immutableIDs bool
}

func (lig *lazyItemGetter) GetData(
	ctx context.Context,
	errs *fault.Bus,
) (io.ReadCloser, *details.ItemInfo, bool, error) {
	itemData, info, err := getItemAndInfo(
		ctx,
		lig.getter,
		lig.userID,
		lig.itemID,
		lig.immutableIDs,
		lig.parentPath)
	if err != nil {
		// If an item was deleted then return an empty file so we don't fail
		// the backup and return a sentinel error when asked for ItemInfo so
		// we don't display the item in the backup.
		//
		// The item will be deleted from kopia on the next backup when the
		// delta token shows it's removed.
		if graph.IsErrDeletedInFlight(err) {
			logger.CtxErr(ctx, err).Info("item not found")
			return nil, nil, true, nil
		}

		err = clues.Stack(err)
		errs.AddRecoverable(ctx, err)

		return nil, nil, false, err
	}

	// Update the mod time to what we already told kopia about. This is required
	// for proper details merging.
	info.Modified = lig.modTime

	return io.NopCloser(bytes.NewReader(itemData)),
		&details.ItemInfo{Exchange: info},
		false,
		nil
}
