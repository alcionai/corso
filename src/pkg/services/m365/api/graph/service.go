package graph

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kauth "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	locationHeader      = "Location"
	rateLimitHeader     = "RateLimit-Limit"
	rateRemainingHeader = "RateLimit-Remaining"
	rateResetHeader     = "RateLimit-Reset"
	retryAfterHeader    = "Retry-After"
	retryAttemptHeader  = "Retry-Attempt"
)

type QueryParams struct {
	Category          path.CategoryType
	ProtectedResource idname.Provider
	TenantID          string
}

// ---------------------------------------------------------------------------
// Interfaces
// ---------------------------------------------------------------------------

type Servicer interface {
	// Client() returns msgraph Service client that can be used to process and execute
	// the majority of the queries to the M365 Backstore
	Client() *msgraphsdkgo.GraphServiceClient
	// Adapter() returns GraphRequest adapter used to process large requests, create batches
	// and page iterators
	Adapter() abstractions.RequestAdapter
}

// ---------------------------------------------------------------------------
// Service Handler
// ---------------------------------------------------------------------------

var _ Servicer = &Service{}

type Service struct {
	adapter abstractions.RequestAdapter
	client  *msgraphsdkgo.GraphServiceClient
}

func NewService(adapter abstractions.RequestAdapter) *Service {
	return &Service{
		adapter: adapter,
		client:  msgraphsdkgo.NewGraphServiceClient(adapter),
	}
}

func (s Service) Adapter() abstractions.RequestAdapter {
	return s.adapter
}

func (s Service) Client() *msgraphsdkgo.GraphServiceClient {
	return s.client
}

// Seraialize writes an M365 parsable object into a byte array using the built-in
// application/json writer within the adapter.
func (s Service) Serialize(object serialization.Parsable) ([]byte, error) {
	writer, err := s.adapter.GetSerializationWriterFactory().GetSerializationWriter("application/json")
	if err != nil || writer == nil {
		return nil, clues.Wrap(err, "creating json serialization writer")
	}

	err = writer.WriteObjectValue("", object)
	if err != nil {
		return nil, clues.Wrap(err, "serializing object")
	}

	return writer.GetSerializedContent()
}

// ---------------------------------------------------------------------------
// Adapter
// ---------------------------------------------------------------------------

// CreateAdapter uses provided credentials to log into M365 using Kiota Azure Library
// with Azure identity package. An adapter object is a necessary to component
// to create a graph api client connection.
func CreateAdapter(
	tenant, client, secret string,
	counter *count.Bus,
	opts ...Option,
) (abstractions.RequestAdapter, error) {
	auth, err := GetAuth(tenant, client, secret)
	if err != nil {
		return nil, err
	}

	httpClient, cc := KiotaHTTPClient(counter, opts...)

	adpt, err := msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return wrapAdapter(adpt, cc), nil
}

func GetAuth(tenant string, client string, secret string) (*kauth.AzureIdentityAuthenticationProvider, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := azidentity.NewClientSecretCredential(tenant, client, secret, nil)
	if err != nil {
		return nil, clues.Wrap(err, "creating m365 client identity")
	}

	auth, err := kauth.NewAzureIdentityAuthenticationProviderWithScopes(
		cred,
		[]string{"https://graph.microsoft.com/.default"})
	if err != nil {
		return nil, clues.Wrap(err, "creating azure authentication")
	}

	return auth, nil
}

// KiotaHTTPClient creates a httpClient with middlewares and timeout configured
// for use in the graph adapter.
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func KiotaHTTPClient(
	counter *count.Bus,
	opts ...Option,
) (*http.Client, *clientConfig) {
	var (
		clientOptions = msgraphsdkgo.GetDefaultClientOptions()
		cc            = populateConfig(opts...)
		middlewares   = kiotaMiddlewares(&clientOptions, cc, counter)
		httpClient    = msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	)

	cc.apply(httpClient)

	return httpClient, cc
}

// ---------------------------------------------------------------------------
// HTTP Client Config
// ---------------------------------------------------------------------------

const (
	defaultDelay             = 3 * time.Second
	defaultHTTPClientTimeout = 1 * time.Hour

	// Default retry count for retry middlewares
	defaultMaxRetries = 3
	// Retry count for graph adapter
	//
	// Bumping retries to 6 since we have noticed that auth token expiry errors
	// may continue to persist even after 3 retries.
	adapterMaxRetries = 6

	// FIXME: This should ideally be 0, but if we set to 0, graph
	// client with automatically set the context timeout to 0 as
	// well which will make the client unusable.
	// https://github.com/microsoft/kiota-http-go/pull/71
	defaultNoTimeout = 48 * time.Hour
)

type clientConfig struct {
	timeout time.Duration
	// MaxConnectionRetries is the number of connection-level retries that
	// attempt to re-run the request due to a broken or closed connection.
	maxConnectionRetries int
	// MaxRetries is the number of middleware retires attempted
	// before returning with failure
	maxRetries int
	// The minimum delay in seconds between retries
	minDelay time.Duration

	appendMiddleware []khttp.Middleware
}

type Option func(*clientConfig)

// populate constructs a clientConfig according to the provided options.
func populateConfig(opts ...Option) *clientConfig {
	cc := clientConfig{
		maxConnectionRetries: adapterMaxRetries,
		maxRetries:           defaultMaxRetries,
		minDelay:             defaultDelay,
		timeout:              defaultHTTPClientTimeout,
	}

	for _, opt := range opts {
		opt(&cc)
	}

	return &cc
}

// apply updates the http.Client with the expected options.
func (c *clientConfig) apply(hc *http.Client) {
	hc.Timeout = c.timeout
}

// NoTimeout sets the httpClient.Timeout to 48 hours (eg: unlimited).
// The resulting client isn't suitable for most queries, due to the
// capacity for a call to persist forever.  This configuration should
// only be used when downloading very large files.
func NoTimeout() Option {
	return func(c *clientConfig) {
		c.timeout = defaultNoTimeout
	}
}

func Timeout(timeout time.Duration) Option {
	return func(c *clientConfig) {
		c.timeout = timeout
	}
}

func MaxRetries(max int) Option {
	return func(c *clientConfig) {
		if max < 0 {
			max = 0
		} else if max > 5 {
			max = 5
		}

		c.maxRetries = max
	}
}

func MinimumBackoff(min time.Duration) Option {
	return func(c *clientConfig) {
		if min < 100*time.Millisecond {
			min = 100 * time.Millisecond
		} else if min > 5*time.Second {
			min = 5 * time.Second
		}

		c.minDelay = min
	}
}

func appendMiddleware(mw ...khttp.Middleware) Option {
	return func(c *clientConfig) {
		if len(mw) > 0 {
			c.appendMiddleware = mw
		}
	}
}

func MaxConnectionRetries(max int) Option {
	return func(c *clientConfig) {
		if max < 0 {
			max = 0
		} else if max > 5 {
			max = 5
		}

		c.maxConnectionRetries = max
	}
}

// ---------------------------------------------------------------------------
// Middleware Control
// ---------------------------------------------------------------------------

// kiotaMiddlewares creates a default slice of middleware for the Graph Client.
func kiotaMiddlewares(
	options *msgraphgocore.GraphClientOptions,
	cc *clientConfig,
	counter *count.Bus,
) []khttp.Middleware {
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
		msgraphgocore.NewGraphTelemetryHandler(options),
		&RetryMiddleware{
			MaxRetries: cc.maxRetries,
			Delay:      cc.minDelay,
		},
		// We use default kiota retry handler for 503 and 504 errors
		khttp.NewRetryHandlerWithOptions(retryOptions),
		khttp.NewRedirectHandler(),
		khttp.NewCompressionHandler(),
		khttp.NewParametersNameDecodingHandler(),
		khttp.NewUserAgentHandler(),
		&LoggingMiddleware{},
	}

	// Optionally add concurrency limiter middleware if it has been initialized.
	if concurrencyLimitMiddlewareSingleton != nil {
		mw = append(mw, concurrencyLimitMiddlewareSingleton)
	}

	throttler := &throttlingMiddleware{
		tf:      newTimedFence(),
		counter: counter,
	}

	mw = append(
		mw,
		throttler,
		&RateLimiterMiddleware{},
		&MetricsMiddleware{
			counter: counter,
		})

	if len(cc.appendMiddleware) > 0 {
		mw = append(mw, cc.appendMiddleware...)
	}

	return mw
}

// ---------------------------------------------------------------------------
// Graph Api Adapter Wrapper
// ---------------------------------------------------------------------------

var _ abstractions.RequestAdapter = &adapterWrap{}

const (
	// Delay between retry attempts
	adapterRetryDelay = 3 * time.Second
)

// adapterWrap takes a GraphRequestAdapter and replaces the Send() function to
// act as a middleware for all http calls.  Certain error conditions never reach
// the the client middleware layer, and therefore miss out on logging and retries.
// By hijacking the Send() call, we can ensure three basic needs:
// 1. Panics generated by the graph client are caught instead of crashing corso.
// 2. Http and Http2 connection closures are retried.
// 3. Error and debug conditions are logged.
type adapterWrap struct {
	abstractions.RequestAdapter
	config     *clientConfig
	retryDelay time.Duration
}

func wrapAdapter(gra *msgraphsdkgo.GraphRequestAdapter, cc *clientConfig) *adapterWrap {
	return &adapterWrap{
		RequestAdapter: gra,
		config:         cc,
		retryDelay:     adapterRetryDelay,
	}
}

// Graph may abruptly close connections, which we should retry.
var connectionEnded = filters.In([]string{
	"connection reset by peer",
	"client connection force closed",
	"read: connection timed out",
})

func (aw *adapterWrap) Send(
	ctx context.Context,
	requestInfo *abstractions.RequestInformation,
	constructor serialization.ParsableFactory,
	errorMappings abstractions.ErrorMappings,
) (sp serialization.Parsable, e error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "graph adapter request"); crErr != nil {
			e = Stack(ctx, crErr)
		}
	}()

	retriedErrors := []string{}

	// This external retry wrapper is unsophisticated, but should
	// only retry under certain circumstances
	// 1. stream errors from http/2, which will fail before we reach
	// client middleware handling.
	// 2. jwt token invalidation, which requires a re-auth that's handled
	// in the Send() call, before reaching client middleware.
	for i := 0; i < aw.config.maxConnectionRetries+1; i++ {
		if i > 0 {
			time.Sleep(aw.retryDelay)
		}

		ictx := clues.Add(
			ctx,
			"request_retry_iter", i,
			"request_start_time", dttm.Now())

		resp, err := aw.RequestAdapter.Send(ictx, requestInfo, constructor, errorMappings)
		if err == nil {
			return resp, nil
		}

		err = stackWithCoreErr(ictx, err, 1)
		e = err

		if IsErrConnectionReset(err) ||
			connectionEnded.Compare(err.Error()) ||
			errors.Is(err, io.ErrUnexpectedEOF) {
			logger.Ctx(ictx).Debug("http connection error")
			events.Inc(events.APICall, "connectionerror")
		} else if errors.Is(err, core.ErrAuthTokenExpired) {
			logger.Ctx(ictx).Debug("bad jwt token")
			events.Inc(events.APICall, "badjwttoken")
		} else if requestInfo.Method.String() == http.MethodGet && IsErrInvalidRequest(err) {
			// Graph may sometimes return a transient 400 response during onedrive
			// and sharepoint backup. This is pending investigation on msft end, retry
			// for now as it's a transient issue. Restrict retries to GET requests only
			// to limit the scope of this fix.
			logger.Ctx(ictx).Debug("invalid request")
			events.Inc(events.APICall, "invalidgetrequest")
		} else if requestInfo.Method.String() == http.MethodGet && errors.Is(err, ErrNotFoundEmptyResp) {
			// We've started seeing 404s with no content being returned for messages
			// message attachments, and events. Attempting to manually fetch the items
			// succeeds. Therefore we want to retry these to see if we can work around
			// the problem.
			logger.Ctx(ictx).Debug("404 with no content")
			events.Inc(events.APICall, "notfoundnocontent")
		} else {
			// exit most errors without retry
			break
		}

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
