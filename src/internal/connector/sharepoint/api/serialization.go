package api

import (
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	betamodels "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// createFromBytes generates an m365 object form bytes.
func createFromBytes(
	bytes []byte,
	createFunc serialization.ParsableFactory,
) (serialization.Parsable, error) {
	parseNode, err := kjson.NewJsonParseNodeFactory().GetRootParseNode("application/json", bytes)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes into base m365 object")
	}

	v, err := parseNode.GetObjectValue(createFunc)
	if err != nil {
		return nil, clues.Wrap(err, "parsing m365 object factory")
	}

	return v, nil
}

func CreateListFromBytes(bytes []byte) (models.Listable, error) {
	parsable, err := createFromBytes(bytes, models.CreateListFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to sharepoint list")
	}

	list := parsable.(models.Listable)

	return list, nil
}

func CreatePageFromBytes(bytes []byte) (betamodels.SitePageable, error) {
	parsable, err := createFromBytes(bytes, betamodels.CreateSitePageFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to sharepoint page")
	}

	page := parsable.(betamodels.SitePageable)

	return page, nil
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
	leg := map[string]struct{}{
		"Attachments":  {},
		"Edit":         {},
		"Content Type": {},
	}

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

		_, isLegacy := leg[displayName]

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

// cloneColumnDefinitionable utility function for encapsulating models.ColumnDefinitionable data
// into new object for upload.
func cloneColumnDefinitionable(orig models.ColumnDefinitionable) models.ColumnDefinitionable {
	newColumn := models.NewColumnDefinition()

	newColumn.SetAdditionalData(orig.GetAdditionalData())
	newColumn.SetBoolean(orig.GetBoolean())
	newColumn.SetCalculated(orig.GetCalculated())
	newColumn.SetChoice(orig.GetChoice())
	newColumn.SetColumnGroup(orig.GetColumnGroup())
	newColumn.SetContentApprovalStatus(orig.GetContentApprovalStatus())
	newColumn.SetCurrency(orig.GetCurrency())
	newColumn.SetDateTime(orig.GetDateTime())
	newColumn.SetDefaultValue(orig.GetDefaultValue())
	newColumn.SetDescription(orig.GetDescription())
	newColumn.SetDisplayName(orig.GetDisplayName())
	newColumn.SetEnforceUniqueValues(orig.GetEnforceUniqueValues())
	newColumn.SetGeolocation(orig.GetGeolocation())
	newColumn.SetHidden(orig.GetHidden())
	newColumn.SetHyperlinkOrPicture(orig.GetHyperlinkOrPicture())
	newColumn.SetIndexed(orig.GetIndexed())
	newColumn.SetIsDeletable(orig.GetIsDeletable())
	newColumn.SetIsReorderable(orig.GetIsReorderable())
	newColumn.SetIsSealed(orig.GetIsSealed())
	newColumn.SetLookup(orig.GetLookup())
	newColumn.SetName(orig.GetName())
	newColumn.SetNumber(orig.GetNumber())
	newColumn.SetOdataType(orig.GetOdataType())
	newColumn.SetPersonOrGroup(orig.GetPersonOrGroup())
	newColumn.SetPropagateChanges(orig.GetPropagateChanges())
	newColumn.SetReadOnly(orig.GetReadOnly())
	newColumn.SetRequired(orig.GetRequired())
	newColumn.SetSourceColumn(orig.GetSourceColumn())
	newColumn.SetSourceContentType(orig.GetSourceContentType())
	newColumn.SetTerm(orig.GetTerm())
	newColumn.SetText(orig.GetText())
	newColumn.SetThumbnail(orig.GetThumbnail())
	newColumn.SetType(orig.GetType())
	newColumn.SetValidation(orig.GetValidation())

	return newColumn
}

// CloneListItem creates a new `SharePoint.ListItem` and stores the original item's
// M365 data into it set fields.
// - https://learn.microsoft.com/en-us/graph/api/resources/listitem?view=graph-rest-1.0
func CloneListItem(orig models.ListItemable) models.ListItemable {
	newItem := models.NewListItem()
	newFieldData := retrieveFieldData(orig.GetFields())

	newItem.SetAdditionalData(orig.GetAdditionalData())
	newItem.SetAnalytics(orig.GetAnalytics())
	newItem.SetContentType(orig.GetContentType())
	newItem.SetCreatedBy(orig.GetCreatedBy())
	newItem.SetCreatedByUser(orig.GetCreatedByUser())
	newItem.SetCreatedDateTime(orig.GetCreatedDateTime())
	newItem.SetDescription(orig.GetDescription())
	// ETag cannot be carried forward
	newItem.SetFields(newFieldData)
	newItem.SetLastModifiedBy(orig.GetLastModifiedBy())
	newItem.SetLastModifiedByUser(orig.GetLastModifiedByUser())
	newItem.SetLastModifiedDateTime(orig.GetLastModifiedDateTime())
	newItem.SetOdataType(orig.GetOdataType())
	// parentReference and SharePointIDs cause error on upload.
	// POST Command will link items to the created list.
	newItem.SetVersions(orig.GetVersions())

	return newItem
}

// retrieveFieldData utility function to clone raw listItem data from the embedded
// additionalData map
// Further details on FieldValueSets:
// - https://learn.microsoft.com/en-us/graph/api/resources/fieldvalueset?view=graph-rest-1.0
func retrieveFieldData(orig models.FieldValueSetable) models.FieldValueSetable {
	fields := models.NewFieldValueSet()
	additionalData := make(map[string]any)
	fieldData := orig.GetAdditionalData()

	// M365 Book keeping values removed during new Item Creation
	// Removed Values:
	// -- Prefixes -> @odata.context : absolute path to previous list
	// .           -> @odata.etag : Embedded link to Prior M365 ID
	// -- String Match: Read-Only Fields
	// -> id : previous un
	for key, value := range fieldData {
		if strings.HasPrefix(key, "_") || strings.HasPrefix(key, "@") ||
			key == "Edit" || key == "Created" || key == "Modified" ||
			strings.Contains(key, "LookupId") || strings.Contains(key, "ChildCount") || strings.Contains(key, "LinkTitle") {
			continue
		}

		additionalData[key] = value
	}

	fields.SetAdditionalData(additionalData)

	return fields
}
