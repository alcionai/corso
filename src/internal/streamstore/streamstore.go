// streamstore implements helpers to store large
// data streams in a repository
package streamstore

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/common/prefixmatcher"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/internal/kopia"
	"github.com/alcionai/canario/src/internal/kopia/inject"
	"github.com/alcionai/canario/src/internal/stats"
	"github.com/alcionai/canario/src/pkg/backup/identity"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/path"
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
func (ss *storeStreamer) Write(
	ctx context.Context,
	reasons []identity.Reasoner,
	errs *fault.Bus,
) (string, error) {
	ctx = clues.Add(ctx, "snapshot_type", "stream store")

	id, err := write(ctx, ss.kw, reasons, ss.dbcs, errs)
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

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Streamer is the core interface for all types of data streamed to and
// from the store.
type Streamer interface {
	Collector
	Writer
	Reader
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
	Write(context.Context, []identity.Reasoner, *fault.Bus) (string, error)
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
	item       data.Item
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

// Items() always returns a channel with a single data.Item
// representing the object to be persisted
func (dc *streamCollection) Items(context.Context, *fault.Bus) <-chan data.Item {
	items := make(chan data.Item, 1)
	defer close(items)
	items <- dc.item

	return items
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
		return nil, clues.StackWC(ctx, err)
	}

	// TODO: We could use an io.Pipe here to avoid a double copy but that
	// makes error handling a bit complicated
	bs, err := col.mr.Marshal()
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "marshalling body")
	}

	item, err := data.NewPrefetchedItem(
		io.NopCloser(bytes.NewReader(bs)),
		col.itemName,
		time.Now())
	if err != nil {
		return nil, clues.StackWC(ctx, err)
	}

	dc := streamCollection{
		folderPath: p,
		item:       item,
	}

	return &dc, nil
}

// write persists bytes to the store
func write(
	ctx context.Context,
	bup inject.BackupConsumer,
	reasons []identity.Reasoner,
	dbcs []data.BackupCollection,
	errs *fault.Bus,
) (string, error) {
	ctx = clues.Add(ctx, "collection_source", "streamstore")

	backupStats, _, _, err := bup.ConsumeBackupCollections(
		ctx,
		reasons,
		nil,
		dbcs,
		prefixmatcher.NopReader[map[string]struct{}](),
		nil,
		false,
		count.New(),
		errs)
	if err != nil {
		return "", clues.Wrap(err, "storing marshalled bytes in repository")
	}

	return backupStats.SnapshotID, nil
}

// read retrieves an object from the store
func read(
	ctx context.Context,
	snapshotID string,
	tenantID string,
	service path.ServiceType,
	col Collectable,
	rer inject.RestoreProducer,
	errs *fault.Bus,
) error {
	// construct the path of the container
	p, err := path.Builder{}.
		Append(col.itemName).
		ToStreamStorePath(tenantID, col.purpose, service, true)
	if err != nil {
		return clues.StackWC(ctx, err)
	}

	pd, err := p.Dir()
	if err != nil {
		return clues.StackWC(ctx, err)
	}

	ctx = clues.Add(ctx, "snapshot_id", snapshotID)

	cs, err := rer.ProduceRestoreCollections(
		ctx,
		snapshotID,
		[]path.RestorePaths{
			{
				StoragePath: p,
				RestorePath: pd,
			},
		},
		&stats.ByteCounter{},
		errs)
	if err != nil {
		return clues.Wrap(err, "retrieving data")
	}

	// Expect only 1 data collection
	if len(cs) != 1 {
		return clues.NewWC(ctx, "unexpected collection count").
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
			return clues.NewWC(ctx, "context cancelled waiting for data")

		case itemData, ok := <-items:
			if !ok {
				if !found {
					return clues.NewWC(ctx, "no data found")
				}

				return nil
			}

			if err := col.Unmr(itemData.ToReader()); err != nil {
				return clues.WrapWC(ctx, err, "unmarshalling data")
			}

			found = true
		}
	}
}
