package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alcionai/clues"
	backoff "github.com/cenkalti/backoff/v4"
	khttp "github.com/microsoft/kiota-http-go"
	"golang.org/x/time/rate"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// Logging
// ---------------------------------------------------------------------------

// LoggingMiddleware can be used to log the http request sent by the graph client
type LoggingMiddleware struct{}

// well-known path names used by graph api calls
// used to un-hide path elements in a pii.SafeURL
var SafeURLPathParams = pii.MapWithPlurals(
	//nolint:misspell
	"alltime",
	"analytics",
	"archive",
	"beta",
	"calendargroup",
	"calendar",
	"calendarview",
	"channel",
	"childfolder",
	"children",
	"clone",
	"column",
	"contactfolder",
	"contact",
	"contenttype",
	"delta",
	"drive",
	"event",
	"group",
	"inbox",
	"instance",
	"invitation",
	"item",
	"joinedteam",
	"label",
	"list",
	"mailfolder",
	"member",
	"message",
	"notification",
	"page",
	"primarychannel",
	"root",
	"security",
	"site",
	"subscription",
	"team",
	"unarchive",
	"user",
	"v1.0")

//	well-known safe query parameters used by graph api calls
//
// used to un-hide query params in a pii.SafeURL
var SafeURLQueryParams = map[string]struct{}{
	"deltatoken":    {},
	"startdatetime": {},
	"enddatetime":   {},
	"$count":        {},
	"$expand":       {},
	"$filter":       {},
	"$select":       {},
	"$top":          {},
}

func LoggableURL(url string) pii.SafeURL {
	return pii.SafeURL{
		URL:           url,
		SafePathElems: SafeURLPathParams,
		SafeQueryKeys: SafeURLQueryParams,
	}
}

func (handler *LoggingMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	ctx := clues.Add(
		req.Context(),
		"method", req.Method,
		"url", LoggableURL(req.URL.String()),
		"request_len", req.ContentLength)

	// call the next middleware
	resp, err := pipeline.Next(req, middlewareIndex)

	if strings.Contains(req.URL.String(), "users//") {
		logger.Ctx(ctx).Error("malformed request url: missing resource")
	}

	if resp == nil {
		return resp, err
	}

	ctx = clues.Add(ctx, "status", resp.Status, "statusCode", resp.StatusCode)
	log := logger.Ctx(ctx)

	// Return immediately if the response is good (2xx).
	// If api logging is toggled, log a body-less dump of the request/resp.
	if (resp.StatusCode / 100) == 2 {
		if logger.DebugAPIFV || os.Getenv(log2xxGraphRequestsEnvKey) != "" {
			log.Debugw("2xx graph api resp", "response", getRespDump(ctx, resp, os.Getenv(log2xxGraphResponseEnvKey) != ""))
		}

		return resp, err
	}

	// Log errors according to api debugging configurations.
	// When debugging is toggled, every non-2xx is recorded with a response dump.
	// Otherwise, throttling cases and other non-2xx responses are logged
	// with a slimmer reference for telemetry/supportability purposes.
	if logger.DebugAPIFV || os.Getenv(logGraphRequestsEnvKey) != "" {
		log.Errorw("non-2xx graph api response", "response", getRespDump(ctx, resp, true))
		return resp, err
	}

	msg := fmt.Sprintf("graph api error: %s", resp.Status)

	// special case for supportability: log all throttling cases.
	if resp.StatusCode == http.StatusTooManyRequests {
		log = log.With(
			"limit", resp.Header.Get(rateLimitHeader),
			"remaining", resp.Header.Get(rateRemainingHeader),
			"reset", resp.Header.Get(rateResetHeader),
			"retry-after", resp.Header.Get(retryAfterHeader))
	} else if resp.StatusCode/100 == 4 || resp.StatusCode == http.StatusServiceUnavailable {
		log = log.With("response", getRespDump(ctx, resp, true))
	}

	log.Info(msg)

	return resp, err
}

func getRespDump(ctx context.Context, resp *http.Response, getBody bool) string {
	respDump, err := httputil.DumpResponse(resp, getBody)
	if err != nil {
		logger.CtxErr(ctx, err).Error("dumping http response")
	}

	return string(respDump)
}

// ---------------------------------------------------------------------------
// Retry & Backoff
// ---------------------------------------------------------------------------

// RetryHandler handles transient HTTP responses and retries the request given the retry options
type RetryHandler struct {
	// The maximum number of times a request can be retried
	MaxRetries int
	// The delay in seconds between retries
	Delay time.Duration
}

func (middleware RetryHandler) retryRequest(
	ctx context.Context,
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
	resp *http.Response,
	executionCount int,
	cumulativeDelay time.Duration,
	exponentialBackoff *backoff.ExponentialBackOff,
	respErr error,
) (*http.Response, error) {
	if (respErr != nil || middleware.isRetriableErrorCode(req, resp.StatusCode)) &&
		middleware.isRetriableRequest(req) &&
		executionCount < middleware.MaxRetries {
		executionCount++

		delay := middleware.getRetryDelay(req, resp, exponentialBackoff)

		cumulativeDelay += delay

		req.Header.Set(retryAttemptHeader, strconv.Itoa(executionCount))

		timer := time.NewTimer(delay)

		select {
		case <-ctx.Done():
			// Don't retry if the context is marked as done, it will just error out
			// when we attempt to send the retry anyway.
			return resp, ctx.Err()

		// Will exit switch-block so the remainder of the code doesn't need to be
		// indented.
		case <-timer.C:
		}

		response, err := pipeline.Next(req, middlewareIndex)
		if err != nil && !IsErrTimeout(err) && !IsErrConnectionReset(err) {
			return response, Stack(ctx, err).With("retry_count", executionCount)
		}

		return middleware.retryRequest(ctx,
			pipeline,
			middlewareIndex,
			req,
			response,
			executionCount,
			cumulativeDelay,
			exponentialBackoff,
			err)
	}

	if respErr != nil {
		return nil, Stack(ctx, respErr).With("retry_count", executionCount)
	}

	return resp, nil
}

func (middleware RetryHandler) isRetriableErrorCode(req *http.Request, code int) bool {
	return code == http.StatusInternalServerError || code == http.StatusServiceUnavailable
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

// Intercept implements the interface and evaluates whether to retry a failed request.
func (middleware RetryHandler) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	ctx := req.Context()

	response, err := pipeline.Next(req, middlewareIndex)
	if err != nil && !IsErrTimeout(err) {
		return response, Stack(ctx, err)
	}

	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.InitialInterval = middleware.Delay
	exponentialBackOff.Reset()

	response, err = middleware.retryRequest(
		ctx,
		pipeline,
		middlewareIndex,
		req,
		response,
		0,
		0,
		exponentialBackOff,
		err)
	if err != nil {
		return nil, Stack(ctx, err)
	}

	return response, nil
}

// We're trying to keep calls below the 10k-per-10-minute threshold.
// 15 tokens every second nets 900 per minute.  That's 9000 every 10 minutes,
// which is a bit below the mark.
// But suppose we have a minute-long dry spell followed by a 10 minute tsunami.
// We'll have built up 900 tokens in reserve, so the first 900 calls go through
// immediately.  Over the next 10 minutes, we'll partition out the other calls
// at a rate of 900-per-minute, ending at a total of 9900.  Theoretically, if
// the volume keeps up after that, we'll always stay between 9000 and 9900 out
// of 10k.
const (
	perSecond = 15
	maxCap    = 900
)

// Single, global rate limiter at this time.  Refinements for method (creates,
// versus reads) or service can come later.
var limiter = rate.NewLimiter(perSecond, maxCap)

// QueueRequest will allow the request to occur immediately if we're under the
// 1k-calls-per-minute rate.  Otherwise, the call will wait in a queue until
// the next token set is available.
func QueueRequest(ctx context.Context) {
	if err := limiter.Wait(ctx); err != nil {
		logger.CtxErr(ctx, err).Error("graph middleware waiting on the limiter")
	}
}

// ---------------------------------------------------------------------------
// Rate Limiting
// ---------------------------------------------------------------------------

// ThrottleControlMiddleware is used to ensure we don't overstep 10k-per-10-min
// request limits.
type ThrottleControlMiddleware struct{}

func (handler *ThrottleControlMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	QueueRequest(req.Context())
	return pipeline.Next(req, middlewareIndex)
}

// MetricsMiddleware aggregates per-request metrics on the events bus
type MetricsMiddleware struct{}

func (handler *MetricsMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	var (
		start     = time.Now()
		resp, err = pipeline.Next(req, middlewareIndex)
		status    = "nil-resp"
	)

	if resp != nil {
		status = resp.Status
	}

	events.Inc(events.APICall)
	events.Inc(events.APICall, status)
	events.Since(start, events.APICall)
	events.Since(start, events.APICall, status)

	return resp, err
}
