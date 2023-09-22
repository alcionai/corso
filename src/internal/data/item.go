package data

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/alcionai/clues"
	"github.com/spatialcurrent/go-lazy/pkg/lazy"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

func NewDeletedItem(itemID string) Item {
	return &prefetchedItem{
		id:      itemID,
		deleted: true,
		// TODO(ashmrtn): This really doesn't need to be set since deleted items are
		// never passed to the actual storage engine. Setting it for now so tests
		// don't break.
		modTime: time.Now(),
	}
}

func NewPrefetchedItem(
	reader io.ReadCloser,
	itemID string,
	info details.ItemInfo,
) Item {
	return &prefetchedItem{
		id:      itemID,
		reader:  reader,
		info:    info,
		modTime: info.Modified(),
	}
}

// prefetchedItem represents a single item retrieved from the remote service.
type prefetchedItem struct {
	id     string
	reader io.ReadCloser
	info   details.ItemInfo
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

func (i prefetchedItem) Info() (details.ItemInfo, error) {
	return i.info, nil
}

func (i prefetchedItem) ModTime() time.Time {
	return i.modTime
}

type ItemDataGetter interface {
	GetData(context.Context) (io.ReadCloser, *details.ItemInfo, bool, error)
}

func NewLazyItem(
	ctx context.Context,
	itemGetter ItemDataGetter,
	itemID string,
	modTime time.Time,
	errs *fault.Bus,
) Item {
	return &lazyItem{
		ctx:        ctx,
		id:         itemID,
		itemGetter: itemGetter,
		modTime:    modTime,
		errs:       errs,
	}
}

// lazyItem represents a single item retrieved from the remote service. It
// lazily fetches the item's data when the first call to ToReader().Read() is
// made.
type lazyItem struct {
	ctx        context.Context
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

func (i lazyItem) ID() string {
	return i.id
}

func (i *lazyItem) ToReader() io.ReadCloser {
	return lazy.NewLazyReadCloser(func() (io.ReadCloser, error) {
		reader, info, delInFlight, err := i.itemGetter.GetData(i.ctx)
		if err != nil {
			err = clues.Stack(err)
			i.errs.AddRecoverable(i.ctx, err)

			return nil, err
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

			return io.NopCloser(bytes.NewReader([]byte{})), nil
		}

		i.info = info

		return reader, nil
	})
}

func (i lazyItem) Deleted() bool {
	return false
}

func (i lazyItem) Info() (details.ItemInfo, error) {
	if i.delInFlight {
		return details.ItemInfo{}, clues.Stack(ErrNotFound).WithClues(i.ctx)
	} else if i.info == nil {
		return details.ItemInfo{}, clues.New("requesting ItemInfo before data retrieval").
			WithClues(i.ctx)
	}

	return *i.info, nil
}

func (i lazyItem) ModTime() time.Time {
	return i.modTime
}
