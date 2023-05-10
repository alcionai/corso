package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
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
// methods
// ---------------------------------------------------------------------------

// CreateCalendar makes an event Calendar with the name in the user's M365 exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-calendars?view=graph-rest-1.0&tabs=go
func (c Events) CreateCalendar(
	ctx context.Context,
	user, calendarName string,
) (models.Calendarable, error) {
	requestbody := models.NewCalendar()
	requestbody.SetName(&calendarName)

	mdl, err := c.Stable.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating calendar")
	}

	return mdl, nil
}

// DeleteContainer removes a calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func (c Events) DeleteContainer(
	ctx context.Context,
	user, calendarID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

func (c Events) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	service, err := c.service()
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	ofc, err := optionsForCalendarsByID([]string{"name", "owner"})
	if err != nil {
		return nil, graph.Wrap(ctx, err, "setting event calendar options")
	}

	cal, err := service.Client().UsersById(userID).CalendarsById(containerID).Get(ctx, ofc)
	if err != nil {
		return nil, graph.Stack(ctx, err).WithClues(ctx)
	}

	return graph.CalendarDisplayable{Calendarable: cal}, nil
}

// GetContainerByName fetches a calendar by name
func (c Events) GetContainerByName(
	ctx context.Context,
	userID, name string,
) (models.Calendarable, error) {
	filter := fmt.Sprintf("name eq '%s'", name)
	options := &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemCalendarsRequestBuilderGetQueryParameters{
			Filter: &filter,
		},
	}

	ctx = clues.Add(ctx, "calendar_name", name)

	resp, err := c.Stable.Client().UsersById(userID).Calendars().Get(ctx, options)
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

	if err := checkIDAndName(cd); err != nil {
		return nil, err
	}

	return cal, nil
}

// GetItem retrieves an Eventable item.
func (c Events) GetItem(
	ctx context.Context,
	user, itemID string,
	immutableIDs bool,
	errs *fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		err      error
		event    models.Eventable
		header   = buildPreferHeaders(false, immutableIDs)
		itemOpts = &users.ItemEventsEventItemRequestBuilderGetRequestConfiguration{
			Headers: header,
		}
	)

	event, err = c.Stable.Client().UsersById(user).EventsById(itemID).Get(ctx, itemOpts)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	if ptr.Val(event.GetHasAttachments()) || HasAttachments(event.GetBody()) {
		options := &users.ItemEventsItemAttachmentsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemEventsItemAttachmentsRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
			Headers: header,
		}

		attached, err := c.LargeItem.
			Client().
			UsersById(user).
			EventsById(itemID).
			Attachments().
			Get(ctx, options)
		if err != nil {
			return nil, nil, graph.Wrap(ctx, err, "event attachment download")
		}

		event.SetAttachments(attached.GetValue())
	}

	return event, EventInfo(event), nil
}

// EnumerateContainers iterates through all of the users current
// calendars, converting each to a graph.CacheFolder, and
// calling fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Events) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
	errs *fault.Bus,
) error {
	service, err := c.service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	ofc, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return graph.Wrap(ctx, err, "setting calendar options")
	}

	el := errs.Local()
	builder := service.Client().UsersById(userID).Calendars()

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, ofc)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, cal := range resp.GetValue() {
			if el.Failure() != nil {
				break
			}

			cd := CalendarDisplayable{Calendarable: cal}
			if err := checkIDAndName(cd); err != nil {
				errs.AddRecoverable(graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
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
			if err := fn(temp); err != nil {
				errs.AddRecoverable(graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemCalendarsRequestBuilder(link, service.Adapter())
	}

	return el.Failure()
}

const (
	eventBetaDeltaURLTemplate = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events/delta"
)

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &eventPager{}

type eventPager struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

func NewEventPager(
	ctx context.Context,
	gs graph.Servicer,
	user, calendarID string,
	immutableIDs bool,
) (itemPager, error) {
	options := &users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		Headers: buildPreferHeaders(true, immutableIDs),
	}

	builder := gs.Client().UsersById(user).CalendarsById(calendarID).Events()
	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		gri, err := builder.ToGetRequestInformation(ctx, options)
		if err != nil {
			logger.CtxErr(ctx, err).Error("getting builder info")
		} else {
			logger.Ctx(ctx).
				Infow("builder path-parameters", "path_parameters", gri.PathParameters)
		}
	}

	return &eventPager{gs, builder, options}, nil
}

func (p *eventPager) getPage(ctx context.Context) (api.PageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *eventPager) setNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *eventPager) reset(context.Context) {}

func (p *eventPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ itemPager = &eventDeltaPager{}

type eventDeltaPager struct {
	gs         graph.Servicer
	user       string
	calendarID string
	builder    *users.ItemCalendarsItemEventsDeltaRequestBuilder
	options    *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration
}

func NewEventDeltaPager(
	ctx context.Context,
	gs graph.Servicer,
	user, calendarID, deltaURL string,
	immutableIDs bool,
) (itemPager, error) {
	options := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration{
		Headers: buildPreferHeaders(true, immutableIDs),
	}

	var builder *users.ItemCalendarsItemEventsDeltaRequestBuilder

	if deltaURL == "" {
		builder = getEventDeltaBuilder(ctx, gs, user, calendarID, options)
	} else {
		builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(deltaURL, gs.Adapter())
	}

	return &eventDeltaPager{gs, user, calendarID, builder, options}, nil
}

func getEventDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	user string,
	calendarID string,
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
	rawURL := fmt.Sprintf(eventBetaDeltaURLTemplate, user, calendarID)
	builder := users.NewItemCalendarsItemEventsDeltaRequestBuilder(rawURL, gs.Adapter())

	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		gri, err := builder.ToGetRequestInformation(ctx, options)
		if err != nil {
			logger.CtxErr(ctx, err).Error("getting builder info")
		} else {
			logger.Ctx(ctx).
				Infow("builder path-parameters", "path_parameters", gri.PathParameters)
		}
	}

	return builder
}

func (p *eventDeltaPager) getPage(ctx context.Context) (api.PageLinker, error) {
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
	p.builder = getEventDeltaBuilder(ctx, p.gs, p.user, p.calendarID, p.options)
}

func (p *eventDeltaPager) valuesIn(pl api.PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

func (c Events) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, calendarID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	ctx = clues.Add(
		ctx,
		"container_id", calendarID)

	pager, err := NewEventPager(ctx, service, user, calendarID, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating non-delta pager")
	}

	deltaPager, err := NewEventDeltaPager(ctx, service, user, calendarID, oldDelta, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating delta pager")
	}

	return getAddedAndRemovedItemIDs(ctx, service, pager, deltaPager, oldDelta, canMakeDeltaQueries)
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

// Serialize transforms the event into a byte slice.
func (c Events) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	user, itemID string,
) ([]byte, error) {
	event, ok := item.(models.Eventable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("item is not an Eventable: %T", item))
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(event.GetId()))

	var (
		err    error
		writer = kjson.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", event); err != nil {
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
