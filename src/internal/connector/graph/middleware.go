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

// 1 MB
const logMBLimit = 1 * 1048576

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
		req.Context(),
		"method", req.Method,
		"url", LoggableURL(req.URL.String()),
		"request_content_len", req.ContentLength,
		"resp_status", resp.Status,
		"resp_status_code", resp.StatusCode,
		"resp_content_len", resp.ContentLength)

	var (
		log       = logger.Ctx(ctx)
		respClass = resp.StatusCode / 100

		// special cases where we always dump the response body, since the response
		// details might be critical to understanding the response when debugging.
		// * 400-bad-request
		// * 403-forbidden
		logBody = logger.DebugAPIFV ||
			os.Getenv(logGraphRequestsEnvKey) != "" ||
			resp.StatusCode == http.StatusBadRequest ||
			resp.StatusCode == http.StatusForbidden
	)

	// special case: always info-level status 429 logs
	if resp.StatusCode == http.StatusTooManyRequests {
		log.With("response", getRespDump(ctx, resp, logBody)).
			Info("graph api throttling")

		return resp, err
	}

	// Log api calls according to api debugging configurations.
	switch respClass {
	case 2:
		if logBody {
			// only dump the body if it's under a size limit.  We don't want to copy gigs into memory for a log.
			dump := getRespDump(ctx, resp, os.Getenv(log2xxGraphResponseEnvKey) != "" && resp.ContentLength < logMBLimit)
			log.Infow("2xx graph api resp", "response", dump)
		}
	case 3:
		log.With("redirect_location", LoggableURL(resp.Header.Get(locationHeader))).
			With("response", getRespDump(ctx, resp, false)).
			Info("graph api redirect: " + resp.Status)
	default:
		log.With("response", getRespDump(ctx, resp, logBody)).
			Error("graph api error: " + resp.Status)
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

	// Alternate approach
	retriable := IsErrTimeout(err) || IsErrConnectionReset(err) ||
		(resp != nil && (resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5))

	if !retriable {
		return resp, stackReq(ctx, req, resp, err)
	}

	// Existing approach
	// if err != nil && !IsErrTimeout(err) && !IsErrConnectionReset(err) {
	// 	return resp, stackReq(ctx, req, resp, err)
	// }

	// if resp != nil && resp.StatusCode/100 != 4 && resp.StatusCode/100 != 5 {
	// 	return resp, err
	// }

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

	if priorErr != nil {
		return nil, stackReq(ctx, req, nil, priorErr)
	}

	return resp, nil
}

var retryableRespCodes = []int{
	http.StatusInternalServerError,
	http.StatusBadGateway,
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

	events.IncN(xmrui, events.APICall, xmruHeader)

	return resp, err
}
