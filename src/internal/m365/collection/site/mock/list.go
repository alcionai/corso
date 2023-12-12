package mock

import (
	"context"
	"errors"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
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

type ListRestoreHandler struct {
	Err error
}

func (lh *ListRestoreHandler) PostList(
	ctx context.Context,
	listName string,
	oldListByteArray []byte,
) (models.Listable, error) {
	newListName := listName

	oldList, err := betaAPI.CreateListFromBytes(oldListByteArray)
	if err != nil {
		return nil, errors.New("error while creating old list")
	}

	if name, ok := ptr.ValOK(oldList.GetDisplayName()); ok {
		nameParts := strings.Split(listName, "_")
		if len(nameParts) > 0 {
			nameParts[len(nameParts)-1] = name
			newListName = strings.Join(nameParts, "_")
		}
	}

	newList := betaAPI.ToListable(oldList, newListName)

	return newList, lh.Err
}

func (lh *ListRestoreHandler) PostListItem(
	ctx context.Context,
	listID string,
	oldListByteArray []byte,
) ([]models.ListItemable, error) {
	oldList, err := betaAPI.CreateListFromBytes(oldListByteArray)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "creating old list to get list items")
	}

	contents := make([]models.ListItemable, 0)

	for _, itm := range oldList.GetItems() {
		temp := betaAPI.CloneListItem(itm)
		contents = append(contents, temp)
	}

	return contents, lh.Err
}
