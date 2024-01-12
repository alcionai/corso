package sharepoint

import (
	"context"
	"errors"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/site"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections will restore the specified data collections into OneDrive
func (h *sharepointHandler) ConsumeRestoreCollections(
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
		graph.LimiterCfg{Service: path.SharePointService})

	var (
		deets = &details.Builder{}
		lrh   = drive.NewSiteRestoreHandler(
			h.apiClient,
			rcc.Selector.PathService())
		listsRh = site.NewListsRestoreHandler(
			rcc.ProtectedResource.ID(),
			h.apiClient.Lists())
		restoreMetrics support.CollectionMetrics

		caches = drive.NewRestoreCaches(h.backupDriveIDNames)

		el = errs.Local()
		cl = ctr.Local()
	)

	// Reorder collections so that the parents directories are created
	// before the child directories; a requirement for permissions.
	data.SortRestoreCollections(dcs)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		if el.Failure() != nil {
			break
		}

		var (
			err      error
			category = dc.FullPath().Category()
			metrics  support.CollectionMetrics
			ictx     = clues.Add(ctx,
				"category", category,
				"restore_location", clues.Hide(rcc.RestoreConfig.Location),
				"resource_owner", clues.Hide(dc.FullPath().ProtectedResource()),
				"full_path", dc.FullPath())
			collisionKeyToItemID map[string]string
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
			err = caches.Populate(ctx, h.apiClient.Users(), h.apiClient.Groups(), lrh, rcc.ProtectedResource.ID(), errs)
			if err != nil {
				return nil, nil, clues.Wrap(err, "initializing restore caches")
			}

			metrics, err = drive.RestoreCollection(
				ictx,
				lrh,
				rcc,
				dc,
				caches,
				deets,
				control.DefaultRestoreContainerName(dttm.HumanReadableDriveItem),
				errs,
				ctr)

		case path.ListsCategory:
			collisionKeyToItemID, err = listsRh.GetListsByCollisionKey(ictx)
			if err != nil {
				el.AddRecoverable(ictx, clues.Wrap(err, "building lists collision map"))
				continue
			}

			metrics, err = site.RestoreListCollection(
				ictx,
				listsRh,
				dc,
				rcc.RestoreConfig.Location,
				deets,
				collisionKeyToItemID,
				rcc.RestoreConfig.OnCollision,
				cl,
				errs)

		case path.PagesCategory:
			metrics, err = site.RestorePageCollection(
				ictx,
				h.apiClient.Stable,
				dc,
				rcc.RestoreConfig.Location,
				deets,
				errs)

		default:
			return nil, nil, clues.Wrap(clues.New(category.String()), "category not supported").With("category", category)
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if err != nil {
			el.AddRecoverable(ctx, err)
		}

		if errors.Is(err, context.Canceled) {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		restoreMetrics,
		rcc.RestoreConfig.Location)

	return deets.Details(), status.ToCollectionStats(), el.Failure()
}
