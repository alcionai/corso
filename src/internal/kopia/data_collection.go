package kopia

import (
	"io"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/path"
)

var (
	_ data.Collection  = &kopiaDataCollection{}
	_ data.ByteCounter = &kopiaDataCollection{}
	_ data.Stream      = &kopiaDataStream{}
	_ data.StreamSize  = &kopiaDataStream{}
)

type kopiaDataCollection struct {
	path       path.Path
	streams    []data.Stream
	countBytes int64
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

func (kdc *kopiaDataCollection) CountBytes(i int64) {
	kdc.countBytes += i
}

func (kdc kopiaDataCollection) BytesCounted() int64 {
	return kdc.countBytes
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

func (kds kopiaDataStream) Size() int64 {
	return kds.size
}
