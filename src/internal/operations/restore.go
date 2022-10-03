package operations

import (
	"context"
	"runtime/trace"
	"time"

	"github.com/google/uuid"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
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
	stats.Errs
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
	started           bool
	readErr, writeErr error

	// a transient value only used to pair up start-end events.
	restoreID string
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (restoreDetails *details.Details, err error) {
	defer trace.StartRegion(ctx, "operations:restore:run").End()

	var (
		parseErrs *multierror.Error
		opStats   = restoreStats{
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

	ds, b, err := op.store.GetDetailsFromBackupID(ctx, op.BackupID)
	if err != nil {
		err = errors.Wrap(err, "getting backup details for restore")
		opStats.readErr = err

		return nil, err
	}

	op.bus.Event(
		ctx,
		events.RestoreStart,
		map[string]any{
			events.StartTime:        startTime,
			events.BackupID:         op.BackupID,
			events.BackupCreateTime: b.CreationTime,
			events.RestoreID:        opStats.restoreID,
			// TODO: restore options,
		},
	)

	paths, err := formatDetailsForRestoration(ctx, op.Selectors, ds)
	if err != nil {
		opStats.readErr = err
		return nil, err
	}

	logger.Ctx(ctx).Infof("Discovered %d items in backup %s to restore", len(paths), op.BackupID)

	dcs, err := op.kopia.RestoreMultipleItems(ctx, b.SnapshotID, paths, opStats.bytesRead)
	if err != nil {
		err = errors.Wrap(err, "retrieving service data")

		parseErrs = multierror.Append(parseErrs, err)
		opStats.readErr = parseErrs.ErrorOrNil()

		return nil, err
	}

	opStats.readErr = parseErrs.ErrorOrNil()
	opStats.cs = dcs
	opStats.resourceCount = len(data.ResourceOwnerSet(dcs))

	// restore those collections using graph
	gc, err := connector.NewGraphConnector(ctx, op.account)
	if err != nil {
		err = errors.Wrap(err, "connecting to graph api")
		opStats.writeErr = err

		return nil, err
	}

	restoreDetails, err = gc.RestoreDataCollections(ctx, op.Selectors, op.Destination, dcs)
	if err != nil {
		err = errors.Wrap(err, "restoring service data")
		opStats.writeErr = err

		return nil, err
	}

	opStats.started = true
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

	if !opStats.started {
		op.Status = Failed

		return multierror.Append(
			errors.New("errors prevented the operation from processing"),
			opStats.readErr,
			opStats.writeErr)
	}

	op.Results.ReadErrors = opStats.readErr
	op.Results.WriteErrors = opStats.writeErr

	op.Results.BytesRead = opStats.bytesRead.NumBytes
	op.Results.ItemsRead = len(opStats.cs) // TODO: file count, not collection count
	op.Results.ItemsWritten = opStats.gc.Successful
	op.Results.ResourceOwners = opStats.resourceCount

	op.bus.Event(
		ctx,
		events.RestoreEnd,
		map[string]any{
			events.BackupID:      op.BackupID,
			events.DataRetrieved: op.Results.BytesRead,
			events.Duration:      op.Results.CompletedAt.Sub(op.Results.StartedAt),
			events.EndTime:       op.Results.CompletedAt,
			events.ItemsRead:     op.Results.ItemsRead,
			events.ItemsWritten:  op.Results.ItemsWritten,
			events.Resources:     op.Results.ResourceOwners,
			events.RestoreID:     opStats.restoreID,
			events.Service:       op.Selectors.Service.String(),
			events.StartTime:     op.Results.StartedAt,
			events.Status:        op.Status,
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
	var fds *details.Details

	switch sel.Service {
	case selectors.ServiceExchange:
		er, err := sel.ToExchangeRestore()
		if err != nil {
			return nil, err
		}

		// format the details and retrieve the items from kopia
		fds = er.Reduce(ctx, deets)
		if len(fds.Entries) == 0 {
			return nil, selectors.ErrorNoMatchingItems
		}

	case selectors.ServiceOneDrive:
		odr, err := sel.ToOneDriveRestore()
		if err != nil {
			return nil, err
		}

		// format the details and retrieve the items from kopia
		fds = odr.Reduce(ctx, deets)
		if len(fds.Entries) == 0 {
			return nil, selectors.ErrorNoMatchingItems
		}

	default:
		return nil, errors.Errorf("Service %s not supported", sel.Service)
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

	return paths, nil
}
