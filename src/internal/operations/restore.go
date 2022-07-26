package operations

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/data"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/stats"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/selectors"
	"github.com/alcionai/corso/pkg/store"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation

	BackupID  model.ID           `json:"backupID"`
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
	opts Options,
	kw *kopia.Wrapper,
	sw *store.Wrapper,
	acct account.Account,
	backupID model.ID,
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
	cs                []data.DataCollection
	gc                *support.ConnectorOperationStatus
	readErr, writeErr error
}

// Run begins a synchronous restore operation.
// todo (keepers): return stats block in first param.
func (op *RestoreOperation) Run(ctx context.Context) error {
	// TODO: persist initial state of restoreOperation in modelstore

	// persist operation results to the model store on exit
	stats := restoreStats{}
	defer op.persistResults(time.Now(), &stats)

	// retrieve the restore point details
	d, b, err := op.store.GetDetailsFromBackupID(ctx, op.BackupID)
	if err != nil {
		stats.readErr = errors.Wrap(err, "getting backup details for restore")
		return stats.readErr
	}

	er, err := op.Selectors.ToExchangeRestore()
	if err != nil {
		stats.readErr = err
		return err
	}

	// format the details and retrieve the items from kopia
	fds := er.Reduce(d)
	// todo: use path pkg for this
	fdsPaths := fds.Paths()
	paths := make([][]string, len(fdsPaths))
	for i := range fdsPaths {
		paths[i] = strings.Split(fdsPaths[i], "/")
	}
	dcs, err := op.kopia.RestoreMultipleItems(ctx, b.SnapshotID, paths)
	if err != nil {
		stats.readErr = errors.Wrap(err, "retrieving service data")
		return stats.readErr
	}
	stats.cs = dcs

	// restore those collections using graph
	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		stats.writeErr = errors.Wrap(err, "connecting to graph api")
		return stats.writeErr
	}

	if err := gc.RestoreMessages(ctx, dcs); err != nil {
		stats.writeErr = errors.Wrap(err, "restoring service data")
		return stats.writeErr
	}
	stats.gc = gc.Status()

	op.Status = Successful
	return nil
}

// writes the restoreOperation outcome to the modelStore.
func (op *RestoreOperation) persistResults(
	started time.Time,
	stats *restoreStats,
) {
	op.Status = Successful
	if stats.readErr != nil || stats.writeErr != nil {
		op.Status = Failed
	}
	op.Results.ReadErrors = stats.readErr
	op.Results.WriteErrors = stats.writeErr

	op.Results.ItemsRead = len(stats.cs) // TODO: file count, not collection count

	if stats.gc != nil {
		op.Results.ItemsWritten = stats.gc.ObjectCount
	}

	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	// TODO: persist operation to modelstore
}
