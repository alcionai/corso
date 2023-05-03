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
	"golang.org/x/exp/slices"
	"golang.org/x/time/rate"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/logger"
)

type nexter interface {
	Next(req *http.Request, middlewareIndex int) (*http.Response, error)
}

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

func (mw *LoggingMiddleware) Intercept(
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

// RetryMiddleware handles transient HTTP responses and retries the request given the retry options
type RetryMiddleware struct {
	// The maximum number of times a request can be retried
	MaxRetries int
	// The delay in seconds between retries
	Delay time.Duration
}

// Intercept implements the interface and evaluates whether to retry a failed request.
func (mw RetryMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	ctx := req.Context()

	resp, err := pipeline.Next(req, middlewareIndex)
	if err != nil && !IsErrTimeout(err) && !IsErrConnectionReset(err) {
		return resp, stackReq(ctx, req, resp, err)
	}

	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.InitialInterval = mw.Delay
	exponentialBackOff.Reset()

	resp, err = mw.retryRequest(
		ctx,
		pipeline,
		middlewareIndex,
		req,
		resp,
		0,
		0,
		exponentialBackOff,
		err)
	if err != nil {
		return nil, stackReq(ctx, req, resp, err)
	}

	return resp, nil
}

func (mw RetryMiddleware) retryRequest(
	ctx context.Context,
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
	resp *http.Response,
	executionCount int,
	cumulativeDelay time.Duration,
	exponentialBackoff *backoff.ExponentialBackOff,
	priorErr error,
) (*http.Response, error) {
	status := "unknown_resp_status"
	statusCode := -1

	if resp != nil {
		status = resp.Status
		statusCode = resp.StatusCode
	}

	ctx = clues.Add(
		ctx,
		"prev_resp_status", status,
		"retry_count", executionCount)

	// only retry under certain conditions:
	// 1, there was an error.  2, the resp and/or status code match retriable conditions.
	// 3, the request is retriable.
	// 4, we haven't hit our max retries already.
	if (priorErr != nil || mw.isRetriableRespCode(ctx, resp, statusCode)) &&
		mw.isRetriableRequest(req) &&
		executionCount < mw.MaxRetries {
		executionCount++

		delay := mw.getRetryDelay(req, resp, exponentialBackoff)

		cumulativeDelay += delay

		req.Header.Set(retryAttemptHeader, strconv.Itoa(executionCount))

		timer := time.NewTimer(delay)

		select {
		case <-ctx.Done():
			// Don't retry if the context is marked as done, it will just error out
			// when we attempt to send the retry anyway.
			return resp, clues.Stack(ctx.Err()).WithClues(ctx)

		case <-timer.C:
		}

		nextResp, err := pipeline.Next(req, middlewareIndex)
		if err != nil && !IsErrTimeout(err) && !IsErrConnectionReset(err) {
			return nextResp, stackReq(ctx, req, nextResp, err)
		}

		return mw.retryRequest(ctx,
			pipeline,
			middlewareIndex,
			req,
			nextResp,
			executionCount,
			cumulativeDelay,
			exponentialBackoff,
			err)
	}

	if priorErr != nil {
		return nil, stackReq(ctx, req, nil, priorErr)
	}

	return resp, nil
}

var retryableRespCodes = []int{
	http.StatusInternalServerError,
	http.StatusServiceUnavailable,
	http.StatusBadGateway,
	http.StatusGatewayTimeout,
}

func (mw RetryMiddleware) isRetriableRespCode(ctx context.Context, resp *http.Response, code int) bool {
	if slices.Contains(retryableRespCodes, code) {
		return true
	}

	// not a status code, but the message itself might indicate a connectivity issue that
	// can be retried independent of the status code.
	return strings.Contains(
		strings.ToLower(getRespDump(ctx, resp, true)),
		strings.ToLower(string(IOErrDuringRead)))
}

func (mw RetryMiddleware) isRetriableRequest(req *http.Request) bool {
	isBodiedMethod := req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH"
	if isBodiedMethod && req.Body != nil {
		return req.ContentLength != -1
	}

	return true
}

func (mw RetryMiddleware) getRetryDelay(
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

func (mw *ThrottleControlMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	QueueRequest(req.Context())
	return pipeline.Next(req, middlewareIndex)
}

// MetricsMiddleware aggregates per-request metrics on the events bus
type MetricsMiddleware struct{}

func (mw *MetricsMiddleware) Intercept(
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
