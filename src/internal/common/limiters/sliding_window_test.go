package limiters

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type SlidingWindowUnitTestSuite struct {
	tester.Suite
}

func TestSlidingWindowLimiterSuite(t *testing.T) {
	suite.Run(t, &SlidingWindowUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

// TestWaitBasic tests the Wait() functionality of the limiter with multiple
// concurrent requests.
func (suite *SlidingWindowUnitTestSuite) TestWaitBasic() {
	var (
		t          = suite.T()
		windowSize = 1 * time.Second
		// Assume slide interval is equal to window size for simplicity.
		slideInterval   = 1 * time.Second
		capacity        = 100
		startTime       = time.Now()
		numRequests     = 3 * capacity
		wg              sync.WaitGroup
		mu              sync.Mutex
		intervalToCount = make(map[time.Duration]int)
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	s, err := NewSlidingWindowLimiter(windowSize, slideInterval, capacity)
	require.NoError(t, err)

	defer s.Shutdown()

	// Check if all tokens are available for use post initialization.
	require.Equal(t, capacity, len(s.(*slidingWindow).permits))

	// Make concurrent requests to the limiter
	for i := 0; i < numRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := s.Wait(ctx)
			require.NoError(t, err)

			// Number of seconds since startTime
			bucket := time.Since(startTime).Truncate(windowSize)

			mu.Lock()
			intervalToCount[bucket]++
			mu.Unlock()
		}()
	}

	wg.Wait()

	// Verify that number of requests allowed in each window is less than or equal
	// to window capacity
	for _, c := range intervalToCount {
		require.True(t, c <= capacity, "count: %d, capacity: %d", c, capacity)
	}
}

// TestWaitSliding tests the sliding window functionality of the limiter with distributed
// Wait() calls.
func (suite *SlidingWindowUnitTestSuite) TestWaitSliding() {
	var (
		t             = suite.T()
		windowSize    = 1 * time.Second
		slideInterval = 10 * time.Millisecond
		capacity      = 100
		// Test will run for duration of 2 windowSize.
		numRequests = 2 * capacity
		wg          sync.WaitGroup
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	s, err := NewSlidingWindowLimiter(windowSize, slideInterval, capacity)
	require.NoError(t, err)

	defer s.Shutdown()

	// Make concurrent requests to the limiter
	for i := 0; i < numRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Sleep for a random duration to spread out requests over multiple slide
			// intervals & windows, so that we can test the sliding window logic better.
			// Without this, the requests will be bunched up in the very first intervals
			// of the 2 windows. Rest of the intervals will be empty.
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)

			err := s.Wait(ctx)
			require.NoError(t, err)
		}()
	}
	wg.Wait()

	// Verify that number of requests allowed in each window is less than or equal
	// to window capacity
	sw := s.(*slidingWindow)
	data := append(sw.prev.count, sw.curr.count...)

	sums := slidingSum(data, sw.numIntervals)

	for _, sum := range sums {
		fmt.Printf("sum: %d\n", sum)
		require.True(t, sum <= capacity, "sum: %d, capacity: %d", sum, capacity)
	}
}

func (suite *SlidingWindowUnitTestSuite) TestContextCancellation() {
	var (
		t             = suite.T()
		windowSize    = 100 * time.Millisecond
		slideInterval = 10 * time.Millisecond
		wg            sync.WaitGroup
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// Initialize limiter with capacity = 0 to test context cancellations.
	s, err := NewSlidingWindowLimiter(windowSize, slideInterval, 0)
	require.NoError(t, err)

	defer s.Shutdown()

	ctx, cancel := context.WithTimeout(ctx, 2*windowSize)
	defer cancel()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err := s.Wait(ctx)
		require.Equal(t, context.DeadlineExceeded, err)
	}()

	wg.Wait()
}

func (suite *SlidingWindowUnitTestSuite) TestNewSlidingWindowLimiter() {
	tests := []struct {
		name          string
		windowSize    time.Duration
		slideInterval time.Duration
		capacity      int
		expectErr     assert.ErrorAssertionFunc
	}{
		{
			name:          "Invalid window size",
			windowSize:    0,
			slideInterval: 10 * time.Millisecond,
			capacity:      100,
			expectErr:     assert.Error,
		},
		{
			name:          "Invalid slide interval",
			windowSize:    100 * time.Millisecond,
			slideInterval: 0,
			capacity:      100,
			expectErr:     assert.Error,
		},
		{
			name:          "Slide interval > window size",
			windowSize:    10 * time.Millisecond,
			slideInterval: 100 * time.Millisecond,
			capacity:      100,
			expectErr:     assert.Error,
		},
		{
			name:          "Invalid capacity",
			windowSize:    100 * time.Millisecond,
			slideInterval: 10 * time.Millisecond,
			capacity:      -1,
			expectErr:     assert.Error,
		},
		{
			name:          "Valid parameters",
			windowSize:    100 * time.Millisecond,
			slideInterval: 10 * time.Millisecond,
			capacity:      100,
			expectErr:     assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			s, err := NewSlidingWindowLimiter(
				test.windowSize,
				test.slideInterval,
				test.capacity)
			if s != nil {
				defer s.Shutdown()
			}

			test.expectErr(t, err)
		})
	}
}

func slidingSum(data []int, w int) []int {
	var (
		sum = 0
		res = make([]int, len(data)-w+1)
	)

	for i := 0; i < w; i++ {
		sum += data[i]
	}

	res[0] = sum

	for i := 1; i < len(data)-w+1; i++ {
		sum = sum - data[i-1] + data[i+w-1]
		res[i] = sum
	}

	return res
}
