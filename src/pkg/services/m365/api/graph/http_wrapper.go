package graph

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
	"github.com/pkg/errors"
	"golang.org/x/net/http2"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// HTTP wrapper config
// ---------------------------------------------------------------------------

const (
	httpWrapperRetryDelay = 3 * time.Second
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
func NewHTTPWrapper(
	counter *count.Bus,
	opts ...Option,
) *httpWrapper {
	var (
		cc = populateConfig(opts...)
		rt = customTransport{
			n: pipeline{
				middlewares: internalMiddleware(cc, counter),
				transport:   defaultTransport(),
			},
		}
		redirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		hc = &http.Client{
			CheckRedirect: redirect,
			Transport:     rt,
		}
	)

	cc.apply(hc)

	return &httpWrapper{
		client:     hc,
		config:     cc,
		retryDelay: httpWrapperRetryDelay,
	}
}

// NewNoTimeoutHTTPWrapper constructs a http wrapper with no context timeout.
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func NewNoTimeoutHTTPWrapper(
	counter *count.Bus,
	opts ...Option,
) *httpWrapper {
	opts = append(opts, NoTimeout())
	return NewHTTPWrapper(counter, opts...)
}

// ---------------------------------------------------------------------------
// requests
// ---------------------------------------------------------------------------

// Request does the provided request.
func (hw httpWrapper) Request(
	ctx context.Context,
	method, url string,
	body io.Reader,
	headers map[string]string,
) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
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

	if hw.config.requesterAuth != nil {
		err := hw.config.requesterAuth.addAuthToHeaders(ctx, url, req.Header)
		if err != nil {
			return nil, clues.Wrap(err, "setting request auth headers")
		}
	}

	retriedErrors := []string{}

	var e error

	// stream errors from http/2 will fail before we reach
	// client middleware handling, therefore we don't get to
	// make use of the retry middleware.  This external
	// retry wrapper is unsophisticated, but should only
	// retry in the event of a `stream error`, which is not
	// a common expectation.
	for i := 0; i < hw.config.maxConnectionRetries+1; i++ {
		if i > 0 {
			time.Sleep(hw.retryDelay)
		}

		ictx := clues.Add(
			ctx,
			"request_retry_iter", i,
			"request_start_time", dttm.Now())

		resp, err := hw.client.Do(req)
		if err == nil {
			logResp(ictx, resp)
			return resp, nil
		}

		err = stackWithCoreErr(ictx, err, 1)
		e = err

		var http2StreamErr http2.StreamError
		if !errors.As(err, &http2StreamErr) {
			// exit most errors without retry
			break
		}

		logger.Ctx(ictx).Debug("http2 stream error")
		events.Inc(events.APICall, "streamerror")

		retriedErrors = append(retriedErrors, err.Error())
	}

	e = clues.Stack(e).
		With(
			"retried_errors", retriedErrors,
			"request_end_time", dttm.Now()).
		WithTrace(1).
		OrNil()

	// no chance of a non-error return here.
	// we handle that inside the loop.
	return nil, e
}

// ---------------------------------------------------------------------------
// constructor internals
// ---------------------------------------------------------------------------

type (
	httpWrapper struct {
		client     *http.Client
		config     *clientConfig
		retryDelay time.Duration
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

func internalMiddleware(
	cc *clientConfig,
	counter *count.Bus,
) []khttp.Middleware {
	throttler := &throttlingMiddleware{
		tf:      newTimedFence(),
		counter: counter,
	}

	retryOptions := khttp.RetryHandlerOptions{
		ShouldRetry: func(
			delay time.Duration,
			executionCount int,
			request *http.Request,
			response *http.Response,
		) bool {
			return true
		},
		MaxRetries:   cc.maxRetries,
		DelaySeconds: int(cc.minDelay.Seconds()),
	}

	mw := []khttp.Middleware{
		&RetryMiddleware{
			MaxRetries: cc.maxRetries,
			Delay:      cc.minDelay,
		},
		// We use default kiota retry handler for 503 and 504 errors
		khttp.NewRetryHandlerWithOptions(retryOptions),
		khttp.NewRedirectHandler(),
		&LoggingMiddleware{},
		throttler,
		&RateLimiterMiddleware{},
		&MetricsMiddleware{
			counter: counter,
		},
	}

	if len(cc.appendMiddleware) > 0 {
		mw = append(mw, cc.appendMiddleware...)
	}

	return mw
}
