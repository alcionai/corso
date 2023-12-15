package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type ListHandler struct {
	ListItem models.Listable
	Err      error
}

func (lh *ListHandler) GetItemByID(
	ctx context.Context,
	itemID string,
) (models.Listable, *details.SharePointInfo, error) {
	ls := models.NewList()

	lh.ListItem = ls
	lh.ListItem.SetId(ptr.To(itemID))

	info := &details.SharePointInfo{
		ItemName: itemID,
	}

	return ls, info, lh.Err
}
