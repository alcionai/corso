package eml

// This package helps convert from the json response
// received from Graph API to .eml format (rfc0822).

// RFC
// Original: https://www.ietf.org/rfc/rfc0822.txt
// New: https://datatracker.ietf.org/doc/html/rfc5322
// Extension for MIME: https://www.ietf.org/rfc/rfc1521.txt

// Data missing from backup:
// SetReturnPath SetPriority SetListUnsubscribe SetDkim
// AddAlternative SetDSN (and any other X-MS specific headers)

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/converters/ics"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	addressFormat = `"%s" <%s>`
	dateFormat    = "2006-01-02 15:04:05 MST" // from xhit/go-simple-mail
)

func formatAddress(entry models.EmailAddressable) string {
	name := ptr.Val(entry.GetName())
	email := ptr.Val(entry.GetAddress())

	if len(name) == 0 && len(email) == 0 {
		return ""
	}

	if len(email) == 0 {
		return fmt.Sprintf(`"%s"`, name)
	}

	if name == email || len(name) == 0 {
		return email
	}

	return fmt.Sprintf(addressFormat, name, email)
}

// getICalData converts the emails to an event so that ical generation
// can generate from it.
func getICalData(ctx context.Context, data models.Messageable) (string, error) {
	msg, ok := data.(*models.EventMessageRequest)
	if !ok {
		return "", clues.NewWC(ctx, "unexpected message type").
			With("interface_type", fmt.Sprintf("%T", data))
	}

	// This method returns nil if data is not pulled using the necessary expand property
	// .../messages/<message_id>/?expand=Microsoft.Graph.EventMessage/Event
	// Also works for emails which are a result of someone accepting an
	// invite. If we add this expand query parameter value when directly
	// fetching a cancellation mail, the request fails.  It however looks
	// to be OK to run when listing emails although it gives empty({})
	// event value for cancellations.
	// TODO(meain): cancelled event details are available when pulling .eml
	if mevent := msg.GetEvent(); mevent != nil {
		return ics.FromEventable(ctx, mevent)
	}

	// Exceptions(modifications) are covered under this, although graph just sends the
	// exception event and not the parent, which what eml obtained from graph also contains
	if ptr.Val(msg.GetMeetingMessageType()) != models.MEETINGREQUEST_MEETINGMESSAGETYPE {
		// We don't have event data if it not "REQUEST" type.
		// Both cancellation and acceptance does not return enough
		// information to recreate an event.
		return "", nil
	}

	// If data was not fetch with an expand property, then we can
	// approximate the details with the following
	event := models.NewEvent()
	event.SetId(msg.GetId())
	event.SetCreatedDateTime(msg.GetCreatedDateTime())
	event.SetLastModifiedDateTime(msg.GetLastModifiedDateTime())
	event.SetIsAllDay(msg.GetIsAllDay())
	event.SetStart(msg.GetStartDateTime())
	event.SetEnd(msg.GetEndDateTime())
	event.SetRecurrence(msg.GetRecurrence())
	// event.SetIsCancelled()
	event.SetSubject(msg.GetSubject())
	event.SetBodyPreview(msg.GetBodyPreview())
	event.SetBody(msg.GetBody())

	// https://learn.microsoft.com/en-us/graph/api/resources/eventmessage?view=graph-rest-1.0
	// In addition, Outlook automatically creates an event instance in
	// the invitee's calendar, with the showAs property as tentative.
	event.SetShowAs(ptr.To(models.TENTATIVE_FREEBUSYSTATUS))

	event.SetCategories(msg.GetCategories())
	event.SetWebLink(msg.GetWebLink())
	event.SetOrganizer(msg.GetFrom())

	// NOTE: If an event was previously created and we added people to
	// it, the original list of attendee are not available.
	atts := []models.Attendeeable{}

	for _, to := range msg.GetToRecipients() {
		att := models.NewAttendee()
		att.SetEmailAddress(to.GetEmailAddress())
		att.SetTypeEscaped(ptr.To(models.REQUIRED_ATTENDEETYPE))
		atts = append(atts, att)
	}

	for _, cc := range msg.GetCcRecipients() {
		att := models.NewAttendee()
		att.SetEmailAddress(cc.GetEmailAddress())
		att.SetTypeEscaped(ptr.To(models.OPTIONAL_ATTENDEETYPE))
		atts = append(atts, att)
	}

	// bcc did not show up in my tests, but adding for completeness
	for _, bcc := range msg.GetBccRecipients() {
		att := models.NewAttendee()
		att.SetEmailAddress(bcc.GetEmailAddress())
		att.SetTypeEscaped(ptr.To(models.OPTIONAL_ATTENDEETYPE))
		atts = append(atts, att)
	}

	event.SetAttendees(atts)

	event.SetLocation(msg.GetLocation())
	// event.SetSensitivity() // unavailable in msg
	event.SetImportance(msg.GetImportance())
	// event.SetOnlineMeeting() // not available in eml either
	event.SetAttachments(msg.GetAttachments())

	return ics.FromEventable(ctx, event)
}

// FromJSON converts a Messageable (as json) to .eml format
func FromJSON(ctx context.Context, body []byte) (string, error) {
	ctx = clues.Add(ctx, "body_len", len(body))

	data, err := api.BytesToMessageable(body)
	if err != nil {
		return "", clues.WrapWC(ctx, err, "converting to messageble")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(data.GetId()))

	email := mail.NewMSG()
	email.Encoding = mail.EncodingBase64 // Doing it to be safe for when we have eventMessage (newline issues)
	email.AllowDuplicateAddress = true   // More "correct" conversion
	email.AddBccToHeader = true          // Don't ignore Bcc
	email.AllowEmptyAttachments = true   // Don't error on empty attachments
	email.UseProvidedAddress = true      // Don't try to parse the email address

	if data.GetFrom() != nil {
		email.SetFrom(formatAddress(data.GetFrom().GetEmailAddress()))
	}

	if data.GetToRecipients() != nil {
		for _, recipient := range data.GetToRecipients() {
			email.AddTo(formatAddress(recipient.GetEmailAddress()))
		}
	}

	if data.GetCcRecipients() != nil {
		for _, recipient := range data.GetCcRecipients() {
			email.AddCc(formatAddress(recipient.GetEmailAddress()))
		}
	}

	if data.GetBccRecipients() != nil {
		for _, recipient := range data.GetBccRecipients() {
			email.AddBcc(formatAddress(recipient.GetEmailAddress()))
		}
	}

	if data.GetReplyTo() != nil {
		rts := data.GetReplyTo()
		if len(rts) > 1 {
			logger.Ctx(ctx).
				With("reply_to_count", len(rts)).
				Warn("more than 1 Reply-To, adding only the first one")
		}

		if len(rts) != 0 {
			email.SetReplyTo(formatAddress(rts[0].GetEmailAddress()))
		}
	}

	if data.GetSubject() != nil {
		email.SetSubject(ptr.Val(data.GetSubject()))
	}

	if data.GetSentDateTime() != nil {
		email.SetDate(ptr.Val(data.GetSentDateTime()).Format(dateFormat))
	}

	if data.GetBody() != nil {
		if data.GetBody().GetContentType() != nil {
			var contentType mail.ContentType

			switch data.GetBody().GetContentType().String() {
			case "html":
				contentType = mail.TextHTML
			case "text":
				contentType = mail.TextPlain
			default:
				// https://learn.microsoft.com/en-us/graph/api/resources/itembody?view=graph-rest-1.0#properties
				// This should not be possible according to the documentation
				logger.Ctx(ctx).
					With("body_type", data.GetBody().GetContentType().String()).
					Info("unknown body content type")

				contentType = mail.TextPlain
			}

			email.SetBody(contentType, ptr.Val(data.GetBody().GetContent()))
		}
	}

	if data.GetAttachments() != nil {
		for _, attachment := range data.GetAttachments() {
			kind := ptr.Val(attachment.GetContentType())

			bytes, err := attachment.GetBackingStore().Get("contentBytes")
			if err != nil {
				return "", clues.WrapWC(ctx, err, "failed to get attachment bytes").
					With("kind", kind)
			}

			if bytes == nil {
				// Some attachments have an "item" field instead of
				// "contentBytes". There are items like contacts, emails
				// or calendar events which will not be a normal format
				// and will have to be converted to a text format.
				// TODO(meain): Handle custom attachments
				// https://github.com/alcionai/corso/issues/4772
				logger.Ctx(ctx).
					With("attachment_id", ptr.Val(attachment.GetId())).
					Info("unhandled attachment type")

				continue
			}

			bts, ok := bytes.([]byte)
			if !ok {
				return "", clues.WrapWC(ctx, err, "invalid content bytes").
					With("kind", kind).
					With("interface_type", fmt.Sprintf("%T", bytes))
			}

			name := ptr.Val(attachment.GetName())

			contentID, err := attachment.GetBackingStore().Get("contentId")
			if err != nil {
				return "", clues.WrapWC(ctx, err, "getting content id for attachment").
					With("kind", kind)
			}

			if contentID != nil {
				cids, _ := str.AnyToString(contentID)
				if len(cids) > 0 {
					name = cids
				}
			}

			email.Attach(&mail.File{
				// cannot use filename as inline attachment will not get mapped properly
				Name:     name,
				MimeType: kind,
				Data:     bts,
				Inline:   ptr.Val(attachment.GetIsInline()),
			})
		}
	}

	switch data.(type) {
	case *models.EventMessageResponse, *models.EventMessage:
		// We can't handle this as of now, not enough information
		// TODO: Fetch event object from graph when fetching email
	case *models.CalendarSharingMessage:
		// TODO: Parse out calendar sharing message
		// https://github.com/alcionai/corso/issues/5041
	case *models.EventMessageRequest:
		cal, err := getICalData(ctx, data)
		if err != nil {
			return "", clues.Wrap(err, "getting ical attachment")
		}

		if len(cal) > 0 {
			email.AddAlternative(mail.TextCalendar, cal)
		}
	}

	if err = email.GetError(); err != nil {
		return "", clues.WrapWC(ctx, err, "converting to eml")
	}

	return email.GetMessage(), nil
}
