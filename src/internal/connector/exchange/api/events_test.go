package api_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/exchange/api/mock"
	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type EventsAPIUnitSuite struct {
	tester.Suite
}

func TestEventsAPIUnitSuite(t *testing.T) {
	suite.Run(t, &EventsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// TestEventInfo verifies that searchable event metadata
// can be properly retrieved from a models.Eventable object
func (suite *EventsAPIUnitSuite) TestEventInfo() {
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
					bytes        = exchMock.EventBytes("Test Mock")
					future       = time.Now().UTC().AddDate(0, 0, 1)
					eventTime    = time.Date(future.Year(), future.Month(), future.Day(), future.Hour(), 0, 0, 0, time.UTC)
					eventEndTime = eventTime.Add(30 * time.Minute)
					event, err   = support.CreateEventFromBytes(bytes)
				)

				require.NoError(suite.T(), err, clues.ToCore(err))

				return event, &details.ExchangeInfo{
					ItemType:   details.ExchangeEvent,
					Subject:    subject,
					Organizer:  organizer,
					EventStart: eventTime,
					EventEnd:   eventEndTime,
				}
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			event, expected := test.evtAndRP()
			result := api.EventInfo(event)

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

type EventAPIE2ESuite struct {
	tester.Suite
	credentials account.M365Config
	ac          api.Client
	user        string
}

// We do end up mocking the actual request, but creating the rest
// similar to E2E suite
func TestEventAPIE2ESuite(t *testing.T) {
	suite.Run(t, &EventAPIE2ESuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *EventAPIE2ESuite) SetupSuite() {
	t := suite.T()

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365
	suite.ac, err = mock.NewClient(m365)
	require.NoError(t, err, clues.ToCore(err))

	suite.user = tester.M365UserID(suite.T())
}

func (suite *EventAPIE2ESuite) TestPaginationErrorConditions() {
	did := "directory-id"

	type errResp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	tests := []struct {
		name      string
		prevDelta bool
		setupf    func()
	}{
		{
			name: "direct error on dleta",
			setupf: func() {
				gock.New("https://graph.microsoft.com").
					Get("/beta/users/" + suite.user + "/calendars/" + did + "/events/delta$").
					Reply(404)
			},
		},
		{
			name:      "delta reset",
			prevDelta: true,
			setupf: func() {
				gock.New("https://graph.microsoft.com").
					Get("/fakedelta").
					Reply(403).
					JSON(map[string]errResp{"error": {Code: "SyncStateNotFound", Message: "..."}})

				gock.New("https://graph.microsoft.com").
					Get("/beta/users/" + suite.user + "/calendars/" + did + "/events/delta$").
					Reply(404)
			},
		},
		{
			name:      "box full",
			prevDelta: true,
			setupf: func() {
				gock.New("https://graph.microsoft.com").
					Get("/fakedelta").
					Reply(403).
					JSON(map[string]errResp{
						"error": {
							Code:    "ErrorQuotaExceeded",
							Message: "The process failed to get the correct properties.",
						},
					})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/" + suite.user + "/calendars/" + did + "/events$").
					Reply(404)
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			defer gock.Off()
			tt.setupf()

			delta := ""
			if tt.prevDelta {
				delta = "https://graph.microsoft.com/fakedelta"
			}

			pgr, err := api.NewEventPager(suite.ac.Stable, suite.user, did, delta, false)
			require.NoError(suite.T(), err, "create pager")

			_, _, _, err = api.GetAddedAndRemovedItemIDsFromPager(ctx, delta, &pgr)

			// just a unique enough check
			assert.True(
				suite.T(),
				err.Error() == "The server returned an unexpected status code with no response body: 404",
				"get 404",
			)
			assert.True(suite.T(), gock.IsDone(), "all mocks used")
		})
	}
}
