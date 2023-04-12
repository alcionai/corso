package graph

import (
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kauth "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	logGraphRequestsEnvKey    = "LOG_GRAPH_REQUESTS"
	log2xxGraphRequestsEnvKey = "LOG_2XX_GRAPH_REQUESTS"
	log2xxGraphResponseEnvKey = "LOG_2XX_GRAPH_RESPONSES"
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
// Adapter
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// HTTP Client Config
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

// ---------------------------------------------------------------------------
// Middleware Control
// ---------------------------------------------------------------------------

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
