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
