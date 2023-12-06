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
	"github.com/alcionai/corso/src/internal/tester"
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
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
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

func fileAt(
	parentSuffix any,
	fileSuffixes ...any,
) models.DriveItemable {
	return driveItem(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		parentDir(folderName(parentSuffix)),
		folderID(parentSuffix),
		isFile)
}

func fileAtRoot(
	fileSuffixes ...any,
) models.DriveItemable {
	return driveItem(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		parentDir(),
		rootID,
		isFile)
}

func fileWURLAtRoot(
	url string,
	isDeleted bool,
	fileSuffixes ...any,
) models.DriveItemable {
	di := driveFile(parentDir(), rootID, fileSuffixes...)
	di.SetAdditionalData(map[string]any{
		"@microsoft.graph.downloadUrl": url,
	})

	if isDeleted {
		di.SetDeleted(models.NewDeleted())
	}

	return di
}

func fileWSizeAtRoot(
	size int64,
	fileSuffixes ...any,
) models.DriveItemable {
	return driveItemWSize(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		parentDir(),
		rootID,
		size,
		isFile)
}

func fileWSizeAt(
	size int64,
	parentSuffix any,
	fileSuffixes ...any,
) models.DriveItemable {
	return driveItemWSize(
		fileID(fileSuffixes...),
		fileName(fileSuffixes...),
		parentDir(folderName(parentSuffix)),
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

func driveRootFolder() models.DriveItemable {
	rootFolder := models.NewDriveItem()
	rootFolder.SetName(ptr.To(rootName))
	rootFolder.SetId(ptr.To(rootID))
	rootFolder.SetRoot(models.NewRoot())
	rootFolder.SetFolder(models.NewFolder())

	return rootFolder
}

func folderAtRoot(
	folderSuffixes ...any,
) models.DriveItemable {
	return driveItem(
		folderID(folderSuffixes...),
		folderName(folderSuffixes...),
		parentDir(),
		rootID,
		isFolder)
}

func folderAt(
	parentSuffix any,
	folderSuffixes ...any,
) models.DriveItemable {
	return driveItem(
		folderID(folderSuffixes...),
		folderName(folderSuffixes...),
		parentDir(folderName(parentSuffix)),
		folderID(parentSuffix),
		isFolder)
}

// ---------------------------------------------------------------------------
// id, name, path factories
// ---------------------------------------------------------------------------

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

func fullPathPath(t *testing.T, elems ...string) path.Path {
	p, err := path.FromDataLayerPath(fullPath(elems...), false)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func driveFullPath(driveID any, elems ...string) string {
	return toPath(append(
		[]string{
			tenant,
			path.OneDriveService.String(),
			user,
			path.FilesCategory.String(),
			odConsts.DriveFolderPrefixBuilder(id(drive, driveID)).String(),
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
		[]string{odConsts.DriveFolderPrefixBuilder(id(drive, driveID)).String()},
		elems...)...)
}

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

func aPage(items ...models.DriveItemable) mock.NextPage {
	return mock.NextPage{
		Items: append([]models.DriveItemable{driveRootFolder()}, items...),
	}
}

func aPageWReset(items ...models.DriveItemable) mock.NextPage {
	return mock.NextPage{
		Items: append([]models.DriveItemable{driveRootFolder()}, items...),
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
		prevDeltas[driveID] = id(delta, "prev")
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
}

func aColl(
	curr, prev path.Path,
	fileIDs ...string,
) *collectionAssertion {
	ids := make([]string, 0, 2*len(fileIDs))

	for _, fUD := range fileIDs {
		ids = append(ids, fUD+metadata.DataFileSuffix)
		ids = append(ids, fUD+metadata.MetaFileSuffix)
	}

	return &collectionAssertion{
		curr:    curr,
		prev:    prev,
		state:   data.StateOf(prev, curr, count.New()),
		fileIDs: ids,
	}
}

// to aggregate all collection-related expectations in the backup
// map collection path -> collection state -> assertion
type expectedCollections struct {
	assertions  map[string]*collectionAssertion
	doNotMerge  assert.BoolAssertionFunc
	hasURLCache assert.ValueAssertionFunc
}

func expectCollections(
	doNotMerge bool,
	hasURLCache bool,
	colls ...*collectionAssertion,
) expectedCollections {
	as := map[string]*collectionAssertion{}

	for _, coll := range colls {
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
}

// ---------------------------------------------------------------------------
// delta trees
// ---------------------------------------------------------------------------

func defaultTreePfx(t *testing.T) path.Path {
	fpb := fullPathPath(t).ToBuilder()
	fpe := fpb.Elements()
	fpe = fpe[:len(fpe)-1]
	fpb = path.Builder{}.Append(fpe...)

	p, err := path.FromDataLayerPath(fpb.String(), false)
	require.NoErrorf(
		t,
		err,
		"err processing path:\n\terr %+v\n\tpath %q",
		clues.ToCore(err),
		fpb)

	return p
}

func defaultLoc() path.Elements {
	return path.NewElements("root:/foo/bar/baz/qux/fnords/smarf/voi/zumba/bangles/howdyhowdyhowdy")
}

func newTree(t *testing.T) *folderyMcFolderFace {
	return newFolderyMcFolderFace(defaultTreePfx(t), rootID)
}

func treeWithRoot(t *testing.T) *folderyMcFolderFace {
	tree := newFolderyMcFolderFace(defaultTreePfx(t), rootID)
	rootey := newNodeyMcNodeFace(nil, rootID, rootName, false)
	tree.root = rootey
	tree.folderIDToNode[rootID] = rootey

	return tree
}

func treeAfterReset(t *testing.T) *folderyMcFolderFace {
	tree := newFolderyMcFolderFace(defaultTreePfx(t), rootID)
	tree.reset()

	return tree
}

func treeWithFoldersAfterReset(t *testing.T) *folderyMcFolderFace {
	tree := treeWithFolders(t)
	tree.hadReset = true

	return tree
}

func treeWithTombstone(t *testing.T) *folderyMcFolderFace {
	tree := treeWithRoot(t)
	tree.tombstones[folderID()] = newNodeyMcNodeFace(nil, folderID(), "", false)

	return tree
}

func treeWithFolders(t *testing.T) *folderyMcFolderFace {
	tree := treeWithRoot(t)

	parent := newNodeyMcNodeFace(tree.root, folderID("parent"), folderName("parent"), true)
	tree.folderIDToNode[parent.id] = parent
	tree.root.children[parent.id] = parent

	f := newNodeyMcNodeFace(parent, folderID(), folderName(), false)
	tree.folderIDToNode[f.id] = f
	parent.children[f.id] = f

	return tree
}

func treeWithFileAtRoot(t *testing.T) *folderyMcFolderFace {
	tree := treeWithRoot(t)
	tree.root.files[fileID()] = custom.ToCustomDriveItem(fileAtRoot())
	tree.fileIDToParentID[fileID()] = rootID

	return tree
}

func treeWithDeletedFile(t *testing.T) *folderyMcFolderFace {
	tree := treeWithRoot(t)
	tree.deleteFile(fileID("d"))

	return tree
}

func treeWithFileInFolder(t *testing.T) *folderyMcFolderFace {
	tree := treeWithFolders(t)
	tree.folderIDToNode[folderID()].files[fileID()] = custom.ToCustomDriveItem(fileAt(folder))
	tree.fileIDToParentID[fileID()] = folderID()

	return tree
}

func treeWithFileInTombstone(t *testing.T) *folderyMcFolderFace {
	tree := treeWithTombstone(t)
	tree.tombstones[folderID()].files[fileID()] = custom.ToCustomDriveItem(fileAt("tombstone"))
	tree.fileIDToParentID[fileID()] = folderID()

	return tree
}

// root -> idx(folder, parent) -> folderID()
// one item at each dir
// one tombstone: idx(folder, tombstone)
// one item in the tombstone
// one deleted item
func fullTree(t *testing.T) *folderyMcFolderFace {
	return fullTreeWithNames("parent", "tombstone")(t)
}

func fullTreeWithNames(
	parentFolderX, tombstoneX any,
) func(t *testing.T) *folderyMcFolderFace {
	return func(t *testing.T) *folderyMcFolderFace {
		ctx, flush := tester.NewContext(t)
		defer flush()

		tree := treeWithRoot(t)

		// file in root
		df := driveFile(parentDir(), rootID, "r")
		err := tree.addFile(
			rootID,
			fileID("r"),
			custom.ToCustomDriveItem(df))
		require.NoError(t, err, clues.ToCore(err))

		// root -> idx(folder, parent)
		err = tree.setFolder(ctx, rootID, folderID(parentFolderX), folderName(parentFolderX), false)
		require.NoError(t, err, clues.ToCore(err))

		// file in idx(folder, parent)
		df = driveFile(parentDir(folderName(parentFolderX)), folderID(parentFolderX), "p")
		err = tree.addFile(
			folderID(parentFolderX),
			fileID("p"),
			custom.ToCustomDriveItem(df))
		require.NoError(t, err, clues.ToCore(err))

		// idx(folder, parent) -> id(folder)
		err = tree.setFolder(ctx, folderID(parentFolderX), folderID(), folderName(), false)
		require.NoError(t, err, clues.ToCore(err))

		// file in id(folder)
		df = driveFile(parentDir(folderName()), folderID())
		err = tree.addFile(
			folderID(),
			fileID(),
			custom.ToCustomDriveItem(df))
		require.NoError(t, err, clues.ToCore(err))

		// tombstone - have to set a non-tombstone folder first, then add the item, then tombstone the folder
		err = tree.setFolder(ctx, rootID, folderID(tombstoneX), folderName(tombstoneX), false)
		require.NoError(t, err, clues.ToCore(err))

		// file in tombstone
		df = driveFile(parentDir(folderName(tombstoneX)), folderID(tombstoneX), "t")
		err = tree.addFile(
			folderID(tombstoneX),
			fileID("t"),
			custom.ToCustomDriveItem(df))
		require.NoError(t, err, clues.ToCore(err))

		err = tree.setTombstone(ctx, folderID(tombstoneX))
		require.NoError(t, err, clues.ToCore(err))

		// deleted file
		tree.deleteFile(fileID("d"))

		return tree
	}
}

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
