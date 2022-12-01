package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	cdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/contacts/delta"
	mdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/messages/delta"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

const nextLinkKey = "@odata.nextLink"

// FilterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func FilterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.ExchangeScope,
) error {
	var (
		collectionType = CategoryToOptionIdentifier(scope.Category().PathType())
		errs           error
	)

	for _, c := range resolver.Items() {
		dirPath, ok := pathAndMatch(qp, c, scope)
		if ok {
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
		}
	}

	for directoryID, col := range collections {
		fetchFunc, err := getFetchIDFunc(scope.Category().PathType())
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

		jobs, err := fetchFunc(ctx, col.service, qp.ResourceOwner, directoryID)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.ResourceOwner,
				err,
				errs,
			)
		}

		col.jobs = append(col.jobs, jobs...)
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

// FetchIDFunc collection of helper functions which return a list of strings
// from a response.
type FetchIDFunc func(ctx context.Context, gs graph.Service, user, containerID string) ([]string, error)

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
) ([]string, error) {
	ids := []string{}

	response, err := gs.Client().
		UsersById(user).
		CalendarsById(calendarID).
		Events().Get(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateEventCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return nil, errors.Wrap(err, "iterator creation failure during fetchEventIDs")
	}

	var errs *multierror.Error

	err = pageIterator.Iterate(ctx, func(pageItem any) bool {
		entry, ok := pageItem.(graph.Idable)
		if !ok {
			errs = multierror.Append(errs, errors.New("item without GetId() call"))
			return true
		}

		if entry.GetId() == nil {
			errs = multierror.Append(errs, errors.New("item with nil ID"))
			return true
		}

		ids = append(ids, *entry.GetId())

		return true
	})

	if err != nil {
		return nil, errors.Wrap(
			err,
			support.ConnectorStackErrorTrace(err)+
				" :fetching events from calendar "+calendarID,
		)
	}

	return ids, errs.ErrorOrNil()
}

// FetchContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func FetchContactIDsFromDirectory(ctx context.Context, gs graph.Service, user, directoryID string) ([]string, error) {
	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForContactFoldersItem([]string{"parentFolderId"})
	if err != nil {
		return nil, errors.Wrap(err, "getting query options")
	}

	builder := gs.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Delta()

	for {
		// TODO(ashmrtn): Update to pass options once graph SDK dependency is updated.
		resp, err := sendContactsDeltaGet(ctx, builder, options, gs.Adapter())
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
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

		nextLinkIface := resp.GetAdditionalData()[nextLinkKey]
		if nextLinkIface == nil {
			break
		}

		nextLink := nextLinkIface.(*string)
		if len(*nextLink) == 0 {
			break
		}

		builder = cdelta.NewDeltaRequestBuilder(*nextLink, gs.Adapter())
	}

	return ids, errs.ErrorOrNil()
}

// FetchMessageIDsFromDirectory function that returns a list of  all the m365IDs of the exchange.Mail
// of the targeted directory
func FetchMessageIDsFromDirectory(
	ctx context.Context,
	gs graph.Service,
	user, directoryID string,
) ([]string, error) {
	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForFolderMessages([]string{"id"})
	if err != nil {
		return nil, errors.Wrap(err, "getting query options")
	}

	builder := gs.Client().
		UsersById(user).
		MailFoldersById(directoryID).
		Messages().
		Delta()

	for {
		// TODO(ashmrtn): Update to pass options once graph SDK dependency is updated.
		resp, err := sendMessagesDeltaGet(ctx, builder, options, gs.Adapter())
		if err != nil {
			return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
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

		nextLinkIface := resp.GetAdditionalData()[nextLinkKey]
		if nextLinkIface == nil {
			break
		}

		nextLink := nextLinkIface.(*string)
		if len(*nextLink) == 0 {
			break
		}

		builder = mdelta.NewDeltaRequestBuilder(*nextLink, gs.Adapter())
	}

	return ids, errs.ErrorOrNil()
}
