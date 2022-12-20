package api

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
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

type eventPager struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

type eventWrapper struct {
	models.EventCollectionResponseable
}

func (ew eventWrapper) GetOdataDeltaLink() *string {
	return nil
}

func (p *eventPager) getPage(ctx context.Context) (pageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	return eventWrapper{resp}, err
}

func (p *eventPager) setNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventPager) idsIn(pl pageLinker) ([]getIDer, error) {
	return toIders(pl)
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

// DeleteCalendar removes calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func (c Events) DeleteCalendar(
	ctx context.Context,
	user, calendarID string,
) error {
	return c.stable.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
}

// RetrieveEventDataForUser is a GraphRetrievalFunc that returns event data.
func (c Events) RetrieveEventDataForUser(
	ctx context.Context,
	user, m365ID string,
) (serialization.Parsable, error) {
	return c.stable.Client().UsersById(user).EventsById(m365ID).Get(ctx, nil)
}

func (c Client) GetAllCalendarNamesForUser(
	ctx context.Context,
	user string,
) (serialization.Parsable, error) {
	options, err := optionsForCalendars([]string{"name", "owner"})
	if err != nil {
		return nil, err
	}

	return c.stable.Client().UsersById(user).Calendars().Get(ctx, options)
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

	var errs *multierror.Error

	ofc, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return errors.Wrapf(err, "options for event calendars")
	}

	builder := service.Client().UsersById(userID).Calendars()

	for {
		resp, err := builder.Get(ctx, ofc)
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

func (c Events) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, calendarID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var errs *multierror.Error

	errUpdater := func(err error) {
		errs = multierror.Append(errs, errors.Wrap(err, "calendar "+calendarID))
	}

	options, err := optionsForEventsByCalendar([]string{"id"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	builder := service.Client().UsersById(user).CalendarsById(calendarID).Events()
	pgr := &eventPager{service, builder, options}

	added, _, _, err := getContainerIDs(ctx, pgr, errUpdater)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	// Events don't have a delta endpoint so just return an empty string.
	return added, nil, DeltaUpdate{}, errs.ErrorOrNil()
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
