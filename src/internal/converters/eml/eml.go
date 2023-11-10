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
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	addressFormat = "%s <%s>"
	dateFormat    = "2006-01-02 15:04:05 MST" // from xhit/go-simple-mail
)

func formatAddress(entry models.EmailAddressable) string {
	name := ptr.Val(entry.GetName())
	email := ptr.Val(entry.GetAddress())

	if name == email {
		return email
	}

	return fmt.Sprintf(addressFormat, name, email)
}

// FromJSON converts a Messageable (as json) to .eml format
func FromJSON(ctx context.Context, body []byte) (string, error) {
	data, err := api.BytesToMessageable(body)
	if err != nil {
		return "", clues.Wrap(err, "converting to messageble")
	}

	email := mail.NewMSG()

	if data.GetFrom() != nil {
		email.SetFrom(formatAddress(data.GetFrom().GetEmailAddress()))

		if email.Error != nil {
			return "", clues.Wrap(email.Error, "adding from address").
				With("id", ptr.Val(data.GetId()), "from", data.GetFrom())
		}
	}

	if data.GetToRecipients() != nil {
		for _, recipient := range data.GetToRecipients() {
			email.AddTo(formatAddress(recipient.GetEmailAddress()))

			if email.Error != nil {
				return "", clues.Wrap(email.Error, "adding to address").
					With("id", ptr.Val(data.GetId()), "to", recipient)
			}
		}
	}

	if data.GetCcRecipients() != nil {
		for _, recipient := range data.GetCcRecipients() {
			email.AddCc(formatAddress(recipient.GetEmailAddress()))

			if email.Error != nil {
				return "", clues.Wrap(email.Error, "adding cc address").
					With("id", ptr.Val(data.GetId()), "cc", recipient)
			}
		}
	}

	if data.GetBccRecipients() != nil {
		for _, recipient := range data.GetBccRecipients() {
			email.AddBcc(formatAddress(recipient.GetEmailAddress()))

			if email.Error != nil {
				return "", clues.Wrap(email.Error, "adding bcc address").
					With("id", ptr.Val(data.GetId()), "bcc", recipient)
			}
		}
	}

	if data.GetReplyTo() != nil {
		rts := data.GetReplyTo()
		if len(rts) > 1 {
			logger.Ctx(ctx).
				With("id", ptr.Val(data.GetId()),
					"reply_to_count", len(rts)).
				Warn("more than 1 Reply-To, adding only the first one")
		}

		if len(rts) != 0 {
			email.SetReplyTo(formatAddress(rts[0].GetEmailAddress()))
		}
	}

	if data.GetSubject() != nil {
		email.SetSubject(ptr.Val(data.GetSubject()))

		if email.Error != nil {
			return "", clues.Wrap(email.Error, "adding subject").
				With("id", ptr.Val(data.GetId()), "subject", data.GetSubject())
		}
	}

	if data.GetSentDateTime() != nil {
		email.SetDate(ptr.Val(data.GetSentDateTime()).Format(dateFormat))

		if email.Error != nil {
			return "", clues.Wrap(email.Error, "adding date").
				With("id", ptr.Val(data.GetId()), "date", data.GetSentDateTime())
		}
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
					With("body_type", data.GetBody().GetContentType().String(),
						"id", ptr.Val(data.GetId())).
					Info("unknown body content type")

				contentType = mail.TextPlain
			}

			email.SetBody(contentType, ptr.Val(data.GetBody().GetContent()))

			if email.Error != nil {
				return "", clues.Wrap(email.Error, "adding body").
					With("id", ptr.Val(data.GetId()),
						"body_type", data.GetBody().GetContentType().String(),
						"body_len", len(*data.GetBody().GetContent()))
			}
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

			if email.Error != nil {
				return "", clues.Wrap(email.Error, "adding attachment").
					With("id", ptr.Val(data.GetId()),
						"attachment_id", ptr.Val(attachment.GetId()),
						"attachment_name", ptr.Val(attachment.GetName()))
			}
		}
	}

	return email.GetMessage(), nil
}
