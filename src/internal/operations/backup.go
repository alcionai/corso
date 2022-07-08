package operations

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/selectors"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation

	Results   BackupResults      `json:"results"`
	Selectors selectors.Selector `json:"selectors"`
	Version   string             `json:"version"`

	account account.Account
}

// BackupResults aggregate the details of the result of the operation.
type BackupResults struct {
	summary
	metrics
	Backup *backup.Backup `json:"backup"`
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts Options,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	acct account.Account,
	selector selectors.Selector,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, kw, ms),
		Selectors: selector,
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

// aggregates stats from the backup.Run().
// primarily used so that the defer can take in a
// pointer wrapping the values, while those values
// get populated asynchronously.
type backupStats struct {
	k                 *kopia.BackupStats
	gc                *support.ConnectorOperationStatus
	readErr, writeErr error
}

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) error {
	// TODO: persist initial state of backupOperation in modelstore

	// persist operation results to the model store on exit
	stats := backupStats{}
	defer op.persistResults(time.Now(), &stats)

	// retrieve data from the producer
	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		stats.readErr = err
		return errors.Wrap(err, "connecting to graph api")
	}

	var cs []connector.DataCollection
	cs, err = gc.ExchangeDataCollection(ctx, op.Selectors)
	if err != nil {
		stats.readErr = err
		return errors.Wrap(err, "retrieving service data")
	}
	stats.gc = gc.Status()

	// hand the results to the consumer
	var details *backup.Details
	stats.k, details, err = op.kopia.BackupCollections(ctx, cs)
	if err != nil {
		stats.writeErr = err
		return errors.Wrap(err, "backing up service data")
	}

	err = op.createBackupModels(ctx, stats.k.SnapshotID, details)
	if err != nil {
		stats.writeErr = err
		return err
	}
	return nil
}

func (op *BackupOperation) createBackupModels(ctx context.Context, snapID string, details *backup.Details) error {
	err := op.modelStore.Put(ctx, kopia.BackupDetailsModel, &details.DetailsModel)
	if err != nil {
		return errors.Wrap(err, "creating backupdetails model")
	}

	rp := backup.New(snapID, string(details.ModelStoreID))

	err = op.modelStore.Put(ctx, kopia.BackupModel, rp)
	if err != nil {
		return errors.Wrap(err, "creating backup model")
	}

	op.Results.Backup = rp

	return nil
}

// writes the backupOperation outcome to the modelStore.
func (op *BackupOperation) persistResults(
	started time.Time,
	stats *backupStats,
) {
	op.Status = Successful
	if stats.readErr != nil || stats.writeErr != nil {
		op.Status = Failed
	}

	op.Results.ReadErrors = stats.readErr
	op.Results.WriteErrors = stats.writeErr

	if stats.gc != nil {
		op.Results.ItemsRead = stats.gc.ObjectCount
	}
	if stats.k != nil {
		op.Results.ItemsWritten = stats.k.TotalFileCount
	}

	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	// TODO: persist operation to modelstore
}
