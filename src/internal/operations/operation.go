package operations

import (
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/store"
)

// OpStatus describes the current status of an operation.
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
type OpStatus int

//go:generate stringer -type=OpStatus -linecomment
const (
	Unknown    OpStatus = iota // Status Unknown
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
	Status  OpStatus        `json:"status"`

	bus   events.Eventer
	kopia *kopia.Wrapper
	Store *store.Wrapper
}

func newOperation(
	opts control.Options,
	bus events.Eventer,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
) operation {
	return operation{
		CreatedAt: time.Now(),
		Errors:    fault.New(opts.FailureHandling == control.FailFast),
		Options:   opts,

		bus:   bus,
		kopia: kw,
		Store: sw,

		Status: InProgress,
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return clues.New("missing kopia connection")
	}

	if op.Store == nil {
		return clues.New("missing modelstore")
	}

	return nil
}
