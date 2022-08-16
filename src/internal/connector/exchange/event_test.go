package exchange

import (
	"testing"
	"time"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EventSuite struct {
	suite.Suite
}

func TestEventSuite(t *testing.T) {
	suite.Run(t, &EventSuite{})
}

func (suite *EventSuite) TestEventInfo() {
	initial := time.Now()
	now := initial.Format(common.StandardTimeFormat)
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
			name: "Received only",
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
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			event, expected := test.evtAndRP()
			suite.Equal(expected, EventInfo(event))
		})
	}
}
