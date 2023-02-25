package kopia

import (
	"context"
	"io"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.RestoreCollection = &kopiaDataCollection{}
	_ data.Stream            = &kopiaDataStream{}
)

type kopiaDataCollection struct {
	path         path.Path
	streams      []data.Stream
	snapshotRoot fs.Entry
	counter      ByteCounter
}

func (kdc *kopiaDataCollection) Items(
	ctx context.Context,
	_ *fault.Bus, // unused, just matching the interface
) <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for _, s := range kdc.streams {
			res <- s
		}
	}()

	return res
}

func (kdc kopiaDataCollection) FullPath() path.Path {
	return kdc.path
}

func (kdc kopiaDataCollection) Fetch(
	ctx context.Context,
	name string,
) (data.Stream, error) {
	if kdc.snapshotRoot == nil {
		return nil, clues.New("no snapshot root")
	}

	p, err := kdc.FullPath().Append(name, true)
	if err != nil {
		return nil, clues.Wrap(err, "creating item path")
	}

	// TODO(ashmrtn): We could possibly hold a reference to the folder this
	// collection corresponds to, but that requires larger changes for the
	// creation of these collections.
	return getItemStream(ctx, p, kdc.snapshotRoot, kdc.counter)
}

type kopiaDataStream struct {
	reader io.ReadCloser
	uuid   string
	size   int64
}

func (kds kopiaDataStream) ToReader() io.ReadCloser {
	return kds.reader
}

func (kds kopiaDataStream) UUID() string {
	return kds.uuid
}

func (kds kopiaDataStream) Deleted() bool {
	return false
}

func (kds kopiaDataStream) Size() int64 {
	return kds.size
}
