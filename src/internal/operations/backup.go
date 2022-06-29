package operations

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/account"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation

	Results BackupResults `json:"results"`
	Targets []string      `json:"selectors"` // todo: replace with Selectors
	Version string        `json:"version"`

	account account.Account
}

// BackupResults aggregate the details of the result of the operation.
type BackupResults struct {
	operationSummary
	operationMetrics
	// todo: RestorePoint RestorePoint
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts OperationOpts,
	kw *kopia.KopiaWrapper,
	acct account.Account,
	targets []string,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, kw),
		Targets:   targets,
		Version:   "v0",
		account:   acct,
	}
	if err := op.validate(); err != nil {
		return BackupOperation{}, err
	}

	return op, nil
}

func (op BackupOperation) validate() error {
	return op.operation.validate()
}

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (*kopia.BackupStats, error) {
	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to graph api")
	}

	cs, err := gc.ExchangeDataCollection(ctx, op.Targets[0])
	if err != nil {
		return nil, errors.Wrap(err, "retrieving service data")
	}

	// todo: utilize stats
	stats, err := op.kopia.BackupCollections(ctx, cs)
	if err != nil {
		return nil, errors.Wrap(err, "backing up service data")
	}

	op.Status = Successful
	return stats, nil
}
