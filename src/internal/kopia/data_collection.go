package kopia

import (
	"io"

	"github.com/alcionai/corso/internal/connector/data"
)

var _ data.DataCollection = &kopiaDataCollection{}
var _ data.DataStream = &kopiaDataStream{}

type kopiaDataCollection struct {
	path    []string
	streams []data.DataStream
}

func (kdc *kopiaDataCollection) Items() <-chan data.DataStream {
	res := make(chan data.DataStream)
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
