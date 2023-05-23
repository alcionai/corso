package onedrive

import (
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type URLCacheIntegrationSuite struct {
	tester.Suite
	service graph.Servicer
	user    string
	driveID string
}

func TestURLCacheIntegrationSuite(t *testing.T) {
	suite.Run(t, &URLCacheIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *URLCacheIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	suite.service = loadTestService(t)
	suite.user = tester.SecondaryM365UserID(t)

	pager, err := PagerForSource(OneDriveSource, suite.service, suite.user, nil)
	require.NoError(t, err, clues.ToCore(err))

	odDrives, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	require.NoError(t, err, clues.ToCore(err))
	require.Greaterf(t, len(odDrives), 0, "user %s does not have a drive", suite.user)
	suite.driveID = ptr.Val(odDrives[0].GetId())
}

// Basic test for urlCache. Create some files in onedrive, then access them via
// url cache
func (suite *URLCacheIntegrationSuite) TestURLCacheBasic() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	srv := suite.service
	driveID := suite.driveID

	// Create a new test folder
	root, err := srv.Client().Drives().ByDriveId(driveID).Root().Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	newFolderName := tester.DefaultTestRestoreDestination("folder").ContainerName

	newFolder, err := CreateItem(
		ctx,
		srv,
		driveID,
		ptr.Val(root.GetId()),
		newItem(newFolderName, true))
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, newFolder.GetId())

	// Delete folder on exit
	defer func() {
		ictx := clues.Add(ctx, "folder_id", ptr.Val(newFolder.GetId()))

		err := DeleteItem(ictx, loadTestService(t), driveID, ptr.Val(newFolder.GetId()))
		if err != nil {
			logger.CtxErr(ictx, err).Errorw("deleting folder")
		}
	}()

	// Create a bunch of files in the new folder
	var items []models.DriveItemable

	for i := 0; i < 10; i++ {
		newItemName := "testItem_" + dttm.FormatNow(dttm.SafeForTesting)

		item, err := CreateItem(
			ctx,
			srv,
			driveID,
			ptr.Val(newFolder.GetId()),
			newItem(newItemName, false))
		if err != nil {
			// Something bad happened, skip this item
			return
		}

		items = append(items, item)
	}

	// Create a new URL cache
	cache := newURLCache(
		suite.driveID,
		1*time.Hour,
		collectDriveItems,
		defaultItemPager)

	// Launch parallel requests to the cache, one per item
	var wg sync.WaitGroup
	for i := 0; i < len(items); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Read item from URL cache
			itemProps, err := cache.getItemProperties(
				ctx,
				srv,
				ptr.Val(items[i].GetId()))

			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, itemProps)
			require.NotEmpty(t, itemProps.downloadURL)
			require.Equal(t, false, itemProps.isDeleted)

			// Validate download URL
			c := graph.NewNoTimeoutHTTPWrapper()

			resp, err := c.Request(
				ctx,
				http.MethodGet,
				itemProps.downloadURL,
				nil,
				nil)

			require.NoError(t, err, clues.ToCore(err))
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}(i)
	}
	wg.Wait()

	// Validate that <= 1 delta queries were made
	require.LessOrEqual(t, cache.deltaQueryCount, 1)
}
