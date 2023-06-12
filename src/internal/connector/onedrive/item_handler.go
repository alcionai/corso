package onedrive

import (
	"context"
	"net/http"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// backup
// ---------------------------------------------------------------------------

var _ BackupHandler = &itemBackupHandler{}

type itemBackupHandler struct {
	ac api.Drives
}

func (h itemBackupHandler) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return h.ac.Get(ctx, url, headers)
}

func (h itemBackupHandler) PathPrefix(
	tenantID, resourceOwner, driveID string,
) (path.Path, error) {
	return path.Build(
		tenantID,
		resourceOwner,
		path.OneDriveService,
		path.FilesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h itemBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID, resourceOwner string,
) (path.Path, error) {
	return folders.ToDataLayerOneDrivePath(tenantID, resourceOwner, false)
}

func (h itemBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.OneDriveService, path.FilesCategory
}

func (h itemBackupHandler) NewDrivePager(
	resourceOwner string, fields []string,
) api.DrivePager {
	return h.ac.NewUserDrivePager(resourceOwner, fields)
}

func (h itemBackupHandler) NewItemPager(
	driveID, link string,
	fields []string,
) api.DriveItemEnumerator {
	return h.ac.NewItemPager(driveID, link, fields)
}

func (h itemBackupHandler) AugmentItemInfo(
	dii details.ItemInfo,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	return augmentItemInfo(dii, item, size, parentPath)
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

// ---------------------------------------------------------------------------
// Restore
// ---------------------------------------------------------------------------

var _ RestoreHandler = &itemRestoreHandler{}

type itemRestoreHandler struct {
	ac api.Drives
}

func NewRestoreHandler(ac api.Client) *itemRestoreHandler {
	return &itemRestoreHandler{ac.Drives()}
}

// AugmentItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func (h itemRestoreHandler) AugmentItemInfo(
	dii details.ItemInfo,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	return augmentItemInfo(dii, item, size, parentPath)
}

func (h itemRestoreHandler) NewItemContentUpload(
	ctx context.Context,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	return h.ac.NewItemContentUpload(ctx, driveID, itemID)
}

func (h itemRestoreHandler) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return h.ac.DeleteItemPermission(ctx, driveID, itemID, permissionID)
}

func (h itemRestoreHandler) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return h.ac.PostItemPermissionUpdate(ctx, driveID, itemID, body)
}

func (h itemRestoreHandler) PostItemInContainer(
	ctx context.Context,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
) (models.DriveItemable, error) {
	return h.ac.PostItemInContainer(ctx, driveID, parentFolderID, newItem)
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

// ---------------------------------------------------------------------------
// Common
// ---------------------------------------------------------------------------

func augmentItemInfo(
	dii details.ItemInfo,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	var email, driveName, driveID string

	if item.GetCreatedBy() != nil && item.GetCreatedBy().GetUser() != nil {
		// User is sometimes not available when created via some
		// external applications (like backup/restore solutions)
		ed, ok := item.GetCreatedBy().GetUser().GetAdditionalData()["email"]
		if ok {
			email = *ed.(*string)
		}
	}

	if item.GetParentReference() != nil {
		driveID = ptr.Val(item.GetParentReference().GetDriveId())
		driveName = strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	}

	var pps string
	if parentPath != nil {
		pps = parentPath.String()
	}

	dii.OneDrive = &details.OneDriveInfo{
		Created:    ptr.Val(item.GetCreatedDateTime()),
		DriveID:    driveID,
		DriveName:  driveName,
		ItemName:   ptr.Val(item.GetName()),
		ItemType:   details.OneDriveItem,
		Modified:   ptr.Val(item.GetLastModifiedDateTime()),
		Owner:      email,
		ParentPath: pps,
		Size:       size,
	}

	return dii
}
