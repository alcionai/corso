package graph

import (
	"net/http"
	"sync"

	khttp "github.com/microsoft/kiota-http-go"
)

// ConcurrencyLimiter middleware limits the number of concurrent requests to graph API.
// Currently we only enforce concurrency limit for exchange
type ConcurrencyLimiter struct {
	semaphoreExch chan struct{}
}

var (
	once               sync.Once
	concurrencyLimiter *ConcurrencyLimiter
	// Outlooks expects max 4 concurrent requests
	// https://learn.microsoft.com/en-us/graph/throttling-limits#outlook-service-limits
	maxConcurrentRequestsExchange = 4
)

func InitializeConcurrencyLimiter(capacity int) {
	once.Do(func() {
		if capacity < 1 || capacity > maxConcurrentRequestsExchange {
			capacity = maxConcurrentRequestsExchange
		}

		concurrencyLimiter = &ConcurrencyLimiter{
			semaphoreExch: make(chan struct{}, capacity),
		}
	})
}

func GetConcurrencyLimiter() *ConcurrencyLimiter {
	return concurrencyLimiter
}

func (cl *ConcurrencyLimiter) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	cl.semaphoreExch <- struct{}{}
	defer func() {
		<-cl.semaphoreExch
	}()

	// Call the next middleware in the pipeline
	return pipeline.Next(req, middlewareIndex)
}
