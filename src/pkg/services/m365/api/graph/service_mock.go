package graph

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
)

// authMock implements the
// github.com/microsoft.kiota-abstractions-go/authentication:AuthenticationProvider
// interface.
type authMock struct{}

// AuthenticateRequest is the function called prior to sending a graph API
// request. It ensures the client has the proper authentication to send the
// request. Returning nil allows us to skip the authentication that would
// normally happen.
func (a authMock) AuthenticateRequest(
	context.Context,
	*abstractions.RequestInformation,
	map[string]any,
) error {
	return nil
}

func NewGockService(
	creds account.M365Config,
	counter *count.Bus,
	opts ...Option,
) (*Service, error) {
	// Need to initialize the concurrency limiter else we'll get an error.
	InitializeConcurrencyLimiter(context.Background(), true, 4)

	a, err := CreateGockAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		counter,
		opts...)
	if err != nil {
		return nil, clues.Wrap(err, "generating graph adapter")
	}

	return NewService(a), nil
}

// CreateGockAdapter is similar to graph.CreateAdapter, but with option to
// enable interceptions via gock to make it mockable.
func CreateGockAdapter(
	tenant, client, secret string,
	counter *count.Bus,
	opts ...Option,
) (abstractions.RequestAdapter, error) {
	cc := populateConfig(opts...)
	// We need to manufacture our own graph client since gock.InterceptClient
	// replaces the transport which replaces all the middleware we add. Start with
	// a client without any middleware and then replace the transport completely
	// with the mocked one since there's no real in between that we could use.
	clientOptions := msgraphsdkgo.GetDefaultClientOptions()
	middlewares := kiotaMiddlewares(&clientOptions, cc, counter)

	//nolint:lll
	// This is lifted from
	// https://github.com/microsoft/kiota-http-go/blob/e32eb086c8d28002dcba922f3271d56327ba8b03/pipeline.go#L75
	// which was found by following
	// https://github.com/microsoftgraph/msgraph-sdk-go-core/blob/93a2c8acb7dfff7f3e2791670f51ccb001d7b127/graph_client_factory.go#L26
	httpClient := khttp.GetDefaultClient()
	httpClient.Transport = khttp.NewCustomTransportWithParentTransport(
		// gock's default transport isn't quite the same as the default graph uses
		// but since we're mocking things it's probably ok.
		gock.NewTransport(),
		middlewares...)

	cc.apply(httpClient)

	ng, err := msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		// Use our own mock auth instance. This allows us to completely avoid having
		// to make requests to microsoft servers during testing.
		authMock{},
		nil, nil,
		httpClient)

	return wrapAdapter(ng, cc), clues.Stack(err).OrNil()
}
