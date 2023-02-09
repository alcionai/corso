package graph

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"time"

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
	DelaySeconds int
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
	respErr error,
) (*http.Response, error) {
	if (respErr != nil || middleware.isRetriableErrorCode(req, resp.StatusCode)) &&
		middleware.isRetriableRequest(req) &&
		executionCount < options.MaxRetries &&
		!options.NoRetry &&
		cumulativeDelay < time.Duration(absoluteMaxDelaySeconds)*time.Second {
		executionCount++
		delay := middleware.getRetryDelay(req, resp, options, executionCount)
		cumulativeDelay += delay

		req.Header.Set(retryAttemptHeader, strconv.Itoa(executionCount))

		time.Sleep(delay)

		response, err := pipeline.Next(req, middlewareIndex)
		if err != nil && !IsErrTimeout(err) {
			return response, err
		}

		return middleware.retryRequest(ctx,
			pipeline,
			middlewareIndex,
			options,
			req,
			response,
			executionCount,
			cumulativeDelay,
			err)
	}

	return resp, nil
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

func (middleware RetryHandler) getRetryDelay(req *http.Request,
	resp *http.Response,
	options RetryHandlerOptions,
	executionCount int,
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

	return time.Duration(math.Pow(float64(options.DelaySeconds), float64(executionCount))) * time.Second
}
