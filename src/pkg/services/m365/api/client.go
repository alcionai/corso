package api

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Client is used to fulfill queries that are traditionally
// backed by GraphAPI. A struct is used in this case, instead
// of deferring to pure function wrappers, so that the boundary
// separates the granular implementation of the graphAPI and
// kiota away from the other packages.
type Client struct {
	Credentials account.M365Config

	// The Stable service is re-usable for any non-paged request.
	// This allows us to maintain performance across async requests.
	Stable graph.Servicer

	// The LargeItem graph servicer is configured specifically for
	// downloading large items such as drive item content or outlook
	// mail and event attachments.
	LargeItem graph.Servicer
}

// NewClient produces a new exchange api client.  Must be used in
// place of creating an ad-hoc client struct.
func NewClient(creds account.M365Config) (Client, error) {
	s, err := NewService(creds)
	if err != nil {
		return Client{}, err
	}

	li, err := newLargeItemService(creds)
	if err != nil {
		return Client{}, err
	}

	return Client{creds, s, li}, nil
}

// Service generates a new graph servicer.  New servicers are used for paged
// and other long-running requests instead of the client's stable service,
// so that in-flight state within the adapter doesn't get clobbered.
// Most calls should use the Client.Stable property instead of calling this
// func, unless it is explicitly necessary.
func (c Client) Service() (graph.Servicer, error) {
	return NewService(c.Credentials)
}

func NewService(creds account.M365Config, opts ...graph.Option) (*graph.Service, error) {
	a, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		opts...)
	if err != nil {
		return nil, clues.Wrap(err, "generating graph api adapter")
	}

	return graph.NewService(a), nil
}

func newLargeItemService(creds account.M365Config) (*graph.Service, error) {
	a, err := NewService(creds, graph.NoTimeout())
	if err != nil {
		return nil, clues.Wrap(err, "generating no-timeout graph adapter")
	}

	return a, nil
}
