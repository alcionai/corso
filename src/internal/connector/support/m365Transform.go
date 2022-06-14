package support

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ToMessage transfers all data from old message to new
// message except for the messageId.
func ToMessage(orig models.Messageable) *models.Message {
	message := models.NewMessage()
	message.SetSubject(orig.GetSubject())
	message.SetBodyPreview(orig.GetBodyPreview())
	message.SetBody(orig.GetBody())
	message.SetSentDateTime(orig.GetSentDateTime())
	message.SetReceivedDateTime(orig.GetReceivedDateTime())
	message.SetToRecipients(orig.GetToRecipients())
	message.SetSender(orig.GetSender())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetBccRecipients(orig.GetBccRecipients())
	message.SetCcRecipients(orig.GetCcRecipients())
	message.SetReplyTo(orig.GetReplyTo())
	message.SetFlag(orig.GetFlag())
	message.SetHasAttachments(orig.GetHasAttachments())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetConversationId(orig.GetConversationId())
	message.SetExtensions(orig.GetExtensions())
	message.SetFlag(orig.GetFlag())
	message.SetFrom(orig.GetFrom())
	message.SetImportance(orig.GetImportance())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetInternetMessageId(orig.GetInternetMessageId())
	message.SetInternetMessageHeaders(orig.GetInternetMessageHeaders())
	message.SetIsDeliveryReceiptRequested(orig.GetIsDeliveryReceiptRequested())
	message.SetIsRead(orig.GetIsRead())
	message.SetIsReadReceiptRequested(orig.GetIsReadReceiptRequested())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetMultiValueExtendedProperties(orig.GetMultiValueExtendedProperties())
	message.SetUniqueBody(orig.GetUniqueBody())
	message.SetWebLink(orig.GetWebLink())
	return message

}
