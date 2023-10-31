package mock

import (
	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
)

func NewService(
	creds account.M365Config,
	counter *count.Bus,
	opts ...graph.Option,
) (*graph.Service, error) {
	a, err := CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		counter,
		opts...)
	if err != nil {
		return nil, clues.Wrap(err, "generating graph adapter")
	}

	return graph.NewService(a), nil
}

// CreateAdapter is similar to graph.CreateAdapter, but with option to
// enable interceptions via gock to make it mockable.
func CreateAdapter(
	tenant, client, secret string,
	counter *count.Bus,
	opts ...graph.Option,
) (*msgraphsdkgo.GraphRequestAdapter, error) {
	auth, err := graph.GetAuth(tenant, client, secret)
	if err != nil {
		return nil, err
	}

	httpClient, _ := graph.KiotaHTTPClient(counter, opts...)

	// This makes sure that we are able to intercept any requests via
	// gock. Only necessary for testing.
	gock.InterceptClient(httpClient)

	return msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
}
