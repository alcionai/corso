package exchange

import (
	"context"

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
	errs error,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool

// maybeGetAndPopulateFolderResolver gets a folder resolver if one is available for
// this category of data. If one is not available, returns nil so that other
// logic in the caller can complete as long as they check if the resolver is not
// nil. If an error occurs populating the resolver, returns an error.
func maybeGetAndPopulateFolderResolver(
	ctx context.Context,
	qp graph.QueryParams,
	category path.CategoryType,
) (graph.ContainerResolver, error) {
	var res graph.ContainerResolver

	switch category {
	case path.EmailCategory:
		service, err := createService(qp.Credentials, qp.FailFast)
		if err != nil {
			return nil, err
		}

		res = &mailFolderCache{
			userID: qp.User,
			gs:     service,
		}

	default:
		return nil, nil
	}

	if err := res.Populate(ctx); err != nil {
		return nil, errors.Wrap(err, "populating directory resolver")
	}

	return res, nil
}

func resolveCollectionPath(
	ctx context.Context,
	resolver graph.ContainerResolver,
	tenantID, user, folderID string,
	category path.CategoryType,
) (path.Path, error) {
	if resolver == nil {
		// Allows caller to default to old-style path.
		return nil, errors.WithStack(errNilResolver)
	}

	p, err := resolver.IDToPath(ctx, folderID)
	if err != nil {
		return nil, errors.Wrap(err, "resolving folder ID")
	}

	return p.ToDataLayerExchangePathForCategory(
		tenantID,
		user,
		category,
		false,
	)
}

// IterateSelectAllDescendablesForCollection utility function for
// Iterating through MessagesCollectionResponse or ContactsCollectionResponse,
// objects belonging to any folder are
// placed into a Collection based on the parent folder
func IterateSelectAllDescendablesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		isCategorySet  bool
		collectionType optionIdentifier
		category       path.CategoryType
		resolver       graph.ContainerResolver
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
				errs = support.WrapAndAppend(
					"getting folder resolver for category "+category.String(),
					err,
					errs,
				)
			} else {
				resolver = r
			}

			isCategorySet = true
		}

		entry, ok := pageItem.(descendable)
		if !ok {
			errs = support.WrapAndAppendf(qp.User, errors.New("descendable conversion failure"), errs)
			return true
		}

		// Saving to messages to list. Indexed by folder
		directory := *entry.GetParentFolderId()

		dirPath, err := path.Builder{}.Append(directory).ToDataLayerExchangePathForCategory(
			qp.Credentials.TenantID,
			qp.User,
			category,
			false,
		)
		if err != nil {
			errs = support.WrapAndAppend("converting to resource path", err, errs)
			// This really shouldn't be happening unless we have a bad category.
			return true
		}

		if _, ok = collections[directory]; !ok {
			newPath, err := resolveCollectionPath(
				ctx,
				resolver,
				qp.Credentials.TenantID,
				qp.User,
				directory,
				category,
			)

			if err != nil {
				if !errors.Is(err, errNilResolver) {
					errs = support.WrapAndAppend("", err, errs)
				}
			} else {
				dirPath = newPath
			}

			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errs = support.WrapAndAppend(qp.User, err, errs)
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
func IterateSelectAllEventsFromCalendars(
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	return func(pageItem any) bool {
		if pageItem == nil {
			return true
		}

		shell, ok := pageItem.(models.Calendarable)
		if !ok {
			errs = support.WrapAndAppend(
				qp.User,
				errors.New("calendar event"),
				errs)

			return true
		}

		service, err := createService(qp.Credentials, qp.FailFast)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				errors.Wrap(err, "unable to create service during IterateSelectAllEventsFromCalendars"),
				errs,
			)

			return true
		}

		eventResponseable, err := service.Client().
			UsersById(qp.User).
			CalendarsById(*shell.GetId()).
			Events().Get()
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs,
			)
		}

		directory := shell.GetName()
		owner := shell.GetOwner()

		// Conditional Guard Checks
		if eventResponseable == nil ||
			directory == nil ||
			owner == nil {
			return true
		}

		eventables := eventResponseable.GetValue()
		// Clause is true when Calendar has does not have any events
		if eventables == nil {
			return true
		}

		if _, ok := collections[*directory]; !ok {
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errs = support.WrapAndAppend(qp.User, err, errs)

				return true
			}

			dirPath, err := path.Builder{}.Append(*directory).ToDataLayerExchangePathForCategory(
				qp.Credentials.TenantID,
				qp.User,
				path.EventsCategory,
				false,
			)
			if err != nil {
				// This really shouldn't be happening.
				errs = support.WrapAndAppend("converting to resource path", err, errs)
				return true
			}

			edc := NewCollection(
				qp.User,
				dirPath,
				events,
				service,
				statusUpdater,
			)
			collections[*directory] = &edc
		}

		for _, event := range eventables {
			collections[*directory].AddJob(*event.GetId())
		}

		return true
	}
}

// IterateAndFilterMessagesForCollections is a filtering GraphIterateFunc
// that places exchange mail message ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
func IterateAndFilterMessagesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var isFilterSet bool

	return func(messageItem any) bool {
		if !isFilterSet {
			err := CollectMailFolders(
				ctx,
				qp,
				collections,
				statusUpdater,
			)
			if err != nil {
				errs = support.WrapAndAppend(qp.User, err, errs)
				return false
			}

			isFilterSet = true
		}

		message, ok := messageItem.(descendable)
		if !ok {
			errs = support.WrapAndAppend(qp.User, errors.New("message iteration failure"), errs)
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

func IterateFilterFolderDirectoriesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var (
		service graph.Service
		err     error
	)

	resolver, err := maybeGetAndPopulateFolderResolver(ctx, qp, path.EmailCategory)
	if err != nil {
		errs = support.WrapAndAppend(
			"getting folder resolver for category email",
			err,
			errs,
		)
	}

	return func(folderItem any) bool {
		folder, ok := folderItem.(displayable)
		if !ok {
			errs = support.WrapAndAppend(
				qp.User,
				errors.New("unable to transform folderable item"),
				errs,
			)

			return true
		}
		// Continue to iterate if folder name is empty
		if folder.GetDisplayName() == nil {
			return true
		}

		if !qp.Scope.Matches(selectors.ExchangeMailFolder, *folder.GetDisplayName()) {
			return true
		}

		directory := *folder.GetId()

		dirPath, err := path.Builder{}.Append(directory).ToDataLayerExchangePathForCategory(
			qp.Credentials.TenantID,
			qp.User,
			path.EmailCategory,
			false,
		)
		if err != nil {
			// This really shouldn't be happening.
			errs = support.WrapAndAppend("converting to resource path", err, errs)
			return true
		}

		p, err := resolveCollectionPath(
			ctx,
			resolver,
			qp.Credentials.TenantID,
			qp.User,
			directory,
			path.EmailCategory,
		)

		if err != nil {
			if !errors.Is(err, errNilResolver) {
				errs = support.WrapAndAppend("", err, errs)
			}
		} else {
			dirPath = p
		}

		service, err = createService(qp.Credentials, qp.FailFast)
		if err != nil {
			errs = support.WrapAndAppend(
				*folder.GetDisplayName(),
				errors.Wrap(
					err,
					"unable to create service a folder query service for "+qp.User,
				),
				errs,
			)

			return true
		}

		temp := NewCollection(
			qp.User,
			dirPath,
			messages,
			service,
			statusUpdater,
		)
		collections[directory] = &temp

		return true
	}
}

func IterateSelectAllContactsForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) func(any) bool {
	var isPrimarySet bool

	return func(folderItem any) bool {
		folder, ok := folderItem.(models.ContactFolderable)
		if !ok {
			errs = support.WrapAndAppend(
				qp.User,
				errors.New("casting folderItem to models.ContactFolderable"),
				errs,
			)
		}

		if !isPrimarySet {
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errs = support.WrapAndAppend(
					qp.User,
					errors.Wrap(err, "unable to create service during IterateSelectAllContactsForCollections"),
					errs,
				)

				return true
			}

			contactIDS, err := ReturnContactIDsFromDirectory(service, qp.User, *folder.GetParentFolderId())
			if err != nil {
				errs = support.WrapAndAppend(
					qp.User,
					err,
					errs,
				)

				return true
			}

			dirPath, err := path.Builder{}.Append(*folder.GetParentFolderId()).ToDataLayerExchangePathForCategory(
				qp.Credentials.TenantID,
				qp.User,
				path.ContactsCategory,
				false,
			)
			if err != nil {
				errs = support.WrapAndAppend(
					qp.User,
					err,
					errs,
				)

				return true
			}

			edc := NewCollection(
				qp.User,
				dirPath,
				contacts,
				service,
				statusUpdater,
			)
			edc.jobs = append(edc.jobs, contactIDS...)
			collections["Contacts"] = &edc
			isPrimarySet = true
		}

		service, err := createService(qp.Credentials, qp.FailFast)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs,
			)

			return true
		}

		folderID := *folder.GetId()

		listOfIDs, err := ReturnContactIDsFromDirectory(service, qp.User, folderID)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs,
			)

			return true
		}

		if folder.GetDisplayName() == nil ||
			listOfIDs == nil {
			return true // Invalid state TODO: How should this be named
		}

		dirPath, err := path.Builder{}.Append(*folder.GetId()).ToDataLayerExchangePathForCategory(
			qp.Credentials.TenantID,
			qp.User,
			path.EventsCategory,
			false,
		)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs,
			)
			return true
		}

		edc := NewCollection(
			qp.User,
			dirPath,
			contacts,
			service,
			statusUpdater,
		)
		edc.jobs = append(edc.jobs, listOfIDs...)
		collections[*folder.GetId()] = &edc

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
	errs error,
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
			errs = support.WrapAndAppend(
				errorIdentifier,
				errors.New("struct does not implement displayable"),
				errs,
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

// UpdateListFunc collection of functions which return a list of strings
// from a response.
type UpdateListFunc func(gs graph.Service, user, m365ID string) ([]string, error)

func ReturnContactIDsFromDirectory(gs graph.Service, user, m365ID string) ([]string, error) {
	options, err := optionsForContactFoldersItem([]string{""})
	if err != nil {
		return nil, err
	}

	stringArray := []string{}

	response, err := gs.Client().
		UsersById(user).
		ContactFoldersById(m365ID).
		Contacts().
		GetWithRequestConfigurationAndResponseHandler(options, nil)
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

	if iterateErr := pageIterator.Iterate(callbackFunc); iterateErr != nil {
		return nil, iterateErr
	}

	if err != nil {
		return nil, err
	}

	return stringArray, nil
}
