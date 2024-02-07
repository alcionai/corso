package graph

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/count"
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
	httpClient, cc := KiotaHTTPClient(counter, opts...)

	// This makes sure that we are able to intercept any requests via
	// gock. Only necessary for testing.
	gock.InterceptClient(httpClient)

	ng, err := msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		// Use our own mock auth instance. This allows us to completely avoid having
		// to make requests to microsoft servers during testing.
		authMock{},
		nil, nil,
		httpClient)

	return wrapAdapter(ng, cc), clues.Stack(err).OrNil()
}
