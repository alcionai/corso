package api

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/hashicorp/go-multierror"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
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

	return c.stable.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
}

// DeleteContainer removes a calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func (c Events) DeleteContainer(
	ctx context.Context,
	user, calendarID string,
) error {
	return c.stable.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
}

func (c Events) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	ofc, err := optionsForCalendarsByID([]string{"name", "owner"})
	if err != nil {
		return nil, errors.Wrap(err, "options for event calendar")
	}

	var cal models.Calendarable

	err = graph.RunWithRetry(func() error {
		cal, err = service.Client().UsersById(userID).CalendarsById(containerID).Get(ctx, ofc)
		return err
	})

	if err != nil {
		return nil, err
	}

	return graph.CalendarDisplayable{Calendarable: cal}, nil
}

// GetItem retrieves an Eventable item.
func (c Events) GetItem(
	ctx context.Context,
	user, itemID string,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		event models.Eventable
		err   error
	)

	err = graph.RunWithRetry(func() error {
		event, err = c.stable.Client().UsersById(user).EventsById(itemID).Get(ctx, nil)
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	var (
		errs    *multierror.Error
		options = &users.ItemEventsItemAttachmentsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemEventsItemAttachmentsRequestBuilderGetQueryParameters{
				Expand: []string{"microsoft.graph.itemattachment/item"},
			},
		}
	)

	if *event.GetHasAttachments() || HasAttachments(event.GetBody()) {
		for count := 0; count < numberOfRetries; count++ {
			attached, err := c.largeItem.
				Client().
				UsersById(user).
				EventsById(itemID).
				Attachments().
				Get(ctx, options)
			if err == nil {
				event.SetAttachments(attached.GetValue())
				break
			}

			logger.Ctx(ctx).Debugw("retrying event attachment download", "err", err)
			errs = multierror.Append(errs, err)
		}

		if err != nil {
			logger.Ctx(ctx).Errorw("event attachment download exceeded maximum retries", "err", errs)
			return nil, nil, support.WrapAndAppend(itemID, errors.Wrap(err, "download event attachment"), nil)
		}
	}

	return event, EventInfo(event), nil
}

func (c Client) GetAllCalendarNamesForUser(
	ctx context.Context,
	user string,
) (serialization.Parsable, error) {
	options, err := optionsForCalendars([]string{"name", "owner"})
	if err != nil {
		return nil, err
	}

	var resp models.CalendarCollectionResponseable

	err = graph.RunWithRetry(func() error {
		resp, err = c.stable.Client().UsersById(user).Calendars().Get(ctx, options)
		return err
	})

	return resp, err
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
) error {
	service, err := c.service()
	if err != nil {
		return err
	}

	var (
		resp models.CalendarCollectionResponseable
		errs *multierror.Error
	)

	ofc, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return errors.Wrapf(err, "options for event calendars")
	}

	builder := service.Client().UsersById(userID).Calendars()

	for {
		var err error

		err = graph.RunWithRetry(func() error {
			resp, err = builder.Get(ctx, ofc)
			return err
		})

		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, cal := range resp.GetValue() {
			cd := CalendarDisplayable{Calendarable: cal}
			if err := checkIDAndName(cd); err != nil {
				errs = multierror.Append(err, errs)
				continue
			}

			temp := graph.NewCacheFolder(cd, path.Builder{}.Append(*cd.GetDisplayName()))

			err = fn(temp)
			if err != nil {
				errs = multierror.Append(err, errs)
				continue
			}
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = users.NewItemCalendarsRequestBuilder(*resp.GetOdataNextLink(), service.Adapter())
	}

	return errs.ErrorOrNil()
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
	var (
		resp api.DeltaPageLinker
		err  error
	)

	err = graph.RunWithRetry(func() error {
		resp, err = p.builder.Get(ctx, p.options)
		return err
	})

	return resp, err
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

	var (
		resetDelta bool
		errs       *multierror.Error
	)

	ctx = clues.AddAll(
		ctx,
		"category", selectors.ExchangeEvent,
		"calendar_id", calendarID)

	if len(oldDelta) > 0 {
		builder := users.NewItemCalendarsItemEventsDeltaRequestBuilder(oldDelta, service.Adapter())
		pgr := &eventPager{service, builder, nil}

		added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
		// note: happy path, not the error condition
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, false}, errs.ErrorOrNil()
		}
		// only return on error if it is NOT a delta issue.
		// on bad deltas we retry the call with the regular builder
		if !graph.IsErrInvalidDelta(err) {
			return nil, nil, DeltaUpdate{}, err
		}

		resetDelta = true
		errs = nil
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

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	// Events don't have a delta endpoint so just return an empty string.
	return added, removed, DeltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
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
		return nil, fmt.Errorf("expected Eventable, got %T", item)
	}

	ctx = clues.Add(ctx, "item_id", *event.GetId())

	var (
		err    error
		writer = kioser.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", event); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, errors.Wrap(err, "serializing calendar event")
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
		organizer, subject string
		recurs             bool
		start              = time.Time{}
		end                = time.Time{}
		created            = time.Time{}
	)

	if evt.GetOrganizer() != nil &&
		evt.GetOrganizer().GetEmailAddress() != nil &&
		evt.GetOrganizer().GetEmailAddress().GetAddress() != nil {
		organizer = *evt.GetOrganizer().
			GetEmailAddress().
			GetAddress()
	}

	if evt.GetSubject() != nil {
		subject = *evt.GetSubject()
	}

	if evt.GetRecurrence() != nil {
		recurs = true
	}

	if evt.GetStart() != nil &&
		evt.GetStart().GetDateTime() != nil {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		startTime := *evt.GetStart().GetDateTime() + "Z"

		output, err := common.ParseTime(startTime)
		if err == nil {
			start = output
		}
	}

	if evt.GetEnd() != nil &&
		evt.GetEnd().GetDateTime() != nil {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		endTime := *evt.GetEnd().GetDateTime() + "Z"

		output, err := common.ParseTime(endTime)
		if err == nil {
			end = output
		}
	}

	if evt.GetCreatedDateTime() != nil {
		created = *evt.GetCreatedDateTime()
	}

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeEvent,
		Organizer:   organizer,
		Subject:     subject,
		EventStart:  start,
		EventEnd:    end,
		EventRecurs: recurs,
		Created:     created,
		Modified:    orNow(evt.GetLastModifiedDateTime()),
	}
}
