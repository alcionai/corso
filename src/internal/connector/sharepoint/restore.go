package sharepoint

import (
	"context"
	"fmt"
	"io"
	"runtime/trace"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	service graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.Collection,
	deets *details.Builder,
) (*support.ConnectorOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		restoreErrors  error
	)

	errUpdater := func(id string, err error) {
		restoreErrors = support.WrapAndAppend(id, err, restoreErrors)
	}

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		var (
			metrics  support.CollectionMetrics
			canceled bool
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
			metrics, canceled = onedrive.RestoreCollection(
				ctx,
				service,
				dc,
				onedrive.OneDriveSource,
				dest.ContainerName,
				deets,
				errUpdater)
		case path.ListsCategory:
			metrics, canceled = RestoreCollection(
				ctx,
				service,
				dc,
				dest.ContainerName,
				deets,
				errUpdater,
			)
		default:
			return nil, errors.Errorf("category %s not supported", dc.FullPath().Category())
		}

		restoreMetrics.Combine(metrics)

		if canceled {
			break
		}
	}

	return support.CreateStatus(
			ctx,
			support.Restore,
			len(dcs),
			restoreMetrics,
			restoreErrors,
			dest.ContainerName),
		nil
}

// createRestoreFolders creates the restore folder hieararchy in the specified drive and returns the folder ID
// of the last folder entry given in the hiearchy
func createRestoreFolders(ctx context.Context, service graph.Servicer, siteID string, restoreFolders []string,
) (string, error) {
	// Get Main Drive for Site, Documents
	mainDrive, err := service.Client().SitesById(siteID).Drive().Get(ctx, nil)
	if err != nil {
		return "", errors.Wrapf(
			err,
			"failed to get site drive root. details: %s",
			support.ConnectorStackErrorTrace(err),
		)
	}

	return onedrive.CreateRestoreFolders(ctx, service, *mainDrive.GetId(), restoreFolders)
}

// restoreItem utility function restores a List to the siteID. The name is changed to to Corso_Restore_{timeStame}_name
// API Reference: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=http
// Restored List can be verified within the Site contents
func restoreItem(
	ctx context.Context,
	service graph.Servicer,
	itemData data.Stream,
	siteID, destName string,
) (details.ItemInfo, error) {
	ctx, end := D.Span(ctx, "gc:sharepoint:restoreList", D.Label("item_uuid", itemData.UUID()))
	defer end()

	var (
		dii         = details.ItemInfo{}
		itemName    = itemData.UUID()
		displayName = itemName
	)

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, errors.Wrap(err, "sharepoint restoreItem failed to retrieve bytes from data.Stream")
	}
	// Create Item
	newItem, err := support.CreateListFromBytes(byteArray)
	if err != nil {
		return dii, errors.Wrapf(err, "failed to construct list item %s", itemName)
	}

	// If field "name" is set, this will trigger the following error:
	// invalidRequest Cannot define a 'name' for a list as it is assigned by the server. Instead, provide 'displayName'
	if newItem.GetName() != nil {
		adtlData := newItem.GetAdditionalData()
		adtlData["list_name"] = *newItem.GetName()
		newItem.SetName(nil)
	}

	if newItem.GetDisplayName() != nil {
		displayName = *newItem.GetDisplayName()
	}

	newName := fmt.Sprintf("%s_%s", destName, displayName)
	newItem.SetDisplayName(&newName)

	// Restore to M365 store
	restoredList, err := service.Client().SitesById(siteID).Lists().Post(ctx, newItem, nil)
	if err != nil {
		return dii, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	written := int64(len(byteArray))

	dii.SharePoint = sharePointListInfo(restoredList, written)

	return dii, nil
}

func RestoreCollection(
	ctx context.Context,
	service graph.Servicer,
	dc data.Collection,
	restoreContainerName string,
	deets *details.Builder,
	errUpdater func(string, error),
) (support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:sharepoint:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics   = support.CollectionMetrics{}
		directory = dc.FullPath()
	)

	trace.Log(ctx, "gc:sharepoint:restoreCollection", directory.String())
	siteID := directory.ResourceOwner()

	// Restore items from the collection
	items := dc.Items()

	for {
		select {
		case <-ctx.Done():
			errUpdater("context canceled", ctx.Err())
			return metrics, true

		case itemData, ok := <-items:
			if !ok {
				return metrics, false
			}
			metrics.Objects++

			itemInfo, err := restoreItem(
				ctx,
				service,
				itemData,
				siteID,
				restoreContainerName,
			)
			if err != nil {
				errUpdater(itemData.UUID(), err)
				continue
			}

			metrics.TotalBytes += itemInfo.SharePoint.Size

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				logger.Ctx(ctx).DPanicw("transforming item to full path", "error", err)
				errUpdater(itemData.UUID(), err)

				continue
			}

			deets.Add(
				itemPath.String(),
				itemPath.ShortRef(),
				"",
				true,
				itemInfo)

			metrics.Successes++
		}
	}
}
