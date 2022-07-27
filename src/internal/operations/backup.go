package operations

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/stats"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/store"
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
	stats.ReadWrites
	stats.StartAndEndTime
	BackupID model.ID `json:"backupID"`
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	selector selectors.Selector,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, kw, sw),
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
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	// TODO: persist initial state of backupOperation in modelstore

	// persist operation results to the model store on exit
	var (
		stats   backupStats
		details *details.Details
	)
	defer func() {
		op.persistResults(time.Now(), &stats)

		err = op.createBackupModels(ctx, stats.k.SnapshotID, details)
		if err != nil {
			stats.writeErr = err
			// todo: ^ we're not persisting this yet, except for the error shown to the user.
		}
	}()

	// retrieve data from the producer
	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		stats.readErr = err
		return errors.Wrap(err, "connecting to graph api")
	}

	var cs []data.DataCollection
	cs, err = gc.ExchangeDataCollection(ctx, op.Selectors)
	if err != nil {
		stats.readErr = err
		return errors.Wrap(err, "retrieving service data")
	}

	// hand the results to the consumer
	stats.k, details, err = op.kopia.BackupCollections(ctx, cs)
	if err != nil {
		stats.writeErr = err
		return errors.Wrap(err, "backing up service data")
	}
	stats.gc = gc.AwaitStatus()

	return err
}

// writes the results metrics to the operation results.
// later stored in the manifest using createBackupModels.
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
		op.Results.ItemsRead = stats.gc.Successful
	}
	if stats.k != nil {
		op.Results.ItemsWritten = stats.k.TotalFileCount
	}

	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(ctx context.Context, snapID string, details *details.Details) error {
	err := op.store.Put(ctx, model.BackupDetailsSchema, &details.DetailsModel)
	if err != nil {
		return errors.Wrap(err, "creating backupdetails model")
	}

	b := backup.New(
		snapID, string(details.ModelStoreID), op.Status.String(),
		op.Selectors,
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
	)
	err = op.store.Put(ctx, model.BackupSchema, b)
	if err != nil {
		return errors.Wrap(err, "creating backup model")
	}
	op.Results.BackupID = b.StableID

	return nil
}
