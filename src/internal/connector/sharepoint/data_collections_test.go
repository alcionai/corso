package sharepoint

import (
	"testing"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// consts
// ---------------------------------------------------------------------------

const (
	testBaseDrivePath = "drive/driveID1/root:"
)

type testFolderMatcher struct {
	scope selectors.SharePointScope
}

func (fm testFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.SharePointLibrary)
}

func (fm testFolderMatcher) Matches(path string) bool {
	return fm.scope.Matches(selectors.SharePointLibrary, path)
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type SharePointLibrariesSuite struct {
	suite.Suite
	mockService mockServicer
}

type mockServicer struct{}

func (mock mockServicer) Client() *msgraphsdk.GraphServiceClient {
	return nil
}

func (mock mockServicer) Adapter() *msgraphsdk.GraphRequestAdapter {
	return nil
}

func (mock mockServicer) Serialize(object absser.Parsable) ([]byte, error) {
	return nil, nil
}

func TestSharePointLibrariesSuite(t *testing.T) {
	suite.Run(t, new(SharePointLibrariesSuite))
}

func (suite *SharePointLibrariesSuite) TestUpdateCollections() {
	anyFolder := (&selectors.SharePointBackup{}).Libraries(selectors.Any())[0]

	const (
		tenant = "tenant"
		site   = "site"
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
			testCase: "Single File",
			items: []models.DriveItemable{
				driveItem("file", testBaseDrivePath, true),
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
	}

	for _, test := range tests {
		suite.T().Run(test.testCase, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			paths := map[string]string{}
			c := onedrive.NewCollections(
				tenant,
				site,
				onedrive.SharePointSource,
				testFolderMatcher{test.scope},
				suite.mockService,
				nil,
				control.Options{})
			err := c.UpdateCollections(ctx, "driveID", test.items, paths)
			test.expect(t, err)
			assert.Equal(t, len(test.expectedCollectionPaths), len(c.CollectionMap), "collection paths")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			for _, collPath := range test.expectedCollectionPaths {
				assert.Contains(t, c.CollectionMap, collPath)
			}
		})
	}
}

func driveItem(name string, path string, isFile bool) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&name)

	parentReference := models.NewItemReference()
	parentReference.SetPath(&path)
	item.SetParentReference(parentReference)

	if isFile {
		item.SetFile(models.NewFile())
	}

	return item
}
