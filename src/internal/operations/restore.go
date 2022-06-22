package operations

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/credentials"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation
	Version string

	restorePointID string
	creds          credentials.M365

	Targets []string // something for targets/filter/source/app&users/etc
}

// NewRestoreOperation constructs and validates a restore operation.
func NewRestoreOperation(
	ctx context.Context,
	opts OperationOpts,
	kw *kopia.KopiaWrapper,
	creds credentials.M365,
	restorePointID string,
	targets []string,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation: newOperation(opts, kw),
		Version:   "v0",
		creds:     creds,
		Targets:   targets,
	}
	if err := op.validate(); err != nil {
		return RestoreOperation{}, err
	}

	return op, nil
}

func (op RestoreOperation) validate() error {
	if err := op.creds.Validate(); err != nil {
		return errors.Wrap(err, "invalid credentials")
	}
	return op.operation.validate()
}

// Run begins a synchronous restore operation.
// todo (keepers): return stats block in first param.
func (op *RestoreOperation) Run(ctx context.Context) error {
	dc, err := op.kopia.RestoreSingleItem(ctx, op.restorePointID, op.Targets)
	if err != nil {
		return errors.Wrap(err, "retrieving service data")
	}

	gc, err := connector.NewGraphConnector(op.creds.TenantID, op.creds.ClientID, op.creds.ClientSecret)
	if err != nil {
		return errors.Wrap(err, "connecting to graph api")
	}

	if err := gc.RestoreMessages(ctx, dc); err != nil {
		return errors.Wrap(err, "restoring service data")
	}

	op.Status = Successful
	return nil
}
