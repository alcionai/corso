package graph

import (
	"context"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/jwt"
	"github.com/alcionai/corso/src/internal/common/pii"
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

func logResp(ctx context.Context, resp *http.Response, req *http.Request) {
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

	// Log bearer token iat and exp claims if we hit 401s. This is purely for
	// debugging purposes and will be removed in the future.
	if resp.StatusCode == http.StatusUnauthorized {
		errs := []any{"graph api error: " + resp.Status}

		// As per MSFT docs, the token may have a special format and may not always
		// validate as a JWT. Hence log token lifetime in a best effort manner only.
		iat, exp, err := getTokenLifetime(ctx, req)
		if err != nil {
			errs = append(errs, " getting token lifetime: ", err)
		}

		log.With("response", getRespDump(ctx, resp, logBody)).
			With("token issued at", iat, "token expires at", exp).
			Error(errs...)

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

	var logURL pii.SafeURL

	if req.URL != nil {
		logURL = LoggableURL(req.URL.String())
	}

	return clues.AddTraceName(
		req.Context(),
		"graph-http-middleware",
		"method", req.Method,
		"url", logURL,
		"request_content_len", req.ContentLength)
}

// GetTokenLifetime extracts the JWT token embedded in the request and returns
// the token's issue and expiration times. The token is expected to be in the
// "Authorization" header, with a "Bearer " prefix.  If the token is not present
// or is malformed, an error is returned.
func getTokenLifetime(
	ctx context.Context,
	req *http.Request,
) (time.Time, time.Time, error) {
	if req == nil {
		return time.Time{}, time.Time{}, clues.New("nil request")
	}

	rawToken := req.Header.Get("Authorization")

	// Strip the "Bearer " prefix from the token. This prefix is guaranteed to be
	// present as per msft docs. But even if it's not, the jwt lib will handle
	// malformed tokens gracefully and return an error.
	rawToken = strings.TrimPrefix(rawToken, "Bearer ")
	iat, exp, err := jwt.GetJWTLifetime(ctx, rawToken)

	return iat, exp, clues.Stack(err).OrNil()
}
