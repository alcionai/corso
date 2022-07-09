package exchange

import (
	"time"

	"github.com/alcionai/corso/pkg/backup"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func MessageInfo(msg models.Messageable) *backup.ExchangeInfo {
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
	return &backup.ExchangeInfo{
		Sender:   sender,
		Subject:  subject,
		Received: received,
	}
}
