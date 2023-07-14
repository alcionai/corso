package graph

import (
	"context"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// constructors
// ---------------------------------------------------------------------------

type Requester interface {
	Request(
		ctx context.Context,
		method, url string,
		body io.Reader,
		headers map[string]string,
	) (*http.Response, error)
}

// NewHTTPWrapper produces a http.Client wrapper that ensures
// calls use all the middleware we expect from the graph api client.
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func NewHTTPWrapper(opts ...Option) *httpWrapper {
	var (
		cc = populateConfig(opts...)
		rt = customTransport{
			n: pipeline{
				middlewares: internalMiddleware(cc),
				transport:   defaultTransport(),
			},
		}
		redirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		hc = &http.Client{
			CheckRedirect: redirect,
			Timeout:       defaultHTTPClientTimeout,
			Transport:     rt,
		}
	)

	cc.apply(hc)

	return &httpWrapper{hc}
}

// NewNoTimeoutHTTPWrapper constructs a http wrapper with no context timeout.
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func NewNoTimeoutHTTPWrapper(opts ...Option) *httpWrapper {
	opts = append(opts, NoTimeout())
	return NewHTTPWrapper(opts...)
}

// ---------------------------------------------------------------------------
// requests
// ---------------------------------------------------------------------------

var streamErrRE = regexp.MustCompile(`stream error: stream ID \d+; .+; received from peer`)

// Request does the provided request.
func (hw httpWrapper) Request(
	ctx context.Context,
	method, url string,
	body io.Reader,
	headers map[string]string,
) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, clues.Wrap(err, "new http request")
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//nolint:lll
	// Decorate the traffic
	// See https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online#how-to-decorate-your-http-traffic
	req.Header.Set("User-Agent", "ISV|Alcion|Corso/"+version.Version)

	var resp *http.Response

	i := 0

	// stream errors from http/2 will fail before we reach
	// client middleware handling, therefore we don't get to
	// make use of the retry middleware.  This external
	// retry wrapper is unsophisticated, but should only
	// retry in the event of a `stream error`, which is not
	// a common expectation.
	for i < 3 {
		ictx := clues.Add(ctx, "request_retry_iter", i)

		resp, err = hw.client.Do(req)
		if err != nil && !streamErrRE.MatchString(err.Error()) {
			return nil, Stack(ictx, err)
		}

		if err == nil {
			break
		}

		logger.Ctx(ictx).Debug("retrying after stream error")
		events.Inc(events.APICall, "streamerror")

		time.Sleep(3 * time.Second)
		i++
	}

	if err != nil {
		return nil, Stack(ctx, err)
	}

	return resp, nil
}

// ---------------------------------------------------------------------------
// constructor internals
// ---------------------------------------------------------------------------

type (
	httpWrapper struct {
		client *http.Client
	}

	customTransport struct {
		n nexter
	}

	pipeline struct {
		transport   http.RoundTripper
		middlewares []khttp.Middleware
	}
)

// RoundTrip kicks off the middleware chain and returns a response
func (ct customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return ct.n.Next(req, 0)
}

// Next moves the request object through middlewares in the pipeline
func (pl pipeline) Next(req *http.Request, idx int) (*http.Response, error) {
	if idx < len(pl.middlewares) {
		return pl.middlewares[idx].Intercept(pl, idx+1, req)
	}

	return pl.transport.RoundTrip(req)
}

func defaultTransport() http.RoundTripper {
	defaultTransport := http.DefaultTransport.(*http.Transport).Clone()
	defaultTransport.ForceAttemptHTTP2 = true

	return defaultTransport
}

func internalMiddleware(cc *clientConfig) []khttp.Middleware {
	mw := []khttp.Middleware{
		&RetryMiddleware{
			MaxRetries: cc.maxRetries,
			Delay:      cc.minDelay,
		},
		khttp.NewRetryHandler(),
		khttp.NewRedirectHandler(),
		&LoggingMiddleware{},
		&throttlingMiddleware{newTimedFence()},
		&RateLimiterMiddleware{},
		&MetricsMiddleware{},
	}

	if len(cc.appendMiddleware) > 0 {
		mw = append(mw, cc.appendMiddleware...)
	}

	return mw
}
