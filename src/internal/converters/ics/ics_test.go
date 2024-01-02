package ics

// Useful tools
// https://icalendar.org/validator.html
// https://icalendar.org/rrule-tool.html
// https://balsoftware.net/rrule/

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
)

type ICSUnitSuite struct {
	tester.Suite
}

func TestICSUnitSuite(t *testing.T) {
	suite.Run(t, &ICSUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ICSUnitSuite) TestGetLocationString() {
	table := []struct {
		name   string
		loc    func() models.Locationable
		expect string
	}{
		{
			name: "only displayname",
			loc: func() models.Locationable {
				loc := models.NewLocation()
				loc.SetDisplayName(ptr.To("DisplayName"))
				return loc
			},
			expect: "DisplayName",
		},
		{
			name: "full address",
			loc: func() models.Locationable {
				loc := models.NewLocation()
				loc.SetDisplayName(ptr.To("DisplayName"))

				addr := models.NewPhysicalAddress()
				addr.SetStreet(ptr.To("Street"))
				addr.SetCity(ptr.To("City"))
				addr.SetState(ptr.To("State"))
				addr.SetCountryOrRegion(ptr.To("Country"))
				addr.SetPostalCode(ptr.To("PostalCode"))

				loc.SetAddress(addr)
				return loc
			},
			expect: "DisplayName, Street, City, State, Country, PostalCode",
		},
		{
			name: "displayname and street",
			loc: func() models.Locationable {
				loc := models.NewLocation()
				loc.SetDisplayName(ptr.To("DisplayName"))

				addr := models.NewPhysicalAddress()
				addr.SetStreet(ptr.To("Street"))

				loc.SetAddress(addr)
				return loc
			},
			expect: "DisplayName, Street",
		},
		{
			name: "only street",
			loc: func() models.Locationable {
				loc := models.NewLocation()

				addr := models.NewPhysicalAddress()
				addr.SetStreet(ptr.To("Street"))

				loc.SetAddress(addr)
				return loc
			},
			expect: "Street",
		},
		{
			name: "displayname, city, country",
			loc: func() models.Locationable {
				loc := models.NewLocation()
				loc.SetDisplayName(ptr.To("DisplayName"))

				addr := models.NewPhysicalAddress()
				addr.SetCity(ptr.To("City"))
				addr.SetCountryOrRegion(ptr.To("Country"))

				loc.SetAddress(addr)
				return loc
			},
			expect: "DisplayName, City, Country",
		},
	}

	for _, tt := range table {
		suite.Run(tt.name, func() {
			assert.Equal(suite.T(), tt.expect, getLocationString(tt.loc()))
		})
	}
}

func (suite *ICSUnitSuite) TestGetUTCTime() {
	table := []struct {
		name      string
		timestamp string
		timezone  string
		time      time.Time
		errCheck  require.ErrorAssertionFunc
	}{
		{
			name:      "valid time in UTC",
			timestamp: "2021-01-01T12:00:00Z",
			timezone:  "UTC",
			time:      time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC),
			errCheck:  require.NoError,
		},
		{
			name:      "valid time in IST",
			timestamp: "2021-01-01T12:00:00Z",
			timezone:  "India Standard Time",
			time:      time.Date(2021, 1, 1, 6, 30, 0, 0, time.UTC),
			errCheck:  require.NoError,
		},
		{
			name:      "invalid time",
			timestamp: "invalid",
			timezone:  "UTC",
			time:      time.Time{},
			errCheck:  require.Error,
		},
		{
			name:      "invalid timezone",
			timestamp: "2021-01-01T12:00:00Z",
			timezone:  "invalid",
			time:      time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC),
			errCheck:  require.Error,
		},
	}

	for _, tt := range table {
		suite.Run(tt.name, func() {
			t, err := getUTCTime(tt.timestamp, tt.timezone)
			tt.errCheck(suite.T(), err)

			if !tt.time.Equal(time.Time{}) {
				assert.Equal(suite.T(), tt.time, t)
			}
		})
	}
}

func (suite *ICSUnitSuite) TestGetRecurrencePattern() {
	table := []struct {
		name       string
		recurrence func() models.PatternedRecurrenceable
		expect     string
		errCheck   require.ErrorAssertionFunc
	}{
		{
			name: "daily",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("daily")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=DAILY;INTERVAL=1",
			errCheck: require.NoError,
		},
		{
			name: "weekly",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("weekly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=WEEKLY;INTERVAL=1",
			errCheck: require.NoError,
		},
		{
			name: "weekly with end date",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("weekly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				rng := models.NewRecurrenceRange()

				rrtype, err := models.ParseRecurrenceRangeType("endDate")
				require.NoError(suite.T(), err)

				rng.SetTypeEscaped(rrtype.(*models.RecurrenceRangeType))

				edate := serialization.NewDateOnly(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC))
				rng.SetEndDate(edate)
				rng.SetRecurrenceTimeZone(ptr.To("UTC"))

				rec.SetPattern(pat)
				rec.SetRangeEscaped(rng)

				return rec
			},
			expect:   "FREQ=WEEKLY;INTERVAL=1;UNTIL=20210101T000000Z",
			errCheck: require.NoError,
		},
		{
			name: "weekly with end count",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("weekly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				rng := models.NewRecurrenceRange()

				rrtype, err := models.ParseRecurrenceRangeType("numbered")
				require.NoError(suite.T(), err)

				rng.SetTypeEscaped(rrtype.(*models.RecurrenceRangeType))

				rng.SetNumberOfOccurrences(ptr.To(int32(10)))
				rng.SetRecurrenceTimeZone(ptr.To("UTC"))

				rec.SetPattern(pat)
				rec.SetRangeEscaped(rng)

				return rec
			},
			expect:   "FREQ=WEEKLY;INTERVAL=1;COUNT=10",
			errCheck: require.NoError,
		},
		{
			name: "days of week",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("weekly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				days := []models.DayOfWeek{
					models.MONDAY_DAYOFWEEK,
					models.WEDNESDAY_DAYOFWEEK,
					models.THURSDAY_DAYOFWEEK,
				}

				pat.SetDaysOfWeek(days)

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,WE,TH",
			errCheck: require.NoError,
		},
		{
			name: "daily with custom interval",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("daily")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(2)))

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=DAILY;INTERVAL=2",
			errCheck: require.NoError,
		},
		{
			name: "day of month",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("absoluteMonthly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				pat.SetDayOfMonth(ptr.To(int32(5)))

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=MONTHLY;INTERVAL=1;BYMONTHDAY=5",
			errCheck: require.NoError,
		},
		{
			name: "every 3rd august",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("absoluteYearly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(3)))

				pat.SetMonth(ptr.To(int32(8)))

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=YEARLY;INTERVAL=3;BYMONTH=8",
			errCheck: require.NoError,
		},
		{
			name: "first friday of august every year",
			recurrence: func() models.PatternedRecurrenceable {
				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("relativeYearly")
				require.NoError(suite.T(), err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				pat.SetMonth(ptr.To(int32(8)))
				pat.SetDaysOfWeek([]models.DayOfWeek{models.FRIDAY_DAYOFWEEK})

				wi, err := models.ParseWeekIndex("first")
				require.NoError(suite.T(), err)
				pat.SetIndex(wi.(*models.WeekIndex))

				rec.SetPattern(pat)

				return rec
			},
			expect:   "FREQ=YEARLY;INTERVAL=1;BYMONTH=8;BYDAY=1FR",
			errCheck: require.NoError,
		},
		// TODO(meain): could still use more tests for edge cases of time
	}

	for _, tt := range table {
		suite.Run(tt.name, func() {
			rec, err := getRecurrencePattern(tt.recurrence())
			tt.errCheck(suite.T(), err)

			assert.Equal(suite.T(), tt.expect, rec)
		})
	}
}

func baseEvent() *models.Event {
	e := models.NewEvent()

	e.SetId(ptr.To("mango"))
	e.SetSubject(ptr.To("Subject"))

	start := models.NewDateTimeTimeZone()
	start.SetDateTime(ptr.To(time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC).Format(time.RFC3339)))
	start.SetTimeZone(ptr.To("UTC"))

	end := models.NewDateTimeTimeZone()
	end.SetDateTime(ptr.To(time.Date(2021, 1, 2, 13, 0, 0, 0, time.UTC).Format(time.RFC3339)))
	end.SetTimeZone(ptr.To("UTC"))

	e.SetStart(start)
	e.SetEnd(end)

	return e
}

func (suite *ICSUnitSuite) TestEventConversion() {
	t := suite.T()

	table := []struct {
		name  string
		event func() *models.Event
		check func(string)
	}{
		{
			name: "simple event",
			event: func() *models.Event {
				return baseEvent()
			},
			check: func(out string) {
				assert.Contains(t, out, "BEGIN:VCALENDAR", "beginning of calendar")
				assert.Contains(t, out, "VERSION:2.0", "version")
				assert.Contains(t, out, "PRODID:-//Alcion//Corso", "prodid")
				assert.Contains(t, out, "BEGIN:VEVENT", "beginning of event")
				assert.Contains(t, out, "UID:mango", "uid")
				assert.Contains(t, out, "SUMMARY:Subject", "summary")
				assert.Contains(t, out, "DTSTART:20210101T120000Z", "start time")
				assert.Contains(t, out, "DTEND:20210102T130000Z", "end time")
				assert.Contains(t, out, "END:VEVENT", "end of event")
				assert.Contains(t, out, "END:VCALENDAR", "end of calendar")
			},
		},
		{
			name: "event with created and modified time",
			event: func() *models.Event {
				e := baseEvent()

				e.SetCreatedDateTime(ptr.To(time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC)))
				e.SetLastModifiedDateTime(ptr.To(time.Date(2021, 1, 2, 13, 0, 0, 0, time.UTC)))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "CREATED:20210101T120000Z", "created time")
				assert.Contains(t, out, "LAST-MODIFIED:20210102T130000Z", "modified time")
			},
		},
		{
			name: "all day events",
			event: func() *models.Event {
				e := baseEvent()

				e.SetIsAllDay(ptr.To(true))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "DTSTART;VALUE=DATE:20210101", "start time")
				assert.Contains(t, out, "DTEND;VALUE=DATE:20210102", "end time")
			},
		},
		{
			// All time values should get converted to UTC
			name: "start and end with different timezone",
			event: func() *models.Event {
				e := baseEvent()

				start := models.NewDateTimeTimeZone()
				start.SetDateTime(ptr.To(time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC).Format(time.RFC3339)))
				start.SetTimeZone(ptr.To("India Standard Time"))

				end := models.NewDateTimeTimeZone()
				end.SetDateTime(ptr.To(time.Date(2021, 1, 2, 13, 0, 0, 0, time.UTC).Format(time.RFC3339)))
				end.SetTimeZone(ptr.To("Korea Standard Time"))

				e.SetStart(start)
				e.SetEnd(end)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "DTSTART:20210101T063000Z", "start time")
				assert.Contains(t, out, "DTEND:20210102T040000Z", "end time")
			},
		},
		{
			name: "daily event",
			event: func() *models.Event {
				e := baseEvent()

				rec := models.NewPatternedRecurrence()
				pat := models.NewRecurrencePattern()

				typ, err := models.ParseRecurrencePatternType("daily")
				require.NoError(t, err)

				pat.SetTypeEscaped(typ.(*models.RecurrencePatternType))
				pat.SetInterval(ptr.To(int32(1)))

				rec.SetPattern(pat)

				e.SetRecurrence(rec)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "RRULE:FREQ=DAILY;INTERVAL=1", "recurrence rule")
			},
		},
		{
			name: "cancelled event",
			event: func() *models.Event {
				e := baseEvent()

				e.SetIsCancelled(ptr.To(true))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "STATUS:CANCELLED", "cancelled status")
			},
		},
		{
			name: "draft event",
			event: func() *models.Event {
				e := baseEvent()

				e.SetIsDraft(ptr.To(true))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "STATUS:DRAFT", "draft status")
			},
		},
		{
			name: "text body",
			event: func() *models.Event {
				e := baseEvent()

				body := models.NewItemBody()
				btype, err := models.ParseBodyType("text")
				require.NoError(t, err, "parse body type")

				body.SetContentType(btype.(*models.BodyType))
				body.SetContent(ptr.To("body"))

				e.SetBody(body)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "DESCRIPTION:body", "body")
			},
		},
		{
			name: "html body",
			event: func() *models.Event {
				e := baseEvent()

				body := models.NewItemBody()
				btype, err := models.ParseBodyType("html")
				require.NoError(t, err, "parse body type")

				body.SetContentType(btype.(*models.BodyType))
				body.SetContent(ptr.To("<html><body>body</body></html>"))

				e.SetBodyPreview(ptr.To("body preview"))
				e.SetBody(body)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "DESCRIPTION:body preview", "body preview")
				assert.Contains(t, out, "X-ALT-DESC;FMTTYPE=text/html:<html><body>body</body></html>", "body")
			},
		},
		{
			name: "free busy free",
			event: func() *models.Event {
				e := baseEvent()

				fbs, err := models.ParseFreeBusyStatus("free")
				require.NoError(t, err, "parse free busy status")

				e.SetShowAs(fbs.(*models.FreeBusyStatus))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "FREEBUSY:FREE", "free busy status")
			},
		},
		{
			name: "free busy unknown",
			event: func() *models.Event {
				e := baseEvent()

				fbs, err := models.ParseFreeBusyStatus("unknown")
				require.NoError(t, err, "parse free busy status")

				e.SetShowAs(fbs.(*models.FreeBusyStatus))

				return e
			},
			check: func(out string) {
				assert.NotContains(t, out, "FREEBUSY", "free busy status")
			},
		},
		{
			name: "free busy oof",
			event: func() *models.Event {
				e := baseEvent()

				fbs, err := models.ParseFreeBusyStatus("oof")
				require.NoError(t, err, "parse free busy status")

				e.SetShowAs(fbs.(*models.FreeBusyStatus))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "FREEBUSY:BUSY-UNAVAILABLE", "free busy status")
			},
		},
		{
			name: "categories",
			event: func() *models.Event {
				e := baseEvent()

				e.SetCategories([]string{"cat1", "cat2"})

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "CATEGORIES:cat1", "categories")
				assert.Contains(t, out, "CATEGORIES:cat2", "categories")
			},
		},
		{
			name: "weblink",
			event: func() *models.Event {
				e := baseEvent()

				e.SetWebLink(ptr.To("https://example.com"))

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "URL:https://example.com", "weblink")
			},
		},
		{
			name: "organizer just email",
			event: func() *models.Event {
				e := baseEvent()

				org := models.NewRecipient()

				addr := models.NewEmailAddress()
				addr.SetAddress(ptr.To("user@provider.co"))
				org.SetEmailAddress(addr)

				e.SetOrganizer(org)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "ORGANIZER:user@provider.co", "organizer")
			},
		},
		{
			name: "organizer name and email",
			event: func() *models.Event {
				e := baseEvent()

				org := models.NewRecipient()

				addr := models.NewEmailAddress()
				addr.SetAddress(ptr.To("user@provider.co"))
				addr.SetName(ptr.To("User"))

				org.SetEmailAddress(addr)

				e.SetOrganizer(org)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "ORGANIZER;CN=User:user@provider.co", "organizer")
			},
		},
		{
			name: "location",
			event: func() *models.Event {
				e := baseEvent()

				// full test is done separately
				loc := models.NewLocation()
				loc.SetDisplayName(ptr.To("DisplayName"))

				e.SetLocation(loc)

				return e
			},
			check: func(out string) {
				assert.Contains(t, out, "LOCATION:DisplayName", "location")
			},
		},
	}

	for _, tt := range table {
		suite.Run(tt.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// convert event to bytes
			writer := kjson.NewJsonSerializationWriter()
			defer writer.Close()

			err := writer.WriteObjectValue("", tt.event())
			require.NoError(t, err, "serializing contact")

			bts, err := writer.GetSerializedContent()
			require.NoError(t, err, "getting serialized content")

			e, err := FromJSON(ctx, bts)
			require.NoError(t, err, "converting to ics")

			tt.check(e)
		})
	}
}

// checkAttendee check the ATTENDEE field
// This is required instead of a string check as the fields might
// always be in the same order
func checkAttendee(t *testing.T, out, check, msg string) {
	splitFunc := func(c rune) bool {
		return c == ';' || c == ':'
	}

	line := ""

	for _, l := range strings.Split(out, "\r\n") {
		if !strings.HasPrefix(l, "ATTENDEE") {
			continue
		}

		splits := strings.Split(check, ":")
		if strings.Contains(l, splits[len(splits)-1]) {
			line = l
			break
		}
	}

	if len(line) == 0 {
		assert.Fail(t, fmt.Sprintf("line not found %s", msg))
		return
	}

	as := strings.FieldsFunc(line, splitFunc)
	bs := strings.FieldsFunc(check, splitFunc)

	assert.Equal(t, len(as), len(bs), fmt.Sprintf("length of fields of %s", msg))
	assert.ElementsMatch(t, as, bs, fmt.Sprintf("fields %s", msg))
}

func (suite *ICSUnitSuite) TestAttendees() {
	t := suite.T()

	table := []struct {
		name  string
		att   [][]string
		check func(string)
	}{
		{
			name: "single attendee",
			// email, role, participation
			att: [][]string{{"one@att.co", "", ""}},
			check: func(out string) {
				checkAttendee(t, out, "ATTENDEE;CN=one:mailto:one@att.co", "attendee")
			},
		},
		{
			name: "single attendee with role and participation",
			att:  [][]string{{"one@att.co", "required", "declined"}},
			check: func(out string) {
				checkAttendee(
					t,
					out,
					"ATTENDEE;ROLE=REQ-PARTICIPANT;CN=one;PARTSTAT=DECLINED:mailto:one@att.co",
					"attendee")
			},
		},
		{
			name: "multiple attendees",
			att: [][]string{
				{"one@att.co", "", ""},
				{"two@att.co", "optional", "accepted"},
				{"th@att.co", "resource", "notResponded"}, // th instead of three to prevent split
				{"four@att.co", "required", "tentativelyAccepted"},
			},
			check: func(out string) {
				checkAttendee(t, out, "ATTENDEE;CN=one:mailto:one@att.co", "attendee one")
				checkAttendee(
					t,
					out,
					"ATTENDEE;ROLE=OPT-PARTICIPANT;CN=two;PARTSTAT=ACCEPTED:mailto:two@att.co",
					"attendee two")
				checkAttendee(
					t,
					out,
					"ATTENDEE;ROLE=NON-PARTICIPANT;CN=th;PARTSTAT=NEEDS-ACTION:mailto:th@att.co",
					"attendee th")
				checkAttendee(
					t,
					out,
					"ATTENDEE;ROLE=REQ-PARTICIPANT;CN=four;PARTSTAT=TENTATIVE:mailto:four@att.co",
					"attendee four")
			},
		},
	}

	for _, tt := range table {
		suite.Run(tt.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// convert event to bytes
			writer := kjson.NewJsonSerializationWriter()
			defer writer.Close()

			e := baseEvent()

			atts := make([]models.Attendeeable, len(tt.att))

			for i, a := range tt.att {
				att := models.NewAttendee()

				addr := models.NewEmailAddress()
				addr.SetAddress(ptr.To(a[0]))

				name := strings.Split(a[0], "@")[0]
				addr.SetName(ptr.To(name))

				att.SetEmailAddress(addr)

				if len(a[1]) > 0 {
					atype, err := models.ParseAttendeeType(a[1])
					require.NoError(t, err, "parse attendee type")

					att.SetTypeEscaped(atype.(*models.AttendeeType))
				}

				if len(a[2]) > 0 {
					stat := models.NewResponseStatus()
					resp, err := models.ParseResponseType(a[2])
					require.NoError(t, err, "parse response type")

					stat.SetResponse(resp.(*models.ResponseType))
					att.SetStatus(stat)
				}

				atts[i] = att
			}

			e.SetAttendees(atts)

			err := writer.WriteObjectValue("", e)
			require.NoError(t, err, "serializing contact")

			bts, err := writer.GetSerializedContent()
			require.NoError(t, err, "getting serialized content")

			out, err := FromJSON(ctx, bts)
			require.NoError(t, err, "converting to ics")

			tt.check(out)
		})
	}
}

func checkAttachment(t *testing.T, out, check, msg string) {
	var (
		attachments       = []string{}
		inAttachment      = false
		attachment        = ""
		checkSplits       = strings.Split(check, ":")[0]
		filenameSegment   = ""
		attachmentToCheck = ""
	)

	for _, l := range strings.Split(out, "\r\n") {
		if strings.HasPrefix(l, "ATTACH") {
			inAttachment = true
			attachment = l
		} else if inAttachment {
			if len(l) > 1 || l[0] == ' ' {
				attachment += l[1:]
			}

			inAttachment = false

			attachments = append(attachments, attachment)
		}
	}

	if inAttachment {
		attachments = append(attachments, attachment)
	}

	if len(attachments) == 0 {
		assert.Fail(t, fmt.Sprintf("no attachments found: %s", msg))
		return
	}

	for _, s := range strings.Split(checkSplits, ";") {
		if strings.HasPrefix(s, "FILENAME") {
			filenameSegment = s
			break
		}
	}

	if len(filenameSegment) == 0 {
		assert.Fail(t, fmt.Sprintf("filename not found %s", msg))
		return
	}

	for _, a := range attachments {
		if strings.Contains(a, filenameSegment) {
			attachmentToCheck = a
			break
		}
	}

	if len(attachmentToCheck) == 0 {
		assert.Fail(t, fmt.Sprintf("attachment not found: %s", msg))
		return
	}

	splitFunc := func(c rune) bool {
		return c == ';' || c == ':'
	}

	as := strings.FieldsFunc(attachmentToCheck, splitFunc)
	bs := strings.FieldsFunc(check, splitFunc)

	assert.Equal(t, len(as), len(bs), fmt.Sprintf("length of fields of %s", msg))
	assert.ElementsMatch(t, as, bs, fmt.Sprintf("fields %s", msg))
}

func (suite *ICSUnitSuite) TestAttachments() {
	t := suite.T()

	type attachment struct {
		cid     string // contentid
		name    string
		ctype   string
		content string
		inline  bool
	}

	table := []struct {
		name  string
		att   []attachment
		check func(string)
	}{
		{
			name: "single attachment",
			att: []attachment{
				{"1", "one", "text/plain", "content", false},
			},
			check: func(out string) {
				checkAttachment(
					t,
					out,
					"ATTACH;VALUE=BINARY;FMTTYPE=text/plain;FILENAME=one;ENCODING=base64:Y29udGVudA==",
					"attachment")
			},
		},
		{
			name: "multiple attachments",
			att: []attachment{
				{"1", "one", "text/plain", "content", false},
				{"2", "two", "text/html", "<html><body>content</body></html>", false},
			},
			check: func(out string) {
				base := "ATTACH;FILENAME=one;ENCODING=base64;VALUE=BINARY;FMTTYPE=text/plain:"
				content := base64.StdEncoding.EncodeToString([]byte("content"))

				checkAttachment(
					t,
					out,
					base+content,
					"attachment one")

				base = "ATTACH;FILENAME=two;ENCODING=base64;VALUE=BINARY;FMTTYPE=text/html:"
				content = base64.StdEncoding.EncodeToString([]byte("<html><body>content</body></html>"))
				checkAttachment(
					t,
					out,
					base+content,
					"attachment two")
			},
		},
	}

	for _, tt := range table {
		suite.Run(tt.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// convert event to bytes
			writer := kjson.NewJsonSerializationWriter()
			defer writer.Close()

			e := baseEvent()

			err := writer.WriteObjectValue("", e)
			require.NoError(t, err, "serializing contact")

			bts, err := writer.GetSerializedContent()
			require.NoError(t, err, "getting serialized content")

			parsed := map[string]any{}
			err = json.Unmarshal(bts, &parsed)
			require.NoError(t, err, "unmarshalling json")

			// could not add attachment content without doing this
			atts := make([]map[string]any, len(tt.att))

			for i, a := range tt.att {
				att := map[string]any{
					"@odata.type":  "#microsoft.graph.fileAttachment",
					"name":         a.name,
					"contentType":  a.ctype,
					"contentBytes": base64.StdEncoding.EncodeToString([]byte(a.content)),
					"contentId":    a.cid,
					"isInline":     a.inline,
				}

				atts[i] = att
			}

			parsed["attachments"] = atts

			bts, err = json.Marshal(parsed)
			require.NoError(t, err, "marshalling json")

			out, err := FromJSON(ctx, bts)
			require.NoError(t, err, "converting to ics")

			tt.check(out)
		})
	}
}
