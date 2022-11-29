package exchange

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type EventSuite struct {
	suite.Suite
}

func TestEventSuite(t *testing.T) {
	suite.Run(t, &EventSuite{})
}

// TestEventInfo verifies that searchable event metadata
// can be properly retrieved from a models.Eventable object
func (suite *EventSuite) TestEventInfo() {
	// Exchange stores start/end times in UTC and the below compares hours
	// directly so we need to "normalize" the timezone here.
	initial := time.Now().UTC()
	now := common.FormatTimeWith(initial, common.M365DateTimeTimeZone)

	suite.T().Logf("Initial: %v\nFormatted: %v\n", initial, now)

	tests := []struct {
		name     string
		evtAndRP func() (models.Eventable, *details.ExchangeInfo)
	}{
		{
			name: "Empty event",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				event := models.NewEvent()

				// Start and Modified will always be available in API
				event.SetCreatedDateTime(&initial)
				event.SetLastModifiedDateTime(&initial)

				return event, &details.ExchangeInfo{
					ItemType: details.ExchangeEvent,
				}
			},
		},
		{
			name: "Start time only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				var (
					event    = models.NewEvent()
					dateTime = models.NewDateTimeTimeZone()
				)

				event.SetCreatedDateTime(&initial)
				event.SetLastModifiedDateTime(&initial)
				dateTime.SetDateTime(&now)
				event.SetStart(dateTime)

				return event, &details.ExchangeInfo{
					ItemType:   details.ExchangeEvent,
					Received:   initial,
					EventStart: initial,
				}
			},
		},
		{
			name: "Start and end time only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				var (
					event     = models.NewEvent()
					startTime = models.NewDateTimeTimeZone()
					endTime   = models.NewDateTimeTimeZone()
				)

				event.SetCreatedDateTime(&initial)
				event.SetLastModifiedDateTime(&initial)
				startTime.SetDateTime(&now)
				event.SetStart(startTime)

				nowp30m := common.FormatTimeWith(initial.Add(30*time.Minute), common.M365DateTimeTimeZone)
				endTime.SetDateTime(&nowp30m)
				event.SetEnd(endTime)

				return event, &details.ExchangeInfo{
					ItemType:   details.ExchangeEvent,
					Received:   initial,
					EventStart: initial,
					EventEnd:   initial.Add(30 * time.Minute),
				}
			},
		},
		{
			name: "Subject Only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				var (
					subject = "Hello Corso"
					event   = models.NewEvent()
				)

				event.SetCreatedDateTime(&initial)
				event.SetLastModifiedDateTime(&initial)
				event.SetSubject(&subject)

				return event, &details.ExchangeInfo{
					ItemType: details.ExchangeEvent,
					Subject:  subject,
				}
			},
		},
		{
			name: "Using mockable",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				var (
					organizer    = "foobar3@8qzvrj.onmicrosoft.com"
					subject      = " Test Mock Review + Lunch"
					bytes        = mockconnector.GetDefaultMockEventBytes("Test Mock")
					future       = time.Now().UTC().AddDate(0, 0, 1)
					eventTime    = time.Date(future.Year(), future.Month(), future.Day(), future.Hour(), 0, 0, 0, time.UTC)
					eventEndTime = eventTime.Add(30 * time.Minute)
					event, err   = support.CreateEventFromBytes(bytes)
				)

				require.NoError(suite.T(), err)

				return event, &details.ExchangeInfo{
					ItemType:   details.ExchangeEvent,
					Subject:    subject,
					Organizer:  organizer,
					EventStart: eventTime,
					EventEnd:   eventEndTime,
					Size:       10,
				}
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			event, expected := test.evtAndRP()
			result := EventInfo(event, 10)

			assert.Equal(t, expected.Subject, result.Subject, "subject")
			assert.Equal(t, expected.Sender, result.Sender, "sender")

			expStartYear, expStartMonth, _ := expected.EventStart.Date() // Day not used at certain times of the day
			expStartHr, expStartMin, expStartSec := expected.EventStart.Clock()
			recvStartYear, recvStartMonth, _ := result.EventStart.Date()
			recvStartHr, recvStartMin, recvStartSec := result.EventStart.Clock()

			assert.Equal(t, expStartYear, recvStartYear, "year")
			assert.Equal(t, expStartMonth, recvStartMonth, "month")
			assert.Equal(t, expStartHr, recvStartHr, "hour")
			assert.Equal(t, expStartMin, recvStartMin, "minute")
			assert.Equal(t, expStartSec, recvStartSec, "second")

			expEndYear, expEndMonth, _ := expected.EventEnd.Date() // Day not used at certain times of the day
			expEndHr, expEndMin, expEndSec := expected.EventEnd.Clock()
			recvEndYear, recvEndMonth, _ := result.EventEnd.Date()
			recvEndHr, recvEndMin, recvEndSec := result.EventEnd.Clock()

			assert.Equal(t, expEndYear, recvEndYear, "year")
			assert.Equal(t, expEndMonth, recvEndMonth, "month")
			assert.Equal(t, expEndHr, recvEndHr, "hour")
			assert.Equal(t, expEndMin, recvEndMin, "minute")
			assert.Equal(t, expEndSec, recvEndSec, "second")
		})
	}
}
