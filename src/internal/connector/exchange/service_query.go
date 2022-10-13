package exchange

import (
	"context"
	"fmt"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
)

// GraphQuery represents functions which perform exchange-specific queries
// into M365 backstore. Responses -> returned items will only contain the information
// that is included in the options
// TODO: use selector or path for granularity into specific folders or specific date ranges
type GraphQuery func(ctx context.Context, gs graph.Service, userID string) (absser.Parsable, error)

// GetAllMessagesForUser is a GraphQuery function for receiving all messages for a single user
func GetAllMessagesForUser(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	selecting := []string{"id", "parentFolderId"}

	options, err := optionsForMessages(selecting)
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Messages().Get(ctx, options)
}

// GetAllContactsForUser is a GraphQuery function for querying all the contacts in a user's account
func GetAllContactsForUser(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	selecting := []string{"parentFolderId"}

	options, err := optionsForContacts(selecting)
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Contacts().Get(ctx, options)
}

// GetAllFolderDisplayNamesForUser is a GraphQuery function for getting FolderId and display
// names for Mail Folder. All other information for the MailFolder object is omitted.
func GetAllFolderNamesForUser(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForMailFolders([]string{"displayName"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).MailFolders().Get(ctx, options)
}

func GetAllCalendarNamesForUser(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForCalendars([]string{"name", "owner"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Calendars().Get(ctx, options)
}

// GetAllContactFolderNamesForUser is a GraphQuery function for getting ContactFolderId
// and display names for contacts. All other information is omitted.
// Does not return the primary Contact Folder
func GetAllContactFolderNamesForUser(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForContactFolders([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).ContactFolders().Get(ctx, options)
}

// GetAllUsersForTenant is a GraphQuery for retrieving all the UserCollectionResponse with
// that contains the UserID and email for each user. All other information is omitted
func GetAllUsersForTenant(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	selecting := []string{"userPrincipalName"}

	options, err := optionsForUsers(selecting)
	if err != nil {
		return nil, err
	}

	return gs.Client().Users().Get(ctx, options)
}

// GetAllEvents for User. Default returns EventResponseCollection for future events.
// of the time that the call was made. 'calendar' option must be present to gain
// access to additional data map in future calls.
func GetAllEventsForUser(ctx context.Context, gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForEvents([]string{"id", "calendar"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Events().Get(ctx, options)
}

// GraphRetrievalFunctions are functions from the Microsoft Graph API that retrieve
// the default associated data of a M365 object. This varies by object. Additional
// Queries must be run to obtain the omitted fields.
type GraphRetrievalFunc func(ctx context.Context, gs graph.Service, user, m365ID string) (absser.Parsable, error)

// RetrieveContactDataForUser is a GraphRetrievalFun that returns all associated fields.
func RetrieveContactDataForUser(ctx context.Context, gs graph.Service, user, m365ID string) (absser.Parsable, error) {
	return gs.Client().UsersById(user).ContactsById(m365ID).Get(ctx, nil)
}

// RetrieveEventDataForUser is a GraphRetrievalFunc that returns event data.
// Calendarable and attachment fields are omitted due to size
func RetrieveEventDataForUser(ctx context.Context, gs graph.Service, user, m365ID string) (absser.Parsable, error) {
	return gs.Client().UsersById(user).EventsById(m365ID).Get(ctx, nil)
}

// RetrieveMessageDataForUser is a GraphRetrievalFunc that returns message data.
// Attachment field is omitted due to size.
func RetrieveMessageDataForUser(ctx context.Context, gs graph.Service, user, m365ID string) (absser.Parsable, error) {
	return gs.Client().UsersById(user).MessagesById(m365ID).Get(ctx, nil)
}

// CollectFolders is a utility function for creating Collections based off parameters found
// in the ExchangeScope found in the graph.QueryParams
// TODO(ashmrtn): This may not need to do the query if we decide the cache
// should always:
//   1. be passed in
//   2. be populated with all folders for the user
func CollectFolders(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
) error {
	var (
		query             GraphQuery
		transformer       absser.ParsableFactory
		queryService, err = createService(qp.Credentials, qp.FailFast)
	)

	if err != nil {
		return errors.Wrapf(
			err,
			"unable to create graph.Service within CollectFolders service for "+qp.User,
		)
	}

	option := scopeToOptionIdentifier(qp.Scope)
	switch option {
	case messages:
		query = GetAllFolderNamesForUser
		transformer = models.CreateMailFolderCollectionResponseFromDiscriminatorValue
	case contacts:
		query = GetAllContactFolderNamesForUser
		transformer = models.CreateContactFolderCollectionResponseFromDiscriminatorValue
	case events:
		query = GetAllCalendarNamesForUser
		transformer = models.CreateCalendarCollectionResponseFromDiscriminatorValue
	default:
		return fmt.Errorf("unsupported option %s used in CollectFolders", option)
	}

	response, err := query(ctx, queryService, qp.User)
	if err != nil {
		return fmt.Errorf(
			"unable to query mail folder for %s: details: %s",
			qp.User,
			support.ConnectorStackErrorTrace(err),
		)
	}

	// Iterator required to ensure all potential folders are inspected
	// when the breadth of the folder space is large
	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		&queryService.adapter,
		transformer)
	if err != nil {
		return errors.Wrap(err, "unable to create iterator during mail folder query service")
	}

	errUpdater := func(id string, e error) {
		err = support.WrapAndAppend(id, e, err)
	}

	callbackFunc := IterateFilterContainersForCollections(
		ctx,
		qp,
		errUpdater,
		collections,
		statusUpdater,
		resolver,
	)

	iterateFailure := pageIterator.Iterate(ctx, callbackFunc)
	if iterateFailure != nil {
		err = support.WrapAndAppend(qp.User+" iterate failure", iterateFailure, err)
	}

	return err
}
