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
	_ containerCreator = &eventRestoreHandler{}
	_ itemRestorer     = &eventRestoreHandler{}
)

type eventRestoreHandler struct {
	ac     api.Events
	ip     itemPoster[models.Eventable]
	ap     attachmentPoster
	userID string
}

func newEventRestoreHandler(
	ac api.Client,
	userID string,
) eventRestoreHandler {
	ace := ac.Events()

	return eventRestoreHandler{
		ac:     ace,
		ip:     ace,
		ap:     ace,
		userID: userID,
	}
}

func (h eventRestoreHandler) newContainerCache() graph.ContainerResolver {
	return &eventCalendarCache{
		userID: h.userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h eventRestoreHandler) CreateContainer(
	ctx context.Context,
	userID, containerName, parentContainerID string,
) (graph.Container, error) {
	return h.ac.CreateContainer(ctx, userID, containerName, parentContainerID)
}

func (h eventRestoreHandler) CanGetContainerByName() bool {
	return true
}

func (h eventRestoreHandler) GetContainerByName(
	ctx context.Context,
	userID, containerName string,
) (graph.Container, error) {
	return h.ac.GetContainerByName(ctx, userID, containerName)
}

func (h eventRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	destinationID string,
	errs *fault.Bus,
) (*details.ExchangeInfo, error) {
	event, err := api.BytesToEventable(body)
	if err != nil {
		return nil, clues.Wrap(err, "creating event from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(event.GetId()))

	event = toEventSimplified(event)

	var attachments []models.Attachmentable

	if ptr.Val(event.GetHasAttachments()) {
		attachments = event.GetAttachments()
		event.SetAttachments([]models.Attachmentable{})
	}

	item, err := h.ip.PostItem(ctx, h.userID, destinationID, event)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring mail message")
	}

	err = uploadAttachments(
		ctx,
		h.ap,
		attachments,
		h.userID,
		destinationID,
		ptr.Val(item.GetId()),
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	info := api.EventInfo(event)
	info.Size = int64(len(body))

	return info, nil
}
