package exchange

import (
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
	mscalendars "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars"
	mscevents "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events"
	mscontactfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders"
	mscontactfolderitem "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item"
	mscontactfolderchild "github.com/microsoftgraph/msgraph-sdk-go/users/item/contactfolders/item/childfolders"
	mscontacts "github.com/microsoftgraph/msgraph-sdk-go/users/item/contacts"
	msevents "github.com/microsoftgraph/msgraph-sdk-go/users/item/events"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msfolderitem "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	msitem "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"
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

// Delta requests for mail and contacts have the same parameters and config
// structs.
type DeltaRequestBuilderGetQueryParameters struct {
	Count   *bool    `uriparametername:"%24count"`
	Filter  *string  `uriparametername:"%24filter"`
	Orderby []string `uriparametername:"%24orderby"`
	Search  *string  `uriparametername:"%24search"`
	Select  []string `uriparametername:"%24select"`
	Skip    *int32   `uriparametername:"%24skip"`
	Top     *int32   `uriparametername:"%24top"`
}

type DeltaRequestBuilderGetRequestConfiguration struct {
	Headers         map[string]string
	Options         []abs.RequestOption
	QueryParameters *DeltaRequestBuilderGetQueryParameters
}

func optionsForFolderMessages(moreOps []string) (*DeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, messages)
	if err != nil {
		return nil, err
	}

	requestParameters := &DeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &DeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

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

// optionsForCalendars places allowed options for exchange.Calendar object
// @param moreOps should reflect elements from fieldsForCalendars
// @return is first call in Calendars().GetWithRequestConfigurationAndResponseHandler
func optionsForCalendars(moreOps []string) (
	*mscalendars.CalendarsRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, calendars)
	if err != nil {
		return nil, err
	}

	requestParams := &mscalendars.CalendarsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mscalendars.CalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}

	return options, nil
}

// optionsForContactFolders places allowed options for exchange.ContactFolder object
// @return is first call in ContactFolders().GetWithRequestConfigurationAndResponseHandler
func optionsForContactFolders(moreOps []string) (
	*mscontactfolder.ContactFoldersRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &mscontactfolder.ContactFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mscontactfolder.ContactFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFolderByID(moreOps []string) (
	*mscontactfolderitem.ContactFolderItemRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &mscontactfolderitem.ContactFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mscontactfolderitem.ContactFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
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

// optionsForMailFoldersItem transforms the options into a more dynamic call for MailFoldersById.
// moreOps is a []string of options(e.g. "displayName", "isHidden")
// Returns first call in MailFoldersById().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFoldersItem(
	moreOps []string,
) (*msfolderitem.MailFolderItemRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msfolderitem.MailFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msfolderitem.MailFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContactFoldersItem is the same as optionsForContacts.
func optionsForContactFoldersItem(
	moreOps []string,
) (*DeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, contacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &DeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &DeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForEvents ensures valid option inputs for exchange.Events
// @return is first call in Events().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForCalendarEvents(moreOps []string) (*mscevents.EventsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, events)
	if err != nil {
		return nil, err
	}

	requestParameters := &mscevents.EventsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mscevents.EventsRequestBuilderGetRequestConfiguration{
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

// optionsForContactChildFolders builds a contacts child folders request.
func optionsForContactChildFolders(
	moreOps []string,
) (*mscontactfolderchild.ChildFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, contacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &mscontactfolderchild.ChildFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &mscontactfolderchild.ChildFoldersRequestBuilderGetRequestConfiguration{
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
