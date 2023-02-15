package api

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

// ---------------------------------------------------------------------------
// common types and consts
// ---------------------------------------------------------------------------

// DeltaUpdate holds the results of a current delta token.  It normally
// gets produced when aggregating the addition and removal of items in
// a delta-queriable folder.
type DeltaUpdate struct {
	// the deltaLink itself
	URL string
	// true if the old delta was marked as invalid
	Reset bool
}

// GraphQuery represents functions which perform exchange-specific queries
// into M365 backstore. Responses -> returned items will only contain the information
// that is included in the options
// TODO: use selector or path for granularity into specific folders or specific date ranges
type GraphQuery func(ctx context.Context, userID string) (serialization.Parsable, error)

// GraphRetrievalFunctions are functions from the Microsoft Graph API that retrieve
// the default associated data of a M365 object. This varies by object. Additional
// Queries must be run to obtain the omitted fields.
type GraphRetrievalFunc func(
	ctx context.Context,
	user, m365ID string,
) (serialization.Parsable, error)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// Client is used to fulfill the interface for exchange
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

	// The largeItem graph servicer is configured specifically for
	// downloading large items.  Specifically for use when handling
	// attachments, and for no other use.
	largeItem graph.Servicer
}

// NewClient produces a new exchange api client.  Must be used in
// place of creating an ad-hoc client struct.
func NewClient(creds account.M365Config) (Client, error) {
	s, err := newService(creds)
	if err != nil {
		return Client{}, err
	}

	li, err := newLargeItemService(creds)
	if err != nil {
		return Client{}, err
	}

	return Client{creds, s, li}, nil
}

// service generates a new service.  Used for paged and other long-running
// requests instead of the client's stable service, so that in-flight state
// within the adapter doesn't get clobbered
func (c Client) service() (*graph.Service, error) {
	s, err := newService(c.Credentials)
	return s, err
}

func newService(creds account.M365Config) (*graph.Service, error) {
	a, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret)
	if err != nil {
		return nil, errors.Wrap(err, "generating no-timeout graph adapter")
	}

	return graph.NewService(a), nil
}

func newLargeItemService(creds account.M365Config) (*graph.Service, error) {
	a, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		graph.NoTimeout())
	if err != nil {
		return nil, errors.Wrap(err, "generating no-timeout graph adapter")
	}

	return graph.NewService(a), nil
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// checkIDAndName is a helper function to ensure that
// the ID and name pointers are set prior to being called.
func checkIDAndName(c graph.Container) error {
	id := ptr.Val(c.GetId())
	if len(id) == 0 {
		return errors.New("container missing ID")
	}

	dn := ptr.Val(c.GetDisplayName())
	if len(dn) == 0 {
		return clues.New("container missing display name").With("container_id", id)
	}

	return nil
}

func HasAttachments(body models.ItemBodyable) bool {
	if body.GetContent() == nil || body.GetContentType() == nil ||
		*body.GetContentType() == models.TEXT_BODYTYPE || len(*body.GetContent()) == 0 {
		return false
	}

	return strings.Contains(ptr.Val(body.GetContent()), "src=\"cid:")
}
