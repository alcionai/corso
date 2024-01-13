package operations

import (
	"context"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
)

// RetentionConfigOperation wraps an operation with restore-specific props.
type RetentionConfigOperation struct {
	operation
	Results RetentionConfigResults
	rcOpts  repository.Retention
}

// RetentionConfigResults aggregate the details of the results of the operation.
type RetentionConfigResults struct {
	stats.StartAndEndTime
}

// NewRetentionConfigOperation constructs and validates an operation to change
// retention parameters.
func NewRetentionConfigOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	rcOpts repository.Retention,
	bus events.Eventer,
) (RetentionConfigOperation, error) {
	op := RetentionConfigOperation{
		operation: newOperation(opts, bus, count.New(), kw, nil),
		rcOpts:    rcOpts,
	}

	// Don't run validation because we don't populate the model store.

	return op, nil
}

func (op *RetentionConfigOperation) Run(ctx context.Context) (err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "retention_config"); crErr != nil {
			err = crErr
		}
	}()

	op.Results.StartedAt = time.Now()

	// TODO(ashmrtn): Send telemetry?

	return op.do(ctx)
}

func (op *RetentionConfigOperation) do(ctx context.Context) error {
	defer func() {
		op.Results.CompletedAt = time.Now()
	}()

	err := op.operation.kopia.SetRetentionParameters(ctx, op.rcOpts)
	if err != nil {
		op.Status = Failed
		return clues.Wrap(err, "running retention config operation")
	}

	op.Status = Completed

	return nil
}
