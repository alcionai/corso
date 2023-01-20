package api

import (
	"fmt"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/users"
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
) (*users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForMessages)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForCalendars places allowed options for exchange.Calendar object
// @param moreOps should reflect elements from fieldsForCalendars
// @return is first call in Calendars().GetWithRequestConfigurationAndResponseHandler
func optionsForCalendars(moreOps []string) (
	*users.ItemCalendarsRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForCalendars)
	if err != nil {
		return nil, err
	}
	// should be a CalendarsRequestBuilderGetRequestConfiguration
	requestParams := &users.ItemCalendarsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}

	return options, nil
}

// optionsForContactFolders places allowed options for exchange.ContactFolder object
// @return is first call in ContactFolders().GetWithRequestConfigurationAndResponseHandler
func optionsForContactFolders(moreOps []string) (
	*users.ItemContactFoldersRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemContactFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemContactFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFolderByID(moreOps []string) (
	*users.ItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemContactFoldersContactFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForMailFolders transforms the options into a more dynamic call for MailFolders.
// @param moreOps is a []string of options(e.g. "displayName", "isHidden")
// @return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(
	moreOps []string,
) (*users.ItemMailFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemMailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemMailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForMailFoldersItem transforms the options into a more dynamic call for MailFoldersById.
// moreOps is a []string of options(e.g. "displayName", "isHidden")
// Returns first call in MailFoldersById().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFoldersItem(
	moreOps []string,
) (*users.ItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForFolders)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

func optionsForContactFoldersItemDelta(
	moreOps []string,
) (*users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForEvents ensures a valid option inputs for `exchange.Events` when selected from within a Calendar
func optionsForEventsByCalendarDelta(
	moreOps []string,
) (*users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForEvents)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContactChildFolders builds a contacts child folders request.
func optionsForContactChildFolders(
	moreOps []string,
) (*users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemContactFoldersItemChildFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return options, nil
}

// optionsForContacts transforms options into select query for MailContacts
// @return is the first call in Contacts().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForContacts(moreOps []string) (*users.ItemContactsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemContactsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemContactsRequestBuilderGetRequestConfiguration{
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
