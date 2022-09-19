package exchange

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

	now := initial.Format(common.StandardTimeFormat)
	suite.T().Logf("Initial: %v\nFormatted: %v\n", initial, now)

	tests := []struct {
		name     string
		evtAndRP func() (models.Eventable, *details.ExchangeInfo)
	}{
		{
			name: "Empty event",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				i := &details.ExchangeInfo{ItemType: details.ExchangeEvent}
				return models.NewEvent(), i
			},
		},
		{
			name: "Start time only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				event := models.NewEvent()
				dateTime := models.NewDateTimeTimeZone()
				dateTime.SetDateTime(&now)
				event.SetStart(dateTime)
				full, err := time.Parse(common.StandardTimeFormat, now)
				require.NoError(suite.T(), err)
				i := &details.ExchangeInfo{
					ItemType: details.ExchangeEvent,
					Received: full,
				}
				return event, i
			},
		},
		{
			name: "Subject Only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				subject := "Hello Corso"
				event := models.NewEvent()
				event.SetSubject(&subject)
				i := &details.ExchangeInfo{
					ItemType: details.ExchangeEvent,
					Subject:  subject,
				}
				return event, i
			},
		},
		{
			name: "Using mockable",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				bytes := mockconnector.GetMockEventBytes("Test Mock")
				event, err := support.CreateEventFromBytes(bytes)
				require.NoError(suite.T(), err)
				subject := " Test Mock Review + Lunch"
				organizer := "foobar3@8qzvrj.onmicrosoft.com"
				future := time.Now().AddDate(0, 0, 1)
				eventTime := time.Date(2022, future.Month(), future.Day(), 6, 0, 0, 0, time.UTC)
				i := &details.ExchangeInfo{
					ItemType:   details.ExchangeEvent,
					Subject:    subject,
					Organizer:  organizer,
					EventStart: eventTime,
				}
				return event, i
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			event, expected := test.evtAndRP()
			result := EventInfo(event)
			suite.Equal(expected.Subject, result.Subject)
			suite.Equal(expected.Sender, result.Sender)
			expYear, expMonth, _ := expected.EventStart.Date() // Day not used at certain times of the day
			expHr, expMin, expSec := expected.EventStart.Clock()
			recvYear, recvMonth, _ := result.EventStart.Date()
			recvHr, recvMin, recvSec := result.EventStart.Clock()
			suite.Equal(expYear, recvYear)
			suite.Equal(expMonth, recvMonth)
			suite.Equal(expHr, recvHr)
			suite.Equal(expMin, recvMin)
			suite.Equal(expSec, recvSec)
		})
	}
}
