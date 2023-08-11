package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ itemRestorer = &mailRestoreHandler{}

type mailRestoreHandler struct {
	ac api.Mail
}

func newMailRestoreHandler(
	ac api.Client,
) mailRestoreHandler {
	return mailRestoreHandler{
		ac: ac.Mail(),
	}
}

func (h mailRestoreHandler) newContainerCache(userID string) graph.ContainerResolver {
	return &mailContainerCache{
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
	userID, parentContainerID, containerName string,
) (graph.Container, error) {
	if len(parentContainerID) == 0 {
		parentContainerID = api.MsgFolderRoot
	}

	return h.ac.CreateContainer(ctx, userID, parentContainerID, containerName)
}

func (h mailRestoreHandler) GetContainerByName(
	ctx context.Context,
	userID, parentContainerID, containerName string,
) (graph.Container, error) {
	return h.ac.GetContainerByName(ctx, userID, parentContainerID, containerName)
}

func (h mailRestoreHandler) defaultRootContainer() string {
	return api.MsgFolderRoot
}

func (h mailRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	userID, destinationID string,
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.ExchangeInfo, error) {
	return restoreMail(
		ctx,
		h.ac,
		body,
		userID, destinationID,
		collisionKeyToItemID,
		collisionPolicy,
		errs,
		ctr)
}

type mailRestorer interface {
	postItemer[models.Messageable]
	deleteItemer
	attachmentPoster
}

func restoreMail(
	ctx context.Context,
	mr mailRestorer,
	body []byte,
	userID, destinationID string,
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.ExchangeInfo, error) {
	msg, err := api.BytesToMessageable(body)
	if err != nil {
		return nil, clues.Wrap(err, "creating mail from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(msg.GetId()))

	var (
		collisionKey         = api.MailCollisionKey(msg)
		collisionID          string
		shouldDeleteOriginal bool
	)

	if id, ok := collisionKeyToItemID[collisionKey]; ok {
		log := logger.Ctx(ctx).With("collision_key", clues.Hide(collisionKey))
		log.Debug("item collision")

		if collisionPolicy == control.Skip {
			ctr.Inc(count.CollisionSkip)
			log.Debug("skipping item with collision")

			return nil, graph.ErrItemAlreadyExistsConflict
		}

		collisionID = id
		shouldDeleteOriginal = collisionPolicy == control.Replace
	}

	msg = setMessageSVEPs(toMessage(msg))

	attachments := msg.GetAttachments()
	// Item.Attachments --> HasAttachments doesn't always have a value populated when deserialized
	msg.SetAttachments([]models.Attachmentable{})

	item, err := mr.PostItem(ctx, userID, destinationID, msg)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring mail message")
	}

	// mails have no PUT request, and PATCH could retain data that's not
	// associated with the backup item state.  Instead of updating, we
	// post first, then delete.  In case of failure between the two calls,
	// at least we'll have accidentally over-produced data instead of deleting
	// the user's data.
	if shouldDeleteOriginal {
		if err := mr.DeleteItem(ctx, userID, collisionID); err != nil && !graph.IsErrDeletedInFlight(err) {
			return nil, graph.Wrap(ctx, err, "deleting colliding mail message")
		}
	}

	err = uploadAttachments(
		ctx,
		mr,
		attachments,
		userID,
		destinationID,
		ptr.Val(item.GetId()),
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	var size int64

	if msg.GetBody() != nil {
		bc := ptr.Val(msg.GetBody().GetContent())
		size = int64(len(bc))
	}

	if shouldDeleteOriginal {
		ctr.Inc(count.CollisionReplace)
	} else {
		ctr.Inc(count.NewItemCreated)
	}

	return api.MailInfo(msg, size), nil
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

func (h mailRestoreHandler) getItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	m, err := h.ac.GetItemsInContainerByCollisionKey(ctx, userID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}
