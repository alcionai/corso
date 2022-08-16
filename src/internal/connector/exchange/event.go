package exchange

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/backup/details"
)

func EventInfo(evt models.Eventable) *details.ExchangeInfo {
	organizer := ""
	subject := ""
	received := time.Time{}

	if evt.GetOrganizer() != nil &&
		evt.GetOrganizer().GetEmailAddress() != nil &&
		evt.GetOrganizer().GetEmailAddress().GetAddress() != nil {
		organizer = *evt.GetOrganizer().
			GetEmailAddress().
			GetAddress()
	}
	if evt.GetSubject() != nil {
		subject = *evt.GetSubject()
	}
	if evt.GetStart() != nil &&
		evt.GetStart().GetDateTime() != nil {
		timeString := *evt.GetStart().GetDateTime()
		output, err := time.Parse(common.StandardTimeFormat, timeString)
		if err == nil {
			received = output
		}
	}
	return &details.ExchangeInfo{
		Sender:   organizer,
		Subject:  subject,
		Received: received,
	}

}
