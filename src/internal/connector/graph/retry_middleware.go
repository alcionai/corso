package graph

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/alcionai/corso/src/internal/connector/support"
	backoff "github.com/cenkalti/backoff/v4"
	khttp "github.com/microsoft/kiota-http-go"
)

// ---------------------------------------------------------------------------
// Client Middleware
// ---------------------------------------------------------------------------

// RetryHandler handles transient HTTP responses and retries the request given the retry options
type RetryHandler struct {
	// default options to use when evaluating the response
	options RetryHandlerOptions
}

type RetryHandlerOptions struct {
	// request never retried if flag set to true
	NoRetry bool
	// The maximum number of times a request can be retried
	MaxRetries int
	// The delay in seconds between retries
	Delay time.Duration
}

func (middleware RetryHandler) retryRequest(
	ctx context.Context,
	pipeline khttp.Pipeline,
	middlewareIndex int,
	options RetryHandlerOptions,
	req *http.Request,
	resp *http.Response,
	executionCount int,
	cumulativeDelay time.Duration,
	exponentialBackoff *backoff.ExponentialBackOff,
	respErr error,
) (*http.Response, error) {
	if (respErr != nil || middleware.isRetriableErrorCode(req, resp.StatusCode)) &&
		middleware.isRetriableRequest(req) &&
		executionCount < options.MaxRetries &&
		!options.NoRetry &&
		cumulativeDelay < time.Duration(absoluteMaxDelaySeconds)*time.Second {
		executionCount++

		delay := middleware.getRetryDelay(req, resp, exponentialBackoff)

		cumulativeDelay += delay

		req.Header.Set(retryAttemptHeader, strconv.Itoa(executionCount))

		time.Sleep(delay)

		response, err := pipeline.Next(req, middlewareIndex)
		if err != nil && !IsErrTimeout(err) {
			return response, support.ConnectorStackErrorTraceWrap(err, "maximum retries or unretryable")
		}

		return middleware.retryRequest(ctx,
			pipeline,
			middlewareIndex,
			options,
			req,
			response,
			executionCount,
			cumulativeDelay,
			exponentialBackoff,
			err)
	}

	return resp, support.ConnectorStackErrorTraceWrap(respErr, "maximum retries or unretryable")
}

func (middleware RetryHandler) isRetriableErrorCode(req *http.Request, code int) bool {
	return code == http.StatusInternalServerError
}

func (middleware RetryHandler) isRetriableRequest(req *http.Request) bool {
	isBodiedMethod := req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH"
	if isBodiedMethod && req.Body != nil {
		return req.ContentLength != -1
	}

	return true
}

func (middleware RetryHandler) getRetryDelay(
	req *http.Request,
	resp *http.Response,
	exponentialBackoff *backoff.ExponentialBackOff,
) time.Duration {
	var retryAfter string
	if resp != nil {
		retryAfter = resp.Header.Get(retryAfterHeader)
	}

	if retryAfter != "" {
		retryAfterDelay, err := strconv.ParseFloat(retryAfter, 64)
		if err == nil {
			return time.Duration(retryAfterDelay) * time.Second
		}
	} // TODO parse the header if it's a date

	return exponentialBackoff.NextBackOff()
}
