package graph

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ data.BackupCollection = emptyCollection{}

type emptyCollection struct {
	p  path.Path
	su support.StatusUpdater
}

func (c emptyCollection) Items(ctx context.Context, _ *fault.Bus) <-chan data.Stream {
	res := make(chan data.Stream)
	close(res)

	s := support.CreateStatus(ctx, support.Backup, 0, support.CollectionMetrics{}, "")
	c.su(s)

	return res
}

func (c emptyCollection) FullPath() path.Path {
	return c.p
}

func (c emptyCollection) PreviousPath() path.Path {
	return c.p
}

func (c emptyCollection) State() data.CollectionState {
	// This assumes we won't change the prefix path. Could probably use MovedState
	// as well if we do need to change things around.
	return data.NotMovedState
}

func (c emptyCollection) DoNotMergeItems() bool {
	return false
}

// ---------------------------------------------------------------------------
// base collections
// ---------------------------------------------------------------------------

func BaseCollections(
	ctx context.Context,
	colls []data.BackupCollection,
	tenant, rOwner string,
	service path.ServiceType,
	categories map[path.CategoryType]struct{},
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	var (
		res      = []data.BackupCollection{}
		el       = errs.Local()
		lastErr  error
		collKeys = map[string]struct{}{}
	)

	// won't catch deleted collections, since they have no FullPath
	for _, c := range colls {
		if c.FullPath() != nil {
			collKeys[c.FullPath().String()] = struct{}{}
		}
	}

	for cat := range categories {
		ictx := clues.Add(ctx, "base_service", service, "base_category", cat)

		p, err := path.ServicePrefix(tenant, rOwner, service, cat)
		if err != nil {
			// Shouldn't happen.
			err = clues.Wrap(err, "making path").WithClues(ictx)
			el.AddRecoverable(err)
			lastErr = err

			continue
		}

		// only add this collection if it doesn't already exist in the set.
		if _, ok := collKeys[p.String()]; !ok {
			res = append(res, emptyCollection{p: p, su: su})
		}
	}

	return res, lastErr
}

// ---------------------------------------------------------------------------
// prefix migration
// ---------------------------------------------------------------------------

var _ data.BackupCollection = prefixCollection{}

// TODO: move this out of graph.  /data would be a much better owner
// for a generic struct like this.  However, support.StatusUpdater makes
// it difficult to extract from this package in a generic way.
type prefixCollection struct {
	full, prev path.Path
	su         support.StatusUpdater
	state      data.CollectionState
}

func (c prefixCollection) Items(ctx context.Context, _ *fault.Bus) <-chan data.Stream {
	res := make(chan data.Stream)
	close(res)

	s := support.CreateStatus(ctx, support.Backup, 0, support.CollectionMetrics{}, "")
	c.su(s)

	return res
}

func (c prefixCollection) FullPath() path.Path {
	return c.full
}

func (c prefixCollection) PreviousPath() path.Path {
	return c.prev
}

func (c prefixCollection) State() data.CollectionState {
	return c.state
}

func (c prefixCollection) DoNotMergeItems() bool {
	return false
}

// Creates a new collection that only handles prefix pathing.
func NewPrefixCollection(prev, full path.Path, su support.StatusUpdater) (*prefixCollection, error) {
	if prev != nil {
		if len(prev.Item()) > 0 {
			return nil, clues.New("prefix collection previous path contains an item")
		}

		if len(prev.Folders()) > 0 {
			return nil, clues.New("prefix collection previous path contains folders")
		}
	}

	if full != nil {
		if len(full.Item()) > 0 {
			return nil, clues.New("prefix collection full path contains an item")
		}

		if len(full.Folders()) > 0 {
			return nil, clues.New("prefix collection full path contains folders")
		}
	}

	pc := &prefixCollection{
		prev:  prev,
		full:  full,
		su:    su,
		state: data.StateOf(prev, full),
	}

	if pc.state == data.DeletedState {
		return nil, clues.New("collection attempted to delete prefix")
	}

	if pc.state == data.NewState {
		return nil, clues.New("collection attempted to create a new prefix")
	}

	return pc, nil
}

// NewDeletedPrefixCollection creates a new collection that only handles
// deleting the prefix path.
func NewDeletedPrefixCollection(
	prev path.Path,
	su support.StatusUpdater,
) (*prefixCollection, error) {
	if prev == nil {
		return nil, clues.New("nil prefix path")
	}

	if len(prev.Item()) > 0 {
		return nil, clues.New("prefix collection previous path contains an item")
	}

	if len(prev.Folders()) > 0 {
		return nil, clues.New("prefix collection previous path contains folders")
	}

	pc := &prefixCollection{
		prev:  prev,
		full:  nil,
		su:    su,
		state: data.StateOf(prev, nil),
	}

	if pc.state != data.DeletedState {
		return nil, clues.New("collection didn't attempt to delete prefix")
	}

	return pc, nil
}
