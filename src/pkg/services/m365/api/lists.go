package api

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/keys"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var ErrSkippableListTemplate = clues.New("unable to create lists with skippable templates")

type columnDetails struct {
	createFieldName    string
	getFieldName       string
	isPersonColumn     bool
	isLookupColumn     bool
	isMultipleEnabled  bool
	hasDefaultedToText bool
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
// The full documentation concerning SharePoint Lists can
// be found at: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
// Note additional calls are required for the relationships that exist outside of the object properties.

// GetListById is a utility function to populate a SharePoint.List with objects associated with a given siteID.
// @param siteID the M365 ID that represents the SharePoint Site
// Makes additional calls to retrieve the following relationships:
// - Columns
// - ContentTypes
// - List Items
func (c Lists) GetListByID(ctx context.Context,
	siteID, listID string,
) (models.Listable, *details.SharePointInfo, error) {
	list, err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Wrap(ctx, err, "fetching list")
	}

	cols, cTypes, lItems, err := c.getListContents(ctx, siteID, listID)
	if err != nil {
		return nil, nil, graph.Wrap(ctx, err, "getting list contents")
	}

	list.SetColumns(cols)
	list.SetContentTypes(cTypes)
	list.SetItems(lItems)

	return list, ListToSPInfo(list), nil
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
	storedList models.Listable,
	errs *fault.Bus,
) (models.Listable, error) {
	el := errs.Local()

	// this ensure all columns, contentTypes are set to the newList
	newList, columnNames := ToListable(storedList, listName)

	if newList.GetList() != nil &&
		SkipListTemplates.HasKey(ptr.Val(newList.GetList().GetTemplate())) {
		return nil, clues.StackWC(ctx, ErrSkippableListTemplate)
	}

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

	for _, itm := range storedList.GetItems() {
		temp := CloneListItem(itm, columnNames)
		listItems = append(listItems, temp)
	}

	err = c.PostListItems(
		ctx,
		siteID,
		ptr.Val(restoredList.GetId()),
		listItems)
	if err != nil {
		err = graph.Wrap(ctx, err, "creating list item")
		el.AddRecoverable(ctx, err)
	}

	restoredList.SetItems(listItems)

	return restoredList, nil
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

func (c Lists) PatchList(
	ctx context.Context,
	siteID, listID string,
	list models.Listable,
) (models.Listable, error) {
	patchedList, err := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists().
		ByListId(listID).
		Patch(ctx, list, nil)

	return patchedList, graph.Wrap(ctx, err, "patching list").OrNil()
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
func ToListable(orig models.Listable, listName string) (models.Listable, map[string]*columnDetails) {
	newList := models.NewList()

	newList.SetContentTypes(orig.GetContentTypes())
	newList.SetCreatedBy(orig.GetCreatedBy())
	newList.SetCreatedByUser(orig.GetCreatedByUser())
	newList.SetCreatedDateTime(orig.GetCreatedDateTime())
	newList.SetDescription(orig.GetDescription())
	newList.SetDisplayName(ptr.To(listName))
	newList.SetLastModifiedBy(orig.GetLastModifiedBy())
	newList.SetLastModifiedByUser(orig.GetLastModifiedByUser())
	newList.SetLastModifiedDateTime(orig.GetLastModifiedDateTime())
	newList.SetList(orig.GetList())
	newList.SetOdataType(orig.GetOdataType())
	newList.SetParentReference(orig.GetParentReference())

	columns := make([]models.ColumnDefinitionable, 0)
	columnNames := map[string]*columnDetails{TitleColumnName: {
		getFieldName:    TitleColumnName,
		createFieldName: TitleColumnName,
	}}

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
		if readOnly || displayName == TitleColumnName || legacyColumns.HasKey(displayName) {
			continue
		}

		columns = append(columns, cloneColumnDefinitionable(cd, columnNames))
	}

	newList.SetColumns(columns)

	return newList, columnNames
}

// cloneColumnDefinitionable utility function for encapsulating models.ColumnDefinitionable data
// into new object for upload.
func cloneColumnDefinitionable(
	orig models.ColumnDefinitionable,
	columnNames map[string]*columnDetails,
) models.ColumnDefinitionable {
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
	setColumnType(newColumn, orig, columnNames)

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

func setColumnType(
	newColumn *models.ColumnDefinition,
	orig models.ColumnDefinitionable,
	columnNames map[string]*columnDetails,
) {
	colName := ptr.Val(newColumn.GetName())
	colDetails := &columnDetails{}

	// for certain columns like 'person', the column name is say 'personName'.
	// if the list item for that column holds single value,
	// the field data is fetched as '{"personNameLookupId": "10"}'
	// if the list item for that column holds multiple values,
	// the field data is fetched as '{"personName": [{"lookupId": 10}, {"lookupId": 11}]}'.
	// Hence this function helps us to determine which name to use while accessing stored data
	colDetails.getFieldName = colName
	colDetails.createFieldName = colName

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
		colDetails.isLookupColumn = true
		isMultipleEnabled := ptr.Val(orig.GetLookup().GetAllowMultipleValues())
		colDetails.isMultipleEnabled = isMultipleEnabled

		if isMultipleEnabled {
			colDetails.createFieldName = colName + LookupIDFieldNamePart
		} else {
			updatedName := colName + LookupIDFieldNamePart
			colDetails.getFieldName = updatedName
			colDetails.createFieldName = updatedName
		}

		newColumn.SetLookup(orig.GetLookup())
	case orig.GetThumbnail() != nil:
		newColumn.SetThumbnail(orig.GetThumbnail())
	case orig.GetTerm() != nil:
		newColumn.SetTerm(orig.GetTerm())
	case orig.GetPersonOrGroup() != nil:
		colDetails.isPersonColumn = true
		isMultipleEnabled := ptr.Val(orig.GetPersonOrGroup().GetAllowMultipleSelection())
		colDetails.isMultipleEnabled = isMultipleEnabled

		if isMultipleEnabled {
			colDetails.createFieldName = colName + LookupIDFieldNamePart
		} else {
			updatedName := colName + LookupIDFieldNamePart
			colDetails.getFieldName = updatedName
			colDetails.createFieldName = updatedName
		}

		newColumn.SetPersonOrGroup(orig.GetPersonOrGroup())
	default:
		colDetails.hasDefaultedToText = true

		newColumn.SetText(models.NewTextColumn())
	}

	columnNames[colName] = colDetails
}

// CloneListItem creates a new `SharePoint.ListItem` and stores the original item's
// M365 data into it set fields.
// - https://learn.microsoft.com/en-us/graph/api/resources/listitem?view=graph-rest-1.0
func CloneListItem(orig models.ListItemable, columnNames map[string]*columnDetails) models.ListItemable {
	newItem := models.NewListItem()

	// list item data
	newFieldData := retrieveFieldData(orig.GetFields(), columnNames)
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
// Further documentation on FieldValueSets:
// - https://learn.microsoft.com/en-us/graph/api/resources/fieldvalueset?view=graph-rest-1.0
func retrieveFieldData(orig models.FieldValueSetable, columnNames map[string]*columnDetails) models.FieldValueSetable {
	fields := models.NewFieldValueSet()

	additionalData := setAdditionalDataByColumnNames(orig, columnNames)
	if addressField, fieldName, ok := hasAddressFields(additionalData); ok {
		concatenatedAddress := concatenateAddressFields(addressField)
		additionalData[fieldName] = concatenatedAddress
	}

	if hyperLinkField, fieldName, ok := hasHyperLinkFields(additionalData); ok {
		concatenatedHyperlink := concatenateHyperLinkFields(hyperLinkField)
		additionalData[fieldName] = concatenatedHyperlink
	}

	if metadataField, fieldName, ok := hasMetadataFields(additionalData); ok {
		additionalData[fieldName] = concatenateMetadataFields(metadataField)
	}

	fields.SetAdditionalData(additionalData)

	return fields
}

func setAdditionalDataByColumnNames(
	orig models.FieldValueSetable,
	columnNames map[string]*columnDetails,
) map[string]any {
	if orig == nil {
		return make(map[string]any)
	}

	fieldData := orig.GetAdditionalData()
	filteredData := make(map[string]any)

	for _, colDetails := range columnNames {
		if val, ok := fieldData[colDetails.getFieldName]; ok {
			setMultipleEnabledByFieldData(val, colDetails)
			filteredData[colDetails.createFieldName] = val
			populateMultipleValues(val, filteredData, colDetails)
		}

		specifyODataType(filteredData, colDetails, colDetails.createFieldName)
	}

	return filteredData
}

func populateMultipleValues(val any, filteredData map[string]any, colDetails *columnDetails) {
	if !colDetails.isMultipleEnabled {
		return
	}

	if !colDetails.isPersonColumn &&
		!colDetails.isLookupColumn {
		return
	}

	multiNestedFields, ok := val.([]any)
	if !ok || len(multiNestedFields) == 0 {
		return
	}

	lookupIDs := make([]float64, 0)

	checkFields := func(colDetails *columnDetails) []string {
		if colDetails.isLookupColumn {
			return []string{LookupIDKey, LookupValueKey}
		}

		if colDetails.isPersonColumn {
			return []string{LookupIDKey, LookupValueKey, PersonEmailKey}
		}

		return []string{}
	}

	for _, nestedFields := range multiNestedFields {
		if md, ok := nestedFields.(map[string]any); ok && keys.HasKeys(md, checkFields(colDetails)...) {
			lookupID, ok := md[LookupIDKey].(*float64)
			if !ok {
				continue
			}

			lookupIDs = append(lookupIDs, ptr.Val(lookupID))
		}
	}

	filteredData[colDetails.createFieldName] = lookupIDs
}

func setMultipleEnabledByFieldData(val any, colDetails *columnDetails) {
	// already set while column definition
	// not required to determined from field values
	if colDetails.isMultipleEnabled {
		return
	}

	// for columns like 'choice', even though it has an option to hold single/multiple values,
	// the columnDefinition property 'allowMultipleValues' is not available.
	// Hence we determine single/multiple from the actual field data.
	if reflect.TypeOf(val).Kind() == reflect.Slice {
		colDetails.isMultipleEnabled = true
	}
}

// when creating list items with multiple values for a single column
// we let the API know that we are sending a collection.
// Hence this adds an additional field '<columnName>@@odata.type'
// with value depending on type of column.
func specifyODataType(filteredData map[string]any, colDetails *columnDetails, colName string) {
	// text column itself does not allow holding multiple values
	// some columns like 'term'/'managed metadata' have,
	//  but they get defaulted to text column.
	if colDetails.hasDefaultedToText {
		return
	}

	// only specify odata.type for columns holding multiple data
	if !colDetails.isMultipleEnabled {
		return
	}

	switch {
	case colDetails.isPersonColumn, colDetails.isLookupColumn:
		filteredData[colName+ODataTypeFieldNamePart] = ODataTypeFieldNameIntVal
	default:
		filteredData[colName+ODataTypeFieldNamePart] = ODataTypeFieldNameStringVal
	}
}

func hasAddressFields(additionalData map[string]any) (map[string]any, string, bool) {
	for k, v := range additionalData {
		nestedFields, ok := v.(map[string]any)
		if !ok || keys.HasKeys(nestedFields, GeoLocFN) {
			continue
		}

		if keys.HasKeys(nestedFields, addressFieldNames...) {
			return nestedFields, k, true
		}
	}

	return nil, "", false
}

func hasHyperLinkFields(additionalData map[string]any) (map[string]any, string, bool) {
	for fieldName, value := range additionalData {
		nestedFields, ok := value.(map[string]any)
		if !ok {
			continue
		}

		if keys.HasKeys(nestedFields,
			[]string{HyperlinkDescriptionKey, HyperlinkURLKey}...) {
			return nestedFields, fieldName, true
		}
	}

	return nil, "", false
}

func hasMetadataFields(additionalData map[string]any) ([]map[string]any, string, bool) {
	for fieldName, value := range additionalData {
		switch valType := reflect.TypeOf(value).Kind(); valType {
		case reflect.Map:
			metadataFields, areMetadataFields := getMetadataFields(value)
			if areMetadataFields {
				return []map[string]any{metadataFields}, fieldName, true
			}

		case reflect.Slice:
			mmdfs := make([]map[string]any, 0)

			multiMetadataFields, ok := value.([]any)
			if !ok {
				continue
			}

			for _, mdfs := range multiMetadataFields {
				metadataFields, areMetadataFields := getMetadataFields(mdfs)
				if areMetadataFields {
					mmdfs = append(mmdfs, metadataFields)
				}
			}

			if len(mmdfs) > 0 {
				return mmdfs, fieldName, true
			}
		}
	}

	return nil, "", false
}

func getMetadataFields(metadataFieldvalue any) (map[string]any, bool) {
	nestedFields, ok := metadataFieldvalue.(map[string]any)
	return nestedFields, ok && keys.HasKeys(nestedFields, MetadataLabelKey, MetadataTermGUIDKey, MetadataWssIDKey)
}

func concatenateAddressFields(addressFields map[string]any) string {
	parts := make([]string, 0)

	if dispName, ok := addressFields[DisplayNameKey].(*string); ok {
		parts = append(parts, ptr.Val(dispName))
	}

	if fields, ok := addressFields[AddressKey].(map[string]any); ok {
		parts = append(parts, addressKeyToVal(fields, StreetKey))
		parts = append(parts, addressKeyToVal(fields, CityKey))
		parts = append(parts, addressKeyToVal(fields, StateKey))
		parts = append(parts, addressKeyToVal(fields, CountryKey))
		parts = append(parts, addressKeyToVal(fields, PostalCodeKey))
	}

	if coords, ok := addressFields[CoordinatesKey].(map[string]any); ok {
		parts = append(parts, addressKeyToVal(coords, LatitudeKey))
		parts = append(parts, addressKeyToVal(coords, LongitudeKey))
	}

	if len(parts) > 0 {
		return strings.Join(parts, ",")
	}

	return ""
}

func concatenateHyperLinkFields(hyperlinkFields map[string]any) string {
	parts := make([]string, 0)

	if v, err := str.AnyValueToString(HyperlinkURLKey, hyperlinkFields); err == nil {
		parts = append(parts, v)
	}

	if v, err := str.AnyValueToString(HyperlinkDescriptionKey, hyperlinkFields); err == nil {
		parts = append(parts, v)
	}

	if len(parts) > 0 {
		return strings.Join(parts, ",")
	}

	return ""
}

func concatenateMetadataFields(metadataFieldsArr []map[string]any) string {
	labels := make([]string, 0)

	for _, md := range metadataFieldsArr {
		mdVal, ok := md[MetadataLabelKey].(*string)
		if ok {
			labels = append(labels, ptr.Val(mdVal))
		}
	}

	if len(labels) > 0 {
		return strings.Join(labels, ",")
	}

	return ""
}

func addressKeyToVal(fields map[string]any, key string) string {
	if v, err := str.AnyValueToString(key, fields); err == nil {
		return v
	} else if v, ok := fields[key].(*float64); ok {
		return fmt.Sprintf("%v", ptr.Val(v))
	}

	return ""
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

// ListToSPInfo translates models.Listable metadata into searchable content
// List documentation: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
func ListToSPInfo(lst models.Listable) *details.SharePointInfo {
	var (
		name     = ptr.Val(lst.GetDisplayName())
		webURL   = ptr.Val(lst.GetWebUrl())
		created  = ptr.Val(lst.GetCreatedDateTime())
		modified = ptr.Val(lst.GetLastModifiedDateTime())
		count    = len(lst.GetItems())
	)

	template := ""
	if lst.GetList() != nil {
		template = ptr.Val(lst.GetList().GetTemplate())
	}

	return &details.SharePointInfo{
		ItemType: details.SharePointList,
		Modified: modified,
		Created:  created,
		WebURL:   webURL,
		List: &details.ListInfo{
			Name:      name,
			ItemCount: int64(count),
			Template:  template,
		},
	}
}

func listCollisionKeyProps() []string {
	return idAnd("displayName")
}

// Two lists with same name cannot be created,
// hence going by the displayName itself as the collision key.
// Only displayName can be set. name is system generated based on displayName.
func ListCollisionKey(list models.Listable) string {
	if list == nil {
		return ""
	}

	return ptr.Val(list.GetDisplayName())
}
