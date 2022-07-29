package operations

import (
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/store"
)

type opStatus int

//go:generate stringer -type=opStatus -linecomment
const (
	Unknown    opStatus = iota // Status Unknown
	InProgress                 // In Progress
	Successful                 // Successful
	Failed                     // Failed
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

	kopia *kopia.Wrapper
	store *store.Wrapper
}

func newOperation(
	opts control.Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
) operation {
	return operation{
		CreatedAt: time.Now(),
		Options:   opts,
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
