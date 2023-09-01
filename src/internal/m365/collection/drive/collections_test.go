package drive

import (
	"context"
	"strconv"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	pmMock "github.com/alcionai/corso/src/internal/common/prefixmatcher/mock"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	apiMock "github.com/alcionai/corso/src/pkg/services/m365/api/mock"
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

func getDelList(files ...string) map[string]struct{} {
	delList := map[string]struct{}{}
	for _, file := range files {
		delList[file+metadata.DataFileSuffix] = struct{}{}
		delList[file+metadata.MetaFileSuffix] = struct{}{}
	}

	return delList
}

func (suite *OneDriveCollectionsUnitSuite) TestUpdateCollections() {
	anyFolder := (&selectors.OneDriveBackup{}).Folders(selectors.Any())[0]

	const (
		driveID   = "driveID1"
		tenant    = "tenant"
		user      = "user"
		folder    = "/folder"
		folderSub = "/folder/subfolder"
		pkg       = "/package"
	)

	bh := itemBackupHandler{userID: user}
	testBaseDrivePath := odConsts.DriveFolderPrefixBuilder("driveID1").String()
	expectedPath := getExpectedPathGenerator(suite.T(), bh, tenant, testBaseDrivePath)
	expectedStatePath := getExpectedStatePathGenerator(suite.T(), bh, tenant, testBaseDrivePath)

	tests := []struct {
		testCase               string
		items                  []models.DriveItemable
		inputFolderMap         map[string]string
		scope                  selectors.OneDriveScope
		expect                 assert.ErrorAssertionFunc
		expectedCollectionIDs  map[string]statePath
		expectedItemCount      int
		expectedContainerCount int
		expectedFileCount      int
		expectedSkippedCount   int
		expectedMetadataPaths  map[string]string
		expectedExcludes       map[string]struct{}
	}{
		{
			testCase: "Invalid item",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("item", "item", testBaseDrivePath, "root", false, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.Error,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "Single File",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("file", "file", testBaseDrivePath, "root", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// Root folder is skipped since it's always present.
			expectedMetadataPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes: getDelList("file"),
		},
		{
			testCase: "Single Folder",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NewState, folder),
			},
			expectedMetadataPaths: map[string]string{
				"root":   expectedPath(""),
				"folder": expectedPath("/folder"),
			},
			expectedItemCount:      1,
			expectedContainerCount: 2,
			expectedExcludes:       map[string]struct{}{},
		},
		{
			testCase: "Single Package",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"package": expectedStatePath(data.NewState, pkg),
			},
			expectedMetadataPaths: map[string]string{
				"root":    expectedPath(""),
				"package": expectedPath("/package"),
			},
			expectedItemCount:      1,
			expectedContainerCount: 2,
			expectedExcludes:       map[string]struct{}{},
		},
		{
			testCase: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"folder":  expectedStatePath(data.NewState, folder),
				"package": expectedStatePath(data.NewState, pkg),
			},
			expectedItemCount:      5,
			expectedFileCount:      3,
			expectedContainerCount: 3,
			expectedMetadataPaths: map[string]string{
				"root":    expectedPath(""),
				"folder":  expectedPath("/folder"),
				"package": expectedPath("/package"),
			},
			expectedExcludes: getDelList("fileInRoot", "fileInFolder", "fileInPackage"),
		},
		{
			testCase: "contains folder selector",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, "folder", false, true, false),
				driveItem("folder2", "folder", testBaseDrivePath+folderSub, "subfolder", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem("fileInFolder2", "fileInFolder2", testBaseDrivePath+folderSub+folder, "folder2", true, false, false),
				driveItem("fileInFolderPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          (&selectors.OneDriveBackup{}).Folders([]string{"folder"})[0],
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"folder":    expectedStatePath(data.NewState, folder),
				"subfolder": expectedStatePath(data.NewState, folderSub),
				"folder2":   expectedStatePath(data.NewState, folderSub+folder),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			// just "folder" isn't added here because the include check is done on the
			// parent path since we only check later if something is a folder or not.
			expectedMetadataPaths: map[string]string{
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folderSub),
				"folder2":   expectedPath(folderSub + folder),
			},
			expectedExcludes: getDelList("fileInFolder", "fileInFolder2"),
		},
		{
			testCase: "prefix subfolder selector",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, "folder", false, true, false),
				driveItem("folder2", "folder", testBaseDrivePath+folderSub, "subfolder", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem("fileInFolder2", "fileInFolder2", testBaseDrivePath+folderSub+folder, "folder2", true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope: (&selectors.OneDriveBackup{}).
				Folders([]string{"/folder/subfolder"}, selectors.PrefixMatch())[0],
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"subfolder": expectedStatePath(data.NewState, folderSub),
				"folder2":   expectedStatePath(data.NewState, folderSub+folder),
			},
			expectedItemCount:      3,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"subfolder": expectedPath(folderSub),
				"folder2":   expectedPath(folderSub + folder),
			},
			expectedExcludes: getDelList("fileInFolder2"),
		},
		{
			testCase: "match subfolder selector",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, "folder", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, "folder", true, false, false),
				driveItem("fileInSubfolder", "fileInSubfolder", testBaseDrivePath+folderSub, "subfolder", true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, "package", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          (&selectors.OneDriveBackup{}).Folders([]string{"folder/subfolder"})[0],
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"subfolder": expectedStatePath(data.NewState, folderSub),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// No child folders for subfolder so nothing here.
			expectedMetadataPaths: map[string]string{
				"subfolder": expectedPath(folderSub),
			},
			expectedExcludes: getDelList("fileInSubfolder"),
		},
		{
			testCase: "not moved folder tree",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folderSub),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NotMovedState, folder),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folderSub),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "moved folder tree",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.MovedState, folder, "/a-folder"),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath(folderSub),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "moved folder tree with file no previous",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
				driveItem("folder", "folder2", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NewState, "/folder2"),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"root":   expectedPath(""),
				"folder": expectedPath("/folder2"),
			},
			expectedExcludes: getDelList("file"),
		},
		{
			testCase: "moved folder tree with file no previous 1",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.NewState, folder),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"root":   expectedPath(""),
				"folder": expectedPath(folder),
			},
			expectedExcludes: getDelList("file"),
		},
		{
			testCase: "moved folder tree and subfolder 1",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.MovedState, folder, "/a-folder"),
				"subfolder": expectedStatePath(data.MovedState, "/subfolder", "/a-folder/subfolder"),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath("/subfolder"),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "moved folder tree and subfolder 2",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("subfolder", "subfolder", testBaseDrivePath, "root", false, true, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.MovedState, folder, "/a-folder"),
				"subfolder": expectedStatePath(data.MovedState, "/subfolder", "/a-folder/subfolder"),
			},
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 3,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath(folder),
				"subfolder": expectedPath("/subfolder"),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "move subfolder when moving parent",
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
					false,
				),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder":    expectedPath("/a-folder"),
				"subfolder": expectedPath("/a-folder/subfolder"),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.MovedState, folder, "/a-folder"),
				"folder2":   expectedStatePath(data.NewState, "/folder2"),
				"subfolder": expectedStatePath(data.MovedState, folderSub, "/a-folder/subfolder"),
			},
			expectedItemCount:      5,
			expectedFileCount:      2,
			expectedContainerCount: 4,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath("/folder"),
				"folder2":   expectedPath("/folder2"),
				"subfolder": expectedPath("/folder/subfolder"),
			},
			expectedExcludes: getDelList("itemInSubfolder", "itemInFolder2"),
		},
		{
			testCase: "moved folder tree multiple times",
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
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":   expectedStatePath(data.NotMovedState, ""),
				"folder": expectedStatePath(data.MovedState, "/folder2", "/a-folder"),
			},
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"folder":    expectedPath("/folder2"),
				"subfolder": expectedPath("/folder2/subfolder"),
			},
			expectedExcludes: getDelList("file"),
		},
		{
			testCase: "deleted folder and package",
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
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"folder":  expectedStatePath(data.DeletedState, folder),
				"package": expectedStatePath(data.DeletedState, pkg),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "delete folder without previous",
			items: []models.DriveItemable{
				driveRootItem("root"),
				delItem("folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{
				"root": expectedPath(""),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "delete folder tree move subfolder",
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
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":      expectedStatePath(data.NotMovedState, ""),
				"folder":    expectedStatePath(data.DeletedState, folder),
				"subfolder": expectedStatePath(data.MovedState, "/subfolder", folderSub),
			},
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"root":      expectedPath(""),
				"subfolder": expectedPath("/subfolder"),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "delete file",
			items: []models.DriveItemable{
				driveRootItem("root"),
				delItem("item", testBaseDrivePath, "root", true, false, false),
			},
			inputFolderMap: map[string]string{
				"root": expectedPath(""),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes: getDelList("item"),
		},
		{
			testCase: "item before parent errors",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("file", "file", testBaseDrivePath+"/folder", "folder", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.Error,
			expectedCollectionIDs: map[string]statePath{
				"root": expectedStatePath(data.NotMovedState, ""),
			},
			expectedItemCount:      0,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"root": expectedPath(""),
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "1 root file, 1 folder, 1 package, 1 good file, 1 malware",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, "root", true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, "root", false, true, false),
				driveItem("package", "package", testBaseDrivePath, "root", false, false, true),
				driveItem("goodFile", "goodFile", testBaseDrivePath+folder, "folder", true, false, false),
				malwareItem("malwareFile", "malwareFile", testBaseDrivePath+folder, "folder", true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionIDs: map[string]statePath{
				"root":    expectedStatePath(data.NotMovedState, ""),
				"folder":  expectedStatePath(data.NewState, folder),
				"package": expectedStatePath(data.NewState, pkg),
			},
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			expectedSkippedCount:   1,
			expectedMetadataPaths: map[string]string{
				"root":    expectedPath(""),
				"folder":  expectedPath("/folder"),
				"package": expectedPath("/package"),
			},
			expectedExcludes: getDelList("fileInRoot", "goodFile"),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.testCase, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				excludes        = map[string]struct{}{}
				outputFolderMap = map[string]string{}
				itemCollection  = map[string]map[string]string{
					driveID: {},
				}
				errs = fault.New(true)
			)

			maps.Copy(outputFolderMap, tt.inputFolderMap)

			c := NewCollections(
				&itemBackupHandler{api.Drives{}, user, tt.scope},
				tenant,
				user,
				nil,
				control.Options{ToggleFeatures: control.Toggles{}})

			c.CollectionMap[driveID] = map[string]*Collection{}

			err := c.UpdateCollections(
				ctx,
				driveID,
				"General",
				tt.items,
				tt.inputFolderMap,
				outputFolderMap,
				excludes,
				itemCollection,
				false,
				errs)
			tt.expect(t, err, clues.ToCore(err))
			assert.Equal(t, len(tt.expectedCollectionIDs), len(c.CollectionMap[driveID]), "total collections")
			assert.Equal(t, tt.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, tt.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, tt.expectedContainerCount, c.NumContainers, "container count")
			assert.Equal(t, tt.expectedSkippedCount, len(errs.Skipped()), "skipped items")

			for id, sp := range tt.expectedCollectionIDs {
				if !assert.Containsf(t, c.CollectionMap[driveID], id, "missing collection with id %s", id) {
					// Skip collections we don't find so we don't get an NPE.
					continue
				}

				assert.Equalf(t, sp.state, c.CollectionMap[driveID][id].State(), "state for collection %s", id)
				assert.Equalf(t, sp.curPath, c.CollectionMap[driveID][id].FullPath(), "current path for collection %s", id)
				assert.Equalf(t, sp.prevPath, c.CollectionMap[driveID][id].PreviousPath(), "prev path for collection %s", id)
			}

			assert.Equal(t, tt.expectedMetadataPaths, outputFolderMap, "metadata paths")
			assert.Equal(t, tt.expectedExcludes, excludes, "exclude list")
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
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
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
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
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
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
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
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {},
							},
						),
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
							graph.DeltaURLsFileName,
							map[string]string{
								driveID1: "",
							},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
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
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							graph.DeltaURLsFileName,
							map[string]string{driveID2: deltaURL2},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID2: {
									folderID2: path2,
								},
							},
						),
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
			name: "BadFormat",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]string{driveID1: deltaURL1},
						),
					}
				},
			},
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
		},
		{
			// Unexpected files are logged and skipped. They don't cause an error to
			// be returned.
			name: "BadFileName",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
						graph.NewMetadataEntry(
							"foo",
							map[string]string{driveID1: deltaURL1},
						),
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
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID2: path2,
								},
							},
						),
					}
				},
			},
			expectedDeltas:       nil,
			expectedPaths:        nil,
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
		},
		{
			name: "DriveAlreadyFound_Deltas",
			cols: []func() []graph.MetadataCollectionEntry{
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL1},
						),
						graph.NewMetadataEntry(
							graph.PreviousPathFileName,
							map[string]map[string]string{
								driveID1: {
									folderID1: path1,
								},
							},
						),
					}
				},
				func() []graph.MetadataCollectionEntry {
					return []graph.MetadataCollectionEntry{
						graph.NewMetadataEntry(
							graph.DeltaURLsFileName,
							map[string]string{driveID1: deltaURL2},
						),
					}
				},
			},
			expectedDeltas:       nil,
			expectedPaths:        nil,
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			cols := []data.RestoreCollection{}

			for _, c := range test.cols {
				mc, err := graph.MakeMetadataCollection(
					tenant,
					user,
					path.OneDriveService,
					path.FilesCategory,
					c(),
					func(*support.ControllerOperationStatus) {})
				require.NoError(t, err, clues.ToCore(err))

				cols = append(cols, data.NoFetchRestoreCollection{Collection: mc})
			}

			deltas, paths, canUsePreviousBackup, err := deserializeMetadata(ctx, cols)
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

	_, _, canUsePreviousBackup, err := deserializeMetadata(ctx, []data.RestoreCollection{fc})
	require.NoError(t, err)
	require.False(t, canUsePreviousBackup)
}

func (suite *OneDriveCollectionsUnitSuite) TestGet() {
	var (
		tenant = "a-tenant"
		user   = "a-user"
		empty  = ""
		next   = "next"
		delta  = "delta1"
		delta2 = "delta2"
	)

	metadataPath, err := path.Builder{}.ToServiceCategoryMetadataPath(
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
		bh = itemBackupHandler{userID: user}

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
		items                map[string][]apiMock.PagerResult[models.DriveItemable]
		canUsePreviousBackup bool
		errCheck             assert.ErrorAssertionFunc
		prevFolderPaths      map[string]map[string]string
		// Collection name -> set of item IDs. We can't check item data because
		// that's not mocked out. Metadata is checked separately.
		expectedCollections map[string]map[data.CollectionState][]string
		expectedDeltaURLs   map[string]string
		expectedFolderPaths map[string]map[string]string
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"), // will be present, not needed
							delItem("file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_NoFolderDeltas_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {"file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths:      map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							driveItem("file", "file2", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths:      map[string]map[string]string{},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							driveItem("file", "file2", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			}),
		},
		{
			name:   "OneDrive_OneItemPage_EmptyDelta_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &empty, // probably will never happen with graph
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs: map[string]string{},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
				driveID2: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root2"),
							driveItem("folder2", "folder", driveBasePath2, "root2", false, true, false),
							driveItem("file2", "file", driveBasePath2+"/folder", "folder2", true, false, false),
						},
						DeltaLink: &delta2,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
				driveID2: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath2, "root", false, true, false),
							driveItem("file2", "file", driveBasePath2+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta2,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: assert.AnError,
					},
				},
			},
			canUsePreviousBackup: false,
			errCheck:             assert.Error,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: nil,
			expectedDeltaURLs:   nil,
			expectedFolderPaths: nil,
			expectedDelList:     nil,
		},
		{
			name:   "OneDrive_OneItemPage_DeltaError",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {"file"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			name:   "OneDrive_TwoItemPage_DeltaError",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1:          {data.NotMovedState: {"file"}},
				expectedPath1("/folder"): {data.NewState: {"folder", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			name:   "OneDrive_TwoItemPage_NoDeltaError",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {
					"root": rootFolderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1:          {data.NotMovedState: {"file"}},
				expectedPath1("/folder"): {data.NewState: {"folder", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file", "file2"),
			}),
			doNotMergeItems: map[string]bool{},
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_DeleteNonExistentFolder",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder2", "folder2", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder2", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder2", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							malwareItem("malware", "malware", driveBasePath1+"/folder", "folder", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
							malwareItem("malware2", "malware2", driveBasePath1+"/folder", "folder", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file", "file2"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			name:   "One Drive Delta Error Deleted Folder In New Results",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file2", driveBasePath1+"/folder2", "folder2", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("folder2", driveBasePath1, "root", false, true, false),
							delItem("file2", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta2,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			name:   "One Drive Delta Error Random Folder Delete",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("folder", driveBasePath1, "root", false, true, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			name:   "One Drive Delta Error Random Item Delete",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Err: getDeltaError(),
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
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
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("folder", driveBasePath1, "root", false, true, false),
							delItem("file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta2,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta2,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			name:   "One Drive Item Made And Deleted",
			drives: []models.Driveable{drive1},
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						NextLink: &next,
					},
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder"}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("folder", driveBasePath1, "root", false, true, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"),
							delItem("file", driveBasePath1, "root", true, false, false),
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
			},
			expectedDeltaURLs: map[string]string{
				driveID1: delta,
			},
			expectedFolderPaths: map[string]map[string]string{
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
			items: map[string][]apiMock.PagerResult[models.DriveItemable]{
				driveID1: {
					{
						Values: []models.DriveItemable{
							driveRootItem("root"), // will be present
						},
						DeltaLink: &delta,
					},
				},
			},
			canUsePreviousBackup: true,
			errCheck:             assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
				driveID2: {"root": rootFolderPath2},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NotMovedState: {}},
				rootFolderPath2: {data.DeletedState: {}},
			},
			expectedDeltaURLs: map[string]string{driveID1: delta},
			expectedFolderPaths: map[string]map[string]string{
				driveID1: {"root": rootFolderPath1},
			},
			expectedDelList: pmMock.NewPrefixMap(map[string]map[string]struct{}{}),
			doNotMergeItems: map[string]bool{
				rootFolderPath2: true,
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

			itemPagers := map[string]api.DeltaPager[models.DriveItemable]{}

			for driveID := range test.items {
				itemPagers[driveID] = &apiMock.DeltaPager[models.DriveItemable]{
					ToReturn: test.items[driveID],
				}
			}

			mbh := mock.DefaultOneDriveBH("a-user")
			mbh.DrivePagerV = mockDrivePager
			mbh.ItemPagerV = itemPagers

			c := NewCollections(
				mbh,
				tenant,
				user,
				func(*support.ControllerOperationStatus) {},
				control.Options{ToggleFeatures: control.Toggles{}})

			prevDelta := "prev-delta"
			mc, err := graph.MakeMetadataCollection(
				tenant,
				user,
				path.OneDriveService,
				path.FilesCategory,
				[]graph.MetadataCollectionEntry{
					graph.NewMetadataEntry(
						graph.DeltaURLsFileName,
						map[string]string{
							driveID1: prevDelta,
							driveID2: prevDelta,
						}),
					graph.NewMetadataEntry(
						graph.PreviousPathFileName,
						test.prevFolderPaths),
				},
				func(*support.ControllerOperationStatus) {},
			)
			assert.NoError(t, err, "creating metadata collection", clues.ToCore(err))

			prevMetadata := []data.RestoreCollection{data.NoFetchRestoreCollection{Collection: mc}}
			errs := fault.New(true)

			delList := prefixmatcher.NewStringSetBuilder()

			cols, canUsePreviousBackup, err := c.Get(ctx, prevMetadata, delList, errs)
			test.errCheck(t, err)
			assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup, "can use previous backup")
			assert.Equal(t, test.expectedSkippedCount, len(errs.Skipped()))

			if err != nil {
				return
			}

			collectionCount := 0
			for _, baseCol := range cols {
				var folderPath string
				if baseCol.State() != data.DeletedState {
					folderPath = baseCol.FullPath().String()
				} else {
					folderPath = baseCol.PreviousPath().String()
				}

				if folderPath == metadataPath.String() {
					deltas, paths, _, err := deserializeMetadata(
						ctx,
						[]data.RestoreCollection{
							data.NoFetchRestoreCollection{Collection: baseCol},
						})
					if !assert.NoError(t, err, "deserializing metadata", clues.ToCore(err)) {
						continue
					}

					assert.Equal(t, test.expectedDeltaURLs, deltas, "delta urls")
					assert.Equal(t, test.expectedFolderPaths, paths, "folder paths")

					continue
				}

				collectionCount++

				// TODO(ashmrtn): We should really be getting items in the collection
				// via the Items() channel, but we don't have a way to mock out the
				// actual item fetch yet (mostly wiring issues). The lack of that makes
				// this check a bit more bittle since internal details can change.
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
					"state: %d, path: %s",
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

			expectedCollectionCount := 0
			for _, ec := range test.expectedCollections {
				expectedCollectionCount += len(ec)
			}

			assert.Equal(t, expectedCollectionCount, collectionCount, "number of collections")

			test.expectedDelList.AssertEqual(t, delList)
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

func getDeltaError() error {
	syncStateNotFound := "SyncStateNotFound" // TODO(meain): export graph.errCodeSyncStateNotFound
	me := odataerrors.NewMainError()
	me.SetCode(&syncStateNotFound)

	deltaError := odataerrors.NewODataError()
	deltaError.SetErrorEscaped(me)

	return deltaError
}

func (suite *OneDriveCollectionsUnitSuite) TestCollectItems() {
	next := "next"
	delta := "delta"
	prevDelta := "prev-delta"

	table := []struct {
		name             string
		items            []apiMock.PagerResult[models.DriveItemable]
		deltaURL         string
		prevDeltaSuccess bool
		prevDelta        string
		err              error
	}{
		{
			name:     "delta on first run",
			deltaURL: delta,
			items: []apiMock.PagerResult[models.DriveItemable]{
				{DeltaLink: &delta},
			},
			prevDeltaSuccess: true,
			prevDelta:        prevDelta,
		},
		{
			name:     "empty prev delta",
			deltaURL: delta,
			items: []apiMock.PagerResult[models.DriveItemable]{
				{DeltaLink: &delta},
			},
			prevDeltaSuccess: false,
			prevDelta:        "",
		},
		{
			name:     "next then delta",
			deltaURL: delta,
			items: []apiMock.PagerResult[models.DriveItemable]{
				{NextLink: &next},
				{DeltaLink: &delta},
			},
			prevDeltaSuccess: true,
			prevDelta:        prevDelta,
		},
		{
			name:     "invalid prev delta",
			deltaURL: delta,
			items: []apiMock.PagerResult[models.DriveItemable]{
				{Err: getDeltaError()},
				{DeltaLink: &delta}, // works on retry
			},
			prevDelta:        prevDelta,
			prevDeltaSuccess: false,
		},
		{
			name: "fail a normal delta query",
			items: []apiMock.PagerResult[models.DriveItemable]{
				{NextLink: &next},
				{Err: assert.AnError},
			},
			prevDelta:        prevDelta,
			prevDeltaSuccess: true,
			err:              assert.AnError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			itemPager := &apiMock.DeltaPager[models.DriveItemable]{
				ToReturn: test.items,
			}

			collectorFunc := func(
				ctx context.Context,
				driveID, driveName string,
				driveItems []models.DriveItemable,
				oldPaths map[string]string,
				newPaths map[string]string,
				excluded map[string]struct{},
				itemCollection map[string]map[string]string,
				doNotMergeItems bool,
				errs *fault.Bus,
			) error {
				return nil
			}

			delta, _, _, err := collectItems(
				ctx,
				itemPager,
				"",
				"General",
				collectorFunc,
				map[string]string{},
				test.prevDelta,
				fault.New(true))

			require.ErrorIs(t, err, test.err, "delta fetch err", clues.ToCore(err))
			require.Equal(t, test.deltaURL, delta.URL, "delta url")
			require.Equal(t, !test.prevDeltaSuccess, delta.Reset, "delta reset")
		})
	}
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

			itemPagers := map[string]api.DeltaPager[models.DriveItemable]{}
			itemPagers[driveID] = &apiMock.DeltaPager[models.DriveItemable]{}

			mbh := mock.DefaultOneDriveBH("test-user")
			mbh.ItemPagerV = itemPagers

			c := NewCollections(
				mbh,
				"test-tenant",
				"test-user",
				nil,
				control.Options{ToggleFeatures: control.Toggles{}})

			if _, ok := c.CollectionMap[driveID]; !ok {
				c.CollectionMap[driveID] = map[string]*Collection{}
			}

			// Add a few collections
			for i := 0; i < collCount; i++ {
				coll, err := NewCollection(
					&itemBackupHandler{api.Drives{}, "test-user", anyFolder},
					nil,
					nil,
					driveID,
					nil,
					control.Options{ToggleFeatures: control.Toggles{}},
					CollectionScopeFolder,
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
