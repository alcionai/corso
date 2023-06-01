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
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ itemRestorer = &mailRestoreHandler{}

type mailRestoreHandler struct {
	ac api.Mail
	ip itemPoster[models.Messageable]
}

func newMailRestoreHandler(
	ac api.Client,
) mailRestoreHandler {
	acm := ac.Mail()

	return mailRestoreHandler{
		ac: acm,
		ip: acm,
	}
}

func (h mailRestoreHandler) newContainerCache(userID string) graph.ContainerResolver {
	return &mailFolderCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h mailRestoreHandler) formatRestoreDestination(
	destinationContainerName string,
	collectionFullPath path.Path,
) *path.Builder {
	return path.Builder{}.Append(destinationContainerName).Append(collectionFullPath.Folders()...)
}

func (h mailRestoreHandler) CreateContainer(
	ctx context.Context,
	userID, containerName, parentContainerID string,
) (graph.Container, error) {
	if len(parentContainerID) == 0 {
		parentContainerID = rootFolderAlias
	}

	return h.ac.CreateContainer(ctx, userID, containerName, parentContainerID)
}

func (h mailRestoreHandler) containerSearcher() containerByNamer {
	return nil
}

// always returns rootFolderAlias
func (h mailRestoreHandler) orRootContainer(string) string {
	return rootFolderAlias
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
		h.ac,
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
	svlep := make([]models.SingleValueLegacyExtendedPropertyable, 0)

	// prevent "resending" of the mail in the graph api backstore
	sv1 := models.NewSingleValueLegacyExtendedProperty()
	sv1.SetId(ptr.To(MailRestorePropertyTag))
	sv1.SetValue(ptr.To(RestoreCanonicalEnableValue))
	svlep = append(svlep, sv1)

	// establish the sent date
	if msg.GetSentDateTime() != nil {
		sv2 := models.NewSingleValueLegacyExtendedProperty()
		sv2.SetId(ptr.To(MailSendDateTimeOverrideProperty))
		sv2.SetValue(ptr.To(dttm.FormatToLegacy(ptr.Val(msg.GetSentDateTime()))))

		svlep = append(svlep, sv2)
	}

	// establish the received Date
	if msg.GetReceivedDateTime() != nil {
		sv3 := models.NewSingleValueLegacyExtendedProperty()
		sv3.SetId(ptr.To(MailReceiveDateTimeOverriveProperty))
		sv3.SetValue(ptr.To(dttm.FormatToLegacy(ptr.Val(msg.GetReceivedDateTime()))))

		svlep = append(svlep, sv3)
	}

	msg.SetSingleValueExtendedProperties(svlep)

	return msg
}
