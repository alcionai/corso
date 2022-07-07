package operations

import (
	"context"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/account"
)

// RestoreOperation wraps an operation with restore-specific props.
type RestoreOperation struct {
	operation

	RestorePointID string         `json:"restorePointID"`
	Results        RestoreResults `json:"results"`
	Targets        []string       `json:"selectors"` // todo: replace with Selectors
	Version        string         `json:"bersion"`

	account account.Account
}

// RestoreResults aggregate the details of the results of the operation.
type RestoreResults struct {
	summary
	metrics
}

// NewRestoreOperation constructs and validates a restore operation.
func NewRestoreOperation(
	ctx context.Context,
	opts Options,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	acct account.Account,
	restorePointID string,
	targets []string,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation:      newOperation(opts, kw, ms),
		RestorePointID: restorePointID,
		Targets:        targets,
		Version:        "v0",
		account:        acct,
	}
	if err := op.validate(); err != nil {
		return RestoreOperation{}, err
	}

	return op, nil
}

func (op RestoreOperation) validate() error {
	return op.operation.validate()
}

// Run begins a synchronous restore operation.
// todo (keepers): return stats block in first param.
func (op *RestoreOperation) Run(ctx context.Context) error {
	// TODO: persist initial state of restoreOperation in modelstore

	var (
		cs                []connector.DataCollection
		readErr, writeErr error
	)

	// persist operation results to the model store on exit
	defer op.persistResults(time.Now(), cs, readErr, writeErr)

	dc, readErr := op.kopia.RestoreSingleItem(ctx, op.RestorePointID, op.Targets)
	if readErr != nil {
		return errors.Wrap(readErr, "retrieving service data")
	}

	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		writeErr = multierror.Append(writeErr, err)
		return errors.Wrap(err, "connecting to graph api")
	}

	if err := gc.RestoreMessages(ctx, dc); err != nil {
		writeErr = multierror.Append(writeErr, err)
		return errors.Wrap(err, "restoring service data")
	}

	op.Status = Successful
	return nil
}

// writes the restoreOperation outcome to the modelStore.
func (op *RestoreOperation) persistResults(
	started time.Time,
	cs []connector.DataCollection,
	readErr, writeErr error,
) {
	op.Status = Successful
	if readErr != nil || writeErr != nil {
		op.Status = Failed
	}

	op.Results.ItemsRead = len(cs) // TODO: file count, not collection count
	op.Results.ReadErrors = readErr
	op.Results.ItemsWritten = -1 // TODO: get write count from GC
	op.Results.WriteErrors = writeErr

	op.Results.StartedAt = started
	op.Results.CompletedAt = time.Now()

	// TODO: persist operation to modelstore
}
