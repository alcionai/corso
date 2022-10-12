package exchange

import (
	"context"
	"strings"

	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// GraphIterateFuncs are iterate functions to be used with the M365 iterators (e.g. msgraphgocore.NewPageIterator)
// @returns a callback func that works with msgraphgocore.PageIterator.Iterate function
type GraphIterateFunc func(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool

// IterateSelectAllEventsForCollections
// utility function for iterating through events
// and storing events in collections based on
// the calendarID which originates from M365.
// @param pageItem is a CalendarCollectionResponse possessing two populated fields:
// - id - M365 ID
// - Name - Calendar Name
func IterateSelectEventsFromCalendars(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		isEnabled bool
		err       error
		resolver  graph.ContainerResolver
		category  = path.EventsCategory
	)

	return func(pageItem any) bool {
		if !isEnabled {
			resolver, err = PopulateExchangeContainerResolver(ctx, qp, category, false)
			if err != nil {
				errUpdater(qp.User, err)
				return false
			}

			rootPath, ok := checkRoot(qp, category)
			if ok {
				s, err := createService(qp.Credentials, qp.FailFast)
				if err != nil {
					errUpdater(qp.User, err)
					return true
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

			for _, c := range resolver.GetCacheFolders() {
				dirPath, ok := pathAndMatch(qp, category, c)
				if ok {
					// Create only those that match
					service, err := createService(qp.Credentials, qp.FailFast)
					if err != nil {
						errUpdater(qp.User, err)
						return true
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

			isEnabled = true
		}

		// Should be able to complete on the first run
		for key, col := range collections {
			eventIDs, err := ReturnEventIDsFromCalendar(ctx, col.service, qp.User, key)
			if err != nil {
				errUpdater(
					qp.User,
					errors.Wrap(err, support.ConnectorStackErrorTrace(err)))

				continue
			}

			col.jobs = append(col.jobs, eventIDs...)
		}

		return false
	}
}

// IterateAndFilterDescendablesForCollections is a filtering GraphIterateFunc
// that places exchange objectsids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
func IterateAndFilterDescendablesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		isFilterSet    bool
		err            error
		resolver       graph.ContainerResolver
		collectionType = scopeToOptionIdentifier(qp.Scope)
		category       path.CategoryType
	)

	return func(descendItem any) bool {
		if !isFilterSet {
			if qp.Scope.IncludesCategory(selectors.ExchangeMail) {
				category = path.EmailCategory
			}

			if qp.Scope.IncludesCategory(selectors.ExchangeContact) {
				category = path.ContactsCategory
			}

			resolver, err = PopulateExchangeContainerResolver(ctx, qp, category, false)
			if err != nil {
				errUpdater("getting folder resolver for category "+path.EmailCategory.String(), err)
				return false
			}

			if category == path.ContactsCategory {
				rootPath, ok := checkRoot(qp, category)
				if ok {
					s, err := createService(qp.Credentials, qp.FailFast)
					if err != nil {
						errUpdater(qp.User, err)
						return true
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

			for _, c := range resolver.GetCacheFolders() {
				// Create receive all
				dirPath, ok := pathAndMatch(qp, category, c)
				if ok {
					// Create only those that match
					service, err := createService(qp.Credentials, qp.FailFast)
					if err != nil {
						errUpdater(qp.User, err)
						return true
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

			isFilterSet = true
		}

		message, ok := descendItem.(graph.Descendable)
		if !ok {
			errUpdater(qp.User, errors.New("casting messageItem to descendable"))
			return true
		}
		// Saving only messages for the created directories
		directory := *message.GetParentFolderId()

		if _, ok = collections[directory]; !ok {
			return true
		}

		collections[directory].AddJob(*message.GetId())

		return true
	}
}

// IDistFunc collection of helper functions which return a list of strings
// from a response.
type IDListFunc func(ctx context.Context, gs graph.Service, user, m365ID string) ([]string, error)

// ReturnContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func ReturnContactIDsFromDirectory(ctx context.Context, gs graph.Service, user, directoryID string) ([]string, error) {
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
	nameContains, rootID string,
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
			temp := graph.CreateCalendarDisplayable(cal, rootID)
			containers[*temp.GetDisplayName()] = temp
		}

		return true
	}
}

// ReturnEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
func ReturnEventIDsFromCalendar(
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
