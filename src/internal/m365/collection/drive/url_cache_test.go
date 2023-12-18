package drive

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

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
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *URLCacheIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.user = tconfig.SecondaryM365UserID(t)

	acct := tconfig.NewM365Account(t)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.ac, err = api.NewClient(
		creds,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	drive, err := suite.ac.Users().GetDefaultDrive(ctx, suite.user)
	require.NoError(t, err, clues.ToCore(err))

	suite.driveID = ptr.Val(drive.GetId())
}

// Basic test for urlCache. Create some files in onedrive, then access them via
// url cache
func (suite *URLCacheIntegrationSuite) TestURLCacheBasic() {
	var (
		t             = suite.T()
		ac            = suite.ac.Drives()
		driveID       = suite.driveID
		newFolderName = testdata.DefaultRestoreConfig("folder").Location
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// Create a new test folder
	root, err := ac.GetRootFolder(ctx, driveID)
	require.NoError(t, err, clues.ToCore(err))

	newFolder, err := ac.PostItemInContainer(
		ctx,
		driveID,
		ptr.Val(root.GetId()),
		api.NewDriveItem(newFolderName, true),
		control.Copy)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, newFolder.GetId())

	nfid := ptr.Val(newFolder.GetId())

	// Get the previous delta to feed into url cache
	pager := ac.EnumerateDriveItemsDelta(
		ctx,
		suite.driveID,
		"",
		api.CallConfig{
			Select: api.URLCacheDriveItemProps(),
		})

	// We need to go through all the pages of results so we don't get stuck. This
	// is the only way to get a delta token since getting one requires going
	// through all request results.
	//
	//revive:disable-next-line:empty-block
	for _, _, done := pager.NextPage(); !done; _, _, done = pager.NextPage() {
	}

	du, err := pager.Results()
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, du.URL)

	// Create a bunch of files in the new folder
	var items []models.DriveItemable

	for i := 0; i < 5; i++ {
		newItemName := "test_url_cache_basic_" + dttm.FormatNow(dttm.SafeForTesting)

		item, err := ac.PostItemInContainer(
			ctx,
			driveID,
			nfid,
			api.NewDriveItem(newItemName, false),
			control.Copy)
		require.NoError(t, err, clues.ToCore(err))

		items = append(items, item)
	}

	// Create a new URL cache with a long TTL
	uc, err := newURLCache(
		suite.driveID,
		du.URL,
		1*time.Hour,
		suite.ac.Drives(),
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	// Launch parallel requests to the cache, one per item
	var wg sync.WaitGroup
	for i := 0; i < len(items); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Read item from URL cache
			props, err := uc.getItemProperties(
				ctx,
				ptr.Val(items[i].GetId()))
			require.NoError(t, err, clues.ToCore(err))

			require.NotNil(t, props)
			require.NotEmpty(t, props.downloadURL)
			require.Equal(t, false, props.isDeleted)

			// Validate download URL
			c := graph.NewNoTimeoutHTTPWrapper(count.New())

			resp, err := c.Request(
				ctx,
				http.MethodGet,
				props.downloadURL,
				nil,
				nil)
			require.NoError(t, err, clues.ToCore(err))

			require.NotNil(t, resp)
			require.NotNil(t, resp.Body)

			defer func(rc io.ReadCloser) {
				if rc != nil {
					rc.Close()
				}
			}(resp.Body)

			require.Equal(t, http.StatusOK, resp.StatusCode)
		}(i)
	}
	wg.Wait()

	// Validate that exactly 1 delta query was made by url cache
	require.Equal(t, 1, uc.refreshCount)

	// Validate that the prev delta base stays the same
	require.Equal(t, du.URL, uc.prevDelta)
}

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type URLCacheUnitSuite struct {
	tester.Suite
}

func TestURLCacheUnitSuite(t *testing.T) {
	suite.Run(t, &URLCacheUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *URLCacheUnitSuite) TestGetItemProperties() {
	d := drive()

	aURL := func(n int) string {
		return fmt.Sprintf("https://dummy%d.com", n)
	}

	table := []struct {
		name              string
		pages             []nextPage
		pagerErr          error
		expectedItemProps map[string]itemProps
		expectErr         assert.ErrorAssertionFunc
		expect            func(*testing.T, *urlCache, time.Time)
	}{
		{
			name: "single item in cache",
			pages: []nextPage{
				aPage(d.fileWURLAtRoot(aURL(1), false, 1)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: aURL(1),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 1, len(uc.idToProps))
			},
		},
		{
			name: "multiple items in cache",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3),
					d.fileWURLAtRoot(aURL(4), false, 4),
					d.fileWURLAtRoot(aURL(5), false, 5)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: aURL(1),
					isDeleted:   false,
				},
				fileID(2): {
					downloadURL: aURL(2),
					isDeleted:   false,
				},
				fileID(3): {
					downloadURL: aURL(3),
					isDeleted:   false,
				},
				fileID(4): {
					downloadURL: aURL(4),
					isDeleted:   false,
				},
				fileID(5): {
					downloadURL: aURL(5),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 5, len(uc.idToProps))
			},
		},
		{
			name: "multiple pages",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3)),
				aPage(
					d.fileWURLAtRoot(aURL(4), false, 4),
					d.fileWURLAtRoot(aURL(5), false, 5)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: aURL(1),
					isDeleted:   false,
				},
				fileID(2): {
					downloadURL: aURL(2),
					isDeleted:   false,
				},
				fileID(3): {
					downloadURL: aURL(3),
					isDeleted:   false,
				},
				fileID(4): {
					downloadURL: aURL(4),
					isDeleted:   false,
				},
				fileID(5): {
					downloadURL: aURL(5),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 5, len(uc.idToProps))
			},
		},
		{
			name: "multiple pages with resets",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(-1), false, -1),
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3)),
				aReset(),
				aPage(
					d.fileWURLAtRoot(aURL(0), false, 0),
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3)),
				aPage(
					d.fileWURLAtRoot(aURL(4), false, 4),
					d.fileWURLAtRoot(aURL(5), false, 5)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: aURL(1),
					isDeleted:   false,
				},
				fileID(2): {
					downloadURL: aURL(2),
					isDeleted:   false,
				},
				fileID(3): {
					downloadURL: aURL(3),
					isDeleted:   false,
				},
				fileID(4): {
					downloadURL: aURL(4),
					isDeleted:   false,
				},
				fileID(5): {
					downloadURL: aURL(5),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 6, len(uc.idToProps))
			},
		},
		{
			name: "multiple pages with resets and combo reset+items in page",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(0), false, 0),
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3)),
				aPageWReset(
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3)),
				aPage(
					d.fileWURLAtRoot(aURL(4), false, 4),
					d.fileWURLAtRoot(aURL(5), false, 5)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: aURL(1),
					isDeleted:   false,
				},
				fileID(2): {
					downloadURL: aURL(2),
					isDeleted:   false,
				},
				fileID(3): {
					downloadURL: aURL(3),
					isDeleted:   false,
				},
				fileID(4): {
					downloadURL: aURL(4),
					isDeleted:   false,
				},
				fileID(5): {
					downloadURL: aURL(5),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 5, len(uc.idToProps))
			},
		},
		{
			name: "duplicate items with potentially new urls",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(3), false, 3),
					d.fileWURLAtRoot(aURL(100), false, 1),
					d.fileWURLAtRoot(aURL(200), false, 2)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: aURL(100),
					isDeleted:   false,
				},
				fileID(2): {
					downloadURL: aURL(200),
					isDeleted:   false,
				},
				fileID(3): {
					downloadURL: aURL(3),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 3, len(uc.idToProps))
			},
		},
		{
			name: "deleted items",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.fileWURLAtRoot(aURL(2), false, 2),
					d.fileWURLAtRoot(aURL(1), true, 1)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(1): {
					downloadURL: "",
					isDeleted:   true,
				},
				fileID(2): {
					downloadURL: aURL(2),
					isDeleted:   false,
				},
			},
			expectErr: assert.NoError,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 2, len(uc.idToProps))
			},
		},
		{
			name: "item not found in cache",
			pages: []nextPage{
				aPage(d.fileWURLAtRoot(aURL(1), false, 1)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(2): {},
			},
			expectErr: assert.Error,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 1, len(uc.idToProps))
			},
		},
		{
			name: "delta query error",
			pages: []nextPage{
				aPage(),
			},
			pagerErr: errors.New("delta query error"),
			expectedItemProps: map[string]itemProps{
				fileID(1): {},
				fileID(2): {},
			},
			expectErr: assert.Error,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Equal(t, time.Time{}, uc.lastRefreshTime)
				assert.NotZero(t, uc.refreshCount)
				assert.Equal(t, 0, len(uc.idToProps))
			},
		},
		{
			name: "folder item",
			pages: []nextPage{
				aPage(
					d.fileWURLAtRoot(aURL(1), false, 1),
					d.folderAtRoot(2)),
			},
			expectedItemProps: map[string]itemProps{
				fileID(2): {},
			},
			expectErr: assert.Error,
			expect: func(t *testing.T, uc *urlCache, startTime time.Time) {
				assert.Greater(t, uc.lastRefreshTime, startTime)
				assert.Equal(t, 1, uc.refreshCount)
				assert.Equal(t, 1, len(uc.idToProps))
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			for _, numConcurrentReqs := range []int{1, 2, 32} {
				crTestName := fmt.Sprintf("%d_concurrent_reqs", numConcurrentReqs)
				suite.Run(crTestName, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					drive := drive()

					driveEnumer := driveEnumerator(
						drive.newEnumer().
							withErr(test.pagerErr).
							with(
								delta(deltaURL, test.pagerErr).
									with(test.pages...)))

					cache, err := newURLCache(
						drive.id,
						"",
						1*time.Hour,
						driveEnumer,
						count.New(),
						fault.New(true))
					require.NoError(t, err, clues.ToCore(err))

					var wg sync.WaitGroup
					wg.Add(numConcurrentReqs)

					startTime := time.Now()

					for i := 0; i < numConcurrentReqs; i++ {
						go func(ti int) {
							defer wg.Done()

							for id, expected := range test.expectedItemProps {
								time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

								props, err := cache.getItemProperties(ctx, id)
								test.expectErr(t, err, clues.ToCore(err))
								assert.Equal(t, expected, props)
							}
						}(i)
					}

					wg.Wait()

					test.expect(t, cache, startTime)
				})
			}
		})
	}
}

// Test needsRefresh
func (suite *URLCacheUnitSuite) TestNeedsRefresh() {
	var (
		t               = suite.T()
		refreshInterval = 1 * time.Second
		drv             = drive()
	)

	cache, err := newURLCache(
		drv.id,
		"",
		refreshInterval,
		&enumerateDriveItemsDelta{},
		count.New(),
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

func (suite *URLCacheUnitSuite) TestNewURLCache() {
	drv := drive()

	table := []struct {
		name       string
		driveID    string
		refreshInt time.Duration
		itemPager  EnumerateDriveItemsDeltaer
		errors     *fault.Bus
		expectErr  require.ErrorAssertionFunc
	}{
		{
			name:       "invalid driveID",
			driveID:    "",
			refreshInt: 1 * time.Hour,
			itemPager:  &enumerateDriveItemsDelta{},
			errors:     fault.New(true),
			expectErr:  require.Error,
		},
		{
			name:       "invalid refresh interval",
			driveID:    drv.id,
			refreshInt: 100 * time.Millisecond,
			itemPager:  &enumerateDriveItemsDelta{},
			errors:     fault.New(true),
			expectErr:  require.Error,
		},
		{
			name:       "invalid item enumerator",
			driveID:    drv.id,
			refreshInt: 1 * time.Hour,
			itemPager:  nil,
			errors:     fault.New(true),
			expectErr:  require.Error,
		},
		{
			name:       "valid",
			driveID:    drv.id,
			refreshInt: 1 * time.Hour,
			itemPager:  &enumerateDriveItemsDelta{},
			errors:     fault.New(true),
			expectErr:  require.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			_, err := newURLCache(
				test.driveID,
				"",
				test.refreshInt,
				test.itemPager,
				count.New(),
				test.errors)

			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
