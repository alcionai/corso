package exchange

import (
	"context"
	"fmt"

	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var errNilResolver = errors.New("nil resolver")

// GraphIterateFuncs are iterate functions to be used with the M365 iterators (e.g. msgraphgocore.NewPageIterator)
// @returns a callback func that works with msgraphgocore.PageIterator.Iterate function
type GraphIterateFunc func(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool

// IterateSelectAllDescendablesForCollection utility function for
// Iterating through MessagesCollectionResponse or ContactsCollectionResponse,
// objects belonging to any folder are
// placed into a Collection based on the parent folder
func IterateSelectAllDescendablesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		isCategorySet  bool
		collectionType optionIdentifier
		category       path.CategoryType
		resolver       graph.ContainerResolver
		dirPath        path.Path
		err            error
	)

	return func(pageItem any) bool {
		// Defines the type of collection being created within the function
		if !isCategorySet {
			if qp.Scope.IncludesCategory(selectors.ExchangeMail) {
				collectionType = messages
				category = path.EmailCategory
			}

			if qp.Scope.IncludesCategory(selectors.ExchangeContact) {
				collectionType = contacts
				category = path.ContactsCategory
			}

			if r, err := maybeGetAndPopulateFolderResolver(ctx, qp, category); err != nil {
				errUpdater("getting folder resolver for category "+category.String(), err)
			} else {
				resolver = r
			}

			isCategorySet = true
		}

		entry, ok := pageItem.(descendable)
		if !ok {
			errUpdater(qp.User, errors.New("descendable conversion failure"))
			return true
		}

		// Saving to messages to list. Indexed by folder
		directory := *entry.GetParentFolderId()

		if _, ok = collections[directory]; !ok {
			dirPath, err = getCollectionPath(
				ctx,
				qp,
				resolver,
				directory,
				category,
			)
			if err != nil {
				errUpdater(
					"failure during IterateSelectAllDescendablesForCollections",
					err,
				)
			}

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
			collections[directory] = &edc
		}

		collections[directory].AddJob(*entry.GetId())

		return true
	}
}

// IterateSelectAllEventsForCollections
// utility function for iterating through events
// and storing events in collections based on
// the calendarID which originates from M365.
// @param pageItem is a CalendarCollectionResponse possessing two populated fields:
// - id - M365 ID
// - Name - Calendar Name
func IterateSelectAllEventsFromCalendars(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		isEnabled bool
		service   graph.Service
	)

	return func(pageItem any) bool {
		if !isEnabled {
			// Create Collections based on qp.Scope
			err := CollectFolders(ctx, qp, collections, statusUpdater)
			if err != nil {
				errUpdater(
					qp.User,
					errors.Wrap(err, support.ConnectorStackErrorTrace(err)),
				)

				return false
			}

			service, err = createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errUpdater(qp.User, err)
				return false
			}

			isEnabled = true
		}

		pageItem = CreateCalendarDisplayable(pageItem)

		calendar, ok := pageItem.(displayable)
		if !ok {
			errUpdater(
				qp.User,
				fmt.Errorf("unable to parse pageItem into CalendarDisplayable: %T", pageItem),
			)
		}

		if calendar.GetDisplayName() == nil {
			return true
		}

		collection, ok := collections[*calendar.GetDisplayName()]
		if !ok {
			return true
		}

		eventIDs, err := ReturnEventIDsFromCalendar(ctx, service, qp.User, *calendar.GetId())
		if err != nil {
			errUpdater(
				qp.User,
				errors.Wrap(err, support.ConnectorStackErrorTrace(err)))

			return true
		}

		collection.jobs = append(collection.jobs, eventIDs...)

		return true
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
	var isFilterSet bool

	return func(descendItem any) bool {
		if !isFilterSet {
			err := CollectFolders(
				ctx,
				qp,
				collections,
				statusUpdater,
			)
			if err != nil {
				errUpdater(qp.User, err)
				return false
			}

			isFilterSet = true
		}

		message, ok := descendItem.(descendable)
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

func getCategoryAndValidation(es *selectors.ExchangeScope) (
	optionIdentifier,
	path.CategoryType,
	func(namePtr *string) bool) {
	var (
		option   = scopeToOptionIdentifier(*es)
		category path.CategoryType
		validate func(namePtr *string) bool
	)

	switch option {
	case messages:
		category = path.EmailCategory
		validate = func(namePtr *string) bool {
			if namePtr == nil {
				return true
			}

			return !es.Matches(selectors.ExchangeMailFolder, *namePtr)
		}
	case contacts:
		category = path.ContactsCategory
		validate = func(namePtr *string) bool {
			if namePtr == nil {
				return true
			}

			return !es.Matches(selectors.ExchangeContactFolder, *namePtr)
		}
	case events:
		category = path.EventsCategory
		validate = func(namePtr *string) bool {
			if namePtr == nil {
				return true
			}

			return !es.Matches(selectors.ExchangeEventCalendar, *namePtr)
		}
	}

	return option, category, validate
}

func IterateFilterContainersForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		resolver    graph.ContainerResolver
		isSet       bool
		collectPath string
		err         error
		option      optionIdentifier
		category    path.CategoryType
		validate    func(*string) bool
	)

	return func(folderItem any) bool {
		if !isSet {
			option, category, validate = getCategoryAndValidation(&qp.Scope)

			resolver, err = maybeGetAndPopulateFolderResolver(ctx, qp, category)
			if err != nil {
				errUpdater("getting folder resolver for category "+category.String(), err)
			}

			isSet = true
		}

		if option == events {
			folderItem = CreateCalendarDisplayable(folderItem)
		}

		folder, ok := folderItem.(displayable)
		if !ok {
			errUpdater(qp.User,
				fmt.Errorf("unable to convert input of %T for category: %s", folderItem, category.String()),
			)
		}

		if validate(folder.GetDisplayName()) {
			return true
		}

		if option == messages {
			collectPath = *folder.GetId()
		} else {
			collectPath = *folder.GetDisplayName()
		}

		dirPath, err := getCollectionPath(
			ctx,
			qp,
			resolver,
			collectPath,
			category,
		)
		if err != nil {
			errUpdater(
				"failure converting path during IterateFilterFolderDirectoriesForCollections",
				err,
			)

			return true
		}

		service, err := createService(qp.Credentials, qp.FailFast)
		if err != nil {
			errUpdater(
				*folder.GetDisplayName(),
				errors.Wrap(err, "creating service to iterate filterFolder directories for user: "+qp.User))

			return true
		}

		temp := NewCollection(
			qp.User,
			dirPath,
			option,
			service,
			statusUpdater,
		)
		collections[*folder.GetDisplayName()] = &temp

		return true
	}
}

func IterateSelectAllContactsForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errUpdater func(string, error),
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		isPrimarySet bool
		service      graph.Service
	)

	return func(folderItem any) bool {
		folder, ok := folderItem.(models.ContactFolderable)
		if !ok {
			errUpdater(
				qp.User,
				errors.New("casting folderItem to models.ContactFolderable"),
			)
		}

		if !isPrimarySet && folder.GetParentFolderId() != nil {
			err := CollectFolders(
				ctx,
				qp,
				collections,
				statusUpdater,
			)
			if err != nil {
				errUpdater(qp.User, err)
				return false
			}

			service, err = createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errUpdater(
					qp.User,
					errors.Wrap(err, "unable to create service during IterateSelectAllContactsForCollections"),
				)

				return true
			}

			// Create and Populate Default Contacts folder Collection if true
			if qp.Scope.Matches(selectors.ExchangeContactFolder, DefaultContactFolder) {
				dirPath, err := path.Builder{}.Append(DefaultContactFolder).ToDataLayerExchangePathForCategory(
					qp.Credentials.TenantID,
					qp.User,
					path.ContactsCategory,
					false,
				)
				if err != nil {
					errUpdater(
						qp.User,
						err,
					)

					return false
				}

				edc := NewCollection(
					qp.User,
					dirPath,
					contacts,
					service,
					statusUpdater,
				)

				listOfIDs, err := ReturnContactIDsFromDirectory(ctx, service, qp.User, *folder.GetParentFolderId())
				if err != nil {
					errUpdater(
						qp.User,
						err,
					)

					return false
				}

				edc.jobs = append(edc.jobs, listOfIDs...)
				collections[DefaultContactFolder] = &edc
				isPrimarySet = true
			}
		}

		if folder.GetDisplayName() == nil {
			// This should never happen. Skipping to avoid kernel panic
			return true
		}

		collection, ok := collections[*folder.GetDisplayName()]
		if !ok {
			return true // Not included
		}

		listOfIDs, err := ReturnContactIDsFromDirectory(ctx, service, qp.User, *folder.GetId())
		if err != nil {
			errUpdater(
				qp.User,
				err,
			)

			return true
		}

		collection.jobs = append(collection.jobs, listOfIDs...)

		return true
	}
}

// iterateFindContainerID is a utility function that supports finding
// M365 folders objects that matches the folderName. Iterator callback function
// will work on folderCollection responses whose objects implement
// the displayable interface. If folder exists, the function updates the
// containerID memory address that was passed in.
// @param containerName is the string representation of the folder, directory or calendar holds
// the underlying M365 objects
func iterateFindContainerID(
	containerID **string,
	containerName, errorIdentifier string,
	isCalendar bool,
	errUpdater func(string, error),
) func(any) bool {
	return func(entry any) bool {
		if isCalendar {
			entry = CreateCalendarDisplayable(entry)
		}

		// True when pagination needs more time to get additional responses or
		// when entry is not able to be converted into a displayable
		if entry == nil {
			return true
		}

		folder, ok := entry.(displayable)
		if !ok {
			errUpdater(
				errorIdentifier,
				errors.New("struct does not implement displayable"),
			)

			return true
		}

		// Display name not set on folder
		if folder.GetDisplayName() == nil {
			return true
		}

		if containerName == *folder.GetDisplayName() {
			if folder.GetId() == nil {
				return true // invalid folder
			}

			*containerID = folder.GetId()

			return false
		}

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
