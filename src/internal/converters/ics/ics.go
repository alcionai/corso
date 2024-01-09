package ics

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	ics "github.com/arran4/golang-ical"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// This package is used to convert json response from graph to ics
// Ref: https://icalendar.org/
// Ref: https://www.rfc-editor.org/rfc/rfc5545
// Ref: https://learn.microsoft.com/en-us/graph/api/resources/event?view=graph-rest-1.0

// TODO: Items not handled
// locations (different from location)
// exceptions and modifications

// Field in the backed up data that we cannot handle
// allowNewTimeProposals, hideAttendees, importance, isOnlineMeeting,
// isOrganizer, isReminderOn, onlineMeeting, onlineMeetingProvider,
// onlineMeetingUrl, originalEndTimeZone, originalStart,
// originalStartTimeZone, reminderMinutesBeforeStart, responseRequested,
// responseStatus, sensitivity

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

	// TODO: Handle different location types
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
		return time.Now(), clues.Wrap(err, "parsing time")
	}

	timezone, ok := GraphTimeZoneToTZ[tz]
	if !ok {
		return it, clues.New("unknown timezone")
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now(), clues.Wrap(err, "loading timezone")
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
				// NOTE: We convert just a date into date+time in a
				// different timezone which will cause it to not be just
				// a date anymore.
				endTime, err := getUTCTime(end.String(), ptr.Val(rrange.GetRecurrenceTimeZone()))
				if err != nil {
					return "", clues.Wrap(err, "parsing end time")
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
		return "", clues.Wrap(err, "converting to eventable")
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
	if exceptionOcurrances != nil {
		for _, occ := range exceptionOcurrances.([]any) {
			instance, ok := occ.(map[string]any)
			if !ok {
				return "", clues.New("converting exception instance to map[string]any").
					With("type", fmt.Sprintf("%T", instance))
			}

			exBody, err := json.Marshal(instance)
			if err != nil {
				return "", clues.Wrap(err, "marshalling exception instance")
			}

			exception, err := api.BytesToEventable(exBody)
			if err != nil {
				return "", clues.Wrap(err, "converting to eventable")
			}

			exICalEvent := cal.AddEvent(id)
			start := exception.GetOriginalStart() // will always be in UTC

			exICalEvent.AddProperty(ics.ComponentProperty(ics.PropertyRecurrenceId), start.Format(iCalDateTimeFormat))

			err = updateEventProperties(ctx, exception, exICalEvent)
			if err != nil {
				return "", clues.WrapWC(ctx, err, "updating exception event properties")
			}
		}
	}

	return cal.Serialize(), nil
}

func updateEventProperties(ctx context.Context, event models.Eventable, iCalEvent *ics.VEvent) error {
	created := event.GetCreatedDateTime()
	if created != nil {
		iCalEvent.SetCreatedTime(ptr.Val(created))
	}

	modified := event.GetLastModifiedDateTime()
	if modified != nil {
		iCalEvent.SetModifiedAt(ptr.Val(modified))
	}

	allDay := ptr.Val(event.GetIsAllDay())

	startString := event.GetStart().GetDateTime()
	startTimezone := event.GetStart().GetTimeZone()

	if startString != nil {
		start, err := getUTCTime(ptr.Val(startString), ptr.Val(startTimezone))
		if err != nil {
			return clues.Wrap(err, "parsing start time")
		}

		if allDay {
			iCalEvent.SetStartAt(start, ics.WithValue(string(ics.ValueDataTypeDate)))
		} else {
			iCalEvent.SetStartAt(start)
		}
	}

	endString := event.GetEnd().GetDateTime()
	endTimezone := event.GetEnd().GetTimeZone()

	if endString != nil {
		end, err := getUTCTime(ptr.Val(endString), ptr.Val(endTimezone))
		if err != nil {
			return clues.Wrap(err, "parsing end time")
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
			return clues.WrapWC(ctx, err, "generating RRULE")
		}

		iCalEvent.AddRrule(pattern)
	}

	cancelled := event.GetIsCancelled()
	if cancelled != nil {
		iCalEvent.SetStatus(ics.ObjectStatusCancelled)
	}

	draft := event.GetIsDraft()
	if draft != nil {
		iCalEvent.SetStatus(ics.ObjectStatusDraft)
	}

	summary := event.GetSubject()
	if summary != nil {
		iCalEvent.SetSummary(ptr.Val(summary))
	}

	// TODO: Emojies currently don't seem to be read properly by Outlook
	bodyPreview := ptr.Val(event.GetBodyPreview())

	if event.GetBody() != nil {
		description := ptr.Val(event.GetBody().GetContent())
		contentType := event.GetBody().GetContentType().String()

		if len(description) > 0 && contentType == "text" {
			iCalEvent.SetDescription(description)
		} else if len(description) > 0 {
			// https://stackoverflow.com/a/859475
			iCalEvent.SetDescription(bodyPreview)

			if contentType == "html" {
				desc := strings.ReplaceAll(description, "\r\n", "")
				desc = strings.ReplaceAll(desc, "\n", "")
				iCalEvent.AddProperty("X-ALT-DESC", desc, ics.WithFmtType("text/html"))
			}
		}
	}

	showAs := ptr.Val(event.GetShowAs()).String()
	if len(showAs) > 0 && showAs != "unknown" {
		var status ics.FreeBusyTimeType

		switch showAs {
		case "free":
			status = ics.FreeBusyTimeTypeFree
		case "tentative":
			status = ics.FreeBusyTimeTypeBusyTentative
		case "busy":
			status = ics.FreeBusyTimeTypeBusy
		case "oof", "workingElsewhere": // this is just best effort conversion
			status = ics.FreeBusyTimeTypeBusyUnavailable
		}

		iCalEvent.AddProperty(ics.ComponentPropertyFreebusy, string(status))
	}

	categories := event.GetCategories()
	for _, category := range categories {
		iCalEvent.AddProperty(ics.ComponentPropertyCategories, category)
	}

	// According to the RFC, this property may be used in a calendar
	// component to convey a location where a more dynamic rendition of
	// the calendar information associated with the calendar component
	// can be found.
	url := ptr.Val(event.GetWebLink())
	if len(url) > 0 {
		iCalEvent.SetURL(url)
	}

	organizer := event.GetOrganizer()
	if organizer != nil {
		name := ptr.Val(organizer.GetEmailAddress().GetName())
		addr := ptr.Val(organizer.GetEmailAddress().GetAddress())

		// TODO: What to do if we only have a name?
		if len(name) > 0 && len(addr) > 0 {
			iCalEvent.SetOrganizer(addr, ics.WithCN(name))
		} else if len(addr) > 0 {
			iCalEvent.SetOrganizer(addr)
		}
	}

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

	location := getLocationString(event.GetLocation())
	if len(location) > 0 {
		iCalEvent.SetLocation(location)
	}

	// TODO Handle different attachment type (file, item and reference)
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
			return clues.Wrap(err, "getting attachment content")
		}

		content, ok := cb.([]uint8)
		if !ok {
			return clues.NewWC(ctx, "getting attachment content string")
		}

		props = append(props, ics.WithEncoding("base64"), ics.WithValue("BINARY"))
		if len(contentType) > 0 {
			props = append(props, ics.WithFmtType(contentType))
		}

		// TODO: Inline attachments don't show up in Outlook
		inline := ptr.Val(attachment.GetIsInline())
		if inline {
			cidv, err := attachment.GetBackingStore().Get("contentId")
			if err != nil {
				return clues.Wrap(err, "getting attachment content id")
			}

			cid, err := str.AnyToString(cidv)
			if err != nil {
				return clues.Wrap(err, "getting attachment content id string")
			}

			props = append(props, keyValues("CID", cid))
		}

		iCalEvent.AddAttachment(base64.StdEncoding.EncodeToString(content), props...)
	}

	cancelldedOcurrances := event.GetAdditionalData()["cancelledOccurrences"]
	if cancelldedOcurrances != nil {
		for _, occ := range cancelldedOcurrances.([]any) {
			instance, err := str.AnyToString(occ)
			if err != nil {
				return clues.Wrap(err, "getting cancelled occurrence id")
			}

			splits := strings.Split(instance, ".")
			if len(splits) < 2 {
				return clues.NewWC(ctx, "invalid cancelled occurrence id").With("id", instance)
			}

			startStr := splits[len(splits)-1]

			// the data just contains date and no time which seems to work
			start, err := getUTCTime(startStr, ptr.Val(startTimezone))
			if err != nil {
				return clues.Wrap(err, "parsing cancelled event date")
			}

			iCalEvent.AddExdate(start.Format(iCalDateTimeFormat))
		}
	}

	return nil
}
