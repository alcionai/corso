package graph

import (
	"net"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kauth "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	logGraphRequestsEnvKey    = "LOG_GRAPH_REQUESTS"
	log2xxGraphRequestsEnvKey = "LOG_2XX_GRAPH_REQUESTS"
	log2xxGraphResponseEnvKey = "LOG_2XX_GRAPH_RESPONSES"
	defaultMaxRetries         = 3
	defaultDelay              = 3 * time.Second
	locationHeader            = "Location"
	rateLimitHeader           = "RateLimit-Limit"
	rateRemainingHeader       = "RateLimit-Remaining"
	rateResetHeader           = "RateLimit-Reset"
	retryAfterHeader          = "Retry-After"
	retryAttemptHeader        = "Retry-Attempt"
	defaultHTTPClientTimeout  = 1 * time.Hour
)

// AllMetadataFileNames produces the standard set of filenames used to store graph
// metadata such as delta tokens and folderID->path references.
func AllMetadataFileNames() []string {
	return []string{DeltaURLsFileName, PreviousPathFileName}
}

type QueryParams struct {
	Category      path.CategoryType
	ResourceOwner idname.Provider
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
	opts ...Option,
) (*msgraphsdkgo.GraphRequestAdapter, error) {
	auth, err := GetAuth(tenant, client, secret)
	if err != nil {
		return nil, err
	}

	httpClient := KiotaHTTPClient(opts...)

	return msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
}

func GetAuth(tenant string, client string, secret string) (*kauth.AzureIdentityAuthenticationProvider, error) {
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

	return auth, nil
}

// KiotaHTTPClient creates a httpClient with middlewares and timeout configured
// for use in the graph adapter.
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func KiotaHTTPClient(opts ...Option) *http.Client {
	var (
		clientOptions   = msgraphsdkgo.GetDefaultClientOptions()
		cc              = populateConfig(opts...)
		middlewares     = kiotaMiddlewares(&clientOptions, cc)
		httpClient      = msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
		customTransport = khttp.NewCustomTransportWithParentTransport(defaultHttpTransport(), middlewares...)
	)

	httpClient.Transport = customTransport
	httpClient.Timeout = defaultHTTPClientTimeout

	cc.apply(httpClient)

	return httpClient
}

func defaultHttpTransport() *http.Transport {
	defaultTransport := khttp.GetDefaultTransport().(*http.Transport).Clone()

	// default: 10 seconds
	defaultTransport.TLSHandshakeTimeout = 30 * time.Second

	// default: 30 seconds
	nd := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	defaultTransport.DialContext = nd.DialContext

	return defaultTransport
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

	appendMiddleware []khttp.Middleware
}

type Option func(*clientConfig)

// populate constructs a clientConfig according to the provided options.
func populateConfig(opts ...Option) *clientConfig {
	cc := clientConfig{
		maxRetries: defaultMaxRetries,
		minDelay:   defaultDelay,
	}

	for _, opt := range opts {
		opt(&cc)
	}

	return &cc
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
func NoTimeout() Option {
	return func(c *clientConfig) {
		c.noTimeout = true
	}
}

func MaxRetries(max int) Option {
	return func(c *clientConfig) {
		c.overrideRetryCount = true
		c.maxRetries = max
	}
}

func MinimumBackoff(dur time.Duration) Option {
	return func(c *clientConfig) {
		c.minDelay = dur
	}
}

func appendMiddleware(mw ...khttp.Middleware) Option {
	return func(c *clientConfig) {
		if len(mw) > 0 {
			c.appendMiddleware = mw
		}
	}
}

// ---------------------------------------------------------------------------
// Middleware Control
// ---------------------------------------------------------------------------

// kiotaMiddlewares creates a default slice of middleware for the Graph Client.
func kiotaMiddlewares(
	options *msgraphgocore.GraphClientOptions,
	cc *clientConfig,
) []khttp.Middleware {
	mw := []khttp.Middleware{}

	// Optionally add concurrency limiter middleware if it has been initialized
	if concurrencyLim != nil {
		mw = append(mw, concurrencyLim)
	}

	mw = append(mw, []khttp.Middleware{
		msgraphgocore.NewGraphTelemetryHandler(options),
		&RetryMiddleware{
			MaxRetries: cc.maxRetries,
			Delay:      cc.minDelay,
		},
		khttp.NewRetryHandler(),
		khttp.NewRedirectHandler(),
		khttp.NewCompressionHandler(),
		khttp.NewParametersNameDecodingHandler(),
		khttp.NewUserAgentHandler(),
		&LoggingMiddleware{},
		&RateLimiterMiddleware{},
		&MetricsMiddleware{},
	}...)

	if len(cc.appendMiddleware) > 0 {
		mw = append(mw, cc.appendMiddleware...)
	}

	return mw
}
