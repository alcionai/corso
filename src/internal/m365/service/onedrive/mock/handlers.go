package mock

import (
	"context"
	"net/http"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// Backup Handler
// ---------------------------------------------------------------------------

type BackupHandler struct {
	ItemInfo details.ItemInfo

	GI  GetsItem
	GIP GetsItemPermission

	PathPrefixFn  pathPrefixer
	PathPrefixErr error

	MetadataPathPrefixFn  metadataPathPrefixer
	MetadataPathPrefixErr error

	CanonPathFn  canonPather
	CanonPathErr error

	ResourceOwner string
	Service       path.ServiceType
	Category      path.CategoryType

	DrivePagerV api.Pager[models.Driveable]
	// driveID -> itemPager
	ItemPagerV map[string]api.DeltaPager[models.DriveItemable]

	LocationIDFn locationIDer

	getCall  int
	GetResps []*http.Response
	GetErrs  []error
}

func DefaultOneDriveBH(resourceOwner string) *BackupHandler {
	return &BackupHandler{
		ItemInfo: details.ItemInfo{
			OneDrive:  &details.OneDriveInfo{},
			Extension: &details.ExtensionData{},
		},
		GI:                   GetsItem{Err: clues.New("not defined")},
		GIP:                  GetsItemPermission{Err: clues.New("not defined")},
		PathPrefixFn:         defaultOneDrivePathPrefixer,
		MetadataPathPrefixFn: defaultOneDriveMetadataPathPrefixer,
		CanonPathFn:          defaultOneDriveCanonPather,
		ResourceOwner:        resourceOwner,
		Service:              path.OneDriveService,
		Category:             path.FilesCategory,
		LocationIDFn:         defaultOneDriveLocationIDer,
		GetResps:             []*http.Response{nil},
		GetErrs:              []error{clues.New("not defined")},
	}
}

func DefaultSharePointBH(resourceOwner string) *BackupHandler {
	return &BackupHandler{
		ItemInfo: details.ItemInfo{
			SharePoint: &details.SharePointInfo{},
			Extension:  &details.ExtensionData{},
		},
		GI:                   GetsItem{Err: clues.New("not defined")},
		GIP:                  GetsItemPermission{Err: clues.New("not defined")},
		PathPrefixFn:         defaultSharePointPathPrefixer,
		MetadataPathPrefixFn: defaultSharePointMetadataPathPrefixer,
		CanonPathFn:          defaultSharePointCanonPather,
		ResourceOwner:        resourceOwner,
		Service:              path.SharePointService,
		Category:             path.LibrariesCategory,
		LocationIDFn:         defaultSharePointLocationIDer,
		GetResps:             []*http.Response{nil},
		GetErrs:              []error{clues.New("not defined")},
	}
}

func (h BackupHandler) PathPrefix(tID, driveID string) (path.Path, error) {
	pp, err := h.PathPrefixFn(tID, h.ResourceOwner, driveID)
	if err != nil {
		return nil, err
	}

	return pp, h.PathPrefixErr
}

func (h BackupHandler) MetadataPathPrefix(tID string) (path.Path, error) {
	pp, err := h.MetadataPathPrefixFn(tID, h.ResourceOwner)
	if err != nil {
		return nil, err
	}

	return pp, h.PathPrefixErr
}

func (h BackupHandler) CanonicalPath(pb *path.Builder, tID string) (path.Path, error) {
	cp, err := h.CanonPathFn(pb, tID, h.ResourceOwner)
	if err != nil {
		return nil, err
	}

	return cp, h.CanonPathErr
}

func (h BackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.Service, h.Category
}

func (h BackupHandler) NewDrivePager(string, []string) api.Pager[models.Driveable] {
	return h.DrivePagerV
}

func (h BackupHandler) NewItemPager(driveID string, _ string, _ []string) api.DeltaPager[models.DriveItemable] {
	return h.ItemPagerV[driveID]
}

func (h BackupHandler) FormatDisplayPath(_ string, pb *path.Builder) string {
	return "/" + pb.String()
}

func (h BackupHandler) NewLocationIDer(driveID string, elems ...string) details.LocationIDer {
	return h.LocationIDFn(driveID, elems...)
}

func (h BackupHandler) AugmentItemInfo(details.ItemInfo, models.DriveItemable, int64, *path.Builder) details.ItemInfo {
	return h.ItemInfo
}

func (h *BackupHandler) Get(context.Context, string, map[string]string) (*http.Response, error) {
	c := h.getCall
	h.getCall++

	// allows mockers to only populate the errors slice
	if h.GetErrs[c] != nil {
		return nil, h.GetErrs[c]
	}

	return h.GetResps[c], h.GetErrs[c]
}

func (h BackupHandler) GetItem(ctx context.Context, _, _ string) (models.DriveItemable, error) {
	return h.GI.GetItem(ctx, "", "")
}

func (h BackupHandler) GetItemPermission(
	ctx context.Context,
	_, _ string,
) (models.PermissionCollectionResponseable, error) {
	return h.GIP.GetItemPermission(ctx, "", "")
}

type canonPather func(*path.Builder, string, string) (path.Path, error)

var defaultOneDriveCanonPather = func(pb *path.Builder, tID, ro string) (path.Path, error) {
	return pb.ToDataLayerOneDrivePath(tID, ro, false)
}

var defaultSharePointCanonPather = func(pb *path.Builder, tID, ro string) (path.Path, error) {
	return pb.ToDataLayerSharePointPath(tID, ro, path.LibrariesCategory, false)
}

type (
	pathPrefixer         func(tID, ro, driveID string) (path.Path, error)
	metadataPathPrefixer func(tID, ro string) (path.Path, error)
)

var defaultOneDrivePathPrefixer = func(tID, ro, driveID string) (path.Path, error) {
	return path.Build(
		tID,
		ro,
		path.OneDriveService,
		path.FilesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

var defaultOneDriveMetadataPathPrefixer = func(tID, ro string) (path.Path, error) {
	return path.Builder{}.ToServiceCategoryMetadataPath(
		tID,
		ro,
		path.OneDriveService,
		path.FilesCategory,
		false)
}

var defaultSharePointPathPrefixer = func(tID, ro, driveID string) (path.Path, error) {
	return path.Build(
		tID,
		ro,
		path.SharePointService,
		path.LibrariesCategory,
		false,
		odConsts.DrivesPathDir,
		driveID,
		odConsts.RootPathDir)
}

var defaultSharePointMetadataPathPrefixer = func(tID, ro string) (path.Path, error) {
	return path.Builder{}.ToServiceCategoryMetadataPath(
		tID,
		ro,
		path.SharePointService,
		path.LibrariesCategory,
		false)
}

type locationIDer func(string, ...string) details.LocationIDer

var defaultOneDriveLocationIDer = func(driveID string, elems ...string) details.LocationIDer {
	return details.NewOneDriveLocationIDer(driveID, elems...)
}

var defaultSharePointLocationIDer = func(driveID string, elems ...string) details.LocationIDer {
	return details.NewSharePointLocationIDer(driveID, elems...)
}

func (h BackupHandler) IsAllPass() bool {
	return true
}

func (h BackupHandler) IncludesDir(string) bool {
	return true
}

// ---------------------------------------------------------------------------
// Get Itemer
// ---------------------------------------------------------------------------

type GetsItem struct {
	Item models.DriveItemable
	Err  error
}

func (m GetsItem) GetItem(
	_ context.Context,
	_, _ string,
) (models.DriveItemable, error) {
	return m.Item, m.Err
}

// ---------------------------------------------------------------------------
// Get Item Permissioner
// ---------------------------------------------------------------------------

type GetsItemPermission struct {
	Perm models.PermissionCollectionResponseable
	Err  error
}

func (m GetsItemPermission) GetItemPermission(
	_ context.Context,
	_, _ string,
) (models.PermissionCollectionResponseable, error) {
	return m.Perm, m.Err
}

// ---------------------------------------------------------------------------
// Restore Handler
// --------------------------------------------------------------------------

type RestoreHandler struct {
	ItemInfo details.ItemInfo

	CollisionKeyMap map[string]api.DriveItemIDType

	CalledDeleteItem   bool
	CalledDeleteItemOn string
	DeleteItemErr      error

	CalledPostItem bool
	PostItemResp   models.DriveItemable
	PostItemErr    error

	DrivePagerV api.Pager[models.Driveable]

	PostDriveResp models.Driveable
	PostDriveErr  error

	UploadSessionErr error
}

func (h RestoreHandler) PostDrive(
	ctx context.Context,
	protectedResourceID, driveName string,
) (models.Driveable, error) {
	return h.PostDriveResp, h.PostDriveErr
}

func (h RestoreHandler) NewDrivePager(string, []string) api.Pager[models.Driveable] {
	return h.DrivePagerV
}

func (h *RestoreHandler) AugmentItemInfo(
	details.ItemInfo,
	models.DriveItemable,
	int64,
	*path.Builder,
) details.ItemInfo {
	return h.ItemInfo
}

func (h *RestoreHandler) GetItemsInContainerByCollisionKey(
	context.Context,
	string, string,
) (map[string]api.DriveItemIDType, error) {
	return h.CollisionKeyMap, nil
}

func (h *RestoreHandler) DeleteItem(
	_ context.Context,
	_, itemID string,
) error {
	h.CalledDeleteItem = true
	h.CalledDeleteItemOn = itemID

	return h.DeleteItemErr
}

func (h *RestoreHandler) DeleteItemPermission(
	context.Context,
	string, string, string,
) error {
	return nil
}

func (h *RestoreHandler) NewItemContentUpload(
	context.Context,
	string, string,
) (models.UploadSessionable, error) {
	return models.NewUploadSession(), h.UploadSessionErr
}

func (h *RestoreHandler) PostItemPermissionUpdate(
	context.Context,
	string, string,
	*drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return drives.NewItemItemsItemInviteResponse(), nil
}

func (h *RestoreHandler) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	return nil, clues.New("not implemented")
}

func (h *RestoreHandler) PostItemInContainer(
	context.Context,
	string, string,
	models.DriveItemable,
	control.CollisionPolicy,
) (models.DriveItemable, error) {
	h.CalledPostItem = true
	return h.PostItemResp, h.PostItemErr
}

func (h *RestoreHandler) GetFolderByName(
	context.Context,
	string, string, string,
) (models.DriveItemable, error) {
	return models.NewDriveItem(), nil
}

func (h *RestoreHandler) GetRootFolder(
	context.Context,
	string,
) (models.DriveItemable, error) {
	return models.NewDriveItem(), nil
}
