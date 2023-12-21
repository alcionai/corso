package mock

// ---------------------------------------------------------------------------
// 								>>> TODO <<<
//              https://github.com/alcionai/corso/issues/4846
// This file's functions are duplicated into /drive/helper_test.go, which
// should act as the clear primary owner of that functionality.  However,
// packages outside of /drive (such as sharepoint) depend on these helpers
// for test functionality.  We'll want to unify the two at some point.
// In the meantime, make sure you're referencing and updating the correct
// set of helpers (prefer the /drive version over this one).
// ---------------------------------------------------------------------------

import (
	"context"
	"fmt"
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
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

// ---------------------------------------------------------------------------
// Backup Handler
// ---------------------------------------------------------------------------

type BackupHandler[T any] struct {
	ItemInfo details.ItemInfo
	// FIXME: this is a hacky solution.  Better to use an interface
	// and plug in the selector scope there.
	Sel selectors.Selector

	DriveItemEnumeration EnumerateDriveItemsDelta

	GI  GetsItem
	GIP GetsItemPermission

	PathPrefixFn  pathPrefixer
	PathPrefixErr error

	MetadataPathPrefixFn  metadataPathPrefixer
	MetadataPathPrefixErr error

	CanonPathFn  canonPather
	CanonPathErr error

	ProtectedResource idname.Provider
	Service           path.ServiceType
	Category          path.CategoryType

	// driveID -> itemPager
	ItemPagerV map[string]pagers.DeltaHandler[models.DriveItemable]

	LocationIDFn locationIDer

	getCall  int
	GetResps []*http.Response
	GetErrs  []error

	RootFolder models.DriveItemable
}

func stubRootFolder() models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(ptr.To(odConsts.RootPathDir))
	item.SetId(ptr.To(odConsts.RootID))
	item.SetRoot(models.NewRoot())
	item.SetFolder(models.NewFolder())

	return item
}

func DefaultOneDriveBH(resourceOwner string) *BackupHandler[models.DriveItemable] {
	sel := selectors.NewOneDriveBackup([]string{resourceOwner})
	sel.Include(sel.AllData())

	return &BackupHandler[models.DriveItemable]{
		ItemInfo: details.ItemInfo{
			OneDrive:  &details.OneDriveInfo{},
			Extension: &details.ExtensionData{},
		},
		Sel:                  sel.Selector,
		DriveItemEnumeration: EnumerateDriveItemsDelta{},
		GI:                   GetsItem{Err: clues.New("not defined")},
		GIP:                  GetsItemPermission{Err: clues.New("not defined")},
		PathPrefixFn:         defaultOneDrivePathPrefixer,
		MetadataPathPrefixFn: defaultOneDriveMetadataPathPrefixer,
		CanonPathFn:          defaultOneDriveCanonPather,
		ProtectedResource:    idname.NewProvider(resourceOwner, resourceOwner),
		Service:              path.OneDriveService,
		Category:             path.FilesCategory,
		LocationIDFn:         defaultOneDriveLocationIDer,
		GetResps:             []*http.Response{nil},
		GetErrs:              []error{clues.New("not defined")},
		RootFolder:           stubRootFolder(),
	}
}

func DefaultSharePointBH(resourceOwner string) *BackupHandler[models.DriveItemable] {
	sel := selectors.NewOneDriveBackup([]string{resourceOwner})
	sel.Include(sel.AllData())

	return &BackupHandler[models.DriveItemable]{
		ItemInfo: details.ItemInfo{
			SharePoint: &details.SharePointInfo{},
			Extension:  &details.ExtensionData{},
		},
		Sel:                  sel.Selector,
		GI:                   GetsItem{Err: clues.New("not defined")},
		GIP:                  GetsItemPermission{Err: clues.New("not defined")},
		PathPrefixFn:         defaultSharePointPathPrefixer,
		MetadataPathPrefixFn: defaultSharePointMetadataPathPrefixer,
		CanonPathFn:          defaultSharePointCanonPather,
		ProtectedResource:    idname.NewProvider(resourceOwner, resourceOwner),
		Service:              path.SharePointService,
		Category:             path.LibrariesCategory,
		LocationIDFn:         defaultSharePointLocationIDer,
		GetResps:             []*http.Response{nil},
		GetErrs:              []error{clues.New("not defined")},
		RootFolder:           stubRootFolder(),
	}
}

func DefaultDriveBHWith(
	resource string,
	enumerator EnumerateDriveItemsDelta,
) *BackupHandler[models.DriveItemable] {
	mbh := DefaultOneDriveBH(resource)
	mbh.DriveItemEnumeration = enumerator

	return mbh
}

func (h BackupHandler[T]) PathPrefix(tID, driveID string) (path.Path, error) {
	pp, err := h.PathPrefixFn(tID, h.ProtectedResource.ID(), driveID)
	if err != nil {
		return nil, err
	}

	return pp, h.PathPrefixErr
}

func (h BackupHandler[T]) MetadataPathPrefix(tID string) (path.Path, error) {
	pp, err := h.MetadataPathPrefixFn(tID, h.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	return pp, h.MetadataPathPrefixErr
}

func (h BackupHandler[T]) CanonicalPath(pb *path.Builder, tID string) (path.Path, error) {
	cp, err := h.CanonPathFn(pb, tID, h.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	return cp, h.CanonPathErr
}

func (h BackupHandler[T]) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.Service, h.Category
}

func (h BackupHandler[T]) NewDrivePager(string, []string) pagers.NonDeltaHandler[models.Driveable] {
	return h.DriveItemEnumeration.DrivePager()
}

func (h BackupHandler[T]) FormatDisplayPath(_ string, pb *path.Builder) string {
	return "/" + pb.String()
}

func (h BackupHandler[T]) NewLocationIDer(driveID string, elems ...string) details.LocationIDer {
	return h.LocationIDFn(driveID, elems...)
}

func (h BackupHandler[T]) AugmentItemInfo(
	details.ItemInfo,
	idname.Provider,
	*custom.DriveItem,
	int64,
	*path.Builder,
) details.ItemInfo {
	return h.ItemInfo
}

func (h *BackupHandler[T]) Get(context.Context, string, map[string]string) (*http.Response, error) {
	c := h.getCall
	h.getCall++

	// allows mockers to only populate the errors slice
	if h.GetErrs[c] != nil {
		return nil, h.GetErrs[c]
	}

	return h.GetResps[c], h.GetErrs[c]
}

func (h BackupHandler[T]) EnumerateDriveItemsDelta(
	ctx context.Context,
	driveID, prevDeltaLink string,
	cc api.CallConfig,
) pagers.NextPageResulter[models.DriveItemable] {
	return h.DriveItemEnumeration.EnumerateDriveItemsDelta(
		ctx,
		driveID,
		prevDeltaLink,
		cc)
}

func (h BackupHandler[T]) GetItem(ctx context.Context, _, _ string) (models.DriveItemable, error) {
	return h.GI.GetItem(ctx, "", "")
}

func (h BackupHandler[T]) GetItemPermission(
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
	return path.BuildMetadata(
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
	return path.BuildMetadata(
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
	return details.NewSharePointLocationIDer(path.LibrariesCategory, driveID, elems...)
}

func (h BackupHandler[T]) IsAllPass() bool {
	scope := h.Sel.Includes[0]
	return selectors.IsAnyTarget(selectors.SharePointScope(scope), selectors.SharePointLibraryFolder) ||
		selectors.IsAnyTarget(selectors.OneDriveScope(scope), selectors.OneDriveFolder)
}

func (h BackupHandler[T]) IncludesDir(dir string) bool {
	scope := h.Sel.Includes[0]
	return selectors.SharePointScope(scope).Matches(selectors.SharePointLibraryFolder, dir) ||
		selectors.OneDriveScope(scope).Matches(selectors.OneDriveFolder, dir)
}

func (h BackupHandler[T]) GetRootFolder(context.Context, string) (models.DriveItemable, error) {
	return h.RootFolder, nil
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
// Drive Items Enumerator
// ---------------------------------------------------------------------------

type NextPage struct {
	Items []models.DriveItemable
	Reset bool
}

type EnumerateDriveItemsDelta struct {
	DrivePagers map[string]*DeltaDriveEnumerator
}

func DriveEnumerator(
	ds ...*DeltaDriveEnumerator,
) EnumerateDriveItemsDelta {
	enumerator := EnumerateDriveItemsDelta{
		DrivePagers: map[string]*DeltaDriveEnumerator{},
	}

	for _, drive := range ds {
		enumerator.DrivePagers[drive.Drive.ID] = drive
	}

	return enumerator
}

func (en EnumerateDriveItemsDelta) EnumerateDriveItemsDelta(
	_ context.Context,
	driveID, _ string,
	_ api.CallConfig,
) pagers.NextPageResulter[models.DriveItemable] {
	iterator := en.DrivePagers[driveID]
	return iterator.nextDelta()
}

func (en EnumerateDriveItemsDelta) DrivePager() *apiMock.Pager[models.Driveable] {
	ds := []models.Driveable{}

	for _, dp := range en.DrivePagers {
		ds = append(ds, dp.Drive.Able)
	}

	return &apiMock.Pager[models.Driveable]{
		ToReturn: []apiMock.PagerResult[models.Driveable]{
			{Values: ds},
		},
	}
}

func (en EnumerateDriveItemsDelta) Drives() []*DeltaDrive {
	ds := []*DeltaDrive{}

	for _, dp := range en.DrivePagers {
		ds = append(ds, dp.Drive)
	}

	return ds
}

type DeltaDrive struct {
	ID   string
	Able models.Driveable
}

func Drive(driveSuffix ...any) *DeltaDrive {
	driveID := id("drive", driveSuffix...)

	able := models.NewDrive()
	able.SetId(ptr.To(driveID))
	able.SetName(ptr.To(name("drive", driveSuffix...)))

	return &DeltaDrive{
		ID:   driveID,
		Able: able,
	}
}

func (dd *DeltaDrive) NewEnumer() *DeltaDriveEnumerator {
	cp := &DeltaDrive{}

	*cp = *dd

	return &DeltaDriveEnumerator{Drive: cp}
}

type DeltaDriveEnumerator struct {
	Drive        *DeltaDrive
	idx          int
	DeltaQueries []*DeltaQuery
	Err          error
}

func (dde *DeltaDriveEnumerator) With(ds ...*DeltaQuery) *DeltaDriveEnumerator {
	dde.DeltaQueries = ds
	return dde
}

// WithErr adds an error that is always returned in the last delta index.
func (dde *DeltaDriveEnumerator) WithErr(err error) *DeltaDriveEnumerator {
	dde.Err = err
	return dde
}

func (dde *DeltaDriveEnumerator) nextDelta() *DeltaQuery {
	if dde.idx == len(dde.DeltaQueries) {
		// at the end of the enumeration, return an empty page with no items,
		// not even the root.  This is what graph api would do to signify an absence
		// of changes in the delta.
		lastDU := dde.DeltaQueries[dde.idx-1].DeltaUpdate

		return &DeltaQuery{
			DeltaUpdate: lastDU,
			Pages: []NextPage{{
				Items: []models.DriveItemable{},
			}},
			Err: dde.Err,
		}
	}

	if dde.idx > len(dde.DeltaQueries) {
		// a panic isn't optimal here, but since this mechanism is internal to testing,
		// it's an acceptable way to have the tests ensure we don't over-enumerate deltas.
		panic(fmt.Sprintf("delta index %d larger than count of delta iterations in mock", dde.idx))
	}

	pages := dde.DeltaQueries[dde.idx]

	dde.idx++

	return pages
}

var _ pagers.NextPageResulter[models.DriveItemable] = &DeltaQuery{}

type DeltaQuery struct {
	idx         int
	Pages       []NextPage
	DeltaUpdate pagers.DeltaUpdate
	Err         error
}

func Delta(
	resultDeltaID string,
	err error,
) *DeltaQuery {
	return &DeltaQuery{
		DeltaUpdate: pagers.DeltaUpdate{URL: resultDeltaID},
		Err:         err,
	}
}

func DeltaWReset(
	resultDeltaID string,
	err error,
) *DeltaQuery {
	return &DeltaQuery{
		DeltaUpdate: pagers.DeltaUpdate{
			URL:   resultDeltaID,
			Reset: true,
		},
		Err: err,
	}
}

func (dq *DeltaQuery) With(
	pages ...NextPage,
) *DeltaQuery {
	dq.Pages = pages
	return dq
}

func (dq *DeltaQuery) NextPage() ([]models.DriveItemable, bool, bool) {
	if dq.idx >= len(dq.Pages) {
		return nil, false, true
	}

	np := dq.Pages[dq.idx]
	dq.idx++

	return np.Items, np.Reset, false
}

func (dq *DeltaQuery) Cancel() {}

func (dq *DeltaQuery) Results() (pagers.DeltaUpdate, error) {
	return dq.DeltaUpdate, dq.Err
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

	DrivePagerV pagers.NonDeltaHandler[models.Driveable]

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

func (h RestoreHandler) NewDrivePager(string, []string) pagers.NonDeltaHandler[models.Driveable] {
	return h.DrivePagerV
}

func (h *RestoreHandler) AugmentItemInfo(
	details.ItemInfo,
	idname.Provider,
	*custom.DriveItem,
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

// assumption is only one suffix per id.  Mostly using
// the variadic as an "optional" extension.
func id(v string, suffixes ...any) string {
	id := fmt.Sprintf("id_%s", v)

	// a bit weird, but acts as a quality of life
	// that allows some funcs to take in the `file`
	// or `folder` or etc monikers as the suffix
	// without producing weird outputs.
	if len(suffixes) == 1 {
		sfx0, ok := suffixes[0].(string)
		if ok && sfx0 == v {
			return id
		}
	}

	for _, sfx := range suffixes {
		id = fmt.Sprintf("%s_%v", id, sfx)
	}

	return id
}

// assumption is only one suffix per name.  Mostly using
// the variadic as an "optional" extension.
func name(v string, suffixes ...any) string {
	name := fmt.Sprintf("n_%s", v)

	// a bit weird, but acts as a quality of life
	// that allows some funcs to take in the `file`
	// or `folder` or etc monikers as the suffix
	// without producing weird outputs.
	if len(suffixes) == 1 {
		sfx0, ok := suffixes[0].(string)
		if ok && sfx0 == v {
			return name
		}
	}

	for _, sfx := range suffixes {
		name = fmt.Sprintf("%s_%v", name, sfx)
	}

	return name
}
