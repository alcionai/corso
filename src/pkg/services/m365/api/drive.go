package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Drives() Drives {
	return Drives{c}
}

// Drives is an interface-compliant provider of the client.
type Drives struct {
	Client
}

// ---------------------------------------------------------------------------
// Folders
// ---------------------------------------------------------------------------

const itemByPathRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"

var ErrFolderNotFound = clues.New("folder not found")

// GetFolderByName will lookup the specified folder by name within the parentFolderID folder.
func (c Drives) GetFolderByName(
	ctx context.Context,
	driveID, parentFolderID, folderID string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folderID)
	builder := drives.NewItemItemsDriveItemItemRequestBuilder(rawURL, c.Stable.Adapter())

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

func (c Drives) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	root, err := c.Stable.
		Client().
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
// Items
// ---------------------------------------------------------------------------

// generic drive item getter
func (c Drives) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	di, err := c.Stable.
		Client().
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

func (c Drives) PostItem(
	ctx context.Context,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	session := drives.NewItemItemsItemCreateUploadSessionPostRequestBody()

	r, err := c.Stable.
		Client().
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

func (c Drives) PatchItem(
	ctx context.Context,
	driveID, itemID string,
	item models.DriveItemable,
) error {
	_, err := c.Stable.
		Client().
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

func (c Drives) PutItemContent(
	ctx context.Context,
	driveID, itemID string,
	content []byte,
) error {
	_, err := c.Stable.
		Client().
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
func (c Drives) DeleteItem(
	ctx context.Context,
	driveID, itemID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := c.Service()
	if err != nil {
		return graph.Wrap(ctx, err, "creating adapter to delete item permission")
	}

	err = srv.
		Client().
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

// ---------------------------------------------------------------------------
// Permissions
// ---------------------------------------------------------------------------

func (c Drives) GetItemPermission(
	ctx context.Context,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	perm, err := c.Stable.
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

func (c Drives) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	ctx = graph.ConsumeNTokens(ctx, graph.PermissionsLC)

	itm, err := c.Stable.
		Client().
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

func (c Drives) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := c.Service()
	if err != nil {
		return graph.Wrap(ctx, err, "creating adapter to delete item permission")
	}

	err = srv.
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
