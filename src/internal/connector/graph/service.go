package graph

import (
	"context"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/alcionai/corso/src/internal/connector/support"
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
	logGraphRequestsEnvKey = "LOG_GRAPH_REQUESTS"
	numberOfRetries        = 1
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
		return nil, errors.Wrap(err, "writeObjecValue serialization")
	}

	return writer.GetSerializedContent()
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

type clientConfig struct {
	noTimeout bool
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

// CreateAdapter uses provided credentials to log into M365 using Kiota Azure Library
// with Azure identity package. An adapter object is a necessary to component
// to create  *msgraphsdk.GraphServiceClient
func CreateAdapter(tenant, client, secret string, opts ...option) (*msgraphsdk.GraphRequestAdapter, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := azidentity.NewClientSecretCredential(tenant, client, secret, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating m365 client secret credentials")
	}

	auth, err := ka.NewAzureIdentityAuthenticationProviderWithScopes(
		cred,
		[]string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating new AzureIdentityAuthentication")
	}

	httpClient := HTTPClient(opts...)

	return msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth, nil, nil, httpClient)
}

// HTTPClient creates the httpClient with middlewares and timeout configured
//
// Re-use of http clients is critical, or else we leak OS resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
func HTTPClient(opts ...option) *http.Client {
	clientOptions := msgraphsdk.GetDefaultClientOptions()
	middlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)
	middlewares = append(middlewares, &LoggingMiddleware{})
	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = time.Minute * 3

	(&clientConfig{}).
		populate(opts...).
		apply(httpClient)

	return httpClient
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

// Idable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of an ID.
type Idable interface {
	GetId() *string
}

// Descendable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a "parent folder".
type Descendable interface {
	Idable
	GetParentFolderId() *string
}

// Displayable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a display name.
type Displayable interface {
	Idable
	GetDisplayName() *string
}

type Container interface {
	Descendable
	Displayable
}

// ContainerResolver houses functions for getting information about containers
// from remote APIs (i.e. resolve folder paths with Graph API). Resolvers may
// cache information about containers.
type ContainerResolver interface {
	// IDToPath takes an m365 container ID and converts it to a hierarchical path
	// to that container. The path has a similar format to paths on the local
	// file system.
	IDToPath(ctx context.Context, m365ID string) (*path.Builder, error)
	// Populate performs initialization steps for the resolver
	// @param ctx is necessary param for Graph API tracing
	// @param baseFolderID represents the M365ID base that the resolver will
	// conclude its search. Default input is "".
	Populate(ctx context.Context, baseFolderID string, baseContainerPather ...string) error

	// PathInCache performs a look up of a path reprensentation
	// and returns the m365ID of directory iff the pathString
	// matches the path of a container within the cache.
	// @returns bool represents if m365ID was found.
	PathInCache(pathString string) (string, bool)

	AddToCache(ctx context.Context, m365Container Container) error

	// Items returns the containers in the cache.
	Items() []CachedContainer
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

// Run a function with retries
func RunWithRetry(run func() error) error {
	var err error

	for i := 0; i < numberOfRetries; i++ {
		err = run()
		if err == nil {
			return nil
		}

		// only retry on timeouts and 500-internal-errors.
		if !(IsErrTimeout(err) || IsInternalServerError(err)) {
			break
		}

		if i < numberOfRetries {
			time.Sleep(time.Duration(0*(i+2)) * time.Second)
		}
	}

	return support.ConnectorStackErrorTraceWrap(err, "maximum retries or unretryable")
}
