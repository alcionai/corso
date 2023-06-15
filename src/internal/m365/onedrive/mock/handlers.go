package mock

import (
	"context"
	"net/http"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/backup/details"
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

	CanonPathFn  canonPather
	CanonPathErr error

	Service  path.ServiceType
	Category path.CategoryType

	DrivePagerV api.DrivePager
	// driveID -> itemPager
	ItemPagerV map[string]api.DriveItemEnumerator

	LocationIDFn locationIDer

	getCall  int
	GetResps []*http.Response
	GetErrs  []error
}

func DefaultOneDriveBH() *BackupHandler {
	return &BackupHandler{
		ItemInfo:     details.ItemInfo{OneDrive: &details.OneDriveInfo{}},
		GI:           GetsItem{Err: clues.New("not defined")},
		GIP:          GetsItemPermission{Err: clues.New("not defined")},
		PathPrefixFn: defaultOneDrivePathPrefixer,
		CanonPathFn:  defaultOneDriveCanonPather,
		Service:      path.OneDriveService,
		Category:     path.FilesCategory,
		LocationIDFn: defaultOneDriveLocationIDer,
		GetResps:     []*http.Response{nil},
		GetErrs:      []error{clues.New("not defined")},
	}
}

func DefaultSharePointBH() *BackupHandler {
	return &BackupHandler{
		ItemInfo:     details.ItemInfo{SharePoint: &details.SharePointInfo{}},
		GI:           GetsItem{Err: clues.New("not defined")},
		GIP:          GetsItemPermission{Err: clues.New("not defined")},
		PathPrefixFn: defaultSharePointPathPrefixer,
		CanonPathFn:  defaultSharePointCanonPather,
		Service:      path.SharePointService,
		Category:     path.LibrariesCategory,
		LocationIDFn: defaultSharePointLocationIDer,
		GetResps:     []*http.Response{nil},
		GetErrs:      []error{clues.New("not defined")},
	}
}

func (h BackupHandler) PathPrefix(tID, ro, driveID string) (path.Path, error) {
	pp, err := h.PathPrefixFn(tID, ro, driveID)
	if err != nil {
		return nil, err
	}

	return pp, h.PathPrefixErr
}

func (h BackupHandler) CanonicalPath(pb *path.Builder, tID, ro string) (path.Path, error) {
	cp, err := h.CanonPathFn(pb, tID, ro)
	if err != nil {
		return nil, err
	}

	return cp, h.CanonPathErr
}

func (h BackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.Service, h.Category
}

func (h BackupHandler) NewDrivePager(string, []string) api.DrivePager {
	return h.DrivePagerV
}

func (h BackupHandler) NewItemPager(driveID string, _ string, _ []string) api.DriveItemEnumerator {
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

type pathPrefixer func(tID, ro, driveID string) (path.Path, error)

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
