package sharepoint_test

import (
	"context"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// consts, mocks
// ---------------------------------------------------------------------------

const (
	testBaseDrivePath = "drive/driveID1/root:"
)

type testFolderMatcher struct {
	scope selectors.SharePointScope
}

func (fm testFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.SharePointFolder)
}

func (fm testFolderMatcher) Matches(path string) bool {
	return fm.scope.Matches(selectors.SharePointFolder, path)
}

type MockGraphService struct{}

func (ms *MockGraphService) Client() *msgraphsdk.GraphServiceClient {
	return nil
}

func (ms *MockGraphService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return nil
}

func (ms *MockGraphService) ErrPolicy() bool {
	return false
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type SharePointLibrariesIntegrationSuite struct {
	suite.Suite
	ctx context.Context
}

func TestSharePointLibrariesIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorSharePointTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(SharePointLibrariesIntegrationSuite))
}

func (suite *SharePointLibrariesIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs,
	)

	require.NoError(suite.T(), err)
}

func (suite *SharePointLibrariesIntegrationSuite) TestUpdateCollections() {
	anyFolder := (&selectors.SharePointBackup{}).Folders(selectors.Any(), selectors.Any())[0]

	const (
		tenant    = "tenant"
		site      = "site"
		folder    = "/folder"
		folderSub = "/folder/subfolder"
		pkg       = "/package"
	)

	tests := []struct {
		testCase                string
		items                   []models.DriveItemable
		scope                   selectors.SharePointScope
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
				site,
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
				site,
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
			scope:  (&selectors.SharePointBackup{}).Folders(selectors.Any(), []string{"folder"})[0],
			expect: assert.NoError,
			expectedCollectionPaths: append(
				expectedPathAsSlice(
					suite.T(),
					tenant,
					site,
					testBaseDrivePath+"/folder",
				),
				expectedPathAsSlice(
					suite.T(),
					tenant,
					site,
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
			scope: (&selectors.SharePointBackup{}).
				Folders(selectors.Any(), []string{"/folder/subfolder"}, selectors.PrefixMatch())[0],
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				site,
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
			scope:  (&selectors.SharePointBackup{}).Folders(selectors.Any(), []string{"folder/subfolder"})[0],
			expect: assert.NoError,
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				site,
				testBaseDrivePath+folderSub,
			),
			expectedItemCount:      2,
			expectedFileCount:      1,
			expectedContainerCount: 1,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.testCase, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			c := onedrive.NewCollections(
				tenant,
				site,
				onedrive.SharePointSource,
				testFolderMatcher{test.scope},
				&MockGraphService{},
				nil)
			err := c.UpdateCollections(ctx, "driveID", test.items)
			test.expect(t, err)
			assert.Equal(t, len(test.expectedCollectionPaths), len(c.CollectionMap), "collection paths")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			for _, collPath := range test.expectedCollectionPaths {
				assert.Contains(t, c.CollectionMap, collPath)
			}
			t.Fail()
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

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func expectedPathAsSlice(t *testing.T, tenant, user string, rest ...string) []string {
	res := make([]string, 0, len(rest))

	for _, r := range rest {
		p, err := onedrive.GetCanonicalPath(r, tenant, user, onedrive.SharePointSource)
		require.NoError(t, err)

		res = append(res, p.String())
	}

	return res
}
