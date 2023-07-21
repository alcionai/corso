package mock

import (
	"context"
	"io"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// stream
// ---------------------------------------------------------------------------

var _ data.Stream = &Stream{}

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

// ---------------------------------------------------------------------------
// collection
// ---------------------------------------------------------------------------

var (
	_ data.Collection        = &Collection{}
	_ data.BackupCollection  = &Collection{}
	_ data.RestoreCollection = &Collection{}
)

type Collection struct{}

func (c Collection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Stream {
	return nil
}

func (c Collection) FullPath() path.Path {
	return nil
}

func (c Collection) PreviousPath() path.Path {
	return nil
}

func (c Collection) State() data.CollectionState {
	return data.NewState
}

func (c Collection) DoNotMergeItems() bool {
	return true
}

func (c Collection) FetchItemByName(ctx context.Context, name string) (data.Stream, error) {
	return &Stream{}, clues.New("not implemented")
}
