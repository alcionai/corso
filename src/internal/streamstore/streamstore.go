// streamstore implements helpers to store large
// data streams in a repository
package streamstore

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Streamer is the core interface for all types of data streamed to and
// from the store.
type Streamer interface {
	Writer
	Reader
	Delete(context.Context, string) error
}

type Reader interface {
	Read(context.Context, string, Unmarshaller, *fault.Bus) error
}

type Writer interface {
	Write(context.Context, Marshaller, *fault.Bus) (string, error)
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

type backuper interface {
	BackupCollections(
		ctx context.Context,
		bases []kopia.IncrementalBase,
		cs []data.BackupCollection,
		excluded map[string]struct{},
		tags map[string]string,
		buildTreeWithBase bool,
		errs *fault.Bus,
	) (*kopia.BackupStats, *details.Builder, map[string]kopia.PrevRefs, error)
}

// write persists bytes to the store
func write(
	ctx context.Context,
	tenantID string,
	service path.ServiceType,
	collectionPurpose string,
	itemName string,
	bup backuper,
	mr Marshaller,
	errs *fault.Bus,
) (string, error) {
	// construct the path of the container
	p, err := path.Builder{}.
		ToStreamStorePath(tenantID, collectionPurpose, service, false)
	if err != nil {
		return "", clues.Stack(err).WithClues(ctx)
	}

	// TODO: We could use an io.Pipe here to avoid a double copy but that
	// makes error handling a bit complicated
	bs, err := mr.Marshal()
	if err != nil {
		return "", clues.Wrap(err, "marshalling body").WithClues(ctx)
	}

	dc := &streamCollection{
		folderPath: p,
		item: &streamItem{
			name: itemName,
			data: bs,
		},
	}

	backupStats, _, _, err := bup.BackupCollections(
		ctx,
		nil,
		[]data.BackupCollection{dc},
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
	tenantID string,
	service path.ServiceType,
	collectionPurpose string,
	itemName string,
	id string,
	rer restorer,
	umr Unmarshaller,
	errs *fault.Bus,
) error {
	// construct the path of the container
	p, err := path.Builder{}.
		Append(itemName).
		ToStreamStorePath(tenantID, collectionPurpose, service, true)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	cs, err := rer.RestoreMultipleItems(
		ctx,
		id,
		[]path.Path{p},
		&stats.ByteCounter{},
		errs)
	if err != nil {
		return errors.Wrap(err, "retrieving data")
	}

	// Expect only 1 data collection
	if len(cs) != 1 {
		return clues.New("greater than 1 collection found").
			WithClues(ctx).
			With("collection_count", len(cs))
	}

	var (
		col   = cs[0]
		found = false
		items = col.Items(ctx, errs)
	)

	for {
		select {
		case <-ctx.Done():
			return clues.New("context cancelled waiting for data").WithClues(ctx)

		case itemData, ok := <-items:
			if !ok {
				if !found {
					return clues.New("no backup found").WithClues(ctx)
				}

				return nil
			}

			if err := umr(itemData.ToReader()); err != nil {
				return clues.Wrap(err, "unmarshalling data").WithClues(ctx)
			}

			found = true
		}
	}
}
