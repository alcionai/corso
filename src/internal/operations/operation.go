package operations

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// opStatus describes the current status of an operation.
// InProgress - the standard value for any process that has not
// arrived at an end state.  The end states are Failed, Completed,
// or NoData.
//
// Failed - the operation was unable to begin processing data at all.
// No items have been written by the consumer.
//
// Completed - the operation was able to process one or more of the
// items in the request. Both partial success (0 < N < len(items)
// errored) and total success (0 errors) are set as Completed.
//
// NoData - only occurs when no data was involved in an operation.
// For example, if a backup is requested for a specific user's
// mail, but that account contains zero mail messages, the backup
// contains No Data.
type opStatus int

//go:generate stringer -type=opStatus -linecomment
const (
	Unknown    opStatus = iota // Status Unknown
	InProgress                 // In Progress
	Completed                  // Completed
	Failed                     // Failed
	NoData                     // No Data
)

// --------------------------------------------------------------------------------
// Operation Core
// --------------------------------------------------------------------------------

// An operation tracks the in-progress workload of some long-running process.
// Specific processes (eg: backups, restores, etc) are expected to wrap operation
// with process specific details.
type operation struct {
	CreatedAt time.Time `json:"createdAt"`

	Errors  *fault.Bus      `json:"errors"`
	Options control.Options `json:"options"`
	Status  opStatus        `json:"status"`

	bus   events.Eventer
	kopia *kopia.Wrapper
	store *store.Wrapper
}

func newOperation(
	opts control.Options,
	bus events.Eventer,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
) operation {
	return operation{
		CreatedAt: time.Now(),
		Errors:    fault.New(opts.FailFast),
		Options:   opts,

		bus:   bus,
		kopia: kw,
		store: sw,

		Status: InProgress,
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return clues.New("missing kopia connection")
	}

	if op.store == nil {
		return clues.New("missing modelstore")
	}

	return nil
}

// produces a graph connector.
func connectToM365(
	ctx context.Context,
	sel selectors.Selector,
	acct account.Account,
	errs *fault.Bus,
) (*connector.GraphConnector, error) {
	complete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Connecting to M365"))
	defer func() {
		complete <- struct{}{}
		close(complete)
		closer()
	}()

	// retrieve data from the producer
	resource := connector.Users
	if sel.Service == selectors.ServiceSharePoint {
		resource = connector.Sites
	}

	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, resource, errs)
	if err != nil {
		return nil, err
	}

	return gc, nil
}
