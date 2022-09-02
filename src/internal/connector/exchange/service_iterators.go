package exchange

import (
	"context"
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/path"
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
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	graphStatusChannel chan<- *support.ConnectorOperationStatus,
) func(any) bool

// IterateSelectAllDescendablesForCollection utility function for
// Iterating through MessagesCollectionResponse or ContactsCollectionResponse,
// objects belonging to any folder are
// placed into a Collection based on the parent folder
func IterateSelectAllDescendablesForCollections(
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	var (
		isCategorySet  bool
		collectionType optionIdentifier
		category       path.CategoryType
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

			isCategorySet = true
		}

		entry, ok := pageItem.(descendable)
		if !ok {
			errs = support.WrapAndAppendf(qp.User, errors.New("descendable conversion failure"), errs)
			return true
		}
		// Saving to messages to list. Indexed by folder
		directory := *entry.GetParentFolderId()
		if _, ok = collections[directory]; !ok {
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errs = support.WrapAndAppend(qp.User, err, errs)
				return true
			}

			edc := NewCollection(
				qp.User,
				[]string{qp.Credentials.TenantID, qp.User, category.String(), directory},
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
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	return func(eventItem any) bool {
		event, ok := eventItem.(models.Eventable)
		if !ok {
			errs = support.WrapAndAppend(
				qp.User,
				errors.New("event iteration failure"),
				errs,
			)

			return true
		}

		adtl := event.GetAdditionalData()

		value, ok := adtl["calendar@odata.associationLink"]
		if !ok {
			errs = support.WrapAndAppend(
				qp.User,
				fmt.Errorf("%s: does not support calendar look up", *event.GetId()),
				errs,
			)

			return true
		}

		link, ok := value.(*string)
		if !ok || link == nil {
			errs = support.WrapAndAppend(
				qp.User,
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
				qp.User,
				errors.Wrap(err, *event.GetId()),
				errs,
			)

			return true
		}

		if _, ok := collections[directory]; !ok {
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errs = support.WrapAndAppend(qp.User, err, errs)
				return true
			}

			edc := NewCollection(
				qp.User,
				[]string{qp.Credentials.TenantID, qp.User, path.EventsCategory.String(), directory},
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
	ctx context.Context,
	qp graph.QueryParams,
	errs error,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	var isFilterSet bool

	return func(messageItem any) bool {
		if !isFilterSet {
			err := CollectMailFolders(
				ctx,
				qp,
				collections,
				statusCh,
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
			[]string{qp.Credentials.TenantID, qp.User, path.EmailCategory.String(), directory},
			messages,
			service,
			statusCh,
		)
		collections[directory] = &temp

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

			if entry == nil {
				return true
			}
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
