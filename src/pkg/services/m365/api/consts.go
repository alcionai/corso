package api

import (
	"github.com/alcionai/corso/src/internal/common/keys"
)

// Well knwon Folder Names
// Mail Definitions: https://docs.microsoft.com/en-us/graph/api/resources/mailfolder?view=graph-rest-1.0
const (
	DefaultCalendar = "Calendar"
	DefaultContacts = "Contacts"
	MailInbox       = "Inbox"
	MsgFolderRoot   = "msgfolderroot"

	// Kiota JSON invalid JSON error message.
	invalidJSON = "invalid json type"
)

// ************** Lists starts *****************

const (
	AttachmentsColumnName       = "Attachments"
	EditColumnName              = "Edit"
	ContentTypeColumnName       = "ContentType"
	CreatedColumnName           = "Created"
	ModifiedColumnName          = "Modified"
	AuthorLookupIDColumnName    = "AuthorLookupId"
	EditorLookupIDColumnName    = "EditorLookupId"
	AppAuthorLookupIDColumnName = "AppAuthorLookupId"
	TitleColumnName             = "Title"

	ContentTypeColumnDisplayName = "Content Type"

	AddressFieldName          = "address"
	NestedCityFieldName       = "city"
	NestedCountryFieldName    = "countryOrRegion"
	NestedPostalCodeFieldName = "postalCode"
	NestedStateFieldName      = "state"
	NestedStreetFieldName     = "street"
	CoordinatesFieldName      = "coordinates"
	NestedLatitudeFieldName   = "latitude"
	NestedLongitudeFieldName  = "longitude"
	DisplayNameFieldName      = "displayName"
	LocationURIFieldName      = "locationUri"
	UniqueIDFieldName         = "uniqueId"

	CountryOrRegionFieldName = "CountryOrRegion"
	StateFieldName           = "State"
	CityFieldName            = "City"
	PostalCodeFieldName      = "PostalCode"
	StreetFieldName          = "Street"
	GeoLocFieldName          = "GeoLoc"
	DispNameFieldName        = "DispName"

	HyperlinkDescriptionFieldName = "Description"
	HyperlinkURLFieldName         = "Url"

	LinkTitleFieldNamePart  = "LinkTitle"
	ChildCountFieldNamePart = "ChildCount"
	LookupIDFieldNamePart   = "LookupId"

	ReadOnlyOrHiddenFieldNamePrefix = "_"
	DescoratorFieldNamePrefix       = "@"

	WebTemplateExtensionsListTemplate = "webTemplateExtensionsList"
	// This issue https://github.com/alcionai/corso/issues/4932
	// tracks to backup/restore supportability of `documentLibrary` templated lists
	DocumentLibraryListTemplate = "documentLibrary"
	SharingLinksListTemplate    = "sharingLinks"
	AccessRequestsListTemplate  = "accessRequest"
)

var addressFieldNames = []string{
	AddressFieldName,
	CoordinatesFieldName,
	DisplayNameFieldName,
	LocationURIFieldName,
	UniqueIDFieldName,
}

var legacyColumns = keys.Set{
	AttachmentsColumnName:        {},
	EditColumnName:               {},
	ContentTypeColumnDisplayName: {},
}

var SkipListTemplates = keys.Set{
	WebTemplateExtensionsListTemplate: {},
	DocumentLibraryListTemplate:       {},
	SharingLinksListTemplate:          {},
	AccessRequestsListTemplate:        {},
}

// ************** Lists ends *****************
