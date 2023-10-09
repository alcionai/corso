package data

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/spatialcurrent/go-lazy/pkg/lazy"

	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

var (
	_ Item        = &unindexedPrefetchedItem{}
	_ ItemModTime = &unindexedPrefetchedItem{}

	_ Item        = &prefetchedItem{}
	_ ItemInfo    = &prefetchedItem{}
	_ ItemModTime = &prefetchedItem{}

	_ Item        = &unindexedLazyItem{}
	_ ItemModTime = &unindexedLazyItem{}

	_ Item        = &lazyItem{}
	_ ItemInfo    = &lazyItem{}
	_ ItemModTime = &lazyItem{}
)

func NewDeletedItem(itemID string) Item {
	return &unindexedPrefetchedItem{
		id:      itemID,
		deleted: true,
		// TODO(ashmrtn): This really doesn't need to be set since deleted items are
		// never passed to the actual storage engine. Setting it for now so tests
		// don't break.
		modTime: time.Now().UTC(),
	}
}

func NewUnindexedPrefetchedItem(
	reader io.ReadCloser,
	itemID string,
	modTime time.Time,
) (*unindexedPrefetchedItem, error) {
	r, err := readers.NewVersionedBackupReader(
		readers.SerializationFormat{Version: readers.DefaultSerializationVersion},
		reader)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return &unindexedPrefetchedItem{
		id:      itemID,
		reader:  r,
		modTime: modTime,
	}, nil
}

// unindexedPrefetchedItem represents a single item retrieved from the remote
// service.
//
// This item doesn't implement ItemInfo so it's safe to use for items like
// metadata that shouldn't appear in backup details.
type unindexedPrefetchedItem struct {
	id     string
	reader io.ReadCloser
	// modTime is the modified time of the item. It should match the modTime in
	// info if info is present. Here as a separate field so that deleted items
	// don't error out by trying to source it from info.
	modTime time.Time

	// deleted flags if this item has been removed in the remote service and
	// should be removed in storage as well.
	deleted bool
}

func (i unindexedPrefetchedItem) ID() string {
	return i.id
}

func (i *unindexedPrefetchedItem) ToReader() io.ReadCloser {
	return i.reader
}

func (i unindexedPrefetchedItem) Deleted() bool {
	return i.deleted
}

func (i unindexedPrefetchedItem) ModTime() time.Time {
	return i.modTime
}

func NewPrefetchedItem(
	reader io.ReadCloser,
	itemID string,
	info details.ItemInfo,
) (*prefetchedItem, error) {
	inner, err := NewUnindexedPrefetchedItem(reader, itemID, info.Modified())
	if err != nil {
		return nil, clues.Stack(err)
	}

	return &prefetchedItem{
		unindexedPrefetchedItem: inner,
		info:                    info,
	}, nil
}

// prefetchedItem represents a single item retrieved from the remote service.
//
// This item implements ItemInfo so it should be used for things that need to
// appear in backup details.
type prefetchedItem struct {
	*unindexedPrefetchedItem
	info details.ItemInfo
}

func (i prefetchedItem) Info() (details.ItemInfo, error) {
	return i.info, nil
}

type ItemDataGetter interface {
	GetData(
		context.Context,
		*fault.Bus,
	) (io.ReadCloser, *details.ItemInfo, bool, error)
}

func NewUnindexedLazyItem(
	ctx context.Context,
	itemGetter ItemDataGetter,
	itemID string,
	modTime time.Time,
	errs *fault.Bus,
) *unindexedLazyItem {
	return &unindexedLazyItem{
		ctx:        ctx,
		id:         itemID,
		itemGetter: itemGetter,
		modTime:    modTime,
		errs:       errs,
	}
}

// unindexedLazyItem represents a single item retrieved from the remote service.
// It lazily fetches the item's data when the first call to ToReader().Read() is
// made.
//
// This item doesn't implement ItemInfo so it's safe to use for items like
// metadata that shouldn't appear in backup details.
type unindexedLazyItem struct {
	ctx        context.Context
	mu         sync.Mutex
	id         string
	errs       *fault.Bus
	itemGetter ItemDataGetter

	modTime time.Time
	// info holds the details information for this item. Store a pointer in this
	// struct so we can tell if it's been set already or not.
	//
	// This also helps with garbage collection because now the golang garbage
	// collector can collect the lazyItem struct once the storage engine is done
	// with it. The ItemInfo struct needs to stick around until the end of the
	// backup though as backup details is written last.
	info *details.ItemInfo

	delInFlight bool
}

func (i *unindexedLazyItem) ID() string {
	return i.id
}

func (i *unindexedLazyItem) ToReader() io.ReadCloser {
	return lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
		// Don't allow getting Item info while trying to initialize said info.
		// GetData could be a long running call, but in theory nothing should happen
		// with the item until a reader is returned anyway.
		i.mu.Lock()
		defer i.mu.Unlock()

		reader, info, delInFlight, err := i.itemGetter.GetData(i.ctx, i.errs)
		if err != nil {
			return nil, clues.Stack(err)
		}

		format := readers.SerializationFormat{
			Version: readers.DefaultSerializationVersion,
		}

		// If an item was deleted then return an empty file so we don't fail the
		// backup and return a sentinel error when asked for ItemInfo so we don't
		// display the item in the backup.
		//
		// The item will be deleted from storage on the next backup when either the
		// delta token shows it's removed or we do a full backup (token expired
		// etc.) and the item isn't enumerated in that set.
		if delInFlight {
			logger.Ctx(i.ctx).Info("item not found")

			i.delInFlight = true
			format.DelInFlight = true
			r, err := readers.NewVersionedBackupReader(format)

			return r, clues.Stack(err).OrNil()
		}

		i.info = info

		r, err := readers.NewVersionedBackupReader(format, reader)

		return r, clues.Stack(err).OrNil()
	})
}

func (i *unindexedLazyItem) Deleted() bool {
	return false
}

func (i *unindexedLazyItem) ModTime() time.Time {
	return i.modTime
}

func NewLazyItem(
	ctx context.Context,
	itemGetter ItemDataGetter,
	itemID string,
	modTime time.Time,
	errs *fault.Bus,
) *lazyItem {
	return &lazyItem{
		unindexedLazyItem: NewUnindexedLazyItem(
			ctx,
			itemGetter,
			itemID,
			modTime,
			errs),
	}
}

// lazyItem represents a single item retrieved from the remote service. It
// lazily fetches the item's data when the first call to ToReader().Read() is
// made.
//
// This item implements ItemInfo so it should be used for things that need to
// appear in backup details.
type lazyItem struct {
	*unindexedLazyItem
}

func (i *lazyItem) Info() (details.ItemInfo, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.delInFlight {
		return details.ItemInfo{}, clues.Stack(ErrNotFound).WithClues(i.ctx)
	} else if i.info == nil {
		return details.ItemInfo{}, clues.New("requesting ItemInfo before data retrieval").
			WithClues(i.ctx)
	}

	return *i.info, nil
}
