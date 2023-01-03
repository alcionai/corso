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
)

// CreateCalendar makes an event Calendar with the name in the user's M365 exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-calendars?view=graph-rest-1.0&tabs=go
func (c Client) CreateCalendar(
	ctx context.Context,
	user, calendarName string,
) (models.Calendarable, error) {
	requestbody := models.NewCalendar()
	requestbody.SetName(&calendarName)

	return c.stable.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
}

// DeleteCalendar removes calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func (c Client) DeleteCalendar(
	ctx context.Context,
	user, calendarID string,
) error {
	return c.stable.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
}

// RetrieveEventDataForUser is a GraphRetrievalFunc that returns event data.
// Calendarable and attachment fields are omitted due to size
func (c Client) RetrieveEventDataForUser(
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

// TODO: we want this to be the full handler, not only the builder.
// but this halfway point minimizes changes for now.
func (c Client) GetCalendarsBuilder(
	ctx context.Context,
	userID string,
	optionalFields ...string,
) (
	*users.ItemCalendarsRequestBuilder,
	*users.ItemCalendarsRequestBuilderGetRequestConfiguration,
	*graph.Service,
	error,
) {
	service, err := c.service()
	if err != nil {
		return nil, nil, nil, err
	}

	ofcf, err := optionsForCalendars(optionalFields)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "options for event calendars: %v", optionalFields)
	}

	builder := service.Client().UsersById(userID).Calendars()

	return builder, ofcf, service, nil
}

// FetchEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
func (c Client) FetchEventIDsFromCalendar(
	ctx context.Context,
	user, calendarID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForEventsByCalendar([]string{"id"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	builder := service.Client().UsersById(user).CalendarsById(calendarID).Events()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			if err := graph.IsErrDeletedInFlight(err); err != nil {
				return nil, nil, DeltaUpdate{}, err
			}

			return nil, nil, DeltaUpdate{}, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, item := range resp.GetValue() {
			if item.GetId() == nil {
				errs = multierror.Append(
					errs,
					errors.Errorf("event with nil ID in calendar %s", calendarID),
				)

				// TODO(ashmrtn): Handle fail-fast.
				continue
			}

			ids = append(ids, *item.GetId())
		}

		nextLink := resp.GetOdataNextLink()
		if nextLink == nil || len(*nextLink) == 0 {
			break
		}

		builder = users.NewItemCalendarsItemEventsRequestBuilder(*nextLink, service.Adapter())
	}

	// Events don't have a delta endpoint so just return an empty string.
	return ids, nil, DeltaUpdate{}, errs.ErrorOrNil()
}
