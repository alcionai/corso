package kopia

import (
	"io"

	"github.com/alcionai/corso/internal/connector"
)

var _ connector.DataCollection = &singleItemCollection{}
var _ connector.DataStream = &kopiaDataStream{}

// singleItemCollection implements DataCollection but only returns a single
// DataStream. It is not safe for concurrent use.
type singleItemCollection struct {
	path   []string
	stream connector.DataStream
	used   bool
}

func (sic *singleItemCollection) Items() <-chan connector.DataStream {
	if sic.used {
		return nil
	}

	sic.used = true
	res := make(chan connector.DataStream, 1)
	res <- sic.stream
	close(res)
	return res
}

func (sic singleItemCollection) FullPath() []string {
	return append([]string{}, sic.path...)
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
