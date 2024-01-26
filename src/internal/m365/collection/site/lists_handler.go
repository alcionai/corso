package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var _ backupHandler = &listsBackupHandler{}

type listsBackupHandler struct {
	ac api.Lists
	qp graph.QueryParams
}

func NewListsBackupHandler(
	qp graph.QueryParams,
	ac api.Lists,
) listsBackupHandler {
	return listsBackupHandler{
		ac: ac,
		qp: qp,
	}
}

func (bh listsBackupHandler) CanonicalPath(
	storageDirFolders path.Elements,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			bh.qp.TenantID,
			bh.qp.ProtectedResource.ID(),
			path.SharePointService,
			path.ListsCategory,
			false)
}

func (bh listsBackupHandler) GetItemByID(
	ctx context.Context,
	itemID string,
) (models.Listable, *details.SharePointInfo, error) {
	return bh.ac.GetListByID(ctx, bh.qp.ProtectedResource.ID(), itemID)
}

func (bh listsBackupHandler) GetItems(ctx context.Context, cc api.CallConfig) ([]models.Listable, error) {
	return bh.ac.GetLists(ctx, bh.qp.ProtectedResource.ID(), cc)
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
