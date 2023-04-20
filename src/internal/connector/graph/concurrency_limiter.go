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

var once sync.Once
var concurrencyLimiter *ConcurrencyLimiterMiddleware

func GetConcurrencyLimiterMiddleware(capacity int) (*ConcurrencyLimiterMiddleware, error) {
	// TODO: add capacity checks
	// Default to a capacity if invalid. Don't return an error
	once.Do(func() {
		concurrencyLimiter = &ConcurrencyLimiterMiddleware{
			semaphoreExch: make(chan struct{}, capacity),
		}
	})
	return concurrencyLimiter, nil
}

// TODO: Probably also need a reinitialize or delete for sem?

func (cl *ConcurrencyLimiterMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	// Acquire a slot in the semaphore
	cl.semaphoreExch <- struct{}{}
	defer func() {
		<-cl.semaphoreExch
	}()

	// Call the next middleware in the pipeline
	return pipeline.Next(req, middlewareIndex)
}
