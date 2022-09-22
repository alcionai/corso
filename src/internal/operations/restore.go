package operations

import (
	"context"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation

	BackupID  model.StableID     `json:"backupID"`
	Results   RestoreResults     `json:"results"`
	Selectors selectors.Selector `json:"selectors"` // todo: replace with Selectors
	Version   string             `json:"version"`

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
	bus events.Eventer,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation: newOperation(opts, bus, kw, sw),
		BackupID:  backupID,
		Selectors: sel,
		Version:   "v0",
		account:   acct,
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
	resourceCount     int
	started           bool
	readErr, writeErr error
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (err error) {
	startTime := time.Now()

	// TODO: persist initial state of restoreOperation in modelstore
	op.bus.Event(
		ctx,
		events.RestoreStart,
		map[string]any{
			events.StartTime: startTime,
			events.Service:   op.Selectors.Service.String(),
			events.BackupID:  op.BackupID,
			// TODO: initial backup ID,
			// TODO: events.ExchangeResources: <count of resources>,
			// TODO: source backup time,
			// TODO: restore options,
		},
	)

	// persist operation results to the model store on exit
	opStats := restoreStats{}
	// TODO: persist results?

	defer func() {
		err = op.persistResults(ctx, startTime, &opStats)
		if err != nil {
			return
		}
	}()

	// retrieve the restore point details
	d, b, err := op.store.GetDetailsFromBackupID(ctx, op.BackupID)
	if err != nil {
		err = errors.Wrap(err, "getting backup details for restore")
		opStats.readErr = err

		return err
	}

	var fds *details.Details

	switch op.Selectors.Service {
	case selectors.ServiceExchange:
		er, err := op.Selectors.ToExchangeRestore()
		if err != nil {
			opStats.readErr = err
			return err
		}

		// format the details and retrieve the items from kopia
		fds = er.Reduce(ctx, d)
		if len(fds.Entries) == 0 {
			return selectors.ErrorNoMatchingItems
		}

	case selectors.ServiceOneDrive:
		// TODO: Reduce `details` here when we add support for OneDrive restore filters
		fds = d
	default:
		return errors.Errorf("Service %s not supported", op.Selectors.Service)
	}

	fdsPaths := fds.Paths()
	paths := make([]path.Path, len(fdsPaths))

	var parseErrs *multierror.Error

	for i := range fdsPaths {
		p, err := path.FromDataLayerPath(fdsPaths[i], true)
		if err != nil {
			parseErrs = multierror.Append(
				parseErrs,
				errors.Wrap(err, "parsing details entry path"),
			)

			continue
		}

		paths[i] = p
	}

	dcs, err := op.kopia.RestoreMultipleItems(ctx, b.SnapshotID, paths)
	if err != nil {
		err = errors.Wrap(err, "retrieving service data")

		parseErrs = multierror.Append(parseErrs, err)
		opStats.readErr = parseErrs.ErrorOrNil()

		return err
	}

	opStats.readErr = parseErrs.ErrorOrNil()
	opStats.cs = dcs
	opStats.resourceCount = len(data.ResourceOwnerSet(dcs))

	// restore those collections using graph
	gc, err := connector.NewGraphConnector(ctx, op.account)
	if err != nil {
		err = errors.Wrap(err, "connecting to graph api")
		opStats.writeErr = err

		return err
	}

	err = gc.RestoreDataCollections(ctx, op.Selectors, dcs)
	if err != nil {
		err = errors.Wrap(err, "restoring service data")
		opStats.writeErr = err

		return err
	}

	opStats.started = true
	opStats.gc = gc.AwaitStatus()
	logger.Ctx(ctx).Debug(gc.PrintableStatus())

	return nil
}

// writes the restoreOperation outcome to the modelStore.
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
	op.Results.ItemsRead = len(opStats.cs) // TODO: file count, not collection count
	op.Results.ItemsWritten = opStats.gc.Successful
	op.Results.ResourceOwners = opStats.resourceCount

	op.bus.Event(
		ctx,
		events.RestoreEnd,
		map[string]any{
			// TODO: RestoreID
			events.BackupID:     op.BackupID,
			events.Service:      op.Selectors.Service.String(),
			events.Status:       op.Status,
			events.StartTime:    op.Results.StartedAt,
			events.EndTime:      op.Results.CompletedAt,
			events.Duration:     op.Results.CompletedAt.Sub(op.Results.StartedAt),
			events.ItemsRead:    op.Results.ItemsRead,
			events.ItemsWritten: op.Results.ItemsWritten,
			events.Resources:    op.Results.ResourceOwners,
			// TODO: events.ExchangeDataObserved: <amount of data retrieved>,
		},
	)

	return nil
}
