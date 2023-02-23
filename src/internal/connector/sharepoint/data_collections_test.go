package sharepoint

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// consts
// ---------------------------------------------------------------------------

const (
	testBaseDrivePath = "drives/driveID1/root:"
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
	tester.Suite
}

func TestSharePointLibrariesSuite(t *testing.T) {
	suite.Run(t, &SharePointLibrariesSuite{Suite: tester.NewUnitSuite(t)})
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
		expectedCollectionIDs   []string
		expectedCollectionPaths []string
		expectedItemCount       int
		expectedContainerCount  int
		expectedFileCount       int
	}{
		{
			testCase: "Single File",
			items: []models.DriveItemable{
				driveRootItem("root"),
				driveItem("file", testBaseDrivePath, "root", true),
			},
			scope:                 anyFolder,
			expect:                assert.NoError,
			expectedCollectionIDs: []string{"root"},
			expectedCollectionPaths: expectedPathAsSlice(
				suite.T(),
				tenant,
				site,
				testBaseDrivePath,
			),
			expectedItemCount:      1,
			expectedFileCount:      1,
			expectedContainerCount: 1,
		},
	}

	for _, test := range tests {
		suite.Run(test.testCase, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			paths := map[string]string{}
			newPaths := map[string]string{}
			excluded := map[string]struct{}{}
			c := onedrive.NewCollections(
				graph.HTTPClient(graph.NoTimeout()),
				tenant,
				site,
				onedrive.SharePointSource,
				testFolderMatcher{test.scope},
				&MockGraphService{},
				nil,
				control.Options{})
			err := c.UpdateCollections(
				ctx,
				"driveID1",
				"General",
				test.items,
				paths,
				newPaths,
				excluded,
				map[string]string{},
				true,
			)
			test.expect(t, err)
			assert.Equal(t, len(test.expectedCollectionIDs), len(c.CollectionMap), "collection paths")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			for _, collPath := range test.expectedCollectionIDs {
				assert.Contains(t, c.CollectionMap, collPath)
			}
			for _, col := range c.CollectionMap {
				assert.Contains(t, test.expectedCollectionPaths, col.FullPath().String())
			}
		})
	}
}

func driveItem(name, parentPath, parentID string, isFile bool) models.DriveItemable {
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&name)

	parentReference := models.NewItemReference()
	parentReference.SetPath(&parentPath)
	parentReference.SetId(&parentID)
	item.SetParentReference(parentReference)

	if isFile {
		item.SetFile(models.NewFile())
	}

	return item
}

func driveRootItem(id string) models.DriveItemable {
	name := "root"
	item := models.NewDriveItem()
	item.SetName(&name)
	item.SetId(&id)
	item.SetRoot(models.NewRoot())

	return item
}

type SharePointPagesSuite struct {
	tester.Suite
}

func TestSharePointPagesSuite(t *testing.T) {
	suite.Run(t, &SharePointPagesSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
			tester.CorsoGraphConnectorTests,
			tester.CorsoGraphConnectorSharePointTests),
	})
}

func (suite *SharePointPagesSuite) TestCollectPages() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	siteID := tester.M365SiteID(t)
	a := tester.NewM365Account(t)
	account, err := a.M365Config()
	require.NoError(t, err)

	col, err := collectPages(
		ctx,
		account,
		nil,
		siteID,
		&MockGraphService{},
		control.Defaults(),
		fault.New(true))
	assert.NoError(t, err)
	assert.NotEmpty(t, col)
}
