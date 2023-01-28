package sharepoint

import (
	"context"
	"fmt"
	"io"
	"runtime/trace"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

//----------------------------------------------------------------------------
// SharePoint Restore WorkFlow:
// - RestoreCollections called by GC component
// -- Collections are iterated within, Control Flow Switch
// -- Switch:
// ---- Libraries restored via the same workflow as oneDrive
// ---- Lists call RestoreCollection()
// ----> for each data.Stream within  Collection.Items()
// ----> restoreListItems() is called
// Restored List can be found in the Site's `Site content` page
// Restored Libraries can be found within the Site's `Pages` page
//------------------------------------------

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
		case path.PagesCategory:
			errorMessage := fmt.Sprintf("restore of %s not supported", dc.FullPath().Category())
			logger.Ctx(ctx).Error(errorMessage)

			return nil, errors.New(errorMessage)

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

	var (
		dii      = details.ItemInfo{}
		listName = itemData.UUID()
	)

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, errors.Wrap(err, "sharepoint restoreItem failed to retrieve bytes from data.Stream")
	}
	// Create Item
	oldList, err := support.CreateListFromBytes(byteArray)
	if err != nil {
		return dii, errors.Wrapf(err, "failed to build list item %s", listName)
	}

	if oldList.GetDisplayName() != nil {
		listName = *oldList.GetDisplayName()
	}

	newName := fmt.Sprintf("%s_%s", destName, listName)
	newList := support.ToListable(oldList, newName)

	contents := make([]models.ListItemable, 0)

	for _, itm := range oldList.GetItems() {
		temp := support.CloneListItem(itm)
		contents = append(contents, temp)
	}

	newList.SetItems(contents)

	// Restore to List base to M365 back store
	restoredList, err := service.Client().SitesById(siteID).Lists().Post(ctx, newList, nil)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"failure to create list foundation ID: %s API Error Details: %s",
			itemData.UUID(),
			support.ConnectorStackErrorTrace(err),
		)

		return dii, errors.Wrap(err, errorMsg)
	}

	// Uploading of ListItems is conducted after the List is restored
	// Reference: https://learn.microsoft.com/en-us/graph/api/listitem-create?view=graph-rest-1.0&tabs=http
	if len(contents) > 0 {
		for _, lItem := range contents {
			_, err := service.Client().
				SitesById(siteID).
				ListsById(*restoredList.GetId()).
				Items().
				Post(ctx, lItem, nil)
			if err != nil {
				errorMsg := fmt.Sprintf(
					"listItem  failed for listID %s. Details: %s. Content: %v",
					*restoredList.GetId(),
					support.ConnectorStackErrorTrace(err),
					lItem.GetAdditionalData(),
				)

				return dii, errors.Wrap(err, errorMsg)
			}
		}
	}

	dii.SharePoint = sharePointListInfo(restoredList, int64(len(byteArray)))

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

			itemInfo, err := restoreListItem(
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

func restoreSitePage(
	ctx context.Context,
	service *betasdk.Service,
	itemData data.Stream,
	siteID, destName string,
) (details.ItemInfo, error) {
	ctx, end := D.Span(ctx, "gc:sharepoint:restorePage", D.Label("item_uuid", itemData.UUID()))
	defer end()

	var (
		dii      = details.ItemInfo{}
		pageID   = itemData.UUID()
		pageName = pageID
	)

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, errors.Wrap(err, "sharepoint restorePage failed to restore bytes from data.Stream")
	}

	// Hydrate Page
	page, err := support.CreatePageFromBytes(byteArray)
	if err != nil {
		return dii, errors.Wrapf(err, "failed to create Page object %s", pageID)
	}

	pageNamePtr := page.GetName()
	if pageNamePtr != nil {
		pageName = *pageNamePtr
	}

	newName := fmt.Sprintf("%s_%s", destName, pageName)
	page.SetName(&newName)

	// Restore is a 2-Step Process in Graph API
	// 1. Create the Page on the site
	// 2. Publish the site
	// See: https://learn.microsoft.com/en-us/graph/api/sitepage-create?view=graph-rest-beta
	restoredPage, err := service.Client().SitesById(siteID).Pages().Post(ctx, page, nil)
	if err != nil {
		sendErr := support.ConnectorStackErrorTraceWrap(
			err,
			"failure to create page from ID: %s"+pageName+" API Error Details",
		)

		return dii, sendErr
	}

	// Publish page to make visible
	// See https://learn.microsoft.com/en-us/graph/api/sitepage-publish?view=graph-rest-beta
	if restoredPage.GetWebUrl() == nil {
		return dii, fmt.Errorf("created page %s did not return webURL from API", *restoredPage.GetId())
	}

	pageID := *restoredPage.GetId()

	err = service.Client().SitesById(siteID).PagesById(pageID).Publish().Post(ctx, nil)
	if err != nil {
		return dii, support.ConnectorStackErrorTraceWrap(
			err,
			"unable to publish page ID: "+pageID+" API Error Details",
		)
	}

	dii.SharePoint = sharePointPageInfo(restoredPage, int64(len(byteArray)))
	// Storing new pageID in unused field.
	dii.SharePoint.ParentPath = pageID

	return dii, nil
}
