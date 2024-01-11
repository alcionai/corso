package graph

import (
	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
)

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
	auth, err := GetAuth(tenant, client, secret)
	if err != nil {
		return nil, err
	}

	httpClient, cc := KiotaHTTPClient(counter, opts...)

	// This makes sure that we are able to intercept any requests via
	// gock. Only necessary for testing.
	gock.InterceptClient(httpClient)

	ng, err := msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)

	return wrapAdapter(ng, cc), clues.Stack(err).OrNil()
}
