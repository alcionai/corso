package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ itemRestorer = &contactRestoreHandler{}

type contactRestoreHandler struct {
	ac api.Contacts
	ip itemPoster[models.Contactable]
}

func newContactRestoreHandler(
	ac api.Client,
) contactRestoreHandler {
	return contactRestoreHandler{
		ac: ac.Contacts(),
		ip: ac.Contacts(),
	}
}

func (h contactRestoreHandler) newContainerCache(userID string) graph.ContainerResolver {
	return &contactFolderCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h contactRestoreHandler) formatRestoreDestination(
	destinationContainerName string,
	_ path.Path, // contact folders cannot be nested
) *path.Builder {
	return path.Builder{}.Append(destinationContainerName)
}

func (h contactRestoreHandler) CreateContainer(
	ctx context.Context,
	userID, containerName, _ string, // parent container not used
) (graph.Container, error) {
	return h.ac.CreateContainer(ctx, userID, containerName, "")
}

func (h contactRestoreHandler) containerSearcher() containerByNamer {
	return nil
}

// always returns the provided value
func (h contactRestoreHandler) orRootContainer(c string) string {
	return c
}

func (h contactRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	userID, destinationID string,
	errs *fault.Bus,
) (*details.ExchangeInfo, error) {
	contact, err := api.BytesToContactable(body)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact from bytes")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))

	item, err := h.ip.PostItem(ctx, userID, destinationID, contact)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring mail message")
	}

	info := api.ContactInfo(item)
	info.Size = int64(len(body))

	return info, nil
}
