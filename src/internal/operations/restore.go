package operations

import (
	"context"
	"strings"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/stats"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/store"
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
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation: newOperation(opts, kw, sw),
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
	started           bool
	readErr, writeErr error
}

// Run begins a synchronous restore operation.
func (op *RestoreOperation) Run(ctx context.Context) (err error) {
	// TODO: persist initial state of restoreOperation in modelstore
	// persist operation results to the model store on exit
	opStats := restoreStats{}
	// TODO: persist results?
	defer func() {
		err = op.persistResults(time.Now(), &opStats)
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

	er, err := op.Selectors.ToExchangeRestore()
	if err != nil {
		opStats.readErr = err
		return err
	}

	// format the details and retrieve the items from kopia
	fds := er.Reduce(d)
	if len(fds.Entries) == 0 {
		return errors.New("nothing to restore: no items in the backup match the provided selectors")
	}

	// todo: use path pkg for this
	fdsPaths := fds.Paths()
	paths := make([][]string, len(fdsPaths))

	for i := range fdsPaths {
		paths[i] = strings.Split(fdsPaths[i], "/")
	}

	dcs, err := op.kopia.RestoreMultipleItems(ctx, b.SnapshotID, paths)
	if err != nil {
		err = errors.Wrap(err, "retrieving service data")
		opStats.readErr = err

		return err
	}

	opStats.cs = dcs

	// restore those collections using graph
	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		err = errors.Wrap(err, "connecting to graph api")
		opStats.writeErr = err

		return err
	}

	err = gc.RestoreMessages(ctx, dcs)
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

	return nil
}
