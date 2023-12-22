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
// attendees
// locations (can we have multiple locations)
// recurrence
// showAs (with limitations)

// Field in the backed up data that we cannot handle
// allowNewTimeProposals, hideAttendees, importance, isOnlineMeeting,
// isOrganizer, isReminderOn, onlineMeeting, onlineMeetingProvider,
// onlineMeetingUrl, originalEndTimeZone, originalStart,
// originalStartTimeZone, reminderMinutesBeforeStart, responseRequested,
// responseStatus, sensitivity

func FromJSON(ctx context.Context, body []byte) (string, error) {
	data, err := api.BytesToEventable(body)
	if err != nil {
		return "", clues.Wrap(err, "converting to eventable")
	}

	cal := ics.NewCalendar()
	cal.SetProductId("-//Alcion//Corso") // Does this have to be customizible?

	id := data.GetId() // XXX: iCalUId?
	event := cal.AddEvent(ptr.Val(id))

	cal.SetMethod(ics.MethodRequest) // TODO: validate

	created := data.GetCreatedDateTime()
	if created != nil {
		event.SetCreatedTime(ptr.Val(created))
	}

	modified := data.GetLastModifiedDateTime()
	if modified != nil {
		event.SetModifiedAt(ptr.Val(modified))
	}

	startString := data.GetStart().GetDateTime()
	if startString != nil {
		// TODO: Handle timezone for start and end
		start, err := time.Parse(string(dttm.M365DateTimeTimeZone), ptr.Val(startString))
		if err != nil {
			return "", clues.Wrap(err, "parsing start time")
		}

		event.SetStartAt(start)

	}

	endString := data.GetEnd().GetDateTime()
	if endString != nil {
		end, err := time.Parse(string(dttm.M365DateTimeTimeZone), ptr.Val(endString))
		if err != nil {
			return "", clues.Wrap(err, "parsing end time")
		}

		event.SetEndAt(end)
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
	description := data.GetBody().GetContent()
	if description != nil {
		event.SetDescription(ptr.Val(description))
	}

	categories := data.GetCategories()
	if categories != nil {
		for _, category := range categories {
			event.AddProperty(ics.ComponentPropertyCategories, category)
		}
	}

	url := data.GetWebLink()
	if url != nil {
		event.SetURL(ptr.Val(url))
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

	// TODO: We should ideally encode full address
	location := data.GetLocation().GetDisplayName()
	if location != nil {
		event.SetLocation(ptr.Val(location))
	}

	attachments := data.GetAttachments()
	if attachments != nil {
		for _, attachment := range attachments {
			contentType := ptr.Val(attachment.GetContentType())
			// TODO: Can we have name?
			// name := attachment.GetName()

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

			props := []ics.PropertyParameter{}
			props = append(props, ics.WithEncoding("base64"))
			if len(contentType) > 0 {
				props = append(props, ics.WithFmtType(contentType))
			}

			event.AddAttachment(base64.StdEncoding.EncodeToString(content), props...)
		}
	}

	return cal.Serialize(), nil
}
