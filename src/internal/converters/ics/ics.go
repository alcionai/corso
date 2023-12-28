package ics

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	ics "github.com/arran4/golang-ical"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

func keyValues(key, value string) *ics.KeyValues {
	return &ics.KeyValues{
		Key:   key,
		Value: []string{value},
	}
}

func getLocationString(location models.Locationable) string {
	loc := ""

	dn := ptr.Val(location.GetDisplayName())
	addr := location.GetAddress()
	street := ptr.Val(addr.GetStreet())
	city := ptr.Val(addr.GetCity())
	state := ptr.Val(addr.GetState())
	country := ptr.Val(addr.GetCountryOrRegion())
	postal := ptr.Val(addr.GetPostalCode())

	segments := []string{dn, street, city, state, country, postal}

	nonEmpty := []string{}
	for _, seg := range segments {
		if len(seg) > 0 {
			nonEmpty = append(nonEmpty, seg)
		}
	}

	if len(nonEmpty) > 0 {
		loc = strings.Join(nonEmpty, ", ")
	}

	return loc
}

func getUTCTime(ts, tz string) (time.Time, error) {
	// Timezone is always converted to UTC. We should do
	// this everywhere. This is the easiest way to ensure we
	// have the correct time everywhere and is the same(need
	// to same according to spec).
	it, err := dttm.ParseTime(ts)
	if err != nil {
		return time.Now(), clues.Wrap(err, "parsing time")
	}

	timezone := GraphTimeZoneToTZ[tz]
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
func getReccurencePattern(recurrence models.PatternedRecurrenceable) (string, error) {
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

	day := ptr.Val(pat.GetDayOfMonth())
	if day > 0 {
		recurComponents = append(recurComponents, "BYMONTHDAY="+fmt.Sprint(day))
	}

	dow := pat.GetDaysOfWeek()
	if dow != nil {
		dowComponents := []string{}
		for _, day := range dow {
			dowComponents = append(dowComponents, GraphToICalDOW[day.String()])
		}

		recurComponents = append(recurComponents, "BYDAY="+strings.Join(dowComponents, ","))
	}

	// TODO: I saw this being prepended to BYDAY. Valiate.
	index := pat.GetIndex()
	if index != nil &&
		(ptr.Val(freq) == models.RELATIVEMONTHLY_RECURRENCEPATTERNTYPE ||
			ptr.Val(freq) == models.RELATIVEYEARLY_RECURRENCEPATTERNTYPE) {
		recurComponents = append(recurComponents, "BYSETPOS="+fmt.Sprint(GraphToICalIndex[index.String()]))
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

				recurComponents = append(recurComponents, "UNTIL="+endTime.Format("20060102T150405Z"))
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
	data, err := api.BytesToEventable(body)
	if err != nil {
		return "", clues.Wrap(err, "converting to eventable")
	}

	cal := ics.NewCalendar()
	cal.SetProductId("-//Alcion//Corso") // Does this have to be customizable?

	id := data.GetId() // XXX: iCalUId?
	event := cal.AddEvent(ptr.Val(id))

	created := data.GetCreatedDateTime()
	if created != nil {
		event.SetCreatedTime(ptr.Val(created))
	}

	modified := data.GetLastModifiedDateTime()
	if modified != nil {
		event.SetModifiedAt(ptr.Val(modified))
	}

	allDay := ptr.Val(data.GetIsAllDay())

	startString := data.GetStart().GetDateTime()
	timeZone := data.GetStart().GetTimeZone()
	if startString != nil {
		start, err := getUTCTime(ptr.Val(startString), ptr.Val(timeZone))
		if err != nil {
			return "", clues.Wrap(err, "parsing start time")
		}

		if allDay {
			event.SetStartAt(start, ics.WithValue(string(ics.ValueDataTypeDate)))
		} else {
			event.SetStartAt(start)
		}

	}

	endString := data.GetEnd().GetDateTime()
	timeZone = data.GetEnd().GetTimeZone()
	if endString != nil {
		end, err := getUTCTime(ptr.Val(endString), ptr.Val(timeZone))
		if err != nil {
			return "", clues.Wrap(err, "parsing end time")
		}

		if allDay {
			event.SetEndAt(end, ics.WithValue(string(ics.ValueDataTypeDate)))
		} else {
			event.SetEndAt(end)
		}
	}

	recurrence := data.GetRecurrence()
	if recurrence != nil {
		pattern, err := getReccurencePattern(recurrence)
		if err != nil {
			return "", clues.WrapWC(ctx, err, "generating RRULE")
		}

		event.AddRrule(pattern)
	}

	cancelled := data.GetIsCancelled()
	if cancelled != nil {
		event.SetStatus(ics.ObjectStatusCancelled)
	}

	draft := data.GetIsDraft()
	if draft != nil {
		event.SetStatus(ics.ObjectStatusDraft)
	}

	summary := data.GetSubject()
	if summary != nil {
		event.SetSummary(ptr.Val(summary))
	}

	// TODO: Emojies seem to mess up the ics file
	bodyPreview := ptr.Val(data.GetBodyPreview())
	description := ptr.Val(data.GetBody().GetContent())
	contentType := data.GetBody().GetContentType().String()
	if len(description) > 0 && contentType == "text" && bodyPreview == description {
		event.SetDescription(description)
	} else {
		// https://stackoverflow.com/a/859475
		event.SetDescription(bodyPreview)
		if contentType == "html" {
			desc := strings.ReplaceAll(description, "\r\n", "")
			desc = strings.ReplaceAll(desc, "\n", "")
			event.AddProperty("X-ALT-DESC", desc, ics.WithFmtType("text/html"))
		}
	}

	showAs := ptr.Val(data.GetShowAs()).String()
	if len(showAs) > 0 {
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

		event.AddProperty(ics.ComponentPropertyFreebusy, string(status))
	}

	categories := data.GetCategories()
	for _, category := range categories {
		event.AddProperty(ics.ComponentPropertyCategories, category)
	}

	// According to the RFC, this property may be used in a calendar
	// component to convey a location where a more dynamic rendition of
	// the calendar information associated with the calendar component
	// can be found.
	url := ptr.Val(data.GetWebLink())
	if len(url) > 0 {
		event.SetURL(url)
	}

	organizer := data.GetOrganizer()
	if organizer != nil {
		name := ptr.Val(organizer.GetEmailAddress().GetName())
		addr := ptr.Val(organizer.GetEmailAddress().GetAddress())

		// TODO: What to do if we only have a name?
		if len(name) > 0 && len(addr) > 0 {
			event.SetOrganizer(addr, ics.WithCN(name))
		} else if len(addr) > 0 {
			event.SetOrganizer(addr)
		}
	}

	attendees := data.GetAttendees()
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

		// Time when a status change occurred is not recorded
		status := ptr.Val(attendee.GetStatus().GetResponse()).String()
		if len(status) > 0 && status != "none" {
			var pstat ics.ParticipationStatus
			switch status {
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

		addr := ptr.Val(attendee.GetEmailAddress().GetAddress())
		event.AddAttendee(addr, props...)
	}

	location := getLocationString(data.GetLocation())
	if len(location) > 0 {
		event.SetLocation(location)
	}

	attachments := data.GetAttachments()
	if attachments != nil {
		for _, attachment := range attachments {
			props := []ics.PropertyParameter{}
			contentType := ptr.Val(attachment.GetContentType())
			name := ptr.Val(attachment.GetName())

			if len(name) > 0 {
				// TODO: FILENAME does not seem to be parsed by Outlook
				props = append(props,
					&ics.KeyValues{
						Key:   "FILENAME",
						Value: []string{name},
					})
			}

			cb, err := attachment.GetBackingStore().Get("contentBytes")
			if err != nil {
				return "", clues.Wrap(err, "getting attachment content")
			}

			content, ok := cb.([]uint8)
			if !ok {
				return "", clues.NewWC(ctx, "getting attachment content string")
			}

			props = append(props, ics.WithEncoding("base64"), ics.WithValue("BINARY"))
			if len(contentType) > 0 {
				props = append(props, ics.WithFmtType(contentType))
			}

			// TODO: Inline attachments still don't show up in Outlook
			inline := ptr.Val(attachment.GetIsInline())
			if inline {
				cidv, err := attachment.GetBackingStore().Get("contentId")
				if err != nil {
					return "", clues.Wrap(err, "getting attachment content id")
				}

				cid, err := str.AnyToString(cidv)
				if err != nil {
					return "", clues.Wrap(err, "getting attachment content id string")
				}

				props = append(props, keyValues("CID", cid))
			}

			event.AddAttachment(base64.StdEncoding.EncodeToString(content), props...)
		}
	}

	return cal.Serialize(), nil
}
