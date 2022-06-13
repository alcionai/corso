package support

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func SwapMessage(orig models.Messageable) *models.Message {
	requestBody := models.NewMessage()
	subject := orig.GetSubject()
	requestBody.SetSubject(subject)
	bp := orig.GetBodyPreview()
	requestBody.SetBodyPreview(bp)
	body := orig.GetBody()
	requestBody.SetBody(body)
	sTime := orig.GetSentDateTime()
	rTime := orig.GetReceivedDateTime()
	requestBody.SetSentDateTime(sTime)
	requestBody.SetReceivedDateTime(rTime)
	requestBody.SetToRecipients(orig.GetToRecipients())
	requestBody.SetSender(orig.GetSender())
	requestBody.SetInferenceClassification(orig.GetInferenceClassification())
	requestBody.SetBccRecipients(orig.GetBccRecipients())
	requestBody.SetCcRecipients(orig.GetCcRecipients())
	requestBody.SetReplyTo(orig.GetReplyTo())
	requestBody.SetFlag(orig.GetFlag())
	requestBody.SetHasAttachments(orig.GetHasAttachments())
	requestBody.SetParentFolderId(orig.GetParentFolderId())
	notTrue := false
	requestBody.SetIsDraft(&notTrue)
	return requestBody

}
