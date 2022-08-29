package exchange

import (
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/selectors"
)

// descendable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a "parent folder".
type descendable interface {
	GetId() *string
	GetParentFolderId() *string
}

// displayable represents objects that implement msgraph-sdk-fo/models.entityable
// and have the concept of a display name.
type displayable interface {
	GetId() *string
	GetDisplayName() *string
}

// GraphIterateFuncs are iterate functions to be used with the M365 iterators (e.g. msgraphgocore.NewPageIterator)
// @returns a callback func that works with msgraphgocore.PageIterator.Iterate function
type GraphIterateFunc func(
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	graphStatusChannel chan<- *support.ConnectorOperationStatus,
) func(any) bool

// IterateSelectAllDescendablesForCollection utility function for
// Iterating through MessagesCollectionResponse or ContactsCollectionResponse,
// objects belonging to any folder are
// placed into a Collection based on the parent folder
func IterateSelectAllDescendablesForCollections(
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	var (
		isCategorySet  bool
		collectionType optionIdentifier
		category       string
	)

	return func(pageItem any) bool {
		// Defines the type of collection being created within the function
		if !isCategorySet {
			if scope.IncludesCategory(selectors.ExchangeMail) {
				collectionType = messages
				category = mailCategory
			}

			if scope.IncludesCategory(selectors.ExchangeContact) {
				collectionType = contacts
				category = contactsCategory
			}

			isCategorySet = true
		}

		entry, ok := pageItem.(descendable)
		if !ok {
			errs = support.WrapAndAppendf(user, errors.New("descendable conversion failure"), errs)
			return true
		}
		// Saving to messages to list. Indexed by folder
		directory := *entry.GetParentFolderId()
		if _, ok = collections[directory]; !ok {
			service, err := createService(credentials, failFast)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return true
			}

			edc := NewCollection(
				user,
				[]string{credentials.TenantID, user, category, directory},
				collectionType,
				service,
				statusCh,
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
func IterateSelectAllEventsForCollections(
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	return func(eventItem any) bool {
		event, ok := eventItem.(models.Eventable)
		if !ok {
			errs = support.WrapAndAppend(
				user,
				errors.New("event iteration failure"),
				errs,
			)

			return true
		}

		adtl := event.GetAdditionalData()

		value, ok := adtl["calendar@odata.associationLink"]
		if !ok {
			errs = support.WrapAndAppend(
				user,
				fmt.Errorf("%s: does not support calendar look up", *event.GetId()),
				errs,
			)

			return true
		}

		link, ok := value.(*string)
		if !ok || link == nil {
			errs = support.WrapAndAppend(
				user,
				fmt.Errorf("%s: unable to obtain calendar event data", *event.GetId()),
				errs,
			)

			return true
		}
		// calendars and events are not easily correlated
		// helper function retrieves calendarID from url
		directory, err := parseCalendarIDFromEvent(*link)
		if err != nil {
			errs = support.WrapAndAppend(
				user,
				errors.Wrap(err, *event.GetId()),
				errs,
			)

			return true
		}

		if _, ok := collections[directory]; !ok {
			service, err := createService(credentials, failFast)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return true
			}

			edc := NewCollection(
				user,
				[]string{credentials.TenantID, user, eventsCategory, directory},
				events,
				service,
				statusCh,
			)
			collections[directory] = &edc
		}

		collections[directory].AddJob(*event.GetId())

		return true
	}
}

// IterateAndFilterMessagesForCollections is a filtering GraphIterateFunc
// that places exchange mail message ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
func IterateAndFilterMessagesForCollections(
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	var isFilterSet bool

	return func(messageItem any) bool {
		if !isFilterSet {
			err := CollectMailFolders(
				scope,
				user,
				collections,
				credentials,
				failFast,
				statusCh,
			)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return false
			}

			isFilterSet = true
		}

		message, ok := messageItem.(descendable)
		if !ok {
			errs = support.WrapAndAppend(user, errors.New("message iteration failure"), errs)
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
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	var (
		service graph.Service
		err     error
	)

	return func(folderItem any) bool {
		folder, ok := folderItem.(displayable)
		if !ok {
			errs = support.WrapAndAppend(
				user,
				errors.New("unable to transform folderable item"),
				errs,
			)

			return true
		}
		// Continue to iterate if folder name is empty
		if folder.GetDisplayName() == nil {
			return true
		}
		if !scope.Matches(selectors.ExchangeMailFolder, *folder.GetDisplayName()) {
			return true
		}

		directory := *folder.GetId()

		service, err = createService(credentials, failFast)
		if err != nil {
			errs = support.WrapAndAppend(
				*folder.GetDisplayName(),
				errors.Wrap(
					err,
					"unable to create service a folder query service for "+user,
				),
				errs,
			)

			return true
		}

		temp := NewCollection(
			user,
			[]string{credentials.TenantID, user, mailCategory, directory},
			messages,
			service,
			statusCh,
		)
		collections[directory] = &temp

		return true
	}
}

// iterateFindFolderID is a utility function that supports finding
// M365 folders objects that matches the folderName. Iterator callback function
// will work on folderCollection responses whose objects implement
// the displayable interface. If folder exists, the function updates the
// folderID memory address that was passed in.
func iterateFindFolderID(
	category optionIdentifier,
	folderID **string,
	folderName, errorIdentifier string,
	errs error,
) func(any) bool {
	return func(entry any) bool {
		switch category {
		case messages, contacts:
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

			name := *folder.GetDisplayName()
			if folderName == name {
				if folder.GetId() == nil {
					return true // invalid folder
				}

				*folderID = folder.GetId()

				return false
			}

			return true

		default:
			return false
		}
	}
}
