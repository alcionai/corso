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
func CreateCalendar(ctx context.Context, gs graph.Servicer, user, calendarName string) (models.Calendarable, error) {
	requestbody := models.NewCalendar()
	requestbody.SetName(&calendarName)

	return gs.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
}

// DeleteCalendar removes calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func DeleteCalendar(ctx context.Context, gs graph.Servicer, user, calendarID string) error {
	return gs.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
}

// RetrieveEventDataForUser is a GraphRetrievalFunc that returns event data.
// Calendarable and attachment fields are omitted due to size
func RetrieveEventDataForUser(
	ctx context.Context,
	gs graph.Servicer,
	user, m365ID string,
) (serialization.Parsable, error) {
	return gs.Client().UsersById(user).EventsById(m365ID).Get(ctx, nil)
}

func GetAllCalendarNamesForUser(ctx context.Context, gs graph.Servicer, user string) (serialization.Parsable, error) {
	options, err := optionsForCalendars([]string{"name", "owner"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Calendars().Get(ctx, options)
}

// TODO: we want this to be the full handler, not only the builder.
// but this halfway point minimizes changes for now.
func GetCalendarsBuilder(
	ctx context.Context,
	gs graph.Servicer,
	userID string,
	optionalFields ...string,
) (
	*users.ItemCalendarsRequestBuilder,
	*users.ItemCalendarsRequestBuilderGetRequestConfiguration,
	error,
) {
	ofcf, err := optionsForCalendars(optionalFields)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "options for event calendars: %v", optionalFields)
	}

	builder := gs.Client().
		UsersById(userID).
		Calendars()

	return builder, ofcf, nil
}

// FetchEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
func FetchEventIDsFromCalendar(
	ctx context.Context,
	gs graph.Servicer,
	user, calendarID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForEventsByCalendar([]string{"id"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	builder := gs.Client().
		UsersById(user).
		CalendarsById(calendarID).
		Events()

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

		builder = users.NewItemCalendarsItemEventsRequestBuilder(*nextLink, gs.Adapter())
	}

	// Events don't have a delta endpoint so just return an empty string.
	return ids, nil, DeltaUpdate{}, errs.ErrorOrNil()
}
