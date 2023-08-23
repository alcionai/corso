package site

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/trace"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
		lrh            = drive.NewLibraryRestoreHandler(ac, rcc.Selector.PathService())
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
				"resource_owner", clues.Hide(dc.FullPath().ResourceOwner()),
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
				ac.Stable,
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
	service graph.Servicer,
	itemData data.Item,
	siteID, destName string,
) (details.ItemInfo, error) {
	ctx, end := diagnostics.Span(ctx, "m365:sharepoint:restoreList", diagnostics.Label("item_uuid", itemData.ID()))
	defer end()

	ctx = clues.Add(ctx, "list_item_id", itemData.ID())

	var (
		dii      = details.ItemInfo{}
		listName = itemData.ID()
	)

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, clues.Wrap(err, "reading backup data").WithClues(ctx)
	}

	oldList, err := betaAPI.CreateListFromBytes(byteArray)
	if err != nil {
		return dii, clues.Wrap(err, "creating item").WithClues(ctx)
	}

	if name, ok := ptr.ValOK(oldList.GetDisplayName()); ok {
		listName = name
	}

	var (
		newName  = fmt.Sprintf("%s_%s", destName, listName)
		newList  = betaAPI.ToListable(oldList, newName)
		contents = make([]models.ListItemable, 0)
	)

	for _, itm := range oldList.GetItems() {
		temp := betaAPI.CloneListItem(itm)
		contents = append(contents, temp)
	}

	newList.SetItems(contents)

	// Restore to List base to M365 back store
	restoredList, err := service.Client().Sites().BySiteId(siteID).Lists().Post(ctx, newList, nil)
	if err != nil {
		return dii, graph.Wrap(ctx, err, "restoring list")
	}

	// Uploading of ListItems is conducted after the List is restored
	// Reference: https://learn.microsoft.com/en-us/graph/api/listitem-create?view=graph-rest-1.0&tabs=http
	if len(contents) > 0 {
		for _, lItem := range contents {
			_, err := service.Client().
				Sites().
				BySiteId(siteID).
				Lists().
				ByListId(ptr.Val(restoredList.GetId())).
				Items().
				Post(ctx, lItem, nil)
			if err != nil {
				return dii, graph.Wrap(ctx, err, "restoring list items").
					With("restored_list_id", ptr.Val(restoredList.GetId()))
			}
		}
	}

	dii.SharePoint = ListToSPInfo(restoredList, int64(len(byteArray)))

	return dii, nil
}

func RestoreListCollection(
	ctx context.Context,
	service graph.Servicer,
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
		siteID    = directory.ResourceOwner()
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
			return metrics, clues.Stack(ctx.Err()).WithClues(ctx)

		case itemData, ok := <-items:
			if !ok {
				return metrics, nil
			}
			metrics.Objects++

			itemInfo, err := restoreListItem(
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
				el.AddRecoverable(ctx, clues.Wrap(err, "appending item to full path").WithClues(ctx))
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
		siteID    = directory.ResourceOwner()
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
			return metrics, clues.Stack(ctx.Err()).WithClues(ctx)

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
				el.AddRecoverable(ctx, clues.Wrap(err, "appending item to full path").WithClues(ctx))
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
