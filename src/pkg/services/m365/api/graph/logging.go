package graph

import (
	"context"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// 1 MB
	logMBLimit                = 1 * 1024 * 1024
	logGraphRequestsEnvKey    = "LOG_GRAPH_REQUESTS"
	log2xxGraphRequestsEnvKey = "LOG_2XX_GRAPH_REQUESTS"
	log2xxGraphResponseEnvKey = "LOG_2XX_GRAPH_RESPONSES"
)

// special cases where we always dump the response body, since the response
// details might be critical to understanding the response when debugging.
func shouldLogRespBody(resp *http.Response) bool {
	return logger.DebugAPIFV ||
		os.Getenv(logGraphRequestsEnvKey) != "" ||
		resp.StatusCode > 399
}

func logResp(ctx context.Context, resp *http.Response) {
	var (
		log       = logger.Ctx(ctx)
		respClass = resp.StatusCode / 100

		// special cases where we always dump the response body, since the response
		// details might be critical to understanding the response when debugging.
		logBody = shouldLogRespBody(resp)
	)

	// special case: always info-level status 429 logs
	if resp.StatusCode == http.StatusTooManyRequests {
		log.With("response", getRespDump(ctx, resp, logBody)).
			Info("graph api throttling")
		return
	}

	// Log api calls according to api debugging configurations.
	switch respClass {
	case 2:
		// only log 2xx's if we want the full response body.
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
}

func getRespDump(ctx context.Context, resp *http.Response, getBody bool) string {
	respDump, err := httputil.DumpResponse(resp, getBody)
	if err != nil {
		logger.CtxErr(ctx, err).Error("dumping http response")
	}

	return string(respDump)
}

func getReqCtx(req *http.Request) context.Context {
	if req == nil {
		return context.Background()
	}

	return clues.Add(
		req.Context(),
		"method", req.Method,
		"url", LoggableURL(req.URL.String()),
		"request_content_len", req.ContentLength)
}
