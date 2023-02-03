package operations

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
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
	stats.Errs // deprecated in place of fault.Errors in the base operation.
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
	cs                []data.Collection
	gc                *support.ConnectorOperationStatus
	bytesRead         *stats.ByteCounter
	resourceCount     int
	readErr, writeErr error

	// a transient value only used to pair up start-end events.
	restoreID string
}

type restorer interface {
	RestoreMultipleItems(
		ctx context.Context,
		snapshotID string,
		paths []path.Path,
		bc kopia.ByteCounter,
	) ([]data.Collection, error)
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (restoreDetails *details.Details, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = clues.Wrap(r.(error), "panic recovery").WithClues(ctx).With("stacktrace", debug.Stack())
		}
	}()

	ctx, end := D.Span(ctx, "operations:restore:run")
	defer func() {
		end()
		// wait for the progress display to clean up
		observe.Complete()
	}()

	ctx = clues.AddAll(
		ctx,
		"tenant_id", op.account.ID(), // TODO: pii
		"backup_id", op.BackupID,
		"service", op.Selectors.Service)

	deets, err := op.do(ctx)
	if err != nil {
		logger.Ctx(ctx).
			With("err", err).
			Errorw("restore operation", clues.InErr(err).Slice()...)

		return nil, err
	}

	logger.Ctx(ctx).Infow("completed restore", "results", op.Results)

	return deets, nil
}

func (op *RestoreOperation) do(ctx context.Context) (restoreDetails *details.Details, err error) {
	var (
		opStats = restoreStats{
			bytesRead: &stats.ByteCounter{},
			restoreID: uuid.NewString(),
		}
		startTime = time.Now()
	)

	defer func() {
		err = op.persistResults(ctx, startTime, &opStats)
		if err != nil {
			return
		}
	}()

	detailsStore := streamstore.New(op.kopia, op.account.ID(), op.Selectors.PathService())

	bup, deets, err := getBackupAndDetailsFromID(
		ctx,
		op.BackupID,
		op.store,
		detailsStore,
	)
	if err != nil {
		opStats.readErr = errors.Wrap(err, "restore")
		return nil, opStats.readErr
	}

	ctx = clues.Add(ctx, "resource_owner", bup.Selector.DiscreteOwner)

	op.bus.Event(
		ctx,
		events.RestoreStart,
		map[string]any{
			events.StartTime:        startTime,
			events.BackupID:         op.BackupID,
			events.BackupCreateTime: bup.CreationTime,
			events.RestoreID:        opStats.restoreID,
		},
	)

	paths, err := formatDetailsForRestoration(ctx, op.Selectors, deets)
	if err != nil {
		opStats.readErr = err
		return nil, err
	}

	ctx = clues.Add(ctx, "details_paths", len(paths))
	observe.Message(ctx, observe.Safe(fmt.Sprintf("Discovered %d items in backup %s to restore", len(paths), op.BackupID)))

	kopiaComplete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Enumerating items in repository"))
	defer closer()
	defer close(kopiaComplete)

	dcs, err := op.kopia.RestoreMultipleItems(ctx, bup.SnapshotID, paths, opStats.bytesRead)
	if err != nil {
		opStats.readErr = errors.Wrap(err, "retrieving service data")
		return nil, opStats.readErr
	}
	kopiaComplete <- struct{}{}

	ctx = clues.Add(ctx, "coll_count", len(dcs))
	opStats.cs = dcs
	opStats.resourceCount = len(data.ResourceOwnerSet(dcs))

	gc, err := connectToM365(ctx, op.Selectors, op.account)
	if err != nil {
		opStats.readErr = errors.Wrap(err, "connecting to M365")
		return nil, opStats.readErr
	}

	restoreComplete, closer := observe.MessageWithCompletion(ctx, observe.Safe("Restoring data"))
	defer closer()
	defer close(restoreComplete)

	restoreDetails, err = gc.RestoreDataCollections(
		ctx,
		op.account,
		op.Selectors,
		op.Destination,
		dcs)
	if err != nil {
		opStats.writeErr = errors.Wrap(err, "restoring service data")
		return nil, opStats.writeErr
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
	op.Results.ReadErrors = opStats.readErr
	op.Results.WriteErrors = opStats.writeErr

	op.Status = Completed

	if opStats.readErr != nil || opStats.writeErr != nil {
		op.Status = Failed

		return multierror.Append(
			errors.New("errors prevented the operation from processing"),
			opStats.readErr,
			opStats.writeErr)
	}

	op.Results.BytesRead = opStats.bytesRead.NumBytes
	op.Results.ItemsRead = len(opStats.cs) // TODO: file count, not collection count
	op.Results.ResourceOwners = opStats.resourceCount

	if opStats.gc == nil {
		op.Status = Failed
		return errors.New("data restoration never completed")
	}

	if opStats.readErr == nil && opStats.writeErr == nil && opStats.gc.Successful == 0 {
		op.Status = NoData
	}

	op.Results.ItemsWritten = opStats.gc.Successful

	dur := op.Results.CompletedAt.Sub(op.Results.StartedAt)

	op.bus.Event(
		ctx,
		events.RestoreEnd,
		map[string]any{
			events.BackupID:      op.BackupID,
			events.DataRetrieved: op.Results.BytesRead,
			events.Duration:      dur,
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

	return nil
}

// formatDetailsForRestoration reduces the provided detail entries according to the
// selector specifications.
func formatDetailsForRestoration(
	ctx context.Context,
	sel selectors.Selector,
	deets *details.Details,
) ([]path.Path, error) {
	fds, err := sel.Reduce(ctx, deets)
	if err != nil {
		return nil, err
	}

	var (
		errs     *multierror.Error
		fdsPaths = fds.Paths()
		paths    = make([]path.Path, len(fdsPaths))
	)

	for i := range fdsPaths {
		p, err := path.FromDataLayerPath(fdsPaths[i], true)
		if err != nil {
			errs = multierror.Append(
				errs,
				errors.Wrap(err, "parsing details entry path"),
			)

			continue
		}

		paths[i] = p
	}

	if errs != nil {
		return nil, errs
	}

	return paths, nil
}
