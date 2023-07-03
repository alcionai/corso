package sharepoint

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// consts
// ---------------------------------------------------------------------------

var testBaseDrivePath = path.Builder{}.Append(
	odConsts.DrivesPathDir,
	"driveID1",
	odConsts.RootPathDir)

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type LibrariesBackupUnitSuite struct {
	tester.Suite
}

func TestLibrariesBackupUnitSuite(t *testing.T) {
	suite.Run(t, &LibrariesBackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *LibrariesBackupUnitSuite) TestUpdateCollections() {
	anyFolder := (&selectors.SharePointBackup{}).LibraryFolders(selectors.Any())[0]

	const (
		tenantID = "tenant"
		site     = "site"
		driveID  = "driveID1"
	)

	pb := path.Builder{}.Append(testBaseDrivePath.Elements()...)
	ep, err := libraryBackupHandler{}.CanonicalPath(pb, tenantID, site)
	require.NoError(suite.T(), err, clues.ToCore(err))

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
				driveRootItem(odConsts.RootID),
				driveItem("file", testBaseDrivePath.String(), odConsts.RootID, true),
			},
			scope:                   anyFolder,
			expect:                  assert.NoError,
			expectedCollectionIDs:   []string{odConsts.RootID},
			expectedCollectionPaths: []string{ep.String()},
			expectedItemCount:       1,
			expectedFileCount:       1,
			expectedContainerCount:  1,
		},
	}

	for _, test := range tests {
		suite.Run(test.testCase, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				paths     = map[string]string{}
				newPaths  = map[string]string{}
				excluded  = map[string]struct{}{}
				itemColls = map[string]map[string]string{
					driveID: {},
				}
				collMap = map[string]map[string]*onedrive.Collection{
					driveID: {},
				}
			)

			c := onedrive.NewCollections(
				&libraryBackupHandler{api.Drives{}, test.scope},
				tenantID,
				site,
				nil,
				control.Defaults())

			c.CollectionMap = collMap

			err := c.UpdateCollections(
				ctx,
				driveID,
				"General",
				test.items,
				paths,
				newPaths,
				excluded,
				itemColls,
				true,
				fault.New(true))

			test.expect(t, err, clues.ToCore(err))
			assert.Equal(t, len(test.expectedCollectionIDs), len(c.CollectionMap), "collection paths")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")

			for _, collPath := range test.expectedCollectionIDs {
				assert.Contains(t, c.CollectionMap[driveID], collPath)
			}

			for _, col := range c.CollectionMap[driveID] {
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
	item.SetFolder(models.NewFolder())

	return item
}

type SharePointPagesSuite struct {
	tester.Suite
}

func TestSharePointPagesSuite(t *testing.T) {
	suite.Run(t, &SharePointPagesSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *SharePointPagesSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, false, 4)
}

func (suite *SharePointPagesSuite) TestCollectPages() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID = tester.M365SiteID(t)
		a      = tester.NewM365Account(t)
	)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	col, err := collectPages(
		ctx,
		creds,
		ac,
		mock.NewProvider(siteID, siteID),
		&MockGraphService{},
		control.Defaults(),
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, col)
}
