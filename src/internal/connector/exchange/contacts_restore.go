package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var (
	_ containerCreator = &contactRestoreHandler{}
	_ itemRestorer     = &contactRestoreHandler{}
)

type contactRestoreHandler struct {
	ac     api.Contacts
	ip     itemPoster[models.Contactable]
	userID string
}

func newContactRestoreHandler(
	ac api.Client,
	userID string,
) contactRestoreHandler {
	return contactRestoreHandler{
		ac:     ac.Contacts(),
		ip:     ac.Contacts(),
		userID: userID,
	}
}

func (h contactRestoreHandler) newContainerCache() graph.ContainerResolver {
	return &contactFolderCache{
		userID: h.userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h contactRestoreHandler) CreateContainer(
	ctx context.Context,
	userID, containerName, parentContainerID string,
) (graph.Container, error) {
	return h.ac.CreateContainer(ctx, userID, containerName, parentContainerID)
}

func (h contactRestoreHandler) CanGetContainerByName() bool {
	return false
}

func (h contactRestoreHandler) GetContainerByName(
	ctx context.Context,
	userID, parentContainerID string,
) (graph.Container, error) {
	return nil, clues.New("not supported yet")
}

func (h contactRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	destinationID string,
	errs *fault.Bus,
) (*details.ExchangeInfo, error) {
	contact, err := api.BytesToContactable(body)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact from bytes")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))

	item, err := h.ip.PostItem(ctx, h.userID, destinationID, contact)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring mail message")
	}

	info := api.ContactInfo(item)
	info.Size = int64(len(body))

	return info, nil
}
