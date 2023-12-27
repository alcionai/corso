package ics

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	ics "github.com/arran4/golang-ical"
)

// This package is used to convert json response from graph to ics
// Ref: https://icalendar.org/
// Ref: https://www.rfc-editor.org/rfc/rfc5545
// Ref: https://learn.microsoft.com/en-us/graph/api/resources/event?view=graph-rest-1.0

// TODO: Items not handled
// locations (can we have multiple locations)
// recurrence
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
		start, err := time.Parse(string(dttm.M365DateTimeTimeZone), ptr.Val(startString))
		if err != nil {
			return "", clues.Wrap(err, "parsing start time")
		}

		// TODO: Timezone from graph is UTC or Indian Standard Time,
		// but we need values like Asia/Kolkata
		props := []ics.PropertyParameter{keyValues(string(ics.PropertyTzid), ptr.Val(timeZone))}
		if allDay {
			props = append(props, ics.WithValue(string(ics.ValueDataTypeDate)))
		}

		event.SetStartAt(start, props...)

	}

	endString := data.GetEnd().GetDateTime()
	timeZone = data.GetEnd().GetTimeZone()
	if endString != nil {
		end, err := time.Parse(string(dttm.M365DateTimeTimeZone), ptr.Val(endString))
		if err != nil {
			return "", clues.Wrap(err, "parsing end time")
		}

		props := []ics.PropertyParameter{keyValues(string(ics.PropertyTzid), ptr.Val(timeZone))}
		if allDay {
			props = append(props, ics.WithValue(string(ics.ValueDataTypeDate)))
		}
		event.SetEndAt(end, props...)
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

	// Description could be HTML, but we have not way to differentiate
	// in the output
	description := ptr.Val(data.GetBody().GetContent())
	if len(description) > 0 {
		event.SetDescription(description)
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

	// TODO: We should ideally encode full address
	location := data.GetLocation().GetDisplayName()
	if location != nil {
		event.SetLocation(ptr.Val(location))
	}

	attachments := data.GetAttachments()
	if attachments != nil {
		for _, attachment := range attachments {
			props := []ics.PropertyParameter{}
			contentType := ptr.Val(attachment.GetContentType())
			name := ptr.Val(attachment.GetName())

			if len(name) > 0 {
				props = append(props,
					&ics.KeyValues{
						Key:   "FILENAME",
						Value: []string{name},
					})
			}

			// TODO: What is the deal with inline?
			// inline := ptr.Val(attachment.GetIsInline())

			cb, err := attachment.GetBackingStore().Get("contentBytes")
			if err != nil {
				return "", clues.Wrap(err, "getting attachment content")
			}

			content, ok := cb.([]uint8)
			if !ok {
				return "", clues.NewWC(ctx, "getting attachment content string")
			}

			props = append(props, ics.WithEncoding("base64"))
			if len(contentType) > 0 {
				props = append(props, ics.WithFmtType(contentType))
			}

			event.AddAttachment(base64.StdEncoding.EncodeToString(content), props...)
		}
	}

	return cal.Serialize(), nil
}
