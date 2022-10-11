package exchange

import (
	"context"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/internal/connector/graph"
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
