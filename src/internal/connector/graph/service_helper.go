package graph

import (
	nethttp "net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	az "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	ka "github.com/microsoft/kiota-authentication-azure-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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
func CreateHTTPClient() *nethttp.Client {
	clientOptions := msgraphsdk.GetDefaultClientOptions()
	middlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)
	middlewares = append(middlewares, &LoggingMiddleware{})
	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = time.Second * 90

	return httpClient
}

// LoggingMiddleware can be used to log the http request sent by the graph client
type LoggingMiddleware struct{}

// Intercept implements the RequestInterceptor interface and decodes the parameters name
func (handler *LoggingMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *nethttp.Request,
) (*nethttp.Response, error) {
	var (
		ctx       = req.Context()
		resp, err = pipeline.Next(req, middlewareIndex)
		metadata  = []any{}
	)

	if resp != nil && (resp.StatusCode/100) != 2 {
		respDump, _ := httputil.DumpResponse(resp, true)
		metadata = []any{
			"method", req.Method,
			"url", req.URL,
			"requestLen", req.ContentLength,
			"status", resp.Status,
			"statusCode", resp.StatusCode,
			"request", string(respDump),
		}
	}

	if logger.DebugAPI || os.Getenv(logGraphRequestsEnvKey) != "" {
		logger.Ctx(ctx).Debugw("non-2xx graph api response", metadata...)
	}

	return resp, err
}

func StringToPathCategory(input string) path.CategoryType {
	param := strings.ToLower(input)

	switch param {
	case "email":
		return path.EmailCategory
	case "contacts":
		return path.ContactsCategory
	case "events":
		return path.EventsCategory
	case "files":
		return path.FilesCategory
	case "libraries":
		return path.LibrariesCategory
	default:
		return path.UnknownCategory
	}
}
