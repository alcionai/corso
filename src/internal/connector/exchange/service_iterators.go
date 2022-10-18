package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

// GraphIterateFuncs are iterate functions to be used with the M365 iterators (e.g. msgraphgocore.NewPageIterator)
// @returns a callback func that works with msgraphgocore.PageIterator.Iterate function
type GraphSetCollectionFunc func(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
) error

// IterateSelectAllEventsForCollections
// utility function for iterating through events
// and storing events in collections based on
// the calendarID which originates from M365.
// @param pageItem is a CalendarCollectionResponse possessing two populated fields:
// - id - M365 ID
// - Name - Calendar Name
func FilterEventsFromCalendars(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
) error {
	category := path.EventsCategory

	rootPath, ok := checkRoot(qp, category)
	if ok {
		s, err := createService(qp.Credentials, qp.FailFast)
		if err != nil {
			return errors.Wrap(err, qp.User+" unable to create service FilterEventsFromCalendar")
		}

		concrete := resolver.(*eventCalendarCache)
		edc := NewCollection(
			qp.User,
			rootPath,
			events,
			s,
			statusUpdater,
		)
		collections[concrete.rootID] = &edc
	}

	for _, c := range resolver.Items() {
		dirPath, ok := pathAndMatch(qp, category, c)
		if ok {
			// Create only those that match
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				return errors.Wrap(err, qp.User+" unable to create service FilterEventsFromCalendar")
			}

			edc := NewCollection(
				qp.User,
				dirPath,
				events,
				service,
				statusUpdater,
			)
			collections[*c.GetId()] = &edc
		}
	}

	var errs error
	for key, col := range collections {
		eventIDs, err := FetchEventIDsFromCalendar(ctx, col.service, qp.User, key)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				errors.Wrap(err, support.ConnectorStackErrorTrace(err)),
				errs)

			continue
		}

		col.jobs = append(col.jobs, eventIDs...)
	}

	return errs
}

// IterateAndFilterDescendablesForCollections is a filtering GraphIterateFunc
// that places exchange objectsids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
func FilterDescendablesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
) error {
	category := graph.ScopeToPathCategory(qp.Scope)
	collectionType := categoryToOptionIdentifier(category)

	if category == path.ContactsCategory {
		rootPath, ok := checkRoot(qp, category)
		if ok {
			s, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				return errors.Wrap(err, "failed to create service for collection")
			}

			concrete := resolver.(*contactFolderCache)
			edc := NewCollection(
				qp.User,
				rootPath,
				collectionType,
				s,
				statusUpdater,
			)
			collections[concrete.rootID] = &edc
		}
	}

	for _, c := range resolver.Items() {
		// Create receive all
		dirPath, ok := pathAndMatch(qp, category, c)
		if ok {
			// Create only those that match
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				return errors.Wrap(err, "failed to create service for colleciton")
			}

			edc := NewCollection(
				qp.User,
				dirPath,
				collectionType,
				service,
				statusUpdater,
			)
			collections[*c.GetId()] = &edc
		}
	}

	var errs error
	for directoryID, col := range collections {
		fetchFunc, err := getFetchIDFunc(category)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs)

			if qp.FailFast {
				return errs
			}
			continue
		}
		jobs, err := fetchFunc(ctx, col.service, qp.User, directoryID)
		if err != nil {
			err = errors.Wrap(err, qp.User)
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
			errUpdater("", errors.New("casting item to models.ContactFolderable"))
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
			errUpdater("failure during IterativeCollectCalendarContainers",
				errors.New("casting item to models.Calendarable"))
			return false
		}

		include := len(nameContains) == 0 ||
			strings.Contains(*cal.GetName(), nameContains)
		if include {
			temp := graph.CreateCalendarDisplayable(cal)
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
		return nil, fmt.Errorf("")
	}
}

// ReturnEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
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
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateEventCollectionResponseFromDiscriminatorValue,
	)

	callbackFunc := func(pageItem any) bool {
		entry, ok := pageItem.(models.Eventable)
		if !ok {
			err = errors.New("casting pageItem to models.Eventable")
			return false
		}

		ids = append(ids, *entry.GetId())

		return true
	}

	if iterateErr := pageIterator.Iterate(ctx, callbackFunc); iterateErr != nil {
		return nil,
			errors.Wrap(iterateErr, support.ConnectorStackErrorTrace(err))
	}

	if err != nil {
		return nil, err
	}

	return ids, nil
}

// ReturnContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func FetchContactIDsFromDirectory(ctx context.Context, gs graph.Service, user, directoryID string) ([]string, error) {
	options, err := optionsForContactFoldersItem([]string{"parentFolderId"})
	if err != nil {
		return nil, err
	}

	stringArray := []string{}

	response, err := gs.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Get(ctx, options)
	if err != nil {
		return nil, err
	}

	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateContactCollectionResponseFromDiscriminatorValue,
	)

	callbackFunc := func(pageItem any) bool {
		entry, ok := pageItem.(models.Contactable)
		if !ok {
			err = errors.New("casting pageItem to models.Contactable")
			return false
		}

		stringArray = append(stringArray, *entry.GetId())

		return true
	}

	if iterateErr := pageIterator.Iterate(ctx, callbackFunc); iterateErr != nil {
		return nil, iterateErr
	}

	if err != nil {
		return nil, err
	}

	return stringArray, nil
}

func FetchMessageIDsFromDirectory(
	ctx context.Context,
	gs graph.Service,
	user, directoryID string,
) ([]string, error) {
	stringArray := []string{}

	options, err := optionsForFolderMessages([]string{"id"})
	if err != nil {
		return nil, errors.Wrap(err, "getting query options")
	}

	response, err := gs.Client().
		UsersById(user).
		MailFoldersById(directoryID).
		Messages().
		Get(ctx, options)
	if err != nil {
		return nil, errors.Wrap(
			errors.Wrap(err, support.ConnectorStackErrorTrace(err)),
			"initial folder query",
		)
	}

	pageIter, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateMessageCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating graph iterator")
	}

	var errs *multierror.Error

	err = pageIter.Iterate(ctx, func(pageItem any) bool {
		item, ok := pageItem.(graph.Idable)
		if !ok {
			errs = multierror.Append(errs, errors.New("item without ID function"))
			return true
		}

		if item.GetId() == nil {
			errs = multierror.Append(errs, errors.New("item with nil ID"))
			return true
		}

		stringArray = append(stringArray, *item.GetId())

		return true
	})

	return stringArray, nil
}
