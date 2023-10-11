package drive

import (
	"context"
	"net/http"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// backup
// ---------------------------------------------------------------------------

type baseItemHandler struct {
	ac api.Drives
}

func (h baseItemHandler) NewDrivePager(
	resourceOwner string, fields []string,
) api.Pager[models.Driveable] {
	return h.ac.NewUserDrivePager(resourceOwner, fields)
}

// AugmentItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func (h baseItemHandler) AugmentItemInfo(
	dii details.ItemInfo,
	resource idname.Provider,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	var pps string

	if parentPath != nil {
		pps = parentPath.String()
	}

	driveName, driveID := getItemDriveInfo(item)

	dii.Extension = &details.ExtensionData{}
	dii.OneDrive = &details.OneDriveInfo{
		Created:    ptr.Val(item.GetCreatedDateTime()),
		DriveID:    driveID,
		DriveName:  driveName,
		ItemName:   ptr.Val(item.GetName()),
		ItemType:   details.OneDriveItem,
		Modified:   ptr.Val(item.GetLastModifiedDateTime()),
		Owner:      getItemCreator(item),
		ParentPath: pps,
		Size:       size,
	}

	return dii
}

var _ BackupHandler = &itemBackupHandler{}

type itemBackupHandler struct {
	baseItemHandler
	userID string
	scope  selectors.OneDriveScope
}

func NewItemBackupHandler(ac api.Drives, userID string, scope selectors.OneDriveScope) *itemBackupHandler {
	return &itemBackupHandler{
		baseItemHandler: baseItemHandler{
			ac: ac,
		},
		userID: userID,
		scope:  scope,
	}
}

func (h itemBackupHandler) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return h.ac.Get(ctx, url, headers)
}

func (h itemBackupHandler) PathPrefix(
	tenantID, driveID string,
) (path.Path, error) {
	return path.Build(
		tenantID,
		h.userID,
		path.OneDriveService,
		path.FilesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h itemBackupHandler) MetadataPathPrefix(
	tenantID string,
) (path.Path, error) {
	p, err := path.BuildMetadata(
		tenantID,
		h.userID,
		path.OneDriveService,
		path.FilesCategory,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path")
	}

	return p, nil
}

func (h itemBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.ToDataLayerOneDrivePath(tenantID, h.userID, false)
}

func (h itemBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.OneDriveService, path.FilesCategory
}

func (h itemBackupHandler) FormatDisplayPath(
	_ string, // drive name not displayed for onedrive
	pb *path.Builder,
) string {
	return "/" + pb.String()
}

func (h itemBackupHandler) NewLocationIDer(
	driveID string,
	elems ...string,
) details.LocationIDer {
	return details.NewOneDriveLocationIDer(driveID, elems...)
}

func (h itemBackupHandler) GetItemPermission(
	ctx context.Context,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	return h.ac.GetItemPermission(ctx, driveID, itemID)
}

func (h itemBackupHandler) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	return h.ac.GetItem(ctx, driveID, itemID)
}

func (h itemBackupHandler) IsAllPass() bool {
	return h.scope.IsAny(selectors.OneDriveFolder)
}

func (h itemBackupHandler) IncludesDir(dir string) bool {
	return h.scope.Matches(selectors.OneDriveFolder, dir)
}

func (h itemBackupHandler) EnumerateDriveItemsDelta(
	ctx context.Context,
	driveID, prevDeltaLink string,
	selectProps []string,
) ([]models.DriveItemable, api.DeltaUpdate, error) {
	return h.ac.EnumerateDriveItemsDelta(ctx, driveID, prevDeltaLink, selectProps)
}

// ---------------------------------------------------------------------------
// Restore
// ---------------------------------------------------------------------------

var _ RestoreHandler = &itemRestoreHandler{}

type itemRestoreHandler struct {
	baseItemHandler
}

func NewRestoreHandler(ac api.Client) *itemRestoreHandler {
	return &itemRestoreHandler{
		baseItemHandler: baseItemHandler{
			ac: ac.Drives(),
		},
	}
}

func (h itemRestoreHandler) PostDrive(
	context.Context,
	string, string,
) (models.Driveable, error) {
	return nil, clues.New("creating drives in oneDrive is not supported")
}

func (h itemRestoreHandler) DeleteItem(
	ctx context.Context,
	driveID, itemID string,
) error {
	return h.ac.DeleteItem(ctx, driveID, itemID)
}

func (h itemRestoreHandler) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return h.ac.DeleteItemPermission(ctx, driveID, itemID, permissionID)
}

func (h itemRestoreHandler) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]api.DriveItemIDType, error) {
	m, err := h.ac.GetItemsInContainerByCollisionKey(ctx, driveID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (h itemRestoreHandler) NewItemContentUpload(
	ctx context.Context,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	return h.ac.NewItemContentUpload(ctx, driveID, itemID)
}

func (h itemRestoreHandler) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return h.ac.PostItemPermissionUpdate(ctx, driveID, itemID, body)
}

func (h itemRestoreHandler) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	return h.ac.PostItemLinkShareUpdate(ctx, driveID, itemID, body)
}

func (h itemRestoreHandler) PostItemInContainer(
	ctx context.Context,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
	onCollision control.CollisionPolicy,
) (models.DriveItemable, error) {
	return h.ac.PostItemInContainer(ctx, driveID, parentFolderID, newItem, onCollision)
}

func (h itemRestoreHandler) GetFolderByName(
	ctx context.Context,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	return h.ac.GetFolderByName(ctx, driveID, parentFolderID, folderName)
}

func (h itemRestoreHandler) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	return h.ac.GetRootFolder(ctx, driveID)
}
