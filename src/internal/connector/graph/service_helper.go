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
		return resp, err
	}

	// special case for supportability: log all throttling cases.
	if resp.StatusCode == http.StatusTooManyRequests {
		logger.Ctx(ctx).Infow("graph api throttling", "method", req.Method, "url", req.URL)
	}

	if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
		respDump, _ := httputil.DumpResponse(resp, true)

		metadata := []any{
			"method", req.Method,
			"url", req.URL,
			"requestLen", req.ContentLength,
			"status", resp.Status,
			"statusCode", resp.StatusCode,
			"request", string(respDump),
		}

		logger.Ctx(ctx).Errorw("non-2xx graph api response", metadata...)
	}

	return resp, err
}
