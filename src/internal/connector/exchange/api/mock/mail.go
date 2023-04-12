package mock

import (
	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

// Similar to api.CreateAdapter, but with option to enable mocking via gock
func CreateMockableAdapter(
	tenant, client, secret string,
	opts ...graph.Option,
) (*msgraphsdkgo.GraphRequestAdapter, error) {
	auth, err := graph.GetAuth(tenant, client, secret)
	if err != nil {
		return nil, err
	}

	httpClient := graph.HTTPClient(opts...)

	// This makes sure that we are able to intercept any requests via
	// gock. Only necessary for testing.
	gock.InterceptClient(httpClient)

	return msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
}

func NewMockableService(creds account.M365Config, opts ...graph.Option) (*graph.Service, error) {
	a, err := CreateMockableAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		opts...)
	if err != nil {
		return nil, clues.Wrap(err, "generating graph adapter")
	}

	return graph.NewService(a), nil
}

// NewMockableClient produces a new exchange api client that can be
// mocked using gock.  Must be used in place of creating an ad-hoc
// client struct.
func NewMockableClient(creds account.M365Config) (api.Client, error) {
	s, err := NewMockableService(creds)
	if err != nil {
		return api.Client{}, err
	}

	li, err := NewMockableService(creds, graph.NoTimeout())
	if err != nil {
		return api.Client{}, err
	}

	return api.Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
	}, nil
}
