package onedrive

import (
	"strings"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	testBaseDrivePath = "drive/driveID1/root:"
)

func expectedPathAsSlice(t *testing.T, tenant, user string, rest ...string) []string {
	res := make([]string, 0, len(rest))

	for _, r := range rest {
		p, err := GetCanonicalPath(r, tenant, user, OneDriveSource)
		require.NoError(t, err)

		res = append(res, p.String())
	}

	return res
}

type OneDriveCollectionsSuite struct {
	suite.Suite
}

func TestOneDriveCollectionsSuite(t *testing.T) {
	suite.Run(t, new(OneDriveCollectionsSuite))
}

func (suite *OneDriveCollectionsSuite) TestGetCanonicalPath() {
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
		suite.T().Run(test.name, func(t *testing.T) {
			p := strings.Join(test.dir, "/")
			result, err := GetCanonicalPath(p, tenant, resourceOwner, test.source)
			test.expectErr(t, err)
			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *OneDriveCollectionsSuite) TestUpdateCollections() {
	anyFolder := (&selectors.OneDriveBackup{}).Folders(selectors.Any())[0]

	const (
		tenant    = "tenant"
		user      = "user"
		folder    = "/folder"
		folderSub = "/folder/subfolder"
		pkg       = "/package"
	)

	tests := []struct {
		testCase                string
		items                   []models.DriveItemable
		inputFolderMap          map[string]string
		scope                   selectors.OneDriveScope
		expect                  assert.ErrorAssertionFunc
		expectedCollectionPaths []string
		expectedItemCount       int
		expectedContainerCount  int
		expectedFileCount       int
		expectedMetadataPaths   map[string]string
		expectedExcludes        map[string]struct{}
	}{
		{
			testCase: "Invalid item",
			items: []models.DriveItemable{
				driveItem("item", "item", testBaseDrivePath, false, false, false),
			},
			inputFolderMap:        map[string]string{},
			scope:                 anyFolder,
			expect:                assert.Error,
			expectedMetadataPaths: map[string]string{},
			expectedExcludes:      map[string]struct{}{},
		},
		{
			testCase: "Single File",
			items: []models.DriveItemable{
				driveItem("file", "file", testBaseDrivePath, true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// Root folder is skipped since it's always present.
			expectedMetadataPaths: map[string]string{},
			expectedExcludes:      map[string]struct{}{},
		},
		{
			testCase: "Single Folder",
			items: []models.DriveItemable{
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedMetadataPaths: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
			},
			expectedItemCount:      1,
			expectedContainerCount: 1,
			expectedExcludes:       map[string]struct{}{},
		},
		{
			testCase: "Single Package",
			items: []models.DriveItemable{
				driveItem("package", "package", testBaseDrivePath, false, false, true),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedMetadataPaths: map[string]string{
				"package": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/package",
				)[0],
			},
			expectedItemCount:      1,
			expectedContainerCount: 1,
			expectedExcludes:       map[string]struct{}{},
		},
		{
			testCase: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
				driveItem("package", "package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          anyFolder,
			expect:         assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
				testBaseDrivePath+folder,
				testBaseDrivePath+pkg,
			),
			expectedItemCount:      5,
			expectedFileCount:      3,
			expectedContainerCount: 3,
			expectedMetadataPaths: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"package": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/package",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "contains folder selector",
			items: []models.DriveItemable{
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, false, true, false),
				driveItem("folder2", "folder", testBaseDrivePath+folderSub, false, true, false),
				driveItem("package", "package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInFolder2", "fileInFolder2", testBaseDrivePath+folderSub+folder, true, false, false),
				driveItem("fileInFolderPackage", "fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          (&selectors.OneDriveBackup{}).Folders([]string{"folder"})[0],
			expect:         assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath+"/folder",
				testBaseDrivePath+folderSub,
				testBaseDrivePath+folderSub+folder,
			),
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 3,
			// just "folder" isn't added here because the include check is done on the
			// parent path since we only check later if something is a folder or not.
			expectedMetadataPaths: map[string]string{
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder",
				)[0],
				"folder2": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder/folder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "prefix subfolder selector",
			items: []models.DriveItemable{
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, false, true, false),
				driveItem("folder2", "folder", testBaseDrivePath+folderSub, false, true, false),
				driveItem("package", "package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInFolder2", "fileInFolder2", testBaseDrivePath+folderSub+folder, true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope: (&selectors.OneDriveBackup{}).
				Folders([]string{"/folder/subfolder"}, selectors.PrefixMatch())[0],
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath+folderSub,
				testBaseDrivePath+folderSub+folder,
			),
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 2,
			expectedMetadataPaths: map[string]string{
				"folder2": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder/folder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "match subfolder selector",
			items: []models.DriveItemable{
				driveItem("fileInRoot", "fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath+folder, false, true, false),
				driveItem("package", "package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", "fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInSubfolder", "fileInSubfolder", testBaseDrivePath+folderSub, true, false, false),
				driveItem("fileInPackage", "fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			inputFolderMap: map[string]string{},
			scope:          (&selectors.OneDriveBackup{}).Folders([]string{"folder/subfolder"})[0],
			expect:         assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath+folderSub,
			),
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
			// No child folders for subfolder so nothing here.
			expectedMetadataPaths: map[string]string{},
			expectedExcludes:      map[string]struct{}{},
		},
		{
			testCase: "not moved folder tree",
			items: []models.DriveItemable{
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder",
				)[0],
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "moved folder tree",
			items: []models.DriveItemable{
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/a-folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/a-folder/subfolder",
				)[0],
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "moved folder tree and subfolder 1",
			items: []models.DriveItemable{
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath, false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/a-folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/a-folder/subfolder",
				)[0],
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/subfolder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "moved folder tree and subfolder 2",
			items: []models.DriveItemable{
				driveItem("subfolder", "subfolder", testBaseDrivePath, false, true, false),
				driveItem("folder", "folder", testBaseDrivePath, false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/a-folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/a-folder/subfolder",
				)[0],
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedItemCount:      2,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/subfolder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "deleted folder and package",
			items: []models.DriveItemable{
				delItem("folder", testBaseDrivePath, false, true, false),
				delItem("package", testBaseDrivePath, false, false, true),
			},
			inputFolderMap: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"package": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/package",
				)[0],
			},
			scope:                   anyFolder,
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{},
			expectedItemCount:       0,
			expectedFileCount:       0,
			expectedContainerCount:  0,
			expectedMetadataPaths:   map[string]string{},
			expectedExcludes:        map[string]struct{}{},
		},
		{
			testCase: "delete folder tree move subfolder",
			items: []models.DriveItemable{
				delItem("folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", "subfolder", testBaseDrivePath, false, true, false),
			},
			inputFolderMap: map[string]string{
				"folder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				)[0],
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder/subfolder",
				)[0],
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
			),
			expectedItemCount:      1,
			expectedFileCount:      0,
			expectedContainerCount: 1,
			expectedMetadataPaths: map[string]string{
				"subfolder": expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/subfolder",
				)[0],
			},
			expectedExcludes: map[string]struct{}{},
		},
		{
			testCase: "delete file",
			items: []models.DriveItemable{
				delItem("item", testBaseDrivePath, true, false, false),
			},
			inputFolderMap:          map[string]string{},
			scope:                   anyFolder,
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{},
			expectedItemCount:       1,
			expectedFileCount:       1,
			expectedContainerCount:  0,
			expectedMetadataPaths:   map[string]string{},
			expectedExcludes: map[string]struct{}{
				"item": {},
			},
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.testCase, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			excludes := map[string]struct{}{}
			outputFolderMap := map[string]string{}
			maps.Copy(outputFolderMap, tt.inputFolderMap)
			c := NewCollections(
				graph.HTTPClient(graph.NoTimeout()),
				tenant,
				user,
				OneDriveSource,
				testFolderMatcher{tt.scope},
				&MockGraphService{},
				nil,
				control.Options{})

			err := c.UpdateCollections(
				ctx,
				"driveID",
				"General",
				tt.items,
				tt.inputFolderMap,
				outputFolderMap,
				excludes,
			)
			tt.expect(t, err)
			assert.Equal(t, len(tt.expectedCollectionPaths), len(c.CollectionMap), "collection paths")
			assert.Equal(t, tt.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, tt.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, tt.expectedContainerCount, c.NumContainers, "container count")
			for _, collPath := range tt.expectedCollectionPaths {
				assert.Contains(t, c.CollectionMap, collPath)
			}

			assert.Equal(t, tt.expectedMetadataPaths, outputFolderMap)
			assert.Equal(t, tt.expectedExcludes, excludes)
		})
	}
}

func driveItem(id string, name string, path string, isFile, isFolder, isPackage bool) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&id)

	parentReference := models.NewItemReference()
	parentReference.SetPath(&path)
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

// delItem creates a DriveItemable that is marked as deleted. path must be set
// to the base drive path.
func delItem(id string, path string, isFile, isFolder, isPackage bool) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetId(&id)
	item.SetDeleted(models.NewDeleted())

	parentReference := models.NewItemReference()
	parentReference.SetPath(&path)
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
