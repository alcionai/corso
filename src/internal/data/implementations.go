package data

import (
	"context"

	"github.com/alcionai/clues"

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
