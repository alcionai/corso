package exchange

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/backup/details"
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
				return models.NewEvent(), &details.ExchangeInfo{}
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
				return event, &details.ExchangeInfo{Received: full}
			},
		},
		{
			name: "Subject Only",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				subject := "Hello Corso"
				event := models.NewEvent()
				event.SetSubject(&subject)
				return event, &details.ExchangeInfo{Subject: subject}
			},
		},
		{
			name: "Using mockable",
			evtAndRP: func() (models.Eventable, *details.ExchangeInfo) {
				bytes := mockconnector.GetMockEventBytes("Test Mock")
				event, err := support.CreateEventFromBytes(bytes)
				require.NoError(suite.T(), err)
				subject := " Test MockReview + Lunch"
				sender := "foobar3@8qzvrj.onmicrosoft.com"
				eventTime := time.Date(2022, time.April, 28, 3, 41, 58, 0, time.UTC)
				return event, &details.ExchangeInfo{
					Subject:    subject,
					Sender:     sender,
					EventStart: eventTime,
				}

			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			event, expected := test.evtAndRP()
			result := EventInfo(event)
			suite.Equal(expected.Subject, result.Subject)
			suite.Equal(expected.Sender, result.Sender)
			e_yr, e_mon, e_day := expected.EventStart.Date()
			e_hr, e_min, e_sec := expected.EventStart.Clock()
			r_yr, r_mon, r_day := result.EventStart.Date()
			r_hr, r_min, r_sec := result.EventStart.Clock()
			suite.Equal(e_yr, r_yr)
			suite.Equal(e_mon, r_mon)
			suite.Equal(e_day, r_day)
			suite.Equal(e_hr, r_hr)
			suite.Equal(e_min, r_min)
			suite.Equal(e_sec, r_sec)
		})
	}
}
