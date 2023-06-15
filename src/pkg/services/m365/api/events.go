package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

// EnumerateContainers iterates through all of the users current
// calendars, converting each to a graph.CacheFolder, and
// calling fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Events) EnumerateContainers(
	ctx context.Context,
	userID, baseContainerID string,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	var (
		el     = errs.Local()
		config = &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemCalendarsRequestBuilderGetQueryParameters{
				Select: idAnd("name"),
			},
		}
		builder = c.Stable.
			Client().
			Users().
			ByUserId(userID).
			Calendars()
	)

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, config)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, cal := range resp.GetValue() {
			if el.Failure() != nil {
				break
			}

			cd := CalendarDisplayable{Calendarable: cal}
			if err := graph.CheckIDAndName(cd); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(cal.GetId()),
				"container_name", ptr.Val(cal.GetName()))

			temp := graph.NewCacheFolder(
				cd,
				path.Builder{}.Append(ptr.Val(cd.GetId())),          // storage path
				path.Builder{}.Append(ptr.Val(cd.GetDisplayName()))) // display location
			if err := fn(&temp); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemCalendarsRequestBuilder(link, c.Stable.Adapter())
	}

	return el.Failure()
}

const (
	eventDeltaBetaURLTemplate = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events/delta"
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

	rawURL := fmt.Sprintf(eventExceptionsBetaURLTemplate, userID, itemID)
	builder := users.NewItemEventsEventItemRequestBuilder(rawURL, c.Stable.Adapter())

	event, err = builder.Get(ctx, config)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	attachments, err := c.getAttachments(ctx, event, immutableIDs, userID, itemID)
	if err != nil {
		return nil, nil, err
	}

	event.SetAttachments(attachments)

	// Fetch attachments for exceptions
	exceptionOccurrences := event.GetAdditionalData()["exceptionOccurrences"]
	if exceptionOccurrences != nil {
		eo, ok := exceptionOccurrences.([]interface{})
		if !ok {
			return nil, nil, clues.New("converting exceptionOccurrences to []interface{}").
				With("type", fmt.Sprintf("%T", exceptionOccurrences))
		}

		for _, inst := range eo {
			evi, ok := inst.(map[string]any)
			if !ok {
				return nil, nil, clues.New("converting instance to map[string]interface{}").
					With("type", fmt.Sprintf("%T", inst))
			}

			evt, err := EventFromMap(evi)
			if err != nil {
				return nil, nil, clues.Wrap(err, "parsing exception event")
			}

			// OPTIMIZATION: We don't have to store any of the
			// attachments that carry over from the original
			attachments, err := c.getAttachments(ctx, evt, immutableIDs, userID, ptr.Val(evt.GetId()))
			if err != nil {
				return nil, nil, clues.Wrap(err, "getting exception attachments").
					With("exception_event_id", ptr.Val(evt.GetId()))
			}

			// This odd roundabout way of doing this is required as
			// the json serialization at the end does not serialize if
			// you just pass in a models.Attachmentable
			atts := []map[string]interface{}{}

			for _, att := range attachments {
				writer := kjson.NewJsonSerializationWriter()
				defer writer.Close()

				if err := writer.WriteObjectValue("", att); err != nil {
					return nil, nil, graph.Stack(ctx, err)
				}

				ats, err := writer.GetSerializedContent()
				if err != nil {
					return nil, nil, graph.Wrap(ctx, err, "serializing event")
				}

				atm := map[string]interface{}{}

				err = json.Unmarshal(ats, &atm)
				if err != nil {
					return nil, nil, clues.Wrap(err, "unmarshalling serialized attachment")
				}

				atts = append(atts, atm)
			}

			evi["attachments"] = atts
		}
	}

	return event, EventInfo(event), nil
}

func (c Events) getAttachments(
	ctx context.Context,
	event models.Eventable,
	immutableIDs bool,
	userID string,
	itemID string,
) ([]models.Attachmentable, error) {
	if !ptr.Val(event.GetHasAttachments()) && !HasAttachments(event.GetBody()) {
		return nil, nil
	}

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

func (c Events) UpdateItem(
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
	size int64,
	body models.Attachmentable,
) (models.UploadSessionable, error) {
	bs, err := GetAttachmentContent(body)
	if err != nil {
		return nil, clues.Wrap(err, "serializing attachment content").WithClues(ctx)
	}

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
		return nil, graph.Wrap(ctx, err, "uploading large event attachment")
	}

	url := ptr.Val(us.GetUploadUrl())
	w := graph.NewLargeItemWriter(parentItemID, url, size)
	copyBuffer := make([]byte, graph.AttachmentChunkSize)

	_, err = io.CopyBuffer(w, bytes.NewReader(bs), copyBuffer)
	if err != nil {
		return nil, clues.Wrap(err, "buffering large attachment content").WithClues(ctx)
	}

	return us, nil
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &eventPager{}

type eventPager struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventPager(
	ctx context.Context,
	userID, containerID string,
	immutableIDs bool,
) (itemPager, error) {
	options := &users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Events()

	return &eventPager{c.Stable, builder, options}, nil
}

func (p *eventPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Eventable]{PageLinkValuer: resp}, nil
}

func (p *eventPager) setNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *eventPager) reset(context.Context) {}

func (p *eventPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ itemPager = &eventDeltaPager{}

type eventDeltaPager struct {
	gs          graph.Servicer
	userID      string
	containerID string
	builder     *users.ItemCalendarsItemEventsDeltaRequestBuilder
	options     *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventDeltaPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) (itemPager, error) {
	options := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferPageSize(maxDeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	var builder *users.ItemCalendarsItemEventsDeltaRequestBuilder

	if oldDelta == "" {
		builder = getEventDeltaBuilder(ctx, c.Stable, userID, containerID, options)
	} else {
		builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	}

	return &eventDeltaPager{c.Stable, userID, containerID, builder, options}, nil
}

func getEventDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	userID, containerID string,
	options *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemCalendarsItemEventsDeltaRequestBuilder {
	// Graph SDK only supports delta queries against events on the beta version, so we're
	// manufacturing use of the beta version url to make the call instead.
	// See: https://learn.microsoft.com/ko-kr/graph/api/event-delta?view=graph-rest-beta&tabs=http
	// Note that the delta item body is skeletal compared to the actual event struct.  Lucky
	// for us, we only need the item ID.  As a result, even though we hacked the version, the
	// response body parses properly into the v1.0 structs and complies with our wanted interfaces.
	// Likewise, the NextLink and DeltaLink odata tags carry our hack forward, so the rest of the code
	// works as intended (until, at least, we want to _not_ call the beta anymore).
	rawURL := fmt.Sprintf(eventDeltaBetaURLTemplate, userID, containerID)
	builder := users.NewItemCalendarsItemEventsDeltaRequestBuilder(rawURL, gs.Adapter())

	return builder
}

func (p *eventDeltaPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *eventDeltaPager) setNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventDeltaPager) reset(ctx context.Context) {
	p.builder = getEventDeltaBuilder(ctx, p.gs, p.userID, p.containerID, p.options)
}

func (p *eventDeltaPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

func (c Events) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	ctx = clues.Add(ctx, "container_id", containerID)

	pager, err := c.NewEventPager(ctx, userID, containerID, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating non-delta pager")
	}

	deltaPager, err := c.NewEventDeltaPager(ctx, userID, containerID, oldDelta, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating delta pager")
	}

	return getAddedAndRemovedItemIDs(ctx, c.Stable, pager, deltaPager, oldDelta, canMakeDeltaQueries)
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
