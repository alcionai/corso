package exchange

import (
	"fmt"

	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/path"
)

// -----------------------------------------------------------------------
// Constant Section
// Defines the allowable strings that can be passed into
// selectors for M365 objects
// -----------------------------------------------------------------------
var (
	fieldsForCalendars = map[string]int{
		"changeKey":         1,
		"events":            2,
		"id":                3,
		"isDefaultCalendar": 4,
		"name":              5,
		"owner":             6,
	}

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
	calendars
	events
	messages
	users
	contacts
)

func CategoryToOptionIdentifier(category path.CategoryType) optionIdentifier {
	switch category {
	case path.EmailCategory:
		return messages
	case path.ContactsCategory:
		return contacts
	case path.EventsCategory:
		return events
	default:
		return unknown
	}
}

// -----------------------------------------------------------------------
// exchange.Query Option Section
// These functions can be used to filter a response on M365
// Graph queries and reduce / filter the amount of data returned
// which reduces the overall latency of complex calls
// -----------------------------------------------------------------------

func optionsForFolderMessagesDelta(
	moreOps []string,
) (*msuser.UsersItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, messages)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForCalendars places allowed options for exchange.Calendar object
// @param moreOps should reflect elements from fieldsForCalendars
// @return is first call in Calendars().GetWithRequestConfigurationAndResponseHandler
func optionsForCalendars(moreOps []string) (
	*msuser.UsersItemCalendarsRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, calendars)
	if err != nil {
		return nil, err
	}
	// should be a CalendarsRequestBuilderGetRequestConfiguration
	requestParams := &msuser.UsersItemCalendarsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemCalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}

	return options, nil
}

// optionsForContactFolders places allowed options for exchange.ContactFolder object
// @return is first call in ContactFolders().GetWithRequestConfigurationAndResponseHandler
func optionsForContactFolders(moreOps []string) (
	*msuser.UsersItemContactFoldersRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemContactFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemContactFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFolderByID(moreOps []string) (
	*msuser.UsersItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemContactFoldersContactFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForMailFolders transforms the options into a more dynamic call for MailFolders.
// @param moreOps is a []string of options(e.g. "displayName", "isHidden")
// @return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(
	moreOps []string,
) (*msuser.UsersItemMailFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemMailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemMailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForMailFoldersItem transforms the options into a more dynamic call for MailFoldersById.
// moreOps is a []string of options(e.g. "displayName", "isHidden")
// Returns first call in MailFoldersById().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFoldersItem(
	moreOps []string,
) (*msuser.UsersItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFoldersItemDelta(
	moreOps []string,
) (*msuser.UsersItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, contacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &msuser.UsersItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForEvents ensures valid option inputs for exchange.Events
// @return is first call in Events().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForEvents(moreOps []string) (*msuser.UsersItemEventsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, events)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemEventsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemEventsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForEvents ensures a valid option inputs for `exchange.Events` when selected from within a Calendar
func optionsForEventsByCalendar(
	moreOps []string,
) (*msuser.UsersItemCalendarsItemEventsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, events)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemCalendarsItemEventsRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &msuser.UsersItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContactChildFolders builds a contacts child folders request.
func optionsForContactChildFolders(
	moreOps []string,
) (*msuser.UsersItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, contacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemContactFoldersItemChildFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContacts transforms options into select query for MailContacts
// @return is the first call in Contacts().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForContacts(moreOps []string) (*msuser.UsersItemContactsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, contacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.UsersItemContactsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.UsersItemContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// buildOptions - Utility Method for verifying if select options are valid for the m365 object type
// @return is a pair. The first is a string literal of allowable options based on the object type,
// the second is an error. An error is returned if an unsupported option or optionIdentifier was used
func buildOptions(options []string, optID optionIdentifier) ([]string, error) {
	var (
		allowedOptions  map[string]int
		returnedOptions = []string{"id"}
	)

	switch optID {
	case calendars:
		allowedOptions = fieldsForCalendars
	case contacts:
		allowedOptions = fieldsForContacts
	case events:
		allowedOptions = fieldsForEvents
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
			return nil, fmt.Errorf("unsupported element passed to buildOptions: %v", entry)
		}

		returnedOptions = append(returnedOptions, entry)
	}

	return returnedOptions, nil
}
