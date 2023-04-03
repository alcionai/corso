package graph

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alcionai/clues"
	backoff "github.com/cenkalti/backoff/v4"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kauth "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"golang.org/x/time/rate"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	logGraphRequestsEnvKey    = "LOG_GRAPH_REQUESTS"
	log2xxGraphRequestsEnvKey = "LOG_2XX_GRAPH_REQUESTS"
	retryAttemptHeader        = "Retry-Attempt"
	retryAfterHeader          = "Retry-After"
	defaultMaxRetries         = 3
	defaultDelay              = 3 * time.Second
	rateLimitHeader           = "RateLimit-Limit"
	rateRemainingHeader       = "RateLimit-Remaining"
	rateResetHeader           = "RateLimit-Reset"
	defaultHTTPClientTimeout  = 1 * time.Hour
)

// AllMetadataFileNames produces the standard set of filenames used to store graph
// metadata such as delta tokens and folderID->path references.
func AllMetadataFileNames() []string {
	return []string{DeltaURLsFileName, PreviousPathFileName}
}

type QueryParams struct {
	Category      path.CategoryType
	ResourceOwner string
	Credentials   account.M365Config
}

// ---------------------------------------------------------------------------
// Service Handler
// ---------------------------------------------------------------------------

var _ Servicer = &Service{}

type Service struct {
	adapter *msgraphsdkgo.GraphRequestAdapter
	client  *msgraphsdkgo.GraphServiceClient
}

func NewService(adapter *msgraphsdkgo.GraphRequestAdapter) *Service {
	return &Service{
		adapter: adapter,
		client:  msgraphsdkgo.NewGraphServiceClient(adapter),
	}
}

func (s Service) Adapter() *msgraphsdkgo.GraphRequestAdapter {
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
// Helpers
// ---------------------------------------------------------------------------

type clientConfig struct {
	noTimeout bool
	// MaxRetries before failure
	maxRetries int
	// The minimum delay in seconds between retries
	minDelay           time.Duration
	overrideRetryCount bool
}

type option func(*clientConfig)

// populate constructs a clientConfig according to the provided options.
func (c *clientConfig) populate(opts ...option) *clientConfig {
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// apply updates the http.Client with the expected options.
func (c *clientConfig) applyMiddlewareConfig() (retry int, delay time.Duration) {
	retry = defaultMaxRetries
	if c.overrideRetryCount {
		retry = c.maxRetries
	}

	delay = defaultDelay
	if c.minDelay > 0 {
		delay = c.minDelay
	}

	return
}

// apply updates the http.Client with the expected options.
func (c *clientConfig) apply(hc *http.Client) {
	if c.noTimeout {
		// FIXME: This should ideally be 0, but if we set to 0, graph
		// client with automatically set the context timeout to 0 as
		// well which will make the client unusable.
		// https://github.com/microsoft/kiota-http-go/pull/71
		hc.Timeout = 48 * time.Hour
	}
}

// NoTimeout sets the httpClient.Timeout to 0 (unlimited).
// The resulting client isn't suitable for most queries, due to the
// capacity for a call to persist forever.  This configuration should
// only be used when downloading very large files.
func NoTimeout() option {
	return func(c *clientConfig) {
		c.noTimeout = true
	}
}

func MaxRetries(max int) option {
	return func(c *clientConfig) {
		c.overrideRetryCount = true
		c.maxRetries = max
	}
}

func MinimumBackoff(dur time.Duration) option {
	return func(c *clientConfig) {
		c.minDelay = dur
	}
}

// CreateAdapter uses provided credentials to log into M365 using Kiota Azure Library
// with Azure identity package. An adapter object is a necessary to component
// to create  *msgraphsdk.GraphServiceClient
func CreateAdapter(
	tenant, client, secret string,
	opts ...option,
) (*msgraphsdkgo.GraphRequestAdapter, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := azidentity.NewClientSecretCredential(tenant, client, secret, nil)
	if err != nil {
		return nil, clues.Wrap(err, "creating m365 client identity")
	}

	auth, err := kauth.NewAzureIdentityAuthenticationProviderWithScopes(
		cred,
		[]string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		return nil, clues.Wrap(err, "creating azure authentication")
	}

	httpClient := HTTPClient(opts...)

	return msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
}

// HTTPClient creates the httpClient with middlewares and timeout configured
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func HTTPClient(opts ...option) *http.Client {
	clientOptions := msgraphsdkgo.GetDefaultClientOptions()
	clientconfig := (&clientConfig{}).populate(opts...)
	noOfRetries, minRetryDelay := clientconfig.applyMiddlewareConfig()
	middlewares := GetKiotaMiddlewares(&clientOptions, noOfRetries, minRetryDelay)
	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = defaultHTTPClientTimeout

	clientconfig.apply(httpClient)

	return httpClient
}

// GetDefaultMiddlewares creates a new default set of middlewares for the Kiota request adapter
func GetMiddlewares(maxRetry int, delay time.Duration) []khttp.Middleware {
	return []khttp.Middleware{
		&RetryHandler{
			// The maximum number of times a request can be retried
			MaxRetries: maxRetry,
			// The delay in seconds between retries
			Delay: delay,
		},
		khttp.NewRetryHandler(),
		khttp.NewRedirectHandler(),
		khttp.NewCompressionHandler(),
		khttp.NewParametersNameDecodingHandler(),
		khttp.NewUserAgentHandler(),
		&LoggingMiddleware{},
		&ThrottleControlMiddleware{},
		&MetricsMiddleware{},
	}
}

// GetKiotaMiddlewares creates a default slice of middleware for the Graph Client.
func GetKiotaMiddlewares(
	options *msgraphgocore.GraphClientOptions,
	maxRetry int,
	minDelay time.Duration,
) []khttp.Middleware {
	kiotaMiddlewares := GetMiddlewares(maxRetry, minDelay)
	graphMiddlewares := []khttp.Middleware{
		msgraphgocore.NewGraphTelemetryHandler(options),
	}
	graphMiddlewaresLen := len(graphMiddlewares)
	resultMiddlewares := make([]khttp.Middleware, len(kiotaMiddlewares)+graphMiddlewaresLen)
	copy(resultMiddlewares, graphMiddlewares)
	copy(resultMiddlewares[graphMiddlewaresLen:], kiotaMiddlewares)

	return resultMiddlewares
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
	Adapter() *msgraphsdkgo.GraphRequestAdapter
}

// ---------------------------------------------------------------------------
// Client Middleware
// ---------------------------------------------------------------------------

// LoggingMiddleware can be used to log the http request sent by the graph client
type LoggingMiddleware struct{}

// well-known path names used by graph api calls
// used to un-hide path elements in a pii.SafeURL
var safePathParams = pii.MapWithPlurals(
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
var safeQueryParams = map[string]struct{}{
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
		SafePathElems: safePathParams,
		SafeQueryKeys: safeQueryParams,
	}
}

func (handler *LoggingMiddleware) Intercept(
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
		if logger.DebugAPI || os.Getenv(log2xxGraphRequestsEnvKey) != "" {
			log.Debugw("2xx graph api resp", "response", getRespDump(ctx, resp, false))
		}

		return resp, err
	}

	// Log errors according to api debugging configurations.
	// When debugging is toggled, every non-2xx is recorded with a response dump.
	// Otherwise, throttling cases and other non-2xx responses are logged
	// with a slimmer reference for telemetry/supportability purposes.
	if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
		log.Errorw("non-2xx graph api response", "response", getRespDump(ctx, resp, true))
		return resp, err
	}

	msg := fmt.Sprintf("graph api error: %s", resp.Status)

	// special case for supportability: log all throttling cases.
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
		log = log.With(
			"limit", resp.Header.Get(rateLimitHeader),
			"remaining", resp.Header.Get(rateRemainingHeader),
			"reset", resp.Header.Get(rateResetHeader),
			"retry-after", resp.Header.Get(retryAfterHeader))
	} else if resp.StatusCode/100 == 4 {
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

// Intercept implements the interface and evaluates whether to retry a failed request.
func (middleware RetryHandler) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	ctx := req.Context()

	response, err := pipeline.Next(req, middlewareIndex)
	if err != nil && !IsErrTimeout(err) {
		return response, Stack(ctx, err)
	}

	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.InitialInterval = middleware.Delay
	exponentialBackOff.Reset()

	response, err = middleware.retryRequest(
		ctx,
		pipeline,
		middlewareIndex,
		req,
		response,
		0,
		0,
		exponentialBackOff,
		err)
	if err != nil {
		return nil, Stack(ctx, err)
	}

	return response, nil
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

// ThrottleControlMiddleware is used to ensure we don't overstep 10k-per-10-min
// request limits.
type ThrottleControlMiddleware struct{}

func (handler *ThrottleControlMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	QueueRequest(req.Context())
	return pipeline.Next(req, middlewareIndex)
}

// MetricsMiddleware aggregates per-request metrics on the events bus
type MetricsMiddleware struct{}

func (handler *MetricsMiddleware) Intercept(
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
