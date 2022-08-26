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

const (
	mailCategory     = "mail"
	contactsCategory = "contacts"
	eventsCategory   = "events"
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
	tenant string,
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	graphStatusChannel chan<- *support.ConnectorOperationStatus,
) func(any) bool

// IterateSelectAllMessageForCollection utility function for
// Iterating through MessagesCollectionResponse
// During iteration, messages belonging to any folder are
// placed into a Collection based on the parent folder
func IterateSelectAllMessagesForCollections(
	tenant string,
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	return func(messageItem any) bool {
		// Defines the type of collection being created within the function
		collectionType := messages
		user := scope.Get(selectors.ExchangeUser)[0]

		message, ok := messageItem.(models.Messageable)
		if !ok {
			errs = support.WrapAndAppendf(user, errors.New("message iteration failure"), errs)
			return true
		}
		// Saving to messages to list. Indexed by folder
		directory := *message.GetParentFolderId()
		if _, ok = collections[directory]; !ok {
			service, err := createService(credentials, failFast)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return true
			}
			edc := NewCollection(
				user,
				[]string{tenant, user, mailCategory, directory},
				collectionType,
				service,
				statusCh,
			)
			collections[directory] = &edc
		}
		collections[directory].AddJob(*message.GetId())
		return true
	}
}

// IterateSelectAllEventsForCollections
// utility function for iterating through events
// and storing events in collections based on
// the calendarID which originates from M365.
func IterateSelectAllEventsForCollections(
	tenant string,
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
				[]string{tenant, user, eventsCategory, directory},
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

// IterateAllContactsForCollection GraphIterateFunc for moving through
// a ContactsCollectionsResponse using the msgraphgocore paging interface.
// Contacts Ids are placed into a collection based upon the parent folder
func IterateAllContactsForCollection(
	tenant string,
	user string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	return func(contactsItem any) bool {
		contact, ok := contactsItem.(models.Contactable)
		if !ok {
			errs = support.WrapAndAppend(user, errors.New("contact iteration failure"), errs)
			return true
		}
		directory := *contact.GetParentFolderId()
		if _, ok := collections[directory]; !ok {
			service, err := createService(credentials, failFast)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return true
			}
			edc := NewCollection(
				user,
				[]string{tenant, user, contactsCategory, directory},
				contacts,
				service,
				statusCh,
			)
			collections[directory] = &edc
		}
		collections[directory].AddJob(*contact.GetId())
		return true
	}
}

// IterateAndFilterMessagesForCollections is a filtering GraphIterateFunc
// that places exchange mail message ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
func IterateAndFilterMessagesForCollections(
	tenant string,
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
				tenant,
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

		message, ok := messageItem.(models.Messageable)
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
	tenant string,
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
		folder, ok := folderItem.(models.MailFolderable)
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
		if !scope.Contains(selectors.ExchangeMailFolder, *folder.GetDisplayName()) {
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
			[]string{tenant, user, mailCategory, directory},
			messages,
			service,
			statusCh,
		)
		collections[directory] = &temp

		return true
	}
}
