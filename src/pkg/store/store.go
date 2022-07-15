package store

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
)

var _ Storer = &kopia.ModelStore{}

type (
	deleter interface {
		Delete(ctx context.Context, s model.Schema, id model.ID) error
		DeleteWithModelStoreID(ctx context.Context, id manifest.ID) error
	}
	getter interface {
		Get(ctx context.Context, s model.Schema, id model.ID, data model.Model) error
		GetIDsForType(ctx context.Context, s model.Schema, tags map[string]string) ([]*model.BaseModel, error)
		GetWithModelStoreID(ctx context.Context, s model.Schema, id manifest.ID, data model.Model) error
	}
	updater interface {
		Put(ctx context.Context, s model.Schema, m model.Model) error
		Update(ctx context.Context, s model.Schema, m model.Model) error
	}
	Storer interface {
		getter
		updater
		deleter
	}
)

type Wrapper struct {
	Storer
}

func NewKopiaStore(kMS *kopia.ModelStore) *Wrapper {
	return &Wrapper{kMS}
}
