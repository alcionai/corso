package graph

import (
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	logGraphRequestsEnvKey = "LOG_GRAPH_REQUESTS"
)

// CreateAdapter uses provided credentials to log into M365 using Kiota Azure Library
// with Azure identity package. An adapter object is a necessary to component
// to create  *msgraphsdk.GraphServiceClient
func CreateAdapter(tenant, client, secret string) (*msgraphsdk.GraphRequestAdapter, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := az.NewClientSecretCredential(tenant, client, secret, nil)
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

	httpClient := CreateHTTPClient()

	return msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth, nil, nil, httpClient)
}

// CreateHTTPClient creates the httpClient with middlewares and timeout configured
func CreateHTTPClient() *http.Client {
	clientOptions := msgraphsdk.GetDefaultClientOptions()
	middlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)
	middlewares = append(middlewares, &LoggingMiddleware{})
	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = time.Second * 90

	return httpClient
}

// LargeItemClient generates a client that's configured to handle
// large file downloads.  This client isn't suitable for other queries
// due to loose restrictions on timeouts and such.
//
// Re-use of http clients is critical, or else we leak os resources
// and consume relatively unbound socket connections.  It is important
// to centralize this client to be passed downstream where api calls
// can utilize it on a per-download basis.
//
// TODO: this should get owned by an API client layer, not the GC itself.
func LargeItemClient() *http.Client {
	httpClient := CreateHTTPClient()
	httpClient.Timeout = 0 // infinite timeout for pulling large files

	return httpClient
}

// ---------------------------------------------------------------------------
// Logging Middleware
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

	if (resp.StatusCode / 100) == 2 {
		if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
			respDump, _ := httputil.DumpResponse(resp, false)

			metadata := []any{
				"idx", middlewareIndex,
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

	if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
		respDump, _ := httputil.DumpResponse(resp, true)

		metadata := []any{
			"idx", middlewareIndex,
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
