package eml

// This package helps convert from the json response
// received from Graph API to .eml format (rfc0822).
// Ref: https://www.ietf.org/rfc/rfc0822.txt
// Ref: https://datatracker.ietf.org/doc/html/rfc5322
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
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	fromFormat = "%s <%s>"
	dateFormat = "2006-01-02 15:04:05 MST" // from xhit/go-simple-mail
)

// toEml converts a Messageable to .eml format
func toEml(data models.Messageable) (string, error) {
	email := mail.NewMSG()

	if data.GetFrom() != nil {
		email.SetFrom(
			fmt.Sprintf(
				fromFormat,
				ptr.Val(data.GetFrom().GetEmailAddress().GetName()),
				ptr.Val(data.GetFrom().GetEmailAddress().GetAddress())))
	}

	if data.GetToRecipients() != nil {
		for _, recipient := range data.GetToRecipients() {
			email.AddTo(
				fmt.Sprintf(
					fromFormat,
					ptr.Val(recipient.GetEmailAddress().GetName()),
					ptr.Val(recipient.GetEmailAddress().GetAddress())))
		}
	}

	if data.GetCcRecipients() != nil {
		for _, recipient := range data.GetCcRecipients() {
			email.AddCc(
				fmt.Sprintf(
					fromFormat,
					ptr.Val(recipient.GetEmailAddress().GetName()),
					ptr.Val(recipient.GetEmailAddress().GetAddress())))
		}
	}

	if data.GetBccRecipients() != nil {
		for _, recipient := range data.GetBccRecipients() {
			email.AddBcc(
				fmt.Sprintf(
					fromFormat,
					ptr.Val(recipient.GetEmailAddress().GetName()),
					ptr.Val(recipient.GetEmailAddress().GetAddress())))
		}
	}

	if data.GetReplyTo() != nil {
		rts := data.GetReplyTo()
		if len(rts) > 1 {
			logger.Ctx(context.TODO()).
				With("id", ptr.Val(data.GetId()),
					"reply_to_count", len(rts)).
				Warn("more than 1 reply to")
		} else if len(rts) != 0 {
			email.SetReplyTo(
				fmt.Sprintf(
					fromFormat,
					ptr.Val(rts[0].GetEmailAddress().GetName()),
					ptr.Val(rts[0].GetEmailAddress().GetAddress())))
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
				logger.Ctx(context.TODO()).
					With("body_type", data.GetBody().GetContentType().String(),
						"id", ptr.Val(data.GetId())).
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
				return "", clues.Wrap(err, "failed to get attachment bytes")
			}

			bts, ok := bytes.([]byte)
			if !ok {
				return "", clues.Wrap(err, "invalid content bytes")
			}

			email.Attach(&mail.File{
				Name:     ptr.Val(attachment.GetName()),
				MimeType: kind,
				Data:     bts,
				Inline:   ptr.Val(attachment.GetIsInline()),
			})
		}
	}

	return email.GetMessage(), nil
}
