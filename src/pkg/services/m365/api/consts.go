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

	AddressKey     = "address"
	CoordinatesKey = "coordinates"
	DisplayNameKey = "displayName"
	LocationURIKey = "locationUri"
	UniqueIDKey    = "uniqueId"

	// entries that are nested within a second layer
	CityKey       = "city"
	CountryKey    = "countryOrRegion"
	PostalCodeKey = "postalCode"
	StateKey      = "state"
	StreetKey     = "street"
	LatitudeKey   = "latitude"
	LongitudeKey  = "longitude"

	CountryOrRegionFN = "CountryOrRegion"
	StateFN           = "State"
	CityFN            = "City"
	PostalCodeFN      = "PostalCode"
	StreetFN          = "Street"
	GeoLocFN          = "GeoLoc"
	DispNameFN        = "DispName"

	HyperlinkDescriptionKey = "Description"
	HyperlinkURLKey         = "Url"

	LookupIDKey    = "LookupId"
	LookupValueKey = "LookupValue"

	PersonEmailKey = "Email"

	MetadataLabelKey    = "Label"
	MetadataTermGUIDKey = "TermGuid"
	MetadataWssIDKey    = "WssId"

	LinkTitleFieldNamePart  = "LinkTitle"
	ChildCountFieldNamePart = "ChildCount"
	LookupIDFieldNamePart   = "LookupId"

	ODataTypeFieldNamePart      = "@odata.type"
	ODataTypeFieldNameStringVal = "Collection(Edm.String)"
	ODataTypeFieldNameIntVal    = "Collection(Edm.Int32)"

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
	AddressKey,
	CoordinatesKey,
	DisplayNameKey,
	LocationURIKey,
	UniqueIDKey,
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
