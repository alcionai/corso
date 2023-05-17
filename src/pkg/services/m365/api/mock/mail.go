package mock

import (
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/mock"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// NewClient produces a new exchange api client that can be
// mocked using gock.
func NewClient(creds account.M365Config) (api.Client, error) {
	s, err := mock.NewService(creds)
	if err != nil {
		return api.Client{}, err
	}

	li, err := mock.NewService(creds, graph.NoTimeout())
	if err != nil {
		return api.Client{}, err
	}

	return api.Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
	}, nil
}
