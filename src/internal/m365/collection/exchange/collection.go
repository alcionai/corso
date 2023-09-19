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
	"github.com/spatialcurrent/go-lazy/pkg/lazy"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.BackupCollection = &prefetchCollection{}
	_ data.Item             = &Item{}
	_ data.ItemInfo         = &Item{}
	_ data.ItemModTime      = &Item{}
)

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4
)

func NewBaseCollection(
	curr, prev path.Path,
	location *path.Builder,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) baseCollection {
	return baseCollection{
		ctrl:            ctrlOpts,
		doNotMergeItems: doNotMergeItems,
		fullPath:        curr,
		locationPath:    location,
		prevPath:        prev,
		state:           data.StateOf(prev, curr),
	}
}

// baseCollection contains basic functionality like returning path, location,
// and state information. It can be embedded in other implementations to provide
// this functionality.
//
// Functionality like how items are fetched is left to the embedding struct.
type baseCollection struct {
	ctrl control.Options

	// FullPath is the current hierarchical path used by this collection.
	fullPath path.Path

	// PrevPath is the previous hierarchical path used by this collection.
	// It may be the same as fullPath, if the folder was not renamed or
	// moved.  It will be empty on its first retrieval.
	prevPath path.Path

	// LocationPath contains the path with human-readable display names.
	// IE: "/Inbox/Important" instead of "/abcdxyz123/algha=lgkhal=t"
	locationPath *path.Builder

	state data.CollectionState

	// doNotMergeItems should only be true if the old delta token expired.
	doNotMergeItems bool
}

// FullPath returns the baseCollection's fullPath []string
func (col *baseCollection) FullPath() path.Path {
	return col.fullPath
}

// LocationPath produces the baseCollection's full path, but with display names
// instead of IDs in the folders.  Only populated for Calendars.
func (col *baseCollection) LocationPath() *path.Builder {
	return col.locationPath
}

func (col baseCollection) PreviousPath() path.Path {
	return col.prevPath
}

func (col baseCollection) State() data.CollectionState {
	return col.state
}

func (col baseCollection) DoNotMergeItems() bool {
	return col.doNotMergeItems
}

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
	bc baseCollection,
	user string,
	items itemGetterSerializer,
	added map[string]time.Time,
	removed []string,
	validModTimes bool,
	statusUpdater support.StatusUpdater,
) data.BackupCollection {
	add := maps.Clone(added)
	remove := make(map[string]struct{}, len(removed))

	// Remove any deleted IDs from the set of added IDs because items that are
	// deleted and then restored will have a different ID than they did
	// originally.
	//
	// TODO(ashmrtn): If we switch to immutable IDs then we'll need to handle this
	// sort of operation in the pager since this would become order-dependent
	// unless Graph started consolidating the changes into a single delta result.
	for _, r := range removed {
		delete(add, r)

		remove[r] = struct{}{}
	}

	if !validModTimes {
		return &prefetchCollection{
			baseCollection: bc,
			user:           user,
			added:          add,
			removed:        remove,
			getter:         items,
			statusUpdater:  statusUpdater,
		}
	}

	return &lazyFetchCollection{
		baseCollection: bc,
		user:           user,
		added:          add,
		removed:        remove,
		getter:         items,
		statusUpdater:  statusUpdater,
	}
}

// prefetchCollection implements the interface from data.BackupCollection
// Structure holds data for an Exchange application for a single user
type prefetchCollection struct {
	baseCollection

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
			"category", col.FullPath().Category().String())
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
			col.FullPath().Category().HumanString(),
			col.LocationPath().Elements())
		defer close(colProgress)
	}

	semaphoreCh := make(chan struct{}, col.ctrl.Parallelism.ItemFetch)
	defer close(semaphoreCh)

	// delete all removed items
	for id := range col.removed {
		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			stream <- &Item{
				id:      id,
				modTime: time.Now().UTC(), // removed items have no modTime entry.
				deleted: true,
			}

			atomic.AddInt64(&success, 1)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	parentPath := col.LocationPath().String()

	// add any new items
	for id := range col.added {
		if errs.Failure() != nil {
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
				col.ctrl.ToggleFeatures.ExchangeImmutableIDs,
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
					errs.AddRecoverable(ctx, clues.Wrap(err, "fetching item").Label(fault.LabelForceNoBackupCreation))
				}

				return
			}

			stream <- &Item{
				id:      id,
				message: itemData,
				info:    info,
				modTime: info.Modified,
			}

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
	baseCollection

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
		totalBytes  int64
		wg          sync.WaitGroup
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
			totalBytes,
			col.FullPath().Folder(false),
			errs.Failure())
	}()

	if len(col.added)+len(col.removed) > 0 {
		colProgress = observe.CollectionProgress(
			ctx,
			col.FullPath().Category().String(),
			col.LocationPath().Elements())
		defer close(colProgress)
	}

	// delete all removed items
	for id := range col.removed {
		stream <- &Item{
			id:      id,
			modTime: time.Now().UTC(), // removed items have no modTime entry.
			deleted: true,
		}

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
			"category", col.FullPath().Category().String())

		stream <- &lazyItem{
			ctx:          ictx,
			userID:       user,
			id:           id,
			getter:       col.getter,
			modTime:      modTime,
			immutableIDs: col.ctrl.ToggleFeatures.ExchangeImmutableIDs,
			parentPath:   parentPath,
			errs:         errs,
		}

		atomic.AddInt64(&success, 1)

		if colProgress != nil {
			colProgress <- struct{}{}
		}
	}

	wg.Wait()
}

// Item represents a single item retrieved from exchange
type Item struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
	info    *details.ExchangeInfo // temporary change to bring populate function into directory
	// TODO(ashmrtn): Can probably eventually be sourced from info as there's a
	// request to provide modtime in ItemInfo structs.
	modTime time.Time

	// true if the item was marked by graph as deleted.
	deleted bool
}

func (i *Item) ID() string {
	return i.id
}

func (i *Item) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(i.message))
}

func (i Item) Deleted() bool {
	return i.deleted
}

func (i *Item) Info() (details.ItemInfo, error) {
	return details.ItemInfo{Exchange: i.info}, nil
}

func (i *Item) ModTime() time.Time {
	return i.modTime
}

func NewItem(
	identifier string,
	dataBytes []byte,
	detail details.ExchangeInfo,
	modTime time.Time,
) Item {
	return Item{
		id:      identifier,
		message: dataBytes,
		info:    &detail,
		modTime: modTime,
	}
}

// lazyItem represents a single item retrieved from exchange that lazily fetches
// the item's data when the first call to ToReader().Read() is made.
type lazyItem struct {
	ctx        context.Context
	userID     string
	id         string
	parentPath string
	getter     itemGetterSerializer
	errs       *fault.Bus

	modTime time.Time
	// info holds the Exchnage-specific details information for this item. Store
	// a pointer in this struct so the golang garbage collector can collect the
	// Item struct once kopia is done with it. The ExchangeInfo struct needs to
	// stick around until the end of the backup though as backup details is
	// written last.
	info *details.ExchangeInfo

	immutableIDs bool

	delInFlight bool
}

func (i lazyItem) ID() string {
	return i.id
}

func (i *lazyItem) ToReader() io.ReadCloser {
	return lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
		itemData, info, err := getItemAndInfo(
			i.ctx,
			i.getter,
			i.userID,
			i.ID(),
			i.immutableIDs,
			i.parentPath)
		if err != nil {
			// If an item was deleted then return an empty file so we don't fail
			// the backup and return a sentinel error when asked for ItemInfo so
			// we don't display the item in the backup.
			//
			// The item will be deleted from kopia on the next backup when the
			// delta token shows it's removed.
			if graph.IsErrDeletedInFlight(err) {
				logger.CtxErr(i.ctx, err).Info("item not found")

				i.delInFlight = true

				return io.NopCloser(bytes.NewReader([]byte{})), nil
			}

			err = clues.Stack(err)
			i.errs.AddRecoverable(i.ctx, err)

			return nil, err
		}

		i.info = info
		// Update the mod time to what we already told kopia about. This is required
		// for proper details merging.
		i.info.Modified = i.modTime

		return io.NopCloser(bytes.NewReader(itemData)), nil
	})
}

func (i lazyItem) Deleted() bool {
	return false
}

func (i lazyItem) Info() (details.ItemInfo, error) {
	if i.delInFlight {
		return details.ItemInfo{}, clues.Stack(data.ErrNotFound).WithClues(i.ctx)
	} else if i.info == nil {
		return details.ItemInfo{}, clues.New("requesting ItemInfo before data retrieval").
			WithClues(i.ctx)
	}

	return details.ItemInfo{Exchange: i.info}, nil
}

func (i lazyItem) ModTime() time.Time {
	return i.modTime
}
