package mock

import (
	"github.com/h2non/gock"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

// CreateAdapter is similar to graph.CreateAdapter, but with option to
// enable interceptions via gock to make it mockable.
func CreateAdapter(
	tenant, client, secret string,
	opts ...graph.Option,
) (*msgraphsdkgo.GraphRequestAdapter, error) {
	auth, err := graph.GetAuth(tenant, client, secret)
	if err != nil {
		return nil, err
	}

	httpClient := graph.KiotaHTTPClient(opts...)

	// This makes sure that we are able to intercept any requests via
	// gock. Only necessary for testing.
	gock.InterceptClient(httpClient)

	return msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
}
