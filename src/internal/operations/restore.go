package operations

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation

	BackupID    model.StableID             `json:"backupID"`
	Results     RestoreResults             `json:"results"`
	Selectors   selectors.Selector         `json:"selectors"`
	Destination control.RestoreDestination `json:"destination"`
	Version     string                     `json:"version"`

	account account.Account
}

// RestoreResults aggregate the details of the results of the operation.
type RestoreResults struct {
	stats.ReadWrites
	stats.StartAndEndTime
}

// NewRestoreOperation constructs and validates a restore operation.
func NewRestoreOperation(
	ctx context.Context,
	opts control.Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	backupID model.StableID,
	sel selectors.Selector,
	dest control.RestoreDestination,
	bus events.Eventer,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation:   newOperation(opts, bus, kw, sw),
		BackupID:    backupID,
		Selectors:   sel,
		Destination: dest,
		Version:     "v0",
		account:     acct,
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
	cs            []data.RestoreCollection
	gc            *support.ConnectorOperationStatus
	bytesRead     *stats.ByteCounter
	resourceCount int

	// a transient value only used to pair up start-end events.
	restoreID string
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (restoreDetails *details.Details, err error) {
	defer func() {
		if crErr := crash.Recovery(ctx, recover()); crErr != nil {
			err = crErr
		}
	}()

	var (
		opStats = restoreStats{
			bytesRead: &stats.ByteCounter{},
			restoreID: uuid.NewString(),
		}
		start  = time.Now()
		sstore = streamstore.NewStreamer(op.kopia, op.account.ID(), op.Selectors.PathService())
	)

	// -----
	// Setup
	// -----

	ctx, end := D.Span(ctx, "operations:restore:run")
	defer func() {
		end()
		// wait for the progress display to clean up
		observe.Complete()
	}()

	ctx = clues.Add(
		ctx,
		"tenant_id", op.account.ID(), // TODO: pii
		"backup_id", op.BackupID,
		"service", op.Selectors.Service,
		"destination_container", op.Destination.ContainerName)

	// -----
	// Execution
	// -----

	deets, err := op.do(ctx, &opStats, sstore, start)
	if err != nil {
		// No return here!  We continue down to persistResults, even in case of failure.
		logger.Ctx(ctx).
			With("err", err).
			Errorw("running restore", clues.InErr(err).Slice()...)
		op.Errors.Fail(errors.Wrap(err, "running restore"))
	}

	LogFaultErrors(ctx, op.Errors.Errors(), "running restore")

	// -----
	// Persistence
	// -----

	err = op.persistResults(ctx, start, &opStats)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "persisting restore results"))
		return nil, op.Errors.Failure()
	}

	logger.Ctx(ctx).Infow("completed restore", "results", op.Results)

	return deets, nil
}

func (op *RestoreOperation) do(
	ctx context.Context,
	opStats *restoreStats,
	detailsStore streamstore.Reader,
	start time.Time,
) (*details.Details, error) {
	bup, deets, err := getBackupAndDetailsFromID(
		ctx,
		op.BackupID,
		op.store,
		detailsStore,
		op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "getting backup and details")
	}

	observe.Message(ctx, observe.Safe("Restoring"), observe.Bullet, observe.PII(bup.Selector.DiscreteOwner))

	paths, err := formatDetailsForRestoration(ctx, bup.Version, op.Selectors, deets, op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "formatting paths from details")
	}

	ctx = clues.Add(
		ctx,
		"resource_owner", bup.Selector.DiscreteOwner,
		"details_paths", len(paths),
		"backup_snapshot_id", bup.SnapshotID)

	op.bus.Event(
		ctx,
		events.RestoreStart,
		map[string]any{
			events.StartTime:        start,
			events.BackupID:         op.BackupID,
			events.BackupCreateTime: bup.CreationTime,
			events.RestoreID:        opStats.restoreID,
		})

	observe.Message(ctx, observe.Safe(fmt.Sprintf("Discovered %d items in backup %s to restore", len(paths), op.BackupID)))
	logger.Ctx(ctx).With("selectors", op.Selectors).Info("restoring selection")

	kopiaComplete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Enumerating items in repository"))
	defer closer()
	defer close(kopiaComplete)

	dcs, err := op.kopia.RestoreMultipleItems(ctx, bup.SnapshotID, paths, opStats.bytesRead, op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving collections from repository")
	}

	kopiaComplete <- struct{}{}

	ctx = clues.Add(ctx, "coll_count", len(dcs))

	// should always be 1, since backups are 1:1 with resourceOwners.
	opStats.resourceCount = 1
	opStats.cs = dcs

	gc, err := connectToM365(ctx, op.Selectors, op.account, op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to M365")
	}

	restoreComplete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Restoring data"))
	defer closer()
	defer close(restoreComplete)

	restoreDetails, err := gc.RestoreDataCollections(
		ctx,
		bup.Version,
		op.account,
		op.Selectors,
		op.Destination,
		op.Options,
		dcs,
		op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "restoring collections")
	}

	restoreComplete <- struct{}{}

	opStats.gc = gc.AwaitStatus()

	logger.Ctx(ctx).Debug(gc.PrintableStatus())

	return restoreDetails, nil
}

// persists details and statistics about the restore operation.
func (op *RestoreOperation) persistResults(
	ctx context.Context,
	started time.Time,
	opStats *restoreStats,
) error {
	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	op.Status = Completed

	if op.Errors.Failure() != nil {
		op.Status = Failed
	}

	op.Results.BytesRead = opStats.bytesRead.NumBytes
	op.Results.ItemsRead = len(opStats.cs) // TODO: file count, not collection count
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.gc == nil {
		op.Status = Failed
		return errors.New("restoration never completed")
	}

	if op.Status != Failed && opStats.gc.Metrics.Successes == 0 {
		op.Status = NoData
	}

	op.Results.ItemsWritten = opStats.gc.Metrics.Successes

	op.bus.Event(
		ctx,
		events.RestoreEnd,
		map[string]any{
			events.BackupID:      op.BackupID,
			events.DataRetrieved: op.Results.BytesRead,
			events.Duration:      op.Results.CompletedAt.Sub(op.Results.StartedAt),
			events.EndTime:       common.FormatTime(op.Results.CompletedAt),
			events.ItemsRead:     op.Results.ItemsRead,
			events.ItemsWritten:  op.Results.ItemsWritten,
			events.Resources:     op.Results.ResourceOwners,
			events.RestoreID:     opStats.restoreID,
			events.Service:       op.Selectors.Service.String(),
			events.StartTime:     common.FormatTime(op.Results.StartedAt),
			events.Status:        op.Status.String(),
		},
	)

	return op.Errors.Failure()
}

// formatDetailsForRestoration reduces the provided detail entries according to the
// selector specifications.
func formatDetailsForRestoration(
	ctx context.Context,
	backupVersion int,
	sel selectors.Selector,
	deets *details.Details,
	errs *fault.Bus,
) ([]path.Path, error) {
	fds, err := sel.Reduce(ctx, deets, errs)
	if err != nil {
		return nil, err
	}

	var (
		fdsPaths  = fds.Paths()
		paths     = make([]path.Path, len(fdsPaths))
		shortRefs = make([]string, len(fdsPaths))
		el        = errs.Local()
	)

	for i := range fdsPaths {
		if el.Failure() != nil {
			break
		}

		p, err := path.FromDataLayerPath(fdsPaths[i], true)
		if err != nil {
			el.AddRecoverable(clues.
				Wrap(err, "parsing details path after reduction").
				WithMap(clues.In(ctx)).
				With("path", fdsPaths[i]))

			continue
		}

		paths[i] = p
		shortRefs[i] = p.ShortRef()
	}

	if sel.Service == selectors.ServiceOneDrive {
		paths, err = onedrive.AugmentRestorePaths(backupVersion, paths)
		if err != nil {
			return nil, clues.Wrap(err, "augmenting paths")
		}
	}

	logger.Ctx(ctx).With("short_refs", shortRefs).Infof("found %d details entries to restore", len(shortRefs))

	return paths, el.Failure()
}
