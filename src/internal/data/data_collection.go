package data

import (
	"context"
	"io"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ------------------------------------------------------------------------------------------------
// standard ifaces
// ------------------------------------------------------------------------------------------------

var ErrNotFound = clues.New("not found")

type CollectionState int

const (
	NewState = CollectionState(iota)
	NotMovedState
	MovedState
	DeletedState
)

// A Collection represents the set of data within a single logical location
// denoted by FullPath.
type Collection interface {
	// Items returns a channel from which items in the collection can be read.
	// Each returned struct contains the next item in the collection
	// The channel is closed when there are no more items in the collection or if
	// an unrecoverable error caused an early termination in the sender.
	Items(ctx context.Context, errs *fault.Bus) <-chan Stream
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

type FetchItemByNamer interface {
	// Fetch retrieves an item with the given name from the Collection if it
	// exists. Items retrieved with Fetch may still appear in the channel returned
	// by Items().
	FetchItemByName(ctx context.Context, name string) (Stream, error)
}

// NoFetchRestoreCollection is a wrapper for a Collection that returns
// ErrNotFound for all Fetch calls.
type NoFetchRestoreCollection struct {
	Collection
	FetchItemByNamer
}

func (c NoFetchRestoreCollection) FetchItemByName(context.Context, string) (Stream, error) {
	return nil, ErrNotFound
}

type FetchRestoreCollection struct {
	Collection
	FetchItemByNamer
}

// Stream represents a single item within a Collection
// that can be consumed as a stream (it embeds io.Reader)
type Stream interface {
	// ToReader returns an io.Reader for the DataStream
	ToReader() io.ReadCloser
	// UUID provides a unique identifier for this data
	UUID() string
	// Deleted returns true if the item represented by this Stream has been
	// deleted and should be removed from the current in-progress backup.
	Deleted() bool
}

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

// StreamInfo is used to provide service specific
// information about the Stream
type StreamInfo interface {
	Info() details.ItemInfo
}

// StreamSize is used to provide size
// information about the Stream
type StreamSize interface {
	Size() int64
}

// StreamModTime is used to provide the modified time of the stream's data.
type StreamModTime interface {
	ModTime() time.Time
}

// StateOf lets us figure out the state of the collection from the
// previous and current path
func StateOf(prev, curr path.Path) CollectionState {
	if curr == nil || len(curr.String()) == 0 {
		return DeletedState
	}

	if prev == nil || len(prev.String()) == 0 {
		return NewState
	}

	if curr.String() != prev.String() {
		return MovedState
	}

	return NotMovedState
}

// ExportCollection is the interface that is returned to the SDK consumer
type ExportCollection interface {
	// GetBasePath gets the base path of the collection
	GetBasePath() string

	// GetItems gets the items within the collection(folder)
	GetItems(context.Context) <-chan ExportItem
}

type ExportItemData struct {
	Name string
	Body io.ReadCloser
}

// ExportItem is the item that is returned to the SDK consumer
type ExportItem struct {
	ID    string
	Data  ExportItemData
	Error error
}
