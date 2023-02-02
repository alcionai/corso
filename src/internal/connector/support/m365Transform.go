package support

import (
	"fmt"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

const itmAttachment = "#microsoft.graph.itemAttachment"

// CloneMessageableFields places data from original data into new message object.
// SingleLegacyValueProperty is not populated during this operation
func CloneMessageableFields(orig, message models.Messageable) models.Messageable {
	message.SetAttachments(orig.GetAttachments())
	message.SetAdditionalData(orig.GetAdditionalData())
	message.SetBccRecipients(orig.GetBccRecipients())
	message.SetBody(orig.GetBody())
	message.SetBodyPreview(orig.GetBodyPreview())
	message.SetCcRecipients(orig.GetCcRecipients())
	message.SetConversationId(orig.GetConversationId())
	message.SetConversationIndex(orig.GetConversationIndex())
	message.SetExtensions(orig.GetExtensions())
	message.SetFlag(orig.GetFlag())
	message.SetFrom(orig.GetFrom())
	message.SetHasAttachments(orig.GetHasAttachments())
	message.SetImportance(orig.GetImportance())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetInternetMessageHeaders(orig.GetInternetMessageHeaders())
	message.SetInternetMessageId(orig.GetInternetMessageId())
	message.SetIsDeliveryReceiptRequested(orig.GetIsDeliveryReceiptRequested())
	message.SetIsDraft(orig.GetIsDraft())
	message.SetIsRead(orig.GetIsRead())
	message.SetIsReadReceiptRequested(orig.GetIsReadReceiptRequested())
	message.SetMultiValueExtendedProperties(orig.GetMultiValueExtendedProperties())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetReceivedDateTime(orig.GetReceivedDateTime())
	message.SetReplyTo(orig.GetReplyTo())
	message.SetSender(orig.GetSender())
	message.SetSentDateTime(orig.GetSentDateTime())
	message.SetSubject(orig.GetSubject())
	message.SetToRecipients(orig.GetToRecipients())
	message.SetUniqueBody(orig.GetUniqueBody())
	message.SetWebLink(orig.GetWebLink())

	return message
}

func ToMessage(orig models.Messageable) models.Messageable {
	message := models.NewMessage()
	temp := CloneMessageableFields(orig, message)

	aMessage, ok := temp.(*models.Message)
	if !ok {
		return nil
	}

	return aMessage
}

// ToEventSimplified transforms an event to simplifed restore format
// To overcome some of the MS Graph API challenges, the event object is modified in the following ways:
//   - Instead of adding attendees and generating spurious notifications,
//     add a summary of attendees at the beginning to the event before the original body content
//   - event.attendees is set to an empty list
func ToEventSimplified(orig models.Eventable) models.Eventable {
	attendees := FormatAttendees(orig, *orig.GetBody().GetContentType() == models.HTML_BODYTYPE)
	orig.SetAttendees([]models.Attendeeable{})
	origBody := orig.GetBody()
	newContent := insertStringToBody(origBody, attendees)
	newBody := models.NewItemBody()
	newBody.SetContentType(origBody.GetContentType())
	newBody.SetAdditionalData(origBody.GetAdditionalData())
	newBody.SetOdataType(origBody.GetOdataType())
	newBody.SetContent(&newContent)
	orig.SetBody(newBody)

	return orig
}

type getContenter interface {
	GetContent() *string
	GetContentType() *models.BodyType
}

// insertStringToBody helper function to insert text into models.bodyable
// @returns string containing the content string of altered body.
func insertStringToBody(body getContenter, newContent string) string {
	if body.GetContent() == nil ||
		len(*body.GetContent()) == 0 ||
		body.GetContentType() == nil {
		return ""
	}

	content := *body.GetContent()

	switch *body.GetContentType() {
	case models.TEXT_BODYTYPE:
		return newContent + content

	case models.HTML_BODYTYPE:
		arr := strings.Split(content, "<body>")
		if len(arr) < 2 {
			// malformed html; can't be sure where to insert attendees.
			return newContent + content
		}

		prefix := arr[0] + "<body>"
		interior := arr[1]
		splitOnCloseAngle := strings.Split(interior, ">")

		if len(splitOnCloseAngle) < 3 {
			// no inner elements in body, just insert the new content
			return prefix + newContent + strings.Join(arr[1:], "")
		}

		prefix += splitOnCloseAngle[0] + ">"
		suffix := strings.Join(splitOnCloseAngle[1:], ">")

		return prefix + newContent + suffix
	}

	return newContent + content
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

		if cd.GetDisplayName() != nil {
			displayName = *cd.GetDisplayName()
		}

		if cd.GetReadOnly() != nil {
			readOnly = *cd.GetReadOnly()
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

// ToItemAttachment transforms internal item, OutlookItemables, into
// objects that are able to be uploaded into M365.
// Supported Internal Items:
// - Events
func ToItemAttachment(orig models.Attachmentable) (models.Attachmentable, error) {
	// First things first. Find out where all the information is
	transform, ok := orig.(models.ItemAttachmentable)
	supported := "#microsoft.graph.event"

	if !ok {
		return nil, fmt.Errorf("transforming attachment to item attachment")
	}

	item := transform.GetItem()
	itemType := item.GetOdataType()

	switch *itemType {
	case supported:
		event := item.(models.Eventable)

		newEvent, err := sanitizeEvent(event)
		if err != nil {
			return nil, err
		}

		transform.SetItem(newEvent)

		return transform, nil
	default:
		return nil, fmt.Errorf("")
	}
}

func sanitizeEvent(orig models.Eventable) (models.Eventable, error) {
	newEvent := models.NewEvent()
	newEvent.SetAttendees(orig.GetAttendees())
	newEvent.SetBody(orig.GetBody())
	newEvent.SetBodyPreview(orig.GetBodyPreview())
	newEvent.SetCalendar(orig.GetCalendar())
	newEvent.SetCreatedDateTime(orig.GetCreatedDateTime())
	newEvent.SetEnd(orig.GetEnd())
	newEvent.SetHasAttachments(orig.GetHasAttachments())
	newEvent.SetHideAttendees(orig.GetHideAttendees())
	newEvent.SetImportance(orig.GetImportance())
	newEvent.SetIsAllDay(orig.GetIsAllDay())
	newEvent.SetIsOnlineMeeting(orig.GetIsOnlineMeeting())
	newEvent.SetLocation(orig.GetLocation())
	newEvent.SetLocations(orig.GetLocations())
	newEvent.SetSensitivity(orig.GetSensitivity())
	newEvent.SetReminderMinutesBeforeStart(orig.GetReminderMinutesBeforeStart())
	newEvent.SetStart(orig.GetStart())
	newEvent.SetSubject(orig.GetSubject())
	newEvent.SetType(orig.GetType())

	// Sanitation
	adtl := orig.GetAdditionalData()
	adtl["isOrganizer"] = orig.GetIsOrganizer()
	adtl["isDraft"] = orig.GetIsDraft()

	newEvent.SetIsOrganizer(nil)
	newEvent.SetIsDraft(nil)
	newEvent.SetAdditionalData(adtl)

	attached := orig.GetAttachments()
	attachments := make([]models.Attachmentable, len(attached))

	for _, ax := range attached {
		if *ax.GetOdataType() == "#microsoft.graph.itemAttachment" {
			newAttachment, err := ToItemAttachment(ax)
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, newAttachment)

			continue
		}

		attachments = append(attachments, ax)
	}

	newEvent.SetAttachments(attachments)

	return newEvent, nil
}
