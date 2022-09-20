package operations

import (
	"context"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
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
	BackupID model.StableID `json:"backupID"`
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	selector selectors.Selector,
	bus events.Eventer,
) (BackupOperation, error) {
	op := BackupOperation{
		operation: newOperation(opts, bus, kw, sw),
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
	started           bool
	readErr, writeErr error
}

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	var (
		opStats       backupStats
		backupDetails *details.Details
		startTime     = time.Now()
	)
	// TODO: persist initial state of backupOperation in modelstore

	op.bus.Event(
		ctx,
		events.BackupStart,
		map[string]any{
			events.StartTime: startTime,
			// TODO: initial backup ID,
			// TODO: events.ExchangeResources: <count of resources>,
		},
	)

	// persist operation results to the model store on exit
	defer func() {
		err = op.persistResults(startTime, &opStats)
		if err != nil {
			return
		}

		err = op.createBackupModels(ctx, opStats.k.SnapshotID, backupDetails)
		if err != nil {
			// todo: we're not persisting this yet, except for the error shown to the user.
			opStats.writeErr = err
		}
	}()

	// retrieve data from the producer
	gc, err := connector.NewGraphConnector(ctx, op.account)
	if err != nil {
		err = errors.Wrap(err, "connecting to graph api")
		opStats.readErr = err

		return err
	}

	cs, err := gc.DataCollections(ctx, op.Selectors)
	if err != nil {
		err = errors.Wrap(err, "retrieving service data")
		opStats.readErr = err

		return err
	}

	// hand the results to the consumer
	opStats.k, backupDetails, err = op.kopia.BackupCollections(ctx, cs)
	if err != nil {
		err = errors.Wrap(err, "backing up service data")
		opStats.writeErr = err

		return err
	}

	opStats.started = true
	opStats.gc = gc.AwaitStatus()

	return err
}

// writes the results metrics to the operation results.
// later stored in the manifest using createBackupModels.
func (op *BackupOperation) persistResults(
	started time.Time,
	opStats *backupStats,
) error {
	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	op.Status = Completed
	if !opStats.started {
		op.Status = Failed

		return multierror.Append(
			errors.New("errors prevented the operation from processing"),
			opStats.readErr,
			opStats.writeErr)
	}

	op.Results.ReadErrors = opStats.readErr
	op.Results.WriteErrors = opStats.writeErr
	op.Results.ItemsRead = opStats.gc.Successful
	op.Results.ItemsWritten = opStats.k.TotalFileCount

	return nil
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	snapID string,
	backupDetails *details.Details,
) error {
	if backupDetails == nil {
		return errors.New("no backup details to record")
	}

	err := op.store.Put(ctx, model.BackupDetailsSchema, &backupDetails.DetailsModel)
	if err != nil {
		return errors.Wrap(err, "creating backupdetails model")
	}

	b := backup.New(
		snapID, string(backupDetails.ModelStoreID), op.Status.String(),
		op.Selectors,
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
	)

	err = op.store.Put(ctx, model.BackupSchema, b)
	if err != nil {
		return errors.Wrap(err, "creating backup model")
	}

	op.Results.BackupID = b.ID

	op.bus.Event(
		ctx,
		events.BackupEnd,
		map[string]any{
			events.BackupID:  b.ID,
			events.Status:    op.Status,
			events.StartTime: op.Results.StartedAt,
			events.EndTime:   op.Results.CompletedAt,
			events.Duration:  op.Results.CompletedAt.Sub(op.Results.StartedAt),
			// TODO: events.ExchangeResources: <count of resources>,
			// TODO: events.ExchangeDataObserved: <amount of data retrieved>,
			// TODO: events.ExchangeDataStored: <amount of data stored>,
		},
	)

	return nil
}
