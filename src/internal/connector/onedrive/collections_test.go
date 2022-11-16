package onedrive

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	testBaseDrivePath = "drive/driveID1/root:"
)

func expectedPathAsSlice(t *testing.T, tenant, user string, rest ...string) []string {
	res := make([]string, 0, len(rest))

	for _, r := range rest {
		p, err := getCanonicalPath(r, tenant, user, OneDriveSource)
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

func (suite *OneDriveCollectionsSuite) TestUpdateCollections() {
	anyFolder := (&selectors.OneDriveBackup{}).Folders(selectors.Any(), selectors.Any())[0]

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
		scope                   selectors.OneDriveScope
		expect                  assert.ErrorAssertionFunc
		expectedCollectionPaths []string
		expectedItemCount       int
		expectedContainerCount  int
		expectedFileCount       int
	}{
		{
			testCase: "Invalid item",
			items: []models.DriveItemable{
				driveItem("item", testBaseDrivePath, false, false, false),
			},
			scope:  anyFolder,
			expect: assert.Error,
		},
		{
			testCase: "Single File",
			items: []models.DriveItemable{
				driveItem("file", testBaseDrivePath, true, false, false),
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
			expectedFileCount:      1,
			expectedContainerCount: 1,
		},
		{
			testCase: "Single Folder",
			items: []models.DriveItemable{
				driveItem("folder", testBaseDrivePath, false, true, false),
			},
			scope:                   anyFolder,
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{},
		},
		{
			testCase: "Single Package",
			items: []models.DriveItemable{
				driveItem("package", testBaseDrivePath, false, false, true),
			},
			scope:                   anyFolder,
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{},
		},
		{
			testCase: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveItem("fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", testBaseDrivePath, false, true, false),
				driveItem("package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			scope:  anyFolder,
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath,
				testBaseDrivePath+folder,
				testBaseDrivePath+pkg,
			),
			expectedItemCount:      6,
			expectedFileCount:      3,
			expectedContainerCount: 3,
		},
		{
			testCase: "contains folder selector",
			items: []models.DriveItemable{
				driveItem("fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", testBaseDrivePath+folder, false, true, false),
				driveItem("folder", testBaseDrivePath+folderSub, false, true, false),
				driveItem("package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInFolder2", testBaseDrivePath+folderSub+folder, true, false, false),
				driveItem("fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			scope:  (&selectors.OneDriveBackup{}).Folders(selectors.Any(), []string{"folder"})[0],
			expect: assert.NoError,
			expectedCollectionPaths: append(
				expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+"/folder",
				),
				expectedPathAsSlice(
					suite.T(),
					tenant,
					user,
					testBaseDrivePath+folderSub+folder,
				)...,
			),
			expectedItemCount:      4,
			expectedFileCount:      2,
			expectedContainerCount: 2,
		},
		{
			testCase: "prefix subfolder selector",
			items: []models.DriveItemable{
				driveItem("fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", testBaseDrivePath+folder, false, true, false),
				driveItem("folder", testBaseDrivePath+folderSub, false, true, false),
				driveItem("package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInFolder2", testBaseDrivePath+folderSub+folder, true, false, false),
				driveItem("fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			scope: (&selectors.OneDriveBackup{}).
				Folders(selectors.Any(), []string{"/folder/subfolder"}, selectors.PrefixMatch())[0],
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath+folderSub+folder,
			),
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
		},
		{
			testCase: "match subfolder selector",
			items: []models.DriveItemable{
				driveItem("fileInRoot", testBaseDrivePath, true, false, false),
				driveItem("folder", testBaseDrivePath, false, true, false),
				driveItem("subfolder", testBaseDrivePath+folder, false, true, false),
				driveItem("package", testBaseDrivePath, false, false, true),
				driveItem("fileInFolder", testBaseDrivePath+folder, true, false, false),
				driveItem("fileInSubfolder", testBaseDrivePath+folderSub, true, false, false),
				driveItem("fileInPackage", testBaseDrivePath+pkg, true, false, false),
			},
			scope:  (&selectors.OneDriveBackup{}).Folders(selectors.Any(), []string{"folder/subfolder"})[0],
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				user,
				testBaseDrivePath+folderSub,
			),
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.testCase, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			c := NewCollections(tenant, user, OneDriveSource, testFolderMatcher{tt.scope}, &MockGraphService{}, nil)
			err := c.updateCollections(ctx, "driveID", tt.items)
			tt.expect(t, err)
			assert.Equal(t, len(tt.expectedCollectionPaths), len(c.collectionMap), "collection paths")
			assert.Equal(t, tt.expectedItemCount, c.numItems, "item count")
			assert.Equal(t, tt.expectedFileCount, c.numFiles, "file count")
			assert.Equal(t, tt.expectedContainerCount, c.numContainers, "container count")
			for _, collPath := range tt.expectedCollectionPaths {
				assert.Contains(t, c.collectionMap, collPath)
			}
		})
	}
}

func driveItem(name string, path string, isFile, isFolder, isPackage bool) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&name)

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
