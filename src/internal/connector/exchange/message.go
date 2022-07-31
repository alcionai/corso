package exchange

import (
	"time"

	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func MessageInfo(msg models.Messageable) *details.ExchangeInfo {
	sender := ""
	subject := ""
	received := time.Time{}
	if msg.GetSender() != nil &&
		msg.GetSender().GetEmailAddress() != nil &&
		msg.GetSender().GetEmailAddress().GetAddress() != nil {
		sender = *msg.GetSender().GetEmailAddress().GetAddress()
	}
	if msg.GetSubject() != nil {
		subject = *msg.GetSubject()
	}
	if msg.GetReceivedDateTime() != nil {
		received = *msg.GetReceivedDateTime()
	}
	return &details.ExchangeInfo{
		Sender:   sender,
		Subject:  subject,
		Received: received,
	}
}
