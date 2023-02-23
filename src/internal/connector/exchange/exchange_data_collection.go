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
	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.BackupCollection = &Collection{}
	_ data.Stream           = &Stream{}
	_ data.StreamInfo       = &Stream{}
	_ data.StreamModTime    = &Stream{}
)

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4

	// Outlooks expects max 4 concurrent requests
	// https://learn.microsoft.com/en-us/graph/throttling-limits#outlook-service-limits
	urlPrefetchChannelBufferSize = 4
)

type itemer interface {
	GetItem(
		ctx context.Context,
		user, itemID string,
		errs *fault.Errors,
	) (serialization.Parsable, *details.ExchangeInfo, error)
	Serialize(
		ctx context.Context,
		item serialization.Parsable,
		user, itemID string,
	) ([]byte, error)
}

// Collection implements the interface from data.Collection
// Structure holds data for an Exchange application for a single user
type Collection struct {
	// M365 user
	user string // M365 user
	data chan data.Stream

	// added is a list of existing item IDs that were added to a container
	added map[string]struct{}
	// removed is a list of item IDs that were deleted from, or moved out, of a container
	removed map[string]struct{}

	items itemer

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
	// Currently only implemented for Exchange Calendars.
	locationPath path.Path

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
	curr, prev, location path.Path,
	category path.CategoryType,
	items itemer,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) Collection {
	collection := Collection{
		added:           make(map[string]struct{}, 0),
		category:        category,
		ctrl:            ctrlOpts,
		data:            make(chan data.Stream, collectionChannelBufferSize),
		doNotMergeItems: doNotMergeItems,
		fullPath:        curr,
		items:           items,
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
func (col *Collection) Items(ctx context.Context, errs *fault.Errors) <-chan data.Stream {
	go col.streamItems(ctx, errs)
	return col.data
}

// FullPath returns the Collection's fullPath []string
func (col *Collection) FullPath() path.Path {
	return col.fullPath
}

// LocationPath produces the Collection's full path, but with display names
// instead of IDs in the folders.  Only populated for Calendars.
func (col *Collection) LocationPath() path.Path {
	return col.locationPath
}

// TODO(ashmrtn): Fill in with previous path once GraphConnector compares old
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
func (col *Collection) streamItems(ctx context.Context, errs *fault.Errors) {
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
		col.finishPopulation(ctx, int(success), totalBytes, errs.Err())
	}()

	if len(col.added)+len(col.removed) > 0 {
		var closer func()
		colProgress, closer = observe.CollectionProgress(
			ctx,
			col.fullPath.Category().String(),
			observe.PII(user),
			observe.PII(col.fullPath.Folder(false)))

		go closer()

		defer func() {
			close(colProgress)
		}()
	}

	// Limit the max number of active requests to GC
	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	// delete all removed items
	for id := range col.removed {
		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			col.data <- &Stream{
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
		if errs.Err() != nil {
			break
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			item, info, err := getItemWithRetries(
				ctx,
				user,
				id,
				col.items,
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
					errs.Add(clues.Wrap(err, "fetching item"))
				}

				return
			}

			data, err := col.items.Serialize(ctx, item, user, id)
			if err != nil {
				errs.Add(clues.Wrap(err, "serializing item"))
				return
			}

			info.Size = int64(len(data))

			col.data <- &Stream{
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

// get an item while handling retry and backoff.
func getItemWithRetries(
	ctx context.Context,
	userID, itemID string,
	items itemer,
	errs *fault.Errors,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	item, info, err := items.GetItem(ctx, userID, itemID, errs)
	if err != nil {
		return nil, nil, err
	}

	return item, info, nil
}

// terminatePopulateSequence is a utility function used to close a Collection's data channel
// and to send the status update through the channel.
func (col *Collection) finishPopulation(
	ctx context.Context,
	success int,
	totalBytes int64,
	err error,
) {
	close(col.data)

	attempted := len(col.added) + len(col.removed)
	status := support.CreateStatus(ctx,
		support.Backup,
		1,
		support.CollectionMetrics{
			Objects:    attempted,
			Successes:  success,
			TotalBytes: totalBytes,
		},
		err,
		col.fullPath.Folder(false))

	logger.Ctx(ctx).Debugw("done streaming items", "status", status.String())

	col.statusUpdater(status)
}

// Stream represents a single item retrieved from exchange
type Stream struct {
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

func (od *Stream) UUID() string {
	return od.id
}

func (od *Stream) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.message))
}

func (od Stream) Deleted() bool {
	return od.deleted
}

func (od *Stream) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: od.info}
}

func (od *Stream) ModTime() time.Time {
	return od.modTime
}

// NewStream constructor for exchange.Stream object
func NewStream(identifier string, dataBytes []byte, detail details.ExchangeInfo, modTime time.Time) Stream {
	return Stream{
		id:      identifier,
		message: dataBytes,
		info:    &detail,
		modTime: modTime,
	}
}
