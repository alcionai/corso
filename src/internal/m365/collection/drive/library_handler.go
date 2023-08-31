package drive

import (
	"context"
	"net/http"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ BackupHandler = &libraryBackupHandler{}

type libraryBackupHandler struct {
	ac      api.Drives
	siteID  string
	scope   selectors.SharePointScope
	service path.ServiceType
}

func NewLibraryBackupHandler(
	ac api.Drives,
	siteID string,
	scope selectors.SharePointScope,
	service path.ServiceType,
) libraryBackupHandler {
	return libraryBackupHandler{ac, siteID, scope, service}
}

func (h libraryBackupHandler) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return h.ac.Get(ctx, url, headers)
}

func (h libraryBackupHandler) PathPrefix(
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

func (h libraryBackupHandler) MetadataPathPrefix(
	tenantID string,
) (path.Path, error) {
	p, err := path.Builder{}.ToServiceCategoryMetadataPath(
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

func (h libraryBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.ToDataLayerPath(tenantID, h.siteID, h.service, path.LibrariesCategory, false)
}

func (h libraryBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.service, path.LibrariesCategory
}

func (h libraryBackupHandler) NewDrivePager(
	resourceOwner string,
	fields []string,
) api.Pager[models.Driveable] {
	return h.ac.NewSiteDrivePager(resourceOwner, fields)
}

func (h libraryBackupHandler) NewItemPager(
	driveID, link string,
	fields []string,
) api.DeltaPager[models.DriveItemable] {
	return h.ac.NewDriveItemDeltaPager(driveID, link, fields)
}

func (h libraryBackupHandler) AugmentItemInfo(
	dii details.ItemInfo,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	return augmentItemInfo(dii, h.service, item, size, parentPath)
}

// constructWebURL is a helper function for recreating the webURL
// for the originating SharePoint site. Uses the additionalData map
// from a models.DriveItemable that possesses a downloadURL within the map.
// Returns "" if the map is nil or key is not present.
func constructWebURL(adtl map[string]any) string {
	var (
		desiredKey = "@microsoft.graph.downloadUrl"
		sep        = `/_layouts`
		url        string
	)

	if adtl == nil {
		return url
	}

	r := adtl[desiredKey]
	point, ok := r.(*string)

	if !ok {
		return url
	}

	value := ptr.Val(point)
	if len(value) == 0 {
		return url
	}

	temp := strings.Split(value, sep)
	url = temp[0]

	return url
}

func (h libraryBackupHandler) FormatDisplayPath(
	driveName string,
	pb *path.Builder,
) string {
	return "/" + driveName + "/" + pb.String()
}

func (h libraryBackupHandler) NewLocationIDer(
	driveID string,
	elems ...string,
) details.LocationIDer {
	// TODO(meain): path related changes for groups
	return details.NewSharePointLocationIDer(driveID, elems...)
}

func (h libraryBackupHandler) GetItemPermission(
	ctx context.Context,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	return h.ac.GetItemPermission(ctx, driveID, itemID)
}

func (h libraryBackupHandler) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	return h.ac.GetItem(ctx, driveID, itemID)
}

func (h libraryBackupHandler) IsAllPass() bool {
	return h.scope.IsAny(selectors.SharePointLibraryFolder)
}

func (h libraryBackupHandler) IncludesDir(dir string) bool {
	return h.scope.Matches(selectors.SharePointLibraryFolder, dir)
}

// ---------------------------------------------------------------------------
// Restore
// ---------------------------------------------------------------------------

var _ RestoreHandler = &libraryRestoreHandler{}

type libraryRestoreHandler struct {
	ac      api.Client
	service path.ServiceType
}

func NewLibraryRestoreHandler(ac api.Client, service path.ServiceType) libraryRestoreHandler {
	return libraryRestoreHandler{ac, service}
}

func (h libraryRestoreHandler) PostDrive(
	ctx context.Context,
	siteID, driveName string,
) (models.Driveable, error) {
	return h.ac.Lists().PostDrive(ctx, siteID, driveName)
}

func (h libraryRestoreHandler) NewDrivePager(
	resourceOwner string,
	fields []string,
) api.Pager[models.Driveable] {
	return h.ac.Drives().NewSiteDrivePager(resourceOwner, fields)
}

func (h libraryRestoreHandler) AugmentItemInfo(
	dii details.ItemInfo,
	item models.DriveItemable,
	size int64,
	parentPath *path.Builder,
) details.ItemInfo {
	return augmentItemInfo(dii, h.service, item, size, parentPath)
}

func (h libraryRestoreHandler) DeleteItem(
	ctx context.Context,
	driveID, itemID string,
) error {
	return h.ac.Drives().DeleteItem(ctx, driveID, itemID)
}

func (h libraryRestoreHandler) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return h.ac.Drives().DeleteItemPermission(ctx, driveID, itemID, permissionID)
}

func (h libraryRestoreHandler) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]api.DriveItemIDType, error) {
	m, err := h.ac.Drives().GetItemsInContainerByCollisionKey(ctx, driveID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (h libraryRestoreHandler) NewItemContentUpload(
	ctx context.Context,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	return h.ac.Drives().NewItemContentUpload(ctx, driveID, itemID)
}

func (h libraryRestoreHandler) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return h.ac.Drives().PostItemPermissionUpdate(ctx, driveID, itemID, body)
}

func (h libraryRestoreHandler) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	return h.ac.Drives().PostItemLinkShareUpdate(ctx, driveID, itemID, body)
}

func (h libraryRestoreHandler) PostItemInContainer(
	ctx context.Context,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
	onCollision control.CollisionPolicy,
) (models.DriveItemable, error) {
	return h.ac.Drives().PostItemInContainer(ctx, driveID, parentFolderID, newItem, onCollision)
}

func (h libraryRestoreHandler) GetFolderByName(
	ctx context.Context,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	return h.ac.Drives().GetFolderByName(ctx, driveID, parentFolderID, folderName)
}

func (h libraryRestoreHandler) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	return h.ac.Drives().GetRootFolder(ctx, driveID)
}
