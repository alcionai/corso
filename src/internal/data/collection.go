package data

import (
	"context"
	"io"
	"time"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Collections
// ---------------------------------------------------------------------------

// A Collection represents the set of data within a single logical location
// denoted by FullPath.
type Collection interface {
	// Items returns a channel from which items in the collection can be read.
	// Each returned struct contains the next item in the collection
	// The channel is closed when there are no more items in the collection or if
	// an unrecoverable error caused an early termination in the sender.
	Items(ctx context.Context, errs *fault.Bus) <-chan Item
	// FullPath returns a path struct that acts as a metadata tag for this
	// Collection.
	FullPath() path.Path
}

// BackupCollection is an extension of Collection that is used during backups.
type BackupCollection interface {
	Collection
	// PreviousPath returns the path.Path this collection used to reside at
	// (according to the M365 ID for the container) if the collection was moved or
	// renamed. Returns nil if the collection is new.
	PreviousPath() path.Path
	// State represents changes to the Collection compared to the last backup
	// involving the Collection. State changes are based on the M365 ID of the
	// Collection, not just the path the collection resides at. Collections that
	// are in the same location as they were in the previous backup should be
	// marked as NotMovedState. Renaming or reparenting the Collection counts as
	// Moved. Collections marked as Deleted will be removed from the current
	// backup along with all items and Collections below them in the hierarchy
	// unless said items/Collections were moved.
	State() CollectionState
	// DoNotMergeItems informs kopia that the collection is rebuilding its contents
	// from scratch, and that any items currently stored at the previousPath should
	// be skipped during the process of merging historical data into the new backup.
	// This flag is normally expected to be false.  It should only be flagged under
	// specific circumstances.  Example: if the link token used for incremental queries
	// expires or otherwise becomes unusable, thus requiring the backup producer to
	// re-discover all data in the container.  This flag only affects the path of the
	// collection, and does not cascade to subfolders.
	DoNotMergeItems() bool
}

// RestoreCollection is an extension of Collection that is used during restores.
type RestoreCollection interface {
	Collection
	FetchItemByNamer
}

// ---------------------------------------------------------------------------
// Items
// ---------------------------------------------------------------------------

// Item represents a single item within a Collection
type Item interface {
	// ToReader returns an io.Reader with the item's data
	ToReader() io.ReadCloser
	// ID provides a unique identifier for this item
	ID() string
	// Deleted returns true if the item represented by this Stream has been
	// deleted and should be removed from the current in-progress backup.
	Deleted() bool
}

// ItemInfo returns the details.ItemInfo for the item.
type ItemInfo interface {
	Info() details.ItemInfo
}

// ItemSize returns the size of the item in bytes.
type ItemSize interface {
	Size() int64
}

// ItemModTime provides the last modified time of the item.
//
// If an item implements ItemModTime and ItemInfo it should return the same
// value here as in item.Info().Modified().
type ItemModTime interface {
	ModTime() time.Time
}

type FetchItemByNamer interface {
	// Fetch retrieves an item with the given name from the Collection if it
	// exists. Items retrieved with Fetch may still appear in the channel returned
	// by Items().
	FetchItemByName(ctx context.Context, name string) (Item, error)
}

// ---------------------------------------------------------------------------
// Paths
// ---------------------------------------------------------------------------

// LocationPather provides a LocationPath describing the path with Display Names
// instead of canonical IDs
type LocationPather interface {
	LocationPath() *path.Builder
}

// PreviousLocationPather provides both the current location of the collection
// as well as the location of the item in the previous backup.
//
// TODO(ashmrtn): If we guarantee that we persist the location of collections in
// addition to the path of the item then we could just have a single
// *LocationPather interface with current and previous location functions.
type PreviousLocationPather interface {
	LocationPather
	PreviousLocationPath() details.LocationIDer
}
