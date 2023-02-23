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

func (c emptyCollection) Items(ctx context.Context, errs *fault.Errors) <-chan data.Stream {
	res := make(chan data.Stream)
	close(res)

	s := support.CreateStatus(ctx, support.Backup, 0, support.CollectionMetrics{}, nil, "")
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
	tenant, user string,
	service path.ServiceType,
	categories map[path.CategoryType]struct{},
	su support.StatusUpdater,
) ([]data.BackupCollection, error) {
	var (
		errs = []error{}
		res  = []data.BackupCollection{}
	)

	for cat := range categories {
		p, err := path.Builder{}.Append("tmp").ToDataLayerPath(tenant, user, service, cat, false)
		if err != nil {
			// Shouldn't happen.
			errs = append(
				errs,
				clues.Wrap(err, "making path").With("service", service, "category", cat))

			continue
		}

		p, err = p.Dir()
		if err != nil {
			// Shouldn't happen.
			errs = append(
				errs,
				clues.Wrap(err, "getting base prefix").With("serivce", service, "category", cat))

			continue
		}

		// Pop off the last path element because we just want the prefix.
		res = append(res, emptyCollection{p: p, su: su})
	}

	return res, nil
}
