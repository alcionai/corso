package site

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/trace"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections will restore the specified data collections into OneDrive
func ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	ac api.Client,
	backupDriveIDNames idname.Cacher,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
	ctr *count.Bus,
) (*support.ControllerOperationStatus, error) {
	var (
		lrh            = drive.NewSiteRestoreHandler(ac, rcc.Selector.PathService())
		listsRh        = NewListsRestoreHandler(rcc.ProtectedResource.ID(), ac.Lists())
		restoreMetrics support.CollectionMetrics
		caches         = drive.NewRestoreCaches(backupDriveIDNames)
		el             = errs.Local()
	)

	err := caches.Populate(ctx, lrh, rcc.ProtectedResource.ID())
	if err != nil {
		return nil, clues.Wrap(err, "initializing restore caches")
	}

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
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
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
			metrics, err = RestoreListCollection(
				ictx,
				listsRh,
				dc,
				rcc.RestoreConfig.Location,
				deets,
				errs)

		case path.PagesCategory:
			metrics, err = RestorePageCollection(
				ictx,
				ac.Stable,
				dc,
				rcc.RestoreConfig.Location,
				deets,
				errs)

		default:
			return nil, clues.Wrap(clues.New(category.String()), "category not supported").With("category", category)
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

	return status, el.Failure()
}

// restoreListItem utility function restores a List to the siteID.
// The name is changed to to {DestName}_{name}
// API Reference: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=http
// Restored List can be verified within the Site contents.
func restoreListItem(
	ctx context.Context,
	rh restoreHandler,
	itemData data.Item,
	siteID, destName string,
	errs *fault.Bus,
) (details.ItemInfo, error) {
	ctx, end := diagnostics.Span(ctx, "m365:sharepoint:restoreList", diagnostics.Label("item_uuid", itemData.ID()))
	defer end()

	ctx = clues.Add(ctx, "list_item_id", itemData.ID())

	var (
		dii      = details.ItemInfo{}
		listName = itemData.ID()
	)

	bytes, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, clues.WrapWC(ctx, err, "reading backup data")
	}

	newName := fmt.Sprintf("%s_%s", destName, listName)

	// Restore to List base to M365 back store
	restoredList, err := rh.PostList(ctx, newName, bytes, errs)
	if err != nil {
		return dii, graph.Wrap(ctx, err, "restoring list")
	}

	dii.SharePoint = api.ListToSPInfo(restoredList)

	return dii, nil
}

func RestoreListCollection(
	ctx context.Context,
	rh restoreHandler,
	dc data.RestoreCollection,
	restoreContainerName string,
	deets *details.Builder,
	errs *fault.Bus,
) (support.CollectionMetrics, error) {
	ctx, end := diagnostics.Span(ctx, "m365:sharepoint:restoreListCollection", diagnostics.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics   = support.CollectionMetrics{}
		directory = dc.FullPath()
		siteID    = directory.ProtectedResource()
		items     = dc.Items(ctx, errs)
		el        = errs.Local()
	)

	trace.Log(ctx, "m365:sharepoint:restoreListCollection", directory.String())

	for {
		if el.Failure() != nil {
			break
		}

		select {
		case <-ctx.Done():
			return metrics, clues.StackWC(ctx, ctx.Err())

		case itemData, ok := <-items:
			if !ok {
				return metrics, nil
			}
			metrics.Objects++

			itemInfo, err := restoreListItem(
				ctx,
				rh,
				itemData,
				siteID,
				restoreContainerName,
				errs)
			if errors.Is(err, api.ErrSkippableListTemplate) {
				// should never be encountered as lists with skippable template are not backed up
				// this is an additional check
				logger.Ctx(ctx).Info("failed to create listItem due to skippable template")
				continue
			}

			if err != nil {
				logger.CtxErr(ctx, err).Info("failed to create listItem")
				el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "failed to create listItem"))

				continue
			}

			metrics.Bytes += itemInfo.SharePoint.Size

			itemPath, err := dc.FullPath().AppendItem(itemData.ID())
			if err != nil {
				logger.CtxErr(ctx, err).Info("failed to append item id to full path")
				el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "appending item to full path"))

				continue
			}

			err = deets.Add(
				itemPath,
				&path.Builder{}, // TODO: implement locationRef
				itemInfo)
			if err != nil {
				// Not critical enough to need to stop restore operation.
				logger.Ctx(ctx).Infow("accounting for restored item", "error", err)
			}

			metrics.Successes++
		}
	}

	return metrics, el.Failure()
}

// RestorePageCollection handles restoration of an individual site page collection.
// returns:
// - the collection's item and byte count metrics
// - the context cancellation station. True iff context is canceled.
func RestorePageCollection(
	ctx context.Context,
	gs graph.Servicer,
	dc data.RestoreCollection,
	restoreContainerName string,
	deets *details.Builder,
	errs *fault.Bus,
) (support.CollectionMetrics, error) {
	var (
		metrics   = support.CollectionMetrics{}
		directory = dc.FullPath()
		siteID    = directory.ProtectedResource()
	)

	trace.Log(ctx, "m365:sharepoint:restorePageCollection", directory.String())
	ctx, end := diagnostics.Span(ctx, "m365:sharepoint:restorePageCollection", diagnostics.Label("path", dc.FullPath()))

	defer end()

	var (
		el      = errs.Local()
		service = betaAPI.NewBetaService(gs.Adapter())
		items   = dc.Items(ctx, errs)
	)

	for {
		if el.Failure() != nil {
			break
		}

		select {
		case <-ctx.Done():
			return metrics, clues.StackWC(ctx, ctx.Err())

		case itemData, ok := <-items:
			if !ok {
				return metrics, nil
			}
			metrics.Objects++

			itemInfo, err := betaAPI.RestoreSitePage(
				ctx,
				service,
				itemData,
				siteID,
				restoreContainerName)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

			metrics.Bytes += itemInfo.SharePoint.Size

			itemPath, err := dc.FullPath().AppendItem(itemData.ID())
			if err != nil {
				el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "appending item to full path"))
				continue
			}

			err = deets.Add(
				itemPath,
				&path.Builder{}, // TODO: implement locationRef
				itemInfo)
			if err != nil {
				// Not critical enough to need to stop restore operation.
				logger.Ctx(ctx).Infow("accounting for restored item", "error", err)
			}

			metrics.Successes++
		}
	}

	return metrics, el.Failure()
}
