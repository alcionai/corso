package kopia

import (
	"io"
	"strings"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/path"
)

var (
	_ data.Collection = &kopiaDataCollection{}
	_ data.Stream     = &kopiaDataStream{}
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

func (kdc kopiaDataCollection) FullPath() []string {
	// TODO(ashmrtn): Update this once data.Collection.FullPath supports
	// path.Path. Assumes no adversarial users that use "/" in their folder names.
	return strings.Split(kdc.path.String(), "/")
}

type kopiaDataStream struct {
	reader io.ReadCloser
	uuid   string
}

func (kds kopiaDataStream) ToReader() io.ReadCloser {
	return kds.reader
}

func (kds kopiaDataStream) UUID() string {
	return kds.uuid
}
