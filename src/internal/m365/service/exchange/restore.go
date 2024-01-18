package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections restores M365 objects in data.RestoreCollection to MSFT
// store through GraphAPI.
func (h *exchangeHandler) ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, *data.CollectionStats, error) {
	if len(dcs) == 0 {
		return nil, nil, clues.WrapWC(ctx, data.ErrNoData, "performing restore")
	}

	// TODO(ashmrtn): We should stop relying on the context for rate limiter stuff
	// and instead configure this when we make the handler instance. We can't
	// initialize it in the NewHandler call right now because those functions
	// aren't (and shouldn't be) returning a context along with the handler. Since
	// that call isn't directly calling into this function even if we did
	// initialize the rate limiter there it would be lost because it wouldn't get
	// stored in an ancestor of the context passed to this function.
	ctx = graph.BindRateLimiterConfig(
		ctx,
		graph.LimiterCfg{Service: path.ExchangeService})

	var (
		deets          = &details.Builder{}
		resourceID     = rcc.ProtectedResource.ID()
		directoryCache = make(map[path.CategoryType]graph.ContainerResolver)
		handlers       = exchange.RestoreHandlers(h.apiClient)
		metrics        support.CollectionMetrics
		el             = errs.Local()
	)

	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			category = dc.FullPath().Category()
			ictx     = clues.Add(
				ctx,
				"restore_category", category,
				"restore_full_path", dc.FullPath())
		)

		handler, ok := handlers[category]
		if !ok {
			el.AddRecoverable(ictx, clues.NewWC(ictx, "unsupported restore path category"))
			continue
		}

		if directoryCache[category] == nil {
			gcr := handler.NewContainerCache(resourceID)
			if err := gcr.Populate(ictx, errs, handler.DefaultRootContainer()); err != nil {
				return nil, nil, clues.Wrap(err, "populating container cache")
			}

			directoryCache[category] = gcr
		}

		restoreFolderPath := handler.FormatRestoreDestination(rcc.RestoreConfig.Location, dc.FullPath())

		ictx = clues.Add(ictx, "restore_folder_path", restoreFolderPath)

		var containerID string

		// Only attempt to create a new folder if it's not the default contacts
		// folder. Contacts is weird in that it allows creating sub folders with the
		// same name as the default contacts folder, but which are technically
		// nested folders.
		//
		// Without this check we'll end up restoring to Contacts/Contacts instead of
		// Contacts if in-place restore is requested and the root Contacts folder
		// had items.
		if category == path.ContactsCategory &&
			len(dc.FullPath().Folders()) == 1 &&
			restoreFolderPath.String() == handler.DefaultRootContainer() {
			logger.Ctx(ictx).Info("using default contact folder")

			containerID = handler.DefaultRootContainer()
		} else {
			logger.Ctx(ictx).Info("creating restore folder")

			newContainerID, gcc, err := exchange.CreateDestination(
				ictx,
				handler,
				restoreFolderPath,
				resourceID,
				directoryCache[category],
				errs)
			if err != nil {
				el.AddRecoverable(ictx, err)
				continue
			}

			directoryCache[category] = gcc
			containerID = newContainerID
		}

		ictx = clues.Add(ictx, "restore_destination_id", containerID)

		collisionKeyToItemID, err := handler.GetItemsInContainerByCollisionKey(ictx, resourceID, containerID)
		if err != nil {
			el.AddRecoverable(ictx, clues.Wrap(err, "building item collision cache"))
			continue
		}

		temp, err := exchange.RestoreCollection(
			ictx,
			handler,
			dc,
			resourceID,
			containerID,
			collisionKeyToItemID,
			rcc.RestoreConfig.OnCollision,
			deets,
			errs,
			ctr)

		metrics = support.CombineMetrics(metrics, temp)

		if err != nil {
			if graph.IsErrTimeout(err) {
				break
			}

			el.AddRecoverable(ictx, err)
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		metrics,
		rcc.RestoreConfig.Location)

	return deets.Details(), status.ToCollectionStats(), el.Failure()
}
