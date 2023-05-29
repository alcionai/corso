package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

// ---------------------------------------------------------------------------
// Drives
// ---------------------------------------------------------------------------

func GetUsersDrive(
	ctx context.Context,
	srv graph.Servicer,
	user string,
) (models.Driveable, error) {
	d, err := srv.Client().
		Users().
		ByUserId(user).
		Drive().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting user's drive")
	}

	return d, nil
}

func GetSitesDefaultDrive(
	ctx context.Context,
	srv graph.Servicer,
	site string,
) (models.Driveable, error) {
	d, err := srv.Client().
		Sites().
		BySiteId(site).
		Drive().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting site's drive")
	}

	return d, nil
}

func GetDriveRoot(
	ctx context.Context,
	srv graph.Servicer,
	driveID string,
) (models.DriveItemable, error) {
	root, err := srv.Client().
		Drives().
		ByDriveId(driveID).
		Root().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting drive root")
	}

	return root, nil
}

// ---------------------------------------------------------------------------
// Drive Items
// ---------------------------------------------------------------------------

// generic drive item getter
func GetDriveItem(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.DriveItemable, error) {
	di, err := srv.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting item")
	}

	return di, nil
}

func PostDriveItem(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	session := drives.NewItemItemsItemCreateUploadSessionPostRequestBody()

	r, err := srv.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		CreateUploadSession().
		Post(ctx, session, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "uploading drive item")
	}

	return r, nil
}

func PatchDriveItem(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
	item models.DriveItemable,
) error {
	_, err := srv.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Patch(ctx, item, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "patching drive item")
	}

	return nil
}

func PutDriveItemContent(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
	content []byte,
) error {
	_, err := srv.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Content().
		Put(ctx, content, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "uploading drive item content")
	}

	return nil
}

// deletes require unique http clients
// https://github.com/alcionai/corso/issues/2707
func DeleteDriveItem(
	ctx context.Context,
	gs graph.Servicer,
	driveID, itemID string,
) error {
	err := gs.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting item").With("item_id", itemID)
	}

	return nil
}

const itemByPathRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"

var ErrFolderNotFound = clues.New("folder not found")

// GetFolderByName will lookup the specified folder by name within the parentFolderID folder.
func GetFolderByName(
	ctx context.Context,
	srv graph.Servicer,
	driveID, parentFolderID, folder string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folder)
	builder := drives.NewItemItemsDriveItemItemRequestBuilder(rawURL, srv.Adapter())

	foundItem, err := builder.Get(ctx, nil)
	if err != nil {
		if graph.IsErrDeletedInFlight(err) {
			return nil, graph.Stack(ctx, clues.Stack(ErrFolderNotFound, err))
		}

		return nil, graph.Wrap(ctx, err, "getting folder")
	}

	// Check if the item found is a folder, fail the call if not
	if foundItem.GetFolder() == nil {
		return nil, graph.Wrap(ctx, ErrFolderNotFound, "item is not a folder")
	}

	return foundItem, nil
}

// ---------------------------------------------------------------------------
// Permissions
// ---------------------------------------------------------------------------

func GetItemPermission(
	ctx context.Context,
	service graph.Servicer,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	perm, err := service.
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting item metadata").With("item_id", itemID)
	}

	return perm, nil
}

func PostItemPermissionUpdate(
	ctx context.Context,
	service graph.Servicer,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	ctx = graph.ConsumeNTokens(ctx, graph.PermissionsLC)

	itm, err := service.Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Invite().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "posting permissions")
	}

	return itm, nil
}

func DeleteDriveItemPermission(
	ctx context.Context,
	creds account.M365Config,
	driveID, itemID, permissionID string,
) error {
	a, err := graph.CreateAdapter(creds.AzureTenantID, creds.AzureClientID, creds.AzureClientSecret)
	if err != nil {
		return graph.Wrap(ctx, err, "creating adapter to delete item permission")
	}

	err = graph.NewService(a).
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Permissions().
		ByPermissionId(permissionID).
		Delete(graph.ConsumeNTokens(ctx, graph.PermissionsLC), nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting drive item permission")
	}

	return nil
}
