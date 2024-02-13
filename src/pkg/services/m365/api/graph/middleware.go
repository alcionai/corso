package graph

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alcionai/clues"
	backoff "github.com/cenkalti/backoff/v4"
	khttp "github.com/microsoft/kiota-http-go"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/count"
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
// https://learn.microsoft.com/en-us/graph/api/resources/mailfolder?view=graph-rest-1.0
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
	"clutter",
	"column",
	"conflict",
	"contactfolder",
	"contact",
	"contenttype",
	"conversationhistory",
	"deleteditem",
	"delta",
	"draft",
	"drive",
	"event",
	"group",
	"inbox",
	"instance",
	"invitation",
	"item",
	"joinedteam",
	"junkemail",
	"label",
	"list",
	"localfailure",
	"mailfolder",
	"member",
	"message",
	"msgfolderroot",
	"notification",
	"outbox",
	"page",
	"primarychannel",
	"recoverableitemsdeletion",
	"root",
	"scheduled",
	"searchfolder",
	"security",
	"sentitem",
	"serverfailure",
	"site",
	"subscription",
	"syncissue",
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
	// call the next middleware
	resp, err := pipeline.Next(req, middlewareIndex)
	if resp == nil {
		return resp, err
	}

	ctx := clues.Add(
		getReqCtx(req),
		"resp_status", resp.Status,
		"resp_status_code", resp.StatusCode,
		"resp_content_len", resp.ContentLength)

	logResp(ctx, resp)

	return resp, err
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
	ctx := getReqCtx(req)
	resp, err := pipeline.Next(req, middlewareIndex)

	retriable := IsErrTimeout(err) ||
		IsErrConnectionReset(err) ||
		mw.isRetriableRespCode(ctx, resp)

	if !retriable {
		// Returning a response and an error causes output from either some part of
		// the middleware/graph/golang or the mocking library used for testing.
		// Return one or the other to avoid this.
		err := stackReq(ctx, req, resp, err).OrNil()
		if err != nil {
			return nil, err
		}

		return resp, nil
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

	return resp, stackReq(ctx, req, resp, err).OrNil()
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
	ctx = clues.Add(ctx, "retry_count", executionCount)

	if resp != nil {
		ctx = clues.Add(ctx, "prev_resp_status", resp.Status)
	}

	// only retry if all the following conditions are met:
	// 1, there was a prior error OR the status code match retriable conditions.
	// 3, the request method is retriable.
	// 4, we haven't already hit maximum retries.
	shouldRetry := (priorErr != nil || mw.isRetriableRespCode(ctx, resp)) &&
		mw.isRetriableRequest(req) &&
		executionCount < mw.MaxRetries

	if !shouldRetry {
		return resp, stackReq(ctx, req, resp, priorErr).OrNil()
	}

	executionCount++

	delay := mw.getRetryDelay(req, resp, exponentialBackoff)
	cumulativeDelay += delay

	req.Header.Set(retryAttemptHeader, strconv.Itoa(executionCount))

	timer := time.NewTimer(delay)

	select {
	case <-ctx.Done():
		// Don't retry if the context is marked as done, it will just error out
		// when we attempt to send the retry anyway.
		return resp, clues.StackWC(ctx, ctx.Err())

	case <-timer.C:
	}

	// we have to reset the original body reader for each retry, or else the graph
	// compressor will produce a 0 length body following an error response such
	// as a 500.
	if req.Body != nil {
		if s, ok := req.Body.(io.Seeker); ok {
			if _, err := s.Seek(0, io.SeekStart); err != nil {
				return resp, Wrap(ctx, err, "resetting request body reader")
			}
		} else {
			logger.
				Ctx(getReqCtx(req)).
				Error("body is not an io.Seeker: unable to reset request body")
		}
	}

	nextResp, err := pipeline.Next(req, middlewareIndex)
	if err != nil && !IsErrTimeout(err) && !IsErrConnectionReset(err) {
		return nextResp, stackReq(ctx, req, nextResp, err)
	}

	return mw.retryRequest(
		ctx,
		pipeline,
		middlewareIndex,
		req,
		nextResp,
		executionCount,
		cumulativeDelay,
		exponentialBackoff,
		err)
}

var retryableRespCodes = []int{
	http.StatusInternalServerError,
	http.StatusBadGateway,
}

func (mw RetryMiddleware) isRetriableRespCode(ctx context.Context, resp *http.Response) bool {
	if resp == nil {
		return false
	}

	if slices.Contains(retryableRespCodes, resp.StatusCode) {
		return true
	}

	// prevent the body dump below in case of a 2xx response.
	// There's no reason to check the body on a healthy status.
	if resp.StatusCode/100 != 4 && resp.StatusCode/100 != 5 {
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

	if len(retryAfter) > 0 {
		retryAfterDelay, err := strconv.ParseFloat(retryAfter, 64)
		if err == nil {
			return time.Duration(retryAfterDelay) * time.Second
		}
	} // TODO parse the header if it's a date

	return exponentialBackoff.NextBackOff()
}

// ---------------------------------------------------------------------------
// Metrics
// ---------------------------------------------------------------------------

// MetricsMiddleware aggregates per-request metrics on the events bus
type MetricsMiddleware struct {
	counter *count.Bus
}

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

	if resp == nil {
		return resp, err
	}

	if resp != nil {
		status = resp.Status
	}

	events.Inc(events.APICall)
	events.Inc(events.APICall, status)
	events.Since(start, events.APICall)
	events.Since(start, events.APICall, status)

	// track the graph "resource cost" for each call (if not provided, assume 1)

	// from msoft throttling documentation:
	// x-ms-resource-unit - Indicates the resource unit used for this request. Values are positive integer
	xmru := resp.Header.Get(xmruHeader)
	xmrui, e := strconv.Atoi(xmru)

	if len(xmru) == 0 || e != nil {
		xmrui = 1
	}

	mw.counter.Add(count.APICallTokensConsumed, int64(xmrui))

	events.IncN(xmrui, events.APICall, xmruHeader)

	return resp, err
}
