package site

import (
	"context"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type backupHandler interface {
	getItemByIDer
	getItemser
}

type getItemByIDer interface {
	GetItemByID(ctx context.Context, itemID string) (models.Listable, error)
}

type getItemser interface {
	GetItems(ctx context.Context, cc api.CallConfig) ([]models.Listable, error)
}
