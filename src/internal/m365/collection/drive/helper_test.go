package drive

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

const defaultFileSize int64 = 42

// TODO(ashmrtn): Merge with similar structs in graph and exchange packages.
type oneDriveService struct {
	credentials account.M365Config
	status      support.ControllerOperationStatus
	ac          api.Client
}

func newOneDriveService(credentials account.M365Config) (*oneDriveService, error) {
	ac, err := api.NewClient(
		credentials,
		control.DefaultOptions(),
		count.New())
	if err != nil {
		return nil, err
	}

	service := oneDriveService{
		ac:          ac,
		credentials: credentials,
	}

	return &service, nil
}

func (ods *oneDriveService) updateStatus(status *support.ControllerOperationStatus) {
	if status == nil {
		return
	}

	ods.status = support.MergeStatus(ods.status, *status)
}

func loadTestService(t *testing.T) *oneDriveService {
	a := tconfig.NewM365Account(t)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	service, err := newOneDriveService(creds)
	require.NoError(t, err, clues.ToCore(err))

	return service
}

// ---------------------------------------------------------------------------
// collections
// ---------------------------------------------------------------------------

type statePath struct {
	state    data.CollectionState
	currPath path.Path
	prevPath path.Path
}

func toODPath(t *testing.T, s string) path.Path {
	spl := path.Split(s)
	p, err := path.Builder{}.
		Append(spl[4:]...).
		ToDataLayerPath(
			spl[0],
			spl[2],
			path.OneDriveService,
			path.FilesCategory,
			false)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func asDeleted(t *testing.T, prev string) statePath {
	return statePath{
		state:    data.DeletedState,
		prevPath: toODPath(t, prev),
	}
}

func asMoved(t *testing.T, prev, curr string) statePath {
	return statePath{
		state:    data.MovedState,
		prevPath: toODPath(t, prev),
		currPath: toODPath(t, curr),
	}
}

func asNew(t *testing.T, curr string) statePath {
	return statePath{
		state:    data.NewState,
		currPath: toODPath(t, curr),
	}
}

func asNotMoved(t *testing.T, p string) statePath {
	return statePath{
		state:    data.NotMovedState,
		prevPath: toODPath(t, p),
		currPath: toODPath(t, p),
	}
}

// ---------------------------------------------------------------------------
// misc helpers
// ---------------------------------------------------------------------------

var anyFolderScope = (&selectors.OneDriveBackup{}).Folders(selectors.Any())[0]

type failingColl struct{}

func (f failingColl) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ic := make(chan data.Item)
	defer close(ic)

	errs.AddRecoverable(ctx, assert.AnError)

	return ic
}
func (f failingColl) FullPath() path.Path                                        { return nil }
func (f failingColl) FetchItemByName(context.Context, string) (data.Item, error) { return nil, nil }

func makeExcludeMap(files ...string) map[string]struct{} {
	delList := map[string]struct{}{}
	for _, file := range files {
		delList[file+metadata.DataFileSuffix] = struct{}{}
		delList[file+metadata.MetaFileSuffix] = struct{}{}
	}

	return delList
}

func defaultMetadataPath(t *testing.T) path.Path {
	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(t, err, "making default metadata path", clues.ToCore(err))

	return metadataPath
}

// ---------------------------------------------------------------------------
// limiter
// ---------------------------------------------------------------------------

func minimumLimitOpts() control.Options {
	minLimitOpts := control.DefaultOptions()
	minLimitOpts.PreviewLimits.Enabled = true
	minLimitOpts.PreviewLimits.MaxBytes = 1
	minLimitOpts.PreviewLimits.MaxContainers = 1
	minLimitOpts.PreviewLimits.MaxItems = 1
	minLimitOpts.PreviewLimits.MaxItemsPerContainer = 1
	minLimitOpts.PreviewLimits.MaxPages = 1

	return minLimitOpts
}

// ---------------------------------------------------------------------------
// enumerators
// ---------------------------------------------------------------------------

func collWithMBH(mbh BackupHandler) *Collections {
	return NewCollections(
		mbh,
		tenant,
		idname.NewProvider(user, user),
		func(*support.ControllerOperationStatus) {},
		control.Options{ToggleFeatures: control.Toggles{
			UseDeltaTree: true,
		}},
		count.New())
}

func collWithMBHAndOpts(
	mbh BackupHandler,
	opts control.Options,
) *Collections {
	return NewCollections(
		mbh,
		tenant,
		idname.NewProvider(user, user),
		func(*support.ControllerOperationStatus) {},
		opts,
		count.New())
}

func aPage(items ...models.DriveItemable) nextPage {
	return nextPage{
		Items: append([]models.DriveItemable{rootFolder()}, items...),
	}
}

func aPageWReset(items ...models.DriveItemable) nextPage {
	return nextPage{
		Items: append([]models.DriveItemable{rootFolder()}, items...),
		Reset: true,
	}
}

func aReset(items ...models.DriveItemable) nextPage {
	return nextPage{
		Items: []models.DriveItemable{},
		Reset: true,
	}
}

// ---------------------------------------------------------------------------
// metadata
// ---------------------------------------------------------------------------

func makePrevMetadataColls(
	t *testing.T,
	mbh BackupHandler,
	previousPaths map[string]map[string]string,
) []data.RestoreCollection {
	pathPrefix, err := mbh.MetadataPathPrefix(tenant)
	require.NoError(t, err, clues.ToCore(err))

	prevDeltas := map[string]string{}

	for driveID := range previousPaths {
		prevDeltas[driveID] = deltaURL("prev")
	}

	mdColl, err := graph.MakeMetadataCollection(
		pathPrefix,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(bupMD.DeltaURLsFileName, prevDeltas),
			graph.NewMetadataEntry(bupMD.PreviousPathFileName, previousPaths),
		},
		func(*support.ControllerOperationStatus) {},
		count.New())
	require.NoError(t, err, "creating metadata collection", clues.ToCore(err))

	return []data.RestoreCollection{
		dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mdColl}),
	}
}

func compareMetadata(
	t *testing.T,
	mdColl data.Collection,
	expectDeltas map[string]string,
	expectPrevPaths map[string]map[string]string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	colls := []data.RestoreCollection{
		dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mdColl}),
	}

	deltas, prevs, _, err := deserializeAndValidateMetadata(
		ctx,
		colls,
		count.New(),
		fault.New(true))
	require.NoError(t, err, "deserializing metadata", clues.ToCore(err))

	if expectDeltas != nil {
		assert.Equal(t, expectDeltas, deltas, "delta urls")
	}

	assert.Equal(t, expectPrevPaths, prevs, "previous paths")
}

// ---------------------------------------------------------------------------
// collections
// ---------------------------------------------------------------------------

// for comparisons done by a given collection path
type collectionAssertion struct {
	curr    path.Path
	prev    path.Path
	state   data.CollectionState
	fileIDs []string
	// should never get set by the user.
	// this flag gets flipped when calling assertions.compare.
	// any unseen collection will error on requireNoUnseenCollections
	sawCollection bool

	// used for metadata collection comparison
	deltas    map[string]string
	prevPaths map[string]map[string]string
}

func aColl(
	curr, prev path.Path,
	fileIDs ...string,
) *collectionAssertion {
	ids := make([]string, 0, 2*len(fileIDs))

	for _, fID := range fileIDs {
		ids = append(ids, fID+metadata.DataFileSuffix)
		ids = append(ids, fID+metadata.MetaFileSuffix)
	}

	// should expect all non-root, non-tombstone collections to contain
	// a dir meta file for storing permissions.
	if curr != nil && !strings.HasSuffix(curr.Folder(false), root) {
		ids = append(ids, metadata.DirMetaFileSuffix)
	}

	return &collectionAssertion{
		curr:    curr,
		prev:    prev,
		state:   data.StateOf(prev, curr, count.New()),
		fileIDs: ids,
	}
}

func aMetadata(
	deltas map[string]string,
	prevPaths map[string]map[string]string,
) *collectionAssertion {
	return &collectionAssertion{
		deltas:    deltas,
		prevPaths: prevPaths,
	}
}

// to aggregate all collection-related expectations in the backup
// map collection path -> collection state -> assertion
type expectedCollections struct {
	assertions  map[string]*collectionAssertion
	metadata    *collectionAssertion
	doNotMerge  assert.BoolAssertionFunc
	hasURLCache assert.ValueAssertionFunc
}

func expectCollections(
	doNotMerge bool,
	hasURLCache bool,
	colls ...*collectionAssertion,
) expectedCollections {
	var (
		as = map[string]*collectionAssertion{}
		md *collectionAssertion
	)

	for _, coll := range colls {
		if coll.prevPaths != nil {
			md = coll
			continue
		}

		as[expectFullOrPrev(coll).String()] = coll
	}

	dontMerge := assert.False
	if doNotMerge {
		dontMerge = assert.True
	}

	hasCache := assert.Nil
	if hasURLCache {
		hasCache = assert.NotNil
	}

	return expectedCollections{
		assertions:  as,
		metadata:    md,
		doNotMerge:  dontMerge,
		hasURLCache: hasCache,
	}
}

func (ecs expectedCollections) compare(
	t *testing.T,
	colls []data.BackupCollection,
) {
	for _, coll := range colls {
		ecs.compareColl(t, coll)
	}
}

func (ecs expectedCollections) compareColl(t *testing.T, coll data.BackupCollection) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		itemIDs = []string{}
		p       = fullOrPrevPath(t, coll)
	)

	// check the metadata collection separately
	if coll.FullPath() != nil && coll.FullPath().Equal(defaultMetadataPath(t)) {
		ecs.metadata.sawCollection = true
		compareMetadata(t, coll, ecs.metadata.deltas, ecs.metadata.prevPaths)

		return
	}

	if coll.State() != data.DeletedState {
		for itm := range coll.Items(ctx, fault.New(true)) {
			itemIDs = append(itemIDs, itm.ID())
		}
	}

	expect := ecs.assertions[p.String()]
	require.NotNil(
		t,
		expect,
		"test should have an expected entry for collection with:\n\tstate %q\n\tpath %q",
		coll.State(),
		p)

	expect.sawCollection = true

	assert.ElementsMatchf(
		t,
		expect.fileIDs,
		itemIDs,
		"expected all items to match in collection with:\n\tstate %q\n\tpath %q",
		coll.State(),
		p)

	if expect.prev == nil {
		assert.Nil(t, coll.PreviousPath(), "previous path")
	} else {
		assert.Equal(t, expect.prev, coll.PreviousPath())
	}

	if expect.curr == nil {
		assert.Nil(t, coll.FullPath(), "collection path")
	} else {
		assert.Equal(t, expect.curr, coll.FullPath())
	}

	ecs.doNotMerge(
		t,
		coll.DoNotMergeItems(),
		"expected collection to have the appropariate doNotMerge flag")

	driveColl := coll.(*Collection)

	ecs.hasURLCache(t, driveColl.urlCache, "has a populated url cache handler")
}

// ensure that no collections in the expected set are still flagged
// as sawCollection == false.
func (ecs expectedCollections) requireNoUnseenCollections(t *testing.T) {
	for _, ca := range ecs.assertions {
		require.True(
			t,
			ca.sawCollection,
			"results did not include collection at:\n\tstate %q\t\npath %q",
			ca.state, expectFullOrPrev(ca))
	}

	if ecs.metadata != nil {
		require.True(
			t,
			ecs.metadata.sawCollection,
			"results did not include the metadata collection")
	}
}

// ---------------------------------------------------------------------------
// delta trees
// ---------------------------------------------------------------------------

func defaultTreePfx(t *testing.T, d *deltaDrive) path.Path {
	return d.fullPath(t)
}

func defaultLoc() path.Elements {
	return path.NewElements("root:/foo/bar/baz/qux/fnords/smarf/voi/zumba/bangles/howdyhowdyhowdy")
}

func newTree(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	return newFolderyMcFolderFace(defaultTreePfx(t, d), rootID)
}

func treeWithRoot(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := newFolderyMcFolderFace(defaultTreePfx(t, d), rootID)
	root := custom.ToCustomDriveItem(rootFolder())

	//nolint:forbidigo
	err := tree.setFolder(context.Background(), root)
	require.NoError(t, err, clues.ToCore(err))

	return tree
}

func treeAfterReset(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := newFolderyMcFolderFace(defaultTreePfx(t, d), rootID)
	tree.reset()

	return tree
}

func treeWithFoldersAfterReset(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithFolders(t, d)
	tree.hadReset = true

	return tree
}

func treeWithTombstone(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithRoot(t, d)
	folder := custom.ToCustomDriveItem(d.folderAt(root))

	//nolint:forbidigo
	err := tree.setTombstone(context.Background(), folder)
	require.NoError(t, err, clues.ToCore(err))

	return tree
}

func treeWithFolders(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithRoot(t, d)
	parent := custom.ToCustomDriveItem(d.folderAt(root, "parent"))
	folder := custom.ToCustomDriveItem(d.folderAt("parent"))

	//nolint:forbidigo
	err := tree.setFolder(context.Background(), parent)
	require.NoError(t, err, clues.ToCore(err))

	//nolint:forbidigo
	err = tree.setFolder(context.Background(), folder)
	require.NoError(t, err, clues.ToCore(err))

	return tree
}

func treeWithFileAtRoot(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithRoot(t, d)

	f := custom.ToCustomDriveItem(d.fileAt(root))
	err := tree.addFile(f)
	require.NoError(t, err, clues.ToCore(err))

	return tree
}

func treeWithDeletedFile(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithRoot(t, d)
	tree.deleteFile(fileID("d"))

	return tree
}

func treeWithFileInFolder(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithFolders(t, d)

	f := custom.ToCustomDriveItem(d.fileAt(folder))
	err := tree.addFile(f)
	require.NoError(t, err, clues.ToCore(err))

	return tree
}

func treeWithFileInTombstone(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	tree := treeWithTombstone(t, d)

	// setting these directly, instead of using addFile(),
	// because we can't add files to tombstones.
	tree.tombstones[folderID()].files[fileID()] = custom.ToCustomDriveItem(d.fileAt("tombstone"))
	tree.fileIDToParentID[fileID()] = folderID()

	return tree
}

// root -> idx(folder, parent) -> folderID()
// one item at each dir
// one tombstone: idx(folder, tombstone)
// one item in the tombstone
// one deleted item
func fullTree(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	return fullTreeWithNames("parent", "tombstone")(t, d)
}

func fullTreeWithNames(
	parentFolderSuffix, tombstoneSuffix any,
) func(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
	return func(t *testing.T, d *deltaDrive) *folderyMcFolderFace {
		ctx, flush := tester.NewContext(t)
		defer flush()

		tree := treeWithRoot(t, d)

		// file "r" in root
		df := custom.ToCustomDriveItem(d.fileAt(root, "r"))
		err := tree.addFile(df)
		require.NoError(t, err, clues.ToCore(err))

		// root -> folderID(parentX)
		parent := custom.ToCustomDriveItem(d.folderAt(root, parentFolderSuffix))
		err = tree.setFolder(ctx, parent)
		require.NoError(t, err, clues.ToCore(err))

		// file "p" in folderID(parentX)
		df = custom.ToCustomDriveItem(d.fileAt(parentFolderSuffix, "p"))
		err = tree.addFile(df)
		require.NoError(t, err, clues.ToCore(err))

		// folderID(parentX) -> folderID()
		fld := custom.ToCustomDriveItem(d.folderAt(parentFolderSuffix))
		err = tree.setFolder(ctx, fld)
		require.NoError(t, err, clues.ToCore(err))

		// file "f" in folderID()
		df = custom.ToCustomDriveItem(d.fileAt(folder, "f"))
		err = tree.addFile(df)
		require.NoError(t, err, clues.ToCore(err))

		// tombstone - have to set a non-tombstone folder first,
		// then add the item,
		// then tombstone the folder
		tomb := custom.ToCustomDriveItem(d.folderAt(root, tombstoneSuffix))
		err = tree.setFolder(ctx, tomb)
		require.NoError(t, err, clues.ToCore(err))

		// file "t" in tombstone
		df = custom.ToCustomDriveItem(d.fileAt(tombstoneSuffix, "t"))
		err = tree.addFile(df)
		require.NoError(t, err, clues.ToCore(err))

		err = tree.setTombstone(ctx, tomb)
		require.NoError(t, err, clues.ToCore(err))

		// deleted file "d"
		tree.deleteFile(fileID("d"))

		return tree
	}
}

// ---------------------------------------------------------------------------
// Backup Handler
// ---------------------------------------------------------------------------

type mockBackupHandler[T any] struct {
	ItemInfo details.ItemInfo
	// FIXME: this is a hacky solution.  Better to use an interface
	// and plug in the selector scope there.
	Sel selectors.Selector

	DriveItemEnumeration enumerateDriveItemsDelta

	GI  getsItem
	GIP getsItemPermission

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

func defaultOneDriveBH(resourceOwner string) *mockBackupHandler[models.DriveItemable] {
	sel := selectors.NewOneDriveBackup([]string{resourceOwner})
	sel.Include(sel.AllData())

	return &mockBackupHandler[models.DriveItemable]{
		ItemInfo: details.ItemInfo{
			OneDrive:  &details.OneDriveInfo{},
			Extension: &details.ExtensionData{},
		},
		Sel:                  sel.Selector,
		DriveItemEnumeration: enumerateDriveItemsDelta{},
		GI:                   getsItem{Err: clues.New("not defined")},
		GIP:                  getsItemPermission{Err: clues.New("not defined")},
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

func defaultSharePointBH(resourceOwner string) *mockBackupHandler[models.DriveItemable] {
	sel := selectors.NewOneDriveBackup([]string{resourceOwner})
	sel.Include(sel.AllData())

	return &mockBackupHandler[models.DriveItemable]{
		ItemInfo: details.ItemInfo{
			SharePoint: &details.SharePointInfo{},
			Extension:  &details.ExtensionData{},
		},
		Sel:                  sel.Selector,
		GI:                   getsItem{Err: clues.New("not defined")},
		GIP:                  getsItemPermission{Err: clues.New("not defined")},
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

func defaultDriveBHWith(
	resource string,
	enumerator enumerateDriveItemsDelta,
) *mockBackupHandler[models.DriveItemable] {
	mbh := defaultOneDriveBH(resource)
	mbh.DriveItemEnumeration = enumerator

	return mbh
}

func (h mockBackupHandler[T]) PathPrefix(tID, driveID string) (path.Path, error) {
	pp, err := h.PathPrefixFn(tID, h.ProtectedResource.ID(), driveID)
	if err != nil {
		return nil, err
	}

	return pp, h.PathPrefixErr
}

func (h mockBackupHandler[T]) MetadataPathPrefix(tID string) (path.Path, error) {
	pp, err := h.MetadataPathPrefixFn(tID, h.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	return pp, h.MetadataPathPrefixErr
}

func (h mockBackupHandler[T]) CanonicalPath(pb *path.Builder, tID string) (path.Path, error) {
	cp, err := h.CanonPathFn(pb, tID, h.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	return cp, h.CanonPathErr
}

func (h mockBackupHandler[T]) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.Service, h.Category
}

func (h mockBackupHandler[T]) NewDrivePager(string, []string) pagers.NonDeltaHandler[models.Driveable] {
	return h.DriveItemEnumeration.drivePager()
}

func (h mockBackupHandler[T]) FormatDisplayPath(_ string, pb *path.Builder) string {
	return "/" + pb.String()
}

func (h mockBackupHandler[T]) NewLocationIDer(driveID string, elems ...string) details.LocationIDer {
	return h.LocationIDFn(driveID, elems...)
}

func (h mockBackupHandler[T]) AugmentItemInfo(
	details.ItemInfo,
	idname.Provider,
	*custom.DriveItem,
	int64,
	*path.Builder,
) details.ItemInfo {
	return h.ItemInfo
}

func (h *mockBackupHandler[T]) Get(context.Context, string, map[string]string) (*http.Response, error) {
	c := h.getCall
	h.getCall++

	// allows mockers to only populate the errors slice
	if h.GetErrs[c] != nil {
		return nil, h.GetErrs[c]
	}

	return h.GetResps[c], h.GetErrs[c]
}

func (h mockBackupHandler[T]) EnumerateDriveItemsDelta(
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

func (h mockBackupHandler[T]) GetItem(ctx context.Context, _, _ string) (models.DriveItemable, error) {
	return h.GI.GetItem(ctx, "", "")
}

func (h mockBackupHandler[T]) GetItemPermission(
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
	return details.NewSharePointLocationIDer(driveID, elems...)
}

func (h mockBackupHandler[T]) IsAllPass() bool {
	scope := h.Sel.Includes[0]
	return selectors.IsAnyTarget(selectors.SharePointScope(scope), selectors.SharePointLibraryFolder) ||
		selectors.IsAnyTarget(selectors.OneDriveScope(scope), selectors.OneDriveFolder)
}

func (h mockBackupHandler[T]) IncludesDir(dir string) bool {
	scope := h.Sel.Includes[0]
	return selectors.SharePointScope(scope).Matches(selectors.SharePointLibraryFolder, dir) ||
		selectors.OneDriveScope(scope).Matches(selectors.OneDriveFolder, dir)
}

func (h mockBackupHandler[T]) GetRootFolder(context.Context, string) (models.DriveItemable, error) {
	return h.RootFolder, nil
}

// ---------------------------------------------------------------------------
// Get Itemer
// ---------------------------------------------------------------------------

type getsItem struct {
	Item models.DriveItemable
	Err  error
}

func (m getsItem) GetItem(
	_ context.Context,
	_, _ string,
) (models.DriveItemable, error) {
	return m.Item, m.Err
}

// ---------------------------------------------------------------------------
// Drive Item Enummerator
// ---------------------------------------------------------------------------

type nextPage struct {
	Items []models.DriveItemable
	Reset bool
}

type enumerateDriveItemsDelta struct {
	DrivePagers map[string]*DeltaDriveEnumerator
}

func driveEnumerator(
	ds ...*DeltaDriveEnumerator,
) enumerateDriveItemsDelta {
	enumerator := enumerateDriveItemsDelta{
		DrivePagers: map[string]*DeltaDriveEnumerator{},
	}

	for _, drive := range ds {
		enumerator.DrivePagers[drive.Drive.id] = drive
	}

	return enumerator
}

func (en enumerateDriveItemsDelta) EnumerateDriveItemsDelta(
	_ context.Context,
	driveID, _ string,
	_ api.CallConfig,
) pagers.NextPageResulter[models.DriveItemable] {
	iterator := en.DrivePagers[driveID]
	return iterator.nextDelta()
}

func (en enumerateDriveItemsDelta) drivePager() *apiMock.Pager[models.Driveable] {
	dvs := []models.Driveable{}

	for _, dp := range en.DrivePagers {
		dvs = append(dvs, dp.Drive.able)
	}

	return &apiMock.Pager[models.Driveable]{
		ToReturn: []apiMock.PagerResult[models.Driveable]{
			{Values: dvs},
		},
	}
}

func (en enumerateDriveItemsDelta) getDrives() []*deltaDrive {
	dvs := []*deltaDrive{}

	for _, dp := range en.DrivePagers {
		dvs = append(dvs, dp.Drive)
	}

	return dvs
}

type deltaDrive struct {
	id   string
	able models.Driveable
}

func drive(driveSuffix ...any) *deltaDrive {
	driveID := id(drivePfx, driveSuffix...)

	able := models.NewDrive()
	able.SetId(ptr.To(driveID))
	able.SetName(ptr.To(name(drivePfx, driveSuffix...)))

	return &deltaDrive{
		id:   driveID,
		able: able,
	}
}

func (dd *deltaDrive) newEnumer() *DeltaDriveEnumerator {
	clone := &deltaDrive{}
	*clone = *dd

	return &DeltaDriveEnumerator{Drive: clone}
}

type drivePrevPaths struct {
	id                 string
	folderIDToPrevPath map[string]string
}

func (dd *deltaDrive) newPrevPaths(
	idPathPairs ...string,
) *drivePrevPaths {
	dpp := drivePrevPaths{
		id:                 dd.id,
		folderIDToPrevPath: map[string]string{},
	}

	if len(idPathPairs)%2 == 1 {
		dpp.folderIDToPrevPath["error"] = "idPathPairs had odd count of elements"
		return &dpp
	}

	for i := 0; i < len(idPathPairs); i += 2 {
		dpp.folderIDToPrevPath[idPathPairs[i]] = idPathPairs[i+1]
	}

	return &dpp
}

// transforms 0 or more drivePrevPaths to a map[driveID]map[folderID]prevPathString
func multiDrivePrevPaths(drivePrevs ...*drivePrevPaths) map[string]map[string]string {
	msmss := map[string]map[string]string{}

	for _, dp := range drivePrevs {
		msmss[dp.id] = dp.folderIDToPrevPath
	}

	return msmss
}

// transforms 0 or more drivePrevPaths to a map[driveID]map[folderID]prevPathString
func multiDriveMetadata(
	t *testing.T,
	drivePrevs ...*drivePrevPaths,
) []data.RestoreCollection {
	drc := []data.RestoreCollection{}

	for _, dp := range drivePrevs {
		mdColl := []graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(
				bupMD.DeltaURLsFileName,
				map[string]string{
					dp.id: deltaURL(),
				}),
			graph.NewMetadataEntry(
				bupMD.PreviousPathFileName,
				multiDrivePrevPaths(drivePrevs...)),
		}

		mc, err := graph.MakeMetadataCollection(
			defaultMetadataPath(t),
			mdColl,
			func(*support.ControllerOperationStatus) {},
			count.New())
		require.NoError(t, err, clues.ToCore(err))

		drc = append(drc, dataMock.NewUnversionedRestoreCollection(
			t,
			data.NoFetchRestoreCollection{Collection: mc}))
	}

	return drc
}

type DeltaDriveEnumerator struct {
	Drive        *deltaDrive
	idx          int
	DeltaQueries []*deltaQuery
	Err          error
}

func (dde *DeltaDriveEnumerator) with(ds ...*deltaQuery) *DeltaDriveEnumerator {
	dde.DeltaQueries = ds
	return dde
}

// withErr adds an error that is always returned in the last delta index.
func (dde *DeltaDriveEnumerator) withErr(err error) *DeltaDriveEnumerator {
	dde.Err = err
	return dde
}

func (dde *DeltaDriveEnumerator) nextDelta() *deltaQuery {
	if dde.idx == len(dde.DeltaQueries) {
		// at the end of the enumeration, return an empty page with no items,
		// not even the root.  This is what graph api would do to signify an absence
		// of changes in the delta.
		lastDU := dde.DeltaQueries[dde.idx-1].DeltaUpdate

		return &deltaQuery{
			DeltaUpdate: lastDU,
			Pages: []nextPage{{
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

var _ pagers.NextPageResulter[models.DriveItemable] = &deltaQuery{}

type deltaQuery struct {
	idx         int
	Pages       []nextPage
	DeltaUpdate pagers.DeltaUpdate
	Err         error
}

func delta(
	err error,
	deltaTokenSuffix ...any,
) *deltaQuery {
	return &deltaQuery{
		DeltaUpdate: pagers.DeltaUpdate{URL: deltaURL(deltaTokenSuffix...)},
		Err:         err,
	}
}

func deltaWReset(
	resultDeltaID string,
	err error,
) *deltaQuery {
	return &deltaQuery{
		DeltaUpdate: pagers.DeltaUpdate{
			URL:   resultDeltaID,
			Reset: true,
		},
		Err: err,
	}
}

func (dq *deltaQuery) with(
	pages ...nextPage,
) *deltaQuery {
	dq.Pages = pages
	return dq
}

func (dq *deltaQuery) NextPage() ([]models.DriveItemable, bool, bool) {
	if dq.idx >= len(dq.Pages) {
		return nil, false, true
	}

	np := dq.Pages[dq.idx]
	dq.idx++

	return np.Items, np.Reset, false
}

func (dq *deltaQuery) Cancel() {}

func (dq *deltaQuery) Results() (pagers.DeltaUpdate, error) {
	return dq.DeltaUpdate, dq.Err
}

// ---------------------------------------------------------------------------
// Get Item Permissioner
// ---------------------------------------------------------------------------

type getsItemPermission struct {
	Perm models.PermissionCollectionResponseable
	Err  error
}

func (m getsItemPermission) GetItemPermission(
	_ context.Context,
	_, _ string,
) (models.PermissionCollectionResponseable, error) {
	return m.Perm, m.Err
}

// ---------------------------------------------------------------------------
// Restore Handler
// --------------------------------------------------------------------------

type mockRestoreHandler struct {
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

func (h mockRestoreHandler) PostDrive(
	ctx context.Context,
	protectedResourceID, driveName string,
) (models.Driveable, error) {
	return h.PostDriveResp, h.PostDriveErr
}

func (h mockRestoreHandler) NewDrivePager(string, []string) pagers.NonDeltaHandler[models.Driveable] {
	return h.DrivePagerV
}

func (h *mockRestoreHandler) AugmentItemInfo(
	details.ItemInfo,
	idname.Provider,
	*custom.DriveItem,
	int64,
	*path.Builder,
) details.ItemInfo {
	return h.ItemInfo
}

func (h *mockRestoreHandler) GetItemsInContainerByCollisionKey(
	context.Context,
	string, string,
) (map[string]api.DriveItemIDType, error) {
	return h.CollisionKeyMap, nil
}

func (h *mockRestoreHandler) DeleteItem(
	_ context.Context,
	_, itemID string,
) error {
	h.CalledDeleteItem = true
	h.CalledDeleteItemOn = itemID

	return h.DeleteItemErr
}

func (h *mockRestoreHandler) DeleteItemPermission(
	context.Context,
	string, string, string,
) error {
	return nil
}

func (h *mockRestoreHandler) NewItemContentUpload(
	context.Context,
	string, string,
) (models.UploadSessionable, error) {
	return models.NewUploadSession(), h.UploadSessionErr
}

func (h *mockRestoreHandler) PostItemPermissionUpdate(
	context.Context,
	string, string,
	*drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	return drives.NewItemItemsItemInviteResponse(), nil
}

func (h *mockRestoreHandler) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	return nil, clues.New("not implemented")
}

func (h *mockRestoreHandler) PostItemInContainer(
	context.Context,
	string, string,
	models.DriveItemable,
	control.CollisionPolicy,
) (models.DriveItemable, error) {
	h.CalledPostItem = true
	return h.PostItemResp, h.PostItemErr
}

func (h *mockRestoreHandler) GetFolderByName(
	context.Context,
	string, string, string,
) (models.DriveItemable, error) {
	return models.NewDriveItem(), nil
}

func (h *mockRestoreHandler) GetRootFolder(
	context.Context,
	string,
) (models.DriveItemable, error) {
	return models.NewDriveItem(), nil
}

// ---------------------------------------------------------------------------
// stub drive item factories
// ---------------------------------------------------------------------------

type itemType int

const (
	isFile    itemType = 1
	isFolder  itemType = 2
	isPackage itemType = 3
)

func coreItem(
	id, name, parentPath, parentID string,
	it itemType,
) *models.DriveItem {
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&id)
	item.SetLastModifiedDateTime(ptr.To(time.Now()))

	parentReference := models.NewItemReference()
	parentReference.SetPath(&parentPath)
	parentReference.SetId(&parentID)
	item.SetParentReference(parentReference)

	switch it {
	case isFile:
		item.SetSize(ptr.To[int64](42))
		item.SetFile(models.NewFile())
	case isFolder:
		item.SetFolder(models.NewFolder())
	case isPackage:
		item.SetPackageEscaped(models.NewPackageEscaped())
	}

	return item
}

func driveItem(
	id, name, parentPath, parentID string,
	it itemType,
) models.DriveItemable {
	return coreItem(id, name, parentPath, parentID, it)
}

func driveItemWSize(
	id, name, parentPath, parentID string,
	size int64,
	it itemType,
) models.DriveItemable {
	res := coreItem(id, name, parentPath, parentID, it)
	res.SetSize(ptr.To(size))

	return res
}

func malwareItem(
	id, name, parentPath, parentID string,
	it itemType,
) models.DriveItemable {
	c := coreItem(id, name, parentPath, parentID, it)

	mal := models.NewMalware()
	malStr := "test malware"
	mal.SetDescription(&malStr)

	c.SetMalware(mal)

	return c
}

// delItem creates a DriveItemable that is marked as deleted. path must be set
// to the base drive path.
func delItem(
	id string,
	parentID string,
	it itemType,
) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetId(&id)
	item.SetDeleted(models.NewDeleted())

	parentReference := models.NewItemReference()
	parentReference.SetId(&parentID)
	item.SetParentReference(parentReference)

	switch it {
	case isFile:
		item.SetFile(models.NewFile())
	case isFolder:
		item.SetFolder(models.NewFolder())
	case isPackage:
		item.SetPackageEscaped(models.NewPackageEscaped())
	}

	return item
}

// ---------------------------------------------------------------------------
// file factories
// ---------------------------------------------------------------------------

func fileID(fileSuffixes ...any) string {
	return id(file, fileSuffixes...)
}

func fileName(fileSuffixes ...any) string {
	return name(file, fileSuffixes...)
}

func driveFile(
	parentPath, parentID string,
	fileSuffixes ...any,
) models.DriveItemable {
	return driveItem(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		parentPath,
		parentID,
		isFile)
}

func (dd *deltaDrive) fileAt(
	parentSuffix any,
	fileSuffixes ...any,
) models.DriveItemable {
	if parentSuffix == root {
		return driveItem(
			fileID(fileSuffixes...),
			fileName(fileSuffixes...),
			dd.dir(),
			rootID,
			isFile)
	}

	return driveItem(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		// the file's parent directory isn't used;
		// this parameter is an artifact of the driveItem
		// api and doesn't need to be populated for test
		// success.
		dd.dir(),
		folderID(parentSuffix),
		isFile)
}

func (dd *deltaDrive) fileWURLAtRoot(
	url string,
	isDeleted bool,
	fileSuffixes ...any,
) models.DriveItemable {
	di := driveFile(dd.dir(), rootID, fileSuffixes...)
	di.SetAdditionalData(map[string]any{
		"@microsoft.graph.downloadUrl": url,
	})

	if isDeleted {
		di.SetDeleted(models.NewDeleted())
	}

	return di
}

func (dd *deltaDrive) fileWSizeAt(
	size int64,
	parentSuffix any,
	fileSuffixes ...any,
) models.DriveItemable {
	if parentSuffix == root {
		return driveItemWSize(
			fileID(fileSuffixes...),
			fileName(fileSuffixes...),
			dd.dir(),
			rootID,
			size,
			isFile)
	}

	return driveItemWSize(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		dd.dir(),
		folderID(parentSuffix),
		size,
		isFile)
}

// ---------------------------------------------------------------------------
// folder factories
// ---------------------------------------------------------------------------

func folderID(folderSuffixes ...any) string {
	return id(folder, folderSuffixes...)
}

func folderName(folderSuffixes ...any) string {
	return name(folder, folderSuffixes...)
}

func driveFolder(
	parentPath, parentID string,
	folderSuffixes ...any,
) models.DriveItemable {
	return driveItem(
		folderID(folderSuffixes...),
		folderName(folderSuffixes...),
		parentPath,
		parentID,
		isFolder)
}

func rootFolder() models.DriveItemable {
	rootFolder := models.NewDriveItem()
	rootFolder.SetName(ptr.To(root))
	rootFolder.SetId(ptr.To(rootID))
	rootFolder.SetRoot(models.NewRoot())
	rootFolder.SetFolder(models.NewFolder())

	return rootFolder
}

func (dd *deltaDrive) folderAt(
	parentSuffix any,
	folderSuffixes ...any,
) models.DriveItemable {
	if parentSuffix == root {
		return driveItem(
			folderID(folderSuffixes...),
			folderName(folderSuffixes...),
			dd.dir(),
			rootID,
			isFolder)
	}

	return driveItem(
		folderID(folderSuffixes...),
		folderName(folderSuffixes...),
		// we should be putting in the full location here, not just the
		// parent suffix.  But that full location would be unused because
		// our unit tests don't utilize folder subselection (which is the
		// only reason we need to provide the dir).
		dd.dir(folderName(parentSuffix)),
		folderID(parentSuffix),
		isFolder)
}

func (dd *deltaDrive) packageAtRoot() models.DriveItemable {
	return driveItem(
		folderID(pkg),
		folderName(pkg),
		dd.dir(),
		rootID,
		isPackage)
}

// ---------------------------------------------------------------------------
// id, name, path factories
// ---------------------------------------------------------------------------

func deltaURL(suffixes ...any) string {
	if len(suffixes) > 1 {
		// this should fail any tests.  we could pass in a
		// testing.T instead and fail the call here, but that
		// produces a whole lot of chaff where this check should
		// still get us the expected failure
		return fmt.Sprintf(
			"too many suffixes in the URL; should only be 0 or 1, got %d",
			len(suffixes))
	}

	url := "https://delta.token.url"

	for _, sfx := range suffixes {
		url = fmt.Sprintf("%s?%v", url, sfx)
	}

	return url
}

// assumption is only one suffix per id.  Mostly using
// the variadic as an "optional" extension.
func id(v string, suffixes ...any) string {
	if len(suffixes) > 1 {
		// this should fail any tests.  we could pass in a
		// testing.T instead and fail the call here, but that
		// produces a whole lot of chaff where this check should
		// still get us the expected failure
		return fmt.Sprintf(
			"too many suffixes in the ID; should only be 0 or 1, got %d",
			len(suffixes))
	}

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
	if len(suffixes) > 1 {
		// this should fail any tests.  we could pass in a
		// testing.T instead and fail the call here, but that
		// produces a whole lot of chaff where this check should
		// still get us the expected failure
		return fmt.Sprintf(
			"too many suffixes in the Name; should only be 0 or 1, got %d",
			len(suffixes))
	}

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

func toPath(elems ...string) string {
	es := []string{}
	for _, elem := range elems {
		es = append(es, path.Split(elem)...)
	}

	switch len(es) {
	case 0:
		return ""
	case 1:
		return es[0]
	default:
		return path.Builder{}.Append(es...).String()
	}
}

// produces the full path for the provided drive
func (dd *deltaDrive) strPath(t *testing.T, elems ...string) string {
	return dd.fullPath(t, elems...).String()
}

func (dd *deltaDrive) fullPath(t *testing.T, elems ...string) path.Path {
	p, err := odConsts.DriveFolderPrefixBuilder(dd.id).
		Append(elems...).
		ToDataLayerPath(
			tenant,
			user,
			path.OneDriveService,
			path.FilesCategory,
			false)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

// produces a complete path prefix up to the drive root folder with any
// elements passed in appended to the generated prefix.
func (dd *deltaDrive) dir(elems ...string) string {
	return odConsts.DriveFolderPrefixBuilder(dd.id).
		Append(elems...).
		String()
}

// common item names
const (
	bar       = "bar"
	drivePfx  = "drive"
	fanny     = "fanny"
	file      = "file"
	folder    = "folder"
	foo       = "foo"
	item      = "item"
	malware   = "malware"
	nav       = "nav"
	pkg       = "package"
	rootID    = odConsts.RootID
	root      = odConsts.RootPathDir
	subfolder = "subfolder"
	tenant    = "t"
	user      = "u"
)

// ---------------------------------------------------------------------------
// misc
// ---------------------------------------------------------------------------

func expectFullOrPrev(ca *collectionAssertion) path.Path {
	var p path.Path

	if ca.state != data.DeletedState {
		p = ca.curr
	} else {
		p = ca.prev
	}

	return p
}

func fullOrPrevPath(
	t *testing.T,
	coll data.BackupCollection,
) path.Path {
	var collPath path.Path

	if coll.State() == data.DeletedState {
		collPath = coll.PreviousPath()
	} else {
		collPath = coll.FullPath()
	}

	require.NotNil(
		t,
		collPath,
		"full or prev path are nil for collection with state:\n\t%s",
		coll.State())

	require.False(
		t,
		len(collPath.Elements()) < 4,
		"malformed or missing collection path")

	return collPath
}
