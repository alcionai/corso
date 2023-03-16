package onedrive

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/graph"
	gapi "github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type statePath struct {
	state    data.CollectionState
	curPath  path.Path
	prevPath path.Path
}

func getExpectedStatePathGenerator(
	t *testing.T,
	tenant, user, base string,
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
			p2, err = GetCanonicalPath(base+pths[1], tenant, user, OneDriveSource)
			require.NoError(t, err, clues.ToCore(err))
		}

		p1, err = GetCanonicalPath(base+pths[0], tenant, user, OneDriveSource)
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

func getExpectedPathGenerator(t *testing.T,
	tenant, user, base string,
) func(string) string {
	return func(path string) string {
		p, err := GetCanonicalPath(base+path, tenant, user, OneDriveSource)
		require.NoError(t, err, clues.ToCore(err))

		return p.String()
	}
}

type OneDriveCollectionsUnitSuite struct {
	tester.Suite
}

func TestOneDriveCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, &OneDriveCollectionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveCollectionsUnitSuite) TestGetCanonicalPath() {
	tenant, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		source    driveSource
		dir       []string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "onedrive",
			source:    OneDriveSource,
			dir:       []string{"onedrive"},
			expect:    "tenant/onedrive/resourceOwner/files/onedrive",
			expectErr: assert.NoError,
		},
		{
			name:      "sharepoint",
			source:    SharePointSource,
			dir:       []string{"sharepoint"},
			expect:    "tenant/sharepoint/resourceOwner/libraries/sharepoint",
			expectErr: assert.NoError,
		},
		{
			name:      "unknown",
			source:    unknownDriveSource,
			dir:       []string{"unknown"},
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			p := strings.Join(test.dir, "/")

			result, err := GetCanonicalPath(p, tenant, resourceOwner, test.source)
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func getDelList(files ...string) map[string]struct{} {
	delList := map[string]struct{}{}
	for _, file := range files {
		delList[file+DataFileSuffix] = struct{}{}
		delList[file+MetaFileSuffix] = struct{}{}
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

	testBaseDrivePath := fmt.Sprintf(rootDrivePattern, "driveID1")
	expectedPath := getExpectedPathGenerator(suite.T(), tenant, user, testBaseDrivePath)
	expectedStatePath := getExpectedStatePathGenerator(suite.T(), tenant, user, testBaseDrivePath)

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
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t               = suite.T()
				excludes        = map[string]struct{}{}
				outputFolderMap = map[string]string{}
				itemCollection  = map[string]map[string]string{
					driveID: {},
				}
				errs = fault.New(true)
			)

			maps.Copy(outputFolderMap, tt.inputFolderMap)

			c := NewCollections(
				graph.HTTPClient(graph.NoTimeout()),
				tenant,
				user,
				OneDriveSource,
				testFolderMatcher{tt.scope},
				&MockGraphService{},
				nil,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}})

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
		cols           []func() []graph.MetadataCollectionEntry
		expectedDeltas map[string]string
		expectedPaths  map[string]map[string]string
		errCheck       assert.ErrorAssertionFunc
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
			errCheck: assert.NoError,
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
			expectedDeltas: map[string]string{},
			expectedPaths:  map[string]map[string]string{},
			errCheck:       assert.NoError,
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
			expectedPaths:  map[string]map[string]string{},
			errCheck:       assert.NoError,
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
			expectedDeltas: map[string]string{driveID1: deltaURL1},
			expectedPaths:  map[string]map[string]string{driveID1: {}},
			errCheck:       assert.NoError,
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
			expectedDeltas: map[string]string{},
			expectedPaths:  map[string]map[string]string{},
			errCheck:       assert.NoError,
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
			errCheck: assert.NoError,
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
			expectedDeltas: map[string]string{},
			expectedPaths:  map[string]map[string]string{},
			errCheck:       assert.Error,
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
			errCheck: assert.NoError,
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
			expectedDeltas: nil,
			expectedPaths:  nil,
			errCheck:       assert.Error,
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
			expectedDeltas: nil,
			expectedPaths:  nil,
			errCheck:       assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			cols := []data.RestoreCollection{}

			for _, c := range test.cols {
				mc, err := graph.MakeMetadataCollection(
					tenant,
					user,
					path.OneDriveService,
					path.FilesCategory,
					c(),
					func(*support.ConnectorOperationStatus) {})
				require.NoError(t, err, clues.ToCore(err))

				cols = append(cols, data.NotFoundRestoreCollection{Collection: mc})
			}

			deltas, paths, err := deserializeMetadata(ctx, cols, fault.New(true))
			test.errCheck(t, err)

			assert.Equal(t, test.expectedDeltas, deltas)
			assert.Equal(t, test.expectedPaths, paths)
		})
	}
}

type mockDeltaPageLinker struct {
	link  *string
	delta *string
}

func (pl *mockDeltaPageLinker) GetOdataNextLink() *string {
	return pl.link
}

func (pl *mockDeltaPageLinker) GetOdataDeltaLink() *string {
	return pl.delta
}

type deltaPagerResult struct {
	items     []models.DriveItemable
	nextLink  *string
	deltaLink *string
	err       error
}

type mockItemPager struct {
	// DriveID -> set of return values for queries for that drive.
	toReturn []deltaPagerResult
	getIdx   int
}

func (p *mockItemPager) GetPage(context.Context) (gapi.DeltaPageLinker, error) {
	if len(p.toReturn) <= p.getIdx {
		return nil, assert.AnError
	}

	idx := p.getIdx
	p.getIdx++

	return &mockDeltaPageLinker{
		p.toReturn[idx].nextLink,
		p.toReturn[idx].deltaLink,
	}, p.toReturn[idx].err
}

func (p *mockItemPager) SetNext(string) {}
func (p *mockItemPager) Reset()         {}

func (p *mockItemPager) ValuesIn(gapi.DeltaPageLinker) ([]models.DriveItemable, error) {
	idx := p.getIdx
	if idx > 0 {
		// Return values lag by one since we increment in GetPage().
		idx--
	}

	if len(p.toReturn) <= idx {
		return nil, assert.AnError
	}

	return p.toReturn[idx].items, nil
}

func (suite *OneDriveCollectionsUnitSuite) TestGet() {
	var (
		anyFolder = (&selectors.OneDriveBackup{}).Folders(selectors.Any())[0]
		tenant    = "a-tenant"
		user      = "a-user"
		empty     = ""
		next      = "next"
		delta     = "delta1"
		delta2    = "delta2"
	)

	metadataPath, err := path.Builder{}.ToServiceCategoryMetadataPath(
		tenant,
		user,
		path.OneDriveService,
		path.FilesCategory,
		false,
	)
	require.NoError(suite.T(), err, "making metadata path", clues.ToCore(err))

	driveID1 := uuid.NewString()
	drive1 := models.NewDrive()
	drive1.SetId(&driveID1)
	drive1.SetName(&driveID1)

	driveID2 := uuid.NewString()
	drive2 := models.NewDrive()
	drive2.SetId(&driveID2)
	drive2.SetName(&driveID2)

	var (
		driveBasePath1 = fmt.Sprintf(rootDrivePattern, driveID1)
		driveBasePath2 = fmt.Sprintf(rootDrivePattern, driveID2)

		expectedPath1 = getExpectedPathGenerator(suite.T(), tenant, user, driveBasePath1)
		expectedPath2 = getExpectedPathGenerator(suite.T(), tenant, user, driveBasePath2)

		rootFolderPath1 = expectedPath1("")
		folderPath1     = expectedPath1("/folder")

		rootFolderPath2 = expectedPath2("")
		folderPath2     = expectedPath2("/folder")
	)

	table := []struct {
		name            string
		drives          []models.Driveable
		items           map[string][]deltaPagerResult
		errCheck        assert.ErrorAssertionFunc
		prevFolderPaths map[string]map[string]string
		// Collection name -> set of item IDs. We can't check item data because
		// that's not mocked out. Metadata is checked separately.
		expectedCollections  map[string]map[data.CollectionState][]string
		expectedDeltaURLs    map[string]string
		expectedFolderPaths  map[string]map[string]string
		expectedDelList      map[string]map[string]struct{}
		expectedSkippedCount int
		doNotMergeItems      bool
	}{
		{
			name:   "OneDrive_OneItemPage_DelFileOnly_NoFolders_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"), // will be present, not needed
							delItem("file", driveBasePath1, "root", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoFolders_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors_FileRenamedMultiple",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							driveItem("file", "file2", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			},
		},
		{
			name:   "OneDrive_OneItemPage_NoErrors_FileMovedMultiple",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							driveItem("file", "file2", driveBasePath1, "root", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			},
		},
		{
			name:   "OneDrive_OneItemPage_EmptyDelta_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &empty, // probably will never happen with graph
					},
				},
			},
			errCheck: assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1: {data.NewState: {}},
				folderPath1:     {data.NewState: {"folder", "file"}},
			},
			expectedDeltaURLs:   map[string]string{},
			expectedFolderPaths: map[string]map[string]string{},
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
			},
		},
		{
			name:   "OneDrive_TwoItemPages_NoErrors",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						nextLink: &next,
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file", "file2"),
			},
		},
		{
			name: "TwoDrives_OneItemPageEach_NoErrors",
			drives: []models.Driveable{
				drive1,
				drive2,
			},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
				driveID2: {
					{
						items: []models.DriveItemable{
							driveRootItem("root2"),
							driveItem("folder2", "folder", driveBasePath2, "root2", false, true, false),
							driveItem("file2", "file", driveBasePath2+"/folder", "folder2", true, false, false),
						},
						deltaLink: &delta2,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file"),
				rootFolderPath2: getDelList("file2"),
			},
		},
		{
			name:   "OneDrive_OneItemPage_Errors",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						err: assert.AnError,
					},
				},
			},
			errCheck: assert.Error,
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
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						err: getDeltaError(),
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{},
			doNotMergeItems: true,
		},
		{
			name:   "OneDrive_TwoItemPage_DeltaError",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						err: getDeltaError(),
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						nextLink: &next,
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{},
			doNotMergeItems: true,
		},
		{
			name:   "OneDrive_TwoItemPage_NoDeltaError",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("file", "file", driveBasePath1, "root", true, false, false),
						},
						nextLink: &next,
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {},
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file", "file2"),
			},
			doNotMergeItems: false,
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_DeleteNonExistantFolder",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						err: getDeltaError(),
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder2", "folder2", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder2", "folder2", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{},
			doNotMergeItems: true,
		},
		{
			name:   "OneDrive_OneItemPage_InvalidPrevDelta_AnotherFolderAtDeletedLocation",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						err: getDeltaError(),
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder2", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder2", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
			prevFolderPaths: map[string]map[string]string{
				driveID1: {
					"root":   rootFolderPath1,
					"folder": folderPath1,
				},
			},
			expectedCollections: map[string]map[data.CollectionState][]string{
				rootFolderPath1:          {data.NewState: {}},
				expectedPath1("/folder"): {data.NewState: {"folder2", "file"}},
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
			expectedDelList: map[string]map[string]struct{}{},
			doNotMergeItems: true,
		},
		{
			name:   "OneDrive Two Item Pages with Malware",
			drives: []models.Driveable{drive1},
			items: map[string][]deltaPagerResult{
				driveID1: {
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file", "file", driveBasePath1+"/folder", "folder", true, false, false),
							malwareItem("malware", "malware", driveBasePath1+"/folder", "folder", true, false, false),
						},
						nextLink: &next,
					},
					{
						items: []models.DriveItemable{
							driveRootItem("root"),
							driveItem("folder", "folder", driveBasePath1, "root", false, true, false),
							driveItem("file2", "file2", driveBasePath1+"/folder", "folder", true, false, false),
							malwareItem("malware2", "malware2", driveBasePath1+"/folder", "folder", true, false, false),
						},
						deltaLink: &delta,
					},
				},
			},
			errCheck: assert.NoError,
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
			expectedDelList: map[string]map[string]struct{}{
				rootFolderPath1: getDelList("file", "file2"),
			},
			expectedSkippedCount: 2,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			drivePagerFunc := func(
				source driveSource,
				servicer graph.Servicer,
				resourceOwner string,
				fields []string,
			) (drivePager, error) {
				return &mockDrivePager{
					toReturn: []pagerResult{
						{
							drives: test.drives,
						},
					},
				}, nil
			}

			itemPagerFunc := func(
				servicer graph.Servicer,
				driveID, link string,
			) itemPager {
				return &mockItemPager{
					toReturn: test.items[driveID],
				}
			}

			c := NewCollections(
				graph.HTTPClient(graph.NoTimeout()),
				tenant,
				user,
				OneDriveSource,
				testFolderMatcher{anyFolder},
				&MockGraphService{},
				func(*support.ConnectorOperationStatus) {},
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}},
			)
			c.drivePagerFunc = drivePagerFunc
			c.itemPagerFunc = itemPagerFunc

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
						},
					),
					graph.NewMetadataEntry(
						graph.PreviousPathFileName,
						test.prevFolderPaths,
					),
				},
				func(*support.ConnectorOperationStatus) {},
			)
			assert.NoError(t, err, "creating metadata collection", clues.ToCore(err))

			prevMetadata := []data.RestoreCollection{data.NotFoundRestoreCollection{Collection: mc}}
			errs := fault.New(true)

			cols, delList, err := c.Get(ctx, prevMetadata, errs)
			test.errCheck(t, err)
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
					deltas, paths, err := deserializeMetadata(
						ctx,
						[]data.RestoreCollection{
							data.NotFoundRestoreCollection{Collection: baseCol},
						},
						fault.New(true))
					if !assert.NoError(t, err, "deserializing metadata", clues.ToCore(err)) {
						continue
					}

					assert.Equal(t, test.expectedDeltaURLs, deltas, "delta urls")
					assert.Equal(t, test.expectedFolderPaths, paths, "folder  paths")

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
				assert.Equal(t, test.doNotMergeItems, baseCol.DoNotMergeItems(), "DoNotMergeItems")
			}

			expectedCollectionCount := 0
			for c := range test.expectedCollections {
				for range test.expectedCollections[c] {
					expectedCollectionCount++
				}
			}

			// This check is necessary to make sure we are all the
			// collections we expect it to
			assert.Equal(t, expectedCollectionCount, collectionCount, "number of collections")

			assert.Equal(t, test.expectedDelList, delList, "del list")
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
		item.SetPackage(models.NewPackage_escaped())
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
		item.SetPackage(models.NewPackage_escaped())
	}

	return item
}

func getDeltaError() error {
	syncStateNotFound := "SyncStateNotFound" // TODO(meain): export graph.errCodeSyncStateNotFound
	me := odataerrors.NewMainError()
	me.SetCode(&syncStateNotFound)

	deltaError := odataerrors.NewODataError()
	deltaError.SetError(me)

	return deltaError
}

func (suite *OneDriveCollectionsUnitSuite) TestCollectItems() {
	next := "next"
	delta := "delta"
	prevDelta := "prev-delta"

	table := []struct {
		name             string
		items            []deltaPagerResult
		deltaURL         string
		prevDeltaSuccess bool
		prevDelta        string
		err              error
	}{
		{
			name:     "delta on first run",
			deltaURL: delta,
			items: []deltaPagerResult{
				{deltaLink: &delta},
			},
			prevDeltaSuccess: true,
			prevDelta:        prevDelta,
		},
		{
			name:     "empty prev delta",
			deltaURL: delta,
			items: []deltaPagerResult{
				{deltaLink: &delta},
			},
			prevDeltaSuccess: false,
			prevDelta:        "",
		},
		{
			name:     "next then delta",
			deltaURL: delta,
			items: []deltaPagerResult{
				{nextLink: &next},
				{deltaLink: &delta},
			},
			prevDeltaSuccess: true,
			prevDelta:        prevDelta,
		},
		{
			name:     "invalid prev delta",
			deltaURL: delta,
			items: []deltaPagerResult{
				{err: getDeltaError()},
				{deltaLink: &delta}, // works on retry
			},
			prevDelta:        prevDelta,
			prevDeltaSuccess: false,
		},
		{
			name: "fail a normal delta query",
			items: []deltaPagerResult{
				{nextLink: &next},
				{err: assert.AnError},
			},
			prevDelta:        prevDelta,
			prevDeltaSuccess: true,
			err:              assert.AnError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			itemPager := &mockItemPager{
				toReturn: test.items,
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
