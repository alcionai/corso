package mock

import (
	"context"

	"github.com/alcionai/corso/src/internal/data"
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
