package api

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var legacyColumns = map[string]struct{}{
	"Attachments":  {},
	"Edit":         {},
	"Content Type": {},
}

var readOnlyFieldNames = map[string]struct{}{
	"Attachments":    {},
	"Edit":           {},
	"ContentType":    {},
	"Created":        {},
	"Modified":       {},
	"AuthorLookupId": {},
	"EditorLookupId": {},
}

var addressFieldNames = map[string]struct{}{
	"address":     {},
	"coordinates": {},
	"displayName": {},
	"locationUri": {},
	"uniqueId":    {},
}

var readonlyAddressFieldNames = map[string]struct{}{
	"CountryOrRegion": {},
	"State":           {},
	"City":            {},
	"PostalCode":      {},
	"Street":          {},
	"GeoLoc":          {},
	"DispName":        {},
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

	return cols, cTypes, lItems, nil
}

func (c Lists) PostList(
	ctx context.Context,
	siteID string,
	listName string,
	oldListByteArray []byte,
) (models.Listable, error) {
	newListName := listName

	oldList, err := CreateListFromBytes(oldListByteArray)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "creating old list")
	}

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
	restoredList, err := c.Stable.Client().Sites().BySiteId(siteID).Lists().Post(ctx, newList, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring list")
	}

	return restoredList, nil
}

func (c Lists) PostListItem(
	ctx context.Context,
	siteID, listID string,
	oldListByteArray []byte,
) ([]models.ListItemable, error) {
	oldList, err := CreateListFromBytes(oldListByteArray)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "creating old list to get list items")
	}

	contents := make([]models.ListItemable, 0)

	for _, itm := range oldList.GetItems() {
		temp := CloneListItem(itm)
		contents = append(contents, temp)
	}

	for _, lItem := range contents {
		_, err := c.Stable.
			Client().
			Sites().
			BySiteId(siteID).
			Lists().
			ByListId(listID).
			Items().
			Post(ctx, lItem, nil)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "restoring list items").
				With("restored_list_id", listID)
		}
	}

	return contents, nil
}

func (c Lists) DeleteList(
	ctx context.Context,
	siteID, listID string,
) error {
	if err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Delete(ctx, nil); err != nil {
		return graph.Wrap(ctx, err, "deleting list")
	}

	return nil
}

func CreateListFromBytes(bytes []byte) (models.Listable, error) {
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

		_, isLegacy := legacyColumns[displayName]

		// Skips columns that cannot be uploaded for models.ColumnDefinitionable:
		// - ReadOnly, Title, or Legacy columns: Attachments, Edit, or Content Type
		if readOnly || displayName == "Title" || isLegacy {
			continue
		}

		columns = append(columns, cloneColumnDefinitionable(cd))
	}

	newList.SetColumns(columns)

	return newList
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
	isColumnTypeSet := false

	if orig.GetText() != nil {
		newColumn.SetText(orig.GetText())

		isColumnTypeSet = true
	}

	if orig.GetBoolean() != nil {
		newColumn.SetBoolean(orig.GetBoolean())

		isColumnTypeSet = true
	}

	if orig.GetCalculated() != nil {
		newColumn.SetCalculated(orig.GetCalculated())

		isColumnTypeSet = true
	}

	if orig.GetChoice() != nil {
		newColumn.SetChoice(orig.GetChoice())

		isColumnTypeSet = true
	}

	if orig.GetContentApprovalStatus() != nil {
		newColumn.SetContentApprovalStatus(orig.GetContentApprovalStatus())

		isColumnTypeSet = true
	}

	if orig.GetCurrency() != nil {
		newColumn.SetCurrency(orig.GetCurrency())

		isColumnTypeSet = true
	}

	if orig.GetDateTime() != nil {
		newColumn.SetDateTime(orig.GetDateTime())

		isColumnTypeSet = true
	}

	if orig.GetGeolocation() != nil {
		newColumn.SetGeolocation(orig.GetGeolocation())

		isColumnTypeSet = true
	}

	if orig.GetHyperlinkOrPicture() != nil {
		newColumn.SetHyperlinkOrPicture(orig.GetHyperlinkOrPicture())

		isColumnTypeSet = true
	}

	if orig.GetNumber() != nil {
		newColumn.SetNumber(orig.GetNumber())

		isColumnTypeSet = true
	}

	if orig.GetLookup() != nil {
		newColumn.SetLookup(orig.GetLookup())

		isColumnTypeSet = true
	}

	if orig.GetThumbnail() != nil {
		newColumn.SetThumbnail(orig.GetThumbnail())

		isColumnTypeSet = true
	}

	if orig.GetTerm() != nil {
		newColumn.SetTerm(orig.GetTerm())

		isColumnTypeSet = true
	}

	if orig.GetPersonOrGroup() != nil {
		newColumn.SetPersonOrGroup(orig.GetPersonOrGroup())

		isColumnTypeSet = true
	}

	// defaulting to text type column
	if !isColumnTypeSet {
		textColumn := models.NewTextColumn()

		newColumn.SetText(textColumn)
	}
}

// retrieveFieldData utility function to clone raw listItem data from the embedded
// additionalData map
// Further details on FieldValueSets:
// - https://learn.microsoft.com/en-us/graph/api/resources/fieldvalueset?view=graph-rest-1.0
func retrieveFieldData(orig models.FieldValueSetable) models.FieldValueSetable {
	fields := models.NewFieldValueSet()
	additionalData := make(map[string]any)

	if orig != nil {
		fieldData := orig.GetAdditionalData()

		// M365 Book keeping values removed during new Item Creation
		// Removed Values:
		// -- Prefixes -> @odata.context : absolute path to previous list
		// .           -> @odata.etag : Embedded link to Prior M365 ID
		// -- String Match: Read-Only Fields
		// -> id : previous un

		for key, value := range fieldData {
			_, isReadOnlyField := readOnlyFieldNames[key]

			if strings.HasPrefix(key, "_") ||
				strings.HasPrefix(key, "@") ||
				isReadOnlyField ||
				strings.Contains(key, "LinkTitle") ||
				strings.Contains(key, "ChildCount") {
				continue
			}

			additionalData[key] = value
		}
	}

	retainPrimaryAddressField(additionalData)

	fields.SetAdditionalData(additionalData)

	return fields
}

func retainPrimaryAddressField(additionalData map[string]interface{}) {
	if !hasAddressFields(additionalData) {
		return
	}

	for k := range readonlyAddressFieldNames {
		delete(additionalData, k)
	}
}

func hasAddressFields(additionalData map[string]interface{}) bool {
	for _, value := range additionalData {
		if nestedFields, ok := value.(map[string]interface{}); ok &&
			hasRequiredFields(nestedFields, addressFieldNames) &&
			hasRequiredFields(additionalData, readonlyAddressFieldNames) {
			return true
		}
	}
	return false
}

func hasRequiredFields(data map[string]interface{}, checkFieldNames map[string]struct{}) bool {
	for field := range checkFieldNames {
		if _, exists := data[field]; !exists {
			return false
		}
	}
	return true
}
