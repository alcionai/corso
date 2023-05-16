package graph

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"

	"github.com/alcionai/corso/src/internal/tester"
)

type ConcurrencyMWUnitTestSuite struct {
	tester.Suite
}

func TestConcurrencyLimiterSuite(t *testing.T) {
	suite.Run(t, &ConcurrencyMWUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ConcurrencyMWUnitTestSuite) TestConcurrencyLimiter() {
	t := suite.T()

	maxConcurrentRequests := 4
	cl := generateConcurrencyLimiter(maxConcurrentRequests)
	client := khttp.GetDefaultClient(cl)

	// Server side handler to simulate 429s
	sem := make(chan struct{}, maxConcurrentRequests)
	reqHandler := func(w http.ResponseWriter, r *http.Request) {
		select {
		case sem <- struct{}{}:
			defer func() {
				<-sem
			}()

			time.Sleep(time.Duration(rand.Intn(50)+50) * time.Millisecond)
			w.WriteHeader(http.StatusOK)

			return
		default:
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(reqHandler))
	defer ts.Close()

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resp, err := client.Get(ts.URL)
			require.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}()
	}
	wg.Wait()
}

func (suite *ConcurrencyMWUnitTestSuite) TestInitializeConcurrencyLimiter() {
	t := suite.T()

	InitializeConcurrencyLimiter(2)
	InitializeConcurrencyLimiter(4)

	assert.Equal(t, cap(concurrencyLim.semaphore), 2, "singleton semaphore capacity changed")
}

func (suite *ConcurrencyMWUnitTestSuite) TestGenerateConcurrencyLimiter() {
	tests := []struct {
		name        string
		cap         int
		expectedCap int
	}{
		{
			name:        "valid capacity",
			cap:         2,
			expectedCap: 2,
		},
		{
			name:        "zero capacity",
			cap:         0,
			expectedCap: maxConcurrentRequests,
		},
		{
			name:        "negative capacity",
			cap:         -1,
			expectedCap: maxConcurrentRequests,
		},
		{
			name:        "out of bounds capacity",
			cap:         10,
			expectedCap: maxConcurrentRequests,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			actual := generateConcurrencyLimiter(test.cap)
			assert.Equal(t, cap(actual.semaphore), test.expectedCap,
				"retrieved semaphore capacity vs expected capacity")
		})
	}
}

func (suite *ConcurrencyMWUnitTestSuite) TestTimedFence_Block() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	tf := newTimedFence()

	// raise multiple fences, the longest at 5 seconds
	for i := -5; i < 6; i++ {
		tf.RaiseFence(time.Duration(i) * time.Second)
	}

	// -5..0 get dropped, 1..5 get added
	assert.Len(t, tf.timers, 5)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			err := tf.Block(ctx)
			require.NoError(t, err, clues.ToCore(err))
		}(i)
	}

	wg.Wait()

	// should block for 5 seconds.  comparing to 4 to avoid
	// race condition flakes.
	assert.Less(t, 4.0, time.Since(start).Seconds())
}

func (suite *ConcurrencyMWUnitTestSuite) TestTimedFence_Block_ctxDeadline() {
	ctx, flush := tester.NewContext()
	defer flush()

	ctx, _ = context.WithDeadline(ctx, time.Now().Add(2*time.Second))

	t := suite.T()
	tf := newTimedFence()

	// raise multiple fences, the longest at 10 seconds
	for i := 1; i < 6; i++ {
		tf.RaiseFence(time.Duration(i*2) * time.Second)
	}

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			err := tf.Block(ctx)
			// should error from ctx deadline
			require.Error(t, err, clues.ToCore(err))
		}(i)
	}

	wg.Wait()

	// should block for 2 seconds.  comparing to 3 to avoid
	// race condition flakes.
	assert.Greater(t, 3.0, time.Since(start).Seconds())
}

type mockPipeline struct {
	resp *http.Response
	err  error
}

func (mp mockPipeline) Next(*http.Request, int) (*http.Response, error) {
	return mp.resp, mp.err
}

func (suite *ConcurrencyMWUnitTestSuite) TestThrottlingMiddleware() {
	retryAfterNan := http.Header{}
	retryAfterNan.Set(retryAfterHeader, "brunhuldi")

	retryAfterNeg1 := http.Header{}
	retryAfterNeg1.Set(retryAfterHeader, "-1")

	retryAfter0 := http.Header{}
	retryAfter0.Set(retryAfterHeader, "0")

	retryAfter5 := http.Header{}
	retryAfter5.Set(retryAfterHeader, "5")

	goodPipe := mockPipeline{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
		},
	}

	table := []struct {
		name          string
		pipeline      mockPipeline
		expectMinWait float64
	}{
		{
			name:          "2xx response",
			pipeline:      goodPipe,
			expectMinWait: 0,
		},
		{
			name: "non-429 response",
			pipeline: mockPipeline{
				resp: &http.Response{
					StatusCode: http.StatusBadGateway,
					Header:     retryAfter5,
				},
			},
			expectMinWait: 0,
		},
		{
			name: "429 response w/out retry header",
			pipeline: mockPipeline{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header:     http.Header{},
				},
			},
			expectMinWait: 0,
		},
		{
			name: "429 response w/ nan retry-after",
			pipeline: mockPipeline{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header:     retryAfterNan,
				},
			},
			expectMinWait: 0,
		},
		{
			name: "429 response w/ negative retry-after",
			pipeline: mockPipeline{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header:     retryAfterNeg1,
				},
			},
			expectMinWait: 0,
		},
		{
			name: "429 response w/ zero retry-after",
			pipeline: mockPipeline{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header:     retryAfter0,
				},
			},
			expectMinWait: 0,
		},
		{
			name: "429 response w/ positive retry-after",
			pipeline: mockPipeline{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header:     retryAfter5,
				},
			},
			expectMinWait: 4,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			tm := throttlingMiddleware{newTimedFence()}

			req := &http.Request{}
			req = req.WithContext(ctx)

			start := time.Now()

			_, err := tm.Intercept(test.pipeline, 0, req)
			require.NoError(t, err, clues.ToCore(err))

			_, err = tm.Intercept(goodPipe, 0, req)
			require.NoError(t, err, clues.ToCore(err))

			assert.Less(t, test.expectMinWait, time.Since(start).Seconds())
		})
	}
}
