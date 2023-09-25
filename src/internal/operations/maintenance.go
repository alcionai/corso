package operations

import (
	"context"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/store"
)

// MaintenanceOperation wraps an operation with restore-specific props.
type MaintenanceOperation struct {
	operation
	Results MaintenanceResults
	mOpts   repository.Maintenance
}

// MaintenanceResults aggregate the details of the results of the operation.
type MaintenanceResults struct {
	stats.StartAndEndTime
}

// NewMaintenanceOperation constructs and validates a maintenance operation.
func NewMaintenanceOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	storer store.BackupStorer,
	mOpts repository.Maintenance,
	bus events.Eventer,
) (MaintenanceOperation, error) {
	op := MaintenanceOperation{
		operation: newOperation(opts, bus, count.New(), kw, storer),
		mOpts:     mOpts,
	}

	err := op.validate()

	return op, clues.Stack(err).OrNil()
}

func (op *MaintenanceOperation) Run(ctx context.Context) (err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "maintenance"); crErr != nil {
			err = crErr
		}
	}()

	op.Results.StartedAt = time.Now()

	op.bus.Event(
		ctx,
		events.MaintenanceStart,
		map[string]any{
			events.StartTime: op.Results.StartedAt,
		})

	defer func() {
		op.bus.Event(
			ctx,
			events.MaintenanceEnd,
			map[string]any{
				events.StartTime: op.Results.StartedAt,
				events.Duration:  op.Results.CompletedAt.Sub(op.Results.StartedAt),
				events.EndTime:   dttm.Format(op.Results.CompletedAt),
				events.Status:    op.Status.String(),
				events.Resources: op.mOpts.Type.String(),
			})
	}()

	return op.do(ctx)
}

func (op *MaintenanceOperation) do(ctx context.Context) error {
	defer func() {
		op.Results.CompletedAt = time.Now()
	}()

	err := op.operation.kopia.RepoMaintenance(ctx, op.store, op.mOpts)
	if err != nil {
		op.Status = Failed
		return clues.Wrap(err, "running maintenance operation")
	}

	op.Status = Completed

	return nil
}
