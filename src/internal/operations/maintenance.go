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
)

// MaintenanceOperation wraps an operation with restore-specific props.
type MaintenanceOperation struct {
	operation
	Results          MaintenanceResults
	safety           control.Safety
	force            bool
	quickMaintenance bool
}

// MaintenanceResults aggregate the details of the results of the operation.
type MaintenanceResults struct {
	stats.StartAndEndTime
}

// NewMaintenanceOperation constructs and validates a restore operation.
func NewMaintenanceOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	safety control.Safety,
	quickMaintenance, force bool,
	bus events.Eventer,
) (MaintenanceOperation, error) {
	op := MaintenanceOperation{
		operation:        newOperation(opts, bus, kw, nil),
		safety:           safety,
		quickMaintenance: quickMaintenance,
		force:            force,
	}

	// Don't run validation because we don't populate the model store.

	return op, nil
}

func (op *MaintenanceOperation) Run(ctx context.Context) (err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover(), "maintenance"); crErr != nil {
			err = crErr
		}

		// TODO(ashmrtn): Send success/failure usage stat?

		op.Results.CompletedAt = time.Now()
	}()

	op.Results.StartedAt = time.Now()

	// TODO(ashmrtn): Send usage statistics?

	err = op.operation.kopia.Maintenance(
		ctx,
		op.safety,
		op.quickMaintenance,
		op.force)
	if err != nil {
		return clues.Wrap(err, "running maintenance operation")
	}

	return nil
}
