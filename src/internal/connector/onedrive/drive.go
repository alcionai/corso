package onedrive

import (
	"context"
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/items"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/items/item"
	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/delta"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/logger"
)

var errFolderNotFound = errors.New("folder not found")

const (
	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey           = "@odata.nextLink"
	itemChildrenRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children"
	itemByPathRawURLFmt   = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"
	itemNotFoundErrorCode = "itemNotFound"
)

// Enumerates the drives for the specified user
func drives(ctx context.Context, service graph.Service, user string) ([]models.Driveable, error) {
	r, err := service.Client().UsersById(user).Drives().Get(ctx, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve user drives. user: %s, details: %s",
			user, support.ConnectorStackErrorTrace(err))
	}

	logger.Ctx(ctx).Debugf("Found %d drives for user %s", len(r.GetValue()), user)

	return r.GetValue(), nil
}

// itemCollector functions collect the items found in a drive
type itemCollector func(ctx context.Context, driveID string, driveItems []models.DriveItemable) error

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	service graph.Service,
	driveID string,
	collector itemCollector,
) error {
	// TODO: Specify a timestamp in the delta query
	// https://docs.microsoft.com/en-us/graph/api/driveitem-delta?
	// view=graph-rest-1.0&tabs=http#example-4-retrieving-delta-results-using-a-timestamp
	builder := service.Client().DrivesById(driveID).Root().Delta()

	for {
		r, err := builder.Get(ctx, nil)
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to query drive items. details: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		err = collector(ctx, driveID, r.GetValue())
		if err != nil {
			return err
		}

		// Check if there are more items
		if _, found := r.GetAdditionalData()[nextLinkKey]; !found {
			break
		}

		nextLink := r.GetAdditionalData()[nextLinkKey].(*string)
		logger.Ctx(ctx).Debugf("Found %s nextLink", *nextLink)
		builder = delta.NewDeltaRequestBuilder(*nextLink, service.Adapter())
	}

	return nil
}

// getFolder will lookup the specified folder name under `parentFolderID`
func getFolder(ctx context.Context, service graph.Service, driveID string, parentFolderID string,
	folderName string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folderName)

	builder := item.NewDriveItemItemRequestBuilder(rawURL, service.Adapter())

	item, err := builder.Get(ctx, nil)
	if err != nil {

		if oDataError := support.ODataError(err); oDataError != nil && oDataError.GetError() != nil &&
			oDataError.GetError().GetCode() != nil &&
			*oDataError.GetError().GetCode() == itemNotFoundErrorCode {

			return nil, errors.WithStack(errFolderNotFound)
		}

		return nil, errors.Wrapf(err,
			"failed to get folder %s/%s. details: %s",
			parentFolderID,
			folderName,
			support.ConnectorStackErrorTrace(err),
		)
	}

	// Check if the item found is a folder, fail the call if not
	if item.GetFolder() == nil {
		return nil, errors.WithStack(errFolderNotFound)
	}

	return item, nil
}

// Create a new item in the specified folder
func createItem(ctx context.Context, service graph.Service, driveID string, parentFolderID string,
	item models.DriveItemable,
) (models.DriveItemable, error) {
	// Graph SDK doesn't yet provide a POST method for `/children` so we set the `rawUrl` ourselves as recommended
	// here: https://github.com/microsoftgraph/msgraph-sdk-go/issues/155#issuecomment-1136254310
	rawURL := fmt.Sprintf(itemChildrenRawURLFmt, driveID, parentFolderID)

	builder := items.NewItemsRequestBuilder(rawURL, service.Adapter())

	newItem, err := builder.Post(ctx, item, nil)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to create item. details: %s",
			support.ConnectorStackErrorTrace(err),
		)
	}

	return newItem, nil
}

// newItem initializes a `models.DriveItemable` that can be used as input to `createItem`
func newItem(name string, folder bool) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(&name)

	if folder {
		item.SetFolder(models.NewFolder())
	} else {
		item.SetFile(models.NewFile())
	}

	return item
}
