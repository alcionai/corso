package kopia

import (
	"io"

	"github.com/alcionai/corso/internal/data"
)

var (
	_ data.Collection = &kopiaDataCollection{}
	_ data.Stream     = &kopiaDataStream{}
)

type kopiaDataCollection struct {
	path    []string
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
	return append([]string{}, kdc.path...)
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
