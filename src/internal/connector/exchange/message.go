package exchange

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

func MessageInfo(msg models.Messageable, size int64) *details.ExchangeInfo {
	sender := ""
	subject := ""
	received := time.Time{}
	created := time.Time{}
	modified := time.Time{}

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

	if msg.GetCreatedDateTime() != nil {
		created = *msg.GetCreatedDateTime()
	}

	if msg.GetLastModifiedDateTime() != nil {
		modified = *msg.GetLastModifiedDateTime()
	}

	return &details.ExchangeInfo{
		ItemType: details.ExchangeMail,
		Sender:   sender,
		Subject:  subject,
		Received: received,
		Created:  created,
		Modified: modified,
		Size:     size,
	}
}
