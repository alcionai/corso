package store

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
)

var _ Storer = &kopia.ModelStore{}

type (
	Storer interface {
		Delete(ctx context.Context, s model.Schema, id model.StableID) error
		DeleteWithModelStoreID(ctx context.Context, id manifest.ID) error
		Get(ctx context.Context, s model.Schema, id model.StableID, data model.Model) error
		GetIDsForType(ctx context.Context, s model.Schema, tags map[string]string) ([]*model.BaseModel, error)
		GetWithModelStoreID(ctx context.Context, s model.Schema, id manifest.ID, data model.Model) error
		Put(ctx context.Context, s model.Schema, m model.Model) error
		Update(ctx context.Context, s model.Schema, m model.Model) error
	}
)

type Wrapper struct {
	Storer
}

func NewKopiaStore(kms *kopia.ModelStore) *Wrapper {
	return &Wrapper{kms}
}
