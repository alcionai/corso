package operations

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
)

type opStatus int

const (
	Unknown opStatus = iota
	InProgress
	Successful
	Failed
)

// --------------------------------------------------------------------------------
// Operation Core
// --------------------------------------------------------------------------------

// An operation tracks the in-progress workload of some long-running process.
// Specific processes (eg: backups, restores, etc) are expected to wrap operation
// with process specific details.
type operation struct {
	ID        uuid.UUID `json:"id"`        // system generated identifier
	CreatedAt time.Time `json:"createdAt"` // datetime of the operation's creation
	Options   Options   `json:"options"`
	Status    opStatus  `json:"status"`

	kopia      *kopia.Wrapper
	modelStore *kopia.ModelStore
}

// Options configure some parameters of the operation
type Options struct {
	// todo: collision handling
	// todo: fast fail vs best attempt
}

func newOperation(
	opts Options,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
) operation {
	return operation{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		Options:    opts,
		kopia:      kw,
		modelStore: ms,
		Status:     InProgress,
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return errors.New("missing kopia connection")
	}
	if op.modelStore == nil {
		return errors.New("missing kopia connection")
	}
	return nil
}

// --------------------------------------------------------------------------------
// Results
// --------------------------------------------------------------------------------

// Summary tracks the total files touched and errors produced
// during an operation.
type summary struct {
	ItemsRead    int   `json:"itemsRead,omitempty"`
	ItemsWritten int   `json:"itemsWritten,omitempty"`
	ReadErrors   error `json:"readErrors,omitempty"`
	WriteErrors  error `json:"writeErrors,omitempty"`
}

// Metrics tracks performance details such as timing, throughput, etc.
type metrics struct {
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}
