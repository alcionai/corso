package sharepoint

import (
	"context"
	"fmt"
	"io"
	"runtime/trace"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	discover "github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

//----------------------------------------------------------------------------
// SharePoint Restore WorkFlow:
// - RestoreCollections called by GC component
// -- Collections are iterated within, Control Flow Switch
// -- Switch:
// ---- Libraries restored via the same workflow as oneDrive
// ---- Lists call RestoreCollection()
// ----> for each data.Stream within  RestoreCollection.Items()
// ----> restoreListItems() is called
// Restored List can be found in the Site's `Site content` page
// Restored Libraries can be found within the Site's `Pages` page
//------------------------------------------

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	backupVersion int,
	creds account.M365Config,
	service graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
) (*support.ConnectorOperationStatus, error) {
	var (
		err            error
		restoreMetrics support.CollectionMetrics
	)

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		var (
			category = dc.FullPath().Category()
			metrics  support.CollectionMetrics
			ictx     = clues.Add(ctx,
				"category", category,
				"destination", dest.ContainerName, // TODO: pii
				"resource_owner", dc.FullPath().ResourceOwner()) // TODO: pii
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
			metrics, _, _, err = onedrive.RestoreCollection(
				ictx,
				backupVersion,
				service,
				dc,
				map[string]onedrive.Metadata{}, // Currently permission data is not stored for sharepoint
				onedrive.SharePointSource,
				dest.ContainerName,
				deets,
				map[string]string{},
				false,
				errs)
		case path.ListsCategory:
			metrics, err = RestoreListCollection(
				ictx,
				service,
				dc,
				dest.ContainerName,
				deets,
				errs)
		case path.PagesCategory:
			metrics, err = RestorePageCollection(
				ictx,
				creds,
				dc,
				dest.ContainerName,
				deets,
				errs)
		default:
			return nil, clues.Wrap(clues.New(category.String()), "category not supported")
		}

		restoreMetrics = support.CombineMetrics(restoreMetrics, metrics)

		if err != nil {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		restoreMetrics,
		dest.ContainerName)

	return status, err
}

// createRestoreFolders creates the restore folder hieararchy in the specified drive and returns the folder ID
// of the last folder entry given in the hiearchy
func createRestoreFolders(
	ctx context.Context,
	service graph.Servicer,
	siteID string,
	restoreFolders []string,
) (string, error) {
	// Get Main Drive for Site, Documents
	mainDrive, err := service.Client().SitesById(siteID).Drive().Get(ctx, nil)
	if err != nil {
		return "", graph.Wrap(ctx, err, "getting site drive root")
	}

	return onedrive.CreateRestoreFolders(ctx, service, ptr.Val(mainDrive.GetId()), restoreFolders)
}

// restoreListItem utility function restores a List to the siteID.
// The name is changed to to Corso_Restore_{timeStame}_name
// API Reference: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=http
// Restored List can be verified within the Site contents.
func restoreListItem(
	ctx context.Context,
	service graph.Servicer,
	itemData data.Stream,
	siteID, destName string,
) (details.ItemInfo, error) {
	ctx, end := D.Span(ctx, "gc:sharepoint:restoreList", D.Label("item_uuid", itemData.UUID()))
	defer end()

	ctx = clues.Add(ctx, "list_item_id", itemData.UUID())

	var (
		dii      = details.ItemInfo{}
		listName = itemData.UUID()
	)

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, clues.Wrap(err, "reading backup data").WithClues(ctx)
	}

	oldList, err := support.CreateListFromBytes(byteArray)
	if err != nil {
		return dii, clues.Wrap(err, "creating item").WithClues(ctx)
	}

	if name, ok := ptr.ValOK(oldList.GetDisplayName()); ok {
		listName = name
	}

	var (
		newName  = fmt.Sprintf("%s_%s", destName, listName)
		newList  = support.ToListable(oldList, newName)
		contents = make([]models.ListItemable, 0)
	)

	for _, itm := range oldList.GetItems() {
		temp := support.CloneListItem(itm)
		contents = append(contents, temp)
	}

	newList.SetItems(contents)

	// Restore to List base to M365 back store
	restoredList, err := service.Client().SitesById(siteID).Lists().Post(ctx, newList, nil)
	if err != nil {
		return dii, graph.Wrap(ctx, err, "restoring list")
	}

	// Uploading of ListItems is conducted after the List is restored
	// Reference: https://learn.microsoft.com/en-us/graph/api/listitem-create?view=graph-rest-1.0&tabs=http
	if len(contents) > 0 {
		for _, lItem := range contents {
			_, err := service.Client().
				SitesById(siteID).
				ListsById(ptr.Val(restoredList.GetId())).
				Items().
				Post(ctx, lItem, nil)
			if err != nil {
				return dii, graph.Wrap(ctx, err, "restoring list items").
					With("restored_list_id", ptr.Val(restoredList.GetId()))
			}
		}
	}

	dii.SharePoint = sharePointListInfo(restoredList, int64(len(byteArray)))

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
	ctx, end := D.Span(ctx, "gc:sharepoint:restoreListCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics   = support.CollectionMetrics{}
		directory = dc.FullPath()
		siteID    = directory.ResourceOwner()
		items     = dc.Items(ctx, errs)
		el        = errs.Local()
	)

	trace.Log(ctx, "gc:sharepoint:restoreListCollection", directory.String())

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
				el.AddRecoverable(err)
				continue
			}

			metrics.Bytes += itemInfo.SharePoint.Size

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				el.AddRecoverable(clues.Wrap(err, "appending item to full path").WithClues(ctx))
				continue
			}

			deets.Add(
				itemPath.String(),
				itemPath.ShortRef(),
				"",
				"", // TODO: implement locationRef
				true,
				itemInfo)

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
	creds account.M365Config,
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

	trace.Log(ctx, "gc:sharepoint:restorePageCollection", directory.String())
	ctx, end := D.Span(ctx, "gc:sharepoint:restorePageCollection", D.Label("path", dc.FullPath()))

	defer end()

	adpt, err := graph.CreateAdapter(creds.AzureTenantID, creds.AzureClientID, creds.AzureClientSecret)
	if err != nil {
		return metrics, clues.Wrap(err, "constructing graph client")
	}

	var (
		el      = errs.Local()
		service = discover.NewBetaService(adpt)
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

			itemInfo, err := api.RestoreSitePage(
				ctx,
				service,
				itemData,
				siteID,
				restoreContainerName)
			if err != nil {
				el.AddRecoverable(err)
				continue
			}

			metrics.Bytes += itemInfo.SharePoint.Size

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				el.AddRecoverable(clues.Wrap(err, "appending item to full path").WithClues(ctx))
				continue
			}

			deets.Add(
				itemPath.String(),
				itemPath.ShortRef(),
				"",
				"", // TODO: implement locationRef
				true,
				itemInfo)

			metrics.Successes++
		}
	}

	return metrics, el.Failure()
}
