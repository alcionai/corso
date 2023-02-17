package api

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common"
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

	mdl, err := c.stable.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return mdl, nil
}

// DeleteContainer removes a calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func (c Events) DeleteContainer(
	ctx context.Context,
	user, calendarID string,
) error {
	err := c.stable.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
	if err != nil {
		return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return nil
}

func (c Events) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	service, err := c.service()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	ofc, err := optionsForCalendarsByID([]string{"name", "owner"})
	if err != nil {
		return nil, clues.Wrap(err, "setting event calendar options").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	cal, err := service.Client().UsersById(userID).CalendarsById(containerID).Get(ctx, ofc)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return graph.CalendarDisplayable{Calendarable: cal}, nil
}

// GetItem retrieves an Eventable item.
func (c Events) GetItem(
	ctx context.Context,
	user, itemID string,
	errs *fault.Errors,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		err   error
		event models.Eventable
	)

	event, err = c.stable.Client().UsersById(user).EventsById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	if *event.GetHasAttachments() || HasAttachments(event.GetBody()) {
		options := &users.ItemEventsItemAttachmentsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemEventsItemAttachmentsRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
		}

		attached, err := c.largeItem.
			Client().
			UsersById(user).
			EventsById(itemID).
			Attachments().
			Get(ctx, options)
		if err != nil {
			return nil, nil, clues.Wrap(err, "event attachment download").WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		event.SetAttachments(attached.GetValue())
	}

	return event, EventInfo(event), nil
}

// EnumerateContainers iterates through all of the users current
// calendars, converting each to a graph.CacheFolder, and
// calling fn(cf) on each one.  If fn(cf) errors, the error is
// aggregated into a multierror that gets returned to the caller.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Events) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
	errs *fault.Errors,
) error {
	service, err := c.service()
	if err != nil {
		return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	ofc, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return clues.Wrap(err, "setting calendar options").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	builder := service.Client().UsersById(userID).Calendars()

	for {
		resp, err := builder.Get(ctx, ofc)
		if err != nil {
			return clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		for _, cal := range resp.GetValue() {
			cd := CalendarDisplayable{Calendarable: cal}
			if err := checkIDAndName(cd); err != nil {
				errs.Add(clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...))
				continue
			}

			fctx := clues.AddAll(
				ctx,
				"container_id", ptr.Val(cal.GetId()),
				"container_name", ptr.Val(cal.GetName()))

			temp := graph.NewCacheFolder(
				cd,
				path.Builder{}.Append(ptr.Val(cd.GetId())),          // storage path
				path.Builder{}.Append(ptr.Val(cd.GetDisplayName()))) // display location
			if err := fn(temp); err != nil {
				errs.Add(clues.Stack(err).WithClues(fctx).WithAll(graph.ErrData(err)...))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemCalendarsRequestBuilder(link, service.Adapter())
	}

	return errs.Err()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &eventPager{}

const (
	eventBetaDeltaURLTemplate = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events/delta"
)

type eventPager struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsDeltaRequestBuilder
	options *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration
}

func (p *eventPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	return resp, nil
}

func (p *eventPager) setNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventPager) valuesIn(pl api.DeltaPageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

func (c Events) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, calendarID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var resetDelta bool

	ctx = clues.AddAll(
		ctx,
		"calendar_id", calendarID)

	if len(oldDelta) > 0 {
		var (
			builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(oldDelta, service.Adapter())
			pgr     = &eventPager{service, builder, nil}
		)

		added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
		// note: happy path, not the error condition
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, false}, nil
		}
		// only return on error if it is NOT a delta issue.
		// on bad deltas we retry the call with the regular builder
		if !graph.IsErrInvalidDelta(err) {
			return nil, nil, DeltaUpdate{}, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
		}

		resetDelta = true
	}

	// Graph SDK only supports delta queries against events on the beta version, so we're
	// manufacturing use of the beta version url to make the call instead.
	// See: https://learn.microsoft.com/ko-kr/graph/api/event-delta?view=graph-rest-beta&tabs=http
	// Note that the delta item body is skeletal compared to the actual event struct.  Lucky
	// for us, we only need the item ID.  As a result, even though we hacked the version, the
	// response body parses properly into the v1.0 structs and complies with our wanted interfaces.
	// Likewise, the NextLink and DeltaLink odata tags carry our hack forward, so the rest of the code
	// works as intended (until, at least, we want to _not_ call the beta anymore).
	rawURL := fmt.Sprintf(eventBetaDeltaURLTemplate, user, calendarID)
	builder := users.NewItemCalendarsItemEventsDeltaRequestBuilder(rawURL, service.Adapter())
	pgr := &eventPager{service, builder, nil}

	gri, err := builder.ToGetRequestInformation(ctx, nil)
	if err != nil {
		logger.Ctx(ctx).Errorw("getting builder info", "error", err)
	} else {
		uri, err := gri.GetUri()
		if err != nil {
			logger.Ctx(ctx).Errorw("getting builder uri", "error", err)
		} else {
			logger.Ctx(ctx).Infow("calendar builder", "user", user, "directoryID", calendarID, "uri", uri)
		}
	}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	// Events don't have a delta endpoint so just return an empty string.
	return added, removed, DeltaUpdate{deltaURL, resetDelta}, nil
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
		return nil, clues.Wrap(fmt.Errorf("parseable type: %T", item), "parsable is not an Eventable")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(event.GetId()))

	var (
		err    error
		writer = kioser.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", event); err != nil {
		return nil, clues.Stack(err).WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, clues.Wrap(err, "serializing event").WithClues(ctx).WithAll(graph.ErrData(err)...)
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

		output, err := common.ParseTime(startTime)
		if err == nil {
			start = output
		}
	}

	if evt.GetEnd() != nil && len(ptr.Val(evt.GetEnd().GetDateTime())) > 0 {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		endTime := ptr.Val(evt.GetEnd().GetDateTime()) + "Z"

		output, err := common.ParseTime(endTime)
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
