package api

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Client is used to fulfill the interface for discovery
// queries that are traditionally backed by GraphAPI.  A
// struct is used in this case, instead of deferring to
// pure function wrappers, so that the boundary separates the
// granular implementation of the graphAPI and kiota away
// from the exchange package's broader intents.
type Client struct {
	Credentials account.M365Config

	// The stable service is re-usable for any non-paged request.
	// This allows us to maintain performance across async requests.
	stable graph.Servicer
}

// NewClient produces a new exchange api client.  Must be used in
// place of creating an ad-hoc client struct.
func NewClient(creds account.M365Config) (Client, error) {
	s, err := newService(creds)
	if err != nil {
		return Client{}, err
	}

	return Client{creds, s}, nil
}

// service generates a new service.  Used for paged and other long-running
// requests instead of the client's stable service, so that in-flight state
// within the adapter doesn't get clobbered
func (c Client) service() (*graph.Service, error) {
	return newService(c.Credentials)
}

func newService(creds account.M365Config) (*graph.Service, error) {
	adapter, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
	)
	if err != nil {
		return nil, errors.Wrap(err, "generating graph api service client")
	}

	return graph.NewService(adapter), nil
}
