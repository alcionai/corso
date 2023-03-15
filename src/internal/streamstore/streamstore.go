// streamstore implements helpers to store large
// data streams in a repository
package streamstore

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

var _ Streamer = &storeStreamer{}

type storeStreamer struct {
	kw      *kopia.Wrapper
	tenant  string
	service path.ServiceType
	dbcs    []data.BackupCollection
}

// NewStreamer creates a new streamstore Streamer for stream writing metadata files
// to the store.
func NewStreamer(
	kw *kopia.Wrapper,
	tenant string,
	service path.ServiceType,
) *storeStreamer {
	return &storeStreamer{
		kw:      kw,
		tenant:  tenant,
		service: service,
	}
}

// Collect eagerly searializes the marshalable bytes in the collectable into a
// data.BackupCollection.  The collection is stored within the storeStreamer
// for persistence when Write is called.
func (ss *storeStreamer) Collect(ctx context.Context, col Collectable) error {
	cs, err := collect(ctx, ss.tenant, ss.service, col)
	if err != nil {
		return clues.Wrap(err, "collecting data for stream store")
	}

	ss.dbcs = append(ss.dbcs, cs)

	return nil
}

// Write persists the collected objects in the stream store
func (ss *storeStreamer) Write(ctx context.Context, errs *fault.Bus) (string, error) {
	id, err := write(ctx, ss.kw, ss.dbcs, errs)
	if err != nil {
		return "", clues.Wrap(err, "writing to stream store")
	}

	return id, nil
}

// Read reads a collector object from the kopia repository
func (ss *storeStreamer) Read(ctx context.Context, snapshotID string, col Collectable, errs *fault.Bus) error {
	err := read(ctx, snapshotID, ss.tenant, ss.service, col, ss.kw, errs)
	if err != nil {
		return clues.Wrap(err, "reading from stream store")
	}

	return nil
}

// Delete deletes a `details.Details` object from the kopia repository
func (ss *storeStreamer) Delete(ctx context.Context, detailsID string) error {
	err := ss.kw.DeleteSnapshot(ctx, detailsID)
	if err != nil {
		return errors.Wrap(err, "deleting snapshot in stream store")
	}

	return nil
}

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Streamer is the core interface for all types of data streamed to and
// from the store.
type Streamer interface {
	Collector
	Writer
	Reader
	Delete(context.Context, string) error
}

type CollectorWriter interface {
	Collector
	Writer
}

type Collector interface {
	Collect(context.Context, Collectable) error
}

type Reader interface {
	Read(context.Context, string, Collectable, *fault.Bus) error
}

type Writer interface {
	Write(context.Context, *fault.Bus) (string, error)
}

// Marshallers are used to convert structs into bytes to be persisted in the store.
type Marshaller interface {
	Marshal() ([]byte, error)
}

// Unmarshallers are used to serialize the bytes in the store into the original struct.
type Unmarshaller func(io.ReadCloser) error

// ---------------------------------------------------------------------------
// collection
// ---------------------------------------------------------------------------

// streamCollection is a data.BackupCollection used to persist
// a single data stream
type streamCollection struct {
	// folderPath indicates what level in the hierarchy this collection
	// represents
	folderPath path.Path
	item       *streamItem
}

func (dc *streamCollection) FullPath() path.Path {
	return dc.folderPath
}

func (dc *streamCollection) PreviousPath() path.Path {
	return nil
}

func (dc *streamCollection) State() data.CollectionState {
	return data.NewState
}

func (dc *streamCollection) DoNotMergeItems() bool {
	return false
}

// Items() always returns a channel with a single data.Stream
// representing the object to be persisted
func (dc *streamCollection) Items(context.Context, *fault.Bus) <-chan data.Stream {
	items := make(chan data.Stream, 1)
	defer close(items)
	items <- dc.item

	return items
}

// ---------------------------------------------------------------------------
// item
// ---------------------------------------------------------------------------

type streamItem struct {
	name string
	data []byte
}

func (di *streamItem) UUID() string {
	return di.name
}

func (di *streamItem) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(di.data))
}

func (di *streamItem) Deleted() bool {
	return false
}

// ---------------------------------------------------------------------------
// common reader/writer/deleter
// ---------------------------------------------------------------------------

// collect aggregates a collection of bytes
func collect(
	ctx context.Context,
	tenantID string,
	service path.ServiceType,
	col Collectable,
) (data.BackupCollection, error) {
	// construct the path of the container
	p, err := path.Builder{}.ToStreamStorePath(tenantID, col.purpose, service, false)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	// TODO: We could use an io.Pipe here to avoid a double copy but that
	// makes error handling a bit complicated
	bs, err := col.mr.Marshal()
	if err != nil {
		return nil, clues.Wrap(err, "marshalling body").WithClues(ctx)
	}

	dc := streamCollection{
		folderPath: p,
		item: &streamItem{
			name: col.itemName,
			data: bs,
		},
	}

	return &dc, nil
}

type backuper interface {
	BackupCollections(
		ctx context.Context,
		bases []kopia.IncrementalBase,
		cs []data.BackupCollection,
		globalExcludeSet map[string]map[string]struct{},
		tags map[string]string,
		buildTreeWithBase bool,
		errs *fault.Bus,
	) (*kopia.BackupStats, *details.Builder, map[string]kopia.PrevRefs, error)
}

// write persists bytes to the store
func write(
	ctx context.Context,
	bup backuper,
	dbcs []data.BackupCollection,
	errs *fault.Bus,
) (string, error) {
	backupStats, _, _, err := bup.BackupCollections(
		ctx,
		nil,
		dbcs,
		nil,
		nil,
		false,
		errs)
	if err != nil {
		return "", errors.Wrap(err, "storing marshalled bytes in repository")
	}

	return backupStats.SnapshotID, nil
}

type restorer interface {
	RestoreMultipleItems(
		ctx context.Context,
		snapshotID string,
		paths []path.Path,
		bc kopia.ByteCounter,
		errs *fault.Bus,
	) ([]data.RestoreCollection, error)
}

// read retrieves an object from the store
func read(
	ctx context.Context,
	snapshotID string,
	tenantID string,
	service path.ServiceType,
	col Collectable,
	rer restorer,
	errs *fault.Bus,
) error {
	// construct the path of the container
	p, err := path.Builder{}.
		Append(col.itemName).
		ToStreamStorePath(tenantID, col.purpose, service, true)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	cs, err := rer.RestoreMultipleItems(
		ctx,
		snapshotID,
		[]path.Path{p},
		&stats.ByteCounter{},
		errs)
	if err != nil {
		return errors.Wrap(err, "retrieving data")
	}

	// Expect only 1 data collection
	if len(cs) != 1 {
		return clues.New("unexpected collection count").
			WithClues(ctx).
			With("collection_count", len(cs))
	}

	var (
		c     = cs[0]
		found = false
		items = c.Items(ctx, errs)
	)

	for {
		select {
		case <-ctx.Done():
			return clues.New("context cancelled waiting for data").WithClues(ctx)

		case itemData, ok := <-items:
			if !ok {
				if !found {
					return clues.New("no data found").WithClues(ctx)
				}

				return nil
			}

			if err := col.Unmr(itemData.ToReader()); err != nil {
				return clues.Wrap(err, "unmarshalling data").WithClues(ctx)
			}

			found = true
		}
	}
}
