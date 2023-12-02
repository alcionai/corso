package mock

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

type GetList struct {
	Err error
}

func (m GetList) GetItemByID(ctx context.Context, itemID string) (models.Listable, error) {
	lst := models.NewList()
	lst.SetId(ptr.To(itemID))

	return lst, m.Err
}
