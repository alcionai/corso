package api

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

// ---------------------------------------------------------------------------
// common types
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
type GraphQuery func(ctx context.Context, gs graph.Servicer, userID string) (serialization.Parsable, error)

// GraphRetrievalFunctions are functions from the Microsoft Graph API that retrieve
// the default associated data of a M365 object. This varies by object. Additional
// Queries must be run to obtain the omitted fields.
type GraphRetrievalFunc func(
	ctx context.Context,
	gs graph.Servicer,
	user, m365ID string,
) (serialization.Parsable, error)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

// API is a struct used to fulfill the interface for exchange
// queries that are traditionally backed by GraphAPI.  A
// struct is used in this case, instead of deferring to
// pure function wrappers, so that the boundary separates the
// granular implementation of the graphAPI and kiota away
// from the exchange package's broader intents.
// type API struct{}
