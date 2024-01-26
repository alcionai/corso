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
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

// ---------------------------------------------------------------------------
// backup
// ---------------------------------------------------------------------------

type baseUserDriveHandler struct {
	ac api.Drives
	qp graph.QueryParams
}

func (h baseUserDriveHandler) NewDrivePager(
	fields []string,
) pagers.NonDeltaHandler[models.Driveable] {
	return h.ac.NewUserDrivePager(h.qp.ProtectedResource.ID(), fields)
}

// AugmentItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func (h baseUserDriveHandler) AugmentItemInfo(
	dii details.ItemInfo,
	resource idname.Provider,
	item *custom.DriveItem,
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

var _ BackupHandler = &userDriveBackupHandler{}

type userDriveBackupHandler struct {
	baseUserDriveHandler
	scope selectors.OneDriveScope
}

func NewUserDriveBackupHandler(
	qp graph.QueryParams,
	ac api.Drives,
	scope selectors.OneDriveScope,
) *userDriveBackupHandler {
	return &userDriveBackupHandler{
		baseUserDriveHandler: baseUserDriveHandler{
			ac: ac,
			qp: qp,
		},
		scope: scope,
	}
}

func (h userDriveBackupHandler) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return h.ac.Get(ctx, url, headers)
}

func (h userDriveBackupHandler) PathPrefix(
	driveID string,
) (path.Path, error) {
	return path.Build(
		h.qp.TenantID,
		h.qp.ProtectedResource.ID(),
		path.OneDriveService,
		path.FilesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

func (h userDriveBackupHandler) MetadataPathPrefix() (path.Path, error) {
	p, err := path.BuildMetadata(
		h.qp.TenantID,
		h.qp.ProtectedResource.ID(),
		path.OneDriveService,
		path.FilesCategory,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path")
	}

	return p, nil
}

func (h userDriveBackupHandler) CanonicalPath(
	folders *path.Builder,
) (path.Path, error) {
	return path.Build(
		h.qp.TenantID,
		h.qp.ProtectedResource.ID(),
		path.OneDriveService,
		path.FilesCategory,
		false,
		folders.Elements()...)
}

func (h userDriveBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.OneDriveService, path.FilesCategory
}

func (h userDriveBackupHandler) FormatDisplayPath(
	_ string, // drive name not displayed for onedrive
	pb *path.Builder,
) string {
	return "/" + pb.String()
}

func (h userDriveBackupHandler) NewLocationIDer(
	driveID string,
	elems ...string,
) details.LocationIDer {
	return details.NewOneDriveLocationIDer(driveID, elems...)
}

func (h userDriveBackupHandler) GetItemPermission(
	ctx context.Context,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	return h.ac.GetItemPermission(ctx, driveID, itemID)
}

func (h userDriveBackupHandler) GetItem(
	ctx context.Context,
	driveID, itemID string,
) (models.DriveItemable, error) {
	return h.ac.GetItem(ctx, driveID, itemID)
}

func (h userDriveBackupHandler) IsAllPass() bool {
	return h.scope.IsAny(selectors.OneDriveFolder)
}

func (h userDriveBackupHandler) IncludesDir(dir string) bool {
	return h.scope.Matches(selectors.OneDriveFolder, dir)
}

func (h userDriveBackupHandler) EnumerateDriveItemsDelta(
	ctx context.Context,
	driveID, prevDeltaLink string,
	cc api.CallConfig,
) pagers.NextPageResulter[models.DriveItemable] {
	return h.ac.EnumerateDriveItemsDelta(ctx, driveID, prevDeltaLink, cc)
}

func (h userDriveBackupHandler) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	return h.ac.Drives().GetRootFolder(ctx, driveID)
}

// ---------------------------------------------------------------------------
// Restore
// ---------------------------------------------------------------------------

var _ RestoreHandler = &userDriveRestoreHandler{}

type userDriveRestoreHandler struct {
	baseUserDriveHandler
}

func NewUserDriveRestoreHandler(
	ac api.Client,
) *userDriveRestoreHandler {
	return &userDriveRestoreHandler{
		baseUserDriveHandler: baseUserDriveHandler{
			ac: ac.Drives(),
		},
	}
}

func (h userDriveRestoreHandler) NewDrivePager(
	protectedResourceID string,
	fields []string,
) pagers.NonDeltaHandler[models.Driveable] {
	return h.ac.NewUserDrivePager(protectedResourceID, fields)
}

func (h userDriveRestoreHandler) PostDrive(
	context.Context,
	string, string,
) (models.Driveable, error) {
	return nil, clues.New("creating drives in oneDrive is not supported")
}

func (h userDriveRestoreHandler) DeleteItem(
	ctx context.Context,
	driveID, itemID string,
) error {
	return h.ac.DeleteItem(ctx, driveID, itemID)
}

func (h userDriveRestoreHandler) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return h.ac.DeleteItemPermission(ctx, driveID, itemID, permissionID)
}

func (h userDriveRestoreHandler) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]api.DriveItemIDType, error) {
	m, err := h.ac.GetItemsInContainerByCollisionKey(ctx, driveID, containerID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (h userDriveRestoreHandler) NewItemContentUpload(
	ctx context.Context,
	driveID, itemID string,
) (models.UploadSessionable, error) {
	return h.ac.NewItemContentUpload(ctx, driveID, itemID)
}

func (h userDriveRestoreHandler) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return h.ac.PostItemPermissionUpdate(ctx, driveID, itemID, body)
}

func (h userDriveRestoreHandler) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	return h.ac.PostItemLinkShareUpdate(ctx, driveID, itemID, body)
}

func (h userDriveRestoreHandler) PostItemInContainer(
	ctx context.Context,
	driveID, parentFolderID string,
	newItem models.DriveItemable,
	onCollision control.CollisionPolicy,
) (models.DriveItemable, error) {
	return h.ac.PostItemInContainer(ctx, driveID, parentFolderID, newItem, onCollision)
}

func (h userDriveRestoreHandler) GetFolderByName(
	ctx context.Context,
	driveID, parentFolderID, folderName string,
) (models.DriveItemable, error) {
	return h.ac.GetFolderByName(ctx, driveID, parentFolderID, folderName)
}

func (h userDriveRestoreHandler) GetRootFolder(
	ctx context.Context,
	driveID string,
) (models.DriveItemable, error) {
	return h.ac.GetRootFolder(ctx, driveID)
}
