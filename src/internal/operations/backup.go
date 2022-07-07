package operations

import (
	"context"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/restorepoint"
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
	summary
	metrics
	// todo: RestorePoint RestorePoint
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts Options,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	acct account.Account,
	targets []string,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, kw, ms),
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
	// TODO: persist initial state of backupOperation in modelstore

	var (
		cs                []connector.DataCollection
		stats             = &kopia.BackupStats{}
		readErr, writeErr error
	)

	// persist operation results to the model store on exit
	defer op.persistResults(time.Now(), cs, stats, readErr, writeErr)

	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		readErr = multierror.Append(readErr, err)
		return nil, errors.Wrap(err, "connecting to graph api")
	}

	cs, err = gc.ExchangeDataCollection(ctx, op.Targets[0])
	if err != nil {
		readErr = multierror.Append(readErr, err)
		return nil, errors.Wrap(err, "retrieving service data")
	}

	stats, details, writeErr := op.kopia.BackupCollections(ctx, cs)
	if writeErr != nil {
		return nil, errors.Wrap(err, "backing up service data")
	}

	writeErr = op.createRestorePoint(ctx, stats.SnapshotID, details)
	if writeErr != nil {
		return nil, writeErr
	}

	return stats, nil
}

func (op *BackupOperation) createRestorePoint(ctx context.Context, snapID string, details *restorepoint.Details) error {
	err := op.modelStore.Put(ctx, kopia.RestorePointDetailsModel, details)
	if err != nil {
		return errors.Wrap(err, "creating restorepointdetails model")
	}

	err = op.modelStore.Put(ctx, kopia.RestorePointModel,
		restorepoint.New(snapID, string(details.ModelStoreID)))
	if err != nil {
		return errors.Wrap(err, "creating restorepoint model")
	}
	return nil
}

// writes the backupOperation outcome to the modelStore.
func (op *BackupOperation) persistResults(
	started time.Time,
	cs []connector.DataCollection,
	stats *kopia.BackupStats,
	readErr, writeErr error,
) {
	op.Status = Successful
	if readErr != nil || writeErr != nil {
		op.Status = Failed
	}

	op.Results.ItemsRead = len(cs) // TODO: file count, not collection count
	op.Results.ReadErrors = readErr
	op.Results.ItemsWritten = stats.TotalFileCount
	op.Results.WriteErrors = writeErr

	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	// TODO: persist operation to modelstore
}
