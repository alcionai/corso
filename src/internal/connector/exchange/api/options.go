package api

import (
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
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

const (
	// headerKeyPrefer is used to set query preferences
	headerKeyPrefer = "Prefer"
	// maxPageSizeHeaderFmt is used to indicate max page size
	// preferences
	maxPageSizeHeaderFmt = "odata.maxpagesize=%d"
	// deltaMaxPageSize is the max page size to use for delta queries
	deltaMaxPageSize = 200
	idTypeFmt        = "IdType=%q"
	immutableIDType  = "ImmutableId"
)

// -----------------------------------------------------------------------
// exchange.Query Option Section
// These functions can be used to filter a response on M365
// Graph queries and reduce / filter the amount of data returned
// which reduces the overall latency of complex calls
// -----------------------------------------------------------------------
func optionsForFolderMessages(
	moreOps []string,
	immutableIDs bool,
) (*users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForMessages)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemMailFoldersItemMessagesRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &users.ItemMailFoldersItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
		Headers:         buildPreferHeaders(true, immutableIDs),
	}

	return options, nil
}

func optionsForFolderMessagesDelta(
	moreOps []string,
	immutableIDs bool,
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
		Headers:         buildPreferHeaders(true, immutableIDs),
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

// optionsForCalendarsByID places allowed options for exchange.Calendar object
// @param moreOps should reflect elements from fieldsForCalendars
// @return is first call in Calendars().GetWithRequestConfigurationAndResponseHandler
func optionsForCalendarsByID(moreOps []string) (
	*users.ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration,
	error,
) {
	selecting, err := buildOptions(moreOps, fieldsForCalendars)
	if err != nil {
		return nil, err
	}
	// should be a CalendarsRequestBuilderGetRequestConfiguration
	requestParams := &users.ItemCalendarsCalendarItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &users.ItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
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
	immutableIDs bool,
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
		Headers:         buildPreferHeaders(true, immutableIDs),
	}

	return options, nil
}

func optionsForContactFoldersItem(
	moreOps []string,
	immutableIDs bool,
) (*users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, fieldsForContacts)
	if err != nil {
		return nil, err
	}

	requestParameters := &users.ItemContactFoldersItemContactsRequestBuilderGetQueryParameters{
		Select: selecting,
	}

	options := &users.ItemContactFoldersItemContactsRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
		Headers:         buildPreferHeaders(true, immutableIDs),
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

// buildOptions - Utility Method for verifying if select options are valid for the m365 object type
// @return is a pair. The first is a string literal of allowable options based on the object type,
// the second is an error. An error is returned if an unsupported option or optionIdentifier was used
func buildOptions(fields []string, allowed map[string]struct{}) ([]string, error) {
	returnedOptions := []string{"id"}

	for _, entry := range fields {
		_, ok := allowed[entry]
		if !ok {
			return nil, clues.New("unsupported field: " + entry)
		}
	}

	return append(returnedOptions, fields...), nil
}

// buildPreferHeaders returns the headers we add to item delta page
// requests.
func buildPreferHeaders(pageSize, immutableID bool) *abstractions.RequestHeaders {
	var allHeaders []string

	if pageSize {
		allHeaders = append(allHeaders, fmt.Sprintf(maxPageSizeHeaderFmt, deltaMaxPageSize))
	}

	if immutableID {
		allHeaders = append(allHeaders, fmt.Sprintf(idTypeFmt, immutableIDType))
	}

	headers := abstractions.NewRequestHeaders()
	headers.Add(headerKeyPrefer, strings.Join(allHeaders, ","))

	return headers
}
