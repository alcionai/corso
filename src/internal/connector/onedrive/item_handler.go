package onedrive

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	odConsts "github.com/alcionai/corso/src/internal/connector/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ---------------------------------------------------------------------------
// backup
// ---------------------------------------------------------------------------

var _ BackupHandler = &itemBackupHandler{}

type itemBackupHandler struct {
	ac api.Drives
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

func (h itemBackupHandler) DrivePager(
	resourceOwner string, fields []string,
) api.DrivePager {
	return h.ac.NewUserDrivePager(resourceOwner, fields)
}

func (h itemBackupHandler) ItemPager(
	driveID, link string,
	fields []string,
) api.DriveItemEnumerator {
	return h.ac.NewItemPager(driveID, link, fields)
}

// AugmentItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
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

func (h itemBackupHandler) Requester() graph.Requester            { return h.ac.Requester }
func (h itemBackupHandler) PermissionGetter() GetItemPermissioner { return h.ac }
func (h itemBackupHandler) ItemGetter() GetItemer                 { return h.ac }

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

func (h itemRestoreHandler) FolderByNameGetter() GetFolderByNamer          { return h.ac }
func (h itemRestoreHandler) ItemPoster() PostItemer                        { return h.ac }
func (h itemRestoreHandler) ItemInContainerPoster() PostItemInContainerer  { return h.ac }
func (h itemRestoreHandler) ItemPermissionDeleter() DeleteItemPermissioner { return h.ac }
func (h itemRestoreHandler) ItemPermissionUpdater() UpdateItemPermissioner { return h.ac }
func (h itemRestoreHandler) RootFolderGetter() GetRootFolderer             { return h.ac }

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
