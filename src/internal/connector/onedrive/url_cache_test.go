package onedrive

import (
	"errors"
	"math/rand"
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

type URLCacheUnitSuite struct {
	tester.Suite
}

func TestURLCacheUnitSuite(t *testing.T) {
	suite.Run(t, &URLCacheUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *URLCacheUnitSuite) TestGetItemProperties() {
	deltaString := "delta"
	next := "next"
	driveID := "drive1"

	table := []struct {
		name              string
		pagerResult       map[string][]deltaPagerResult
		expectedItemProps map[string]itemProps
		expectedErr       require.ErrorAssertionFunc
		cacheAssert       func(*urlCache, time.Time)
	}{
		{
			name: "single item in cache",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
						},
						deltaLink: &deltaString,
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"1": {
					downloadURL: "https://dummy1.com",
					isDeleted:   false,
				},
			},
			expectedErr: require.NoError,
			cacheAssert: func(uc *urlCache, startTime time.Time) {
				require.Greater(suite.T(), uc.lastRefreshTime, startTime)
				require.Equal(suite.T(), 1, uc.deltaQueryCount)
				require.Equal(suite.T(), 1, len(uc.idToProps))
			},
		},
		{
			name: "multiple items in cache",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
							fileItem("2", "file2", "root", "root", "https://dummy2.com", false),
							fileItem("3", "file3", "root", "root", "https://dummy3.com", false),
							fileItem("4", "file4", "root", "root", "https://dummy4.com", false),
							fileItem("5", "file5", "root", "root", "https://dummy5.com", false),
						},
						deltaLink: &deltaString,
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"1": {
					downloadURL: "https://dummy1.com",
					isDeleted:   false,
				},
				"2": {
					downloadURL: "https://dummy2.com",
					isDeleted:   false,
				},
				"3": {
					downloadURL: "https://dummy3.com",
					isDeleted:   false,
				},
				"4": {
					downloadURL: "https://dummy4.com",
					isDeleted:   false,
				},
				"5": {
					downloadURL: "https://dummy5.com",
					isDeleted:   false,
				},
			},
			expectedErr: require.NoError,
			cacheAssert: func(uc *urlCache, startTime time.Time) {
				require.Greater(suite.T(), uc.lastRefreshTime, startTime)
				require.Equal(suite.T(), 1, uc.deltaQueryCount)
				require.Equal(suite.T(), 5, len(uc.idToProps))
			},
		},
		{
			name: "duplicate items with potentially new urls",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
							fileItem("2", "file2", "root", "root", "https://dummy2.com", false),
							fileItem("3", "file3", "root", "root", "https://dummy3.com", false),
							fileItem("1", "file1", "root", "root", "https://test1.com", false),
							fileItem("2", "file2", "root", "root", "https://test2.com", false),
						},
						deltaLink: &deltaString,
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"1": {
					downloadURL: "https://test1.com",
					isDeleted:   false,
				},
				"2": {
					downloadURL: "https://test2.com",
					isDeleted:   false,
				},
				"3": {
					downloadURL: "https://dummy3.com",
					isDeleted:   false,
				},
			},
			expectedErr: require.NoError,
			cacheAssert: func(uc *urlCache, startTime time.Time) {
				require.Greater(suite.T(), uc.lastRefreshTime, startTime)
				require.Equal(suite.T(), 1, uc.deltaQueryCount)
				require.Equal(suite.T(), 3, len(uc.idToProps))
			},
		},
		{
			name: "deleted items",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
							fileItem("2", "file2", "root", "root", "https://dummy2.com", false),
							fileItem("1", "file1", "root", "root", "https://dummy1.com", true),
						},
						deltaLink: &deltaString,
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"1": {
					downloadURL: "",
					isDeleted:   true,
				},
				"2": {
					downloadURL: "https://dummy2.com",
					isDeleted:   false,
				},
			},
			expectedErr: require.NoError,
			cacheAssert: func(uc *urlCache, startTime time.Time) {
				require.Greater(suite.T(), uc.lastRefreshTime, startTime)
				require.Equal(suite.T(), 1, uc.deltaQueryCount)
				require.Equal(suite.T(), 2, len(uc.idToProps))
			},
		},
		{
			name: "item not found in cache",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
						},
						deltaLink: &deltaString,
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"2": {},
			},
			expectedErr: require.Error,
			cacheAssert: func(uc *urlCache, startTime time.Time) {
				require.Greater(suite.T(), uc.lastRefreshTime, startTime)
				require.Equal(suite.T(), 1, uc.deltaQueryCount)
				require.Equal(suite.T(), 1, len(uc.idToProps))
			},
		},
		{
			name: "multi-page delta query error",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
						},
						nextLink: &next,
					},
					{
						items: []models.DriveItemable{
							fileItem("2", "file2", "root", "root", "https://dummy2.com", false),
						},
						deltaLink: &deltaString,
						err:       errors.New("delta query error"),
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"1": {},
				"2": {},
			},
			expectedErr: require.Error,
			cacheAssert: func(uc *urlCache, _ time.Time) {
				require.Equal(suite.T(), time.Time{}, uc.lastRefreshTime)
				require.Equal(suite.T(), 0, uc.deltaQueryCount)
				require.Equal(suite.T(), 0, len(uc.idToProps))
			},
		},

		{
			name: "folder item",
			pagerResult: map[string][]deltaPagerResult{
				driveID: {
					{
						items: []models.DriveItemable{
							fileItem("1", "file1", "root", "root", "https://dummy1.com", false),
							driveItem("2", "folder2", "root", "root", false, true, false),
						},
						deltaLink: &deltaString,
					},
				},
			},
			expectedItemProps: map[string]itemProps{
				"2": {},
			},
			expectedErr: require.Error,
			cacheAssert: func(uc *urlCache, startTime time.Time) {
				require.Greater(suite.T(), uc.lastRefreshTime, startTime)
				require.Equal(suite.T(), 1, uc.deltaQueryCount)
				require.Equal(suite.T(), 1, len(uc.idToProps))
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)
			defer flush()

			itemPager := &mockItemPager{
				toReturn: test.pagerResult[driveID],
			}

			cache, err := newURLCache(
				driveID,
				1*time.Hour,
				itemPager,
				fault.New(true))

			require.NoError(suite.T(), err, clues.ToCore(err))

			numConcurrentReq := 100
			var wg sync.WaitGroup
			wg.Add(numConcurrentReq)

			startTime := time.Now()

			for i := 0; i < numConcurrentReq; i++ {
				go func() {
					defer wg.Done()

					for id, expected := range test.expectedItemProps {
						time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

						props, err := cache.getItemProperties(ctx, id)

						test.expectedErr(suite.T(), err, clues.ToCore(err))
						require.Equal(suite.T(), expected, props)
					}
				}()
			}

			wg.Wait()

			test.cacheAssert(cache, startTime)
		})
	}
}

// Test needsRefresh
func (suite *URLCacheUnitSuite) TestNeedsRefresh() {
	driveID := "drive1"
	t := suite.T()
	refreshInterval := 2 * time.Second

	cache, err := newURLCache(
		driveID,
		refreshInterval,
		&mockItemPager{},
		fault.New(true))

	require.NoError(t, err, clues.ToCore(err))

	// cache is empty
	require.True(t, cache.needsRefresh())

	// cache is not empty, but refresh interval has passed
	cache.idToProps["1"] = itemProps{
		downloadURL: "https://dummy1.com",
		isDeleted:   false,
	}

	time.Sleep(refreshInterval)
	require.True(t, cache.needsRefresh())

	// none of the above
	cache.lastRefreshTime = time.Now()
	require.False(t, cache.needsRefresh())

}

// Test newURLCache
func (suite *URLCacheUnitSuite) TestNewURLCache() {
	// table driven tests
	table := []struct {
		name        string
		driveID     string
		refreshInt  time.Duration
		itemPager   api.DriveItemEnumerator
		errors      *fault.Bus
		expectedErr require.ErrorAssertionFunc
	}{
		{
			name:        "invalid driveID",
			driveID:     "",
			refreshInt:  1 * time.Hour,
			itemPager:   &mockItemPager{},
			errors:      fault.New(true),
			expectedErr: require.Error,
		},
		{
			name:        "invalid refresh interval",
			driveID:     "drive1",
			refreshInt:  100 * time.Millisecond,
			itemPager:   &mockItemPager{},
			errors:      fault.New(true),
			expectedErr: require.Error,
		},
		{
			name:        "invalid itemPager",
			driveID:     "drive1",
			refreshInt:  1 * time.Hour,
			itemPager:   nil,
			errors:      fault.New(true),
			expectedErr: require.Error,
		},
		{
			name:        "nil error bus",
			driveID:     "drive1",
			refreshInt:  1 * time.Hour,
			itemPager:   &mockItemPager{},
			errors:      nil,
			expectedErr: require.Error,
		},
		{
			name:        "valid",
			driveID:     "drive1",
			refreshInt:  1 * time.Hour,
			itemPager:   &mockItemPager{},
			errors:      fault.New(true),
			expectedErr: require.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			_, err := newURLCache(
				test.driveID,
				test.refreshInt,
				test.itemPager,
				test.errors)

			test.expectedErr(t, err, clues.ToCore(err))
		})
	}
}
