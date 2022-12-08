package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	nextLinkKey  = "@odata.nextLink"
	deltaLinkKey = "@odata.deltaLink"
)

// getAdditionalDataString gets a string value from the AdditionalData map. If
// the value is not in the map returns an empty string.
func getAdditionalDataString(
	key string,
	addtlData map[string]any,
) string {
	iface := addtlData[key]
	if iface == nil {
		return ""
	}

	value, ok := iface.(*string)
	if !ok {
		return ""
	}

	return *value
}

// FilterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func FilterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]data.Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.ExchangeScope,
) error {
	var (
		errs           error
		collectionType = CategoryToOptionIdentifier(qp.Category)
	)

	for _, c := range resolver.Items() {
		dirPath, ok := pathAndMatch(qp, c, scope)
		if !ok {
			continue
		}

		// Create only those that match
		service, err := createService(qp.Credentials, qp.FailFast)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.ResourceOwner+" FilterContainerAndFillCollection",
				err,
				errs)

			if qp.FailFast {
				return errs
			}
		}

		edc := NewCollection(
			qp.ResourceOwner,
			dirPath,
			collectionType,
			service,
			statusUpdater,
		)
		collections[*c.GetId()] = &edc

		fetchFunc, err := getFetchIDFunc(qp.Category)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.ResourceOwner,
				err,
				errs)

			if qp.FailFast {
				return errs
			}

			continue
		}

		jobs, _, err := fetchFunc(ctx, edc.service, qp.ResourceOwner, *c.GetId())
		if err != nil {
			errs = support.WrapAndAppend(
				qp.ResourceOwner,
				err,
				errs,
			)
		}

		edc.jobs = append(edc.jobs, jobs...)
	}

	return errs
}

func IterativeCollectContactContainers(
	containers map[string]graph.Container,
	nameContains string,
	errUpdater func(string, error),
) func(any) bool {
	return func(entry any) bool {
		folder, ok := entry.(models.ContactFolderable)
		if !ok {
			errUpdater("iterateCollectContactContainers",
				errors.New("casting item to models.ContactFolderable"))
			return false
		}

		include := len(nameContains) == 0 ||
			strings.Contains(*folder.GetDisplayName(), nameContains)

		if include {
			containers[*folder.GetDisplayName()] = folder
		}

		return true
	}
}

func IterativeCollectCalendarContainers(
	containers map[string]graph.Container,
	nameContains string,
	errUpdater func(string, error),
) func(any) bool {
	return func(entry any) bool {
		cal, ok := entry.(models.Calendarable)
		if !ok {
			errUpdater("iterativeCollectCalendarContainers",
				errors.New("casting item to models.Calendarable"))
			return false
		}

		include := len(nameContains) == 0 ||
			strings.Contains(*cal.GetName(), nameContains)
		if include {
			temp := CreateCalendarDisplayable(cal)
			containers[*temp.GetDisplayName()] = temp
		}

		return true
	}
}

// FetchIDFunc collection of helper functions which return a list of all item
// IDs in the given container and a delta token for future requests if the
// container supports fetching delta records.
type FetchIDFunc func(
	ctx context.Context,
	gs graph.Service,
	user, containerID string,
) ([]string, string, error)

func getFetchIDFunc(category path.CategoryType) (FetchIDFunc, error) {
	switch category {
	case path.EmailCategory:
		return FetchMessageIDsFromDirectory, nil
	case path.EventsCategory:
		return FetchEventIDsFromCalendar, nil
	case path.ContactsCategory:
		return FetchContactIDsFromDirectory, nil
	default:
		return nil, fmt.Errorf("category %s not supported by getFetchIDFunc", category)
	}
}

// FetchEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
func FetchEventIDsFromCalendar(
	ctx context.Context,
	gs graph.Service,
	user, calendarID string,
) ([]string, string, error) {
	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForEventsByCalendar([]string{"id"})
	if err != nil {
		return nil, "", err
	}

	builder := gs.Client().
		UsersById(user).
		CalendarsById(calendarID).
		Events()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return nil, "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
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

		builder = msuser.NewUsersItemCalendarsItemEventsRequestBuilder(*nextLink, gs.Adapter())
	}

	// Events don't have a delta endpoint so just return an empty string.
	return ids, "", errs.ErrorOrNil()
}

// FetchContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func FetchContactIDsFromDirectory(
	ctx context.Context,
	gs graph.Service,
	user, directoryID string,
) ([]string, string, error) {
	var (
		errs       *multierror.Error
		ids        []string
		deltaToken string
	)

	options, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"})
	if err != nil {
		return nil, deltaToken, errors.Wrap(err, "getting query options")
	}

	builder := gs.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Delta()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return nil, deltaToken, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, item := range resp.GetValue() {
			if item.GetId() == nil {
				errs = multierror.Append(
					errs,
					errors.Errorf("contact with nil ID in folder %s", directoryID),
				)

				// TODO(ashmrtn): Handle fail-fast.
				continue
			}

			ids = append(ids, *item.GetId())
		}

		addtlData := resp.GetAdditionalData()

		delta := getAdditionalDataString(deltaLinkKey, addtlData)
		if len(delta) > 0 {
			deltaToken = delta
		}

		nextLink := getAdditionalDataString(nextLinkKey, addtlData)
		if len(nextLink) == 0 {
			break
		}

		builder = msuser.NewUsersItemContactFoldersItemContactsDeltaRequestBuilder(nextLink, gs.Adapter())
	}

	return ids, deltaToken, errs.ErrorOrNil()
}

// FetchMessageIDsFromDirectory function that returns a list of  all the m365IDs of the exchange.Mail
// of the targeted directory
func FetchMessageIDsFromDirectory(
	ctx context.Context,
	gs graph.Service,
	user, directoryID string,
) ([]string, string, error) {
	var (
		errs       *multierror.Error
		ids        []string
		deltaToken string
	)

	options, err := optionsForFolderMessagesDelta([]string{"id"})
	if err != nil {
		return nil, deltaToken, errors.Wrap(err, "getting query options")
	}

	builder := gs.Client().
		UsersById(user).
		MailFoldersById(directoryID).
		Messages().
		Delta()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return nil, deltaToken, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, item := range resp.GetValue() {
			if item.GetId() == nil {
				errs = multierror.Append(
					errs,
					errors.Errorf("item with nil ID in folder %s", directoryID),
				)

				// TODO(ashmrtn): Handle fail-fast.
				continue
			}

			ids = append(ids, *item.GetId())
		}

		addtlData := resp.GetAdditionalData()

		delta := getAdditionalDataString(deltaLinkKey, addtlData)
		if len(delta) > 0 {
			deltaToken = delta
		}

		nextLink := getAdditionalDataString(nextLinkKey, addtlData)
		if len(nextLink) == 0 {
			break
		}

		builder = msuser.NewUsersItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, gs.Adapter())
	}

	return ids, deltaToken, errs.ErrorOrNil()
}
