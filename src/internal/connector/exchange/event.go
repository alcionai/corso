package exchange

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/backup/details"
)

// EventInfo searchable metadata for stored event objects.
func EventInfo(evt models.Eventable) *details.ExchangeInfo {
	organizer := ""
	subject := ""
	start := time.Time{}

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
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		timeString := *evt.GetStart().GetDateTime() + "Z"

		output, err := common.ParseTime(timeString)
		if err == nil {
			start = output
		}
	}

	return &details.ExchangeInfo{
		EventRecurs: evt.GetSeriesMasterId() != nil,
		EventStart:  start,
		Organizer:   organizer,
		Subject:     subject,
	}
}
