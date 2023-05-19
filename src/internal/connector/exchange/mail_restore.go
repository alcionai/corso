package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ itemRestorer = &mailRestoreHandler{}

type mailRestoreHandler struct {
	ac api.Mail
	ip itemPoster[models.Messageable]
	ap attachmentPoster
}

func newMailRestoreHandler(
	ac api.Client,
) mailRestoreHandler {
	acm := ac.Mail()

	return mailRestoreHandler{
		ac: acm,
		ip: acm,
		ap: acm,
	}
}

func (h mailRestoreHandler) newContainerCache(userID string) graph.ContainerResolver {
	return &mailFolderCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h mailRestoreHandler) containerFactory() containerCreator {
	return h.ac
}

func (h mailRestoreHandler) containerSearcher() (containerByNamer, bool) {
	return nil, false
}

func (h mailRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	userID, destinationID string,
	errs *fault.Bus,
) (*details.ExchangeInfo, error) {
	msg, err := api.BytesToMessageable(body)
	if err != nil {
		return nil, clues.Wrap(err, "creating mail from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(msg.GetId()))

	msg = setMessageSVEPs(toMessage(msg))

	attachments := msg.GetAttachments()
	// Item.Attachments --> HasAttachments doesn't always have a value populated when deserialized
	msg.SetAttachments([]models.Attachmentable{})

	item, err := h.ip.PostItem(ctx, userID, destinationID, msg)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring mail message")
	}

	err = uploadAttachments(
		ctx,
		h.ap,
		attachments,
		userID,
		destinationID,
		ptr.Val(item.GetId()),
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return api.MailInfo(msg, int64(len(body))), nil
}

func setMessageSVEPs(msg models.Messageable) models.Messageable {
	// Set Extended Properties:
	// 1st: No transmission
	// 2nd: Send Date
	// 3rd: Recv Date
	svlep := make([]models.SingleValueLegacyExtendedPropertyable, 0)
	sv1 := models.NewSingleValueLegacyExtendedProperty()
	sv1.SetId(ptr.To(MailRestorePropertyTag))
	sv1.SetValue(ptr.To(MailRestorePropertyTag))
	svlep = append(svlep, sv1)

	if msg.GetSentDateTime() != nil {
		sv2 := models.NewSingleValueLegacyExtendedProperty()
		sendPropertyValue := dttm.FormatToLegacy(ptr.Val(msg.GetSentDateTime()))
		sendPropertyTag := MailSendDateTimeOverrideProperty
		sv2.SetId(&sendPropertyTag)
		sv2.SetValue(&sendPropertyValue)

		svlep = append(svlep, sv2)
	}

	if msg.GetReceivedDateTime() != nil {
		sv3 := models.NewSingleValueLegacyExtendedProperty()
		recvPropertyValue := dttm.FormatToLegacy(ptr.Val(msg.GetReceivedDateTime()))
		recvPropertyTag := MailReceiveDateTimeOverriveProperty
		sv3.SetId(&recvPropertyTag)
		sv3.SetValue(&recvPropertyValue)

		svlep = append(svlep, sv3)
	}

	msg.SetSingleValueExtendedProperties(svlep)

	return msg
}
