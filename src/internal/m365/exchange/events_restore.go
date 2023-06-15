package exchange

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/dttm"
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
	return &eventContainerCache{
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
		// We cannot use `[]models.Attbachmentable{}` instead of nil
		// for beta endpoint.
		event.SetAttachments(nil)
	}

	item, err := h.ip.PostItem(ctx, userID, destinationID, event)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring calendar item")
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

	// Fix up event instances in case we have a recurring event
	err = updateRecurringEvents(ctx, h.ac, userID, destinationID, ptr.Val(item.GetId()), event)
	if err != nil {
		return nil, clues.Stack(err)
	}

	info := api.EventInfo(event)
	info.Size = int64(len(body))

	return info, nil
}

func updateRecurringEvents(
	ctx context.Context,
	ac api.Events,
	userID, containerID, eventID string,
	event models.Eventable,
) error {
	if event.GetRecurrence() == nil {
		return nil
	}

	// Cancellations and exceptions are currently in additional data
	// but will get their own fields once the beta API lands and
	// should be moved then
	cancelledOccurrences := event.GetAdditionalData()["cancelledOccurrences"]
	exceptionOccurrences := event.GetAdditionalData()["exceptionOccurrences"]

	err := updateCancelledOccurrences(ctx, ac, userID, eventID, cancelledOccurrences)
	if err != nil {
		return clues.Wrap(err, "update cancelled occurrences")
	}

	err = updateExceptionOccurrences(ctx, ac, userID, eventID, exceptionOccurrences)
	if err != nil {
		return clues.Wrap(err, "update exception occurrences")
	}

	return nil
}

// updateExceptionOccurrences take events that have exceptions, uses
// the originalStart date to find the instance and modify it to match
// the backup by updating the instance to match the backed up one
func updateExceptionOccurrences(
	ctx context.Context,
	ac api.Events,
	userID string,
	eventID string,
	exceptionOccurrences any,
) error {
	if exceptionOccurrences == nil {
		return nil
	}

	eo, ok := exceptionOccurrences.([]map[string]any)
	if !ok {
		return clues.New("converting exceptionOccurrences to []map[string]any").
			With("type", fmt.Sprintf("%T", exceptionOccurrences))
	}

	for _, inst := range eo {
		evt, err := api.EventFromMap(inst)
		if err != nil {
			return clues.Wrap(err, "parsing exception event")
		}

		start := ptr.Val(evt.GetOriginalStart())
		startStr := dttm.FormatTo(start, dttm.DateOnly)
		endStr := dttm.FormatTo(start.Add(24*time.Hour), dttm.DateOnly)

		evts, err := ac.GetItemInstances(ctx, userID, eventID, startStr, endStr)
		if err != nil {
			return clues.Wrap(err, "getting instances")
		}

		if len(evts) != 1 {
			return clues.New("invalid number of instances for modified").
				With("count", len(evts), "original_start", start)
		}

		evt = toEventSimplified(evt)

		// TODO(meain): Update attachments (might have to diff the
		// attachments using ids and delete or add). We will have
		// to get the id of the existing attachments, diff them
		// with what we need a then create/delete items kinda like
		// permissions
		_, err = ac.UpdateItem(ctx, userID, ptr.Val(evts[0].GetId()), evt)
		if err != nil {
			return clues.Wrap(err, "updating event instance")
		}
	}
	return nil
}

// updateCancelledOccurrences get the cancelled occurrences which is a
// list of strings of the format "<id>.<date>", parses the date out of
// that and uses the to get the event instance at that date to delete.
func updateCancelledOccurrences(
	ctx context.Context,
	ac api.Events,
	userID string,
	eventID string,
	cancelledOccurrences any,
) error {
	if cancelledOccurrences == nil {
		return nil
	}

	co, ok := cancelledOccurrences.([]*string)
	if !ok {
		return clues.New("converting cancelledOccurrences to []*string").
			With("type", fmt.Sprintf("%T", cancelledOccurrences))
	}

	// OPTIMIZATION: Group instances whose dates are close by
	for _, inst := range co {
		splits := strings.Split(ptr.Val(inst), ".")
		startStr := splits[len(splits)-1]

		start, err := dttm.ParseTime(startStr)
		if err != nil {
			return clues.Wrap(err, "parsing cancelled event date")
		}

		endStr := dttm.FormatTo(start.Add(24*time.Hour), dttm.DateOnly)

		evts, err := ac.GetItemInstances(ctx, userID, eventID, startStr, endStr)
		if err != nil {
			return clues.Wrap(err, "getting instances")
		}

		if len(evts) != 1 {
			return clues.New("invalid number of instances").
				With("count", len(evts))
		}

		err = ac.DeleteItem(ctx, userID, ptr.Val(evts[0].GetId()))
		if err != nil {
			return clues.Wrap(err, "deleting event instance")
		}
	}

	return nil
}
