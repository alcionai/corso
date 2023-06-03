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
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type URLCacheIntegrationSuite struct {
	tester.Suite
	ac      api.Client
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

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.user = tester.SecondaryM365UserID(t)

	acct := tester.NewM365Account(t)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	drive, err := suite.ac.Users().GetDefaultDrive(ctx, suite.user)
	require.NoError(t, err, clues.ToCore(err))

	suite.driveID = ptr.Val(drive.GetId())
}

// Basic test for urlCache. Create some files in onedrive, then access them via
// url cache
func (suite *URLCacheIntegrationSuite) TestURLCacheBasic() {
	var (
		t              = suite.T()
		ac             = suite.ac.Drives()
		driveID        = suite.driveID
		newFolderName  = tester.DefaultTestRestoreDestination("folder").ContainerName
		driveItemPager = suite.ac.Drives().NewItemPager(driveID, "", api.DriveItemSelectDefault())
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// Create a new test folder
	root, err := ac.GetRootFolder(ctx, driveID)
	require.NoError(t, err, clues.ToCore(err))

	newFolder, err := ac.Drives().PostItemInContainer(
		ctx,
		driveID,
		ptr.Val(root.GetId()),
		newItem(newFolderName, true))
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, newFolder.GetId())

	nfid := ptr.Val(newFolder.GetId())

	// Create a bunch of files in the new folder
	var items []models.DriveItemable

	for i := 0; i < 10; i++ {
		newItemName := "test_url_cache_basic_" + dttm.FormatNow(dttm.SafeForTesting)

		item, err := ac.Drives().PostItemInContainer(
			ctx,
			driveID,
			nfid,
			newItem(newItemName, false))
		if err != nil {
			// Something bad happened, skip this item
			continue
		}

		items = append(items, item)
	}

	// Create a new URL cache with a long TTL
	cache, err := newURLCache(
		suite.driveID,
		1*time.Hour,
		driveItemPager,
		fault.New(true))

	require.NoError(t, err, clues.ToCore(err))

	err = cache.refreshCache(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// Launch parallel requests to the cache, one per item
	var wg sync.WaitGroup
	for i := 0; i < len(items); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Read item from URL cache
			props, err := cache.getItemProperties(
				ctx,
				ptr.Val(items[i].GetId()))

			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, props)
			require.NotEmpty(t, props.downloadURL)
			require.Equal(t, false, props.isDeleted)

			// Validate download URL
			c := graph.NewNoTimeoutHTTPWrapper()

			resp, err := c.Request(
				ctx,
				http.MethodGet,
				props.downloadURL,
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
