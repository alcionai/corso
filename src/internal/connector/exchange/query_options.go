package exchange

import (
	"github.com/pkg/errors"

	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	mscontacts "github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"
	msevents "github.com/microsoftgraph/msgraph-sdk-go/users/item/events"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	msitem "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"
)

//-----------------------------------------------------------------------
// Constant Section
// Defines the allowable strings that can be passed into
// selectors for M365 objects
//------------------------------------------------------------
var (
	fieldsForEvents = map[string]int{
		"calendar":          1,
		"end":               2,
		"id":                3,
		"isOnlineMeeting":   4,
		"isReminderOn":      5,
		"responseStatus":    6,
		"responseRequested": 7,
		"showAs":            8,
		"subject":           9,
	}

	fieldsForFolders = map[string]int{
		"childFolderCount": 1,
		"displayName":      2,
		"id":               3,
		"isHidden":         4,
		"parentFolderId":   5,
		"totalItemCount":   6,
		"unreadItemCount":  7,
	}

	fieldsForUsers = map[string]int{
		"birthday":          1,
		"businessPhones":    2,
		"city":              3,
		"companyName":       4,
		"department":        5,
		"displayName":       6,
		"employeeId":        7,
		"id":                8,
		"mail":              9,
		"userPrincipalName": 10,
	}

	fieldsForMessages = map[string]int{
		"conservationId":    1,
		"conversationIndex": 2,
		"parentFolderId":    3,
		"subject":           4,
		"webLink":           5,
		"id":                6,
	}

	fieldsForContacts = map[string]int{
		"id":             1,
		"companyName":    2,
		"department":     3,
		"displayName":    4,
		"fileAs":         5,
		"givenName":      6,
		"manager":        7,
		"parentFolderId": 8,
	}
)

type optionIdentifier int

//go:generate stringer -type=optionIdentifier
const (
	unknown optionIdentifier = iota
	folders
	events
	messages
	users
	contacts
)

//---------------------------------------------------
// exchange.Query Option Section
//------------------------------------------------

// optionsForMessages - used to select allowable options for exchange.Mail types
// @param moreOps is []string of options(e.g. "parentFolderId, subject")
// @return is first call in Messages().GetWithRequestConfigurationAndResponseHandler
func optionsForMessages(moreOps []string) (*msmessage.MessagesRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, messages)
	if err != nil {
		return nil, err
	}
	requestParameters := &msmessage.MessagesRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msmessage.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options, nil
}

// optionsForSingleMessage to select allowable option for a singular exchange.Mail object
// @params moreOps is []string of options (e.g. subject, content.Type)
// @return is first call in MessageById().GetWithRequestConfigurationAndResponseHandler
func OptionsForSingleMessage(moreOps []string) (*msitem.MessageItemRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, messages)
	if err != nil {
		return nil, err
	}
	requestParams := &msitem.MessageItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msitem.MessageItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}
	return options, nil
}

// optionsForMailFolders transforms the options into a more dynamic call for MailFolders.
// @param moreOps is a []string of options(e.g. "displayName", "isHidden")
// @return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(moreOps []string) (*msfolder.MailFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msfolder.MailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msfolder.MailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options, nil
}

// optionsForEvents ensures valid option inputs for exchange.Events
// @return is first call in Events().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForEvents(moreOps []string) (*msevents.EventsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, events)
	if err != nil {
		return nil, err
	}
	requestParameters := &msevents.EventsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msevents.EventsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options, nil
}

// optionsForContacts transforms options into select query for MailContacts
// @return is the first call in Contacts().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForContacts(moreOps []string) (*mscontacts.ContactsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, contacts)
	if err != nil {
		return nil, err
	}
	requestParameters := &mscontacts.ContactsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mscontacts.ContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options, nil
}

func optionsForUsers(moreOps []string) (*msuser.UsersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, users)
	if err != nil {
		return nil, err
	}
	requestParams := &msuser.UsersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}
	return options, nil
}

// buildOptions - Utility Method for verifying if select options are valid for the m365 object type
// @return is a pair. The first is a string literal of allowable options based on the object type,
// the second is an error. An error is returned if an unsupported option or optionIdentifier was used
func buildOptions(options []string, optID optionIdentifier) ([]string, error) {
	var allowedOptions map[string]int
	returnedOptions := []string{"id"}

	switch optID {
	case events:
		allowedOptions = fieldsForEvents
	case contacts:
		allowedOptions = fieldsForContacts
	case folders:
		allowedOptions = fieldsForFolders
	case users:
		allowedOptions = fieldsForUsers
	case messages:
		allowedOptions = fieldsForMessages
	case unknown:
		fallthrough
	default:
		return nil, errors.New("unsupported option")
	}

	for _, entry := range options {
		_, ok := allowedOptions[entry]
		if !ok {
			return nil, errors.New("unsupported option")
		}

		returnedOptions = append(returnedOptions, entry)
	}
	return returnedOptions, nil
}
