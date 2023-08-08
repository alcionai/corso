package mock

import (
	"context"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type RestoreCollection struct {
	data.Collection
	AuxItems map[string]data.Stream
}

func (rc RestoreCollection) FetchItemByName(
	ctx context.Context,
	name string,
) (data.Stream, error) {
	res := rc.AuxItems[name]
	if res == nil {
		return nil, data.ErrNotFound
	}

	return res, nil
}

type BackupCollection struct {
	Path    path.Path
	Loc     *path.Builder
	Streams []data.Stream
	CState  data.CollectionState
}

func (c *BackupCollection) Items(context.Context, *fault.Bus) <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for _, s := range c.Streams {
			res <- s
		}
	}()

	return res
}

func (c BackupCollection) FullPath() path.Path {
	return c.Path
}

func (c BackupCollection) PreviousPath() path.Path {
	return c.Path
}

func (c BackupCollection) LocationPath() *path.Builder {
	return c.Loc
}

func (c BackupCollection) State() data.CollectionState {
	return c.CState
}

func (c BackupCollection) DoNotMergeItems() bool {
	return false
}
