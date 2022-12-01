package operations

import (
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/pkg/control"
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
	CreatedAt time.Time       `json:"createdAt"` // datetime of the operation's creation
	Options   control.Options `json:"options"`
	Status    opStatus        `json:"status"`

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
		Options:   opts,
		bus:       bus,
		kopia:     kw,
		store:     sw,
		Status:    InProgress,
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return errors.New("missing kopia connection")
	}

	if op.store == nil {
		return errors.New("missing modelstore")
	}

	return nil
}
