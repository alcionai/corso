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

func (sic *singleItemCollection) NextItem() (connector.DataStream, error) {
	if sic.used {
		return nil, io.EOF
	}

	sic.used = true
	return sic.stream, nil
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
