package onedrive

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

// Unit tests for urlCache

type URLCacheUnitSuite struct {
	tester.Suite
}

func TestURLCacheUnitSuite(t *testing.T) {
	suite.Run(t, &URLCacheUnitSuite{Suite: tester.NewUnitSuite(t)})
}

/*
Unit test list

1. Test updateCache - against []DI returned by mock delta queries - itempager
	- Duplicate DIs. Latest DI should prevail
	- Deleted DIs
	- DIs with no download URL
	- DIs with download URL
	- DIs with download URL and deleted
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
	-

5. Test updateRefreshTime
	- Validate that refresh time is updated
	- Error case - new refresh time can never be < old refresh time

6. collectDriveItems
	- See collectItems tests

7.
*/

/*
Integration test list - use gock to simulate 401s on refresh interval expiry

1. Test getDownloadURL
2. Test refreshCache

*/
