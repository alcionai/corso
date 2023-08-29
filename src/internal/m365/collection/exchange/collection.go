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
	_ data.BackupCollection = &Collection{}
	_ data.Item             = &Item{}
	_ data.ItemInfo         = &Item{}
	_ data.ItemModTime      = &Item{}
)

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4
)

// Collection implements the interface from data.Collection
// Structure holds data for an Exchange application for a single user
type Collection struct {
	user   string
	stream chan data.Item

	// added is a list of existing item IDs that were added to a container
	added map[string]struct{}
	// removed is a list of item IDs that were deleted from, or moved out, of a container
	removed map[string]struct{}

	getter itemGetterSerializer

	category      path.CategoryType
	statusUpdater support.StatusUpdater
	ctrl          control.Options

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

// NewExchangeDataCollection creates an ExchangeDataCollection.
// State of the collection is set as an observation of the current
// and previous paths.  If the curr path is nil, the state is assumed
// to be deleted.  If the prev path is nil, it is assumed newly created.
// If both are populated, then state is either moved (if they differ),
// or notMoved (if they match).
func NewCollection(
	user string,
	curr, prev path.Path,
	location *path.Builder,
	category path.CategoryType,
	items itemGetterSerializer,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) Collection {
	collection := Collection{
		added:           make(map[string]struct{}, 0),
		category:        category,
		ctrl:            ctrlOpts,
		stream:          make(chan data.Item, collectionChannelBufferSize),
		doNotMergeItems: doNotMergeItems,
		fullPath:        curr,
		getter:          items,
		locationPath:    location,
		prevPath:        prev,
		removed:         make(map[string]struct{}, 0),
		state:           data.StateOf(prev, curr),
		statusUpdater:   statusUpdater,
		user:            user,
	}

	return collection
}

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (col *Collection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	go col.streamItems(ctx, errs)
	return col.stream
}

// FullPath returns the Collection's fullPath []string
func (col *Collection) FullPath() path.Path {
	return col.fullPath
}

// LocationPath produces the Collection's full path, but with display names
// instead of IDs in the folders.  Only populated for Calendars.
func (col *Collection) LocationPath() *path.Builder {
	return col.locationPath
}

// TODO(ashmrtn): Fill in with previous path once the Controller compares old
// and new folder hierarchies.
func (col Collection) PreviousPath() path.Path {
	return col.prevPath
}

func (col Collection) State() data.CollectionState {
	return col.state
}

func (col Collection) DoNotMergeItems() bool {
	return col.doNotMergeItems
}

// ---------------------------------------------------------------------------
// Items() channel controller
// ---------------------------------------------------------------------------

// streamItems is a utility function that uses col.collectionType to be able to serialize
// all the M365IDs defined in the added field. data channel is closed by this function
func (col *Collection) streamItems(ctx context.Context, errs *fault.Bus) {
	var (
		success     int64
		totalBytes  int64
		wg          sync.WaitGroup
		colProgress chan<- struct{}

		user = col.user
		log  = logger.Ctx(ctx).With(
			"service", path.ExchangeService.String(),
			"category", col.category.String())
	)

	defer func() {
		col.finishPopulation(ctx, int(success), totalBytes, errs.Failure())
	}()

	if len(col.added)+len(col.removed) > 0 {
		colProgress = observe.CollectionProgress(
			ctx,
			col.fullPath.Category().String(),
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

			col.stream <- &Item{
				id:      id,
				modTime: time.Now().UTC(), // removed items have no modTime entry.
				deleted: true,
			}

			atomic.AddInt64(&success, 1)
			atomic.AddInt64(&totalBytes, 0)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

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

			item, info, err := col.getter.GetItem(
				ctx,
				user,
				id,
				col.ctrl.ToggleFeatures.ExchangeImmutableIDs,
				fault.New(true)) // temporary way to force a failFast error
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

			data, err := col.getter.Serialize(ctx, item, user, id)
			if err != nil {
				errs.AddRecoverable(ctx, clues.Wrap(err, "serializing item").Label(fault.LabelForceNoBackupCreation))
				return
			}

			// In case of mail the size of data is calc as- size of body content+size of attachment
			// in all other case the size is - total item's serialized size
			if info.Size <= 0 {
				info.Size = int64(len(data))
			}

			info.ParentPath = col.locationPath.String()

			col.stream <- &Item{
				id:      id,
				message: data,
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

// terminatePopulateSequence is a utility function used to close a Collection's data channel
// and to send the status update through the channel.
func (col *Collection) finishPopulation(
	ctx context.Context,
	success int,
	totalBytes int64,
	err error,
) {
	close(col.stream)

	attempted := len(col.added) + len(col.removed)
	status := support.CreateStatus(ctx,
		support.Backup,
		1,
		support.CollectionMetrics{
			Objects:   attempted,
			Successes: success,
			Bytes:     totalBytes,
		},
		col.fullPath.Folder(false))

	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())

	col.statusUpdater(status)
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

func (i *Item) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: i.info}
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
