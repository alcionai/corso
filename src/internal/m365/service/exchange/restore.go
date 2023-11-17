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
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections restores M365 objects in data.RestoreCollection to MSFT
// store through GraphAPI.
func ConsumeRestoreCollections(
	ctx context.Context,
	ac api.Client,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, *support.ControllerOperationStatus, error) {
	if len(dcs) == 0 {
		return nil, support.CreateStatus(ctx, support.Restore, 0, support.CollectionMetrics{}, ""), nil
	}

	var (
		deets          = &details.Builder{}
		resourceID     = rcc.ProtectedResource.ID()
		directoryCache = make(map[path.CategoryType]graph.ContainerResolver)
		handlers       = exchange.RestoreHandlers(ac)
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

		containerID, gcc, err := exchange.CreateDestination(
			ictx,
			handler,
			handler.FormatRestoreDestination(rcc.RestoreConfig.Location, dc.FullPath()),
			resourceID,
			directoryCache[category],
			errs)
		if err != nil {
			el.AddRecoverable(ictx, err)
			continue
		}

		directoryCache[category] = gcc
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

	return deets.Details(), status, el.Failure()
}
