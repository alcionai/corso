package graph

import (
	"context"
	"io"
	"net/http"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"

	"github.com/alcionai/corso/src/internal/version"
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

	resp, err := hw.client.Do(req)
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
	return []khttp.Middleware{
		&RetryMiddleware{
			MaxRetries: cc.maxRetries,
			Delay:      cc.minDelay,
		},
		&LoggingMiddleware{},
		&ThrottleControlMiddleware{},
		&MetricsMiddleware{},
	}
}
