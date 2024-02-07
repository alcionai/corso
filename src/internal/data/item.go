package data

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/alcionai/clues"
	"github.com/spatialcurrent/go-lazy/pkg/lazy"

	"github.com/alcionai/canario/src/internal/common/readers"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/logger"
)

var (
	_ Item        = &prefetchedItem{}
	_ ItemModTime = &prefetchedItem{}

	_ Item        = &prefetchedItemWithInfo{}
	_ ItemInfo    = &prefetchedItemWithInfo{}
	_ ItemModTime = &prefetchedItemWithInfo{}

	_ Item        = &lazyItem{}
	_ ItemModTime = &lazyItem{}

	_ Item        = &lazyItemWithInfo{}
	_ ItemInfo    = &lazyItemWithInfo{}
	_ ItemModTime = &lazyItemWithInfo{}
)

func NewDeletedItem(itemID string) Item {
	return &prefetchedItem{
		id:      itemID,
		deleted: true,
		// TODO(ashmrtn): This really doesn't need to be set since deleted items are
		// never passed to the actual storage engine. Setting it for now so tests
		// don't break.
		modTime: time.Now().UTC(),
	}
}

func NewPrefetchedItem(
	reader io.ReadCloser,
	itemID string,
	modTime time.Time,
) (*prefetchedItem, error) {
	r, err := readers.NewVersionedBackupReader(
		readers.SerializationFormat{Version: readers.DefaultSerializationVersion},
		reader)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return &prefetchedItem{
		id:      itemID,
		reader:  r,
		modTime: modTime,
	}, nil
}

// prefetchedItem represents a single item retrieved from the remote service.
//
// This item doesn't implement ItemInfo so it's safe to use for items like
// metadata that shouldn't appear in backup details.
type prefetchedItem struct {
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

func (i prefetchedItem) ID() string {
	return i.id
}

func (i *prefetchedItem) ToReader() io.ReadCloser {
	return i.reader
}

func (i prefetchedItem) Deleted() bool {
	return i.deleted
}

func (i prefetchedItem) ModTime() time.Time {
	return i.modTime
}

func NewPrefetchedItemWithInfo(
	reader io.ReadCloser,
	itemID string,
	info details.ItemInfo,
) (*prefetchedItemWithInfo, error) {
	inner, err := NewPrefetchedItem(reader, itemID, info.Modified())
	if err != nil {
		return nil, clues.Stack(err)
	}

	return &prefetchedItemWithInfo{
		prefetchedItem: inner,
		info:           info,
	}, nil
}

// prefetchedItemWithInfo represents a single item retrieved from the remote
// service.
//
// This item implements ItemInfo so it should be used for things that need to
// appear in backup details.
type prefetchedItemWithInfo struct {
	*prefetchedItem
	info details.ItemInfo
}

func (i prefetchedItemWithInfo) Info() (details.ItemInfo, error) {
	return i.info, nil
}

type ItemDataGetter interface {
	GetData(
		context.Context,
		*fault.Bus,
	) (io.ReadCloser, *details.ItemInfo, bool, error)
}

func NewLazyItem(
	ctx context.Context,
	itemGetter ItemDataGetter,
	itemID string,
	modTime time.Time,
	counter *count.Bus,
	errs *fault.Bus,
) *lazyItem {
	return &lazyItem{
		ctx:        ctx,
		id:         itemID,
		itemGetter: itemGetter,
		modTime:    modTime,
		counter:    counter,
		errs:       errs,
	}
}

// lazyItem represents a single item retrieved from the remote service. It
// lazily fetches the item's data when the first call to ToReader().Read() is
// made.
//
// This item doesn't implement ItemInfo so it's safe to use for items like
// metadata that shouldn't appear in backup details.
type lazyItem struct {
	ctx        context.Context
	mu         sync.Mutex
	id         string
	counter    *count.Bus
	errs       *fault.Bus
	itemGetter ItemDataGetter

	modTime time.Time
	// info holds the details information for this item. Store a pointer in this
	// struct so we can tell if it's been set already or not.
	//
	// This also helps with garbage collection because now the golang garbage
	// collector can collect the lazyItemWithInfo struct once the storage engine
	// is done with it. The ItemInfo struct needs to stick around until the end of
	// the backup though as backup details is written last.
	info *details.ItemInfo

	delInFlight bool
}

func (i *lazyItem) ID() string {
	return i.id
}

func (i *lazyItem) ToReader() io.ReadCloser {
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
			i.counter.Inc(count.LazyDeletedInFlight)

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

func (i *lazyItem) Deleted() bool {
	return false
}

func (i *lazyItem) ModTime() time.Time {
	return i.modTime
}

func NewLazyItemWithInfo(
	ctx context.Context,
	itemGetter ItemDataGetter,
	itemID string,
	modTime time.Time,
	counter *count.Bus,
	errs *fault.Bus,
) *lazyItemWithInfo {
	return &lazyItemWithInfo{
		lazyItem: NewLazyItem(
			ctx,
			itemGetter,
			itemID,
			modTime,
			counter,
			errs),
	}
}

// lazyItemWithInfo represents a single item retrieved from the remote service.
// It lazily fetches the item's data when the first call to ToReader().Read() is
// made.
//
// This item implements ItemInfo so it should be used for things that need to
// appear in backup details.
type lazyItemWithInfo struct {
	*lazyItem
}

func (i *lazyItemWithInfo) Info() (details.ItemInfo, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.delInFlight {
		return details.ItemInfo{}, clues.StackWC(i.ctx, ErrNotFound)
	} else if i.info == nil {
		return details.ItemInfo{}, clues.NewWC(i.ctx, "requesting ItemInfo before data retrieval")
	}

	return *i.info, nil
}
