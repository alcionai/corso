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
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

type baseSiteHandler struct {
	ac api.Drives
}

func (h baseSiteHandler) NewDrivePager(
	resourceOwner string,
	fields []string,
) pagers.NonDeltaHandler[models.Driveable] {
	return h.ac.NewSiteDrivePager(resourceOwner, fields)
}

func (h baseSiteHandler) AugmentItemInfo(
	dii details.ItemInfo,
	resource idname.Provider,
	item LiteDriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	var pps string

	if parentPath != nil {
		pps = parentPath.String()
	}

	driveName, driveID := getItemDriveInfo(item)

	dii.Extension = &details.ExtensionData{}
	dii.SharePoint = &details.SharePointInfo{
		Created:    ptr.Val(item.GetCreatedDateTime()),
		DriveID:    driveID,
		DriveName:  driveName,
		ItemName:   ptr.Val(item.GetName()),
		ItemType:   details.SharePointLibrary,
		Modified:   ptr.Val(item.GetLastModifiedDateTime()),
		Owner:      getItemCreator(item),
		ParentPath: pps,
		SiteID:     resource.ID(),
		Size:       size,
		WebURL:     resource.Name(),
	}

	return dii
}

var _ BackupHandler = &siteBackupHandler{}

type siteBackupHandler struct {
	baseSiteHandler
	siteID  string
	scope   selectors.SharePointScope
	service path.ServiceType
}

func NewSiteBackupHandler(
	ac api.Drives,
	siteID string,
	scope selectors.SharePointScope,
	service path.ServiceType,
) siteBackupHandler {
	return siteBackupHandler{
		baseSiteHandler: baseSiteHandler{
			ac: ac,
		},
		siteID:  siteID,
		scope:   scope,
		service: service,
	}
}

func (h siteBackupHandler) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return h.ac.Get(ctx, url, headers)
}

func (h siteBackupHandler) PathPrefix(
	tenantID, driveID string,
) (path.Path, error) {
	return path.Build(
		tenantID,
		h.siteID,
		h.service,
		path.LibrariesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h siteBackupHandler) MetadataPathPrefix(
	tenantID string,
) (path.Path, error) {
	p, err := path.BuildMetadata(
		tenantID,
		h.siteID,
		h.service,
		path.LibrariesCategory,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path")
	}

	return p, nil
}

func (h siteBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.ToDataLayerPath(tenantID, h.siteID, h.service, path.LibrariesCategory, false)
}

func (h siteBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.service, path.LibrariesCategory
}

func (h siteBackupHandler) FormatDisplayPath(
	driveName string,
	pb *path.Builder,
) string {
	return "/" + driveName + "/" + pb.String()
}

func (h siteBackupHandler) NewLocationIDer(
	driveID string,
	elems ...string,
) details.LocationIDer {
	return details.NewSharePointLocationIDer(driveID, elems...)
}

func (h siteBackupHandler) GetItemPermission(
	ctx context.Context,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	return h.ac.GetItemPermission(ctx, driveID, itemID)
}

func (h siteBackupHandler) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	return h.ac.GetItem(ctx, driveID, itemID)
}

func (h siteBackupHandler) IsAllPass() bool {
	return h.scope.IsAny(selectors.SharePointLibraryFolder)
}

func (h siteBackupHandler) IncludesDir(dir string) bool {
	return h.scope.Matches(selectors.SharePointLibraryFolder, dir)
}

func (h siteBackupHandler) EnumerateDriveItemsDelta(
	ctx context.Context,
	driveID, prevDeltaLink string,
	cc api.CallConfig,
) pagers.NextPageResulter[models.DriveItemable] {
	return h.ac.EnumerateDriveItemsDelta(ctx, driveID, prevDeltaLink, cc)
}

// ---------------------------------------------------------------------------
// Restore
// ---------------------------------------------------------------------------

var _ RestoreHandler = &siteRestoreHandler{}

type siteRestoreHandler struct {
	baseSiteHandler
	ac      api.Client
	service path.ServiceType
}

func NewSiteRestoreHandler(ac api.Client, service path.ServiceType) siteRestoreHandler {
	return siteRestoreHandler{
		baseSiteHandler: baseSiteHandler{
			ac: ac.Drives(),
		},
		ac:      ac,
		service: service,
	}
}

func (h siteRestoreHandler) PostDrive(
	ctx context.Context,
	siteID, driveName string,
) (models.Driveable, error) {
	return h.ac.Lists().PostDrive(ctx, siteID, driveName)
}

func (h siteRestoreHandler) DeleteItem(
	ctx context.Context,
	driveID, itemID string,
) error {
	return h.ac.Drives().DeleteItem(ctx, driveID, itemID)
}

func (h siteRestoreHandler) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return h.ac.Drives().DeleteItemPermission(ctx, driveID, itemID, permissionID)
}

func (h siteRestoreHandler) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]api.DriveItemIDType, error) {
	m, err := h.ac.Drives().GetItemsInContainerByCollisionKey(ctx, driveID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (h siteRestoreHandler) NewItemContentUpload(
	ctx context.Context,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	return h.ac.Drives().NewItemContentUpload(ctx, driveID, itemID)
}

func (h siteRestoreHandler) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return h.ac.Drives().PostItemPermissionUpdate(ctx, driveID, itemID, body)
}

func (h siteRestoreHandler) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	return h.ac.Drives().PostItemLinkShareUpdate(ctx, driveID, itemID, body)
}

func (h siteRestoreHandler) PostItemInContainer(
	ctx context.Context,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
	onCollision control.CollisionPolicy,
) (models.DriveItemable, error) {
	return h.ac.Drives().PostItemInContainer(ctx, driveID, parentFolderID, newItem, onCollision)
}

func (h siteRestoreHandler) GetFolderByName(
	ctx context.Context,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	return h.ac.Drives().GetFolderByName(ctx, driveID, parentFolderID, folderName)
}

func (h siteRestoreHandler) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	return h.ac.Drives().GetRootFolder(ctx, driveID)
}
