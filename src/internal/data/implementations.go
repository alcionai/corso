package data

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

var ErrNotFound = clues.New("not found")

type CollectionState int

const (
	NewState      CollectionState = 0
	NotMovedState CollectionState = 1
	MovedState    CollectionState = 2
	DeletedState  CollectionState = 3
)

type FetchRestoreCollection struct {
	Collection
	FetchItemByNamer
}

// NoFetchRestoreCollection is a wrapper for a Collection that returns
// ErrNotFound for all Fetch calls.
type NoFetchRestoreCollection struct {
	Collection
}

func (c NoFetchRestoreCollection) FetchItemByName(context.Context, string) (Item, error) {
	return nil, ErrNotFound
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

// -----------------------------------------------------------------------------
// BaseCollection
// -----------------------------------------------------------------------------

func NewBaseCollection(
	curr, prev path.Path,
	location *path.Builder,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) BaseCollection {
	return BaseCollection{
		ctrl:            ctrlOpts,
		doNotMergeItems: doNotMergeItems,
		fullPath:        curr,
		locationPath:    location,
		prevPath:        prev,
		state:           StateOf(prev, curr),
	}
}

// BaseCollection contains basic functionality like returning path, location,
// and state information. It can be embedded in other implementations to provide
// this functionality.
//
// Functionality like how items are fetched is left to the embedding struct.
type BaseCollection struct {
	ctrl control.Options

	// FullPath is the current hierarchical path used by this collection.
	fullPath path.Path

	// PrevPath is the previous hierarchical path used by this collection.
	// It may be the same as fullPath, if the folder was not renamed or
	// moved.  It will be empty on its first retrieval.
	prevPath path.Path

	// LocationPath contains the path with human-readable display names.
	// IE: "/Inbox/Important" instead of "/abcdxyz123/algha=lgkhal=t"
	locationPath *path.Builder

	state CollectionState

	// doNotMergeItems should only be true if the old delta token expired.
	doNotMergeItems bool
}

// FullPath returns the BaseCollection's fullPath []string
func (col *BaseCollection) FullPath() path.Path {
	return col.fullPath
}

// LocationPath produces the BaseCollection's full path, but with display names
// instead of IDs in the folders.  Only populated for Calendars.
func (col *BaseCollection) LocationPath() *path.Builder {
	return col.locationPath
}

func (col BaseCollection) PreviousPath() path.Path {
	return col.prevPath
}

func (col BaseCollection) State() CollectionState {
	return col.state
}

func (col BaseCollection) DoNotMergeItems() bool {
	return col.doNotMergeItems
}

func (col BaseCollection) Opts() control.Options {
	return col.ctrl
}
