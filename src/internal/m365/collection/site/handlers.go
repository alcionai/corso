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

type restoreHandler interface {
	PostLister
	PostListItemer
	DeleteLister
}

type PostLister interface {
	PostList(
		ctx context.Context,
		listName string,
		oldListByteArray []byte,
	) (models.Listable, error)
}

type PostListItemer interface {
	PostListItem(
		ctx context.Context,
		listID string,
		oldListByteArray []byte,
	) ([]models.ListItemable, error)
}

type DeleteLister interface {
	DeleteList(
		ctx context.Context,
		listID string,
	) error
}
