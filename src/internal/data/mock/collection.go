package mock

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/canario/src/internal/common/readers"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/path"
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
	ItemData             []data.Item
	ItemsRecoverableErrs []error
	CState               data.CollectionState

	// For restore
	AuxItems map[string]data.Item
}

func (c Collection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ch := make(chan data.Item)

	go func() {
		defer close(ch)

		el := errs.Local()

		for _, item := range c.ItemData {
			it, ok := item.(*Item)
			if ok && it.ReadErr != nil {
				el.AddRecoverable(ctx, it.ReadErr)
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

var (
	_ data.BackupCollection  = &versionedBackupCollection{}
	_ data.RestoreCollection = &unversionedRestoreCollection{}
	_ data.Item              = &itemWrapper{}
)

type itemWrapper struct {
	data.Item
	reader io.ReadCloser
}

func (i *itemWrapper) ToReader() io.ReadCloser {
	return i.reader
}

func NewUnversionedRestoreCollection(
	t *testing.T,
	col data.RestoreCollection,
) *unversionedRestoreCollection {
	return &unversionedRestoreCollection{
		RestoreCollection: col,
		t:                 t,
	}
}

// unversionedRestoreCollection strips out version format headers on all items.
//
// Wrap data.RestoreCollections in this type if you don't need access to the
// version format header during tests and you know the item readers can't return
// an error.
type unversionedRestoreCollection struct {
	data.RestoreCollection
	t *testing.T
}

func (c *unversionedRestoreCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	res := make(chan data.Item)
	go func() {
		defer close(res)

		for item := range c.RestoreCollection.Items(ctx, errs) {
			r, err := readers.NewVersionedRestoreReader(item.ToReader())
			require.NoError(c.t, err, clues.ToCore(err))

			res <- &itemWrapper{
				Item:   item,
				reader: r,
			}
		}
	}()

	return res
}

func NewVersionedBackupCollection(
	t *testing.T,
	col data.BackupCollection,
) *versionedBackupCollection {
	return &versionedBackupCollection{
		BackupCollection: col,
		t:                t,
	}
}

// versionedBackupCollection injects basic version information on all items.
//
// Wrap data.BackupCollections in this type if you don't need to explicitly set
// the version format header during tests, aren't trying to check reader errors
// cases, and aren't populating backup details.
type versionedBackupCollection struct {
	data.BackupCollection
	t *testing.T
}

func (c *versionedBackupCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	res := make(chan data.Item)
	go func() {
		defer close(res)

		for item := range c.BackupCollection.Items(ctx, errs) {
			r, err := readers.NewVersionedBackupReader(
				readers.SerializationFormat{
					Version: readers.DefaultSerializationVersion,
				},
				item.ToReader())
			require.NoError(c.t, err, clues.ToCore(err))

			res <- &itemWrapper{
				Item:   item,
				reader: r,
			}
		}
	}()

	return res
}
