package mock

import (
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

// NewMockableClient produces a new exchange api client that can be
// mocked using gock.  Must be used in place of creating an ad-hoc
// client struct.
func NewMockableClient(creds account.M365Config) (api.Client, error) {
	s, err := api.NewService(creds, graph.Mockable())
	if err != nil {
		return api.Client{}, err
	}

	li, err := api.NewService(creds, graph.NoTimeout(), graph.Mockable())
	if err != nil {
		return api.Client{}, err
	}

	return api.Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
	}, nil
}
