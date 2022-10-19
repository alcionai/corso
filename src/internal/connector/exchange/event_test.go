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
	initial := time.Now()
	now := common.FormatTime(initial)

	suite.T().Logf("Initial: %v\nFormatted: %v\n", initial, now)

	tests := []struct {
		name     string
		evtAndRP func() (models.Eventable, *details.ExchangeInfo)
	}{
		{
			name: "Empty event",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				return models.NewEvent(), &details.ExchangeInfo{
					ItemType: details.ExchangeEvent,
				}
			},
		},
		{
			name: "Start time only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				var (
					event     = models.NewEvent()
					dateTime  = models.NewDateTimeTimeZone()
					full, err = common.ParseTime(now)
				)

				require.NoError(suite.T(), err)

				dateTime.SetDateTime(&now)
				event.SetStart(dateTime)

				return event, &details.ExchangeInfo{
					ItemType: details.ExchangeEvent,
					Received: full,
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
					organizer  = "foobar3@8qzvrj.onmicrosoft.com"
					subject    = " Test Mock Review + Lunch"
					bytes      = mockconnector.GetDefaultMockEventBytes("Test Mock")
					future     = time.Now().UTC().AddDate(0, 0, 1)
					eventTime  = time.Date(future.Year(), future.Month(), future.Day(), future.Hour(), 0, 0, 0, time.UTC)
					event, err = support.CreateEventFromBytes(bytes)
				)

				require.NoError(suite.T(), err)

				return event, &details.ExchangeInfo{
					ItemType:   details.ExchangeEvent,
					Subject:    subject,
					Organizer:  organizer,
					EventStart: eventTime,
				}
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			event, expected := test.evtAndRP()
			result := EventInfo(event)

			assert.Equal(t, expected.Subject, result.Subject, "subject")
			assert.Equal(t, expected.Sender, result.Sender, "sender")

			expYear, expMonth, _ := expected.EventStart.Date() // Day not used at certain times of the day
			expHr, expMin, expSec := expected.EventStart.Clock()
			recvYear, recvMonth, _ := result.EventStart.Date()
			recvHr, recvMin, recvSec := result.EventStart.Clock()

			assert.Equal(t, expYear, recvYear, "year")
			assert.Equal(t, expMonth, recvMonth, "month")
			assert.Equal(t, expHr, recvHr, "hour")
			assert.Equal(t, expMin, recvMin, "minute")
			assert.Equal(t, expSec, recvSec, "second")
		})
	}
}
