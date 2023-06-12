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

var _ itemRestorer = &eventRestoreHandler{}

type eventRestoreHandler struct {
	ac api.Events
	ip itemPoster[models.Eventable]
}

func newEventRestoreHandler(
	ac api.Client,
) eventRestoreHandler {
	ace := ac.Events()

	return eventRestoreHandler{
		ac: ace,
		ip: ace,
	}
}

func (h eventRestoreHandler) newContainerCache(userID string) graph.ContainerResolver {
	return &eventCalendarCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h eventRestoreHandler) formatRestoreDestination(
	destinationContainerName string,
	_ path.Path, // ignored because calendars cannot be nested
) *path.Builder {
	return path.Builder{}.Append(destinationContainerName)
}

func (h eventRestoreHandler) CreateContainer(
	ctx context.Context,
	userID, containerName, _ string, // parent container not used
) (graph.Container, error) {
	return h.ac.CreateContainer(ctx, userID, containerName, "")
}

func (h eventRestoreHandler) containerSearcher() containerByNamer {
	return h.ac
}

// always returns the provided value
func (h eventRestoreHandler) orRootContainer(c string) string {
	return c
}

func (h eventRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	userID, destinationID string,
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

	item, err := h.ip.PostItem(ctx, userID, destinationID, event)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring mail message")
	}

	err = uploadAttachments(
		ctx,
		h.ac,
		attachments,
		userID,
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
