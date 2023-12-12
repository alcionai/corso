package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
