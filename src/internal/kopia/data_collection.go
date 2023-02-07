package kopia

import (
	"io"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.RestoreCollection = &kopiaDataCollection{}
	_ data.Stream            = &kopiaDataStream{}
)

type kopiaDataCollection struct {
	path    path.Path
	streams []data.Stream
}

func (kdc *kopiaDataCollection) Items() <-chan data.Stream {
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
