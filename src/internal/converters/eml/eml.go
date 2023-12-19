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

// FromJSON converts a Messageable (as json) to .eml format
func FromJSON(ctx context.Context, body []byte) (string, error) {
	data, err := api.BytesToMessageable(body)
	if err != nil {
		return "", clues.Wrap(err, "converting to messageble")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(data.GetId()))

	email := mail.NewMSG()
	email.AllowDuplicateAddress = true // More "correct" conversion
	email.AddBccToHeader = true        // Don't ignore Bcc

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
				return "", clues.WrapWC(ctx, err, "failed to get attachment bytes")
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
				return "", clues.WrapWC(ctx, err, "invalid content bytes")
			}

			if len(bts) == 0 {
				// TODO(meain): pass the data through after
				// https://github.com/xhit/go-simple-mail/issues/96
				logger.Ctx(ctx).
					With("attachment_id", ptr.Val(attachment.GetId())).
					Info("empty attachment")

				continue
			}

			name := ptr.Val(attachment.GetName())

			contentID, err := attachment.GetBackingStore().Get("contentId")
			if err != nil {
				return "", clues.WrapWC(ctx, err, "getting content id for attachment")
			}

			if contentID != nil {
				cids, _ := contentID.(*string)
				if len(*cids) > 0 {
					name = *cids
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

	if err = email.GetError(); err != nil {
		return "", clues.WrapWC(ctx, err, "converting to eml")
	}

	return email.GetMessage(), nil
}
