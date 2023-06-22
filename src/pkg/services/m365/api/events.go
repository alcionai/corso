package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Events() Events {
	return Events{c}
}

// Events is an interface-compliant provider of the client.
type Events struct {
	Client
}

// ---------------------------------------------------------------------------
// containers
// ---------------------------------------------------------------------------

// CreateContainer makes an event Calendar with the name in the user's M365 exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-calendars?view=graph-rest-1.0&tabs=go
func (c Events) CreateContainer(
	ctx context.Context,
	userID, containerName string,
	_ string, // parentContainerID needed for iface, doesn't apply to contacts
) (graph.Container, error) {
	body := models.NewCalendar()
	body.SetName(&containerName)

	container, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating calendar")
	}

	return CalendarDisplayable{Calendarable: container}, nil
}

// DeleteContainer removes a calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func (c Events) DeleteContainer(
	ctx context.Context,
	userID, containerID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

// prefer GetContainerByID where possible.
// use this only in cases where the models.Calendarable
// is required.
func (c Events) GetCalendar(
	ctx context.Context,
	userID, containerID string,
) (models.Calendarable, error) {
	config := &users.ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemCalendarsCalendarItemRequestBuilderGetQueryParameters{
			Select: idAnd("name", "owner"),
		},
	}

	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Get(ctx, config)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// interface-compliant wrapper of GetCalendar
func (c Events) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	cal, err := c.GetCalendar(ctx, userID, containerID)
	if err != nil {
		return nil, err
	}

	return graph.CalendarDisplayable{Calendarable: cal}, nil
}

// GetContainerByName fetches a calendar by name
func (c Events) GetContainerByName(
	ctx context.Context,
	userID, containerName string,
) (graph.Container, error) {
	filter := fmt.Sprintf("name eq '%s'", containerName)
	options := &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemCalendarsRequestBuilderGetQueryParameters{
			Filter: &filter,
		},
	}

	ctx = clues.Add(ctx, "calendar_name", containerName)

	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		Get(ctx, options)
	if err != nil {
		return nil, graph.Stack(ctx, err).WithClues(ctx)
	}

	// We only allow the api to match one calendar with provided name.
	// Return an error if multiple calendars exist (unlikely) or if no calendar
	// is found.
	if len(resp.GetValue()) != 1 {
		err = clues.New("unexpected number of calendars returned").
			With("returned_calendar_count", len(resp.GetValue()))
		return nil, err
	}

	// Sanity check ID and name
	cal := resp.GetValue()[0]
	cd := CalendarDisplayable{Calendarable: cal}

	if err := graph.CheckIDAndName(cd); err != nil {
		return nil, err
	}

	return graph.CalendarDisplayable{Calendarable: cal}, nil
}

func (c Events) PatchCalendar(
	ctx context.Context,
	userID, containerID string,
	body models.Calendarable,
) error {
	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Patch(ctx, body, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "patching event calendar")
	}

	return nil
}

const (
	// Beta version cannot have /calendars/%s for get and Patch
	// https://stackoverflow.com/questions/50492177/microsoft-graph-get-user-calendar-event-with-beta-version
	eventExceptionsBetaURLTemplate = "https://graph.microsoft.com/beta/users/%s/events/%s?$expand=exceptionOccurrences"
	eventPostBetaURLTemplate       = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events"
	eventPatchBetaURLTemplate      = "https://graph.microsoft.com/beta/users/%s/events/%s"
)

// ---------------------------------------------------------------------------
// items
// ---------------------------------------------------------------------------

// GetItem retrieves an Eventable item.
func (c Events) GetItem(
	ctx context.Context,
	userID, itemID string,
	immutableIDs bool,
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		err    error
		event  models.Eventable
		config = &users.ItemEventsEventItemRequestBuilderGetRequestConfiguration{
			Headers: newPreferHeaders(preferImmutableIDs(immutableIDs)),
		}
	)

	// Beta endpoint helps us fetch the event exceptions, but since we
	// don't use the beta SDK, the exceptionOccurrences and
	// cancelledOccurrences end up in AdditionalData
	// https://learn.microsoft.com/en-us/graph/api/resources/event?view=graph-rest-beta#properties
	rawURL := fmt.Sprintf(eventExceptionsBetaURLTemplate, userID, itemID)
	builder := users.NewItemEventsEventItemRequestBuilder(rawURL, c.Stable.Adapter())

	event, err = builder.Get(ctx, config)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	err = validateCancelledOccurrences(event)
	if err != nil {
		return nil, nil, clues.Wrap(err, "verify cancelled occurrences")
	}

	err = fixupExceptionOccurrences(ctx, c, event, immutableIDs, userID)
	if err != nil {
		return nil, nil, clues.Wrap(err, "fixup exception occurrences")
	}

	var attachments []models.Attachmentable
	if ptr.Val(event.GetHasAttachments()) || HasAttachments(event.GetBody()) {
		attachments, err = c.GetAttachments(ctx, immutableIDs, userID, itemID)
		if err != nil {
			return nil, nil, err
		}
	}

	event.SetAttachments(attachments)

	return event, EventInfo(event), nil
}

// fixupExceptionOccurrences gets attachments and converts the data
// into a format that gets serialized when storing to kopia
func fixupExceptionOccurrences(
	ctx context.Context,
	client Events,
	event models.Eventable,
	immutableIDs bool,
	userID string,
) error {
	// Fetch attachments for exceptions
	exceptionOccurrences := event.GetAdditionalData()["exceptionOccurrences"]
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

		evt, err := EventFromMap(instance)
		if err != nil {
			return clues.Wrap(err, "parsing exception event")
		}

		// OPTIMIZATION: We don't have to store any of the
		// attachments that carry over from the original

		var attachments []models.Attachmentable
		if ptr.Val(event.GetHasAttachments()) || HasAttachments(event.GetBody()) {
			attachments, err = client.GetAttachments(ctx, immutableIDs, userID, ptr.Val(evt.GetId()))
			if err != nil {
				return clues.Wrap(err, "getting event instance attachments").
					With("event_instance_id", ptr.Val(evt.GetId()))
			}
		}

		// This odd roundabout way of doing this is required as
		// the json serialization at the end does not serialize if
		// you just pass in a models.Attachmentable
		convertedAttachments := []map[string]interface{}{}

		for _, attachment := range attachments {
			am, err := parseableToMap(attachment)
			if err != nil {
				return clues.Wrap(err, "converting attachment")
			}

			convertedAttachments = append(convertedAttachments, am)
		}

		instance["attachments"] = convertedAttachments
	}

	return nil
}

// Adding checks to ensure that the data is in the format that we expect M365 to return
func validateCancelledOccurrences(event models.Eventable) error {
	cancelledOccurrences := event.GetAdditionalData()["cancelledOccurrences"]
	if cancelledOccurrences != nil {
		co, ok := cancelledOccurrences.([]any)
		if !ok {
			return clues.New("converting cancelledOccurrences to []any").
				With("type", fmt.Sprintf("%T", cancelledOccurrences))
		}

		for _, instance := range co {
			instance, err := str.AnyToString(instance)
			if err != nil {
				return err
			}

			// There might be multiple `.` in the ID and hence >2
			splits := strings.Split(instance, ".")
			if len(splits) < 2 {
				return clues.New("unexpected cancelled event format").
					With("instance", instance)
			}

			startStr := splits[len(splits)-1]

			_, err = dttm.ParseTime(startStr)
			if err != nil {
				return clues.Wrap(err, "parsing cancelled event date")
			}
		}
	}

	return nil
}

func parseableToMap(att serialization.Parsable) (map[string]any, error) {
	var item map[string]any

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	if err := writer.WriteObjectValue("", att); err != nil {
		return nil, err
	}

	ats, err := writer.GetSerializedContent()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ats, &item)
	if err != nil {
		return nil, clues.Wrap(err, "unmarshalling serialized attachment")
	}

	return item, nil
}

func (c Events) GetAttachments(
	ctx context.Context,
	immutableIDs bool,
	userID string,
	itemID string,
) ([]models.Attachmentable, error) {
	config := &users.ItemEventsItemAttachmentsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemEventsItemAttachmentsRequestBuilderGetQueryParameters{
			Expand: []string{"microsoft.graph.itemattachment/item"},
		},
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	attached, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		Events().
		ByEventId(itemID).
		Attachments().
		Get(ctx, config)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "event attachment download")
	}

	return attached.GetValue(), nil
}

func (c Events) DeleteAttachment(
	ctx context.Context,
	userID, calendarID, eventID, attachmentID string,
) error {
	return c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(calendarID).
		Events().
		ByEventId(eventID).
		Attachments().
		ByAttachmentId(attachmentID).
		Delete(ctx, nil)
}

func (c Events) GetItemInstances(
	ctx context.Context,
	userID, itemID string,
	startDate, endDate string,
) ([]models.Eventable, error) {
	config := &users.ItemEventsItemInstancesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemEventsItemInstancesRequestBuilderGetQueryParameters{
			Select:        []string{"id"},
			StartDateTime: ptr.To(startDate),
			EndDateTime:   ptr.To(endDate),
		},
	}

	events, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Events().
		ByEventId(itemID).
		Instances().
		Get(ctx, config)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return events.GetValue(), nil
}

func (c Events) PostItem(
	ctx context.Context,
	userID, containerID string,
	body models.Eventable,
) (models.Eventable, error) {
	rawURL := fmt.Sprintf(eventPostBetaURLTemplate, userID, containerID)
	builder := users.NewItemCalendarsItemEventsRequestBuilder(rawURL, c.Stable.Adapter())

	itm, err := builder.Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating calendar event")
	}

	return itm, nil
}

func (c Events) PatchItem(
	ctx context.Context,
	userID, eventID string,
	body models.Eventable,
) (models.Eventable, error) {
	rawURL := fmt.Sprintf(eventPatchBetaURLTemplate, userID, eventID)
	builder := users.NewItemCalendarsItemEventsEventItemRequestBuilder(rawURL, c.Stable.Adapter())

	itm, err := builder.Patch(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "updating calendar event")
	}

	return itm, nil
}

func (c Events) DeleteItem(
	ctx context.Context,
	userID, itemID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := c.Service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.
		Client().
		Users().
		ByUserId(userID).
		Events().
		ByEventId(itemID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting calendar event")
	}

	return nil
}

func (c Events) PostSmallAttachment(
	ctx context.Context,
	userID, containerID, parentItemID string,
	body models.Attachmentable,
) error {
	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Events().
		ByEventId(parentItemID).
		Attachments().
		Post(ctx, body, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "uploading small event attachment")
	}

	return nil
}

func (c Events) PostLargeAttachment(
	ctx context.Context,
	userID, containerID, parentItemID, itemName string,
	content []byte,
) (string, error) {
	size := int64(len(content))
	session := users.NewItemCalendarEventsItemAttachmentsCreateUploadSessionPostRequestBody()
	session.SetAttachmentItem(makeSessionAttachment(itemName, size))

	us, err := c.LargeItem.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Events().
		ByEventId(parentItemID).
		Attachments().
		CreateUploadSession().
		Post(ctx, session, nil)
	if err != nil {
		return "", graph.Wrap(ctx, err, "uploading large event attachment")
	}

	url := ptr.Val(us.GetUploadUrl())
	w := graph.NewLargeItemWriter(parentItemID, url, size)
	copyBuffer := make([]byte, graph.AttachmentChunkSize)

	_, err = io.CopyBuffer(w, bytes.NewReader(content), copyBuffer)
	if err != nil {
		return "", clues.Wrap(err, "buffering large attachment content").WithClues(ctx)
	}

	return w.ID, nil
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

func BytesToEventable(body []byte) (models.Eventable, error) {
	v, err := createFromBytes(body, models.CreateEventFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to event")
	}

	return v.(models.Eventable), nil
}

func (c Events) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	userID, itemID string,
) ([]byte, error) {
	event, ok := item.(models.Eventable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("item is not an Eventable: %T", item))
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(event.GetId()))

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	if err := writer.WriteObjectValue("", event); err != nil {
		return nil, graph.Stack(ctx, err)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "serializing event")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// CalendarDisplayable is a wrapper that complies with the
// models.Calendarable interface with the graph.Container
// interfaces. Calendars do not have a parentFolderID.
// Therefore, that value will always return nil.
type CalendarDisplayable struct {
	models.Calendarable
}

// GetDisplayName returns the *string of the models.Calendable
// variant:  calendar.GetName()
func (c CalendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

// GetParentFolderId returns the default calendar name address
// EventCalendars have a flat hierarchy and Calendars are rooted
// at the default
//
//nolint:revive
func (c CalendarDisplayable) GetParentFolderId() *string {
	return nil
}

func EventInfo(evt models.Eventable) *details.ExchangeInfo {
	var (
		organizer string
		subject   = ptr.Val(evt.GetSubject())
		recurs    bool
		start     = time.Time{}
		end       = time.Time{}
		created   = ptr.Val(evt.GetCreatedDateTime())
	)

	if evt.GetOrganizer() != nil &&
		evt.GetOrganizer().GetEmailAddress() != nil {
		organizer = ptr.Val(evt.GetOrganizer().GetEmailAddress().GetAddress())
	}

	if evt.GetRecurrence() != nil {
		recurs = true
	}

	if evt.GetStart() != nil && len(ptr.Val(evt.GetStart().GetDateTime())) > 0 {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		startTime := ptr.Val(evt.GetStart().GetDateTime()) + "Z"

		output, err := dttm.ParseTime(startTime)
		if err == nil {
			start = output
		}
	}

	if evt.GetEnd() != nil && len(ptr.Val(evt.GetEnd().GetDateTime())) > 0 {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		endTime := ptr.Val(evt.GetEnd().GetDateTime()) + "Z"

		output, err := dttm.ParseTime(endTime)
		if err == nil {
			end = output
		}
	}

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeEvent,
		Organizer:   organizer,
		Subject:     subject,
		EventStart:  start,
		EventEnd:    end,
		EventRecurs: recurs,
		Created:     created,
		Modified:    ptr.OrNow(evt.GetLastModifiedDateTime()),
	}
}

func EventFromMap(ev map[string]any) (models.Eventable, error) {
	instBytes, err := json.Marshal(ev)
	if err != nil {
		return nil, clues.Wrap(err, "marshaling event exception instance")
	}

	body, err := BytesToEventable(instBytes)
	if err != nil {
		return nil, clues.Wrap(err, "converting exception event bytes to Eventable")
	}

	return body, nil
}

func eventCollisionKeyProps() []string {
	return idAnd("subject")
}

// EventCollisionKey constructs a key from the eventable's creation time, subject, and organizer.
// collision keys are used to identify duplicate item conflicts for handling advanced restoration config.
func EventCollisionKey(item models.Eventable) string {
	if item == nil {
		return ""
	}

	return ptr.Val(item.GetSubject())
}
