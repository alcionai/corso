package graph

import (
	"context"
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
	"github.com/alcionai/corso/src/pkg/selectors"
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

	clientOptions := msgraphsdk.GetDefaultClientOptions()
	middlewares := msgraphgocore.GetDefaultMiddlewaresWithOptions(&clientOptions)

	// When true, additional logging middleware support added for http request
	if os.Getenv(logGraphRequestsEnvKey) != "" {
		middlewares = append(middlewares, &LoggingMiddleware{})
	}

	httpClient := msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	httpClient.Timeout = time.Second * 90

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

// ScopeToPathCategory helper function that maps selectors.ExchangeScope to path.CategoryType
func ScopeToPathCategory(scope selectors.ExchangeScope) path.CategoryType {
	if scope.IncludesCategory(selectors.ExchangeMail) {
		return path.EmailCategory
	}

	if scope.IncludesCategory(selectors.ExchangeContact) {
		return path.ContactsCategory
	}

	if scope.IncludesCategory(selectors.ExchangeEvent) {
		return path.EventsCategory
	}

	return path.UnknownCategory
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
	default:
		return path.UnknownCategory
	}
}
