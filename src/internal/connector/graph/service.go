package graph

import (
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	backoff "github.com/cenkalti/backoff/v4"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	logGraphRequestsEnvKey  = "LOG_GRAPH_REQUESTS"
	numberOfRetries         = 3
	retryAttemptHeader      = "Retry-Attempt"
	retryAfterHeader        = "Retry-After"
	defaultMaxRetries       = 3
	defaultDelay            = 3 * time.Second
	absoluteMaxDelaySeconds = 180
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
	adapter *msgraphsdk.GraphRequestAdapter
	client  *msgraphsdk.GraphServiceClient
}

func NewService(adapter *msgraphsdk.GraphRequestAdapter) *Service {
	return &Service{
		adapter: adapter,
		client:  msgraphsdk.NewGraphServiceClient(adapter),
	}
}

func (s Service) Adapter() *msgraphsdk.GraphRequestAdapter {
	return s.adapter
}

func (s Service) Client() *msgraphsdk.GraphServiceClient {
	return s.client
}

// Seraialize writes an M365 parsable object into a byte array using the built-in
// application/json writer within the adapter.
func (s Service) Serialize(object serialization.Parsable) ([]byte, error) {
	writer, err := s.adapter.GetSerializationWriterFactory().GetSerializationWriter("application/json")
	if err != nil || writer == nil {
		return nil, errors.Wrap(err, "creating json serialization writer")
	}

	err = writer.WriteObjectValue("", object)
	if err != nil {
		return nil, errors.Wrap(err, "serializing object")
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
		hc.Timeout = 0
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
func CreateAdapter(tenant, client, secret string, opts ...option) (*msgraphsdk.GraphRequestAdapter, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := azidentity.NewClientSecretCredential(tenant, client, secret, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating m365 client identity")
	}

	auth, err := ka.NewAzureIdentityAuthenticationProviderWithScopes(
		cred,
		[]string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating azure authentication")
	}

	httpClient := HTTPClient(opts...)

	return msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
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
	clientOptions := msgraphsdk.GetDefaultClientOptions()
	clientconfig := (&clientConfig{}).populate(opts...)
	noOfRetries, minRetryDelay := clientconfig.applyMiddlewareConfig()
	middlewares := GetKiotaMiddlewares(&clientOptions, noOfRetries, minRetryDelay)
	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = time.Minute * 3

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
	}
}

// GetKiotaMiddlewares creates a default slice of middleware for the Graph Client.
func GetKiotaMiddlewares(options *msgraphgocore.GraphClientOptions,
	maxRetry int, minDelay time.Duration,
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
	Client() *msgraphsdk.GraphServiceClient
	// Adapter() returns GraphRequest adapter used to process large requests, create batches
	// and page iterators
	Adapter() *msgraphsdk.GraphRequestAdapter
}

// ---------------------------------------------------------------------------
// Client Middleware
// ---------------------------------------------------------------------------

// LoggingMiddleware can be used to log the http request sent by the graph client
type LoggingMiddleware struct{}

func (handler *LoggingMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	var (
		ctx       = req.Context()
		resp, err = pipeline.Next(req, middlewareIndex)
	)

	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		if strings.Contains(req.URL.String(), "users//") {
			logger.Ctx(ctx).Errorw("malformed request url: missing user", "url", req.URL)
		}
	}

	if resp == nil {
		return resp, err
	}

	// Return immediately if the response is good (2xx).
	// If api logging is toggled, log a body-less dump of the request/resp.
	if (resp.StatusCode / 100) == 2 {
		if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
			respDump, _ := httputil.DumpResponse(resp, false)

			metadata := []any{
				"method", req.Method,
				"status", resp.Status,
				"statusCode", resp.StatusCode,
				"requestLen", req.ContentLength,
				"url", req.URL,
				"response", respDump,
			}

			logger.Ctx(ctx).Debugw("2xx graph api resp", metadata...)
		}

		return resp, err
	}

	// Log errors according to api debugging configurations.
	// When debugging is toggled, every non-2xx is recorded with a respose dump.
	// Otherwise, throttling cases and other non-2xx responses are logged
	// with a slimmer reference for telemetry/supportability purposes.
	if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
		respDump, _ := httputil.DumpResponse(resp, true)

		metadata := []any{
			"method", req.Method,
			"status", resp.Status,
			"statusCode", resp.StatusCode,
			"requestLen", req.ContentLength,
			"url", req.URL,
			"response", string(respDump),
		}

		logger.Ctx(ctx).Errorw("non-2xx graph api response", metadata...)
	} else {
		// special case for supportability: log all throttling cases.
		if resp.StatusCode == http.StatusTooManyRequests {
			logger.Ctx(ctx).Infow("graph api throttling", "method", req.Method, "url", req.URL)
		}

		if resp.StatusCode != http.StatusTooManyRequests && (resp.StatusCode/100) != 2 {
			logger.Ctx(ctx).Infow("graph api error", "status", resp.Status, "method", req.Method, "url", req.URL)
		}
	}

	return resp, err
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
