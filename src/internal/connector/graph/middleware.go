package graph

import (
	"context"
	"io"
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
	"github.com/alcionai/corso/src/pkg/path"
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

// 1 MB
const logMBLimit = 1 * 1048576

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

	ctx = clues.Add(
		ctx,
		"status", resp.Status,
		"statusCode", resp.StatusCode,
		"content_len", resp.ContentLength)

	var (
		log       = logger.Ctx(ctx)
		respClass = resp.StatusCode / 100
		logExtra  = logger.DebugAPIFV || os.Getenv(logGraphRequestsEnvKey) != ""
	)

	// special case: always info log 429 responses
	if resp.StatusCode == http.StatusTooManyRequests {
		log.Infow(
			"graph api throttling",
			"limit", resp.Header.Get(rateLimitHeader),
			"remaining", resp.Header.Get(rateRemainingHeader),
			"reset", resp.Header.Get(rateResetHeader),
			"retry-after", resp.Header.Get(retryAfterHeader))

		return resp, err
	}

	// special case: always dump status-400-bad-request
	if resp.StatusCode == http.StatusBadRequest {
		log.With("response", getRespDump(ctx, resp, true)).
			Error("graph api error: " + resp.Status)

		return resp, err
	}

	// Log api calls according to api debugging configurations.
	switch respClass {
	case 2:
		if logExtra {
			// only dump the body if it's under a size limit.  We don't want to copy gigs into memory for a log.
			dump := getRespDump(ctx, resp, os.Getenv(log2xxGraphResponseEnvKey) != "" && resp.ContentLength < logMBLimit)
			log.Infow("2xx graph api resp", "response", dump)
		}
	case 3:
		log.With("redirect_location", LoggableURL(resp.Header.Get(locationHeader)))

		if logExtra {
			log.With("response", getRespDump(ctx, resp, false))
		}

		log.Info("graph api redirect: " + resp.Status)
	default:
		if logExtra {
			log.With("response", getRespDump(ctx, resp, true))
		}

		log.Error("graph api error: " + resp.Status)
	}

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

	if resp != nil && resp.StatusCode/100 != 4 && resp.StatusCode/100 != 5 {
		return resp, err
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

		// we have to reset the original body reader for each retry, or else the graph
		// compressor will produce a 0 length body following an error response such
		// as a 500.
		if req.Body != nil {
			if s, ok := req.Body.(io.Seeker); ok {
				_, err := s.Seek(0, io.SeekStart)
				if err != nil {
					return nil, Wrap(ctx, err, "resetting request body reader")
				}
			}
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

	// prevent the body dump below in case of a 2xx response.
	// There's no reason to check the body on a healthy status.
	if code/100 != 4 && code/100 != 5 {
		return false
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

//nolint:lll
// ---------------------------------------------------------------------------
// Rate limit controls
// https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online
// ---------------------------------------------------------------------------

const (
	// Default goal is to keep calls below the 10k-per-10-minute threshold.
	// 14 tokens every second nets 840 per minute.  That's 8400 every 10 minutes,
	// which is a bit below the mark.
	// But suppose we have a minute-long dry spell followed by a 10 minute tsunami.
	// We'll have built up 750 tokens in reserve, so the first 750 calls go through
	// immediately.  Over the next 10 minutes, we'll partition out the other calls
	// at a rate of 840-per-minute, ending at a total of 9150.  Theoretically, if
	// the volume keeps up after that, we'll always stay between 8400 and 9150 out
	// of 10k.  Worst case scenario, we have an extra minute of padding to allow
	// up to 9990.
	defaultPerSecond = 14  // 14 * 60 = 840
	defaultMaxCap    = 750 // real cap is 10k-per-10-minutes
	// since drive runs on a per-minute, rather than per-10-minute bucket, we have
	// to keep the max cap equal to the per-second cap.  A large maxCap pool (say,
	// 1200, similar to the per-minute cap) would allow us to make a flood of 2400
	// calls in the first minute, putting us over the per-minute limit.  Keeping
	// the cap at the per-second burst means we only dole out a max of 1240 in one
	// minute (20 cap + 1200 per minute + one burst of padding).
	drivePerSecond = 20 // 20 * 60 = 1200
	driveMaxCap    = 20 // real cap is 1250-per-minute
)

var (
	driveLimiter = rate.NewLimiter(drivePerSecond, driveMaxCap)
	// also used as the exchange service limiter
	defaultLimiter = rate.NewLimiter(defaultPerSecond, defaultMaxCap)
)

type LimiterCfg struct {
	Service path.ServiceType
}

type limiterCfgKey string

const limiterCfgCtxKey limiterCfgKey = "corsoGraphRateLimiterCfg"

func ctxLimiter(ctx context.Context) *rate.Limiter {
	lc, ok := extractRateLimiterConfig(ctx)
	if !ok {
		return defaultLimiter
	}

	switch lc.Service {
	case path.OneDriveService, path.SharePointService:
		return driveLimiter
	default:
		return defaultLimiter
	}
}

func BindRateLimiterConfig(ctx context.Context, lc LimiterCfg) context.Context {
	return context.WithValue(ctx, limiterCfgCtxKey, lc)
}

func extractRateLimiterConfig(ctx context.Context) (LimiterCfg, bool) {
	l := ctx.Value(limiterCfgCtxKey)
	if l == nil {
		return LimiterCfg{}, false
	}

	lc, ok := l.(LimiterCfg)

	return lc, ok
}

type limiterConsumptionKey string

const (
	limiterConsumptionCtxKey       limiterConsumptionKey = "corsoGraphRateLimiterConsumption"
	defaultLimiterConsumption                            = 1
	driveDefaultLimiterConsumption                       = 2
	// limit consumption rate for single-item GETs requests,
	// or delta-based multi-item GETs.
	SingleGetOrDeltaLC = 1
	// limit consumption rate for anything permissions related
	PermissionsLC = 5
)

// ConsumeNTokens ensures any calls using this context will consume
// n rate-limiter tokens.  Default is 1, and this value does not need
// to be established in the context to consume the default tokens.
// This should only get used on a per-call basis, to avoid cross-pollination.
func ConsumeNTokens(ctx context.Context, n int) context.Context {
	return context.WithValue(ctx, limiterConsumptionCtxKey, n)
}

func ctxLimiterConsumption(ctx context.Context, defaultConsumption int) int {
	l := ctx.Value(limiterConsumptionCtxKey)
	if l == nil {
		return defaultConsumption
	}

	lc, ok := l.(int)
	if !ok || lc < 1 {
		return defaultConsumption
	}

	return lc
}

// QueueRequest will allow the request to occur immediately if we're under the
// 1k-calls-per-minute rate.  Otherwise, the call will wait in a queue until
// the next token set is available.
func QueueRequest(ctx context.Context) {
	limiter := ctxLimiter(ctx)
	defaultConsumed := defaultLimiterConsumption

	if limiter == driveLimiter {
		defaultConsumed = driveDefaultLimiterConsumption
	}

	consume := ctxLimiterConsumption(ctx, defaultConsumed)

	if err := limiter.WaitN(ctx, consume); err != nil {
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

// ---------------------------------------------------------------------------
// Metrics
// ---------------------------------------------------------------------------

// MetricsMiddleware aggregates per-request metrics on the events bus
type MetricsMiddleware struct{}

const xmruHeader = "x-ms-resource-unit"

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

	// track the graph "resource cost" for each call (if not provided, assume 1)

	// nil-pointer guard
	if len(resp.Header) == 0 {
		resp.Header = http.Header{}
	}

	// from msoft throttling documentation:
	// x-ms-resource-unit - Indicates the resource unit used for this request. Values are positive integer
	xmru := resp.Header.Get(xmruHeader)
	xmrui, e := strconv.Atoi(xmru)

	if len(xmru) == 0 || e != nil {
		xmrui = 1
	}

	events.IncN(xmrui, events.APICall, xmruHeader)

	return resp, err
}
