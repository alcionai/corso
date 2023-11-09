package drive

import (
	"context"
	"strconv"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
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
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

type statePath struct {
	state    data.CollectionState
	curPath  path.Path
	prevPath path.Path
}

func getExpectedStatePathGenerator(
	t *testing.T,
	bh BackupHandler,
	tenant, base string,
) func(data.CollectionState, ...string) statePath {
	return func(state data.CollectionState, pths ...string) statePath {
		var (
			p1  path.Path
			p2  path.Path
			pp  path.Path
			cp  path.Path
			err error
		)

		if state != data.MovedState {
			require.Len(t, pths, 1, "invalid number of paths to getExpectedStatePathGenerator")
		} else {
			require.Len(t, pths, 2, "invalid number of paths to getExpectedStatePathGenerator")
			pb := path.Builder{}.Append(path.Split(base + pths[1])...)
			p2, err = bh.CanonicalPath(pb, tenant)
			require.NoError(t, err, clues.ToCore(err))
		}

		pb := path.Builder{}.Append(path.Split(base + pths[0])...)
		p1, err = bh.CanonicalPath(pb, tenant)
		require.NoError(t, err, clues.ToCore(err))

		switch state {
		case data.NewState:
			cp = p1
		case data.NotMovedState:
			cp = p1
			pp = p1
		case data.DeletedState:
			pp = p1
		case data.MovedState:
			pp = p2
			cp = p1
		}

		return statePath{
			state:    state,
			curPath:  cp,
			prevPath: pp,
		}
	}
}

func getExpectedPathGenerator(
	t *testing.T,
	bh BackupHandler,
	tenant, base string,
) func(string) string {
	return func(p string) string {
		pb := path.Builder{}.Append(path.Split(base + p)...)
		cp, err := bh.CanonicalPath(pb, tenant)
		require.NoError(t, err, clues.ToCore(err))

		return cp.String()
	}
}

type OneDriveCollectionsUnitSuite struct {
	tester.Suite
}

func TestOneDriveCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, &OneDriveCollectionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func makeExcludeMap(files ...string) map[string]struct{} {
	delList := map[string]struct{}{}
	for _, file := range files {
		delList[file+metadata.DataFileSuffix] = struct{}{}
		delList[file+metadata.MetaFileSuffix] = struct{}{}
	}

	return delList
}

func (suite *OneDriveCollectionsUnitSuite) TestPopulateDriveCollections() {
	anyFolder := (&selectors.OneDriveBackup{}).Folders(selectors.Any())[0]

	const (
		driveID   = "driveID1"
		tenant    = "tenant"
		user      = "user"
		folder    = "/folder"
		subFolder = "/subfolder"
		pkg       = "/package"
	)

	bh := userDriveBackupHandler{userID: user}
	testBaseDrivePath := odConsts.DriveFolderPrefixBuilder("driveID1").String()
	expectedPath := getExpectedPathGenerator(suite.T(), bh, tenant, testBaseDrivePath)
	expectedStatePath := getExpectedStatePathGenerator(suite.T(), bh, tenant, testBaseDrivePath)

	tests := []struct {
		name                     string
		items                    []models.DriveItemable
		inputFolderMap           map[string]string
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
				driveRootItem("root"),
				driveItem("item", "item", testBaseDrivePath, "root", false, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single File",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("file", "file", testBaseDrivePath, "root", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// Root folder is skipped since it's always present.
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes:         makeExcludeMap("file"),
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NewState, folder),
			},
			expectedPrevPaths: map[string]string{
				"root":   expectedPath(""),
				"folder": expectedPath("/folder"),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Folder created twice", // deleted a created with same name in between a backup
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("id1", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("id2", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
				"id2":  expectedStatePath(data.NewState, folder),
			},
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
				"id2":  expectedPath("/folder"),
			},
			expectedItemCount:        1,
			expectedContainerCount:   2,
			expectedExcludes:         map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{},
		},
		{
			name: "Single Package",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"package": expectedStatePath(data.NewState, pkg),
			},
			expectedPrevPaths: map[string]string{
				"root":    expectedPath(""),
				"package": expectedPath("/package"),
			},
			expectedItemCount:      1,
			expectedContainerCount: 2,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				expectedPath("/package"): {},
			},
			expectedCountPackages: 1,
		},
		{
			name: "Single Package with subfolder",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("folder", "folder", testBaseDrivePath+pkg, "package", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+pkg, "package", false, true, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"package":   expectedStatePath(data.NewState, pkg),
				"folder":    expectedStatePath(data.NewState, pkg+folder),
				"subfolder": expectedStatePath(data.NewState, pkg+subFolder),
			},
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"package":   expectedPath(pkg),
				"folder":    expectedPath(pkg + folder),
				"subfolder": expectedPath(pkg + subFolder),
			},
			expectedItemCount:      3,
			expectedContainerCount: 4,
			expectedExcludes:       map[string]struct{}{},
			expectedTopLevelPackages: map[string]struct{}{
				expectedPath(pkg): {},
			},
			expectedCountPackages: 3,
		},
		{
			name: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"folder":  expectedStatePath(data.NewState, folder),
				"package": expectedStatePath(data.NewState, pkg),
			},
			expectedItemCount:      5,
			expectedFileCount:      3,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				"root":    expectedPath(""),
				"folder":  expectedPath("/folder"),
				"package": expectedPath("/package"),
			},
			expectedTopLevelPackages: map[string]struct{}{
				expectedPath("/package"): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap("fileInRoot", "fileInFolder", "fileInPackage"),
		},
		{
			name: "contains folder selector",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, "folder", false, true, false),
				driveItem("folder2", "folder", testBaseDrivePath+folder+subFolder, "subfolder", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem(
					"fileInFolder2",
					"fileInFolder2",
					testBaseDrivePath+folder+subFolder+folder,
					"folder2",
					true,
					false,
					false),
				driveItem("fileInFolderPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{"folder"})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"folder":    expectedStatePath(data.NewState, folder),
				"subfolder": expectedStatePath(data.NewState, folder+subFolder),
				"folder2":   expectedStatePath(data.NewState, folder+subFolder+folder),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			// just "folder" isn't added here because the include check is done on the
			// parent path since we only check later if something is a folder or not.
			expectedPrevPaths: map[string]string{
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folder + subFolder),
				"folder2":   expectedPath(folder + subFolder + folder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("fileInFolder", "fileInFolder2"),
		},
		{
			name: "prefix subfolder selector",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, "folder", false, true, false),
				driveItem("folder2", "folder", testBaseDrivePath+folder+subFolder, "subfolder", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem(
					"fileInFolder2",
					"fileInFolder2",
					testBaseDrivePath+folder+subFolder+folder,
					"folder2",
					true,
					false,
					false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{"/folder/subfolder"}, selectors.PrefixMatch())[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"subfolder": expectedStatePath(data.NewState, folder+subFolder),
				"folder2":   expectedStatePath(data.NewState, folder+subFolder+folder),
			},
			expectedItemCount:      3,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"subfolder": expectedPath(folder + subFolder),
				"folder2":   expectedPath(folder + subFolder + folder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("fileInFolder2"),
		},
		{
			name: "match subfolder selector",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, "folder", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem(
					"fileInSubfolder",
					"fileInSubfolder",
					testBaseDrivePath+folder+subFolder,
					"subfolder",
					true,
					false,
					false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            (&selectors.OneDriveBackup{}).Folders([]string{"folder/subfolder"})[0],
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"subfolder": expectedStatePath(data.NewState, folder+subFolder),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// No child folders for subfolder so nothing here.
			expectedPrevPaths: map[string]string{
				"subfolder": expectedPath(folder + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("fileInSubfolder"),
		},
		{
			name: "not moved folder tree",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folder + subFolder),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NotMovedState, folder),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folder + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.MovedState, folder, "/a-folder"),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folder + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("id1", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("id2", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"id1":       expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
				"id2":  expectedStatePath(data.NewState, folder),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"id2":       expectedPath(folder),
				"subfolder": expectedPath(folder + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup",
			items: []models.DriveItemable{
				driveRootItem("root"),
				delItem("id1", testBaseDrivePath, "root", false, true, false),
				driveItem("id1", "folder", testBaseDrivePath, "root", false, true, false),
				delItem("id1", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"id1":       expectedPath(""),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
				"id1":  expectedStatePath(data.DeletedState, ""),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"subfolder": expectedPath("/a-folder" + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree twice within backup including delete",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("id1", "folder", testBaseDrivePath, "root", false, true, false),
				delItem("id1", testBaseDrivePath, "root", false, true, false),
				driveItem("id2", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"id1":       expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
				"id2":  expectedStatePath(data.NewState, folder),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"id2":       expectedPath(folder),
				"subfolder": expectedPath(folder + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "deleted folder tree twice within backup with addition",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("id1", "folder", testBaseDrivePath, "root", false, true, false),
				delItem("id1", testBaseDrivePath, "root", false, true, false),
				driveItem("id2", "folder", testBaseDrivePath, "root", false, true, false),
				delItem("id2", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"id1":       expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"subfolder": expectedPath(folder + subFolder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree with file no previous",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
				driveItem("folder", "folder2", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NewState, "/folder2"),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":   expectedPath(""),
				"folder": expectedPath("/folder2"),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("file"),
		},
		{
			name: "moved folder tree with file no previous 1",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NewState, folder),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":   expectedPath(""),
				"folder": expectedPath(folder),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("file"),
		},
		{
			name: "moved folder tree and subfolder 1",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.MovedState, folder, "/a-folder"),
				"subfolder": expectedStatePath(data.MovedState, "/subfolder", "/a-folder/subfolder"),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath("/subfolder"),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "moved folder tree and subfolder 2",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("subfolder", "subfolder", testBaseDrivePath, "root", false, true, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.MovedState, folder, "/a-folder"),
				"subfolder": expectedStatePath(data.MovedState, "/subfolder", "/a-folder/subfolder"),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath("/subfolder"),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "move subfolder when moving parent",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder2", "folder2", testBaseDrivePath, "root", false, true, false),
				driveItem("itemInFolder2", "itemInFolder2", testBaseDrivePath+"/folder2", "folder2", true, false, false),
				// Need to see the parent folder first (expected since that's what Graph
				// consistently returns).
				driveItem("folder", "a-folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+"/a-folder", "folder", false, true, false),
				driveItem(
					"itemInSubfolder",
					"itemInSubfolder",
					testBaseDrivePath+"/a-folder/subfolder",
					"subfolder",
					true,
					false,
					false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.MovedState, folder, "/a-folder"),
				"folder2":   expectedStatePath(data.NewState, "/folder2"),
				"subfolder": expectedStatePath(data.MovedState, folder+subFolder, "/a-folder/subfolder"),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 4,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath("/folder"),
				"folder2":   expectedPath("/folder2"),
				"subfolder": expectedPath("/folder/subfolder"),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("itemInSubfolder", "itemInFolder2"),
		},
		{
			name: "moved folder tree multiple times",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
				driveItem("folder", "folder2", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.MovedState, "/folder2", "/a-folder"),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath("/folder2"),
				"subfolder": expectedPath("/folder2/subfolder"),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("file"),
		},
		{
			name: "deleted folder and package",
			items: []models.DriveItemable{
				driveRootItem("root"), // root is always present, but not necessary here
				delItem("folder", testBaseDrivePath, "root", false, true, false),
				delItem("package", testBaseDrivePath, "root", false, false, true),
			},
			inputFolderMap: map[string]string{
				"root":    expectedPath(""),
				"folder":  expectedPath("/folder"),
				"package": expectedPath("/package"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"folder":  expectedStatePath(data.DeletedState, folder),
				"package": expectedStatePath(data.DeletedState, pkg),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder without previous",
			items: []models.DriveItemable{
				driveRootItem("root"),
				delItem("folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"root": expectedPath(""),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete folder tree move subfolder",
			items: []models.DriveItemable{
				driveRootItem("root"),
				delItem("folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath("/folder"),
				"subfolder": expectedPath("/folder/subfolder"),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.DeletedState, folder),
				"subfolder": expectedStatePath(data.MovedState, "/subfolder", folder+subFolder),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedPrevPaths: map[string]string{
				"root":      expectedPath(""),
				"subfolder": expectedPath("/subfolder"),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "delete file",
			items: []models.DriveItemable{
				driveRootItem("root"),
				delItem("item", testBaseDrivePath, "root", true, false, false),
			},
			inputFolderMap: map[string]string{
				"root": expectedPath(""),
			},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         makeExcludeMap("item"),
		},
		{
			name: "item before parent errors",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.Error,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedPrevPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedTopLevelPackages: map[string]struct{}{},
			expectedExcludes:         map[string]struct{}{},
		},
		{
			name: "1 root file, 1 folder, 1 package, 1 good file, 1 malware",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("goodFile", "goodFile", testBaseDrivePath+folder, "folder", true, false, false),
				malwareItem("malwareFile", "malwareFile", testBaseDrivePath+folder, "folder", true, false, false),
			},
			inputFolderMap:   map[string]string{},
			scope:            anyFolder,
			topLevelPackages: map[string]struct{}{},
			expect:           assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"folder":  expectedStatePath(data.NewState, folder),
				"package": expectedStatePath(data.NewState, pkg),
			},
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			expectedSkippedCount:   1,
			expectedPrevPaths: map[string]string{
				"root":    expectedPath(""),
				"folder":  expectedPath("/folder"),
				"package": expectedPath("/package"),
			},
			expectedTopLevelPackages: map[string]struct{}{
				expectedPath("/package"): {},
			},
			expectedCountPackages: 1,
			expectedExcludes:      makeExcludeMap("fileInRoot", "goodFile"),
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				mbh = mock.DefaultOneDriveBH(user)
				du  = pagers.DeltaUpdate{
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
				control.Options{ToggleFeatures: control.Toggles{}})

			c.CollectionMap[driveID] = map[string]*Collection{}

			_, newPrevPaths, err := c.PopulateDriveCollections(
				ctx,
				driveID,
				"General",
				test.inputFolderMap,
				excludes,
				test.topLevelPackages,
				"prevdelta",
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
				assert.Equalf(t, sp.curPath, c.CollectionMap[driveID][id].FullPath(), "current path for collection %s", id)
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

func (suite *OneDriveCollectionsUnitSuite) TestDeserializeMetadata() {
	tenant := "a-tenant"
	user := "a-user"
	driveID1 := "1"
	driveID2 := "2"
	deltaURL1 := "url/1"
	deltaURL2 := "url/2"
	folderID1 := "folder1"
	folderID2 := "folder2"
	path1 := "folder1/path"
	path2 := "folder2/path"

	table := []struct {
		name string
		// Each function returns the set of files for a single data.Collection.
		cols                 []func() []graph.MetadataCollectionEntry
		expectedDeltas       map[string]string
		expectedPaths        map[string]map[string]string
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				driveID1: deltaURL1,
			},
			expectedPaths: map[string]map[string]string{
				driveID1: {
					folderID1: path1,
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
							map[string]string{driveID1: deltaURL1}),
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
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{},
			expectedPaths: map[string]map[string]string{
				driveID1: {
					folderID1: path1,
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {},
							}),
					}
				},
			},
			expectedDeltas:       map[string]string{},
			expectedPaths:        map[string]map[string]string{driveID1: {}},
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
								driveID1: "",
							}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{driveID1: ""},
			expectedPaths: map[string]map[string]string{
				driveID1: {
					folderID1: path1,
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{driveID2: deltaURL2}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID2: {
									folderID2: path2,
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				driveID1: deltaURL1,
				driveID2: deltaURL2,
			},
			expectedPaths: map[string]map[string]string{
				driveID1: {
					folderID1: path1,
				},
				driveID2: {
					folderID2: path2,
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
							map[string]string{driveID1: deltaURL1}),
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
						graph.NewMetadataEntry(
							"foo",
							map[string]string{driveID1: deltaURL1}),
					}
				},
			},
			expectedDeltas: map[string]string{
				driveID1: deltaURL1,
			},
			expectedPaths: map[string]map[string]string{
				driveID1: {
					folderID1: path1,
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID2: path2,
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL2}),
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
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
									folderID2: path1,
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
			name: "DuplicatePreviousPaths_separateDrives",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							}),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							bupMD.DeltaURLsFileName,
							map[string]string{driveID2: deltaURL2}),
						graph.NewMetadataEntry(
							bupMD.PreviousPathFileName,
							map[string]map[string]string{
								driveID2: {
									folderID1: path1,
								},
							}),
					}
				},
			},
			expectedDeltas: map[string]string{
				driveID1: deltaURL1,
				driveID2: deltaURL2,
			},
			expectedPaths: map[string]map[string]string{
				driveID1: {
					folderID1: path1,
				},
				driveID2: {
					folderID1: path1,
				},
			},
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
					func(*support.ControllerOperationStatus) {})
				require.NoError(t, err, clues.ToCore(err))

				cols = append(cols, dataMock.NewUnversionedRestoreCollection(
					t,
					data.NoFetchRestoreCollection{Collection: mc}))
			}

			deltas, paths, canUsePreviousBackup, err := deserializeAndValidateMetadata(ctx, cols, fault.New(true))
			test.errCheck(t, err)
			assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup, "can use previous backup")

			assert.Equal(t, test.expectedDeltas, deltas, "deltas")
			assert.Equal(t, test.expectedPaths, paths, "paths")
		})
	}
}

type failingColl struct{}

func (f failingColl) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ic := make(chan data.Item)
	defer close(ic)

	errs.AddRecoverable(ctx, assert.AnError)

	return ic
}
func (f failingColl) FullPath() path.Path                                        { return nil }
func (f failingColl) FetchItemByName(context.Context, string) (data.Item, error) { return nil, nil }

// This check is to ensure that we don't error out, but still return
// canUsePreviousBackup as false on read errors
func (suite *OneDriveCollectionsUnitSuite) TestDeserializeMetadata_ReadFailure() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	fc := failingColl{}

	_, _, canUsePreviousBackup, err := deserializeAndValidateMetadata(ctx, []data.RestoreCollection{fc}, fault.New(true))
	require.NoError(t, err)
	require.False(t, canUsePreviousBackup)
}

func (suite *OneDriveCollectionsUnitSuite) TestGet() {
	var (
		tenant = "a-tenant"
		user   = "a-user"
		empty  = ""
		delta  = "delta1"
		delta2 = "delta2"
	)

	metadataPath, err := path.BuildMetadata(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false)
	require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))

	driveID1 := "drive-1-" + uuid.NewString()
	drive1 := models.NewDrive()
	drive1.SetId(&driveID1)
	drive1.SetName(&driveID1)

	driveID2 := "drive-2-" + uuid.NewString()
	drive2 := models.NewDrive()
	drive2.SetId(&driveID2)
	drive2.SetName(&driveID2)

	var (
		bh = userDriveBackupHandler{userID: user}

		driveBasePath1 = odConsts.DriveFolderPrefixBuilder(driveID1).String()
		driveBasePath2 = odConsts.DriveFolderPrefixBuilder(driveID2).String()

		expectedPath1 = getExpectedPathGenerator(suite.T(), bh, tenant, driveBasePath1)
		expectedPath2 = getExpectedPathGenerator(suite.T(), bh, tenant, driveBasePath2)

		rootFolderPath1 = expectedPath1("")
		folderPath1     = expectedPath1("/folder")

		rootFolderPath2 = expectedPath2("")
		folderPath2     = expectedPath2("/folder")
	)

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
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"), // will be present, not needed
								delItem("file", driveBasePath1, "root", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: makeExcludeMap("file"),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_NoFolderDeltas_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								driveItem("file", "file", driveBasePath1, "root", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {"file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: makeExcludeMap("file"),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
								driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors_FileRenamedMultiple",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
								driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								driveItem("file", "file2", driveBasePath1+"/folder", "folder", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths:        map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors_FileMovedMultiple",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							driveItem("file", "file2", driveBasePath1, "root", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {"file"}},
				folderPath1:     {data.NewState: {"folder"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: makeExcludeMap("file"),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_EmptyDelta_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: empty, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs: map[string]string{},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive_TwoItemPages_NoErrors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
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
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive_TwoItemPages_WithReset",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
									driveItem("file3", "file3", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
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
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive_TwoItemPages_WithResetCombinedWithItems",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
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
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
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
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
					driveID2: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root2"),
							driveItem("folder2", "folder", driveBasePath2, "root2", false, true, false),
							driveItem("file2", "file", driveBasePath2+"/folder", "folder2", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
				driveID2: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
				rootFolderPath2: {data.NewState: {}},
				folderPath2:     {data.NewState: {"folder2", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
				driveID2: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
				driveID2: {
					"root2":   rootFolderPath2,
					"folder2": folderPath2,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
				rootFolderPath2: true,
				folderPath2:     true,
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
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
					driveID2: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath2, "root", false, true, false),
							driveItem("file2", "file", driveBasePath2+"/folder", "folder", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
				driveID2: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
				rootFolderPath2: {data.NewState: {}},
				folderPath2:     {data.NewState: {"folder", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
				driveID2: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
				driveID2: {
					"root":   rootFolderPath2,
					"folder": folderPath2,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
				rootFolderPath2: true,
				folderPath2:     true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_Errors",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages:       []mock.NextPage{{Items: []models.DriveItemable{}}},
						DeltaUpdate: pagers.DeltaUpdate{},
						Err:         assert.AnError,
					},
				},
			},
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
			previousPaths: map[string]map[string]string{
				driveID1: {},
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
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder2", "folder2", true, false, false),
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
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1:           {data.NewState: {}},
				expectedPath1("/folder"):  {data.DeletedState: {}},
				expectedPath1("/folder2"): {data.NewState: {"folder2", "file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder2": expectedPath1("/folder2"),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1:           true,
				folderPath1:               true,
				expectedPath1("/folder2"): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDeltaCombinedWithItems_DeleteNonExistentFolder",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder2", "folder2", true, false, false),
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
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1:           {data.NewState: {}},
				expectedPath1("/folder"):  {data.DeletedState: {}},
				expectedPath1("/folder2"): {data.NewState: {"folder2", "file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder2": expectedPath1("/folder2"),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1:           true,
				folderPath1:               true,
				expectedPath1("/folder2"): true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								// on the first page, if this is the total data, we'd expect both folder and folder2
								// since new previousPaths merge with the old previousPaths.
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder2", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder2", true, false, false),
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
									driveRootItem("root"),
									driveItem("folder2", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder2", true, false, false),
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
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				expectedPath1("/folder"): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {"folder2", "file"},
				},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder2": expectedPath1("/folder"),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{},
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder2", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder2", true, false, false),
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
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				expectedPath1("/folder"): {
					// Old folder path should be marked as deleted since it should compare
					// by ID.
					data.DeletedState: {},
					data.NewState:     {"folder2", "file"},
				},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder2": expectedPath1("/folder"),
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "OneDrive Two Item Pages with Malware",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
									malwareItem("malware", "malware", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
									malwareItem("malware2", "malware2", driveBasePath1+"/folder", "folder", true, false, false),
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
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
			expectedSkippedCount: 2,
		},
		{
			name:   "One Drive Deleted Folder In New Results With Invalid Delta",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
									driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
									driveItem("file2", "file2", driveBasePath1+"/folder2", "folder2", true, false, false),
								},
							},
							{
								Reset: true,
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
									delItem("folder2", driveBasePath1, "root", false, true, false),
									delItem("file2", driveBasePath1, "root", true, false, false),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder":  folderPath1,
					"folder2": expectedPath1("/folder2"),
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1:           {data.NewState: {}},
				folderPath1:               {data.NotMovedState: {"folder", "file"}},
				expectedPath1("/folder2"): {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1:           true,
				folderPath1:               true,
				expectedPath1("/folder2"): true,
			},
		},
		{
			name:   "One Drive Folder Delete After Invalid Delta",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								delItem("folder", driveBasePath1, "root", false, true, false),
							},
							Reset: true,
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "One Drive Item Delete After Invalid Delta",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("file", driveBasePath1, "root", true, false, false),
								},
								Reset: true,
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
			},
		},
		{
			name:   "One Drive Folder Made And Deleted",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("folder", driveBasePath1, "root", false, true, false),
									delItem("file", driveBasePath1, "root", true, false, false),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
			},
		},
		{
			name:   "One Drive Folder Created -> Deleted -> Created",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("folder", driveBasePath1, "root", false, true, false),
									delItem("file", driveBasePath1, "root", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder1", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file1", "file", driveBasePath1+"/folder", "folder1", true, false, false),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder1", "file1"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder1": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "One Drive Folder Deleted -> Created -> Deleted",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("folder", driveBasePath1, "root", false, true, false),
									delItem("file", driveBasePath1+"/folder", "root", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("folder", driveBasePath1, "root", false, true, false),
									delItem("file", driveBasePath1+"/folder", "root", true, false, false),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {}},
				folderPath1:     {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
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
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("folder", driveBasePath1, "root", false, true, false),
									delItem("file", driveBasePath1, "root", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder1", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file1", "file", driveBasePath1+"/folder", "folder1", true, false, false),
								},
							},
						},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.DeletedState: {}, data.NewState: {"folder1", "file1"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder1": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: false,
				folderPath1:     true,
			},
		},
		{
			name:   "One Drive Item Made And Deleted",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
									driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								},
							},
							{
								Items: []models.DriveItemable{
									driveRootItem("root"),
									delItem("file", driveBasePath1, "root", true, false, false),
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
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
				folderPath1:     true,
			},
		},
		{
			name:   "One Drive Random Folder Delete",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							delItem("folder", driveBasePath1, "root", false, true, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
			},
		},
		{
			name:   "One Drive Random Item Delete",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"),
							delItem("file", driveBasePath1, "root", true, false, false),
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta, Reset: true},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath1: true,
			},
		},
		{
			name:   "TwoPriorDrives_OneTombstoned",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{Items: []models.DriveItemable{
							driveRootItem("root"), // will be present
						}}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
				driveID2: {"root": rootFolderPath2},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {}},
				rootFolderPath2: {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{driveID1: delta},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath2: true,
			},
		},
		{
			name:   "duplicate previous paths in metadata",
			drives: []models.Driveable{drive1, drive2},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					// contains duplicates in previousPath
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
								driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
								driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
								driveItem("file2", "file2", driveBasePath1+"/folder2", "folder2", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta},
					},
					// does not contain duplicates
					driveID2: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								driveItem("folder", "folder", driveBasePath2, "root", false, true, false),
								driveItem("file", "file", driveBasePath2+"/folder", "folder", true, false, false),
								driveItem("folder2", "folder2", driveBasePath2, "root", false, true, false),
								driveItem("file2", "file2", driveBasePath2+"/folder2", "folder2", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta2},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder":  rootFolderPath1 + "/folder",
					"folder2": rootFolderPath1 + "/folder",
					"folder3": rootFolderPath1 + "/folder",
				},
				driveID2: {
					"root":    rootFolderPath2,
					"folder":  rootFolderPath2 + "/folder",
					"folder2": rootFolderPath2 + "/folder2",
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {
					data.NewState: {"folder", "folder2"},
				},
				rootFolderPath1 + "/folder": {
					data.NotMovedState: {"folder", "file"},
				},
				rootFolderPath1 + "/folder2": {
					data.MovedState: {"folder2", "file2"},
				},
				rootFolderPath2: {
					data.NewState: {"folder", "folder2"},
				},
				rootFolderPath2 + "/folder": {
					data.NotMovedState: {"folder", "file"},
				},
				rootFolderPath2 + "/folder2": {
					data.NotMovedState: {"folder2", "file2"},
				},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
				driveID2: delta2,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":    rootFolderPath1,
					"folder":  rootFolderPath1 + "/folder2", // note: this is a bug, but is currently expected
					"folder2": rootFolderPath1 + "/folder2",
					"folder3": rootFolderPath1 + "/folder2",
				},
				driveID2: {
					"root":    rootFolderPath2,
					"folder":  rootFolderPath2 + "/folder",
					"folder2": rootFolderPath2 + "/folder2",
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: makeExcludeMap("file", "file2"),
				rootFolderPath2: makeExcludeMap("file", "file2"),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name:   "out of order item enumeration causes prev path collisions",
			drives: []models.Driveable{drive1},
			enumerator: mock.EnumerateItemsDeltaByDrive{
				DrivePagers: map[string]*mock.DriveItemsDeltaPager{
					driveID1: {
						Pages: []mock.NextPage{{
							Items: []models.DriveItemable{
								driveRootItem("root"),
								driveItem("fanny2", "fanny", driveBasePath1, "root", false, true, false),
								driveItem("file2", "file2", driveBasePath1+"/fanny", "fanny2", true, false, false),
								driveItem("nav", "nav", driveBasePath1, "root", false, true, false),
								driveItem("file", "file", driveBasePath1+"/nav", "nav", true, false, false),
							},
						}},
						DeltaUpdate: pagers.DeltaUpdate{URL: delta},
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			previousPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
					"nav":  rootFolderPath1 + "/fanny",
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {
					data.NewState: {"fanny2"},
				},
				rootFolderPath1 + "/nav": {
					data.MovedState: {"nav", "file"},
				},
				rootFolderPath1 + "/fanny": {
					data.NewState: {"fanny2", "file2"},
				},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedPreviousPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"nav":    rootFolderPath1 + "/nav",
					"fanny2": rootFolderPath1 + "/nav", // note: this is a bug, but currently expected
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: makeExcludeMap("file", "file2"),
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

			mbh := mock.DefaultOneDriveBH("a-user")
			mbh.DrivePagerV = mockDrivePager
			mbh.DriveItemEnumeration = test.enumerator

			c := NewCollections(
				mbh,
				tenant,
				idname.NewProvider(user, user),
				func(*support.ControllerOperationStatus) {},
				control.Options{ToggleFeatures: control.Toggles{}})

			prevDelta := "prev-delta"

			pathPrefix, err := mbh.MetadataPathPrefix(tenant)
			require.NoError(t, err, clues.ToCore(err))

			mc, err := graph.MakeMetadataCollection(
				pathPrefix,
				[]graph.MetadataCollectionEntry{
					graph.NewMetadataEntry(
						bupMD.DeltaURLsFileName,
						map[string]string{
							driveID1: prevDelta,
							driveID2: prevDelta,
						}),
					graph.NewMetadataEntry(
						bupMD.PreviousPathFileName,
						test.previousPaths),
				},
				func(*support.ControllerOperationStatus) {})
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

func coreItem(
	id string,
	name string,
	parentPath string,
	parentID string,
	isFile, isFolder, isPackage bool,
) *models.DriveItem {
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&id)

	parentReference := models.NewItemReference()
	parentReference.SetPath(&parentPath)
	parentReference.SetId(&parentID)
	item.SetParentReference(parentReference)

	switch {
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
	id string,
	name string,
	parentPath string,
	parentID string,
	isFile, isFolder, isPackage bool,
) models.DriveItemable {
	return coreItem(id, name, parentPath, parentID, isFile, isFolder, isPackage)
}

func fileItem(
	id, name, parentPath, parentID, url string,
	deleted bool,
) models.DriveItemable {
	di := driveItem(id, name, parentPath, parentID, true, false, false)
	di.SetAdditionalData(map[string]any{
		"@microsoft.graph.downloadUrl": url,
	})

	if deleted {
		di.SetDeleted(models.NewDeleted())
	}

	return di
}

func malwareItem(
	id string,
	name string,
	parentPath string,
	parentID string,
	isFile, isFolder, isPackage bool,
) models.DriveItemable {
	c := coreItem(id, name, parentPath, parentID, isFile, isFolder, isPackage)

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
	isFile, isFolder, isPackage bool,
) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetId(&id)
	item.SetDeleted(models.NewDeleted())

	parentReference := models.NewItemReference()
	parentReference.SetId(&parentID)
	item.SetParentReference(parentReference)

	switch {
	case isFile:
		item.SetFile(models.NewFile())
	case isFolder:
		item.SetFolder(models.NewFolder())
	case isPackage:
		item.SetPackageEscaped(models.NewPackageEscaped())
	}

	return item
}

func (suite *OneDriveCollectionsUnitSuite) TestAddURLCacheToDriveCollections() {
	driveID := "test-drive"
	collCount := 3
	anyFolder := (&selectors.OneDriveBackup{}).Folders(selectors.Any())[0]

	table := []struct {
		name             string
		items            []apiMock.PagerResult[any]
		deltaURL         string
		prevDeltaSuccess bool
		prevDelta        string
		err              error
	}{
		{
			name: "cache is attached",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			itemPagers := map[string]pagers.DeltaHandler[models.DriveItemable]{}
			itemPagers[driveID] = &apiMock.DeltaPager[models.DriveItemable]{}

			mbh := mock.DefaultOneDriveBH("test-user")
			mbh.ItemPagerV = itemPagers

			c := NewCollections(
				mbh,
				"test-tenant",
				idname.NewProvider("test-user", "test-user"),
				nil,
				control.Options{ToggleFeatures: control.Toggles{}})

			if _, ok := c.CollectionMap[driveID]; !ok {
				c.CollectionMap[driveID] = map[string]*Collection{}
			}

			// Add a few collections
			for i := 0; i < collCount; i++ {
				coll, err := NewCollection(
					&userDriveBackupHandler{
						baseUserDriveHandler: baseUserDriveHandler{
							ac: api.Drives{},
						},
						userID: "test-user",
						scope:  anyFolder,
					},
					idname.NewProvider("", ""),
					nil,
					nil,
					driveID,
					nil,
					control.Options{ToggleFeatures: control.Toggles{}},
					false,
					true,
					nil)
				require.NoError(t, err, clues.ToCore(err))

				c.CollectionMap[driveID][strconv.Itoa(i)] = coll
				require.Equal(t, nil, coll.urlCache, "cache not nil")
			}

			err := c.addURLCacheToDriveCollections(
				ctx,
				driveID,
				"",
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			// Check that all collections have the same cache instance attached
			// to them
			var uc *urlCache
			for _, driveColls := range c.CollectionMap {
				for _, coll := range driveColls {
					require.NotNil(t, coll.urlCache, "cache is nil")
					if uc == nil {
						uc = coll.urlCache.(*urlCache)
					} else {
						require.Equal(t, uc, coll.urlCache, "cache not equal")
					}
				}
			}
		})
	}
}
