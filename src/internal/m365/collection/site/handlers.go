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

type restoreHandler interface {
	PostLister
	PostListItemer
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
