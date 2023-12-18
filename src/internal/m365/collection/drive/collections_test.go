package drive

import (
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
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

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
	d := drive()

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
				rootFolder(),
				driveItem(id(item), name(item), d.dir(), rootID, -1),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, d.strPath(t)),
			},
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: d.strPath(t),
			},
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single File",
			items: []models.DriveItemable{
				rootFolder(),
				driveFile(d.dir(), rootID),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, d.strPath(t)),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// Root folder is skipped since it's always present.
			expectedPrevPaths: map[string]string{
				rootID: d.strPath(t),
			},
			expectedExcludes:         makeExcludeMap(fileID()),
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asNew(t, d.strPath(t, folderName())),
			},
			expectedPrevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder created twice", // deleted a created with same name in between a backup
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
				driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:      asNotMoved(t, d.strPath(t)),
				folderID(2): asNew(t, d.strPath(t, folderName())),
			},
			expectedPrevPaths: map[string]string{
				rootID:      d.strPath(t),
				folderID(2): d.strPath(t, folderName()),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Package",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:  asNotMoved(t, d.strPath(t)),
				id(pkg): asNew(t, d.strPath(t, name(pkg))),
			},
			expectedPrevPaths: map[string]string{
				rootID:  d.strPath(t),
				id(pkg): d.strPath(t, name(pkg)),
			},
			expectedItemCount:      1,
			expectedContainerCount: 2,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				d.strPath(t, name(pkg)): {},
			},
			expectedCountPackages: 1,
		},
		{
			name: "Single Package with subfolder",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
				driveItem(folderID(), folderName(), d.dir(name(pkg)), id(pkg), isFolder),
				driveItem(id(subfolder), name(subfolder), d.dir(name(pkg)), id(pkg), isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, d.strPath(t)),
				id(pkg):       asNew(t, d.strPath(t, name(pkg))),
				folderID():    asNew(t, d.strPath(t, name(pkg), folderName())),
				id(subfolder): asNew(t, d.strPath(t, name(pkg), name(subfolder))),
			},
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				id(pkg):       d.strPath(t, name(pkg)),
				folderID():    d.strPath(t, name(pkg), folderName()),
				id(subfolder): d.strPath(t, name(pkg), name(subfolder)),
			},
			expectedItemCount:      3,
			expectedContainerCount: 4,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				d.strPath(t, name(pkg)): {},
			},
			expectedCountPackages: 3,
		},
		{
			name: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				rootFolder(),
				driveFile(d.dir(), rootID, "inRoot"),
				driveFolder(d.dir(), rootID),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
				driveFile(d.dir(folderName()), folderID(), "inFolder"),
				driveFile(d.dir(name(pkg)), id(pkg), "inPackage"),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asNew(t, d.strPath(t, folderName())),
				id(pkg):    asNew(t, d.strPath(t, name(pkg))),
			},
			expectedItemCount:      5,
			expectedFileCount:      3,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
				id(pkg):    d.strPath(t, name(pkg)),
			},
			expectedTopLevelPackages: map[string]struct{}{
				d.strPath(t, name(pkg)): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap(fileID("inRoot"), fileID("inFolder"), fileID("inPackage")),
		},
		{
			name: "contains folder selector",
			items: []models.DriveItemable{
				rootFolder(),
				driveFile(d.dir(), rootID, "inRoot"),
				driveFolder(d.dir(), rootID),
				driveItem(id(subfolder), name(subfolder), d.dir(folderName()), folderID(), isFolder),
				driveItem(folderID(2), folderName(), d.dir(folderName(), name(subfolder)), id(subfolder), isFolder),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
				driveItem(fileID("inFolder"), fileID("inFolder"), d.dir(folderName()), folderID(), isFile),
				driveItem(fileID("inFolder2"), fileName("inFolder2"), d.dir(folderName(), name(subfolder), folderName()), folderID(2), isFile),
				driveItem(fileID("inFolderPackage"), fileName("inPackage"), d.dir(name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{folderName()})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				folderID():    asNew(t, d.strPath(t, folderName())),
				id(subfolder): asNew(t, d.strPath(t, folderName(), name(subfolder))),
				folderID(2):   asNew(t, d.strPath(t, folderName(), name(subfolder), folderName())),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			// just "folder" isn't added here because the include check is done on the
			// parent path since we only check later if something is a folder or not.
			expectedPrevPaths: map[string]string{
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
				folderID(2):   d.strPath(t, folderName(), name(subfolder), folderName()),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID("inFolder"), fileID("inFolder2")),
		},
		{
			name: "prefix subfolder selector",
			items: []models.DriveItemable{
				rootFolder(),
				driveFile(d.dir(), rootID, "inRoot"),
				driveFolder(d.dir(), rootID),
				driveItem(id(subfolder), name(subfolder), d.dir(folderName()), folderID(), isFolder),
				driveItem(folderID(2), folderName(), d.dir(folderName(), name(subfolder)), id(subfolder), isFolder),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
				driveItem(fileID("inFolder"), fileID("inFolder"), d.dir(folderName()), folderID(), isFile),
				driveItem(fileID("inFolder2"), fileName("inFolder2"), d.dir(folderName(), name(subfolder), folderName()), folderID(2), isFile),
				driveItem(fileID("inFolderPackage"), fileName("inPackage"), d.dir(name(pkg)), id(pkg), isFile),
			},
			previousPaths: map[string]string{},
			scope: (&selectors.OneDriveBackup{}).Folders(
				[]string{toPath(folderName(), name(subfolder))},
				selectors.PrefixMatch())[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(subfolder): asNew(t, d.strPath(t, folderName(), name(subfolder))),
				folderID(2):   asNew(t, d.strPath(t, folderName(), name(subfolder), folderName())),
			},
			expectedItemCount:      3,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
				folderID(2):   d.strPath(t, folderName(), name(subfolder), folderName()),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID("inFolder2")),
		},
		{
			name: "match subfolder selector",
			items: []models.DriveItemable{
				rootFolder(),
				driveFile(d.dir(), rootID),
				driveFolder(d.dir(), rootID),
				driveItem(id(subfolder), name(subfolder), d.dir(folderName()), folderID(), isFolder),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
				driveItem(fileID(1), fileName(1), d.dir(folderName()), folderID(), isFile),
				driveItem(fileID("inSubfolder"), fileName("inSubfolder"), d.dir(folderName(), name(subfolder)), id(subfolder), isFile),
				driveItem(fileID(9), fileName(9), d.dir(name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{toPath(folderName(), name(subfolder))})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(subfolder): asNew(t, d.strPath(t, folderName(), name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// No child folders for subfolder so nothing here.
			expectedPrevPaths: map[string]string{
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID("inSubfolder")),
		},
		{
			name: "not moved folder tree",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asNotMoved(t, d.strPath(t, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asMoved(t, d.strPath(t, folderName("a")), d.strPath(t, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(folderID(1), folderName(), d.dir(), rootID, isFolder),
				driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID(1):   d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:      asNotMoved(t, d.strPath(t)),
				folderID(2): asNew(t, d.strPath(t, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID(2):   d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup",
			items: []models.DriveItemable{
				rootFolder(),
				delItem(folderID(), rootID, isFolder),
				driveItem(folderID(), name(drivePfx), d.dir(), rootID, isFolder),
				delItem(folderID(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asDeleted(t, d.strPath(t, "")),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup including delete",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
				delItem(folderID(), rootID, isFolder),
				driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:      asNotMoved(t, d.strPath(t)),
				folderID(2): asNew(t, d.strPath(t, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID(2):   d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup with addition",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(folderID(1), folderName(), d.dir(), rootID, isFolder),
				delItem(folderID(1), rootID, isFolder),
				driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
				delItem(folderID(2), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID(1):   d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, d.strPath(t)),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree with file no previous",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
				driveItem(fileID(), fileName(), d.dir(folderName()), folderID(), isFile),
				driveItem(folderID(), folderName(2), d.dir(), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asNew(t, d.strPath(t, folderName(2))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName(2)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID()),
		},
		{
			name: "moved folder tree with file no previous 1",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
				driveItem(fileID(), fileName(), d.dir(folderName()), folderID(), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asNew(t, d.strPath(t, folderName())),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID()),
		},
		{
			name: "moved folder tree and subfolder 1",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
				driveItem(id(subfolder), name(subfolder), d.dir(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, d.strPath(t)),
				folderID():    asMoved(t, d.strPath(t, folderName("a")), d.strPath(t, folderName())),
				id(subfolder): asMoved(t, d.strPath(t, folderName("a"), name(subfolder)), d.strPath(t, name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree and subfolder 2",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(id(subfolder), name(subfolder), d.dir(), rootID, isFolder),
				driveFolder(d.dir(), rootID),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, d.strPath(t)),
				folderID():    asMoved(t, d.strPath(t, folderName("a")), d.strPath(t, folderName())),
				id(subfolder): asMoved(t, d.strPath(t, folderName("a"), name(subfolder)), d.strPath(t, name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "move subfolder when moving parent",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(folderID(2), folderName(2), d.dir(), rootID, isFolder),
				driveItem(id(item), name(item), d.dir(folderName(2)), folderID(2), isFile),
				// Need to see the parent folder first (expected since that's what Graph
				// consistently returns).
				driveItem(folderID(), folderName("a"), d.dir(), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), d.dir(folderName("a")), folderID(), isFolder),
				driveItem(id(item, 2), name(item, 2), d.dir(folderName("a"), name(subfolder)), id(subfolder), isFile),
				driveFolder(d.dir(), rootID),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, d.strPath(t)),
				folderID(2):   asNew(t, d.strPath(t, folderName(2))),
				folderID():    asMoved(t, d.strPath(t, folderName("a")), d.strPath(t, folderName())),
				id(subfolder): asMoved(t, d.strPath(t, folderName("a"), name(subfolder)), d.strPath(t, folderName(), name(subfolder))),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 4,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName()),
				folderID(2):   d.strPath(t, folderName(2)),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(item), id(item, 2)),
		},
		{
			name: "moved folder tree multiple times",
			items: []models.DriveItemable{
				rootFolder(),
				driveFolder(d.dir(), rootID),
				driveItem(fileID(), fileName(), d.dir(folderName()), folderID(), isFile),
				driveItem(folderID(), folderName(2), d.dir(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    d.strPath(t, folderName("a")),
				id(subfolder): d.strPath(t, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asMoved(t, d.strPath(t, folderName("a")), d.strPath(t, folderName(2))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName(2)),
				id(subfolder): d.strPath(t, folderName(2), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID()),
		},
		{
			name: "deleted folder and package",
			items: []models.DriveItemable{
				rootFolder(), // root is always present, but not necessary here
				delItem(folderID(), rootID, isFolder),
				delItem(id(pkg), rootID, isPackage),
			},
			previousPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
				id(pkg):    d.strPath(t, name(pkg)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asDeleted(t, d.strPath(t, folderName())),
				id(pkg):    asDeleted(t, d.strPath(t, name(pkg))),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: d.strPath(t),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder without previous",
			items: []models.DriveItemable{
				rootFolder(),
				delItem(folderID(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				rootID: d.strPath(t),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, d.strPath(t)),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: d.strPath(t),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder tree move subfolder",
			items: []models.DriveItemable{
				rootFolder(),
				delItem(folderID(), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), d.dir(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				rootID:        d.strPath(t),
				folderID():    d.strPath(t, folderName()),
				id(subfolder): d.strPath(t, folderName(), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, d.strPath(t)),
				folderID():    asDeleted(t, d.strPath(t, folderName())),
				id(subfolder): asMoved(t, d.strPath(t, folderName(), name(subfolder)), d.strPath(t, name(subfolder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        d.strPath(t),
				id(subfolder): d.strPath(t, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete file",
			items: []models.DriveItemable{
				rootFolder(),
				delItem(id(item), rootID, isFile),
			},
			previousPaths: map[string]string{
				rootID: d.strPath(t),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, d.strPath(t)),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: d.strPath(t),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(item)),
		},
		{
			name: "item before parent errors",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(fileID(), fileName(), d.dir(folderName()), folderID(), isFile),
				driveFolder(d.dir(), rootID),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, d.strPath(t)),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: d.strPath(t),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "1 root file, 1 folder, 1 package, 1 good file, 1 malware",
			items: []models.DriveItemable{
				rootFolder(),
				driveItem(fileID(), fileID(), d.dir(), rootID, isFile),
				driveFolder(d.dir(), rootID),
				driveItem(id(pkg), name(pkg), d.dir(), rootID, isPackage),
				driveItem(fileID("good"), fileName("good"), d.dir(folderName()), folderID(), isFile),
				malwareItem(id(malware), name(malware), d.dir(folderName()), folderID(), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, d.strPath(t)),
				folderID(): asNew(t, d.strPath(t, folderName())),
				id(pkg):    asNew(t, d.strPath(t, name(pkg))),
			},
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			expectedSkippedCount:   1,
			expectedPrevPaths: map[string]string{
				rootID:     d.strPath(t),
				folderID(): d.strPath(t, folderName()),
				id(pkg):    d.strPath(t, name(pkg)),
			},
			expectedTopLevelPackages: map[string]struct{}{
				d.strPath(t, name(pkg)): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap(fileID(), fileID("good")),
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				drive    = drive()
				mbh      = defaultOneDriveBH(user)
				excludes = map[string]struct{}{}
				errs     = fault.New(true)
			)

			mbh.DriveItemEnumeration = driveEnumerator(
				drive.newEnumer().with(
					delta(nil, "notempty").with(
						aPage(test.items...))))

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

			c.CollectionMap[drive.id] = map[string]*Collection{}

			_, newPrevPaths, err := c.PopulateDriveCollections(
				ctx,
				drive.id,
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
				maps.Keys(c.CollectionMap[drive.id]),
				"expected collection IDs")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			assert.Equal(t, test.expectedSkippedCount, len(errs.Skipped()), "skipped item count")

			for id, sp := range test.expectedCollectionIDs {
				if !assert.Containsf(t, c.CollectionMap[drive.id], id, "missing collection with id %s", id) {
					// Skip collections we don't find so we don't get an NPE.
					continue
				}

				assert.Equalf(t, sp.state, c.CollectionMap[drive.id][id].State(), "state for collection %s", id)
				assert.Equalf(t, sp.currPath, c.CollectionMap[drive.id][id].FullPath(), "current path for collection %s", id)
				assert.Equalf(t, sp.prevPath, c.CollectionMap[drive.id][id].PreviousPath(), "prev path for collection %s", id)
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
	t := suite.T()
	d := drive()
	d2 := drive(2)

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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				d.id: deltaURL(),
			},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
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
							map[string]string{d.id: deltaURL()}),
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
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {},
							}),
					}
				},
			},
			expectedDeltas:       map[string]string{},
			expectedPaths:        map[string]map[string]string{d.id: {}},
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
								d.id: "",
							}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{d.id: ""},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{d2.id: deltaURL(2)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d2.id: {
									folderID(2): d2.strPath(t),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				d.id:  deltaURL(),
				d2.id: deltaURL(2),
			},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
				},
				d2.id: {
					folderID(2): d2.strPath(t),
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
							map[string]string{d.id: deltaURL()}),
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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
						graph.NewMetadataEntry(
							"foo",
							map[string]string{d.id: deltaURL()}),
					}
				},
			},
			expectedDeltas: map[string]string{
				d.id: deltaURL(),
			},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(2): d2.strPath(t),
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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{d.id: deltaURL(2)}),
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
							map[string]string{d.id: deltaURL()}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
									folderID(2): d.strPath(t),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				d.id: deltaURL(),
			},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
					folderID(2): d.strPath(t),
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
								d.id: deltaURL(),
							}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d.id: {
									folderID(1): d.strPath(t),
									folderID(2): d.strPath(t),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{d2.id: deltaURL(2)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								d2.id: {
									folderID(1): d.strPath(t),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				d.id:  deltaURL(),
				d2.id: deltaURL(2),
			},
			expectedPaths: map[string]map[string]string{
				d.id: {
					folderID(1): d.strPath(t),
					folderID(2): d.strPath(t),
				},
				d2.id: {
					folderID(1): d.strPath(t),
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

	_, _, canUsePreviousBackup, err := deserializeAndValidateMetadata(
		ctx,
		[]data.RestoreCollection{fc},
		count.New(),
		fault.New(true))
	require.NoError(t, err)
	require.False(t, canUsePreviousBackup)
}

func (suite *CollectionsUnitSuite) TestGet_treeCannotBeUsedWhileIncomplete() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	mbh := defaultOneDriveBH(user)
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	mbh.DriveItemEnumeration = driveEnumerator(
		drive().newEnumer().with(
			delta(nil).with(
				aPage(
					delItem(fileID(), rootID, isFile)))))

	c := collWithMBH(mbh)
	c.ctrl = opts

	_, _, err := c.Get(ctx, nil, nil, fault.New(true))
	require.ErrorIs(t, err, errGetTreeNotImplemented, clues.ToCore(err))
}

func (suite *CollectionsUnitSuite) TestGet() {
	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))

	t := suite.T()
	d := drive(1)
	d2 := drive(2)

	table := []struct {
		name                 string
		enumerator           enumerateDriveItemsDelta
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
			name: "OneDrive_OneItemPage_DelFileOnly_NoFolders_NoErrors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {rootID: d.strPath(t)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NotMovedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {rootID: d.strPath(t)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				d.strPath(t): makeExcludeMap(fileID()),
			}),
		},
		{
			name: "OneDrive_OneItemPage_NoFolderDeltas_NoErrors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							driveFile(d.dir(), rootID))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {rootID: d.strPath(t)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NotMovedState: {fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {rootID: d.strPath(t)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				d.strPath(t): makeExcludeMap(fileID()),
			}),
		},
		{
			name: "OneDrive_OneItemPage_NoErrors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_NoErrors_FileRenamedMultiple",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							driveItem(fileID(), fileName(2), d.dir(folderName()), folderID(), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_NoErrors_FileMovedMultiple",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							driveItem(fileID(), fileName(2), d.dir(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NotMovedState: {fileID()}},
				d.strPath(t, folderName()): {data.NewState: {folderID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				d.strPath(t): makeExcludeMap(fileID()),
			}),
		},
		{
			name: "OneDrive_TwoItemPages_NoErrors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID(), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_TwoItemPages_WithReset",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							driveItem(fileID(3), fileName(3), d.dir(folderName()), folderID(), isFile)),
						aReset(),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID(), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_TwoItemPages_WithResetCombinedWithItems",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPageWReset(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID(), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "TwoDrives_OneItemPageEach_NoErrors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())))),
				d2.newEnumer().with(
					deltaWReset(nil, 2).with(aPage(
						driveItem(folderID(2), folderName(), d2.dir(), rootID, isFolder),
						driveItem(fileID(2), fileName(), d2.dir(folderName()), folderID(2), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
				d2.id:           {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):                {data.NewState: {}},
				d.strPath(t, folderName()):  {data.NewState: {folderID(), fileID()}},
				d2.strPath(t):               {data.NewState: {}},
				d2.strPath(t, folderName()): {data.NewState: {folderID(2), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
				d2.id:           deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
				d2.id: {
					rootID:      d2.strPath(t),
					folderID(2): d2.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):                true,
				d.strPath(t, folderName()):  true,
				d2.strPath(t):               true,
				d2.strPath(t, folderName()): true,
			},
		},
		{
			name: "TwoDrives_DuplicateIDs_OneItemPageEach_NoErrors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())))),
				d2.newEnumer().with(
					deltaWReset(nil, 2).with(
						aPage(
							driveFolder(d2.dir(), rootID),
							driveItem(fileID(2), fileName(), d2.dir(folderName()), folderID(), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
				d2.id:           {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):                {data.NewState: {}},
				d.strPath(t, folderName()):  {data.NewState: {folderID(), fileID()}},
				d2.strPath(t):               {data.NewState: {}},
				d2.strPath(t, folderName()): {data.NewState: {folderID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
				d2.id:           deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
				d2.id: {
					rootID:     d2.strPath(t),
					folderID(): d2.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):                true,
				d.strPath(t, folderName()):  true,
				d2.strPath(t):               true,
				d2.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_Errors",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(assert.AnError))),
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections:   nil,
			expectedDeltaURLs:     nil,
			expectedPreviousPaths: nil,
			expectedDelList:       nil,
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_DeleteNonExistentFolder",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage(
							driveFolder(d.dir(), rootID, 2),
							driveFile(d.dir(folderName(2)), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):                {data.NewState: {}},
				d.strPath(t, folderName()):  {data.DeletedState: {}},
				d.strPath(t, folderName(2)): {data.NewState: {folderID(2), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(2): d.strPath(t, folderName(2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):                true,
				d.strPath(t, folderName()):  true,
				d.strPath(t, folderName(2)): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDeltaCombinedWithItems_DeleteNonExistentFolder",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage(
							driveFolder(d.dir(), rootID, 2),
							driveFile(d.dir(folderName(2)), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):                {data.NewState: {}},
				d.strPath(t, folderName()):  {data.DeletedState: {}},
				d.strPath(t, folderName(2)): {data.NewState: {folderID(2), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(2): d.strPath(t, folderName(2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):                true,
				d.strPath(t, folderName()):  true,
				d.strPath(t, folderName(2)): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
							driveFile(d.dir(folderName()), folderID(2))),
						aReset(),
						aPage(
							driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
							driveFile(d.dir(folderName()), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
				d.strPath(t, folderName()): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {folderID(2), fileID()},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(2): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtExistingLocation",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aReset(),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
				d.strPath(t, folderName()): {
					data.NewState: {folderID(), fileID()},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_ImmediateInvalidPrevDelta_MoveFolderToPreviouslyExistingPath",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage(
							driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
							driveItem(fileID(2), fileName(), d.dir(folderName()), folderID(2), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
				d.strPath(t, folderName()): {
					data.DeletedState: {},
					data.NewState:     {folderID(2), fileID(2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(2): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aReset(),
						aPage(
							driveItem(folderID(2), folderName(), d.dir(), rootID, isFolder),
							driveFile(d.dir(folderName()), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
				d.strPath(t, folderName()): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {folderID(2), fileID()},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(2): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "OneDrive Two Item Pages with Malware",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							malwareItem(id(malware), name(malware), d.dir(folderName()), folderID(), isFile)),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID(), 2),
							malwareItem(id(malware, 2), name(malware, 2), d.dir(folderName()), folderID(), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
			expectedSkippedCount: 2,
		},
		{
			name: "One Drive Deleted Folder In New Results With Invalid Delta",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil, 2).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							driveFolder(d.dir(), rootID, 2),
							driveFile(d.dir(folderName(2)), folderID(2), 2)),
						aReset(),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							delItem(folderID(2), rootID, isFolder),
							delItem(fileName(2), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID():  d.strPath(t, folderName()),
					folderID(2): d.strPath(t, folderName(2)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):                {data.NewState: {}},
				d.strPath(t, folderName()):  {data.NewState: {folderID(), fileID()}},
				d.strPath(t, folderName(2)): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):                true,
				d.strPath(t, folderName()):  true,
				d.strPath(t, folderName(2)): true,
			},
		},
		{
			name: "One Drive Folder Delete After Invalid Delta",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPageWReset(
							delItem(folderID(), rootID, isFolder))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "One Drive Item Delete After Invalid Delta",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPageWReset(
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t): true,
			},
		},
		{
			name: "One Drive Folder Made And Deleted",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil, 2).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t): true,
			},
		},
		{
			name: "One Drive Folder Created -> Deleted -> Created",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil, 2).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile)),
						aPage(
							driveItem(folderID(1), folderName(), d.dir(), rootID, isFolder),
							driveItem(fileID(1), fileName(), d.dir(folderName()), folderID(1), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID(1), fileID(1)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(1): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "One Drive Folder Deleted -> Created -> Deleted",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil, 2).with(
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile)),
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NotMovedState: {}},
				d.strPath(t, folderName()): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name: "One Drive Folder Created -> Deleted -> Created with prev",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil, 2).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile)),
						aPage(
							driveItem(folderID(1), folderName(), d.dir(), rootID, isFolder),
							driveItem(fileID(1), fileName(), d.dir(folderName()), folderID(1), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.DeletedState: {}, data.NewState: {folderID(1), fileID(1)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID(1): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               false,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "One Drive Item Made And Deleted",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())),
						aPage(delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):               {data.NewState: {}},
				d.strPath(t, folderName()): {data.NewState: {folderID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:     d.strPath(t),
					folderID(): d.strPath(t, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t):               true,
				d.strPath(t, folderName()): true,
			},
		},
		{
			name: "One Drive Random Folder Delete",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					deltaWReset(nil).with(
						aPage(
							delItem(folderID(), rootID, isFolder))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t): true,
			},
		},
		{
			name: "One Drive Random Item Delete",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID: d.strPath(t),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d.strPath(t): true,
			},
		},
		{
			name: "TwoPriorDrives_OneTombstoned",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(aPage()))), // root only
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {rootID: d.strPath(t)},
				d2.id:           {rootID: d2.strPath(t)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t):  {data.NotMovedState: {}},
				d2.strPath(t): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{id(drivePfx, 1): deltaURL()},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {rootID: d.strPath(t)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				d2.strPath(t): true,
			},
		},
		{
			name: "duplicate previous paths in metadata",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID()),
							driveFolder(d.dir(), rootID, 2),
							driveFile(d.dir(folderName(2)), folderID(2), 2)))),
				d2.newEnumer().with(
					delta(nil, 2).with(
						aPage(
							driveFolder(d2.dir(), rootID),
							driveFile(d2.dir(folderName()), folderID()),
							driveFolder(d2.dir(), rootID, 2),
							driveFile(d2.dir(folderName(2)), folderID(2), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID():  d.strPath(t, folderName()),
					folderID(2): d.strPath(t, folderName()),
					folderID(3): d.strPath(t, folderName()),
				},
				d2.id: {
					rootID:      d2.strPath(t),
					folderID():  d2.strPath(t, folderName()),
					folderID(2): d2.strPath(t, folderName(2)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {
					data.NewState: {folderID(), folderID(2)},
				},
				d.strPath(t, folderName()): {
					data.NotMovedState: {folderID(), fileID()},
				},
				d.strPath(t, folderName(2)): {
					data.MovedState: {folderID(2), fileID(2)},
				},
				d2.strPath(t): {
					data.NewState: {folderID(), folderID(2)},
				},
				d2.strPath(t, folderName()): {
					data.NotMovedState: {folderID(), fileID()},
				},
				d2.strPath(t, folderName(2)): {
					data.NotMovedState: {folderID(2), fileID(2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
				d2.id:           deltaURL(2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:      d.strPath(t),
					folderID():  d.strPath(t, folderName(2)), // note: this is a bug, but is currently expected
					folderID(2): d.strPath(t, folderName(2)),
					folderID(3): d.strPath(t, folderName(2)),
				},
				d2.id: {
					rootID:      d2.strPath(t),
					folderID():  d2.strPath(t, folderName()),
					folderID(2): d2.strPath(t, folderName(2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				d.strPath(t):  makeExcludeMap(fileID(), fileID(2)),
				d2.strPath(t): makeExcludeMap(fileID(), fileID(2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name: "out of order item enumeration causes prev path collisions",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							driveItem(folderID(fanny, 2), folderName(fanny), d.dir(), rootID, isFolder),
							driveFile(d.dir(folderName(fanny)), folderID(fanny, 2), 2),
							driveFolder(d.dir(), rootID, nav),
							driveFile(d.dir(folderName(nav)), folderID(nav)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:        d.strPath(t),
					folderID(nav): d.strPath(t, folderName(fanny)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {
					data.NewState: {folderID(fanny, 2)},
				},
				d.strPath(t, folderName(nav)): {
					data.MovedState: {folderID(nav), fileID()},
				},
				d.strPath(t, folderName(fanny)): {
					data.NewState: {folderID(fanny, 2), fileID(2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:             d.strPath(t),
					folderID(nav):      d.strPath(t, folderName(nav)),
					folderID(fanny, 2): d.strPath(t, folderName(nav)), // note: this is a bug, but currently expected
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				d.strPath(t): makeExcludeMap(fileID(), fileID(2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name: "out of order item enumeration causes opposite prev path collisions",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							driveFile(d.dir(), rootID, 1),
							driveFolder(d.dir(), rootID, fanny),
							driveFolder(d.dir(), rootID, nav),
							driveFolder(d.dir(folderName(fanny)), folderID(fanny), foo),
							driveItem(folderID(bar), folderName(foo), d.dir(folderName(nav)), folderID(nav), isFolder))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:          d.strPath(t),
					folderID(nav):   d.strPath(t, folderName(nav)),
					folderID(fanny): d.strPath(t, folderName(fanny)),
					folderID(foo):   d.strPath(t, folderName(nav), folderName(foo)),
					folderID(bar):   d.strPath(t, folderName(fanny), folderName(foo)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				d.strPath(t): {
					data.NotMovedState: {fileID(1)},
				},
				d.strPath(t, folderName(nav)): {
					data.NotMovedState: {folderID(nav)},
				},
				d.strPath(t, folderName(nav), folderName(foo)): {
					data.MovedState: {folderID(bar)},
				},
				d.strPath(t, folderName(fanny)): {
					data.NotMovedState: {folderID(fanny)},
				},
				d.strPath(t, folderName(fanny), folderName(foo)): {
					data.MovedState: {folderID(foo)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drivePfx, 1): deltaURL(),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drivePfx, 1): {
					rootID:          d.strPath(t),
					folderID(nav):   d.strPath(t, folderName(nav)),
					folderID(fanny): d.strPath(t, folderName(fanny)),
					folderID(foo):   d.strPath(t, folderName(nav), folderName(foo)), // note: this is a bug, but currently expected
					folderID(bar):   d.strPath(t, folderName(nav), folderName(foo)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				d.strPath(t): makeExcludeMap(fileID(1)),
			}),
			doNotMergeItems: map[string]bool{},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbh := defaultOneDriveBH(user)
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
							id(drivePfx, 1): prevDelta,
							d2.id:           prevDelta,
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

func (suite *CollectionsUnitSuite) TestAddURLCacheToDriveCollections() {
	d := drive(1)
	d2 := drive(2)

	table := []struct {
		name       string
		enumerator enumerateDriveItemsDelta
		errCheck   assert.ErrorAssertionFunc
	}{
		{
			name: "Two drives with unique url cache instances",
			enumerator: driveEnumerator(
				d.newEnumer().with(
					delta(nil).with(
						aPage(
							driveFolder(d.dir(), rootID),
							driveFile(d.dir(folderName()), folderID())))),
				d2.newEnumer().with(
					delta(nil, 2).with(
						aPage(
							driveItem(folderID(2), folderName(), d2.dir(), rootID, isFolder),
							driveItem(fileID(2), fileName(), d2.dir(folderName()), folderID(2), isFile))))),
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

			mbh := defaultOneDriveBH(user)
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
				len(test.enumerator.getDrives()),
				len(caches),
				"expected one cache per drive")
		})
	}
}
