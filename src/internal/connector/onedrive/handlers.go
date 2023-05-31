package onedrive

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ItemInfoAugmenter interface {
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

	PathPrefix(tenantID, resourceOwner, driveID string) (path.Path, error)
	CanonicalPath(
		folders *path.Builder,
		tenantID, resourceOwner string,
	) (path.Path, error)
	ServiceCat() (path.ServiceType, path.CategoryType)
	DrivePager(resourceOwner string, fields []string) api.DrivePager
	ItemPager(driveID, link string, fields []string) api.DriveItemEnumerator
	FormatDisplayPath(driveName string, parentPath *path.Builder) string
	NewLocationIDer(driveID string, elems ...string) details.LocationIDer
	Requester() graph.Requester

	PermissionGetter() GetItemPermissioner
	ItemGetter() GetItemer
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
	ItemInfoAugmenter

	FolderByNameGetter() GetFolderByNamer
	ItemPoster() PostItemer
	ItemInContainerPoster() PostItemInContainerer
	ItemPermissionDeleter() DeleteItemPermissioner
	ItemPermissionUpdater() UpdateItemPermissioner
	RootFolderGetter() GetRootFolderer
}

type PostItemer interface {
	PostItem(
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
	GetRootFolder(
		ctx context.Context,
		driveID string,
	) (models.DriveItemable, error)
}
