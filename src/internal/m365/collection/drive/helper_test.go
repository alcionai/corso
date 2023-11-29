package drive

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

const defaultItemSize int64 = 42

// TODO(ashmrtn): Merge with similar structs in graph and exchange packages.
type oneDriveService struct {
	credentials account.M365Config
	status      support.ControllerOperationStatus
	ac          api.Client
}

func NewOneDriveService(credentials account.M365Config) (*oneDriveService, error) {
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

	service, err := NewOneDriveService(creds)
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
// stub drive items
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

// func fileAtRoot() models.DriveItemable {
// 	return driveItem(id(file), name(file), parentDir(), rootID, isFile)
// }

func fileAt(
	parentX any,
) models.DriveItemable {
	pd := parentDir(namex(folder, parentX))
	pid := idx(folder, parentX)

	if parentX == folder {
		pd = parentDir(name(folder))
		pid = id(folder)
	}

	return driveItem(
		id(file),
		name(file),
		pd,
		pid,
		isFile)
}

func fileAtDeep(
	parentDir, parentID string,
) models.DriveItemable {
	return driveItem(
		id(file),
		name(file),
		parentDir,
		parentID,
		isFile)
}

func filexAtRoot(
	x any,
) models.DriveItemable {
	return driveItem(
		idx(file, x),
		namex(file, x),
		parentDir(),
		rootID,
		isFile)
}

func filexAt(
	x, parentX any,
) models.DriveItemable {
	pd := parentDir(namex(folder, parentX))
	pid := idx(folder, parentX)

	if parentX == folder {
		pd = parentDir(name(folder))
		pid = id(folder)
	}

	return driveItem(
		idx(file, x),
		namex(file, x),
		pd,
		pid,
		isFile)
}

func filexWSizeAtRoot(
	x any,
	size int64,
) models.DriveItemable {
	return driveItemWithSize(
		idx(file, x),
		namex(file, x),
		parentDir(),
		rootID,
		size,
		isFile)
}

func filexWSizeAt(
	x, parentX any,
	size int64,
) models.DriveItemable {
	pd := parentDir(namex(folder, parentX))
	pid := idx(folder, parentX)

	if parentX == folder {
		pd = parentDir(name(folder))
		pid = id(folder)
	}

	return driveItemWithSize(
		idx(file, x),
		namex(file, x),
		pd,
		pid,
		size,
		isFile)
}

func folderAtRoot() models.DriveItemable {
	return driveItem(id(folder), name(folder), parentDir(), rootID, isFolder)
}

func folderAtDeep(
	parentDir, parentID string,
) models.DriveItemable {
	return driveItem(
		id(folder),
		name(folder),
		parentDir,
		parentID,
		isFolder)
}

func folderxAt(
	x, parentX any,
) models.DriveItemable {
	pd := parentDir(namex(folder, parentX))
	pid := idx(folder, parentX)

	if parentX == folder {
		pd = parentDir(name(folder))
		pid = id(folder)
	}

	return driveItem(
		idx(folder, x),
		namex(folder, x),
		pd,
		pid,
		isFolder)
}

func folderxAtRoot(
	x any,
) models.DriveItemable {
	return driveItem(
		idx(folder, x),
		namex(folder, x),
		parentDir(),
		rootID,
		isFolder)
}

func driveItemWithSize(
	id, name, parentPath, parentID string,
	size int64,
	it itemType,
) models.DriveItemable {
	res := coreItem(id, name, parentPath, parentID, it)
	res.SetSize(ptr.To(size))

	return res
}

func fileItem(
	id, name, parentPath, parentID, url string,
	deleted bool,
) models.DriveItemable {
	di := driveItem(id, name, parentPath, parentID, isFile)
	di.SetAdditionalData(map[string]any{
		"@microsoft.graph.downloadUrl": url,
	})

	if deleted {
		di.SetDeleted(models.NewDeleted())
	}

	return di
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

func driveRootItem() models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(ptr.To(rootName))
	item.SetId(ptr.To(rootID))
	item.SetRoot(models.NewRoot())
	item.SetFolder(models.NewFolder())

	return item
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

func id(v string) string {
	return fmt.Sprintf("id_%s_0", v)
}

func idx(v string, sfx any) string {
	return fmt.Sprintf("id_%s_%v", v, sfx)
}

func name(v string) string {
	return fmt.Sprintf("n_%s_0", v)
}

func namex(v string, sfx any) string {
	return fmt.Sprintf("n_%s_%v", v, sfx)
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

func fullPath(elems ...string) string {
	return toPath(append(
		[]string{
			tenant,
			path.OneDriveService.String(),
			user,
			path.FilesCategory.String(),
			odConsts.DriveFolderPrefixBuilder(id(drive)).String(),
		},
		elems...)...)
}

func driveFullPath(driveID any, elems ...string) string {
	return toPath(append(
		[]string{
			tenant,
			path.OneDriveService.String(),
			user,
			path.FilesCategory.String(),
			odConsts.DriveFolderPrefixBuilder(idx(drive, driveID)).String(),
		},
		elems...)...)
}

func parentDir(elems ...string) string {
	return toPath(append(
		[]string{odConsts.DriveFolderPrefixBuilder(id(drive)).String()},
		elems...)...)
}

func driveParentDir(driveID any, elems ...string) string {
	return toPath(append(
		[]string{odConsts.DriveFolderPrefixBuilder(idx(drive, driveID)).String()},
		elems...)...)
}

// just for readability
const (
	doMergeItems    = true
	doNotMergeItems = false
)

// common item names
const (
	bar       = "bar"
	delta     = "delta_url"
	drive     = "drive"
	fanny     = "fanny"
	file      = "file"
	folder    = "folder"
	foo       = "foo"
	item      = "item"
	malware   = "malware"
	nav       = "nav"
	pkg       = "package"
	rootID    = odConsts.RootID
	rootName  = odConsts.RootPathDir
	subfolder = "subfolder"
	tenant    = "t"
	user      = "u"
)

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

// func fullOrPrevPath(
// 	t *testing.T,
// 	coll data.BackupCollection,
// ) path.Path {
// 	var collPath path.Path

// 	if coll.State() != data.DeletedState {
// 		collPath = coll.FullPath()
// 	} else {
// 		collPath = coll.PreviousPath()
// 	}

// 	require.False(
// 		t,
// 		len(collPath.Elements()) < 4,
// 		"malformed or missing collection path")

// 	return collPath
// }

func pagerForDrives(drives ...models.Driveable) *apiMock.Pager[models.Driveable] {
	return &apiMock.Pager[models.Driveable]{
		ToReturn: []apiMock.PagerResult[models.Driveable]{
			{Values: drives},
		},
	}
}

func makePrevMetadataColls(
	t *testing.T,
	mbh BackupHandler,
	previousPaths map[string]map[string]string,
) []data.RestoreCollection {
	pathPrefix, err := mbh.MetadataPathPrefix(tenant)
	require.NoError(t, err, clues.ToCore(err))

	prevDeltas := map[string]string{}

	for driveID := range previousPaths {
		prevDeltas[driveID] = idx(delta, "prev")
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

// func compareMetadata(
// 	t *testing.T,
// 	mdColl data.Collection,
// 	expectDeltas map[string]string,
// 	expectPrevPaths map[string]map[string]string,
// ) {
// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	colls := []data.RestoreCollection{
// 		dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mdColl}),
// 	}

// 	deltas, prevs, _, err := deserializeAndValidateMetadata(
// 		ctx,
// 		colls,
// 		count.New(),
// 		fault.New(true))
// 	require.NoError(t, err, "deserializing metadata", clues.ToCore(err))
// 	assert.Equal(t, expectDeltas, deltas, "delta urls")
// 	assert.Equal(t, expectPrevPaths, prevs, "previous paths")
// }

// for comparisons done by collection state
type stateAssertion struct {
	itemIDs []string
	// should never get set by the user.
	// this flag gets flipped when calling assertions.compare.
	// any unseen collection will error on requireNoUnseenCollections
	// sawCollection bool
}

// for comparisons done by a given collection path
type collectionAssertion struct {
	doNotMerge    assert.BoolAssertionFunc
	states        map[data.CollectionState]*stateAssertion
	excludedItems map[string]struct{}
}

type statesToItemIDs map[data.CollectionState][]string

// TODO(keepers): move excludeItems to a more global position.
func newCollAssertion(
	doNotMerge bool,
	itemsByState statesToItemIDs,
	excludeItems ...string,
) collectionAssertion {
	states := map[data.CollectionState]*stateAssertion{}

	for state, itemIDs := range itemsByState {
		states[state] = &stateAssertion{
			itemIDs: itemIDs,
		}
	}

	dnm := assert.False
	if doNotMerge {
		dnm = assert.True
	}

	return collectionAssertion{
		doNotMerge:    dnm,
		states:        states,
		excludedItems: makeExcludeMap(excludeItems...),
	}
}

// to aggregate all collection-related expectations in the backup
// map collection path -> collection state -> assertion
type collectionAssertions map[string]collectionAssertion

// ensure the provided collection matches expectations as set by the test.
// func (cas collectionAssertions) compare(
// 	t *testing.T,
// 	coll data.BackupCollection,
// 	excludes *prefixmatcher.StringSetMatchBuilder,
// ) {
// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	var (
// 		itemCh  = coll.Items(ctx, fault.New(true))
// 		itemIDs = []string{}
// 	)

// 	p := fullOrPrevPath(t, coll)

// 	for itm := range itemCh {
// 		itemIDs = append(itemIDs, itm.ID())
// 	}

// 	expect := cas[p.String()]
// 	expectState := expect.states[coll.State()]
// 	expectState.sawCollection = true

// 	assert.ElementsMatchf(
// 		t,
// 		expectState.itemIDs,
// 		itemIDs,
// 		"expected all items to match in collection with:\nstate %q\npath %q",
// 		coll.State(),
// 		p)

// 	expect.doNotMerge(
// 		t,
// 		coll.DoNotMergeItems(),
// 		"expected collection to have the appropariate doNotMerge flag")

// 	if result, ok := excludes.Get(p.String()); ok {
// 		assert.Equal(
// 			t,
// 			expect.excludedItems,
// 			result,
// 			"excluded items")
// 	}
// }

// ensure that no collections in the expected set are still flagged
// as sawCollection == false.
// func (cas collectionAssertions) requireNoUnseenCollections(
// 	t *testing.T,
// ) {
// 	for p, withPath := range cas {
// 		for _, state := range withPath.states {
// 			require.True(
// 				t,
// 				state.sawCollection,
// 				"results should have contained collection:\n\t%q\t\n%q",
// 				state, p)
// 		}
// 	}
// }

func aPage(items ...models.DriveItemable) mock.NextPage {
	return mock.NextPage{
		Items: append([]models.DriveItemable{driveRootItem()}, items...),
	}
}

func aPageWReset(items ...models.DriveItemable) mock.NextPage {
	return mock.NextPage{
		Items: append([]models.DriveItemable{driveRootItem()}, items...),
		Reset: true,
	}
}

func aReset(items ...models.DriveItemable) mock.NextPage {
	return mock.NextPage{
		Items: []models.DriveItemable{},
		Reset: true,
	}
}

// ---------------------------------------------------------------------------
// delta trees
// ---------------------------------------------------------------------------

var loc = path.NewElements("root:/foo/bar/baz/qux/fnords/smarf/voi/zumba/bangles/howdyhowdyhowdy")

func treeWithRoot() *folderyMcFolderFace {
	tree := newFolderyMcFolderFace(nil, rootID)
	rootey := newNodeyMcNodeFace(nil, rootID, rootName, false)
	tree.root = rootey
	tree.folderIDToNode[rootID] = rootey

	return tree
}

func treeWithTombstone() *folderyMcFolderFace {
	tree := treeWithRoot()
	tree.tombstones[id(folder)] = newNodeyMcNodeFace(nil, id(folder), "", false)

	return tree
}

func treeWithFolders() *folderyMcFolderFace {
	tree := treeWithRoot()

	parent := newNodeyMcNodeFace(tree.root, idx(folder, "parent"), namex(folder, "parent"), true)
	tree.folderIDToNode[parent.id] = parent
	tree.root.children[parent.id] = parent

	f := newNodeyMcNodeFace(parent, id(folder), name(folder), false)
	tree.folderIDToNode[f.id] = f
	parent.children[f.id] = f

	return tree
}

func treeWithFileAtRoot() *folderyMcFolderFace {
	tree := treeWithRoot()
	tree.root.files[id(file)] = fileyMcFileFace{
		lastModified: time.Now(),
		contentSize:  42,
	}
	tree.fileIDToParentID[id(file)] = rootID

	return tree
}

func treeWithFileInFolder() *folderyMcFolderFace {
	tree := treeWithFolders()
	tree.folderIDToNode[id(folder)].files[id(file)] = fileyMcFileFace{
		lastModified: time.Now(),
		contentSize:  42,
	}
	tree.fileIDToParentID[id(file)] = id(folder)

	return tree
}

func treeWithFileInTombstone() *folderyMcFolderFace {
	tree := treeWithTombstone()
	tree.tombstones[id(folder)].files[id(file)] = fileyMcFileFace{
		lastModified: time.Now(),
		contentSize:  42,
	}
	tree.fileIDToParentID[id(file)] = id(folder)

	return tree
}
