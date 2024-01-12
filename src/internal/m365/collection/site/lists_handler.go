package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

func (bh listsBackupHandler) CanonicalPath(
	storageDirFolders path.Elements,
	tenantID string,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			tenantID,
			bh.protectedResource,
			path.SharePointService,
			path.ListsCategory,
			false)
}

func (bh listsBackupHandler) GetItemByID(
	ctx context.Context,
	itemID string,
) (models.Listable, *details.SharePointInfo, error) {
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
	storedList models.Listable,
	errs *fault.Bus,
) (models.Listable, error) {
	return rh.ac.PostList(ctx, rh.protectedResource, listName, storedList, errs)
}

func (rh listsRestoreHandler) DeleteList(
	ctx context.Context,
	listID string,
) error {
	return rh.ac.DeleteList(ctx, rh.protectedResource, listID)
}

func (rh listsRestoreHandler) GetList(
	ctx context.Context,
	listID string,
) (models.Listable, *details.SharePointInfo, error) {
	return rh.ac.GetListByID(ctx, rh.protectedResource, listID)
}
