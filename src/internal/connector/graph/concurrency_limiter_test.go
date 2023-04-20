package graph

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	khttp "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ConcurrencyLimiterTestSuite struct {
	tester.Suite
}

func TestConcurrencyLimiterSuite(t *testing.T) {
	suite.Run(t, &ConcurrencyLimiterTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ConcurrencyLimiterTestSuite) TestConcurrencyLimiterMiddleware() {
	t := suite.T()
	maxConcurrentRequests := 4

	InitializeConcurrencyLimiter(maxConcurrentRequests)

	cl := GetConcurrencyLimiter()
	client := khttp.GetDefaultClient(cl)

	sem := make(chan struct{}, maxConcurrentRequests)
	reqHandler := func(w http.ResponseWriter, r *http.Request) {
		select {
		case sem <- struct{}{}:
			defer func() {
				<-sem
			}()

			time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)
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
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resp, err := client.Get(ts.URL)
			require.NoError(t, err)
			assert.Equal(t, resp.StatusCode, http.StatusOK)
		}()
	}
	wg.Wait()
}
