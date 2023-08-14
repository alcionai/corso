package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

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

var _ itemRestorer = &contactRestoreHandler{}

type contactRestoreHandler struct {
	ac api.Contacts
}

func newContactRestoreHandler(
	ac api.Client,
) contactRestoreHandler {
	return contactRestoreHandler{
		ac: ac.Contacts(),
	}
}

func (h contactRestoreHandler) NewContainerCache(userID string) graph.ContainerResolver {
	return &contactContainerCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

func (h contactRestoreHandler) FormatRestoreDestination(
	destinationContainerName string,
	_ path.Path, // contact folders cannot be nested
) *path.Builder {
	return path.Builder{}.Append(destinationContainerName)
}

func (h contactRestoreHandler) CreateContainer(
	ctx context.Context,
	userID, _, containerName string, // parent container not used
) (graph.Container, error) {
	return h.ac.CreateContainer(ctx, userID, "", containerName)
}

func (h contactRestoreHandler) GetContainerByName(
	ctx context.Context,
	userID, _, containerName string, // parent container not used
) (graph.Container, error) {
	return h.ac.GetContainerByName(ctx, userID, "", containerName)
}

func (h contactRestoreHandler) DefaultRootContainer() string {
	return api.DefaultContacts
}

func (h contactRestoreHandler) restore(
	ctx context.Context,
	body []byte,
	userID, destinationID string,
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.ExchangeInfo, error) {
	return restoreContact(
		ctx,
		h.ac,
		body,
		userID, destinationID,
		collisionKeyToItemID,
		collisionPolicy,
		errs,
		ctr)
}

type contactRestorer interface {
	postItemer[models.Contactable]
	deleteItemer
}

func restoreContact(
	ctx context.Context,
	cr contactRestorer,
	body []byte,
	userID, destinationID string,
	collisionKeyToItemID map[string]string,
	collisionPolicy control.CollisionPolicy,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.ExchangeInfo, error) {
	// contacts has a weird relationship with its default
	// folder, which is that the folder is treated as invisible
	// in many cases.  If we're restoring to a blank location,
	// we can interpret that as the root.
	if len(destinationID) == 0 {
		destinationID = api.DefaultContacts
	}

	contact, err := api.BytesToContactable(body)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact from bytes")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))

	var (
		collisionKey         = api.ContactCollisionKey(contact)
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

	item, err := cr.PostItem(ctx, userID, destinationID, contact)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "restoring contact")
	}

	// contacts have no PUT request, and PATCH could retain data that's not
	// associated with the backup item state.  Instead of updating, we
	// post first, then delete.  In case of failure between the two calls,
	// at least we'll have accidentally over-produced data instead of deleting
	// the user's data.
	if shouldDeleteOriginal {
		if err := cr.DeleteItem(ctx, userID, collisionID); err != nil && !graph.IsErrDeletedInFlight(err) {
			return nil, graph.Wrap(ctx, err, "deleting colliding contact")
		}
	}

	info := api.ContactInfo(item)
	info.Size = int64(len(body))

	if shouldDeleteOriginal {
		ctr.Inc(count.CollisionReplace)
	} else {
		ctr.Inc(count.NewItemCreated)
	}

	return info, nil
}

func (h contactRestoreHandler) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	m, err := h.ac.GetItemsInContainerByCollisionKey(ctx, userID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}
