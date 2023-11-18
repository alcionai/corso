package drive

import (
	"context"
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	pmMock "github.com/alcionai/corso/src/internal/common/prefixmatcher/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// helpers
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

func driveRootItem(id string) models.DriveItemable {
	name := "root"
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&id)
	item.SetRoot(models.NewRoot())
	item.SetFolder(models.NewFolder())

	return item
}

// delItem creates a DriveItemable that is marked as deleted. path must be set
// to the base drive path.
func delItem(
	id string,
	parentPath string,
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

func fullPath(driveID any, elems ...string) string {
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

func parent(driveID any, elems ...string) string {
	return toPath(append(
		[]string{odConsts.DriveFolderPrefixBuilder(idx(drive, driveID)).String()},
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
	rootName  = "root"
	rootID    = "root_id"
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
// tests
// ---------------------------------------------------------------------------

type CollectionsUnitSuite struct {
	tester.Suite
}

func TestCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, &CollectionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionsUnitSuite) TestPopulateDriveCollections() {
	t := suite.T()

	tests := []struct {
		name                     string
		items                    []models.DriveItemable
		previousPaths            map[string]string
		topLevelPackages         map[string]struct{}
		scope                    selectors.OneDriveScope
		expect                   assert.ErrorAssertionFunc
		expectedCollectionIDs    map[string]statePath
		expectedItemCount        int
		expectedContainerCount   int
		expectedFileCount        int
		expectedSkippedCount     int
		expectedPrevPaths        map[string]string
		expectedExcludes         map[string]struct{}
		expectedTopLevelPackages map[string]struct{}
		expectedCountPackages    int
	}{
		{
			name: "Invalid item",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(item), name(item), parent(drive), rootID, -1),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, fullPath(drive)),
			},
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: fullPath(drive),
			},
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single File",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(file), name(file), parent(drive), rootID, isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, fullPath(drive)),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// Root folder is skipped since it's always present.
			expectedPrevPaths: map[string]string{
				rootID: fullPath(drive),
			},
			expectedExcludes:         makeExcludeMap(id(file)),
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asNew(t, fullPath(drive, name(folder))),
			},
			expectedPrevPaths: map[string]string{
				rootID:     fullPath(drive),
				id(folder): fullPath(drive, name(folder)),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder created twice", // deleted a created with same name in between a backup
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(idx(folder, 2), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:         asNotMoved(t, fullPath(drive)),
				idx(folder, 2): asNew(t, fullPath(drive, name(folder))),
			},
			expectedPrevPaths: map[string]string{
				rootID:         fullPath(drive),
				idx(folder, 2): fullPath(drive, name(folder)),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Package",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:  asNotMoved(t, fullPath(drive)),
				id(pkg): asNew(t, fullPath(drive, name(pkg))),
			},
			expectedPrevPaths: map[string]string{
				rootID:  fullPath(drive),
				id(pkg): fullPath(drive, name(pkg)),
			},
			expectedItemCount:      1,
			expectedContainerCount: 2,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				fullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 1,
		},
		{
			name: "Single Package with subfolder",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
				driveItem(id(folder), name(folder), parent(drive, name(pkg)), id(pkg), isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive, name(pkg)), id(pkg), isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, fullPath(drive)),
				id(pkg):       asNew(t, fullPath(drive, name(pkg))),
				id(folder):    asNew(t, fullPath(drive, name(pkg), name(folder))),
				id(subfolder): asNew(t, fullPath(drive, name(pkg), name(subfolder))),
			},
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(pkg):       fullPath(drive, name(pkg)),
				id(folder):    fullPath(drive, name(pkg), name(folder)),
				id(subfolder): fullPath(drive, name(pkg), name(subfolder)),
			},
			expectedItemCount:      3,
			expectedContainerCount: 4,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				fullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 3,
		},
		{
			name: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(idx(file, "inRoot"), namex(file, "inRoot"), parent(drive), rootID, isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
				driveItem(idx(file, "inFolder"), namex(file, "inFolder"), parent(drive, name(folder)), id(folder), isFile),
				driveItem(idx(file, "inPackage"), namex(file, "inPackage"), parent(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asNew(t, fullPath(drive, name(folder))),
				id(pkg):    asNew(t, fullPath(drive, name(pkg))),
			},
			expectedItemCount:      5,
			expectedFileCount:      3,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:     fullPath(drive),
				id(folder): fullPath(drive, name(folder)),
				id(pkg):    fullPath(drive, name(pkg)),
			},
			expectedTopLevelPackages: map[string]struct{}{
				fullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap(idx(file, "inRoot"), idx(file, "inFolder"), idx(file, "inPackage")),
		},
		{
			name: "contains folder selector",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(idx(file, "inRoot"), namex(file, "inRoot"), parent(drive), rootID, isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive, name(folder)), id(folder), isFolder),
				driveItem(idx(folder, 2), name(folder), parent(drive, name(folder), name(subfolder)), id(subfolder), isFolder),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
				driveItem(idx(file, "inFolder"), idx(file, "inFolder"), parent(drive, name(folder)), id(folder), isFile),
				driveItem(idx(file, "inFolder2"), namex(file, "inFolder2"), parent(drive, name(folder), name(subfolder), name(folder)), idx(folder, 2), isFile),
				driveItem(idx(file, "inFolderPackage"), namex(file, "inPackage"), parent(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{name(folder)})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(folder):     asNew(t, fullPath(drive, name(folder))),
				id(subfolder):  asNew(t, fullPath(drive, name(folder), name(subfolder))),
				idx(folder, 2): asNew(t, fullPath(drive, name(folder), name(subfolder), name(folder))),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			// just "folder" isn't added here because the include check is done on the
			// parent path since we only check later if something is a folder or not.
			expectedPrevPaths: map[string]string{
				id(folder):     fullPath(drive, name(folder)),
				id(subfolder):  fullPath(drive, name(folder), name(subfolder)),
				idx(folder, 2): fullPath(drive, name(folder), name(subfolder), name(folder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(idx(file, "inFolder"), idx(file, "inFolder2")),
		},
		{
			name: "prefix subfolder selector",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(idx(file, "inRoot"), namex(file, "inRoot"), parent(drive), rootID, isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive, name(folder)), id(folder), isFolder),
				driveItem(idx(folder, 2), name(folder), parent(drive, name(folder), name(subfolder)), id(subfolder), isFolder),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
				driveItem(idx(file, "inFolder"), idx(file, "inFolder"), parent(drive, name(folder)), id(folder), isFile),
				driveItem(idx(file, "inFolder2"), namex(file, "inFolder2"), parent(drive, name(folder), name(subfolder), name(folder)), idx(folder, 2), isFile),
				driveItem(idx(file, "inFolderPackage"), namex(file, "inPackage"), parent(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths: map[string]string{},
			scope: (&selectors.OneDriveBackup{}).Folders(
				[]string{toPath(name(folder), name(subfolder))},
				selectors.PrefixMatch())[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(subfolder):  asNew(t, fullPath(drive, name(folder), name(subfolder))),
				idx(folder, 2): asNew(t, fullPath(drive, name(folder), name(subfolder), name(folder))),
			},
			expectedItemCount:      3,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				id(subfolder):  fullPath(drive, name(folder), name(subfolder)),
				idx(folder, 2): fullPath(drive, name(folder), name(subfolder), name(folder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(idx(file, "inFolder2")),
		},
		{
			name: "match subfolder selector",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(file), name(file), parent(drive), rootID, isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive, name(folder)), id(folder), isFolder),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
				driveItem(idx(file, 1), namex(file, 1), parent(drive, name(folder)), id(folder), isFile),
				driveItem(idx(file, "inSubfolder"), namex(file, "inSubfolder"), parent(drive, name(folder), name(subfolder)), id(subfolder), isFile),
				driveItem(idx(file, 9), namex(file, 9), parent(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{toPath(name(folder), name(subfolder))})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(subfolder): asNew(t, fullPath(drive, name(folder), name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// No child folders for subfolder so nothing here.
			expectedPrevPaths: map[string]string{
				id(subfolder): fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(idx(file, "inSubfolder")),
		},
		{
			name: "not moved folder tree",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, name(folder)),
				id(subfolder): fullPath(drive, name(folder), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asNotMoved(t, fullPath(drive, name(folder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(folder):    fullPath(drive, name(folder)),
				id(subfolder): fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, namex(folder, "a")),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asMoved(t, fullPath(drive, namex(folder, "a")), fullPath(drive, name(folder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(folder):    fullPath(drive, name(folder)),
				id(subfolder): fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(idx(folder, 1), name(folder), parent(drive), rootID, isFolder),
				driveItem(idx(folder, 2), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				idx(folder, 1): fullPath(drive, namex(folder, "a")),
				id(subfolder):  fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:         asNotMoved(t, fullPath(drive)),
				idx(folder, 2): asNew(t, fullPath(drive, name(folder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:         fullPath(drive),
				idx(folder, 2): fullPath(drive, name(folder)),
				id(subfolder):  fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				delItem(id(folder), parent(drive), rootID, isFolder),
				driveItem(id(folder), name(drive), parent(drive), rootID, isFolder),
				delItem(id(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asDeleted(t, fullPath(drive, "")),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup including delete",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				delItem(id(folder), parent(drive), rootID, isFolder),
				driveItem(idx(folder, 2), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, namex(folder, "a")),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:         asNotMoved(t, fullPath(drive)),
				idx(folder, 2): asNew(t, fullPath(drive, name(folder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:         fullPath(drive),
				idx(folder, 2): fullPath(drive, name(folder)),
				id(subfolder):  fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup with addition",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(idx(folder, 1), name(folder), parent(drive), rootID, isFolder),
				delItem(idx(folder, 1), parent(drive), rootID, isFolder),
				driveItem(idx(folder, 2), name(folder), parent(drive), rootID, isFolder),
				delItem(idx(folder, 2), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				idx(folder, 1): fullPath(drive, namex(folder, "a")),
				id(subfolder):  fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, fullPath(drive)),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(subfolder): fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree with file no previous",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(file), name(file), parent(drive, name(folder)), id(folder), isFile),
				driveItem(id(folder), namex(folder, 2), parent(drive), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asNew(t, fullPath(drive, namex(folder, 2))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:     fullPath(drive),
				id(folder): fullPath(drive, namex(folder, 2)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(file)),
		},
		{
			name: "moved folder tree with file no previous 1",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(file), name(file), parent(drive, name(folder)), id(folder), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asNew(t, fullPath(drive, name(folder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:     fullPath(drive),
				id(folder): fullPath(drive, name(folder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(file)),
		},
		{
			name: "moved folder tree and subfolder 1",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, namex(folder, "a")),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, fullPath(drive)),
				id(folder):    asMoved(t, fullPath(drive, namex(folder, "a")), fullPath(drive, name(folder))),
				id(subfolder): asMoved(t, fullPath(drive, namex(folder, "a"), name(subfolder)), fullPath(drive, name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(folder):    fullPath(drive, name(folder)),
				id(subfolder): fullPath(drive, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree and subfolder 2",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(subfolder), name(subfolder), parent(drive), rootID, isFolder),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, namex(folder, "a")),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, fullPath(drive)),
				id(folder):    asMoved(t, fullPath(drive, namex(folder, "a")), fullPath(drive, name(folder))),
				id(subfolder): asMoved(t, fullPath(drive, namex(folder, "a"), name(subfolder)), fullPath(drive, name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(folder):    fullPath(drive, name(folder)),
				id(subfolder): fullPath(drive, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "move subfolder when moving parent",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(idx(folder, 2), namex(folder, 2), parent(drive), rootID, isFolder),
				driveItem(id(item), name(item), parent(drive, namex(folder, 2)), idx(folder, 2), isFile),
				// Need to see the parent folder first (expected since that's what Graph
				// consistently returns).
				driveItem(id(folder), namex(folder, "a"), parent(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive, namex(folder, "a")), id(folder), isFolder),
				driveItem(idx(item, 2), namex(item, 2), parent(drive, namex(folder, "a"), name(subfolder)), id(subfolder), isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, namex(folder, "a")),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:         asNotMoved(t, fullPath(drive)),
				idx(folder, 2): asNew(t, fullPath(drive, namex(folder, 2))),
				id(folder):     asMoved(t, fullPath(drive, namex(folder, "a")), fullPath(drive, name(folder))),
				id(subfolder):  asMoved(t, fullPath(drive, namex(folder, "a"), name(subfolder)), fullPath(drive, name(folder), name(subfolder))),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 4,
			expectedPrevPaths: map[string]string{
				rootID:         fullPath(drive),
				id(folder):     fullPath(drive, name(folder)),
				idx(folder, 2): fullPath(drive, namex(folder, 2)),
				id(subfolder):  fullPath(drive, name(folder), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(item), idx(item, 2)),
		},
		{
			name: "moved folder tree multiple times",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(file), name(file), parent(drive, name(folder)), id(folder), isFile),
				driveItem(id(folder), namex(folder, 2), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				id(folder):    fullPath(drive, namex(folder, "a")),
				id(subfolder): fullPath(drive, namex(folder, "a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asMoved(t, fullPath(drive, namex(folder, "a")), fullPath(drive, namex(folder, 2))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(folder):    fullPath(drive, namex(folder, 2)),
				id(subfolder): fullPath(drive, namex(folder, 2), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(file)),
		},
		{
			name: "deleted folder and package",
			items: []models.DriveItemable{
				driveRootItem(rootID), // root is always present, but not necessary here
				delItem(id(folder), parent(drive), rootID, isFolder),
				delItem(id(pkg), parent(drive), rootID, isPackage),
			},
			previousPaths: map[string]string{
				rootID:     fullPath(drive),
				id(folder): fullPath(drive, name(folder)),
				id(pkg):    fullPath(drive, name(pkg)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asDeleted(t, fullPath(drive, name(folder))),
				id(pkg):    asDeleted(t, fullPath(drive, name(pkg))),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: fullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder without previous",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				delItem(id(folder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				rootID: fullPath(drive),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, fullPath(drive)),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: fullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder tree move subfolder",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				delItem(id(folder), parent(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), parent(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				rootID:        fullPath(drive),
				id(folder):    fullPath(drive, name(folder)),
				id(subfolder): fullPath(drive, name(folder), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, fullPath(drive)),
				id(folder):    asDeleted(t, fullPath(drive, name(folder))),
				id(subfolder): asMoved(t, fullPath(drive, name(folder), name(subfolder)), fullPath(drive, name(subfolder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        fullPath(drive),
				id(subfolder): fullPath(drive, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete file",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				delItem(id(item), parent(drive), rootID, isFile),
			},
			previousPaths: map[string]string{
				rootID: fullPath(drive),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, fullPath(drive)),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: fullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(item)),
		},
		{
			name: "item before parent errors",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(file), name(file), parent(drive, name(folder)), id(folder), isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, fullPath(drive)),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: fullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "1 root file, 1 folder, 1 package, 1 good file, 1 malware",
			items: []models.DriveItemable{
				driveRootItem(rootID),
				driveItem(id(file), id(file), parent(drive), rootID, isFile),
				driveItem(id(folder), name(folder), parent(drive), rootID, isFolder),
				driveItem(id(pkg), name(pkg), parent(drive), rootID, isPackage),
				driveItem(idx(file, "good"), namex(file, "good"), parent(drive, name(folder)), id(folder), isFile),
				malwareItem(id(malware), name(malware), parent(drive, name(folder)), id(folder), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, fullPath(drive)),
				id(folder): asNew(t, fullPath(drive, name(folder))),
				id(pkg):    asNew(t, fullPath(drive, name(pkg))),
			},
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			expectedSkippedCount:   1,
			expectedPrevPaths: map[string]string{
				rootID:     fullPath(drive),
				id(folder): fullPath(drive, name(folder)),
				id(pkg):    fullPath(drive, name(pkg)),
			},
			expectedTopLevelPackages: map[string]struct{}{
				fullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap(id(file), idx(file, "good")),
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				driveID = idx(drive, drive)
				mbh     = mock.DefaultOneDriveBH(user)
				du      = pagers.DeltaUpdate{
					URL:   "notempty",
					Reset: false,
				}
				excludes = map[string]struct{}{}
				errs     = fault.New(true)
			)

			mbh.DriveItemEnumeration = mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID: {
						Pages:       []mock.NextPage{{Items: test.items}},
						DeltaUpdate: du,
					},
				},
			}

			sel := selectors.NewOneDriveBackup([]string{user})
			sel.Include([]selectors.OneDriveScope{test.scope})

			mbh.Sel = sel.Selector

			c := NewCollections(
				mbh,
				tenant,
				idname.NewProvider(user, user),
				nil,
				control.Options{ToggleFeatures: control.Toggles{}},
				count.New())

			c.CollectionMap[driveID] = map[string]*Collection{}

			_, newPrevPaths, err := c.PopulateDriveCollections(
				ctx,
				driveID,
				"General",
				test.previousPaths,
				excludes,
				test.topLevelPackages,
				"prevdelta",
				count.New(),
				errs)
			test.expect(t, err, clues.ToCore(err))
			assert.ElementsMatch(
				t,
				maps.Keys(test.expectedCollectionIDs),
				maps.Keys(c.CollectionMap[driveID]),
				"expected collection IDs")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			assert.Equal(t, test.expectedSkippedCount, len(errs.Skipped()), "skipped item count")

			for id, sp := range test.expectedCollectionIDs {
				if !assert.Containsf(t, c.CollectionMap[driveID], id, "missing collection with id %s", id) {
					// Skip collections we don't find so we don't get an NPE.
					continue
				}

				assert.Equalf(t, sp.state, c.CollectionMap[driveID][id].State(), "state for collection %s", id)
				assert.Equalf(t, sp.currPath, c.CollectionMap[driveID][id].FullPath(), "current path for collection %s", id)
				assert.Equalf(t, sp.prevPath, c.CollectionMap[driveID][id].PreviousPath(), "prev path for collection %s", id)
			}

			assert.Equal(t, test.expectedPrevPaths, newPrevPaths, "previous paths")
			assert.Equal(t, test.expectedTopLevelPackages, test.topLevelPackages, "top level packages")
			assert.Equal(t, test.expectedExcludes, excludes, "excluded item IDs map")

			var countPackages int

			for _, drives := range c.CollectionMap {
				for _, col := range drives {
					if col.isPackageOrChildOfPackage {
						countPackages++
					}
				}
			}

			assert.Equal(t, test.expectedCountPackages, countPackages, "count of collections marked as packages")
		})
	}
}

func (suite *CollectionsUnitSuite) TestDeserializeMetadata() {
	table := []struct {
		name string
		// Each function returns the set of files for a single data.Collection.
		cols                 []func() []graph.MetadataCollectionEntry
		expectedDeltas       map[string]string
		expectedPaths        map[string]map[string]string
		expectedAlerts       []string
		canUsePreviousBackup bool
		errCheck             assert.ErrorAssertionFunc
	}{
		{
			name: "SuccessOneDriveAllOneCollection",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive): id(delta),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			name: "MissingPaths",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
					}
				},
			},
			expectedDeltas:       map[string]string{},
			expectedPaths:        map[string]map[string]string{},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			name: "MissingDeltas",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			// An empty path map but valid delta results in metadata being returned
			// since it's possible to have a drive with no folders other than the
			// root.
			name: "EmptyPaths",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {},
							}),
					}
				},
			},
			expectedDeltas:       map[string]string{},
			expectedPaths:        map[string]map[string]string{id(drive): {}},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			// An empty delta map but valid path results in no metadata for that drive
			// being returned since the path map is only useful if we have a valid
			// delta.
			name: "EmptyDeltas",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{
								id(drive): "",
							}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{id(drive): ""},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			name: "SuccessTwoDrivesTwoCollections",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{idx(drive, 2): idx(delta, 2)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								idx(drive, 2): {
									idx(folder, 2): fullPath(2),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive):     id(delta),
				idx(drive, 2): idx(delta, 2),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
				},
				idx(drive, 2): {
					idx(folder, 2): fullPath(2),
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			// Bad formats are logged but skip adding entries to the maps and don't
			// return an error.
			name:           "BadFormat",
			expectedDeltas: map[string]string{},
			expectedPaths:  map[string]map[string]string{},
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]string{id(drive): id(delta)}),
					}
				},
			},
			canUsePreviousBackup: false,
			errCheck:             assert.NoError,
		},
		{
			// Unexpected files are logged and skipped. They don't cause an error to
			// be returned.
			name: "BadFileName",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
						graph.NewMetadataEntry(
							"foo",
							map[string]string{id(drive): id(delta)}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive): id(delta),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			name: "DriveAlreadyFound_Paths",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 2): fullPath(2),
								},
							}),
					}
				},
			},
			expectedDeltas:       map[string]string{},
			expectedPaths:        map[string]map[string]string{},
			canUsePreviousBackup: false,
			errCheck:             assert.NoError,
		},
		{
			name: "DriveAlreadyFound_Deltas",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): idx(delta, 2)}),
					}
				},
			},
			expectedDeltas:       map[string]string{},
			expectedPaths:        map[string]map[string]string{},
			canUsePreviousBackup: false,
			errCheck:             assert.NoError,
		},
		{
			name: "DuplicatePreviousPaths",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
									idx(folder, 2): fullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive): id(delta),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
					idx(folder, 2): fullPath(1),
				},
			},
			expectedAlerts:       []string{fault.AlertPreviousPathCollision},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
		{
			name: "DuplicatePreviousPaths_separateDrives",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{
								id(drive): id(delta),
							}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive): {
									idx(folder, 1): fullPath(1),
									idx(folder, 2): fullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{idx(drive, 2): idx(delta, 2)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								idx(drive, 2): {
									idx(folder, 1): fullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive):     id(delta),
				idx(drive, 2): idx(delta, 2),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					idx(folder, 1): fullPath(1),
					idx(folder, 2): fullPath(1),
				},
				idx(drive, 2): {
					idx(folder, 1): fullPath(1),
				},
			},
			expectedAlerts:       []string{fault.AlertPreviousPathCollision},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			cols := []data.RestoreCollection{}

			for _, c := range test.cols {
				pathPrefix, err := path.BuildMetadata(
					tenant,
					user,
					path.OneDriveService,
					path.FilesCategory,
					false)
				require.NoError(t, err, clues.ToCore(err))

				mc, err := graph.MakeMetadataCollection(
					pathPrefix,
					c(),
					func(*support.ControllerOperationStatus) {},
					count.New())
				require.NoError(t, err, clues.ToCore(err))

				cols = append(cols, dataMock.NewUnversionedRestoreCollection(
					t,
					data.NoFetchRestoreCollection{Collection: mc}))
			}

			fb := fault.New(true)

			deltas, paths, canUsePreviousBackup, err := deserializeAndValidateMetadata(ctx, cols, count.New(), fb)
			test.errCheck(t, err)
			assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup, "can use previous backup")

			assert.Equal(t, test.expectedDeltas, deltas, "deltas")
			assert.Equal(t, test.expectedPaths, paths, "paths")

			alertMsgs := []string{}

			for _, alert := range fb.Alerts() {
				alertMsgs = append(alertMsgs, alert.Message)
			}

			assert.ElementsMatch(t, test.expectedAlerts, alertMsgs, "alert messages")
		})
	}
}

// This check is to ensure that we don't error out, but still return
// canUsePreviousBackup as false on read errors
func (suite *CollectionsUnitSuite) TestDeserializeMetadata_ReadFailure() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	fc := failingColl{}

	_, _, canUsePreviousBackup, err := deserializeAndValidateMetadata(ctx, []data.RestoreCollection{fc}, count.New(), fault.New(true))
	require.NoError(t, err)
	require.False(t, canUsePreviousBackup)
}

func (suite *CollectionsUnitSuite) TestGet() {
	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(idx(drive, 1)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	drive2 := models.NewDrive()
	drive2.SetId(ptr.To(idx(drive, 2)))
	drive2.SetName(ptr.To(namex(drive, 2)))

	table := []struct {
		name                 string
		drives               []models.Driveable
		enumerator           mock.EnumerateItemsDeltaByDrive
		canUsePreviousBackup bool
		errCheck             assert.ErrorAssertionFunc
		previousPaths        map[string]map[string]string
		// Collection name -> set of item IDs. We can't check item data because
		// that's not mocked out. Metadata is checked separately.
		expectedCollections   map[string]map[data.CollectionState][]string
		expectedDeltaURLs     map[string]string
		expectedPreviousPaths map[string]map[string]string
		// Items that should be excluded from the base. Only populated if the delta
		// was valid and there was at least 1 previous folder path.
		expectedDelList      *pmMock.PrefixMap
		expectedSkippedCount int
		// map full or previous path (prefers full) -> bool
		doNotMergeItems map[string]bool
	}{
		{
			name:   "OneDrive_OneItemPage_DelFileOnly_NoFolders_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID), // will be present, not needed
								delItem(id(file), parent(1), rootID, isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {rootID: fullPath(1)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NotMovedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {rootID: fullPath(1)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(id(file)),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_NoFolderDeltas_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(id(file), name(file), parent(1), rootID, isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {rootID: fullPath(1)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NotMovedState: {id(file)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {rootID: fullPath(1)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(id(file)),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
								driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors_FileRenamedMultiple",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
								driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								driveItem(id(file), namex(file, 2), parent(1, name(folder)), id(folder), isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors_FileMovedMultiple",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
							driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
							driveItem(id(file), namex(file, 2), parent(1), rootID, isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NotMovedState: {id(file)}},
				fullPath(1, name(folder)): {data.NewState: {id(folder)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(id(file)),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_EmptyDelta_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
							driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: "", Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file)}},
			},
			expectedDeltaURLs: map[string]string{},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_TwoItemPages_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 2), namex(file, 2), parent(1, name(folder)), id(folder), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file), idx(file, 2)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_TwoItemPages_WithReset",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 2), namex(file, 2), parent(1, name(folder)), id(folder), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file), idx(file, 2)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_TwoItemPages_WithResetCombinedWithItems",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 2), namex(file, 2), parent(1, name(folder)), id(folder), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file), idx(file, 2)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name: "TwoDrives_OneItemPageEach_NoErrors",
			drives: []models.Driveable{
				drive1,
				drive2,
			},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
							driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
					idx(drive, 2): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(idx("root", 2)),
							driveItem(idx(folder, 2), name(folder), parent(2), idx("root", 2), isFolder),
							driveItem(idx(file, 2), name(file), parent(2, name(folder)), idx(folder, 2), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
				idx(drive, 2): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file)}},
				fullPath(2):               {data.NewState: {}},
				fullPath(2, name(folder)): {data.NewState: {idx(folder, 2), idx(file, 2)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
				idx(drive, 2): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
				idx(drive, 2): {
					idx("root", 2): fullPath(2),
					idx(folder, 2): fullPath(2, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
				fullPath(2):               true,
				fullPath(2, name(folder)): true,
			},
		},
		{
			name: "TwoDrives_DuplicateIDs_OneItemPageEach_NoErrors",
			drives: []models.Driveable{
				drive1,
				drive2,
			},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
							driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
					idx(drive, 2): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(id(folder), name(folder), parent(2), rootID, isFolder),
							driveItem(idx(file, 2), name(file), parent(2, name(folder)), id(folder), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
				idx(drive, 2): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file)}},
				fullPath(2):               {data.NewState: {}},
				fullPath(2, name(folder)): {data.NewState: {id(folder), idx(file, 2)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
				idx(drive, 2): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
				idx(drive, 2): {
					rootID:     fullPath(2),
					id(folder): fullPath(2, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
				fullPath(2):               true,
				fullPath(2, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_Errors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages:       []mock.NextPage{{Items: []models.DriveItemable{}}},
						DeltaUpdate: pagers.DeltaUpdate{},
						Err:         assert.AnError,
					},
				},
			},
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections:   nil,
			expectedDeltaURLs:     nil,
			expectedPreviousPaths: nil,
			expectedDelList:       nil,
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_DeleteNonExistentFolder",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):                   {data.NewState: {}},
				fullPath(1, name(folder)):     {data.DeletedState: {}},
				fullPath(1, namex(folder, 2)): {data.NewState: {idx(folder, 2), id(file)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 2): fullPath(1, namex(folder, 2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):                   true,
				fullPath(1, name(folder)):     true,
				fullPath(1, namex(folder, 2)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDeltaCombinedWithItems_DeleteNonExistentFolder",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):                   {data.NewState: {}},
				fullPath(1, name(folder)):     {data.DeletedState: {}},
				fullPath(1, namex(folder, 2)): {data.NewState: {idx(folder, 2), id(file)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 2): fullPath(1, namex(folder, 2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):                   true,
				fullPath(1, name(folder)):     true,
				fullPath(1, namex(folder, 2)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								// on the first page, if this is the total data, we'd expect both folder and folder2
								// since new previousPaths merge with the old previousPaths.
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 2), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), idx(folder, 2), isFile),
								},
							},
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								// but after a delta reset, we treat this as the total end set of folders, which means
								// we don't expect folder to exist any longer.
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 2), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
				fullPath(1, name(folder)): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {idx(folder, 2), id(file)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 2): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtExistingLocation",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
				fullPath(1, name(folder)): {
					data.NewState: {id(folder), id(file)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_ImmediateInvalidPrevDelta_MoveFolderToPreviouslyExistingPath",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 2), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 2), name(file), parent(1, name(folder)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
				fullPath(1, name(folder)): {
					data.DeletedState: {},
					data.NewState:     {idx(folder, 2), idx(file, 2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 2): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 2), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
				fullPath(1, name(folder)): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {idx(folder, 2), id(file)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 2): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "OneDrive Two Item Pages with Malware",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
									malwareItem(id(malware), name(malware), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 2), namex(file, 2), parent(1, name(folder)), id(folder), isFile),
									malwareItem(idx(malware, 2), namex(malware, 2), parent(1, name(folder)), id(folder), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder), id(file), idx(file, 2)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
			expectedSkippedCount: 2,
		},
		{
			name:   "One Drive Deleted Folder In New Results With Invalid Delta",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
									driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
									driveItem(idx(file, 2), namex(file, 2), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								},
							},
							{
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
									delItem(idx(folder, 2), parent(1), rootID, isFolder),
									delItem(namex(file, 2), parent(1), rootID, isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					id(folder):     fullPath(1, name(folder)),
					idx(folder, 2): fullPath(1, namex(folder, 2)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):                   {data.NewState: {}},
				fullPath(1, name(folder)):     {data.NewState: {id(folder), id(file)}},
				fullPath(1, namex(folder, 2)): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):                   true,
				fullPath(1, name(folder)):     true,
				fullPath(1, namex(folder, 2)): true,
			},
		},
		{
			name:   "One Drive Folder Delete After Invalid Delta",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								delItem(id(folder), parent(1), rootID, isFolder),
							},
							Reset: true,
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "One Drive Item Delete After Invalid Delta",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(file), parent(1), rootID, isFile),
								},
								Reset: true,
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1): true,
			},
		},
		{
			name:   "One Drive Folder Made And Deleted",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(folder), parent(1), rootID, isFolder),
									delItem(id(file), parent(1), rootID, isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1): true,
			},
		},
		{
			name:   "One Drive Folder Created -> Deleted -> Created",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(folder), parent(1), rootID, isFolder),
									delItem(id(file), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 1), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 1), name(file), parent(1, name(folder)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {idx(folder, 1), idx(file, 1)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 1): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "One Drive Folder Deleted -> Created -> Deleted",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(folder), parent(1), rootID, isFolder),
									delItem(id(file), parent(1, name(folder)), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(folder), parent(1), rootID, isFolder),
									delItem(id(file), parent(1, name(folder)), rootID, isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NotMovedState: {}},
				fullPath(1, name(folder)): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name:   "One Drive Folder Created -> Deleted -> Created with prev",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(folder), parent(1), rootID, isFolder),
									delItem(id(file), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(idx(folder, 1), name(folder), parent(1), rootID, isFolder),
									driveItem(idx(file, 1), name(file), parent(1, name(folder)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.DeletedState: {}, data.NewState: {idx(folder, 1), idx(file, 1)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					idx(folder, 1): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               false,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "One Drive Item Made And Deleted",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
									driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID),
									delItem(id(file), parent(1), rootID, isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1):               {data.NewState: {}},
				fullPath(1, name(folder)): {data.NewState: {id(folder)}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:     fullPath(1),
					id(folder): fullPath(1, name(folder)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1):               true,
				fullPath(1, name(folder)): true,
			},
		},
		{
			name:   "One Drive Random Folder Delete",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							delItem(id(folder), parent(1), rootID, isFolder),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1): true,
			},
		},
		{
			name:   "One Drive Random Item Delete",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							delItem(id(file), parent(1), rootID, isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID: fullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(1): true,
			},
		},
		{
			name:   "TwoPriorDrives_OneTombstoned",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID), // will be present
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {rootID: fullPath(1)},
				idx(drive, 2): {rootID: fullPath(2)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {data.NotMovedState: {}},
				fullPath(2): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{idx(drive, 1): id(delta)},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {rootID: fullPath(1)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				fullPath(2): true,
			},
		},
		{
			name:   "duplicate previous paths in metadata",
			drives: []models.Driveable{drive1, drive2},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					// contains duplicates in previousPath
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
								driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
								driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
								driveItem(idx(file, 2), namex(file, 2), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
					// does not contain duplicates
					idx(drive, 2): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(id(folder), name(folder), parent(2), rootID, isFolder),
								driveItem(id(file), name(file), parent(2, name(folder)), id(folder), isFile),
								driveItem(idx(folder, 2), namex(folder, 2), parent(2), rootID, isFolder),
								driveItem(idx(file, 2), namex(file, 2), parent(2, namex(folder, 2)), idx(folder, 2), isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					id(folder):     fullPath(1, name(folder)),
					idx(folder, 2): fullPath(1, name(folder)),
					idx(folder, 3): fullPath(1, name(folder)),
				},
				idx(drive, 2): {
					rootID:         fullPath(2),
					id(folder):     fullPath(2, name(folder)),
					idx(folder, 2): fullPath(2, namex(folder, 2)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {
					data.NewState: {id(folder), idx(folder, 2)},
				},
				fullPath(1, name(folder)): {
					data.NotMovedState: {id(folder), id(file)},
				},
				fullPath(1, namex(folder, 2)): {
					data.MovedState: {idx(folder, 2), idx(file, 2)},
				},
				fullPath(2): {
					data.NewState: {id(folder), idx(folder, 2)},
				},
				fullPath(2, name(folder)): {
					data.NotMovedState: {id(folder), id(file)},
				},
				fullPath(2, namex(folder, 2)): {
					data.NotMovedState: {idx(folder, 2), idx(file, 2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
				idx(drive, 2): idx(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:         fullPath(1),
					id(folder):     fullPath(1, namex(folder, 2)), // note: this is a bug, but is currently expected
					idx(folder, 2): fullPath(1, namex(folder, 2)),
					idx(folder, 3): fullPath(1, namex(folder, 2)),
				},
				idx(drive, 2): {
					rootID:         fullPath(2),
					id(folder):     fullPath(2, name(folder)),
					idx(folder, 2): fullPath(2, namex(folder, 2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(id(file), idx(file, 2)),
				fullPath(2): makeExcludeMap(id(file), idx(file, 2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name:   "out of order item enumeration causes prev path collisions",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(idx(fanny, 2), name(fanny), parent(1), rootID, isFolder),
								driveItem(idx(file, 2), namex(file, 2), parent(1, name(fanny)), idx(fanny, 2), isFile),
								driveItem(id(nav), name(nav), parent(1), rootID, isFolder),
								driveItem(id(file), name(file), parent(1, name(nav)), id(nav), isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:  fullPath(1),
					id(nav): fullPath(1, name(fanny)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {
					data.NewState: {idx(fanny, 2)},
				},
				fullPath(1, name(nav)): {
					data.MovedState: {id(nav), id(file)},
				},
				fullPath(1, name(fanny)): {
					data.NewState: {idx(fanny, 2), idx(file, 2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:        fullPath(1),
					id(nav):       fullPath(1, name(nav)),
					idx(fanny, 2): fullPath(1, name(nav)), // note: this is a bug, but currently expected
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(id(file), idx(file, 2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name:   "out of order item enumeration causes prev path collisions",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(idx(fanny, 2), name(fanny), parent(1), rootID, isFolder),
								driveItem(idx(file, 2), namex(file, 2), parent(1, name(fanny)), idx(fanny, 2), isFile),
								driveItem(id(nav), name(nav), parent(1), rootID, isFolder),
								driveItem(id(file), name(file), parent(1, name(nav)), id(nav), isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:  fullPath(1),
					id(nav): fullPath(1, name(fanny)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {
					data.NewState: {idx(fanny, 2)},
				},
				fullPath(1, name(nav)): {
					data.MovedState: {id(nav), id(file)},
				},
				fullPath(1, name(fanny)): {
					data.NewState: {idx(fanny, 2), idx(file, 2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:        fullPath(1),
					id(nav):       fullPath(1, name(nav)),
					idx(fanny, 2): fullPath(1, name(nav)), // note: this is a bug, but currently expected
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(id(file), idx(file, 2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name:   "out of order item enumeration causes opposite prev path collisions",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID),
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(id(fanny), name(fanny), parent(1), rootID, isFolder),
								driveItem(id(nav), name(nav), parent(1), rootID, isFolder),
								driveItem(id(foo), name(foo), parent(1, name(fanny)), id(fanny), isFolder),
								driveItem(id(bar), name(foo), parent(1, name(nav)), id(nav), isFolder),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:    fullPath(1),
					id(nav):   fullPath(1, name(nav)),
					id(fanny): fullPath(1, name(fanny)),
					id(foo):   fullPath(1, name(nav), name(foo)),
					id(bar):   fullPath(1, name(fanny), name(foo)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				fullPath(1): {
					data.NotMovedState: {idx(file, 1)},
				},
				fullPath(1, name(nav)): {
					data.NotMovedState: {id(nav)},
				},
				fullPath(1, name(nav), name(foo)): {
					data.MovedState: {id(bar)},
				},
				fullPath(1, name(fanny)): {
					data.NotMovedState: {id(fanny)},
				},
				fullPath(1, name(fanny), name(foo)): {
					data.MovedState: {id(foo)},
				},
			},
			expectedDeltaURLs: map[string]string{
				idx(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				idx(drive, 1): {
					rootID:    fullPath(1),
					id(nav):   fullPath(1, name(nav)),
					id(fanny): fullPath(1, name(fanny)),
					id(foo):   fullPath(1, name(nav), name(foo)), // note: this is a bug, but currently expected
					id(bar):   fullPath(1, name(nav), name(foo)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				fullPath(1): makeExcludeMap(idx(file, 1)),
			}),
			doNotMergeItems: map[string]bool{},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mockDrivePager := &apiMock.Pager[models.Driveable]{
				ToReturn: []apiMock.PagerResult[models.Driveable]{
					{Values: test.drives},
				},
			}

			mbh := mock.DefaultOneDriveBH(user)
			mbh.DrivePagerV = mockDrivePager
			mbh.DriveItemEnumeration = test.enumerator

			c := NewCollections(
				mbh,
				tenant,
				idname.NewProvider(user, user),
				func(*support.ControllerOperationStatus) {},
				control.Options{ToggleFeatures: control.Toggles{}},
				count.New())

			prevDelta := "prev-delta"

			pathPrefix, err := mbh.MetadataPathPrefix(tenant)
			require.NoError(t, err, clues.ToCore(err))

			mc, err := graph.MakeMetadataCollection(
				pathPrefix,
				[]graph.MetadataCollectionEntry{
					graph.NewMetadataEntry(
						bupMD.DeltaURLsFileName,
						map[string]string{
							idx(drive, 1): prevDelta,
							idx(drive, 2): prevDelta,
						}),
					graph.NewMetadataEntry(
						bupMD.PreviousPathFileName,
						test.previousPaths),
				},
				func(*support.ControllerOperationStatus) {},
				count.New())
			assert.NoError(t, err, "creating metadata collection", clues.ToCore(err))

			prevMetadata := []data.RestoreCollection{
				dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: mc}),
			}
			errs := fault.New(true)

			delList := prefixmatcher.NewStringSetBuilder()

			cols, canUsePreviousBackup, err := c.Get(ctx, prevMetadata, delList, errs)
			test.errCheck(t, err, clues.ToCore(err))
			assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup, "can use previous backup")
			assert.Equal(t, test.expectedSkippedCount, len(errs.Skipped()))

			if err != nil {
				return
			}

			collPaths := []string{}

			for _, baseCol := range cols {
				var folderPath string
				if baseCol.State() != data.DeletedState {
					folderPath = baseCol.FullPath().String()
				} else {
					folderPath = baseCol.PreviousPath().String()
				}

				if folderPath == metadataPath.String() {
					deltas, prevs, _, err := deserializeAndValidateMetadata(
						ctx,
						[]data.RestoreCollection{
							dataMock.NewUnversionedRestoreCollection(
								t,
								data.NoFetchRestoreCollection{Collection: baseCol}),
						},
						count.New(),
						errs)
					if !assert.NoError(t, err, "deserializing metadata", clues.ToCore(err)) {
						continue
					}

					assert.Equal(t, test.expectedDeltaURLs, deltas, "delta urls")
					assert.Equal(t, test.expectedPreviousPaths, prevs, "previous paths")

					continue
				}

				collPaths = append(collPaths, folderPath)

				// TODO: We should really be getting items in the collection
				// via the Items() channel. The lack of that makes this check a bit more
				// bittle since internal details can change.  The wiring to support
				// mocked GetItems is available.  We just haven't plugged it in yet.
				col, ok := baseCol.(*Collection)
				require.True(t, ok, "getting onedrive.Collection handle")

				itemIDs := make([]string, 0, len(col.driveItems))

				for id := range col.driveItems {
					itemIDs = append(itemIDs, id)
				}

				assert.ElementsMatchf(
					t,
					test.expectedCollections[folderPath][baseCol.State()],
					itemIDs,
					"expected elements to match in collection with:\nstate '%d'\npath '%s'",
					baseCol.State(),
					folderPath)

				p := baseCol.FullPath()
				if p == nil {
					p = baseCol.PreviousPath()
				}

				assert.Equalf(
					t,
					test.doNotMergeItems[p.String()],
					baseCol.DoNotMergeItems(),
					"DoNotMergeItems in collection: %s", p)
			}

			expectCollPaths := []string{}

			for cp, c := range test.expectedCollections {
				// add one entry or each expected collection
				for range c {
					expectCollPaths = append(expectCollPaths, cp)
				}
			}

			assert.ElementsMatch(t, expectCollPaths, collPaths, "collection paths")

			test.expectedDelList.AssertEqual(t, delList, "deleted items")
		})
	}
}

// TestGet_PreviewLimits checks that the limits set for preview backups in
// control.Options.ItemLimits are respected. These tests run a reduced set of
// checks that don't examine metadata, collection states, etc. They really just
// check the expected items appear.
func (suite *CollectionsUnitSuite) TestGet_PreviewLimits() {
	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))

	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(idx(drive, 1)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	drive2 := models.NewDrive()
	drive2.SetId(ptr.To(idx(drive, 2)))
	drive2.SetName(ptr.To(namex(drive, 2)))

	table := []struct {
		name       string
		limits     control.PreviewItemLimits
		drives     []models.Driveable
		enumerator mock.EnumerateItemsDeltaByDrive
		// Collection name -> set of item IDs. We can't check item data because
		// that's not mocked out. Metadata is checked separately.
		expectedCollections map[string][]string
	}{
		{
			name: "OneDrive SinglePage ExcludeItemsOverMaxSize",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             5,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID), // will be present, not needed
								driveItemWithSize(idx(file, 1), namex(file, 1), parent(1), rootID, 7, isFile),
								driveItemWithSize(idx(file, 2), namex(file, 2), parent(1), rootID, 1, isFile),
								driveItemWithSize(idx(file, 3), namex(file, 3), parent(1), rootID, 1, isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive SinglePage SingleFolder ExcludeCombinedItemsOverMaxSize",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             3,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID), // will be present, not needed
								driveItemWithSize(idx(file, 1), namex(file, 1), parent(1), rootID, 1, isFile),
								driveItemWithSize(idx(file, 2), namex(file, 2), parent(1), rootID, 2, isFile),
								driveItemWithSize(idx(file, 3), namex(file, 3), parent(1), rootID, 1, isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2)},
			},
		},
		{
			name: "OneDrive SinglePage MultipleFolders ExcludeCombinedItemsOverMaxSize",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             3,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID), // will be present, not needed
								driveItemWithSize(idx(file, 1), namex(file, 1), parent(1), rootID, 1, isFile),
								driveItemWithSize(idx(folder, 1), namex(folder, 1), parent(1), rootID, 1, isFolder),
								driveItemWithSize(idx(file, 2), namex(file, 2), parent(1, namex(folder, 1)), idx(folder, 1), 2, isFile),
								driveItemWithSize(idx(file, 3), namex(file, 3), parent(1, namex(folder, 1)), idx(folder, 1), 1, isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 2)},
			},
		},
		{
			name: "OneDrive SinglePage SingleFolder ItemLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem(rootID), // will be present, not needed
								driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
								driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
								driveItem(idx(file, 4), namex(file, 4), parent(1), rootID, isFile),
								driveItem(idx(file, 5), namex(file, 5), parent(1), rootID, isFile),
								driveItem(idx(file, 6), namex(file, 6), parent(1), rootID, isFile),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive MultiplePages MultipleFolders ItemLimit WithRepeatedItem",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									// Repeated items shouldn't count against the limit.
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 6), namex(file, 6), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 3)},
			},
		},
		{
			name: "OneDrive MultiplePages PageLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             1,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 6), namex(file, 6), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2)},
			},
		},
		{
			name: "OneDrive MultiplePages PerContainerItemLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 1,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				// Root has an additional item. It's hard to fix that in the code
				// though.
				fullPath(1):                   {idx(file, 1), idx(file, 2)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4)},
			},
		},
		{
			name: "OneDrive MultiplePages PerContainerItemLimit ItemUpdated",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 3,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
									driveItem(idx(file, 1), namex(file, 1), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
									// Updated item that shouldn't count against the limit a second time.
									driveItem(idx(file, 2), namex(file, 2), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {},
				fullPath(1, namex(folder, 0)): {idx(folder, 0), idx(file, 1), idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive MultiplePages PerContainerItemLimit MoveItemBetweenFolders",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 2,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									// Put folder 0 at limit.
									driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
									driveItem(idx(file, 3), namex(file, 3), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 0), namex(folder, 0), parent(1), rootID, isFolder),
									// Try to move item from root to folder 0 which is already at the limit.
									driveItem(idx(file, 1), namex(file, 1), parent(1, namex(folder, 0)), idx(folder, 0), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2)},
				fullPath(1, namex(folder, 0)): {idx(folder, 0), idx(file, 3), idx(file, 4)},
			},
		},
		{
			name: "OneDrive MultiplePages ContainerLimit LastContainerSplitAcrossPages",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        2,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
		{
			name: "OneDrive MultiplePages ContainerLimit NextContainerOnSamePage",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        2,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									// This container shouldn't be returned.
									driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
									driveItem(idx(file, 7), namex(file, 7), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
									driveItem(idx(file, 8), namex(file, 8), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
									driveItem(idx(file, 9), namex(file, 9), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
		{
			name: "OneDrive MultiplePages ContainerLimit NextContainerOnNextPage",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        2,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									// This container shouldn't be returned.
									driveItem(idx(folder, 2), namex(folder, 2), parent(1), rootID, isFolder),
									driveItem(idx(file, 7), namex(file, 7), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
									driveItem(idx(file, 8), namex(file, 8), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
									driveItem(idx(file, 9), namex(file, 9), parent(1, namex(folder, 2)), idx(folder, 2), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
		{
			name: "TwoDrives SeparateLimitAccounting",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
				MaxBytes:             999999,
				MaxPages:             999,
			},
			drives: []models.Driveable{drive1, drive2},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
									driveItem(idx(file, 4), namex(file, 4), parent(1), rootID, isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(1), rootID, isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
					idx(drive, 2): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(2), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(2), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(2), rootID, isFile),
									driveItem(idx(file, 4), namex(file, 4), parent(2), rootID, isFile),
									driveItem(idx(file, 5), namex(file, 5), parent(2), rootID, isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1): {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(2): {idx(file, 1), idx(file, 2), idx(file, 3)},
			},
		},
		{
			name: "OneDrive PreviewDisabled MinimumLimitsIgnored",
			limits: control.PreviewItemLimits{
				MaxItems:             1,
				MaxItemsPerContainer: 1,
				MaxContainers:        1,
				MaxBytes:             1,
				MaxPages:             1,
			},
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(file, 1), namex(file, 1), parent(1), rootID, isFile),
									driveItem(idx(file, 2), namex(file, 2), parent(1), rootID, isFile),
									driveItem(idx(file, 3), namex(file, 3), parent(1), rootID, isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 4), namex(file, 4), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem(rootID), // will be present, not needed
									driveItem(idx(folder, 1), namex(folder, 1), parent(1), rootID, isFolder),
									driveItem(idx(file, 5), namex(file, 5), parent(1, namex(folder, 1)), idx(folder, 1), isFile),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta)},
					},
				},
			},
			expectedCollections: map[string][]string{
				fullPath(1):                   {idx(file, 1), idx(file, 2), idx(file, 3)},
				fullPath(1, namex(folder, 1)): {idx(folder, 1), idx(file, 4), idx(file, 5)},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mockDrivePager := &apiMock.Pager[models.Driveable]{
				ToReturn: []apiMock.PagerResult[models.Driveable]{
					{Values: test.drives},
				},
			}

			mbh := mock.DefaultOneDriveBH(user)
			mbh.DrivePagerV = mockDrivePager
			mbh.DriveItemEnumeration = test.enumerator

			opts := control.DefaultOptions()
			opts.PreviewLimits = test.limits

			c := NewCollections(
				mbh,
				tenant,
				idname.NewProvider(user, user),
				func(*support.ControllerOperationStatus) {},
				opts,
				count.New())

			errs := fault.New(true)

			delList := prefixmatcher.NewStringSetBuilder()

			cols, canUsePreviousBackup, err := c.Get(ctx, nil, delList, errs)
			require.NoError(t, err, clues.ToCore(err))

			assert.True(t, canUsePreviousBackup, "can use previous backup")
			assert.Empty(t, errs.Skipped())

			collPaths := []string{}

			for _, baseCol := range cols {
				// There shouldn't be any deleted collections.
				if !assert.NotEqual(
					t,
					data.DeletedState,
					baseCol.State(),
					"collection marked deleted") {
					continue
				}

				folderPath := baseCol.FullPath().String()

				if folderPath == metadataPath.String() {
					continue
				}

				collPaths = append(collPaths, folderPath)

				// TODO: We should really be getting items in the collection
				// via the Items() channel. The lack of that makes this check a bit more
				// bittle since internal details can change.  The wiring to support
				// mocked GetItems is available.  We just haven't plugged it in yet.
				col, ok := baseCol.(*Collection)
				require.True(t, ok, "getting onedrive.Collection handle")

				itemIDs := make([]string, 0, len(col.driveItems))

				for id := range col.driveItems {
					itemIDs = append(itemIDs, id)
				}

				assert.ElementsMatchf(
					t,
					test.expectedCollections[folderPath],
					itemIDs,
					"expected elements to match in collection with path %q",
					folderPath)
			}

			assert.ElementsMatch(
				t,
				maps.Keys(test.expectedCollections),
				collPaths,
				"collection paths")
		})
	}
}

func (suite *CollectionsUnitSuite) TestAddURLCacheToDriveCollections() {
	drive1 := models.NewDrive()
	drive1.SetId(ptr.To(idx(drive, 1)))
	drive1.SetName(ptr.To(namex(drive, 1)))

	drive2 := models.NewDrive()
	drive2.SetId(ptr.To(idx(drive, 2)))
	drive2.SetName(ptr.To(namex(drive, 2)))

	table := []struct {
		name       string
		drives     []models.Driveable
		enumerator mock.EnumerateItemsDeltaByDrive
		errCheck   assert.ErrorAssertionFunc
	}{
		{
			name: "Two drives with unique url cache instances",
			drives: []models.Driveable{
				drive1,
				drive2,
			},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					idx(drive, 1): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(id(folder), name(folder), parent(1), rootID, isFolder),
							driveItem(id(file), name(file), parent(1, name(folder)), id(folder), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: id(delta), Reset: true},
					},
					idx(drive, 2): {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem(rootID),
							driveItem(idx(folder, 2), name(folder), parent(2), rootID, isFolder),
							driveItem(idx(file, 2), name(file), parent(2, name(folder)), idx(folder, 2), isFile),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: idx(delta, 2), Reset: true},
					},
				},
			},
			errCheck: assert.NoError,
		},
		// TODO(pandeyabs): Add a test case to check that the cache is not attached
		// if a drive has more than urlCacheDriveItemThreshold discovered items.
		// This will require creating 300k+ mock items for the test which might take
		// up a lot of memory during the test. Include it after testing out mem usage.
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mockDrivePager := &apiMock.Pager[models.Driveable]{
				ToReturn: []apiMock.PagerResult[models.Driveable]{
					{Values: test.drives},
				},
			}

			mbh := mock.DefaultOneDriveBH(user)
			mbh.DrivePagerV = mockDrivePager
			mbh.DriveItemEnumeration = test.enumerator

			c := NewCollections(
				mbh,
				tenant,
				idname.NewProvider(user, user),
				func(*support.ControllerOperationStatus) {},
				control.Options{ToggleFeatures: control.Toggles{}},
				count.New())

			errs := fault.New(true)
			delList := prefixmatcher.NewStringSetBuilder()

			cols, _, err := c.Get(ctx, nil, delList, errs)
			test.errCheck(t, err)

			// Group collections by drive ID
			colsByDrive := map[string][]*Collection{}

			for _, col := range cols {
				c, ok := col.(*Collection)
				if !ok {
					// skip metadata collection
					continue
				}

				colsByDrive[c.driveID] = append(colsByDrive[c.driveID], c)
			}

			caches := map[*urlCache]struct{}{}

			// Check that the URL cache is attached to each collection.
			// Also check that each drive gets its own cache instance.
			for drive, driveCols := range colsByDrive {
				var uc *urlCache
				for _, col := range driveCols {
					require.NotNil(t, col.urlCache, "cache is nil")

					if uc == nil {
						uc = col.urlCache.(*urlCache)
					} else {
						require.Equal(
							t,
							uc,
							col.urlCache,
							"drive collections have different url cache instances")
					}

					require.Equal(t, drive, uc.driveID, "drive ID mismatch")
				}

				caches[uc] = struct{}{}
			}

			// Check that we have the expected number of caches. One per drive.
			require.Equal(
				t,
				len(test.drives),
				len(caches),
				"expected one cache per drive")
		})
	}
}
