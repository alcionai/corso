package ics

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/alcionai/clues"
	ics "github.com/arran4/golang-ical"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"jaytaylor.com/html2text"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/dttm"
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
	iCalDateTimeFormat = "20060102T150405Z"
	iCalDateFormat     = "20060102"
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

func getUTCTime(ts, tz string) (time.Time, error) {
	// Timezone is always converted to UTC.  This is the easiest way to
	// ensure we have the correct time as the .ics file expects the same
	// timezone everywhere according to the spec.
	it, err := dttm.ParseTime(ts)
	if err != nil {
		return time.Time{}, clues.Wrap(err, "parsing time").With("given_time_string", ts)
	}

	timezone, ok := GraphTimeZoneToTZ[tz]
	if !ok {
		return it, clues.New("unknown timezone").With("timezone", tz)
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, clues.Wrap(err, "loading timezone").
			With("converted_timezone", timezone)
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

				endTime, err := getUTCTime(
					parsedTime.Format(string(dttm.M365DateTimeTimeZone)),
					ptr.Val(rrange.GetRecurrenceTimeZone()))
				if err != nil {
					return "", clues.WrapWC(ctx, err, "parsing end time")
				}

				recurComponents = append(recurComponents, "UNTIL="+endTime.Format(iCalDateTimeFormat))
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

	cal := ics.NewCalendar()
	cal.SetProductId("-//Alcion//Corso") // Does this have to be customizable?

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

		exICalEvent.AddProperty(ics.ComponentProperty(ics.PropertyRecurrenceId), start.Format(iCalDateTimeFormat))

		err = updateEventProperties(ctx, exception, exICalEvent)
		if err != nil {
			return "", clues.Wrap(err, "updating exception event properties")
		}
	}

	return cal.Serialize(), nil
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}

	return true
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

	// DTSTART - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.2.4
	allDay := ptr.Val(event.GetIsAllDay())
	startString := event.GetStart().GetDateTime()
	startTimezone := event.GetStart().GetTimeZone()

	if startString != nil {
		start, err := getUTCTime(ptr.Val(startString), ptr.Val(startTimezone))
		if err != nil {
			return clues.WrapWC(ctx, err, "parsing start time")
		}

		if allDay {
			iCalEvent.SetStartAt(start, ics.WithValue(string(ics.ValueDataTypeDate)))
		} else {
			iCalEvent.SetStartAt(start)
		}
	}

	// DTEND - https://www.rfc-editor.org/rfc/rfc5545#section-3.8.2.2
	endString := event.GetEnd().GetDateTime()
	endTimezone := event.GetEnd().GetTimeZone()

	if endString != nil {
		end, err := getUTCTime(ptr.Val(endString), ptr.Val(endTimezone))
		if err != nil {
			return clues.WrapWC(ctx, err, "parsing end time")
		}

		if allDay {
			iCalEvent.SetEndAt(end, ics.WithValue(string(ics.ValueDataTypeDate)))
		} else {
			iCalEvent.SetEndAt(end)
		}
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
	if cancelled != nil {
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
					stripped, err := html2text.FromString(description, html2text.Options{PrettyTables: true})
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

		addr := ptr.Val(attendee.GetEmailAddress().GetAddress())
		iCalEvent.AddAttendee(addr, props...)
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
	// TODO Handle different attachment types (file, item and reference)
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
		dateStrings = append(dateStrings, date.Format(iCalDateFormat))
	}

	if len(dateStrings) > 0 {
		iCalEvent.AddProperty(ics.ComponentPropertyExdate, strings.Join(dateStrings, ","))
	}

	return nil
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
		start, err := getUTCTime(ds, tz)
		if err != nil {
			return nil, clues.WrapWC(ctx, err, "parsing cancelled event date")
		}

		dates = append(dates, start)
	}

	return dates, nil
}
