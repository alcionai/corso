package ics

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/mail"
	"strings"
	"time"
	"unicode"

	"github.com/alcionai/clues"
	ics "github.com/arran4/golang-ical"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"jaytaylor.com/html2text"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/converters/ics/tzdata"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// This package is used to convert json response from graph to ics
// https://icalendar.org/
// https://www.rfc-editor.org/rfc/rfc5545
// https://www.rfc-editor.org/rfc/rfc2445.txt
// https://learn.microsoft.com/en-us/graph/api/resources/event?view=graph-rest-1.0
// https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxcical/a685a040-5b69-4c84-b084-795113fb4012

// TODO locations: https://github.com/alcionai/corso/issues/5003

const (
	ICalDateTimeFormat    = "20060102T150405"
	ICalDateTimeFormatUTC = "20060102T150405Z"
	ICalDateFormat        = "20060102"
)

func keyValues(key, value string) *ics.KeyValues {
	return &ics.KeyValues{
		Key:   key,
		Value: []string{value},
	}
}

func getLocationString(location models.Locationable) string {
	if location == nil {
		return ""
	}

	dn := ptr.Val(location.GetDisplayName())
	segments := []string{dn}

	addr := location.GetAddress()
	if addr != nil {
		street := ptr.Val(addr.GetStreet())
		city := ptr.Val(addr.GetCity())
		state := ptr.Val(addr.GetState())
		country := ptr.Val(addr.GetCountryOrRegion())
		postal := ptr.Val(addr.GetPostalCode())

		segments = append(segments, street, city, state, country, postal)
	}

	nonEmpty := []string{}

	for _, seg := range segments {
		if len(seg) > 0 {
			nonEmpty = append(nonEmpty, seg)
		}
	}

	return strings.Join(nonEmpty, ", ")
}

func GetUTCTime(ts, tz string) (time.Time, error) {
	var (
		loc *time.Location
		err error
	)

	// Timezone is always converted to UTC.  This is the easiest way to
	// ensure we have the correct time as the .ics file expects the same
	// timezone everywhere according to the spec.
	it, err := dttm.ParseTime(ts)
	if err != nil {
		return time.Time{}, clues.Wrap(err, "parsing time").With("given_time_string", ts)
	}

	loc, err = time.LoadLocation(tz)
	if err != nil {
		timezone, ok := GraphTimeZoneToTZ[tz]
		if !ok {
			return it, clues.New("unknown timezone").With("timezone", tz)
		}

		loc, err = time.LoadLocation(timezone)
		if err != nil {
			return time.Time{}, clues.Wrap(err, "loading timezone").
				With("converted_timezone", timezone)
		}
	}

	// embed timezone
	locTime := time.Date(it.Year(), it.Month(), it.Day(), it.Hour(), it.Minute(), it.Second(), 0, loc)

	return locTime.UTC(), nil
}

// https://www.rfc-editor.org/rfc/rfc5545#section-3.8.5.3
// https://www.rfc-editor.org/rfc/rfc5545#section-3.3.10
// https://learn.microsoft.com/en-us/graph/api/resources/patternedrecurrence?view=graph-rest-1.0
// Ref: https://github.com/closeio/sync-engine/pull/381/files
func getRecurrencePattern(
	ctx context.Context,
	recurrence models.PatternedRecurrenceable,
) (string, error) {
	recurComponents := []string{}
	pat := recurrence.GetPattern()

	freq := pat.GetTypeEscaped()
	if freq != nil {
		switch *freq {
		case models.DAILY_RECURRENCEPATTERNTYPE:
			recurComponents = append(recurComponents, "FREQ=DAILY")
		case models.WEEKLY_RECURRENCEPATTERNTYPE:
			recurComponents = append(recurComponents, "FREQ=WEEKLY")
		case models.ABSOLUTEMONTHLY_RECURRENCEPATTERNTYPE, models.RELATIVEMONTHLY_RECURRENCEPATTERNTYPE:
			recurComponents = append(recurComponents, "FREQ=MONTHLY")
		case models.ABSOLUTEYEARLY_RECURRENCEPATTERNTYPE, models.RELATIVEYEARLY_RECURRENCEPATTERNTYPE:
			recurComponents = append(recurComponents, "FREQ=YEARLY")
		}
	}

	interval := pat.GetInterval()
	if interval != nil {
		recurComponents = append(recurComponents, "INTERVAL="+fmt.Sprint(ptr.Val(interval)))
	}

	month := ptr.Val(pat.GetMonth())
	if month > 0 {
		recurComponents = append(recurComponents, "BYMONTH="+fmt.Sprint(month))
	}

	// This is required if absoluteMonthly or absoluteYearly
	day := ptr.Val(pat.GetDayOfMonth())
	if day > 0 {
		recurComponents = append(recurComponents, "BYMONTHDAY="+fmt.Sprint(day))
	}

	dow := pat.GetDaysOfWeek()
	if dow != nil {
		dowComponents := []string{}

		for _, day := range dow {
			icalday, ok := GraphToICalDOW[day.String()]
			if !ok {
				return "", clues.NewWC(ctx, "unknown day of week").With("day", day.String())
			}

			dowComponents = append(dowComponents, icalday)
		}

		index := pat.GetIndex()
		prefix := ""

		if index != nil &&
			(ptr.Val(freq) == models.RELATIVEMONTHLY_RECURRENCEPATTERNTYPE ||
				ptr.Val(freq) == models.RELATIVEYEARLY_RECURRENCEPATTERNTYPE) {
			prefix = fmt.Sprint(GraphToICalIndex[index.String()])
		}

		recurComponents = append(recurComponents, "BYDAY="+prefix+strings.Join(dowComponents, ","))
	}

	// This is necessary to compute when weekly events recur
	fdow := pat.GetFirstDayOfWeek()
	if fdow != nil {
		icalday, ok := GraphToICalDOW[fdow.String()]
		if !ok {
			return "", clues.NewWC(ctx, "unknown first day of week").With("day", fdow)
		}

		recurComponents = append(recurComponents, "WKST="+icalday)
	}

	rrange := recurrence.GetRangeEscaped()
	if rrange != nil {
		switch ptr.Val(rrange.GetTypeEscaped()) {
		case models.ENDDATE_RECURRENCERANGETYPE:
			end := rrange.GetEndDate()
			if end != nil {
				parsedTime, err := dttm.ParseTime(end.String())
				if err != nil {
					return "", clues.Wrap(err, "parsing recurrence end date").With("recur_end_date", end.String())
				}

				// end date is always computed as end of the day and
				// so add 23 hours 59 minutes 59 seconds as seconds is
				// the resolution we need
				parsedTime = parsedTime.Add(24*time.Hour - 1*time.Second)

				endTime, err := GetUTCTime(
					parsedTime.Format(string(dttm.M365DateTimeTimeZone)),
					ptr.Val(rrange.GetRecurrenceTimeZone()))
				if err != nil {
					return "", clues.WrapWC(ctx, err, "parsing end time")
				}

				recurComponents = append(recurComponents, "UNTIL="+endTime.Format(ICalDateTimeFormatUTC))
			}
		case models.NOEND_RECURRENCERANGETYPE:
			// Nothing to do
		case models.NUMBERED_RECURRENCERANGETYPE:
			count := ptr.Val(rrange.GetNumberOfOccurrences())
			if count > 0 {
				recurComponents = append(recurComponents, "COUNT="+fmt.Sprint(count))
			}
		}
	}

	return strings.Join(recurComponents, ";"), nil
}

func FromJSON(ctx context.Context, body []byte) (string, error) {
	event, err := api.BytesToEventable(body)
	if err != nil {
		return "", clues.WrapWC(ctx, err, "converting to eventable").
			With("body_len", len(body))
	}

	return FromEventable(ctx, event)
}

func FromEventable(ctx context.Context, event models.Eventable) (string, error) {
	cal := ics.NewCalendar()
	cal.SetProductId("-//Alcion//Corso") // Does this have to be customizable?

	err := addTimeZoneComponents(ctx, cal, event)
	if err != nil {
		return "", clues.Wrap(err, "adding timezone components")
	}

	id := ptr.Val(event.GetId())
	iCalEvent := cal.AddEvent(id)

	err = updateEventProperties(ctx, event, iCalEvent)
	if err != nil {
		return "", clues.Wrap(err, "updating event properties")
	}

	exceptionOcurrances := event.GetAdditionalData()["exceptionOccurrences"]
	if exceptionOcurrances == nil {
		return cal.Serialize(), nil
	}

	for _, occ := range exceptionOcurrances.([]any) {
		instance, ok := occ.(map[string]any)
		if !ok {
			return "", clues.NewWC(ctx, "converting exception instance to map[string]any").
				With("interface_type", fmt.Sprintf("%T", instance))
		}

		exBody, err := json.Marshal(instance)
		if err != nil {
			return "", clues.WrapWC(ctx, err, "marshalling exception instance").
				With("instance_id", instance["id"])
		}

		exception, err := api.BytesToEventable(exBody)
		if err != nil {
			return "", clues.WrapWC(ctx, err, "converting to eventable")
		}

		exICalEvent := cal.AddEvent(id)
		start := exception.GetOriginalStart() // will always be in UTC

		exICalEvent.AddProperty(ics.ComponentProperty(ics.PropertyRecurrenceId), start.Format(ICalDateTimeFormatUTC))

		err = updateEventProperties(ctx, exception, exICalEvent)
		if err != nil {
			return "", clues.Wrap(err, "updating exception event properties")
		}
	}

	return cal.Serialize(), nil
}

func getTZDataKeyValues(ctx context.Context, timezone string) (map[string]string, error) {
	template, ok := tzdata.TZData[timezone]
	if !ok {
		return nil, clues.NewWC(ctx, "timezone not found in tz database").
			With("timezone", timezone)
	}

	keyValues := map[string]string{}

	for _, line := range strings.Split(template, "\n") {
		splits := strings.SplitN(line, ":", 2)
		if len(splits) != 2 {
			return nil, clues.NewWC(ctx, "invalid tzdata line").
				With("line", line).
				With("timezone", timezone)
		}

		keyValues[splits[0]] = splits[1]
	}

	return keyValues, nil
}

func addTimeZoneComponents(ctx context.Context, cal *ics.Calendar, event models.Eventable) error {
	// Handling of timezone get a bit tricky when we have to deal with
	// relative recurrence. The issue comes up when we set a recurrence
	// to be something like "repeat every 3rd Tuesday". Tuesday in UTC
	// and in IST will be different and so we cannot just always use UTC.
	//
	// The way this is solved is by using the timezone in the
	// recurrence for start and end timezones as we have to use UTC
	// for UNTIL(mostly).
	// https://www.rfc-editor.org/rfc/rfc5545#section-3.3.10
	timezone, err := getRecurrenceTimezone(ctx, event)
	if err != nil {
		return clues.Stack(err)
	}

	if timezone != time.UTC {
		kvs, err := getTZDataKeyValues(ctx, timezone.String())
		if err != nil {
			return clues.Stack(err)
		}

		tz := cal.AddTimezone(timezone.String())

		for k, v := range kvs {
			tz.AddProperty(ics.ComponentProperty(k), v)
		}
	}

	return nil
}

// getRecurrenceTimezone get the timezone specified by the recurrence
// in the calendar.  It does a normalization pass where we always convert
// the timezone to the value in tzdb If we don't have a recurrence
// timezone, we don't have to use a specific timezone in the export and
// is safe to return UTC from this method.
func getRecurrenceTimezone(ctx context.Context, event models.Eventable) (*time.Location, error) {
	if event.GetRecurrence() != nil {
		timezone := ptr.Val(event.GetRecurrence().GetRangeEscaped().GetRecurrenceTimeZone())

		ctz, ok := GraphTimeZoneToTZ[timezone]
		if ok {
			timezone = ctz
		}

		cannon, ok := CanonicalTimeZoneMap[timezone]
		if ok {
			timezone = cannon
		}

		loc, err := time.LoadLocation(timezone)
		if err != nil {
			return nil, clues.WrapWC(ctx, err, "unknown timezone").
				With("timezone", timezone)
		}

		return loc, nil
	}

	return time.UTC, nil
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}

	return true
}

// Checks if a given string is a valid email address
func isEmail(em string) bool {
	_, err := mail.ParseAddress(em)
	return err == nil
}

func updateEventProperties(ctx context.Context, event models.Eventable, iCalEvent *ics.VEvent) error {
	// CREATED - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.7.1
	created := event.GetCreatedDateTime()
	if created != nil {
		iCalEvent.SetCreatedTime(ptr.Val(created))
	}

	// LAST-MODIFIED - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.7.3
	modified := event.GetLastModifiedDateTime()
	if modified != nil {
		iCalEvent.SetModifiedAt(ptr.Val(modified))
	}

	timezone, err := getRecurrenceTimezone(ctx, event)
	if err != nil {
		return err
	}

	// DTSTART - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.2.4
	allDay := ptr.Val(event.GetIsAllDay())
	startString := event.GetStart().GetDateTime()
	startTimezone := event.GetStart().GetTimeZone()

	if startString != nil {
		start, err := GetUTCTime(ptr.Val(startString), ptr.Val(startTimezone))
		if err != nil {
			return clues.WrapWC(ctx, err, "parsing start time")
		}

		addTime(iCalEvent, ics.ComponentPropertyDtStart, start, allDay, timezone)
	}

	// DTEND - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.2.2
	endString := event.GetEnd().GetDateTime()
	endTimezone := event.GetEnd().GetTimeZone()

	if endString != nil {
		end, err := GetUTCTime(ptr.Val(endString), ptr.Val(endTimezone))
		if err != nil {
			return clues.WrapWC(ctx, err, "parsing end time")
		}

		addTime(iCalEvent, ics.ComponentPropertyDtEnd, end, allDay, timezone)
	}

	recurrence := event.GetRecurrence()
	if recurrence != nil {
		pattern, err := getRecurrencePattern(ctx, recurrence)
		if err != nil {
			return clues.Wrap(err, "generating RRULE")
		}

		iCalEvent.AddRrule(pattern)
	}

	// STATUS - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.11
	cancelled := event.GetIsCancelled()
	if cancelled != nil && ptr.Val(cancelled) {
		iCalEvent.SetStatus(ics.ObjectStatusCancelled)
	}

	// SUMMARY - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.12
	summary := event.GetSubject()
	if summary != nil {
		iCalEvent.SetSummary(ptr.Val(summary))
	}

	// DESCRIPTION - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.5
	if event.GetBody() != nil {
		description := ptr.Val(event.GetBody().GetContent())
		contentType := event.GetBody().GetContentType().String()

		if len(description) > 0 && contentType == "text" {
			iCalEvent.SetDescription(description)
		} else if len(description) > 0 {
			if contentType == "html" {
				// If we have html, we have two routes. If we don't have
				// UTF-8, then we can do an exact reproduction of the
				// original data in outlook by using X-ALT-DESC field and
				// using the html there. But if we have UTF-8, then we
				// have to use DESCRIPTION field and use the content
				// stripped of html there. This because even though the
				// field technically supports UTF-8, Outlook does not
				// seem to work with it.  Exchange does similar things
				// when it attaches the event to an email.

				// nolint:lll
				// https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxcical/d7f285da-9c7a-4597-803b-b74193c898a8
				// X-ALT-DESC field uses "Text" as in https://www.rfc-editor.org/rfc/rfc2445#section-4.3.11
				if isASCII(description) {
					// https://stackoverflow.com/a/859475
					replacer := strings.NewReplacer("\r\n", "\\n", "\n", "\\n")
					desc := replacer.Replace(description)
					iCalEvent.AddProperty("X-ALT-DESC", desc, ics.WithFmtType("text/html"))
				} else {
					// Disable auto wrap, causes huge memory spikes
					// https://github.com/jaytaylor/html2text/issues/48
					prettyTablesOptions := html2text.NewPrettyTablesOptions()
					prettyTablesOptions.AutoWrapText = false

					stripped, err := html2text.FromString(
						description,
						html2text.Options{PrettyTables: true, PrettyTablesOptions: prettyTablesOptions})
					if err != nil {
						return clues.Wrap(err, "converting html to text").
							With("description_length", len(description))
					}

					iCalEvent.SetDescription(stripped)
				}
			}
		}
	}

	// TRANSP - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.2.7
	showAs := ptr.Val(event.GetShowAs()).String()
	if len(showAs) > 0 {
		var transp ics.TimeTransparency

		switch showAs {
		case "free", "unknown":
			transp = ics.TransparencyTransparent
		default:
			transp = ics.TransparencyOpaque
		}

		iCalEvent.AddProperty(ics.ComponentPropertyTransp, string(transp))
	}

	// CATEGORIES - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.2
	categories := event.GetCategories()
	for _, category := range categories {
		iCalEvent.AddProperty(ics.ComponentPropertyCategories, category)
	}

	// URL - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.4.6
	// According to the RFC, this property may be used in a calendar
	// component to convey a location where a more dynamic rendition of
	// the calendar information associated with the calendar component
	// can be found.
	url := ptr.Val(event.GetWebLink())
	if len(url) > 0 {
		iCalEvent.SetURL(url)
	}

	// ORGANIZER - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.4.3
	organizer := event.GetOrganizer()
	if organizer != nil {
		name := ptr.Val(organizer.GetEmailAddress().GetName())
		addr := ptr.Val(organizer.GetEmailAddress().GetAddress())

		// It does not look like we can get just a name without an address
		if len(name) > 0 && len(addr) > 0 {
			iCalEvent.SetOrganizer(addr, ics.WithCN(name))
		} else if len(addr) > 0 {
			iCalEvent.SetOrganizer(addr)
		}
	}

	// ATTENDEE - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.4.1
	attendees := event.GetAttendees()
	for _, attendee := range attendees {
		props := []ics.PropertyParameter{}

		atype := attendee.GetTypeEscaped()
		if atype != nil {
			var role ics.ParticipationRole

			switch atype.String() {
			case "required":
				role = ics.ParticipationRoleReqParticipant
			case "optional":
				role = ics.ParticipationRoleOptParticipant
			case "resource":
				role = ics.ParticipationRoleNonParticipant
			}

			props = append(props, keyValues(string(ics.ParameterRole), string(role)))
		}

		name := ptr.Val(attendee.GetEmailAddress().GetName())
		if len(name) > 0 {
			props = append(props, ics.WithCN(name))
		}

		// Time when a resp change occurred is not recorded
		if attendee.GetStatus() != nil {
			resp := ptr.Val(attendee.GetStatus().GetResponse()).String()
			if len(resp) > 0 && resp != "none" {
				var pstat ics.ParticipationStatus

				switch resp {
				case "accepted", "organizer":
					pstat = ics.ParticipationStatusAccepted
				case "declined":
					pstat = ics.ParticipationStatusDeclined
				case "tentativelyAccepted":
					pstat = ics.ParticipationStatusTentative
				case "notResponded":
					pstat = ics.ParticipationStatusNeedsAction
				}

				props = append(props, keyValues(string(ics.ParameterParticipationStatus), string(pstat)))
			}
		}

		// It is possible that we get non email items like the below
		// one which is an internal representation of the user in the
		// Exchange system. While we can technically output this as an
		// attendee, it is not useful plus other downstream tools like
		// ones to use PST can choke on this.
		// /o=ExchangeLabs/ou=ExchangeAdministrative Group(FY...LT)/cn=Recipients/cn=883...4a-John Doe
		addr := ptr.Val(attendee.GetEmailAddress().GetAddress())
		if isEmail(addr) {
			iCalEvent.AddAttendee(addr, props...)
		} else {
			logger.Ctx(ctx).
				With("attendee_email", addr).
				With("attendee_name", name).
				Info("skipping non email attendee from ics export")
		}
	}

	// LOCATION - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.7
	location := getLocationString(event.GetLocation())
	if len(location) > 0 {
		iCalEvent.SetLocation(location)
	}

	// X-MICROSOFT-LOCATIONDISPLAYNAME (Outlook seems to use this)
	loc := event.GetLocation()
	if loc != nil {
		locationDisplayName := ptr.Val(event.GetLocation().GetDisplayName())
		if len(locationDisplayName) > 0 {
			iCalEvent.AddProperty("X-MICROSOFT-LOCATIONDISPLAYNAME", locationDisplayName)
		}
	}

	// CLASS - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.3
	// Graph also has the value "personal" which is not supported by the spec
	// Default value is "public" (works for "normal")
	sensitivity := ptr.Val(event.GetSensitivity()).String()
	if sensitivity == "private" {
		iCalEvent.AddProperty(ics.ComponentPropertyClass, "PRIVATE")
	} else if sensitivity == "confidential" {
		iCalEvent.AddProperty(ics.ComponentPropertyClass, "CONFIDENTIAL")
	}

	// PRIORITY - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.9
	imp := ptr.Val(event.GetImportance()).String()
	switch imp {
	case "high":
		iCalEvent.AddProperty(ics.ComponentPropertyPriority, "1")
	case "low":
		iCalEvent.AddProperty(ics.ComponentPropertyPriority, "9")
	}

	meeting := event.GetOnlineMeeting()
	if meeting != nil {
		url := ptr.Val(meeting.GetJoinUrl())
		if len(url) > 0 {
			iCalEvent.AddProperty("X-MICROSOFT-SKYPETEAMSMEETINGURL", url)
		}
	}

	// ATTACH - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.1.1
	attachments := event.GetAttachments()
	for _, attachment := range attachments {
		props := []ics.PropertyParameter{}
		contentType := ptr.Val(attachment.GetContentType())
		name := ptr.Val(attachment.GetName())

		if len(name) > 0 {
			// FILENAME does not seem to be parsed by Outlook
			props = append(props,
				&ics.KeyValues{
					Key:   "FILENAME",
					Value: []string{name},
				})
		}

		cb, err := attachment.GetBackingStore().Get("contentBytes")
		if err != nil {
			return clues.WrapWC(ctx, err, "getting attachment content")
		}

		if cb == nil {
			// TODO(meain): Handle non file attachments
			// https://github.com/alcionai/corso/issues/4772
			logger.Ctx(ctx).
				With("attachment_id", ptr.Val(attachment.GetId()),
					"attachment_type", ptr.Val(attachment.GetOdataType())).
				Info("no contentBytes for attachment")

			continue
		}

		content, ok := cb.([]uint8)
		if !ok {
			return clues.NewWC(ctx, "getting attachment content string").
				With("interface_type", fmt.Sprintf("%T", cb))
		}

		props = append(props, ics.WithEncoding("base64"), ics.WithValue("BINARY"))
		if len(contentType) > 0 {
			props = append(props, ics.WithFmtType(contentType))
		}

		// TODO: Inline attachments don't show up in Outlook
		// Inline attachments is not something supported by the spec
		inline := ptr.Val(attachment.GetIsInline())
		if inline {
			cidv, err := attachment.GetBackingStore().Get("contentId")
			if err != nil {
				return clues.WrapWC(ctx, err, "getting attachment content id")
			}

			cid, err := str.AnyToString(cidv)
			if err != nil {
				return clues.WrapWC(ctx, err, "getting attachment content id string").
					With("interface_type", fmt.Sprintf("%T", cidv))
			}

			props = append(props, keyValues("CID", cid))
		}

		iCalEvent.AddAttachment(base64.StdEncoding.EncodeToString(content), props...)
	}

	// EXDATE - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.5.1
	cancelledDates, err := getCancelledDates(ctx, event)
	if err != nil {
		return clues.Wrap(err, "getting cancelled dates").
			With("event_id", event.GetId())
	}

	dateStrings := []string{}
	for _, date := range cancelledDates {
		dateStrings = append(dateStrings, date.Format(ICalDateFormat))
	}

	if len(dateStrings) > 0 {
		iCalEvent.AddProperty(ics.ComponentPropertyExdate, strings.Join(dateStrings, ","))
	}

	return nil
}

func addTime(iCalEvent *ics.VEvent, prop ics.ComponentProperty, tm time.Time, allDay bool, tzLoc *time.Location) {
	if allDay {
		if tzLoc == time.UTC {
			iCalEvent.SetProperty(prop, tm.Format(ICalDateFormat), ics.WithValue(string(ics.ValueDataTypeDate)))
		} else {
			iCalEvent.SetProperty(
				prop,
				tm.In(tzLoc).Format(ICalDateFormat),
				ics.WithValue(string(ics.ValueDataTypeDate)),
				keyValues("TZID", tzLoc.String()))
		}
	} else {
		if tzLoc == time.UTC {
			iCalEvent.SetProperty(prop, tm.Format(ICalDateTimeFormatUTC))
		} else {
			iCalEvent.SetProperty(prop, tm.In(tzLoc).Format(ICalDateTimeFormat), keyValues("TZID", tzLoc.String()))
		}
	}
}

func getCancelledDates(ctx context.Context, event models.Eventable) ([]time.Time, error) {
	dateStrings, err := api.GetCancelledEventDateStrings(event)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "getting cancelled event date strings")
	}

	dates := []time.Time{}
	tz := ptr.Val(event.GetStart().GetTimeZone())

	for _, ds := range dateStrings {
		// the data just contains date and no time which seems to work
		start, err := GetUTCTime(ds, tz)
		if err != nil {
			return nil, clues.WrapWC(ctx, err, "parsing cancelled event date")
		}

		dates = append(dates, start)
	}

	return dates, nil
}
