package limiters

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"

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

	defer goleak.VerifyNone(t)

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

// TestWaitSliding tests the sliding window functionality of the limiter with
// time distributed WaitN() calls.
func (suite *SlidingWindowUnitTestSuite) TestWaitNSliding() {
	tests := []struct {
		Name          string
		windowSize    time.Duration
		slideInterval time.Duration
		capacity      int
		numRequests   int
		n             int
	}{
		{
			Name:          "Request 1 token each",
			windowSize:    100 * time.Millisecond,
			slideInterval: 10 * time.Millisecond,
			capacity:      100,
			numRequests:   200,
			n:             1,
		},
		{
			Name:          "Request 5 tokens each",
			windowSize:    100 * time.Millisecond,
			slideInterval: 10 * time.Millisecond,
			capacity:      1000,
			numRequests:   100,
			n:             5,
		},
	}

	for _, test := range tests {
		suite.Run(test.Name, func() {
			t := suite.T()

			defer goleak.VerifyNone(t)

			ctx, flush := tester.NewContext(t)
			defer flush()

			s, err := NewSlidingWindowLimiter(test.windowSize, test.slideInterval, test.capacity)
			require.NoError(t, err)

			var wg sync.WaitGroup

			// Make concurrent requests to the limiter
			for i := 0; i < test.numRequests; i++ {
				wg.Add(1)

				go func() {
					defer wg.Done()

					// Sleep for a random duration to spread out requests over
					// multiple slide intervals & windows, so that we can test
					// the sliding window logic better.
					// Without this, the requests will be bunched up in the very
					// first interval of the 2 windows. Rest of the intervals
					// will be empty.
					time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)

					err := s.WaitN(ctx, test.n)
					require.NoError(t, err)
				}()
			}
			wg.Wait()

			// Shutdown the ticker before accessing the internal limiter state.
			s.Shutdown()

			// Verify that number of requests allowed in each window is less than or equal
			// to window capacity
			sw := s.(*slidingWindow)
			data := append(sw.prev.count, sw.curr.count...)

			sums := slidingSums(data, sw.numIntervals)

			for _, sum := range sums {
				require.True(t, sum <= test.capacity, "sum: %d, capacity: %d", sum, test.capacity)
			}
		})
	}
}

func (suite *SlidingWindowUnitTestSuite) TestContextCancellation() {
	t := suite.T()

	// Since this test can infinitely block on failure conditions, run it within
	// a time contained eventually block.
	assert.Eventually(t, func() bool {
		var (
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
			require.ErrorIs(t, err, context.DeadlineExceeded)
		}()

		wg.Wait()

		return true
	}, 3*time.Second, 100*time.Millisecond)
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
			name:          "Window not divisible by slide interval",
			windowSize:    100 * time.Millisecond,
			slideInterval: 11 * time.Millisecond,
			capacity:      100,
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

			defer goleak.VerifyNone(t)

			s, err := NewSlidingWindowLimiter(
				test.windowSize,
				test.slideInterval,
				test.capacity)
			if s != nil {
				s.Shutdown()
			}

			test.expectErr(t, err)
		})
	}
}

func slidingSums(data []int, w int) []int {
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

func (suite *SlidingWindowUnitTestSuite) TestShutdown() {
	var (
		t             = suite.T()
		windowSize    = 1 * time.Second
		slideInterval = 1 * time.Second
		capacity      = 100
	)

	defer goleak.VerifyNone(t)

	s, err := NewSlidingWindowLimiter(windowSize, slideInterval, capacity)
	require.NoError(t, err)

	s.Shutdown()

	// Second call to Shutdown() should be a no-op.
	s.Shutdown()
}

// TestReset tests if limiter state is cleared and all tokens are available for
// use post reset.
func (suite *SlidingWindowUnitTestSuite) TestReset() {
	var (
		t             = suite.T()
		windowSize    = 100 * time.Millisecond
		slideInterval = 10 * time.Millisecond
		capacity      = 10
		numRequests   = capacity
		wg            sync.WaitGroup
	)

	defer goleak.VerifyNone(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	s, err := NewSlidingWindowLimiter(windowSize, slideInterval, capacity)
	require.NoError(t, err)

	// Make some requests to the limiter.
	for i := 0; i < numRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := s.Wait(ctx)
			require.NoError(t, err)
		}()
	}
	wg.Wait()

	// Reset the limiter.
	s.Reset()

	// Shutdown the ticker before accessing the internal limiter state.
	s.Shutdown()

	sw := s.(*slidingWindow)

	// Check if state is cleared, and all tokens are available for use post reset.
	require.Equal(t, capacity, len(sw.permits))

	for i := 0; i < sw.numIntervals; i++ {
		require.Equal(t, 0, sw.prev.count[i])
		require.Equal(t, 0, sw.curr.count[i])
	}
}

// TestResetDuringActiveRequests tests if reset is transparent to any active
// requests and they are not affected. It also checks that limiter stays
// within capacity limits post reset.
func (suite *SlidingWindowUnitTestSuite) TestResetDuringActiveRequests() {
	var (
		t             = suite.T()
		windowSize    = 100 * time.Millisecond
		slideInterval = 10 * time.Millisecond
		capacity      = 10
		numRequests   = 10 * capacity
		wg            sync.WaitGroup
	)

	defer goleak.VerifyNone(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	s, err := NewSlidingWindowLimiter(windowSize, slideInterval, capacity)
	require.NoError(t, err)

	// Make some requests to the limiter as well as reset it concurrently
	// in 10:1 request to reset ratio.
	for i := 0; i < numRequests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := s.Wait(ctx)
			require.NoError(t, err)
		}()

		// Launch a reset every 10th iteration.
		if i%10 == 0 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				s.Reset()
			}()
		}
	}
	wg.Wait()

	// Shutdown the ticker before accessing the internal limiter state.
	s.Shutdown()

	// Verify that number of requests allowed in each window is less than or equal
	// to window capacity
	sw := s.(*slidingWindow)
	data := append(sw.prev.count, sw.curr.count...)

	sums := slidingSums(data, sw.numIntervals)

	for _, sum := range sums {
		require.True(t, sum <= capacity, "sum: %d, capacity: %d", sum, capacity)
	}
}
