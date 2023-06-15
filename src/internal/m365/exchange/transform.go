package exchange

import (
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

//==========================================================
// m365Transform.go contains utility functions that
// either add, modify, or remove fields from M365
// objects for interacton with M365 services
//=========================================================

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

func toMessage(orig models.Messageable) models.Messageable {
	message := models.NewMessage()
	return CloneMessageableFields(orig, message)
}

// ToEventSimplified transforms an event to simplified restore format
// To overcome some of the MS Graph API challenges, the event object is modified in the following ways:
//   - Instead of adding attendees and generating spurious notifications,
//     add a summary of attendees at the beginning to the event before the original body content
//   - event.attendees is set to an empty list
func toEventSimplified(orig models.Eventable) models.Eventable {
	attendees := FormatAttendees(orig, ptr.Val(orig.GetBody().GetContentType()) == models.HTML_BODYTYPE)
	orig.SetAttendees([]models.Attendeeable{})
	origBody := orig.GetBody()
	newContent := insertStringToBody(origBody, attendees)
	newBody := models.NewItemBody()
	newBody.SetContentType(origBody.GetContentType())
	newBody.SetOdataType(origBody.GetOdataType())
	newBody.SetContent(&newContent)
	orig.SetBody(newBody)
	// Sanitation steps for Events
	// See: https://github.com/alcionai/corso/issues/2490
	orig.SetTransactionId(nil)
	orig.SetWebLink(nil)
	orig.SetICalUId(nil)
	orig.SetId(nil)

	// Sanitize recurrence timezone.
	if orig.GetRecurrence() != nil {
		recurrenceTimezone := ptr.Val(orig.GetRecurrence().GetRange().GetRecurrenceTimeZone())
		if len(recurrenceTimezone) == 0 {
			orig.GetRecurrence().GetRange().SetRecurrenceTimeZone(nil)
		}
	}

	// Remove exceptions for recuring events
	// TODO(meain): Should we check if they exist before deleting?
	additionalData := origBody.GetAdditionalData()
	delete(additionalData, "cancelledOccurrences")
	delete(additionalData, "exceptionOccurrences")
	newBody.SetAdditionalData(additionalData)

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
		len(ptr.Val(body.GetContent())) == 0 ||
		body.GetContentType() == nil {
		return ""
	}

	content := ptr.Val(body.GetContent())

	switch ptr.Val(body.GetContentType()) {
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

// ===============================================================================================
// Sanitization section
// Set of functions that support ItemAttachemtable object restoration.
// These attachments can be nested as well as possess one of the other
// reference types. To ensure proper upload, each interior`item` requires
// that certain fields be modified.
// ItemAttachment:
// https://learn.microsoft.com/en-us/graph/api/resources/itemattachment?view=graph-rest-1.0
// https://learn.microsoft.com/en-us/exchange/client-developer/exchange-web-services/attachments-and-ews-in-exchange
// https://learn.microsoft.com/en-us/exchange/client-developer/exchange-web-services/folders-and-items-in-ews-in-exchange
// ===============================================================================================
// M365 Models possess a field, OData.Type which indicate
// the represent the intended model in string format.
// The constants listed here identify the supported itemAttachments
// currently supported for Restore operations.
// itemAttachments
// support ODataType values
//
//nolint:lll
const (
	itemAttachment  = "#microsoft.graph.itemAttachment"
	eventItemType   = "#microsoft.graph.event"
	mailItemType    = "#microsoft.graph.message"
	contactItemType = "#microsoft.graph.contact"
)

// toItemAttachment transforms internal item, OutlookItemables, into
// objects that are able to be uploaded into M365.
func toItemAttachment(orig models.Attachmentable) (models.Attachmentable, error) {
	transform, ok := orig.(models.ItemAttachmentable)
	if !ok { // Shouldn't ever happen
		return nil, clues.New("transforming attachment to item attachment")
	}

	item := transform.GetItem()
	itemType := ptr.Val(item.GetOdataType())

	switch itemType {
	case contactItemType:
		contact := item.(models.Contactable)
		revised := sanitizeContact(contact)

		transform.SetItem(revised)

		return transform, nil
	case eventItemType:
		event := item.(models.Eventable)

		newEvent, err := sanitizeEvent(event)
		if err != nil {
			return nil, err
		}

		transform.SetItem(newEvent)

		return transform, nil
	case mailItemType:
		message := item.(models.Messageable)

		newMessage, err := sanitizeMessage(message)
		if err != nil {
			return nil, err
		}

		transform.SetItem(newMessage)

		return transform, nil
	default:
		return nil, clues.New(fmt.Sprintf("unsupported attachment type: %T", itemType))
	}
}

// TODO #2428 (dadam39): re-apply nested attachments for itemAttachments
// func sanitizeAttachments(attached []models.Attachmentable) ([]models.Attachmentable, error) {
// 	attachments := make([]models.Attachmentable, len(attached))

// 	for _, ax := range attached {
// 		if ptr.Val(ax.GetOdataType()) == itemAttachment {
// 			newAttachment, err := ToItemAttachment(ax)
// 			if err != nil {
// 				return nil, err
// 			}

// 			attachments = append(attachments, newAttachment)

// 			continue
// 		}

// 		attachments = append(attachments, ax)
// 	}

// 	return attachments, nil
// }

// sanitizeContact removes fields which prevent a Contact from
// being uploaded as an attachment.
func sanitizeContact(orig models.Contactable) models.Contactable {
	orig.SetParentFolderId(nil)
	orig.SetAdditionalData(nil)

	return orig
}

// sanitizeEvent transfers data into event object and
// removes unique IDs from the M365 object
func sanitizeEvent(orig models.Eventable) (models.Eventable, error) {
	newEvent := models.NewEvent()
	newEvent.SetAttendees(orig.GetAttendees())
	newEvent.SetBody(orig.GetBody())
	newEvent.SetBodyPreview(orig.GetBodyPreview())
	newEvent.SetCalendar(orig.GetCalendar())
	newEvent.SetCreatedDateTime(orig.GetCreatedDateTime())
	newEvent.SetEnd(orig.GetEnd())
	// TODO: dadams39 Nested attachments not supported
	// Upstream: https://github.com/microsoft/kiota-serialization-json-go/issues/61
	newEvent.SetHasAttachments(nil)
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

	// Sanitation NOTE
	// isDraft and isOrganizer *bool ptr's have to be removed completely
	// from JSON in order for POST method to succeed.
	// Current as of 2/2/2023

	newEvent.SetIsOrganizer(nil)
	newEvent.SetIsDraft(nil)
	newEvent.SetAdditionalData(orig.GetAdditionalData())

	// TODO #2428 (dadam39): re-apply nested attachments for itemAttachments
	// Upstream: https://github.com/microsoft/kiota-serialization-json-go/issues/61
	// attachments, err := sanitizeAttachments(message.GetAttachments())
	// if err != nil {
	// 	return nil, err
	// }
	newEvent.SetAttachments(nil)

	return newEvent, nil
}

func sanitizeMessage(orig models.Messageable) (models.Messageable, error) {
	message := toMessage(orig)

	// TODO #2428 (dadam39): re-apply nested attachments for itemAttachments
	// Upstream: https://github.com/microsoft/kiota-serialization-json-go/issues/61
	// attachments, err := sanitizeAttachments(message.GetAttachments())
	// if err != nil {
	// 	return nil, err
	// }
	message.SetAttachments(nil)

	// The following fields are set to nil to
	// not interfere with M365 guard checks.
	message.SetHasAttachments(nil)
	message.SetParentFolderId(nil)
	message.SetInternetMessageHeaders(nil)
	message.SetIsDraft(nil)

	return message, nil
}
