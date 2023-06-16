package exchange

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ itemRestorer = &eventRestoreHandler{}

type eventRestoreHandler struct {
	ac api.Events
}

func newEventRestoreHandler(
	ac api.Client,
) eventRestoreHandler {
	return eventRestoreHandler{
		ac: ac.Events(),
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
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	errs *fault.Bus,
) (*details.ExchangeInfo, error) {
	return restoreEvent(
		ctx,
		h.ac,
		body,
		userID, destinationID,
		collisionKeyToItemID,
		collisionPolicy,
		errs)
}

type eventRestorer interface {
	postItemer[models.Eventable]
	eventInstanceAndAttachmenter
}

func restoreEvent(
	ctx context.Context,
	er eventRestorer,
	body []byte,
	userID, destinationID string,
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	errs *fault.Bus,
) (*details.ExchangeInfo, error) {
	event, err := api.BytesToEventable(body)
	if err != nil {
		return nil, clues.Wrap(err, "creating event from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(event.GetId()))

	var (
		collisionKey = api.EventCollisionKey(event)
		collisionID  string
		replace      bool
	)

	if id, ok := collisionKeyToItemID[collisionKey]; ok {
		log := logger.Ctx(ctx).With("collision_key", clues.Hide(collisionKey))
		log.Debug("item collision")

		if collisionPolicy == control.Skip {
			log.Debug("skipping item with collision")
			return nil, graph.ErrItemAlreadyExistsConflict
		}

		collisionID = id
		replace = collisionPolicy == control.Replace
	}

	event = toEventSimplified(event)

	var attachments []models.Attachmentable

	if ptr.Val(event.GetHasAttachments()) {
		attachments = event.GetAttachments()
		// We cannot use `[]models.Attbachmentable{}` instead of nil
		// for beta endpoint.
		event.SetAttachments(nil)
	}

	item, err := er.PostItem(ctx, userID, destinationID, event)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring event")
	}

	// events have no PUT request, and PATCH could retain data that's not
	// associated with the backup item state.  Instead of updating, we
	// post first, then delete.  In case of failure between the two calls,
	// at least we'll have accidentally over-produced data instead of deleting
	// the user's data.
	if replace {
		if err := er.DeleteItem(ctx, userID, collisionID); err != nil {
			return nil, graph.Wrap(ctx, err, "deleting colliding event")
		}
	}

	err = uploadAttachments(
		ctx,
		er,
		attachments,
		userID,
		destinationID,
		ptr.Val(item.GetId()),
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	// Have to parse event again as we modified the original event and
	// removed cancelled and exceptions events form it
	event, err = api.BytesToEventable(body)
	if err != nil {
		return nil, clues.Wrap(err, "creating event from bytes").WithClues(ctx)
	}

	// Fix up event instances in case we have a recurring event
	err = updateRecurringEvents(
		ctx,
		er,
		userID,
		destinationID,
		ptr.Val(item.GetId()),
		event,
		errs,
	)
	if err != nil {
		return nil, clues.Stack(err)
	}

	info := api.EventInfo(event)
	info.Size = int64(len(body))

	return info, nil
}

func updateRecurringEvents(
	ctx context.Context,
	eiaa eventInstanceAndAttachmenter,
	userID, containerID, itemID string,
	event models.Eventable,
	errs *fault.Bus,
) error {
	if event.GetRecurrence() == nil {
		return nil
	}

	// Cancellations and exceptions are currently in additional data
	// but will get their own fields once the beta API lands and
	// should be moved then
	cancelledOccurrences := event.GetAdditionalData()["cancelledOccurrences"]
	exceptionOccurrences := event.GetAdditionalData()["exceptionOccurrences"]

	err := updateCancelledOccurrences(ctx, eiaa, userID, itemID, cancelledOccurrences)
	if err != nil {
		return clues.Wrap(err, "update cancelled occurrences")
	}

	err = updateExceptionOccurrences(ctx, eiaa, userID, containerID, itemID, exceptionOccurrences, errs)
	if err != nil {
		return clues.Wrap(err, "update exception occurrences")
	}

	return nil
}

type eventInstanceAndAttachmenter interface {
	attachmentGetDeletePoster
	DeleteItem(
		ctx context.Context,
		userID, itemID string,
	) error
	GetItemInstances(
		ctx context.Context,
		userID, itemID string,
		startDate, endDate string,
	) ([]models.Eventable, error)
	PatchItem(
		ctx context.Context,
		userID, eventID string,
		body models.Eventable,
	) (models.Eventable, error)
}

// updateExceptionOccurrences take events that have exceptions, uses
// the originalStart date to find the instance and modify it to match
// the backup by updating the instance to match the backed up one
func updateExceptionOccurrences(
	ctx context.Context,
	eiaa eventInstanceAndAttachmenter,
	userID string,
	containerID string,
	itemID string,
	exceptionOccurrences any,
	errs *fault.Bus,
) error {
	if exceptionOccurrences == nil {
		return nil
	}

	eo, ok := exceptionOccurrences.([]any)
	if !ok {
		return clues.New("converting exceptionOccurrences to []any").
			With("type", fmt.Sprintf("%T", exceptionOccurrences))
	}

	for _, instance := range eo {
		instance, ok := instance.(map[string]any)
		if !ok {
			return clues.New("converting instance to map[string]any").
				With("type", fmt.Sprintf("%T", instance))
		}

		evt, err := api.EventFromMap(instance)
		if err != nil {
			return clues.Wrap(err, "parsing exception event")
		}

		start := ptr.Val(evt.GetOriginalStart())
		startStr := dttm.FormatTo(start, dttm.DateOnly)
		endStr := dttm.FormatTo(start.Add(24*time.Hour), dttm.DateOnly)

		ictx := clues.Add(ctx, "event_instance_id", ptr.Val(evt.GetId()), "event_instance_date", start)

		// Get all instances on the day of the instance which should
		// just the one we need to modify
		instances, err := eiaa.GetItemInstances(ictx, userID, itemID, startStr, endStr)
		if err != nil {
			return clues.Wrap(err, "getting instances")
		}

		// Since the min recurrence interval is 1 day and we are
		// querying for only a single day worth of instances, we
		// should not have more than one instance here.
		if len(instances) != 1 {
			return clues.New("invalid number of instances for modified").
				With("instances_count", len(instances), "search_start", startStr, "search_end", endStr)
		}

		evt = toEventSimplified(evt)

		_, err = eiaa.PatchItem(ictx, userID, ptr.Val(instances[0].GetId()), evt)
		if err != nil {
			return clues.Wrap(err, "updating event instance")
		}

		// We are creating event again from map as `toEventSimplified`
		// removed the attachments and creating a clone from start of
		// the event is non-trivial
		evt, err = api.EventFromMap(instance)
		if err != nil {
			return clues.Wrap(err, "parsing event instance")
		}

		err = updateAttachments(
			ictx,
			eiaa,
			userID,
			containerID,
			ptr.Val(instances[0].GetId()),
			evt,
			errs)
		if err != nil {
			return clues.Wrap(err, "updating event instance attachments")
		}
	}

	return nil
}

type attachmentGetDeletePoster interface {
	attachmentPoster
	GetAttachments(
		ctx context.Context,
		immutableIDs bool,
		userID string,
		itemID string,
	) ([]models.Attachmentable, error)
	DeleteAttachment(
		ctx context.Context,
		userID, calendarID, eventID, attachmentID string,
	) error
}

// updateAttachments updates the attachments of an event to match what
// is present in the backed up event. Ideally we could make use of the
// id of the series master event's attachments to see if we had
// added/removed any attachments, but as soon an event is modified,
// the id changes which makes the ids unusable. In this function, we
// use the name and content bytes to detect the changes. This function
// can be used to update the attachments of any event irrespective of
// whether they are event instances of a series master although for
// newer event, since we probably won't already have any events it
// would be better use Post[Small|Large]Attachment.
func updateAttachments(
	ctx context.Context,
	agdp attachmentGetDeletePoster,
	userID, containerID, eventID string,
	event models.Eventable,
	errs *fault.Bus,
) error {
	el := errs.Local()

	attachments, err := agdp.GetAttachments(ctx, false, userID, eventID)
	if err != nil {
		return clues.Wrap(err, "getting attachments")
	}

	// Delete attachments that are not present in the backup but are
	// present in the event(ones that were automatically inherited
	// from series master).
	for _, att := range attachments {
		if el.Failure() != nil {
			return el.Failure()
		}

		name := ptr.Val(att.GetName())
		id := ptr.Val(att.GetId())

		content, err := api.GetAttachmentContent(att)
		if err != nil {
			return clues.Wrap(err, "getting attachment").With("attachment_id", id)
		}

		found := false

		for _, nAtt := range event.GetAttachments() {
			nName := ptr.Val(nAtt.GetName())

			nContent, err := api.GetAttachmentContent(nAtt)
			if err != nil {
				return clues.Wrap(err, "getting attachment").With("attachment_id", ptr.Val(nAtt.GetId()))
			}

			if name == nName && bytes.Equal(content, nContent) {
				found = true
				break
			}
		}

		if !found {
			err = agdp.DeleteAttachment(ctx, userID, containerID, eventID, id)
			if err != nil {
				logger.CtxErr(ctx, err).With("attachment_name", name).Info("attachment delete failed")
				el.AddRecoverable(ctx, clues.Wrap(err, "deleting event attachment").
					WithClues(ctx).With("attachment_name", name))
			}
		}
	}

	// Upload missing(attachments that are present in the individual
	// instance but not in the series master event) attachments
	for _, att := range event.GetAttachments() {
		name := ptr.Val(att.GetName())
		id := ptr.Val(att.GetId())

		content, err := api.GetAttachmentContent(att)
		if err != nil {
			return clues.Wrap(err, "getting attachment").With("attachment_id", id)
		}

		found := false

		for _, nAtt := range attachments {
			nName := ptr.Val(nAtt.GetName())

			bContent, err := api.GetAttachmentContent(nAtt)
			if err != nil {
				return clues.Wrap(err, "getting attachment").With("attachment_id", ptr.Val(nAtt.GetId()))
			}

			// Max size allowed for an outlook attachment is 150MB
			if name == nName && bytes.Equal(content, bContent) {
				found = true
				break
			}
		}

		if !found {
			err = uploadAttachment(ctx, agdp, userID, containerID, eventID, att)
			if err != nil {
				return clues.Wrap(err, "uploading attachment").
					With("attachment_id", id)
			}
		}
	}

	return el.Failure()
}

// updateCancelledOccurrences get the cancelled occurrences which is a
// list of strings of the format "<id>.<date>", parses the date out of
// that and uses the to get the event instance at that date to delete.
func updateCancelledOccurrences(
	ctx context.Context,
	eiaa eventInstanceAndAttachmenter,
	userID string,
	itemID string,
	cancelledOccurrences any,
) error {
	if cancelledOccurrences == nil {
		return nil
	}

	co, ok := cancelledOccurrences.([]any)
	if !ok {
		return clues.New("converting cancelledOccurrences to []any").
			With("type", fmt.Sprintf("%T", cancelledOccurrences))
	}

	// OPTIMIZATION: We can fetch a date range instead of fetching
	// instances if we have multiple cancelled events which are nearby
	// and reduce the number of API calls that we have to make
	for _, instance := range co {
		instance, err := str.AnyToString(instance)
		if err != nil {
			return err
		}

		splits := strings.Split(instance, ".")

		startStr := splits[len(splits)-1]

		start, err := dttm.ParseTime(startStr)
		if err != nil {
			return clues.Wrap(err, "parsing cancelled event date")
		}

		endStr := dttm.FormatTo(start.Add(24*time.Hour), dttm.DateOnly)

		// Get all instances on the day of the instance which should
		// just the one we need to modify
		instances, err := eiaa.GetItemInstances(ctx, userID, itemID, startStr, endStr)
		if err != nil {
			return clues.Wrap(err, "getting instances")
		}

		// Since the min recurrence interval is 1 day and we are
		// querying for only a single day worth of instances, we
		// should not have more than one instance here.
		if len(instances) != 1 {
			return clues.New("invalid number of instances for cancelled").
				With("instances_count", len(instances), "search_start", startStr, "search_end", endStr)
		}

		err = eiaa.DeleteItem(ctx, userID, ptr.Val(instances[0].GetId()))
		if err != nil {
			return clues.Wrap(err, "deleting event instance")
		}
	}

	return nil
}

func (h eventRestoreHandler) getItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	m, err := h.ac.GetItemsInContainerByCollisionKey(ctx, userID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}
