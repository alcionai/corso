package operations

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/account"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation

	RestorePointID string         `json:"restorePointID"`
	Results        RestoreResults `json:"results"`
	Targets        []string       `json:"selectors"` // todo: replace with Selectors
	Version        string         `json:"bersion"`

	account account.Account
}

// RestoreResults aggregate the details of the results of the operation.
type RestoreResults struct {
	summary
	metrics
}

// NewRestoreOperation constructs and validates a restore operation.
func NewRestoreOperation(
	ctx context.Context,
	opts Options,
	kw *kopia.Wrapper,
	acct account.Account,
	restorePointID string,
	targets []string,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation:      newOperation(opts, kw),
		RestorePointID: restorePointID,
		Targets:        targets,
		Version:        "v0",
		account:        acct,
	}
	if err := op.validate(); err != nil {
		return RestoreOperation{}, err
	}

	return op, nil
}

func (op RestoreOperation) validate() error {
	return op.operation.validate()
}

// aggregates stats from the restore.Run().
// primarily used so that the defer can take in a
// pointer wrapping the values, while those values
// get populated asynchronously.
type restoreStats struct {
	cs                []connector.DataCollection
	gc                *support.ConnectorOperationStatus
	readErr, writeErr error
}

// Run begins a synchronous restore operation.
// todo (keepers): return stats block in first param.
func (op *RestoreOperation) Run(ctx context.Context) error {
	// TODO: persist initial state of restoreOperation in modelstore

	// persist operation results to the model store on exit
	stats := restoreStats{}
	defer op.persistResults(time.Now(), &stats)

	dc, err := op.kopia.RestoreSingleItem(ctx, op.RestorePointID, op.Targets)
	if err != nil {
		stats.readErr = err
		return errors.Wrap(err, "retrieving service data")
	}
	stats.cs = []connector.DataCollection{dc}

	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		stats.writeErr = err
		return errors.Wrap(err, "connecting to graph api")
	}

	if err := gc.RestoreMessages(ctx, dc); err != nil {
		stats.writeErr = err
		return errors.Wrap(err, "restoring service data")
	}

	op.Status = Successful
	return nil
}

// writes the restoreOperation outcome to the modelStore.
func (op *RestoreOperation) persistResults(
	started time.Time,
	stats *restoreStats,
) {
	op.Status = Successful
	if stats.readErr != nil || stats.writeErr != nil {
		op.Status = Failed
	}
	op.Results.ReadErrors = stats.readErr
	op.Results.WriteErrors = stats.writeErr

	op.Results.ItemsRead = len(stats.cs) // TODO: file count, not collection count

	if stats.gc != nil {
		op.Results.ItemsWritten = stats.gc.ObjectCount
	}

	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	// TODO: persist operation to modelstore
}
