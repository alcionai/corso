package drive

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
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
	NewDrivePagerer

	// PathPrefix constructs the service and category specific path prefix for
	// the given values.
	PathPrefix(tenantID, driveID string) (path.Path, error)

	// MetadataPathPrefix returns the prefix path for metadata
	MetadataPathPrefix(tenantID string) (path.Path, error)

	// CanonicalPath constructs the service and category specific path for
	// the given values.
	CanonicalPath(folders *path.Builder, tenantID string) (path.Path, error)

	// ServiceCat returns the service and category used by this implementation.
	ServiceCat() (path.ServiceType, path.CategoryType)
	NewItemPager(driveID, link string, fields []string) api.DeltaPager[models.DriveItemable]
	// FormatDisplayPath creates a human-readable string to represent the
	// provided path.
	FormatDisplayPath(driveName string, parentPath *path.Builder) string
	NewLocationIDer(driveID string, elems ...string) details.LocationIDer

	// scope wrapper funcs
	IsAllPass() bool
	IncludesDir(dir string) bool
}

type NewDrivePagerer interface {
	NewDrivePager(resourceOwner string, fields []string) api.Pager[models.Driveable]
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
	DeleteItemer
	DeleteItemPermissioner
	GetFolderByNamer
	GetItemsByCollisionKeyser
	GetRootFolderer
	ItemInfoAugmenter
	NewDrivePagerer
	NewItemContentUploader
	PostDriver
	PostItemInContainerer
	DeleteItemPermissioner
	UpdateItemPermissioner
	UpdateItemLinkSharer
}

type DeleteItemer interface {
	DeleteItem(
		ctx context.Context,
		driveID, itemID string,
	) error
}

type DeleteItemPermissioner interface {
	DeleteItemPermission(
		ctx context.Context,
		driveID, itemID, permissionID string,
	) error
}

type GetItemsByCollisionKeyser interface {
	// GetItemsInContainerByCollisionKey looks up all items currently in
	// the container, and returns them in a map[collisionKey]itemID.
	// The collision key is uniquely defined by each category of data.
	// Collision key checks are used during restore to handle the on-
	// collision restore configurations that cause the item restore to get
	// skipped, replaced, or copied.
	GetItemsInContainerByCollisionKey(
		ctx context.Context,
		driveID, containerID string,
	) (map[string]api.DriveItemIDType, error)
}

type NewItemContentUploader interface {
	// NewItemContentUpload creates an upload session which is used as a writer
	// for large item content.
	NewItemContentUpload(
		ctx context.Context,
		driveID, itemID string,
	) (models.UploadSessionable, error)
}

type UpdateItemPermissioner interface {
	PostItemPermissionUpdate(
		ctx context.Context,
		driveID, itemID string,
		body *drives.ItemItemsItemInvitePostRequestBody,
	) (drives.ItemItemsItemInviteResponseable, error)
}

type UpdateItemLinkSharer interface {
	PostItemLinkShareUpdate(
		ctx context.Context,
		driveID, itemID string,
		body *drives.ItemItemsItemCreateLinkPostRequestBody,
	) (models.Permissionable, error)
}

type PostDriver interface {
	PostDrive(
		ctx context.Context,
		protectedResourceID, driveName string,
	) (models.Driveable, error)
}

type PostItemInContainerer interface {
	PostItemInContainer(
		ctx context.Context,
		driveID, parentFolderID string,
		newItem models.DriveItemable,
		onCollision control.CollisionPolicy,
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
