package site

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/trace"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

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
	var (
		dii    = details.ItemInfo{}
		itemID = itemData.ID()
	)

	ctx, end := diagnostics.Span(ctx, "m365:sharepoint:restoreList", diagnostics.Label("item_uuid", itemData.ID()))
	defer end()

	ctx = clues.Add(ctx, "list_item_id", itemID)

	bytes, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, clues.WrapWC(ctx, err, "reading backup data")
	}

	storedList, err := api.BytesToListable(bytes)
	if err != nil {
		return dii, clues.WrapWC(ctx, err, "generating list from stored bytes")
	}

	newName := formatListsRestoreDestination(destName, itemID, storedList)

	// Restore to List base to M365 back store
	restoredList, err := rh.PostList(ctx, newName, storedList, errs)
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

// newName is of format: destinationName_listID
// here we replace listID with displayName of list generated from stored list
func formatListsRestoreDestination(destName, itemID string, storedList models.Listable) string {
	part1 := destName
	part2 := itemID

	if dispName, ok := ptr.ValOK(storedList.GetDisplayName()); ok {
		part2 = dispName
	}

	return fmt.Sprintf("%s_%s", part1, part2)
}
