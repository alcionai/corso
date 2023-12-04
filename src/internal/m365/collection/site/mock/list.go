package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

type ListHandler struct {
	ListItem models.Listable
	Err      error
}

func (lh *ListHandler) GetItemByID(ctx context.Context, itemID string) (models.Listable, error) {
	ls := models.NewList()

	lh.ListItem = ls
	lh.ListItem.SetId(ptr.To(itemID))

	return ls, lh.Err
}
