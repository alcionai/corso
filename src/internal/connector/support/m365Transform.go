package support

import (
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

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
	additionalData := make(map[string]interface{})
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
