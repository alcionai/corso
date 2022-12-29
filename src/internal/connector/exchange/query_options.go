package exchange

import (
	"fmt"

	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
)

// -----------------------------------------------------------------------
// Constant Section
// Defines the allowable strings that can be passed into
// selectors for M365 objects
// -----------------------------------------------------------------------
var (
	fieldsForCalendars = map[string]struct{}{
		"changeKey":         {},
		"events":            {},
		"id":                {},
		"isDefaultCalendar": {},
		"name":              {},
		"owner":             {},
	}

	fieldsForEvents = map[string]struct{}{
		"calendar":          {},
		"end":               {},
		"id":                {},
		"isOnlineMeeting":   {},
		"isReminderOn":      {},
		"responseStatus":    {},
		"responseRequested": {},
		"showAs":            {},
		"subject":           {},
	}

	fieldsForFolders = map[string]struct{}{
		"childFolderCount": {},
		"displayName":      {},
		"id":               {},
		"isHidden":         {},
		"parentFolderId":   {},
		"totalItemCount":   {},
		"unreadItemCount":  {},
	}

	fieldsForUsers = map[string]struct{}{
		"birthday":          {},
		"businessPhones":    {},
		"city":              {},
		"companyName":       {},
		"department":        {},
		"displayName":       {},
		"employeeId":        {},
		"id":                {},
		"mail":              {},
		"userPrincipalName": {},
	}

	fieldsForMessages = map[string]struct{}{
		"conservationId":    {},
		"conversationIndex": {},
		"parentFolderId":    {},
		"subject":           {},
		"webLink":           {},
		"id":                {},
		"isRead":            {},
	}

	fieldsForContacts = map[string]struct{}{
		"id":             {},
		"companyName":    {},
		"department":     {},
		"displayName":    {},
		"fileAs":         {},
		"givenName":      {},
		"manager":        {},
		"parentFolderId": {},
	}
)

// -----------------------------------------------------------------------
// exchange.Query Option Section
// These functions can be used to filter a response on M365
// Graph queries and reduce / filter the amount of data returned
// which reduces the overall latency of complex calls
// -----------------------------------------------------------------------

func optionsForFolderMessagesDelta(
	moreOps []string,
) (*msuser.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForMessages)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForCalendars places allowed options for exchange.Calendar object
// @param moreOps should reflect elements from fieldsForCalendars
// @return is first call in Calendars().GetWithRequestConfigurationAndResponseHandler
func optionsForCalendars(moreOps []string) (
	*msuser.ItemCalendarsRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForCalendars)
	if err != nil {
		return nil, err
	}
	// should be a CalendarsRequestBuilderGetRequestConfiguration
	requestParams := &msuser.ItemCalendarsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemCalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}

	return options, nil
}

// optionsForContactFolders places allowed options for exchange.ContactFolder object
// @return is first call in ContactFolders().GetWithRequestConfigurationAndResponseHandler
func optionsForContactFolders(moreOps []string) (
	*msuser.ItemContactFoldersRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemContactFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemContactFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFolderByID(moreOps []string) (
	*msuser.ItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemContactFoldersContactFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForMailFolders transforms the options into a more dynamic call for MailFolders.
// @param moreOps is a []string of options(e.g. "displayName", "isHidden")
// @return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(
	moreOps []string,
) (*msuser.ItemMailFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemMailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemMailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForMailFoldersItem transforms the options into a more dynamic call for MailFoldersById.
// moreOps is a []string of options(e.g. "displayName", "isHidden")
// Returns first call in MailFoldersById().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFoldersItem(
	moreOps []string,
) (*msuser.ItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFoldersItemDelta(
	moreOps []string,
) (*msuser.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &msuser.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForEvents ensures valid option inputs for exchange.Events
// @return is first call in Events().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForEvents(moreOps []string) (*msuser.ItemEventsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForEvents)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemEventsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemEventsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForEvents ensures a valid option inputs for `exchange.Events` when selected from within a Calendar
func optionsForEventsByCalendar(
	moreOps []string,
) (*msuser.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForEvents)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemCalendarsItemEventsRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &msuser.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContactChildFolders builds a contacts child folders request.
func optionsForContactChildFolders(
	moreOps []string,
) (*msuser.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemContactFoldersItemChildFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContacts transforms options into select query for MailContacts
// @return is the first call in Contacts().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForContacts(moreOps []string) (*msuser.ItemContactsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &msuser.ItemContactsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msuser.ItemContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// buildOptions - Utility Method for verifying if select options are valid for the m365 object type
// @return is a pair. The first is a string literal of allowable options based on the object type,
// the second is an error. An error is returned if an unsupported option or optionIdentifier was used
func buildOptions(fields []string, allowed map[string]struct{}) ([]string, error) {
	returnedOptions := []string{"id"}

	for _, entry := range fields {
		_, ok := allowed[entry]
		if !ok {
			return nil, fmt.Errorf("unsupported field: %v", entry)
		}
	}

	return append(returnedOptions, fields...), nil
}
