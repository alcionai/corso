package kopia

import (
	"io"

	"github.com/alcionai/corso/internal/connector"
)

var _ connector.DataCollection = &kopiaDataCollection{}
var _ connector.DataStream = &kopiaDataStream{}

type kopiaDataCollection struct {
	path    []string
	streams []connector.DataStream
}

func (kdc *kopiaDataCollection) Items() <-chan connector.DataStream {
	res := make(chan connector.DataStream)
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
