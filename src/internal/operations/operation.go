package operations

import (
	"time"

	"github.com/google/uuid"
	multierror "github.com/hashicorp/go-multierror"
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
	ID        uuid.UUID     `json:"id"`        // system generated identifier
	CreatedAt time.Time     `json:"createdAt"` // datetime of the operation's creation
	Options   OperationOpts `json:"options"`
	Status    opStatus      `json:"status"`

	kopia *kopia.KopiaWrapper
}

// OperationOpts configure some parameters of the operation
type OperationOpts struct {
	// todo: collision handling
	// todo: fast fail vs best attempt
}

func newOperation(
	opts OperationOpts,
	kw *kopia.KopiaWrapper,
) operation {
	return operation{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Options:   opts,
		kopia:     kw,
		Status:    InProgress,
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return errors.New("missing kopia connection")
	}
	return nil
}

// --------------------------------------------------------------------------------
// Results
// --------------------------------------------------------------------------------

// Summary tracks the total files touched and errors produced
// during an operation.
type operationSummary struct {
	ItemsRead    int              `json:"itemsRead"`
	ItemsWritten int              `json:"itemsWritten"`
	ReadErrors   multierror.Error `json:"readErrors"`
	WriteErrors  multierror.Error `json:"writeErrors"`
}

// Metrics tracks performance details such as timing, throughput, etc.
type operationMetrics struct {
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}
