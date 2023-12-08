package sharepoint

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
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
		siteID   = "site"
		driveID  = "driveID1"
	)

	pb := path.Builder{}.Append(testBaseDrivePath.Elements()...)
	ep, err := drive.NewSiteBackupHandler(api.Drives{}, siteID, nil, path.SharePointService).
		CanonicalPath(pb, tenantID)
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
				mbh      = mock.DefaultSharePointBH(siteID)
				paths    = map[string]string{}
				excluded = map[string]struct{}{}
				collMap  = map[string]map[string]*drive.Collection{
					driveID: {},
				}
				topLevelPackages = map[string]struct{}{}
			)

			mbh.DriveItemEnumeration = mock.DriveEnumerator(
				mock.Drive(driveID).With(
					mock.Delta("notempty", nil).With(mock.NextPage{Items: test.items})))

			c := drive.NewCollections(
				mbh,
				tenantID,
				idname.NewProvider(siteID, siteID),
				nil,
				control.DefaultOptions(),
				count.New())

			c.CollectionMap = collMap

			_, _, err := c.PopulateDriveCollections(
				ctx,
				driveID,
				"General",
				paths,
				excluded,
				topLevelPackages,
				"notempty",
				count.New(),
				fault.New(true))

			test.expect(t, err, clues.ToCore(err))
			assert.Equal(t, len(test.expectedCollectionIDs), len(c.CollectionMap), "collection paths")
			assert.Equal(t, test.expectedItemCount, c.NumItems, "item count")
			assert.Equal(t, test.expectedFileCount, c.NumFiles, "file count")
			assert.Equal(t, test.expectedContainerCount, c.NumContainers, "container count")
			assert.Empty(t, topLevelPackages, "should not find package type folders")

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
