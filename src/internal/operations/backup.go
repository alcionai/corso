package operations

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/credentials"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation
	Version string

	creds credentials.M365

	Targets []string // something for targets/filter/source/app&users/etc
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts OperationOpts,
	kw *kopia.KopiaWrapper,
	creds credentials.M365,
	targets []string,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, kw),
		Version:   "v0",
		creds:     creds,
		Targets:   targets,
	}
	if err := op.validate(); err != nil {
		return BackupOperation{}, err
	}

	return op, nil
}

func (op BackupOperation) validate() error {
	if err := op.creds.Validate(); err != nil {
		return errors.Wrap(err, "invalid credentials")
	}
	return op.operation.validate()
}

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (*kopia.BackupStats, error) {
	gc, err := connector.NewGraphConnector(op.creds.TenantID, op.creds.ClientID, op.creds.ClientSecret)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to graph api")
	}

	cs, err := gc.ExchangeDataCollection(ctx, op.Targets[0])
	if err != nil {
		return nil, errors.Wrap(err, "retrieving application data")
	}

	// todo: utilize stats
	stats, err := op.kopia.BackupCollections(ctx, cs)
	if err != nil {
		return nil, errors.Wrap(err, "backing up application data")
	}

	op.Status = Successful
	return stats, nil
}
