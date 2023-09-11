package mock

import (
	"context"
	"io"
	"time"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Item
// ---------------------------------------------------------------------------

var (
	_ data.Item     = &Item{}
	_ data.ItemInfo = &Item{}
)

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

func (s *Item) Info() (details.ItemInfo, error) {
	return s.ItemInfo, nil
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

type Collection struct {
	Path                 path.Path
	Loc                  *path.Builder
	ItemData             []*Item
	ItemsRecoverableErrs []error
	CState               data.CollectionState
}

func (c Collection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ch := make(chan data.Item)

	go func() {
		defer close(ch)

		el := errs.Local()

		for _, item := range c.ItemData {
			if item.ReadErr != nil {
				el.AddRecoverable(ctx, item.ReadErr)
				continue
			}

			ch <- item
		}
	}()

	for _, err := range c.ItemsRecoverableErrs {
		errs.AddRecoverable(ctx, err)
	}

	return ch
}

func (c Collection) FullPath() path.Path {
	return c.Path
}

func (c Collection) PreviousPath() path.Path {
	return c.Path
}

func (c Collection) LocationPath() *path.Builder {
	return c.Loc
}

func (c Collection) State() data.CollectionState {
	return c.CState
}

func (c Collection) DoNotMergeItems() bool {
	return false
}

func (c Collection) FetchItemByName(
	ctx context.Context,
	name string,
) (data.Item, error) {
	res := c.AuxItems[name]
	if res == nil {
		return nil, data.ErrNotFound
	}

	return res, nil
}

var _ data.RestoreCollection = &RestoreCollection{}

type RestoreCollection struct {
	data.Collection
	AuxItems map[string]data.Item
}

func (rc RestoreCollection) FetchItemByName(
	ctx context.Context,
	name string,
) (data.Item, error) {
	res := rc.AuxItems[name]
	if res == nil {
		return nil, data.ErrNotFound
	}

	return res, nil
}
