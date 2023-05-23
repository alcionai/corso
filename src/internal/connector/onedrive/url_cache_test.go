package onedrive

import (
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// Unit tests for urlCache

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

// Basic test for urlCache
func (suite *URLCacheIntegrationSuite) TestURLCacheBasic() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	srv := suite.service
	driveID := suite.driveID

	// Create a new URL cache
	cache := newURLCache(
		suite.driveID,
		1*time.Hour,
		collectDriveItems,
		defaultItemPager)

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

	// Launch parallel requests to create onedrive files. Read the URLs from
	// the cache and confirm whether they are valid.
	// Verify that only one delta query is made.
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Sleep for some msec, range 0-100
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Create a test file in test folder
			newItemName := "testItem_" + dttm.FormatNow(dttm.SafeForTesting)

			newItem, err := CreateItem(
				ctx,
				srv,
				driveID,
				ptr.Val(newFolder.GetId()),
				newItem(newItemName, false))
			if err != nil {
				// Something bad happened, skip this item
				return
			}

			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, newItem.GetId())
			// Read item from URL cache
			url, err := cache.getDownloadURL(ctx, srv, ptr.Val(newItem.GetId()))
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, url)

			// Validate URL
			// TODO: use downloadItem call once it's refactored to use URLs
			// instead of DriveItemable
			client := graph.NewNoTimeoutHTTPWrapper()
			resp, err := client.Request(ctx, http.MethodGet, url, nil, nil)
			require.NoError(t, err, clues.ToCore(err))
			require.Equal(t, http.StatusOK, resp.StatusCode)
		}()
	}
	wg.Wait()

	// TODO: This can get flaky in presence of graph transient errors
	// See how we handle it in other integration tests.
	require.Equal(t, 1, cache.deltaQueryCount)
}

/*
Unit test list

1. Test updateCache - against []DI returned by mock delta queries - itempager
	- Duplicate DIs. Latest DI should prevail
	- Deleted DIs
	- DIs with no download URL
	- DIs with download URL
	- DIs with download URL and deleted
	- delta query failures
2. Test readCache
	- cache miss - return error
	- cache hit - return URL
	- cache hit, deleted - return deleted error
3. Test needRefresh
	- cache is empty
	- cache is not empty, but refresh interval has passed
	- none of the above
4. Test refreshCache - mock out delta query
	- Semaphore:
		- Validate that only one thread can concurrently refresh the cache
		- nil semaphore - should not panic
	- If cache is already refreshed by another thread, return

5. Test updateRefreshTime
	- Validate that refresh time is updated
	- Error case - new refresh time can never be < old refresh time

6. collectDriveItems
	- See collectItems tests

7. Concurrency tests
	- Stale cache read: Edge cases during refresh interval expiry.
		Readers holding read lock, refresh should block until read lock is released.
		Cache may serve stale cache hits for readers at this time.
		Client should fallback to item GET on eventual 401
	- RW lock tests
		- Multiple concurrent readers, single writer(cache refresher)
		- Multiple potential concurrent writers, multiple readers

*/

/*
Integration test list - use gock to simulate 401s on refresh interval expiry

1. Test getDownloadURL
2. Test refreshCache
3. Test downloadContent - Cache failures should not be treated as fatal error by client
	- Confirm fallback to item GET
	- Deleted item is a tricky one. Need to see how to handle it.
	- use mock cache to simulate cache failures

*/
