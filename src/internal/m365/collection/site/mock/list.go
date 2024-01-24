package mock

import (
	"context"
	"slices"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ListHandler struct {
	protectedResource string
	lists             []models.Listable
	listsMap          map[string]models.Listable
	err               error
}

func NewListHandler(lists []models.Listable, protectedResource string, err error) ListHandler {
	lstMap := make(map[string]models.Listable)
	for _, lst := range lists {
		lstMap[ptr.Val(lst.GetId())] = lst
	}

	return ListHandler{
		protectedResource: protectedResource,
		lists:             lists,
		listsMap:          lstMap,
		err:               err,
	}
}

func (lh ListHandler) GetItemByID(
	ctx context.Context,
	itemID string,
) (models.Listable, *details.SharePointInfo, error) {
	lstInfo := &details.SharePointInfo{
		List: &details.ListInfo{
			Name: itemID,
		},
	}

	lst, ok := lh.listsMap[itemID]
	if ok {
		return lst, lstInfo, lh.err
	}

	listInfo := models.NewListInfo()
	listInfo.SetTemplate(ptr.To("genericList"))

	ls := models.NewList()
	ls.SetId(ptr.To(itemID))
	ls.SetList(listInfo)

	lh.listsMap[itemID] = ls

	return ls, lstInfo, lh.err
}

func (lh ListHandler) GetItems(
	context.Context,
	api.CallConfig,
) ([]models.Listable, error) {
	return lh.lists, lh.err
}

func (lh ListHandler) CanonicalPath(
	storageDirFolders path.Elements,
	tenantID string,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			tenantID,
			lh.protectedResource,
			path.SharePointService,
			path.ListsCategory,
			false)
}

func (lh *ListHandler) Check(t *testing.T, expected []string) {
	listIDs := maps.Keys(lh.listsMap)

	slices.Sort(listIDs)
	slices.Sort(expected)

	assert.Equal(t, expected, listIDs, "expected calls")
}

type ListRestoreHandler struct {
	deleteListErr error
	postListErr   error
	patchListErr  error
}

func NewListRestoreHandler(deleteListErr error, postListErr, patchListErr error) *ListRestoreHandler {
	return &ListRestoreHandler{
		deleteListErr: deleteListErr,
		postListErr:   postListErr,
		patchListErr:  patchListErr,
	}
}

func (lh *ListRestoreHandler) PostList(
	ctx context.Context,
	listName string,
	storedList models.Listable,
	_ *fault.Bus,
) (models.Listable, error) {
	ls, _ := api.ToListable(storedList, listName)
	return ls, lh.postListErr
}

func (lh *ListRestoreHandler) PatchList(
	ctx context.Context,
	listID string,
	list models.Listable,
) (models.Listable, error) {
	return nil, lh.patchListErr
}

func (lh *ListRestoreHandler) GetList(
	ctx context.Context,
	listID string,
) (models.Listable, *details.SharePointInfo, error) {
	ls := models.NewList()
	ls.SetId(ptr.To(listID))

	return ls, api.ListToSPInfo(ls), nil
}

func (lh *ListRestoreHandler) DeleteList(_ context.Context, _ string) error {
	return lh.deleteListErr
}

func (lh *ListRestoreHandler) GetListsByCollisionKey(ctx context.Context) (map[string]string, error) {
	return map[string]string{}, nil
}

func StubLists(ids ...string) []models.Listable {
	lists := make([]models.Listable, 0, len(ids))

	for _, id := range ids {
		listInfo := models.NewListInfo()
		listInfo.SetTemplate(ptr.To("genericList"))

		lst := models.NewList()
		lst.SetDisplayName(ptr.To(id))
		lst.SetId(ptr.To(id))
		lst.SetList(listInfo)

		lists = append(lists, lst)
	}

	return lists
}
