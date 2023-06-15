package onedrive

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ItemInfoAugmenter interface {
	// AugmentItemInfo will populate a details.<Service>Info struct
	// with properties from the drive item.  ItemSize is passed in
	// separately for restore processes because the local itemable
	// doesn't have its size value updated as a side effect of creation,
	// and kiota drops any SetSize update.
	AugmentItemInfo(
		dii details.ItemInfo,
		item models.DriveItemable,
		size int64,
		parentPath *path.Builder,
	) details.ItemInfo
}

// ---------------------------------------------------------------------------
// backup
// ---------------------------------------------------------------------------

type BackupHandler interface {
	ItemInfoAugmenter
	api.Getter
	GetItemPermissioner
	GetItemer

	// PathPrefix constructs the service and category specific path prefix for
	// the given values.
	PathPrefix(tenantID, resourceOwner, driveID string) (path.Path, error)

	// CanonicalPath constructs the service and category specific path for
	// the given values.
	CanonicalPath(
		folders *path.Builder,
		tenantID, resourceOwner string,
	) (path.Path, error)

	// ServiceCat returns the service and category used by this implementation.
	ServiceCat() (path.ServiceType, path.CategoryType)
	NewDrivePager(resourceOwner string, fields []string) api.DrivePager
	NewItemPager(driveID, link string, fields []string) api.DriveItemEnumerator
	// FormatDisplayPath creates a human-readable string to represent the
	// provided path.
	FormatDisplayPath(driveName string, parentPath *path.Builder) string
	NewLocationIDer(driveID string, elems ...string) details.LocationIDer

	// scope wrapper funcs
	IsAllPass() bool
	IncludesDir(dir string) bool
}

type GetItemPermissioner interface {
	GetItemPermission(
		ctx context.Context,
		driveID, itemID string,
	) (models.PermissionCollectionResponseable, error)
}

type GetItemer interface {
	GetItem(
		ctx context.Context,
		driveID, itemID string,
	) (models.DriveItemable, error)
}

// ---------------------------------------------------------------------------
// restore
// ---------------------------------------------------------------------------

type RestoreHandler interface {
	DeleteItemPermissioner
	GetFolderByNamer
	GetRootFolderer
	ItemInfoAugmenter
	NewItemContentUploader
	PostItemInContainerer
	UpdateItemPermissioner
}

type NewItemContentUploader interface {
	// NewItemContentUpload creates an upload session which is used as a writer
	// for large item content.
	NewItemContentUpload(
		ctx context.Context,
		driveID, itemID string,
	) (models.UploadSessionable, error)
}

type DeleteItemPermissioner interface {
	DeleteItemPermission(
		ctx context.Context,
		driveID, itemID, permissionID string,
	) error
}

type UpdateItemPermissioner interface {
	PostItemPermissionUpdate(
		ctx context.Context,
		driveID, itemID string,
		body *drives.ItemItemsItemInvitePostRequestBody,
	) (drives.ItemItemsItemInviteResponseable, error)
}

type PostItemInContainerer interface {
	PostItemInContainer(
		ctx context.Context,
		driveID, parentFolderID string,
		newItem models.DriveItemable,
	) (models.DriveItemable, error)
}

type GetFolderByNamer interface {
	GetFolderByName(
		ctx context.Context,
		driveID, parentFolderID, folderID string,
	) (models.DriveItemable, error)
}

type GetRootFolderer interface {
	// GetRootFolder gets the root folder for the drive.
	GetRootFolder(
		ctx context.Context,
		driveID string,
	) (models.DriveItemable, error)
}
