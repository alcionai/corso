package graph

import (
	"net/http"
	"sync"

	khttp "github.com/microsoft/kiota-http-go"
)

// ConcurrencyLimiterMiddleware limits the number of concurrent requests.
// Currently we only enforce concurrency limit for exchange services
type ConcurrencyLimiterMiddleware struct {
	semaphoreExch chan struct{}
}

// TODO: Instead of creating a singleton and doing per graph API request
// checks for exchange/non-exchange, associate this middleware only with
// exchange workloads by making changes in this code path
// https://github.com/alcionai/corso/blob/e5136ceabbe85c17d72618992c0ce2cf2fa7faf7/src/internal/connector/graph/service.go#L148

var once sync.Once
var concurrencyLimiter *ConcurrencyLimiterMiddleware

func GetConcurrencyLimiterMiddleware(capacity int) (*ConcurrencyLimiterMiddleware, error) {
	// TODO: add capacity checks
	// TODO: Don't return an error. Default to a capacity if invalid.
	once.Do(func() {
		concurrencyLimiter = &ConcurrencyLimiterMiddleware{
			semaphoreExch: make(chan struct{}, capacity),
		}
	})
	return concurrencyLimiter, nil
}

func (cl *ConcurrencyLimiterMiddleware) Intercept(
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
