package exchange

import (
	"context"
	"fmt"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
)

// GraphQuery represents functions which perform exchange-specific queries
// into M365 backstore. Responses -> returned items will only contain the information
// that is included in the options
// TODO: use selector or path for granularity into specific folders or specific date ranges
type GraphQuery func(graph.Service, string) (absser.Parsable, error)

// GetAllMessagesForUser is a GraphQuery function for receiving all messages for a single user
func GetAllMessagesForUser(gs graph.Service, user string) (absser.Parsable, error) {
	selecting := []string{"id", "parentFolderId"}

	options, err := optionsForMessages(selecting)
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Messages().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GetAllContactsForUser is a GraphQuery function for querying all the contacts in a user's account
func GetAllContactsForUser(gs graph.Service, user string) (absser.Parsable, error) {
	selecting := []string{"parentFolderId"}

	options, err := optionsForContacts(selecting)
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Contacts().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GetAllFolderDisplayNamesForUser is a GraphQuery function for getting FolderId and display
// names for Mail Folder. All other information for the MailFolder object is omitted.
func GetAllFolderNamesForUser(gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForMailFolders([]string{"displayName"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

func GetAllCalendarNamesForUser(gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Calendars().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GetAllContactFolderNamesForUser is a GraphQuery function for getting ContactFolderId
// and display names for contacts. All other information is omitted.
// Does not return the primary Contact Folder
func GetAllContactFolderNamesForUser(gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForContactFolders([]string{"displayName"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).ContactFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GetAllUsersForTenant is a GraphQuery for retrieving all the UserCollectionResponse with
// that contains the UserID and email for each user. All other information is omitted
func GetAllUsersForTenant(gs graph.Service, user string) (absser.Parsable, error) {
	selecting := []string{"userPrincipalName"}

	options, err := optionsForUsers(selecting)
	if err != nil {
		return nil, err
	}

	return gs.Client().Users().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GetAllEvents for User. Default returns EventResponseCollection for future events.
// of the time that the call was made. 'calendar' option must be present to gain
// access to additional data map in future calls.
func GetAllEventsForUser(gs graph.Service, user string) (absser.Parsable, error) {
	options, err := optionsForEvents([]string{"id", "calendar"})
	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(user).Events().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GraphRetrievalFunctions are functions from the Microsoft Graph API that retrieve
// the default associated data of a M365 object. This varies by object. Additional
// Queries must be run to obtain the omitted fields.
type GraphRetrievalFunc func(gs graph.Service, user, m365ID string) (absser.Parsable, error)

// RetrieveContactDataForUser is a GraphRetrievalFun that returns all associated fields.
func RetrieveContactDataForUser(gs graph.Service, user, m365ID string) (absser.Parsable, error) {
	return gs.Client().UsersById(user).ContactsById(m365ID).Get()
}

// RetrieveEventDataForUser is a GraphRetrievalFunc that returns event data.
// Calendarable and attachment fields are omitted due to size
func RetrieveEventDataForUser(gs graph.Service, user, m365ID string) (absser.Parsable, error) {
	return gs.Client().UsersById(user).EventsById(m365ID).Get()
}

// RetrieveMessageDataForUser is a GraphRetrievalFunc that returns message data.
// Attachment field is omitted due to size.
func RetrieveMessageDataForUser(gs graph.Service, user, m365ID string) (absser.Parsable, error) {
	return gs.Client().UsersById(user).MessagesById(m365ID).Get()
}

func CollectMailFolders(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
) error {
	queryService, err := createService(qp.Credentials, qp.FailFast)
	if err != nil {
		return errors.New("unable to create a mail folder query service for " + qp.User)
	}

	query, err := GetAllFolderNamesForUser(queryService, qp.User)
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
		query,
		&queryService.adapter,
		models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return errors.Wrap(err, "unable to create iterator during mail folder query service")
	}

	callbackFunc := IterateFilterFolderDirectoriesForCollections(
		ctx,
		qp,
		err,
		collections,
		statusUpdater,
	)

	iterateFailure := pageIterator.Iterate(callbackFunc)
	if iterateFailure != nil {
		err = support.WrapAndAppend(qp.User+" iterate failure", iterateFailure, err)
	}

	return err
}
