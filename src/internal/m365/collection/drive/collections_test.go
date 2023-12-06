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
				driveRootFolder(),
				driveItem(id(item), name(item), driveParentDir(drive), rootID, -1),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, driveFullPath(drive)),
			},
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single File",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFile(driveParentDir(drive), rootID),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, driveFullPath(drive)),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// Root folder is skipped since it's always present.
			expectedPrevPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			expectedExcludes:         makeExcludeMap(fileID()),
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asNew(t, driveFullPath(drive, folderName())),
			},
			expectedPrevPaths: map[string]string{
				rootID:     driveFullPath(drive),
				folderID(): driveFullPath(drive, folderName()),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder created twice", // deleted a created with same name in between a backup
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(folderID(2), folderName(), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:      asNotMoved(t, driveFullPath(drive)),
				folderID(2): asNew(t, driveFullPath(drive, folderName())),
			},
			expectedPrevPaths: map[string]string{
				rootID:      driveFullPath(drive),
				folderID(2): driveFullPath(drive, folderName()),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Package",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:  asNotMoved(t, driveFullPath(drive)),
				id(pkg): asNew(t, driveFullPath(drive, name(pkg))),
			},
			expectedPrevPaths: map[string]string{
				rootID:  driveFullPath(drive),
				id(pkg): driveFullPath(drive, name(pkg)),
			},
			expectedItemCount:      1,
			expectedContainerCount: 2,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				driveFullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 1,
		},
		{
			name: "Single Package with subfolder",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
				driveItem(folderID(), folderName(), driveParentDir(drive, name(pkg)), id(pkg), isFolder),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive, name(pkg)), id(pkg), isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, driveFullPath(drive)),
				id(pkg):       asNew(t, driveFullPath(drive, name(pkg))),
				folderID():    asNew(t, driveFullPath(drive, name(pkg), folderName())),
				id(subfolder): asNew(t, driveFullPath(drive, name(pkg), name(subfolder))),
			},
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				id(pkg):       driveFullPath(drive, name(pkg)),
				folderID():    driveFullPath(drive, name(pkg), folderName()),
				id(subfolder): driveFullPath(drive, name(pkg), name(subfolder)),
			},
			expectedItemCount:      3,
			expectedContainerCount: 4,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				driveFullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 3,
		},
		{
			name: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFile(driveParentDir(drive), rootID, "inRoot"),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
				driveFile(driveParentDir(drive, folderName()), folderID(), "inFolder"),
				driveFile(driveParentDir(drive, name(pkg)), id(pkg), "inPackage"),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asNew(t, driveFullPath(drive, folderName())),
				id(pkg):    asNew(t, driveFullPath(drive, name(pkg))),
			},
			expectedItemCount:      5,
			expectedFileCount:      3,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:     driveFullPath(drive),
				folderID(): driveFullPath(drive, folderName()),
				id(pkg):    driveFullPath(drive, name(pkg)),
			},
			expectedTopLevelPackages: map[string]struct{}{
				driveFullPath(drive, name(pkg)): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap(fileID("inRoot"), fileID("inFolder"), fileID("inPackage")),
		},
		{
			name: "contains folder selector",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFile(driveParentDir(drive), rootID, "inRoot"),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive, folderName()), folderID(), isFolder),
				driveItem(folderID(2), folderName(), driveParentDir(drive, folderName(), name(subfolder)), id(subfolder), isFolder),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
				driveItem(fileID("inFolder"), fileID("inFolder"), driveParentDir(drive, folderName()), folderID(), isFile),
				driveItem(fileID("inFolder2"), fileName("inFolder2"), driveParentDir(drive, folderName(), name(subfolder), folderName()), folderID(2), isFile),
				driveItem(fileID("inFolderPackage"), fileName("inPackage"), driveParentDir(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{folderName()})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				folderID():    asNew(t, driveFullPath(drive, folderName())),
				id(subfolder): asNew(t, driveFullPath(drive, folderName(), name(subfolder))),
				folderID(2):   asNew(t, driveFullPath(drive, folderName(), name(subfolder), folderName())),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			// just "folder" isn't added here because the include check is done on the
			// parent path since we only check later if something is a folder or not.
			expectedPrevPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
				folderID(2):   driveFullPath(drive, folderName(), name(subfolder), folderName()),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID("inFolder"), fileID("inFolder2")),
		},
		{
			name: "prefix subfolder selector",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFile(driveParentDir(drive), rootID, "inRoot"),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive, folderName()), folderID(), isFolder),
				driveItem(folderID(2), folderName(), driveParentDir(drive, folderName(), name(subfolder)), id(subfolder), isFolder),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
				driveItem(fileID("inFolder"), fileID("inFolder"), driveParentDir(drive, folderName()), folderID(), isFile),
				driveItem(fileID("inFolder2"), fileName("inFolder2"), driveParentDir(drive, folderName(), name(subfolder), folderName()), folderID(2), isFile),
				driveItem(fileID("inFolderPackage"), fileName("inPackage"), driveParentDir(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths: map[string]string{},
			scope: (&selectors.OneDriveBackup{}).Folders(
				[]string{toPath(folderName(), name(subfolder))},
				selectors.PrefixMatch())[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(subfolder): asNew(t, driveFullPath(drive, folderName(), name(subfolder))),
				folderID(2):   asNew(t, driveFullPath(drive, folderName(), name(subfolder), folderName())),
			},
			expectedItemCount:      3,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
				folderID(2):   driveFullPath(drive, folderName(), name(subfolder), folderName()),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID("inFolder2")),
		},
		{
			name: "match subfolder selector",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFile(driveParentDir(drive), rootID),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive, folderName()), folderID(), isFolder),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
				driveItem(fileID(1), fileName(1), driveParentDir(drive, folderName()), folderID(), isFile),
				driveItem(fileID("inSubfolder"), fileName("inSubfolder"), driveParentDir(drive, folderName(), name(subfolder)), id(subfolder), isFile),
				driveItem(fileID(9), fileName(9), driveParentDir(drive, name(pkg)), id(pkg), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{toPath(folderName(), name(subfolder))})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				id(subfolder): asNew(t, driveFullPath(drive, folderName(), name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// No child folders for subfolder so nothing here.
			expectedPrevPaths: map[string]string{
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID("inSubfolder")),
		},
		{
			name: "not moved folder tree",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asNotMoved(t, driveFullPath(drive, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asMoved(t, driveFullPath(drive, folderName("a")), driveFullPath(drive, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(folderID(1), folderName(), driveParentDir(drive), rootID, isFolder),
				driveItem(folderID(2), folderName(), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID(1):   driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:      asNotMoved(t, driveFullPath(drive)),
				folderID(2): asNew(t, driveFullPath(drive, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID(2):   driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup",
			items: []models.DriveItemable{
				driveRootFolder(),
				delItem(folderID(), rootID, isFolder),
				driveItem(folderID(), name(drive), driveParentDir(drive), rootID, isFolder),
				delItem(folderID(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asDeleted(t, driveFullPath(drive, "")),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup including delete",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
				delItem(folderID(), rootID, isFolder),
				driveItem(folderID(2), folderName(), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:      asNotMoved(t, driveFullPath(drive)),
				folderID(2): asNew(t, driveFullPath(drive, folderName())),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID(2):   driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup with addition",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(folderID(1), folderName(), driveParentDir(drive), rootID, isFolder),
				delItem(folderID(1), rootID, isFolder),
				driveItem(folderID(2), folderName(), driveParentDir(drive), rootID, isFolder),
				delItem(folderID(2), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID(1):   driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, driveFullPath(drive)),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree with file no previous",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(fileID(), fileName(), driveParentDir(drive, folderName()), folderID(), isFile),
				driveItem(folderID(), folderName(2), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asNew(t, driveFullPath(drive, folderName(2))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:     driveFullPath(drive),
				folderID(): driveFullPath(drive, folderName(2)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID()),
		},
		{
			name: "moved folder tree with file no previous 1",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(fileID(), fileName(), driveParentDir(drive, folderName()), folderID(), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asNew(t, driveFullPath(drive, folderName())),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:     driveFullPath(drive),
				folderID(): driveFullPath(drive, folderName()),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID()),
		},
		{
			name: "moved folder tree and subfolder 1",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, driveFullPath(drive)),
				folderID():    asMoved(t, driveFullPath(drive, folderName("a")), driveFullPath(drive, folderName())),
				id(subfolder): asMoved(t, driveFullPath(drive, folderName("a"), name(subfolder)), driveFullPath(drive, name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree and subfolder 2",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive), rootID, isFolder),
				driveFolder(driveParentDir(drive), rootID),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, driveFullPath(drive)),
				folderID():    asMoved(t, driveFullPath(drive, folderName("a")), driveFullPath(drive, folderName())),
				id(subfolder): asMoved(t, driveFullPath(drive, folderName("a"), name(subfolder)), driveFullPath(drive, name(subfolder))),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "move subfolder when moving parent",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(folderID(2), folderName(2), driveParentDir(drive), rootID, isFolder),
				driveItem(id(item), name(item), driveParentDir(drive, folderName(2)), folderID(2), isFile),
				// Need to see the parent folder first (expected since that's what Graph
				// consistently returns).
				driveItem(folderID(), folderName("a"), driveParentDir(drive), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive, folderName("a")), folderID(), isFolder),
				driveItem(id(item, 2), name(item, 2), driveParentDir(drive, folderName("a"), name(subfolder)), id(subfolder), isFile),
				driveFolder(driveParentDir(drive), rootID),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, driveFullPath(drive)),
				folderID(2):   asNew(t, driveFullPath(drive, folderName(2))),
				folderID():    asMoved(t, driveFullPath(drive, folderName("a")), driveFullPath(drive, folderName())),
				id(subfolder): asMoved(t, driveFullPath(drive, folderName("a"), name(subfolder)), driveFullPath(drive, folderName(), name(subfolder))),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 4,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName()),
				folderID(2):   driveFullPath(drive, folderName(2)),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(item), id(item, 2)),
		},
		{
			name: "moved folder tree multiple times",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(fileID(), fileName(), driveParentDir(drive, folderName()), folderID(), isFile),
				driveItem(folderID(), folderName(2), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				folderID():    driveFullPath(drive, folderName("a")),
				id(subfolder): driveFullPath(drive, folderName("a"), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asMoved(t, driveFullPath(drive, folderName("a")), driveFullPath(drive, folderName(2))),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName(2)),
				id(subfolder): driveFullPath(drive, folderName(2), name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(fileID()),
		},
		{
			name: "deleted folder and package",
			items: []models.DriveItemable{
				driveRootFolder(), // root is always present, but not necessary here
				delItem(folderID(), rootID, isFolder),
				delItem(id(pkg), rootID, isPackage),
			},
			previousPaths: map[string]string{
				rootID:     driveFullPath(drive),
				folderID(): driveFullPath(drive, folderName()),
				id(pkg):    driveFullPath(drive, name(pkg)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asDeleted(t, driveFullPath(drive, folderName())),
				id(pkg):    asDeleted(t, driveFullPath(drive, name(pkg))),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder without previous",
			items: []models.DriveItemable{
				driveRootFolder(),
				delItem(folderID(), rootID, isFolder),
			},
			previousPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, driveFullPath(drive)),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder tree move subfolder",
			items: []models.DriveItemable{
				driveRootFolder(),
				delItem(folderID(), rootID, isFolder),
				driveItem(id(subfolder), name(subfolder), driveParentDir(drive), rootID, isFolder),
			},
			previousPaths: map[string]string{
				rootID:        driveFullPath(drive),
				folderID():    driveFullPath(drive, folderName()),
				id(subfolder): driveFullPath(drive, folderName(), name(subfolder)),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:        asNotMoved(t, driveFullPath(drive)),
				folderID():    asDeleted(t, driveFullPath(drive, folderName())),
				id(subfolder): asMoved(t, driveFullPath(drive, folderName(), name(subfolder)), driveFullPath(drive, name(subfolder))),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				rootID:        driveFullPath(drive),
				id(subfolder): driveFullPath(drive, name(subfolder)),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete file",
			items: []models.DriveItemable{
				driveRootFolder(),
				delItem(id(item), rootID, isFile),
			},
			previousPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, driveFullPath(drive)),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap(id(item)),
		},
		{
			name: "item before parent errors",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(fileID(), fileName(), driveParentDir(drive, folderName()), folderID(), isFile),
				driveFolder(driveParentDir(drive), rootID),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				rootID: asNotMoved(t, driveFullPath(drive)),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				rootID: driveFullPath(drive),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "1 root file, 1 folder, 1 package, 1 good file, 1 malware",
			items: []models.DriveItemable{
				driveRootFolder(),
				driveItem(fileID(), fileID(), driveParentDir(drive), rootID, isFile),
				driveFolder(driveParentDir(drive), rootID),
				driveItem(id(pkg), name(pkg), driveParentDir(drive), rootID, isPackage),
				driveItem(fileID("good"), fileName("good"), driveParentDir(drive, folderName()), folderID(), isFile),
				malwareItem(id(malware), name(malware), driveParentDir(drive, folderName()), folderID(), isFile),
			},
			previousPaths:    map[string]string{},
			scope:            anyFolderScope,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				rootID:     asNotMoved(t, driveFullPath(drive)),
				folderID(): asNew(t, driveFullPath(drive, folderName())),
				id(pkg):    asNew(t, driveFullPath(drive, name(pkg))),
			},
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			expectedSkippedCount:   1,
			expectedPrevPaths: map[string]string{
				rootID:     driveFullPath(drive),
				folderID(): driveFullPath(drive, folderName()),
				id(pkg):    driveFullPath(drive, name(pkg)),
			},
			expectedTopLevelPackages: map[string]struct{}{
				driveFullPath(drive, name(pkg)): {},
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
				drive    = mock.Drive()
				mbh      = mock.DefaultOneDriveBH(user)
				excludes = map[string]struct{}{}
				errs     = fault.New(true)
			)

			mbh.DriveItemEnumeration = mock.DriveEnumerator(
				drive.NewEnumer().With(
					mock.Delta("notempty", nil).With(
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

			c.CollectionMap[drive.ID] = map[string]*Collection{}

			_, newPrevPaths, err := c.PopulateDriveCollections(
				ctx,
				drive.ID,
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
				maps.Keys(c.CollectionMap[drive.ID]),
				"expected collection IDs")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			assert.Equal(t, test.expectedSkippedCount, len(errs.Skipped()), "skipped item count")

			for id, sp := range test.expectedCollectionIDs {
				if !assert.Containsf(t, c.CollectionMap[drive.ID], id, "missing collection with id %s", id) {
					// Skip collections we don't find so we don't get an NPE.
					continue
				}

				assert.Equalf(t, sp.state, c.CollectionMap[drive.ID][id].State(), "state for collection %s", id)
				assert.Equalf(t, sp.currPath, c.CollectionMap[drive.ID][id].FullPath(), "current path for collection %s", id)
				assert.Equalf(t, sp.prevPath, c.CollectionMap[drive.ID][id].PreviousPath(), "prev path for collection %s", id)
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
									folderID(1): driveFullPath(1),
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
					folderID(1): driveFullPath(1),
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
									folderID(1): driveFullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					folderID(1): driveFullPath(1),
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
									folderID(1): driveFullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{id(drive): ""},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					folderID(1): driveFullPath(1),
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
									folderID(1): driveFullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive, 2): id(delta, 2)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive, 2): {
									folderID(2): driveFullPath(2),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive):    id(delta),
				id(drive, 2): id(delta, 2),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					folderID(1): driveFullPath(1),
				},
				id(drive, 2): {
					folderID(2): driveFullPath(2),
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
									folderID(1): driveFullPath(1),
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
					folderID(1): driveFullPath(1),
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
									folderID(1): driveFullPath(1),
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
									folderID(2): driveFullPath(2),
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
									folderID(1): driveFullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive): id(delta, 2)}),
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
									folderID(1): driveFullPath(1),
									folderID(2): driveFullPath(1),
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
					folderID(1): driveFullPath(1),
					folderID(2): driveFullPath(1),
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
									folderID(1): driveFullPath(1),
									folderID(2): driveFullPath(1),
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{id(drive, 2): id(delta, 2)}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								id(drive, 2): {
									folderID(1): driveFullPath(1),
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				id(drive):    id(delta),
				id(drive, 2): id(delta, 2),
			},
			expectedPaths: map[string]map[string]string{
				id(drive): {
					folderID(1): driveFullPath(1),
					folderID(2): driveFullPath(1),
				},
				id(drive, 2): {
					folderID(1): driveFullPath(1),
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

	mbh := mock.DefaultOneDriveBH(user)
	opts := control.DefaultOptions()
	opts.ToggleFeatures.UseDeltaTree = true

	mbh.DriveItemEnumeration = mock.DriveEnumerator(
		mock.Drive().NewEnumer().With(
			mock.Delta(id(delta), nil).With(
				aPage(delItem(fileID(), rootID, isFile)))))

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

	drive1 := mock.Drive(1)
	drive2 := mock.Drive(2)

	table := []struct {
		name                 string
		enumerator           mock.EnumerateDriveItemsDelta
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
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {rootID: driveFullPath(1)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NotMovedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {rootID: driveFullPath(1)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				driveFullPath(1): makeExcludeMap(fileID()),
			}),
		},
		{
			name: "OneDrive_OneItemPage_NoFolderDeltas_NoErrors",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						driveFile(driveParentDir(1), rootID))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {rootID: driveFullPath(1)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NotMovedState: {fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {rootID: driveFullPath(1)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				driveFullPath(1): makeExcludeMap(fileID()),
			}),
		},
		{
			name: "OneDrive_OneItemPage_NoErrors",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID()))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_NoErrors_FileRenamedMultiple",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID()),
						driveItem(fileID(), fileName(2), driveParentDir(1, folderName()), folderID(), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_NoErrors_FileMovedMultiple",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID()),
						driveItem(fileID(), fileName(2), driveParentDir(1), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NotMovedState: {fileID()}},
				driveFullPath(1, folderName()): {data.NewState: {folderID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				driveFullPath(1): makeExcludeMap(fileID()),
			}),
		},
		{
			name: "OneDrive_TwoItemPages_NoErrors",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID(), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_TwoItemPages_WithReset",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID()),
							driveItem(fileID(3), fileName(3), driveParentDir(1, folderName()), folderID(), isFile)),
						aReset(),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID(), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_TwoItemPages_WithResetCombinedWithItems",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPageWReset(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID(), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "TwoDrives_OneItemPageEach_NoErrors",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID())))),
				drive2.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(aPage(
						driveItem(folderID(2), folderName(), driveParentDir(2), rootID, isFolder),
						driveItem(fileID(2), fileName(), driveParentDir(2, folderName()), folderID(2), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
				id(drive, 2): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID()}},
				driveFullPath(2):               {data.NewState: {}},
				driveFullPath(2, folderName()): {data.NewState: {folderID(2), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
				id(drive, 2): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
				id(drive, 2): {
					rootID:      driveFullPath(2),
					folderID(2): driveFullPath(2, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
				driveFullPath(2):               true,
				driveFullPath(2, folderName()): true,
			},
		},
		{
			name: "TwoDrives_DuplicateIDs_OneItemPageEach_NoErrors",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID())))),
				drive2.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(aPage(
						driveFolder(driveParentDir(2), rootID),
						driveItem(fileID(2), fileName(), driveParentDir(2, folderName()), folderID(), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
				id(drive, 2): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID()}},
				driveFullPath(2):               {data.NewState: {}},
				driveFullPath(2, folderName()): {data.NewState: {folderID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
				id(drive, 2): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
				id(drive, 2): {
					rootID:     driveFullPath(2),
					folderID(): driveFullPath(2, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
				driveFullPath(2):               true,
				driveFullPath(2, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_Errors",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta("", assert.AnError))),
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections:   nil,
			expectedDeltaURLs:     nil,
			expectedPreviousPaths: nil,
			expectedDelList:       nil,
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_DeleteNonExistentFolder",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage(
							driveFolder(driveParentDir(1), rootID, 2),
							driveFile(driveParentDir(1, folderName(2)), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):                {data.NewState: {}},
				driveFullPath(1, folderName()):  {data.DeletedState: {}},
				driveFullPath(1, folderName(2)): {data.NewState: {folderID(2), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(2): driveFullPath(1, folderName(2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):                true,
				driveFullPath(1, folderName()):  true,
				driveFullPath(1, folderName(2)): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDeltaCombinedWithItems_DeleteNonExistentFolder",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage(
							driveFolder(driveParentDir(1), rootID, 2),
							driveFile(driveParentDir(1, folderName(2)), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):                {data.NewState: {}},
				driveFullPath(1, folderName()):  {data.DeletedState: {}},
				driveFullPath(1, folderName(2)): {data.NewState: {folderID(2), fileID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(2): driveFullPath(1, folderName(2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):                true,
				driveFullPath(1, folderName()):  true,
				driveFullPath(1, folderName(2)): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveItem(folderID(2), folderName(), driveParentDir(1), rootID, isFolder),
							driveFile(driveParentDir(1, folderName()), folderID(2))),
						aReset(),
						aPage(
							driveItem(folderID(2), folderName(), driveParentDir(1), rootID, isFolder),
							driveFile(driveParentDir(1, folderName()), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
				driveFullPath(1, folderName()): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {folderID(2), fileID()},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(2): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtExistingLocation",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aReset(),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID()))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
				driveFullPath(1, folderName()): {
					data.NewState: {folderID(), fileID()},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_ImmediateInvalidPrevDelta_MoveFolderToPreviouslyExistingPath",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage(
							driveItem(folderID(2), folderName(), driveParentDir(1), rootID, isFolder),
							driveItem(fileID(2), fileName(), driveParentDir(1, folderName()), folderID(2), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
				driveFullPath(1, folderName()): {
					data.DeletedState: {},
					data.NewState:     {folderID(2), fileID(2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(2): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aReset(),
						aPage(
							driveItem(folderID(2), folderName(), driveParentDir(1), rootID, isFolder),
							driveFile(driveParentDir(1, folderName()), folderID(2)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
				driveFullPath(1, folderName()): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {folderID(2), fileID()},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(2): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "OneDrive Two Item Pages with Malware",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID()),
							malwareItem(id(malware), name(malware), driveParentDir(1, folderName()), folderID(), isFile)),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID(), 2),
							malwareItem(id(malware, 2), name(malware, 2), driveParentDir(1, folderName()), folderID(), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(), fileID(), fileID(2)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
			expectedSkippedCount: 2,
		},
		{
			name: "One Drive Deleted Folder In New Results With Invalid Delta",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID()),
							driveFolder(driveParentDir(1), rootID, 2),
							driveFile(driveParentDir(1, folderName(2)), folderID(2), 2)),
						aReset(),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID()),
							delItem(folderID(2), rootID, isFolder),
							delItem(fileName(2), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID():  driveFullPath(1, folderName()),
					folderID(2): driveFullPath(1, folderName(2)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):                {data.NewState: {}},
				driveFullPath(1, folderName()):  {data.NewState: {folderID(), fileID()}},
				driveFullPath(1, folderName(2)): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):                true,
				driveFullPath(1, folderName()):  true,
				driveFullPath(1, folderName(2)): true,
			},
		},
		{
			name: "One Drive Folder Delete After Invalid Delta",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPageWReset(
						delItem(folderID(), rootID, isFolder))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "One Drive Item Delete After Invalid Delta",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPageWReset(
						delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1): true,
			},
		},
		{
			name: "One Drive Folder Made And Deleted",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1): true,
			},
		},
		{
			name: "One Drive Folder Created -> Deleted -> Created",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile)),
						aPage(
							driveItem(folderID(1), folderName(), driveParentDir(1), rootID, isFolder),
							driveItem(fileID(1), fileName(), driveParentDir(1, folderName()), folderID(1), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID(1), fileID(1)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(1): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "One Drive Folder Deleted -> Created -> Deleted",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile)),
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NotMovedState: {}},
				driveFullPath(1, folderName()): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name: "One Drive Folder Created -> Deleted -> Created with prev",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta, 2), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(
							delItem(folderID(), rootID, isFolder),
							delItem(fileID(), rootID, isFile)),
						aPage(
							driveItem(folderID(1), folderName(), driveParentDir(1), rootID, isFolder),
							driveItem(fileID(1), fileName(), driveParentDir(1, folderName()), folderID(1), isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.DeletedState: {}, data.NewState: {folderID(1), fileID(1)}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID(1): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               false,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "One Drive Item Made And Deleted",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(
						aPage(
							driveFolder(driveParentDir(1), rootID),
							driveFile(driveParentDir(1, folderName()), folderID())),
						aPage(delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1):               {data.NewState: {}},
				driveFullPath(1, folderName()): {data.NewState: {folderID()}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:     driveFullPath(1),
					folderID(): driveFullPath(1, folderName()),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1):               true,
				driveFullPath(1, folderName()): true,
			},
		},
		{
			name: "One Drive Random Folder Delete",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.DeltaWReset(id(delta), nil).With(aPage(
						delItem(folderID(), rootID, isFolder))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1): true,
			},
		},
		{
			name: "One Drive Random Item Delete",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						delItem(fileID(), rootID, isFile))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID: driveFullPath(1),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(1): true,
			},
		},
		{
			name: "TwoPriorDrives_OneTombstoned",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage()))), // root only
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {rootID: driveFullPath(1)},
				id(drive, 2): {rootID: driveFullPath(2)},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {data.NotMovedState: {}},
				driveFullPath(2): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{id(drive, 1): id(delta)},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {rootID: driveFullPath(1)},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				driveFullPath(2): true,
			},
		},
		{
			name: "duplicate previous paths in metadata",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID()),
						driveFolder(driveParentDir(1), rootID, 2),
						driveFile(driveParentDir(1, folderName(2)), folderID(2), 2)))),
				drive2.NewEnumer().With(
					mock.Delta(id(delta, 2), nil).With(aPage(
						driveFolder(driveParentDir(2), rootID),
						driveFile(driveParentDir(2, folderName()), folderID()),
						driveFolder(driveParentDir(2), rootID, 2),
						driveFile(driveParentDir(2, folderName(2)), folderID(2), 2))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID():  driveFullPath(1, folderName()),
					folderID(2): driveFullPath(1, folderName()),
					folderID(3): driveFullPath(1, folderName()),
				},
				id(drive, 2): {
					rootID:      driveFullPath(2),
					folderID():  driveFullPath(2, folderName()),
					folderID(2): driveFullPath(2, folderName(2)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {
					data.NewState: {folderID(), folderID(2)},
				},
				driveFullPath(1, folderName()): {
					data.NotMovedState: {folderID(), fileID()},
				},
				driveFullPath(1, folderName(2)): {
					data.MovedState: {folderID(2), fileID(2)},
				},
				driveFullPath(2): {
					data.NewState: {folderID(), folderID(2)},
				},
				driveFullPath(2, folderName()): {
					data.NotMovedState: {folderID(), fileID()},
				},
				driveFullPath(2, folderName(2)): {
					data.NotMovedState: {folderID(2), fileID(2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
				id(drive, 2): id(delta, 2),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:      driveFullPath(1),
					folderID():  driveFullPath(1, folderName(2)), // note: this is a bug, but is currently expected
					folderID(2): driveFullPath(1, folderName(2)),
					folderID(3): driveFullPath(1, folderName(2)),
				},
				id(drive, 2): {
					rootID:      driveFullPath(2),
					folderID():  driveFullPath(2, folderName()),
					folderID(2): driveFullPath(2, folderName(2)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				driveFullPath(1): makeExcludeMap(fileID(), fileID(2)),
				driveFullPath(2): makeExcludeMap(fileID(), fileID(2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name: "out of order item enumeration causes prev path collisions",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						driveItem(folderID(fanny, 2), folderName(fanny), driveParentDir(1), rootID, isFolder),
						driveFile(driveParentDir(1, folderName(fanny)), folderID(fanny, 2), 2),
						driveFolder(driveParentDir(1), rootID, nav),
						driveFile(driveParentDir(1, folderName(nav)), folderID(nav)))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:        driveFullPath(1),
					folderID(nav): driveFullPath(1, folderName(fanny)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {
					data.NewState: {folderID(fanny, 2)},
				},
				driveFullPath(1, folderName(nav)): {
					data.MovedState: {folderID(nav), fileID()},
				},
				driveFullPath(1, folderName(fanny)): {
					data.NewState: {folderID(fanny, 2), fileID(2)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:             driveFullPath(1),
					folderID(nav):      driveFullPath(1, folderName(nav)),
					folderID(fanny, 2): driveFullPath(1, folderName(nav)), // note: this is a bug, but currently expected
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				driveFullPath(1): makeExcludeMap(fileID(), fileID(2)),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name: "out of order item enumeration causes opposite prev path collisions",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						driveFile(driveParentDir(1), rootID, 1),
						driveFolder(driveParentDir(1), rootID, fanny),
						driveFolder(driveParentDir(1), rootID, nav),
						driveFolder(driveParentDir(1, folderName(fanny)), folderID(fanny), foo),
						driveItem(folderID(bar), folderName(foo), driveParentDir(1, folderName(nav)), folderID(nav), isFolder))))),
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:          driveFullPath(1),
					folderID(nav):   driveFullPath(1, folderName(nav)),
					folderID(fanny): driveFullPath(1, folderName(fanny)),
					folderID(foo):   driveFullPath(1, folderName(nav), folderName(foo)),
					folderID(bar):   driveFullPath(1, folderName(fanny), folderName(foo)),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				driveFullPath(1): {
					data.NotMovedState: {fileID(1)},
				},
				driveFullPath(1, folderName(nav)): {
					data.NotMovedState: {folderID(nav)},
				},
				driveFullPath(1, folderName(nav), folderName(foo)): {
					data.MovedState: {folderID(bar)},
				},
				driveFullPath(1, folderName(fanny)): {
					data.NotMovedState: {folderID(fanny)},
				},
				driveFullPath(1, folderName(fanny), folderName(foo)): {
					data.MovedState: {folderID(foo)},
				},
			},
			expectedDeltaURLs: map[string]string{
				id(drive, 1): id(delta),
			},
			expectedPreviousPaths: map[string]map[string]string{
				id(drive, 1): {
					rootID:          driveFullPath(1),
					folderID(nav):   driveFullPath(1, folderName(nav)),
					folderID(fanny): driveFullPath(1, folderName(fanny)),
					folderID(foo):   driveFullPath(1, folderName(nav), folderName(foo)), // note: this is a bug, but currently expected
					folderID(bar):   driveFullPath(1, folderName(nav), folderName(foo)),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				driveFullPath(1): makeExcludeMap(fileID(1)),
			}),
			doNotMergeItems: map[string]bool{},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbh := mock.DefaultOneDriveBH(user)
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
							id(drive, 1): prevDelta,
							id(drive, 2): prevDelta,
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
	drive1 := mock.Drive(1)
	drive2 := mock.Drive(2)

	table := []struct {
		name       string
		enumerator mock.EnumerateDriveItemsDelta
		errCheck   assert.ErrorAssertionFunc
	}{
		{
			name: "Two drives with unique url cache instances",
			enumerator: mock.DriveEnumerator(
				drive1.NewEnumer().With(
					mock.Delta(id(delta), nil).With(aPage(
						driveFolder(driveParentDir(1), rootID),
						driveFile(driveParentDir(1, folderName()), folderID())))),
				drive2.NewEnumer().With(
					mock.Delta(id(delta, 2), nil).With(aPage(
						driveItem(folderID(2), folderName(), driveParentDir(2), rootID, isFolder),
						driveItem(fileID(2), fileName(), driveParentDir(2, folderName()), folderID(2), isFile))))),
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

			mbh := mock.DefaultOneDriveBH(user)
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
				len(test.enumerator.Drives()),
				len(caches),
				"expected one cache per drive")
		})
	}
}
