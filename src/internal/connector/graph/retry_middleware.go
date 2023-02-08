package graph

import (
	"context"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	khttp "github.com/microsoft/kiota-http-go"
)

// RetryHandler handles transient HTTP responses and retries the request given the retry options
type RetryHandler struct {
	// default options to use when evaluating the response
	options RetryHandlerOptions
}

type RetryHandlerOptions struct {
	// The maximum number of times a request can be retried
	MaxRetries int
	// The delay in seconds between retries
	DelaySeconds int
}

var retryKeyValue = abs.RequestOptionKey{
	Key: "RetryHandler",
}

// GetKey returns the key value to be used when the option is added to the request context
func (options *RetryHandlerOptions) GetKey() abs.RequestOptionKey {
	return retryKeyValue
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
	if options.MaxRetries < 1 {
		options.MaxRetries = defaultMaxRetries
	}

	if (respErr != nil || middleware.isRetriableErrorCode(req, resp.StatusCode)) &&
		middleware.isRetriableRequest(req) &&
		executionCount < options.MaxRetries &&
		cumulativeDelay < time.Duration(absoluteMaxDelaySeconds)*time.Second {
		executionCount++
		delay := middleware.getRetryDelay(req, resp, options, executionCount)
		cumulativeDelay += delay

		req.Header.Set(retryAttemptHeader, strconv.Itoa(executionCount))

		if req.Body != nil {
			s, ok := req.Body.(io.Seeker)
			if ok {
				_, err := s.Seek(0, io.SeekStart)
				if err != nil {
					return &http.Response{}, err
				}
			}
		}

		time.Sleep(delay)

		response, err := pipeline.Next(req, middlewareIndex)
		if err != nil && !IsErrTimeout(err) {
			return response, err
		}

		return middleware.retryRequest(ctx, pipeline,
			middlewareIndex, options, req, response, executionCount, cumulativeDelay, err)
	}

	return resp, nil
}

func (middleware RetryHandler) isRetriableErrorCode(req *http.Request, code int) bool {
	return code == http.StatusTooManyRequests ||
		code == http.StatusServiceUnavailable ||
		code == http.StatusGatewayTimeout ||
		code == http.StatusInternalServerError
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
	if options.DelaySeconds < 1 {
		options.DelaySeconds = defaultDelaySeconds
	}

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
