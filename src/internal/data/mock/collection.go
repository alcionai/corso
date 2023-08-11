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
// Item
// ---------------------------------------------------------------------------

var _ data.Item = &Item{}

type Item struct {
	DeletedFlag  bool
	ItemID       string
	ItemInfo     details.ItemInfo
	ItemSize     int64
	ModifiedTime time.Time
	Reader       io.ReadCloser
	ReadErr      error
}

func (s *Item) ID() string {
	return s.ItemID
}

func (s Item) Deleted() bool {
	return s.DeletedFlag
}

func (s *Item) ToReader() io.ReadCloser {
	if s.ReadErr != nil {
		return io.NopCloser(errReader{s.ReadErr})
	}

	return s.Reader
}

func (s *Item) Info() details.ItemInfo {
	return s.ItemInfo
}

func (s *Item) Size() int64 {
	return s.ItemSize
}

func (s *Item) ModTime() time.Time {
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

func (c Collection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
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

func (c Collection) FetchItemByName(ctx context.Context, name string) (data.Item, error) {
	return &Item{}, clues.New("not implemented")
}
