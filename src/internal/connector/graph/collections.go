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

func BaseCollections(
	ctx context.Context,
	tenant, rOwner string,
	service path.ServiceType,
	categories map[path.CategoryType]struct{},
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	var (
		res     = []data.BackupCollection{}
		el      = errs.Local()
		lastErr error
	)

	for cat := range categories {
		ictx := clues.Add(ctx, "base_service", service, "base_category", cat)

		p, err := path.Build(tenant, rOwner, service, cat, false, "tmp")
		if err != nil {
			// Shouldn't happen.
			err = clues.Wrap(err, "making path").WithClues(ictx)
			el.AddRecoverable(err)
			lastErr = err

			continue
		}

		p, err = p.Dir()
		if err != nil {
			// Shouldn't happen.
			err = clues.Wrap(err, "getting base prefix").WithClues(ictx)
			el.AddRecoverable(err)
			lastErr = err

			continue
		}

		// Pop off the last path element because we just want the prefix.
		res = append(res, emptyCollection{p: p, su: su})
	}

	return res, lastErr
}
