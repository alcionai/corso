package graph

import (
	"context"
	nethttp "net/http"
	"net/http/httputil"
	"os"
	"time"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"

	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	logGraphRequestsEnvKey = "LOG_GRAPH_REQUESTS"
)

// CreateAdapter uses provided credentials to log into M365 using Kiota Azure Library
// with Azure identity package.
func CreateAdapter(tenant, client, secret string) (*msgraphsdk.GraphRequestAdapter, error) {
	// Client Provider: Uses Secret for access to tenant-level data
	cred, err := az.NewClientSecretCredential(tenant, client, secret, nil)
	if err != nil {
		return nil, err
	}

	auth, err := ka.NewAzureIdentityAuthenticationProviderWithScopes(
		cred,
		[]string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		return nil, err
	}

	// If the "LOG_GRAPH_REQUESTS" environment variable is not set, return
	// the default client
	if os.Getenv(logGraphRequestsEnvKey) == "" {
		clientOptions := msgraphsdk.GetDefaultClientOptions()
		defaultMiddlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)
		httpClient := msgraphgocore.GetDefaultClient(&clientOptions, defaultMiddlewares...)
		httpClient.Timeout = time.Second * 90
		return msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
			auth, nil, nil, httpClient)
	}

	// Create a client with logging middleware
	clientOptions := msgraphsdk.GetDefaultClientOptions()
	defaultMiddlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)
	middlewares := []khttp.Middleware{&LoggingMiddleware{}}
	middlewares = append(middlewares, defaultMiddlewares...)
	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)

	return msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth, nil, nil, httpClient)
}

// LoggingMiddleware can be used to log the http request sent by the graph client
type LoggingMiddleware struct{}

// Intercept implements the RequestInterceptor interface and decodes the parameters name
func (handler *LoggingMiddleware) Intercept(
	pipeline khttp.Pipeline, middlewareIndex int, req *nethttp.Request,
) (*nethttp.Response, error) {
	requestDump, _ := httputil.DumpRequest(req, true)
	logger.Ctx(context.TODO()).Infof("REQUEST: %s", string(requestDump))

	return pipeline.Next(req, middlewareIndex)
}
