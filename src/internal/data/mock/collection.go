package mock

import (
	"io"
	"time"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

type Stream struct {
	ID           string
	Reader       io.ReadCloser
	ReadErr      error
	ItemSize     int64
	ModifiedTime time.Time
	DeletedFlag  bool
	ItemInfo     details.ItemInfo
}

func (s *Stream) UUID() string {
	return s.ID
}

func (s Stream) Deleted() bool {
	return s.DeletedFlag
}

func (s *Stream) ToReader() io.ReadCloser {
	if s.ReadErr != nil {
		return io.NopCloser(errReader{s.ReadErr})
	}

	return s.Reader
}

func (s *Stream) Info() details.ItemInfo {
	return s.ItemInfo
}

func (s *Stream) Size() int64 {
	return s.ItemSize
}

func (s *Stream) ModTime() time.Time {
	return s.ModifiedTime
}

type errReader struct {
	readErr error
}

func (er errReader) Read([]byte) (int, error) {
	return 0, er.readErr
}
