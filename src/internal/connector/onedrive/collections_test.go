package onedrive

import (
	"context"
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OneDriveCollectionsSuite struct {
	suite.Suite
}

func TestOneDriveCollectionsSuite(t *testing.T) {
	suite.Run(t, new(OneDriveCollectionsSuite))
}

func (suite *OneDriveCollectionsSuite) TestUpdateCollections() {
	tests := []struct {
		testCase                string
		items                   []models.DriveItemable
		expect                  assert.ErrorAssertionFunc
		expectedCollectionPaths []string
		expectedItemCount       int
		expectedFolderCount     int
		expectedPackageCount    int
		expectedFileCount       int
	}{
		{
			testCase: "Invalid item",
			items: []models.DriveItemable{
				driveItem("item", "/root", false, false, false),
			},
			expect: assert.Error,
		},
		{
			testCase: "Single File",
			items: []models.DriveItemable{
				driveItem("file", "/root", true, false, false),
			},
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{"/root"},
			expectedItemCount:       1,
			expectedFileCount:       1,
		},
		{
			testCase: "Single Folder",
			items: []models.DriveItemable{
				driveItem("folder", "/root", false, true, false),
			},
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{"/root", "/root/folder"},
			expectedItemCount:       1,
			expectedFolderCount:     1,
		},
		{
			testCase: "Single Package",
			items: []models.DriveItemable{
				driveItem("package", "/root", false, false, true),
			},
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{"/root", "/root/package"},
			expectedItemCount:       1,
			expectedPackageCount:    1,
		},
		{
			testCase: "1 root file, 1 folder, 1 package, 2 files, 3 collections",
			items: []models.DriveItemable{
				driveItem("fileInRoot", "/root", true, false, false),
				driveItem("folder", "/root", false, true, false),
				driveItem("package", "/root", false, false, true),
				driveItem("fileInFolder", "/root/folder", true, false, false),
				driveItem("fileInPackage", "/root/package", true, false, false),
			},
			expect:                  assert.NoError,
			expectedCollectionPaths: []string{"/root", "/root/folder", "/root/package"},
			expectedItemCount:       5,
			expectedFileCount:       3,
			expectedFolderCount:     1,
			expectedPackageCount:    1,
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.testCase, func(t *testing.T) {
			c := NewCollections("user", &MockGraphService{}, nil)
			err := c.updateCollections(context.Background(), "driveID", tt.items)
			tt.expect(t, err)
			assert.Equal(t, len(tt.expectedCollectionPaths), len(c.collectionMap))
			assert.Equal(t, tt.expectedItemCount, c.numItems)
			assert.Equal(t, tt.expectedFileCount, c.numFiles)
			assert.Equal(t, tt.expectedFolderCount, c.numDirs)
			assert.Equal(t, tt.expectedPackageCount, c.numPackages)
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
