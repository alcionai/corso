// streamstore implements helpers to store large
// data streams in a repository
package streamstore

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
