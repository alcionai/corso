package operations

import (
	"context"
	"fmt"
	"runtime/debug"
	"sort"
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
	cs                []data.RestoreCollection
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
		errs *fault.Errors,
	) ([]data.RestoreCollection, error)
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (restoreDetails *details.Details, err error) {
	defer func() {
		if r := recover(); r != nil {
			var rerr error
			if re, ok := r.(error); ok {
				rerr = re
			} else if re, ok := r.(string); ok {
				rerr = clues.New(re)
			} else {
				rerr = clues.New(fmt.Sprintf("%v", r))
			}

			err = clues.Wrap(rerr, "panic recovery").
				WithClues(ctx).
				With("stacktrace", string(debug.Stack()))
			logger.Ctx(ctx).
				With("err", err).
				Errorw("backup panic", clues.InErr(err).Slice()...)
		}
	}()

	var (
		opStats = restoreStats{
			bytesRead: &stats.ByteCounter{},
			restoreID: uuid.NewString(),
		}
		start        = time.Now()
		detailsStore = streamstore.New(op.kopia, op.account.ID(), op.Selectors.PathService())
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
		"service", op.Selectors.Service)

	// -----
	// Execution
	// -----

	deets, err := op.do(ctx, &opStats, detailsStore, start)
	if err != nil {
		// No return here!  We continue down to persistResults, even in case of failure.
		logger.Ctx(ctx).
			With("err", err).
			Errorw("doing restore", clues.InErr(err).Slice()...)
		op.Errors.Fail(errors.Wrap(err, "doing restore"))
		opStats.readErr = op.Errors.Err()
	}

	// -----
	// Persistence
	// -----

	err = op.persistResults(ctx, start, &opStats)
	if err != nil {
		op.Errors.Fail(errors.Wrap(err, "persisting restore results"))
		opStats.writeErr = op.Errors.Err()

		return nil, op.Errors.Err()
	}

	logger.Ctx(ctx).Infow("completed restore", "results", op.Results)

	return deets, nil
}

func (op *RestoreOperation) do(
	ctx context.Context,
	opStats *restoreStats,
	detailsStore detailsReader,
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

	paths, err := formatDetailsForRestoration(ctx, op.Selectors, deets, op.Errors)
	if err != nil {
		return nil, errors.Wrap(err, "formatting paths from details")
	}

	ctx = clues.Add(
		ctx,
		"resource_owner", bup.Selector.DiscreteOwner,
		"details_paths", len(paths))

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
	// TODO(keepers): remove when fault.Errors handles all iterable error aggregation.
	if opStats.gc.ErrorCount > 0 {
		return nil, opStats.gc.Err
	}

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

		// TODO(keepers): replace with fault.Errors handling.
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
		return errors.New("restoration never completed")
	}

	if opStats.gc.Successful == 0 {
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
	errs *fault.Errors,
) ([]path.Path, error) {
	fds, err := sel.Reduce(ctx, deets, errs)
	if err != nil {
		return nil, err
	}

	var (
		fdsPaths = fds.Paths()
		paths    = make([]path.Path, len(fdsPaths))
		et       = errs.Tracker()
	)

	for i := range fdsPaths {
		if errs.Err() != nil {
			break
		}

		p, err := path.FromDataLayerPath(fdsPaths[i], true)
		if err != nil {
			et.Add(clues.
				Wrap(err, "parsing details path after reduction").
				WithMap(clues.In(ctx)).
				With("path", fdsPaths[i]))

			continue
		}

		paths[i] = p
	}

	// TODO(meain): Move this to onedrive specific component, but as
	// of now the paths can technically be from multiple services

	// This sort is done primarily to order `.meta` files after `.data`
	// files. This is only a necessity for OneDrive as we are storing
	// metadata for files/folders in separate meta files and we the
	// data to be restored before we can restore the metadata.
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].String() < paths[j].String()
	})

	return paths, et.Err()
}
