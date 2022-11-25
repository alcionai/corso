package operations

import (
	"context"
	"time"

	"github.com/google/uuid"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
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
	stats.Errs
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
	resourceCount     int
	started           bool
	readErr, writeErr error
}

// Run begins a synchronous backup operation.
func (op *BackupOperation) Run(ctx context.Context) (err error) {
	ctx, end := D.Span(ctx, "operations:backup:run")
	defer end()

	var (
		opStats       backupStats
		backupDetails *details.Details
		startTime     = time.Now()
	)

	op.Results.BackupID = model.StableID(uuid.NewString())

	op.bus.Event(
		ctx,
		events.BackupStart,
		map[string]any{
			events.StartTime: startTime,
			events.Service:   op.Selectors.Service.String(),
			events.BackupID:  op.Results.BackupID,
		},
	)

	// persist operation results to the model store on exit
	defer func() {
		// wait for the progress display to clean up
		observe.Complete()

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

	complete, closer := observe.MessageWithCompletion("Connecting to M365:")
	defer closer()
	defer close(complete)

	// retrieve data from the producer
	resource := connector.Users
	if op.Selectors.Service == selectors.ServiceSharePoint {
		resource = connector.Sites
	}

	gc, err := connector.NewGraphConnector(ctx, op.account, resource)
	if err != nil {
		err = errors.Wrap(err, "connecting to graph api")
		opStats.readErr = err

		return err
	}
	complete <- struct{}{}

	discoverCh, closer := observe.MessageWithCompletion("Discovering items to backup:")
	defer closer()
	defer close(discoverCh)

	cs, err := gc.DataCollections(ctx, op.Selectors)
	if err != nil {
		err = errors.Wrap(err, "retrieving service data")
		opStats.readErr = err

		return err
	}

	discoverCh <- struct{}{}

	opStats.resourceCount = len(data.ResourceOwnerSet(cs))

	backupCh, closer := observe.MessageWithCompletion("Backing up data:")
	defer closer()
	defer close(backupCh)

	// hand the results to the consumer
	opStats.k, backupDetails, err = op.kopia.BackupCollections(ctx, cs, op.Selectors.PathService())
	if err != nil {
		err = errors.Wrap(err, "backing up service data")
		opStats.writeErr = err

		return err
	}
	backupCh <- struct{}{}

	logger.Ctx(ctx).Debugf("Backed up %d directories and %d files", opStats.k.TotalDirectoryCount, opStats.k.TotalFileCount)

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

	if opStats.readErr == nil && opStats.writeErr == nil && opStats.gc.Successful == 0 {
		op.Status = NoData
	}

	op.Results.ReadErrors = opStats.readErr
	op.Results.WriteErrors = opStats.writeErr

	op.Results.BytesRead = opStats.k.TotalHashedBytes
	op.Results.BytesUploaded = opStats.k.TotalUploadedBytes
	op.Results.ItemsRead = opStats.gc.Successful
	op.Results.ItemsWritten = opStats.k.TotalFileCount
	op.Results.ResourceOwners = opStats.resourceCount

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
		op.Results.BackupID,
		op.Selectors,
		op.Results.ReadWrites,
		op.Results.StartAndEndTime,
	)

	err = op.store.Put(ctx, model.BackupSchema, b)
	if err != nil {
		return errors.Wrap(err, "creating backup model")
	}

	op.bus.Event(
		ctx,
		events.BackupEnd,
		map[string]any{
			events.BackupID:   b.ID,
			events.DataStored: op.Results.BytesUploaded,
			events.Duration:   op.Results.CompletedAt.Sub(op.Results.StartedAt),
			events.EndTime:    op.Results.CompletedAt,
			events.Resources:  op.Results.ResourceOwners,
			events.Service:    op.Selectors.PathService().String(),
			events.StartTime:  op.Results.StartedAt,
			events.Status:     op.Status,
		},
	)

	return nil
}
