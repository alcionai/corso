package mock

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/connector/exchange/api"
	"github.com/alcionai/corso/src/pkg/connector/graph"
	"github.com/alcionai/corso/src/pkg/connector/graph/mock"
)

func NewService(creds account.M365Config, opts ...graph.Option) (*graph.Service, error) {
	a, err := mock.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		opts...)
	if err != nil {
		return nil, clues.Wrap(err, "generating graph adapter")
	}

	return graph.NewService(a), nil
}

// NewClient produces a new exchange api client that can be
// mocked using gock.
func NewClient(creds account.M365Config) (api.Client, error) {
	s, err := NewService(creds)
	if err != nil {
		return api.Client{}, err
	}

	li, err := NewService(creds, graph.NoTimeout())
	if err != nil {
		return api.Client{}, err
	}

	return api.Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
	}, nil
}
