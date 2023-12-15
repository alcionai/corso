package api

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/maps"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

const (
	AttachmentsColumnName    = "Attachments"
	EditColumnName           = "Edit"
	ContentTypeColumnName    = "ContentType"
	CreatedColumnName        = "Created"
	ModifiedColumnName       = "Modified"
	AuthorLookupIDColumnName = "AuthorLookupId"
	EditorLookupIDColumnName = "EditorLookupId"

	ContentTypeColumnDisplayName = "Content Type"

	AddressFieldName     = "address"
	CoordinatesFieldName = "coordinates"
	DisplayNameFieldName = "displayName"
	LocationURIFieldName = "locationUri"
	UniqueIDFieldName    = "uniqueId"

	CountryOrRegionFieldName = "CountryOrRegion"
	StateFieldName           = "State"
	CityFieldName            = "City"
	PostalCodeFieldName      = "PostalCode"
	StreetFieldName          = "Street"
	GeoLocFieldName          = "GeoLoc"
	DispNameFieldName        = "DispName"
	LinkTitleFieldNamePart   = "LinkTitle"
	ChildCountFieldNamePart  = "ChildCount"

	ReadOnlyOrHiddenFieldNamePrefix = "_"
	DescoratorFieldNamePrefix       = "@"
)

var addressFieldNames = []string{
	AddressFieldName,
	CoordinatesFieldName,
	DisplayNameFieldName,
	LocationURIFieldName,
	UniqueIDFieldName,
}

var readOnlyAddressFieldNames = []string{
	CountryOrRegionFieldName,
	StateFieldName,
	CityFieldName,
	PostalCodeFieldName,
	StreetFieldName,
	GeoLocFieldName,
	DispNameFieldName,
}

var legacyColumns = maps.KeySet{
	AttachmentsColumnName:        {},
	EditColumnName:               {},
	ContentTypeColumnDisplayName: {},
}

var readOnlyFieldNames = maps.KeySet{
	AttachmentsColumnName:    {},
	EditColumnName:           {},
	ContentTypeColumnName:    {},
	CreatedColumnName:        {},
	ModifiedColumnName:       {},
	AuthorLookupIDColumnName: {},
	EditorLookupIDColumnName: {},
}

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Lists() Lists {
	return Lists{c}
}

// Lists is an interface-compliant provider of the client.
type Lists struct {
	Client
}

// PostDrive creates a new list of type drive.  Specifically used to create
// documentLibraries for SharePoint Sites.
func (c Lists) PostDrive(
	ctx context.Context,
	siteID, driveName string,
) (models.Driveable, error) {
	list := models.NewList()
	list.SetDisplayName(&driveName)
	list.SetDescription(ptr.To("corso auto-generated restore destination"))

	li := models.NewListInfo()
	li.SetTemplate(ptr.To("documentLibrary"))
	list.SetList(li)

	// creating a list of type documentLibrary will result in the creation
	// of a new drive owned by the given site.
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists()

	newList, err := builder.Post(ctx, list, nil)
	if graph.IsErrItemAlreadyExistsConflict(err) {
		return nil, clues.StackWC(ctx, graph.ErrItemAlreadyExistsConflict, err)
	}

	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating documentLibrary list")
	}

	// drive information is not returned by the list creation.
	drive, err := builder.
		ByListId(ptr.Val(newList.GetId())).
		Drive().
		Get(ctx, nil)

	return drive, graph.Wrap(ctx, err, "fetching created documentLibrary").OrNil()
}

// SharePoint lists represent lists on a site. Inherits additional properties from
// baseItem: https://learn.microsoft.com/en-us/graph/api/resources/baseitem?view=graph-rest-1.0
// The full details concerning SharePoint Lists can
// be found at: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
// Note additional calls are required for the relationships that exist outside of the object properties.

// GetListById is a utility function to populate a SharePoint.List with objects associated with a given siteID.
// @param siteID the M365 ID that represents the SharePoint Site
// Makes additional calls to retrieve the following relationships:
// - Columns
// - ContentTypes
// - List Items
func (c Lists) GetListByID(ctx context.Context, siteID, listID string) (models.Listable, error) {
	list, err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "fetching list")
	}

	cols, cTypes, lItems, err := c.getListContents(ctx, siteID, listID)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting list contents")
	}

	list.SetColumns(cols)
	list.SetContentTypes(cTypes)
	list.SetItems(lItems)

	return list, nil
}

// getListContents utility function to retrieve associated M365 relationships
// which are not included with the standard List query:
// - Columns, ContentTypes, ListItems
func (c Lists) getListContents(ctx context.Context, siteID, listID string) (
	[]models.ColumnDefinitionable,
	[]models.ContentTypeable,
	[]models.ListItemable,
	error,
) {
	cols, err := c.GetListColumns(ctx, siteID, listID, CallConfig{})
	if err != nil {
		return nil, nil, nil, err
	}

	cTypes, err := c.GetContentTypes(ctx, siteID, listID, CallConfig{})
	if err != nil {
		return nil, nil, nil, err
	}

	for i := 0; i < len(cTypes); i++ {
		columnLinks, err := c.GetColumnLinks(ctx, siteID, listID, ptr.Val(cTypes[i].GetId()), CallConfig{})
		if err != nil {
			return nil, nil, nil, err
		}

		cTypes[i].SetColumnLinks(columnLinks)

		cTypeColumns, err := c.GetCTypesColumns(ctx, siteID, listID, ptr.Val(cTypes[i].GetId()), CallConfig{})
		if err != nil {
			return nil, nil, nil, err
		}

		cTypes[i].SetColumns(cTypeColumns)
	}

	lItems, err := c.GetListItems(ctx, siteID, listID, CallConfig{})
	if err != nil {
		return nil, nil, nil, err
	}

	for _, li := range lItems {
		fields, err := c.getListItemFields(ctx, siteID, listID, ptr.Val(li.GetId()))
		if err != nil {
			return nil, nil, nil, err
		}

		li.SetFields(fields)
	}

	return cols, cTypes, lItems, nil
}

func (c Lists) PostList(
	ctx context.Context,
	siteID string,
	listName string,
	oldListByteArray []byte,
) (models.Listable, error) {
	newListName := listName

	oldList, err := BytesToListable(oldListByteArray)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "generating list from stored bytes")
	}

	// the input listName is of format: destinationName_listID
	// here we replace listID with displayName of list generated from stored bytes
	if name, ok := ptr.ValOK(oldList.GetDisplayName()); ok {
		nameParts := strings.Split(listName, "_")
		if len(nameParts) > 0 {
			nameParts[len(nameParts)-1] = name
			newListName = strings.Join(nameParts, "_")
		}
	}

	// this ensure all columns, contentTypes are set to the newList
	newList := ToListable(oldList, newListName)

	// Restore to List base to M365 back store
	restoredList, err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		Post(ctx, newList, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating list")
	}

	listItems := make([]models.ListItemable, 0)

	for _, itm := range oldList.GetItems() {
		temp := CloneListItem(itm)
		listItems = append(listItems, temp)
	}

	err = c.PostListItems(
		ctx,
		siteID,
		ptr.Val(restoredList.GetId()),
		listItems)
	if err == nil {
		restoredList.SetItems(listItems)
		return restoredList, nil
	}

	// [TODO](hitesh) double check if we need to:
	// 1. rollback the entire list
	// 2. restore as much list items possible and add recoverables to fault bus
	// rollback list creation
	err = c.DeleteList(ctx, siteID, ptr.Val(restoredList.GetId()))

	return nil, graph.Wrap(ctx, err, "deleting restored list after items creation failure").OrNil()
}

func (c Lists) PostListItems(
	ctx context.Context,
	siteID, listID string,
	listItems []models.ListItemable,
) error {
	for _, lItem := range listItems {
		_, err := c.Stable.
			Client().
			Sites().
			BySiteId(siteID).
			Lists().
			ByListId(listID).
			Items().
			Post(ctx, lItem, nil)
		if err != nil {
			return graph.Wrap(ctx, err, "creating item in list")
		}
	}

	return nil
}

func (c Lists) DeleteList(
	ctx context.Context,
	siteID, listID string,
) error {
	err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Delete(ctx, nil)

	return graph.Wrap(ctx, err, "deleting list").OrNil()
}

func BytesToListable(bytes []byte) (models.Listable, error) {
	parsable, err := CreateFromBytes(bytes, models.CreateListFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to sharepoint list")
	}

	list := parsable.(models.Listable)

	return list, nil
}

// ToListable utility function to encapsulate stored data for restoration.
// New Listable omits trackable fields such as `id` or `ETag` and other read-only
// objects that are prevented upon upload. Additionally, read-Only columns are
// not attached in this method.
// ListItems are not included in creation of new list, and have to be restored
// in separate call.
func ToListable(orig models.Listable, displayName string) models.Listable {
	newList := models.NewList()

	newList.SetContentTypes(orig.GetContentTypes())
	newList.SetCreatedBy(orig.GetCreatedBy())
	newList.SetCreatedByUser(orig.GetCreatedByUser())
	newList.SetCreatedDateTime(orig.GetCreatedDateTime())
	newList.SetDescription(orig.GetDescription())
	newList.SetDisplayName(&displayName)
	newList.SetLastModifiedBy(orig.GetLastModifiedBy())
	newList.SetLastModifiedByUser(orig.GetLastModifiedByUser())
	newList.SetLastModifiedDateTime(orig.GetLastModifiedDateTime())
	newList.SetList(orig.GetList())
	newList.SetOdataType(orig.GetOdataType())
	newList.SetParentReference(orig.GetParentReference())

	columns := make([]models.ColumnDefinitionable, 0)

	for _, cd := range orig.GetColumns() {
		var (
			displayName string
			readOnly    bool
		)

		if name, ok := ptr.ValOK(cd.GetDisplayName()); ok {
			displayName = name
		}

		if ro, ok := ptr.ValOK(cd.GetReadOnly()); ok {
			readOnly = ro
		}

		// Skips columns that cannot be uploaded for models.ColumnDefinitionable:
		// - ReadOnly, Title, or Legacy columns: Attachments, Edit, or Content Type
		if readOnly || displayName == "Title" || legacyColumns.HasKey(displayName) {
			continue
		}

		columns = append(columns, cloneColumnDefinitionable(cd))
	}

	newList.SetColumns(columns)

	return newList
}

// cloneColumnDefinitionable utility function for encapsulating models.ColumnDefinitionable data
// into new object for upload.
func cloneColumnDefinitionable(orig models.ColumnDefinitionable) models.ColumnDefinitionable {
	newColumn := models.NewColumnDefinition()

	// column attributes
	newColumn.SetName(orig.GetName())
	newColumn.SetOdataType(orig.GetOdataType())
	newColumn.SetPropagateChanges(orig.GetPropagateChanges())
	newColumn.SetReadOnly(orig.GetReadOnly())
	newColumn.SetRequired(orig.GetRequired())
	newColumn.SetAdditionalData(orig.GetAdditionalData())
	newColumn.SetDescription(orig.GetDescription())
	newColumn.SetDisplayName(orig.GetDisplayName())
	newColumn.SetSourceColumn(orig.GetSourceColumn())
	newColumn.SetSourceContentType(orig.GetSourceContentType())
	newColumn.SetHidden(orig.GetHidden())
	newColumn.SetIndexed(orig.GetIndexed())
	newColumn.SetIsDeletable(orig.GetIsDeletable())
	newColumn.SetIsReorderable(orig.GetIsReorderable())
	newColumn.SetIsSealed(orig.GetIsSealed())
	newColumn.SetTypeEscaped(orig.GetTypeEscaped())
	newColumn.SetColumnGroup(orig.GetColumnGroup())
	newColumn.SetEnforceUniqueValues(orig.GetEnforceUniqueValues())

	// column types
	setColumnType(newColumn, orig)

	// Requires nil checks to avoid Graph error: 'General exception while processing'
	defaultValue := orig.GetDefaultValue()
	if defaultValue != nil {
		newColumn.SetDefaultValue(defaultValue)
	}

	validation := orig.GetValidation()
	if validation != nil {
		newColumn.SetValidation(validation)
	}

	return newColumn
}

func setColumnType(newColumn *models.ColumnDefinition, orig models.ColumnDefinitionable) {
	switch {
	case orig.GetText() != nil:
		newColumn.SetText(orig.GetText())
	case orig.GetBoolean() != nil:
		newColumn.SetBoolean(orig.GetBoolean())
	case orig.GetCalculated() != nil:
		newColumn.SetCalculated(orig.GetCalculated())
	case orig.GetChoice() != nil:
		newColumn.SetChoice(orig.GetChoice())
	case orig.GetContentApprovalStatus() != nil:
		newColumn.SetContentApprovalStatus(orig.GetContentApprovalStatus())
	case orig.GetCurrency() != nil:
		newColumn.SetCurrency(orig.GetCurrency())
	case orig.GetDateTime() != nil:
		newColumn.SetDateTime(orig.GetDateTime())
	case orig.GetGeolocation() != nil:
		newColumn.SetGeolocation(orig.GetGeolocation())
	case orig.GetHyperlinkOrPicture() != nil:
		newColumn.SetHyperlinkOrPicture(orig.GetHyperlinkOrPicture())
	case orig.GetNumber() != nil:
		newColumn.SetNumber(orig.GetNumber())
	case orig.GetLookup() != nil:
		newColumn.SetLookup(orig.GetLookup())
	case orig.GetThumbnail() != nil:
		newColumn.SetThumbnail(orig.GetThumbnail())
	case orig.GetTerm() != nil:
		newColumn.SetTerm(orig.GetTerm())
	case orig.GetPersonOrGroup() != nil:
		newColumn.SetPersonOrGroup(orig.GetPersonOrGroup())
	default:
		newColumn.SetText(models.NewTextColumn())
	}
}

// CloneListItem creates a new `SharePoint.ListItem` and stores the original item's
// M365 data into it set fields.
// - https://learn.microsoft.com/en-us/graph/api/resources/listitem?view=graph-rest-1.0
func CloneListItem(orig models.ListItemable) models.ListItemable {
	newItem := models.NewListItem()

	// list item data
	newFieldData := retrieveFieldData(orig.GetFields())
	newItem.SetFields(newFieldData)

	// list item attributes
	newItem.SetAdditionalData(orig.GetAdditionalData())
	newItem.SetDescription(orig.GetDescription())
	newItem.SetCreatedBy(orig.GetCreatedBy())
	newItem.SetCreatedDateTime(orig.GetCreatedDateTime())
	newItem.SetLastModifiedBy(orig.GetLastModifiedBy())
	newItem.SetLastModifiedDateTime(orig.GetLastModifiedDateTime())
	newItem.SetOdataType(orig.GetOdataType())
	newItem.SetAnalytics(orig.GetAnalytics())
	newItem.SetContentType(orig.GetContentType())
	newItem.SetVersions(orig.GetVersions())

	// Requires nil checks to avoid Graph error: 'Invalid request'
	lastCreatedByUser := orig.GetCreatedByUser()
	if lastCreatedByUser != nil {
		newItem.SetCreatedByUser(lastCreatedByUser)
	}

	lastModifiedByUser := orig.GetLastModifiedByUser()
	if lastCreatedByUser != nil {
		newItem.SetLastModifiedByUser(lastModifiedByUser)
	}

	return newItem
}

// retrieveFieldData utility function to clone raw listItem data from the embedded
// additionalData map
// Further details on FieldValueSets:
// - https://learn.microsoft.com/en-us/graph/api/resources/fieldvalueset?view=graph-rest-1.0
func retrieveFieldData(orig models.FieldValueSetable) models.FieldValueSetable {
	fields := models.NewFieldValueSet()
	additionalData := filterAdditionalData(orig)

	retainPrimaryAddressField(additionalData)

	fields.SetAdditionalData(additionalData)

	return fields
}

func filterAdditionalData(orig models.FieldValueSetable) map[string]any {
	if orig == nil {
		return make(map[string]any)
	}

	fieldData := orig.GetAdditionalData()
	filteredData := make(map[string]any)

	for key, value := range fieldData {
		if shouldFilterField(key, value) {
			continue
		}

		filteredData[key] = value
	}

	return filteredData
}

func shouldFilterField(key string, value any) bool {
	return readOnlyFieldNames.HasKey(key) ||
		strings.HasPrefix(key, ReadOnlyOrHiddenFieldNamePrefix) ||
		strings.HasPrefix(key, DescoratorFieldNamePrefix) ||
		strings.Contains(key, LinkTitleFieldNamePart) ||
		strings.Contains(key, ChildCountFieldNamePart)
}

func retainPrimaryAddressField(additionalData map[string]any) {
	if !hasAddressFields(additionalData) {
		return
	}

	for _, k := range readOnlyAddressFieldNames {
		delete(additionalData, k)
	}
}

func hasAddressFields(additionalData map[string]any) bool {
	if !maps.HasKeys(additionalData, readOnlyAddressFieldNames...) {
		return false
	}

	for _, value := range additionalData {
		nestedFields, ok := value.(map[string]any)
		if !ok || maps.HasKeys(nestedFields, GeoLocFieldName) {
			continue
		}

		if maps.HasKeys(nestedFields, addressFieldNames...) {
			return true
		}
	}

	return false
}

func (c Lists) getListItemFields(
	ctx context.Context,
	siteID, listID, itemID string,
) (models.FieldValueSetable, error) {
	prefix := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Items().
		ByListItemId(itemID)

	fields, err := prefix.Fields().Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	return fields, nil
}
