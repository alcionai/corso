package operations

import (
	"context"
	"errors"
	"time"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/google/uuid"
)

type opStatus int

const (
	Unknown opStatus = iota
	InProgress
	Successful
	Failed
)

// An operation tracks the in-progress workload of some long-running process.
// Specific processes (eg: backups, restores, etc) are expected to wrap operation
// with process specific details.
type operation struct {
	ID        uuid.UUID // system generated identifier
	CreatedAt time.Time // datetime of the operation's creation

	options OperationOpts
	kopia   *kopia.KopiaWrapper

	// TODO(rkeepers) deal with circular dependencies here
	// graphConn    GraphConnector  // m365 details

	Status opStatus
	Errors []error
}

type logger interface {
	Debug(context.Context, string)
	Info(context.Context, string)
	Warn(context.Context, string)
	Error(context.Context, string)
}

// OperationOpts configure some parameters of the operation
type OperationOpts struct {
	Logger logger
}

func newOperation(opts OperationOpts, kw *kopia.KopiaWrapper) operation {
	return operation{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		options:   opts,
		kopia:     kw,
		Status:    InProgress,
		Errors:    []error{},
	}
}

func (op operation) validate() error {
	if op.kopia == nil {
		return errors.New("missing kopia connection")
	}
	return nil
}
