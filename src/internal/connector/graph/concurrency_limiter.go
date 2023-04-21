package graph

import (
	"net/http"
	"sync"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
)

// concurrencyLimiter middleware limits the number of concurrent requests to graph API
type concurrencyLimiter struct {
	semaphore chan struct{}
}

var (
	once                  sync.Once
	concurrencyLim        *concurrencyLimiter
	maxConcurrentRequests = 4
)

func generateConcurrencyLimiter(capacity int) *concurrencyLimiter {
	if capacity < 1 || capacity > maxConcurrentRequests {
		capacity = maxConcurrentRequests
	}

	return &concurrencyLimiter{
		semaphore: make(chan struct{}, capacity),
	}
}

func InitializeConcurrencyLimiter(capacity int) {
	once.Do(func() {
		concurrencyLim = generateConcurrencyLimiter(capacity)
	})
}

func (cl *concurrencyLimiter) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	if cl == nil {
		return nil, clues.New("nil concurrency limiter")
	}

	cl.semaphore <- struct{}{}
	defer func() {
		<-cl.semaphore
	}()

	return pipeline.Next(req, middlewareIndex)
}
