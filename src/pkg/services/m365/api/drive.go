package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/errs/core"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
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

const (
	itemByPathRawURLFmt   = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"
	createLinkShareURLFmt = "https://graph.microsoft.com/beta/drives/%s/items/%s/createLink"
)

var ErrFolderNotFound = clues.New("folder not found")

// GetFolderByName will lookup the specified folder by name within the parentFolderID folder.
func (c Drives) GetFolderByName(
	ctx context.Context,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folderName)
	builder := drives.NewItemItemsDriveItemItemRequestBuilder(rawURL, c.Stable.Adapter())

	foundItem, err := builder.Get(ctx, nil)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			err = clues.Stack(ErrFolderNotFound, err)
		}

		return nil, clues.Wrap(err, "getting folder")
	}

	// Check if the item found is a folder, fail the call if not
	if foundItem.GetFolder() == nil {
		return nil, clues.WrapWC(ctx, ErrFolderNotFound, "item is not a folder")
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
		return nil, clues.Wrap(err, "getting drive root")
	}

	return root, nil
}

// TODO: pagination controller needed for completion.
func (c Drives) GetFolderChildren(
	ctx context.Context,
	driveID, folderID string,
) ([]models.DriveItemable, error) {
	response, err := c.Stable.
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(folderID).
		Children().
		Get(ctx, nil)
	if err != nil {
		return nil, clues.Wrap(err, "getting folder children")
	}

	return response.GetValue(), nil
}

// ---------------------------------------------------------------------------
// Items
// ---------------------------------------------------------------------------

// generic drive item getter
func (c Drives) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	options := &drives.ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &drives.ItemItemsDriveItemItemRequestBuilderGetQueryParameters{
			// FIXME: accept a CallConfig instead of hardcoding the select.
			Select: DefaultDriveItemProps(),
		},
	}

	di, err := c.Stable.
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Get(ctx, options)

	return di, clues.Wrap(err, "getting item").OrNil()
}

func (c Drives) NewItemContentUpload(
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

	return r, clues.Wrap(err, "uploading drive item").OrNil()
}

//nolint:lll
const itemChildrenRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s/children?@microsoft.graph.conflictBehavior=%s"

const (
	conflictBehaviorFail    = "fail"
	conflictBehaviorRename  = "rename"
	conflictBehaviorReplace = "replace"
)

// PostItemInContainer creates a new item in the specified folder
func (c Drives) PostItemInContainer(
	ctx context.Context,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
	onCollision control.CollisionPolicy,
) (models.DriveItemable, error) {
	// graph api has no policy for Skip; instead we wrap the same-name failure
	// as a graph.ErrItemAlreadyExistsConflict.
	conflictBehavior := conflictBehaviorFail

	switch onCollision {
	case control.Replace:
		conflictBehavior = conflictBehaviorReplace
	case control.Copy:
		conflictBehavior = conflictBehaviorRename
	}

	// Graph SDK doesn't yet provide a POST method for `/children` so we set the `rawUrl` ourselves as recommended
	// here: https://github.com/microsoftgraph/msgraph-sdk-go/issues/155#issuecomment-1136254310
	rawURL := fmt.Sprintf(itemChildrenRawURLFmt, driveID, parentFolderID, conflictBehavior)
	builder := drives.NewItemItemsRequestBuilder(rawURL, c.Stable.Adapter())

	newItem, err := builder.Post(ctx, newItem, nil)
	if err != nil {
		return nil, clues.Wrap(err, "creating item in folder")
	}

	return newItem, nil
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

	return clues.Wrap(err, "patching drive item").OrNil()
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

	return clues.Wrap(err, "uploading drive item content").OrNil()
}

// deletes require unique http clients
// https://github.com/alcionai/canario/issues/2707
func (c Drives) DeleteItem(
	ctx context.Context,
	driveID, itemID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/canario/issues/2707
	srv, err := c.Service(c.counter)
	if err != nil {
		return clues.WrapWC(ctx, err, "creating adapter to delete item permission")
	}

	err = srv.
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Delete(ctx, nil)

	return clues.Wrap(err, "deleting item").With("item_id", itemID).OrNil()
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

	return perm, clues.Wrap(err, "getting item permissions").OrNil()
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

	return itm, clues.Wrap(err, "posting permissions").OrNil()
}

func (c Drives) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/canario/issues/2707
	srv, err := c.Service(c.counter)
	if err != nil {
		return clues.WrapWC(ctx, err, "creating adapter to delete item permission")
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

	return clues.Wrap(err, "deleting drive item permission").OrNil()
}

func (c Drives) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	ctx = graph.ConsumeNTokens(ctx, graph.PermissionsLC)

	// We are using the beta version of the endpoint. This allows us
	// to add recipients in the same request as well as to make it not
	// send out and email for every link share the user gets added to.
	rawURL := fmt.Sprintf(createLinkShareURLFmt, driveID, itemID)
	builder := drives.NewItemItemsItemCreateLinkRequestBuilder(rawURL, c.Stable.Adapter())

	itm, err := builder.Post(ctx, body, nil)

	return itm, clues.Wrap(err, "creating link share").OrNil()
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// DriveItemCollisionKeyy constructs a key from the item name.
// collision keys are used to identify duplicate item conflicts for handling advanced restoration config.
func DriveItemCollisionKey(item models.DriveItemable) string {
	if item == nil {
		return ""
	}

	return ptr.Val(item.GetName())
}

// NewDriveItem initializes a `models.DriveItemable` with either a folder or file entry.
func NewDriveItem(name string, folder bool) *models.DriveItem {
	itemToCreate := models.NewDriveItem()
	itemToCreate.SetName(&name)

	if folder {
		itemToCreate.SetFolder(models.NewFolder())
	} else {
		itemToCreate.SetFile(models.NewFile())
	}

	return itemToCreate
}
