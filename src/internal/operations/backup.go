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
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

var (
	errBackupNotStarted = errors.New("errors prevented the operation from processing")
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation

	BackupIDs []model.StableID
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
	Status   opStatus
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

	// wait for the progress display to clean up
	defer observe.Complete()

	complete, closer := observe.MessageWithCompletion("Connecting to M365:")
	defer closer()
	defer close(complete)

	// retrieve data from the producer
	gc, err := connector.NewGraphConnector(ctx, op.account)
	if err != nil {
		err = errors.Wrap(err, "connecting to graph api")

		// Errors here don't cause anything to be persisted.
		return multierror.Append(
			errBackupNotStarted,
			err,
		)
	}
	complete <- struct{}{}

	discoverCh, closer := observe.MessageWithCompletion("Discovering items to backup:")
	defer closer()
	defer close(discoverCh)

	cs, err := gc.DataCollections(ctx, op.Selectors)
	if err != nil {
		err = errors.Wrap(err, "retrieving service data")
		return err
	}

	discoverCh <- struct{}{}

	backupCh, closer := observe.MessageWithCompletion("Backing up data:")
	defer closer()
	defer close(backupCh)

	// Group items by (resource owner, service, category) tuple and create a
	// separate backup for each tuple.
	groups := map[string][]data.Collection{}

	for _, c := range cs {
		fp := c.FullPath()
		idx := fp.ResourceOwner() + fp.Service().String() + fp.Category().String()

		groups[idx] = append(groups[idx], c)
	}

	allFailed := true
	allNoData := true

	for _, cols := range groups {
		results := op.backupCollectionsGroup(ctx, cols)

		if results.Status != Failed {
			op.BackupIDs = append(op.BackupIDs, results.BackupID)
		}

		allFailed = allFailed && results.Status == Failed
		allNoData = allNoData && results.Status == NoData
	}

	backupCh <- struct{}{}

	// Wait at the end so any individual backup doesn't deadlock. Seems like the
	// total items read is not reported or persisted anywhere.
	gcStats := gc.AwaitStatus()

	// Check before checking for failures so we pick up when the loop doesn't run
	// because there were no collections.
	if allNoData && gcStats.Successful == 0 {
		op.Status = NoData
	} else if allFailed {
		op.Status = Failed
	} else {
		op.Status = Completed
	}

	return err
}

// backupCollectionsGroup backs up a single set of collections and makes a
// BackupModel for the collections. Each set should correspond to a single
// (resource owner, service, category) tuple.
func (op *BackupOperation) backupCollectionsGroup(
	ctx context.Context,
	cs []data.Collection,
) *BackupResults {
	var (
		err           error
		opStats       backupStats
		backupDetails *details.Details
		results       = &BackupResults{}
		startTime     = time.Now()
		hadData       = len(cs) > 0
	)

	// persist operation results to the model store on exit
	// TODO(ashmrtn): Find some error handling method for this.
	defer func() {
		err = op.persistResults(startTime, &opStats, hadData, results)
		if err != nil {
			return
		}

		err = op.createBackupModels(ctx, results, opStats.k.SnapshotID, backupDetails)
		if err != nil {
			// todo: we're not persisting this yet, except for the error shown to the user.
			opStats.writeErr = err
		}
	}()

	results.BackupID = model.StableID(uuid.NewString())

	op.bus.Event(
		ctx,
		events.BackupStart,
		map[string]any{
			events.StartTime: startTime,
			events.Service:   op.Selectors.Service.String(),
			events.BackupID:  results.BackupID,
		},
	)

	opStats.resourceCount = len(cs)

	// hand the results to the consumer
	opStats.k, backupDetails, err = op.kopia.BackupCollections(ctx, cs, op.Selectors.PathService())
	if err != nil {
		err = errors.Wrap(err, "backing up service data")
		opStats.writeErr = err

		return results
	}

	opStats.started = true

	return results
}

// writes the results metrics to the operation results.
// later stored in the manifest using createBackupModels.
func (op *BackupOperation) persistResults(
	started time.Time,
	opStats *backupStats,
	hadData bool,
	results *BackupResults,
) error {
	results.StartedAt = started
	results.CompletedAt = time.Now()

	op.Status = Completed
	if !opStats.started {
		results.Status = Failed

		return multierror.Append(
			errBackupNotStarted,
			opStats.readErr,
			opStats.writeErr)
	}

	if opStats.readErr == nil && opStats.writeErr == nil && !hadData {
		results.Status = NoData
	}

	results.ReadErrors = opStats.readErr
	results.WriteErrors = opStats.writeErr

	results.BytesRead = opStats.k.TotalHashedBytes
	results.BytesUploaded = opStats.k.TotalUploadedBytes
	// TODO(ashmrtn): Bring back if gc starts returning stats associated with each
	// tuple.
	//results.ItemsRead = opStats.gc.Successful
	results.ItemsWritten = opStats.k.TotalFileCount
	results.ResourceOwners = opStats.resourceCount

	return nil
}

// stores the operation details, results, and selectors in the backup manifest.
func (op *BackupOperation) createBackupModels(
	ctx context.Context,
	results *BackupResults,
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
		snapID, string(backupDetails.ModelStoreID), results.Status.String(),
		results.BackupID,
		op.Selectors,
		results.ReadWrites,
		results.StartAndEndTime,
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
			events.DataStored: results.BytesUploaded,
			events.Duration:   results.CompletedAt.Sub(results.StartedAt),
			events.EndTime:    results.CompletedAt,
			events.Resources:  results.ResourceOwners,
			events.Service:    op.Selectors.PathService().String(),
			events.StartTime:  results.StartedAt,
			events.Status:     op.Status,
		},
	)

	return nil
}
