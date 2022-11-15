package exchange

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// EventInfo searchable metadata for stored event objects.
func EventInfo(evt models.Eventable, size int64) *details.ExchangeInfo {
	var (
		organizer, subject string
		recurs             bool
		start              = time.Time{}
		end                = time.Time{}
		created            = time.Time{}
		modified           = time.Time{}
	)

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

	if evt.GetRecurrence() != nil {
		recurs = true
	}

	if evt.GetStart() != nil &&
		evt.GetStart().GetDateTime() != nil {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		startTime := *evt.GetStart().GetDateTime() + "Z"

		output, err := common.ParseTime(startTime)
		if err == nil {
			start = output
		}
	}

	if evt.GetEnd() != nil &&
		evt.GetEnd().GetDateTime() != nil {
		// timeString has 'Z' literal added to ensure the stored
		// DateTime is not: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		endTime := *evt.GetEnd().GetDateTime() + "Z"

		output, err := common.ParseTime(endTime)
		if err == nil {
			end = output
		}
	}

	if evt.GetCreatedDateTime() != nil {
		created = *evt.GetCreatedDateTime()
	}

	if evt.GetLastModifiedDateTime() != nil {
		modified = *evt.GetLastModifiedDateTime()
	}

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeEvent,
		Organizer:   organizer,
		Subject:     subject,
		EventStart:  start,
		EventEnd:    end,
		EventRecurs: recurs,
		Created:     created,
		Modified:    modified,
		Size:        size,
	}
}
