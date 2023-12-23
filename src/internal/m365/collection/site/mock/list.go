package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type ListHandler struct {
	List models.Listable
	Err  error
}

func (lh *ListHandler) GetItemByID(
	ctx context.Context,
	itemID string,
) (models.Listable, *details.SharePointInfo, error) {
	ls := models.NewList()

	lh.List = ls
	lh.List.SetId(ptr.To(itemID))

	info := &details.SharePointInfo{
		ItemName: itemID,
	}

	return ls, info, lh.Err
}

type ListRestoreHandler struct {
	List models.Listable
	Err  error
}

func (lh *ListRestoreHandler) PostList(
	ctx context.Context,
	listName string,
	storedListBytes []byte,
) (models.Listable, error) {
	ls := models.NewList()

	lh.List = ls
	lh.List.SetDisplayName(ptr.To(listName))

	return lh.List, lh.Err
}
