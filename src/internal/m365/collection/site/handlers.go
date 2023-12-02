package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type backupHandler interface {
	getItemByIDer
}

type getItemByIDer interface {
	GetItemByID(ctx context.Context, itemID string) (models.Listable, error)
}
