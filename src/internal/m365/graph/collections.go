package graph

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ data.BackupCollection = prefixCollection{}

// TODO: move this out of graph.  /data would be a much better owner
// for a generic struct like this.  However, support.StatusUpdater makes
// it difficult to extract from this package in a generic way.
type prefixCollection struct {
	full  path.Path
	prev  path.Path
	su    support.StatusUpdater
	state data.CollectionState
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

		full, err := path.ServicePrefix(tenant, rOwner, service, cat)
		if err != nil {
			// Shouldn't happen.
			err = clues.Wrap(err, "making path").WithClues(ictx)
			el.AddRecoverable(ctx, err)
			lastErr = err

			continue
		}

		// only add this collection if it doesn't already exist in the set.
		if _, ok := collKeys[full.String()]; !ok {
			res = append(res, &prefixCollection{
				prev:  full,
				full:  full,
				su:    su,
				state: data.StateOf(full, full),
			})
		}
	}

	return res, lastErr
}

// ---------------------------------------------------------------------------
// prefix migration
// ---------------------------------------------------------------------------

// Creates a new collection that only handles prefix pathing.
func NewPrefixCollection(
	prev, full path.Path,
	su support.StatusUpdater,
) (*prefixCollection, error) {
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
