package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler = &listsBackupHandler{}

type listsBackupHandler struct {
	ac                api.Lists
	protectedResource string
}

func NewListsBackupHandler(protectedResource string, ac api.Lists) listsBackupHandler {
	return listsBackupHandler{
		ac:                ac,
		protectedResource: protectedResource,
	}
}

func (bh listsBackupHandler) GetItemByID(ctx context.Context, itemID string) (models.Listable, error) {
	return bh.ac.GetListByID(ctx, bh.protectedResource, itemID)
}

func (bh listsBackupHandler) GetItems(ctx context.Context, cc api.CallConfig) ([]models.Listable, error) {
	return bh.ac.GetLists(ctx, bh.protectedResource, cc)
}

var _ restoreHandler = &listsRestoreHandler{}

type listsRestoreHandler struct {
	ac                api.Lists
	protectedResource string
}

func NewListsRestoreHandler(protectedResource string, ac api.Lists) listsRestoreHandler {
	return listsRestoreHandler{
		ac:                ac,
		protectedResource: protectedResource,
	}
}

func (rh listsRestoreHandler) PostList(
	ctx context.Context,
	listName string,
	storedListData []byte,
) (models.Listable, error) {
	return rh.ac.PostList(ctx, rh.protectedResource, listName, storedListData)
}

func (rh listsRestoreHandler) DeleteList(
	ctx context.Context,
	listID string,
) error {
	return rh.ac.DeleteList(ctx, rh.protectedResource, listID)
}
