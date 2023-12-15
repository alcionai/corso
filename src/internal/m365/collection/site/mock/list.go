package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

type ListHandler struct {
	List models.Listable
	Err  error
}

func (lh *ListHandler) GetItemByID(ctx context.Context, itemID string) (models.Listable, error) {
	ls := models.NewList()

	lh.List = ls
	lh.List.SetId(ptr.To(itemID))

	return ls, lh.Err
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
