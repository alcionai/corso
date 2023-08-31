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

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
	now := dttm.FormatTo(initial, dttm.M365DateTimeTimeZone)

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

				nowp30m := dttm.FormatTo(initial.Add(30*time.Minute), dttm.M365DateTimeTimeZone)
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
					event, err   = api.BytesToEventable(bytes)
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

func (suite *EventsAPIUnitSuite) TestBytesToEventable() {
	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "empty bytes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "invalid bytes",
			byteArray:  []byte("Invalid byte stream \"subject:\" Not going to work"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid Event",
			byteArray:  exchMock.EventBytes("Event Test"),
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := api.BytesToEventable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}

type EventsAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestEventsAPIIntgSuite(t *testing.T) {
	suite.Run(t, &EventsAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *EventsAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *EventsAPIIntgSuite) TestEvents_RestoreLargeAttachment() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	userID := tconfig.M365UserID(suite.T())

	folderName := testdata.DefaultRestoreConfig("eventlargeattachmenttest").Location
	evts := suite.its.ac.Events()
	calendar, err := evts.CreateContainer(ctx, userID, "", folderName)
	require.NoError(t, err, clues.ToCore(err))

	tomorrow := time.Now().Add(24 * time.Hour)
	evt := models.NewEvent()
	sdtz := models.NewDateTimeTimeZone()
	edtz := models.NewDateTimeTimeZone()

	evt.SetSubject(ptr.To("Event with attachment"))
	sdtz.SetDateTime(ptr.To(dttm.Format(tomorrow)))
	sdtz.SetTimeZone(ptr.To("UTC"))
	edtz.SetDateTime(ptr.To(dttm.Format(tomorrow.Add(30 * time.Minute))))
	edtz.SetTimeZone(ptr.To("UTC"))
	evt.SetStart(sdtz)
	evt.SetEnd(edtz)

	item, err := evts.PostItem(ctx, userID, ptr.Val(calendar.GetId()), evt)
	require.NoError(t, err, clues.ToCore(err))

	id, err := evts.PostLargeAttachment(
		ctx,
		userID,
		ptr.Val(calendar.GetId()),
		ptr.Val(item.GetId()),
		"raboganm",
		[]byte("mangobar"),
	)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, id, "empty id for large attachment")
}

func (suite *EventsAPIIntgSuite) TestEvents_canFindNonStandardFolder() {
	t := suite.T()

	t.Skip("currently broken: the test user needs to get rotated")

	ctx, flush := tester.NewContext(t)
	defer flush()

	ac := suite.its.ac.Events()
	rc := testdata.DefaultRestoreConfig("api_calendar_discovery")

	cal, err := ac.CreateContainer(ctx, suite.its.user.id, "", rc.Location)
	require.NoError(t, err, clues.ToCore(err))

	var (
		found         bool
		calID         = ptr.Val(cal.GetId())
		findContainer = func(gcc graph.CachedContainer) error {
			if ptr.Val(gcc.GetId()) == calID {
				found = true
			}

			return nil
		}
	)

	err = ac.EnumerateContainers(
		ctx,
		suite.its.user.id,
		"Calendar",
		findContainer,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	require.True(
		t,
		found,
		"the restored container was discovered when enumerating containers.  "+
			"If this fails, the user's calendars have probably broken, "+
			"and the user will need to be rotated")
}

func (suite *EventsAPIIntgSuite) TestEvents_GetContainerByName() {
	table := []struct {
		name      string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "Calendar",
			expectErr: assert.NoError,
		},
		{
			name:      "smarfs",
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := suite.its.ac.
				Events().
				GetContainerByName(ctx, suite.its.user.id, "", test.name)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}

func (suite *EventsAPIIntgSuite) TestEvents_GetContainerByName_mocked() {
	c := models.NewCalendar()
	c.SetId(ptr.To("id"))
	c.SetName(ptr.To("display name"))

	table := []struct {
		name      string
		results   func(*testing.T) map[string]any
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "zero",
			results: func(t *testing.T) map[string]any {
				return parseableToMap(t, models.NewCalendarCollectionResponse())
			},
			expectErr: assert.Error,
		},
		{
			name: "one",
			results: func(t *testing.T) map[string]any {
				mfcr := models.NewCalendarCollectionResponse()
				mfcr.SetValue([]models.Calendarable{c})

				return parseableToMap(t, mfcr)
			},
			expectErr: assert.NoError,
		},
		{
			name: "two",
			results: func(t *testing.T) map[string]any {
				mfcr := models.NewCalendarCollectionResponse()
				mfcr.SetValue([]models.Calendarable{c, c})

				return parseableToMap(t, mfcr)
			},
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)

			defer flush()
			defer gock.Off()

			interceptV1Path("users", "u", "calendars").
				Reply(200).
				JSON(test.results(t))

			_, err := suite.its.gockAC.
				Events().
				GetContainerByName(ctx, "u", "", test.name)
			test.expectErr(t, err, clues.ToCore(err))
			assert.True(t, gock.IsDone())
		})
	}
}
