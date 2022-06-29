package operations

import (
	"context"

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
	kw *kopia.KopiaWrapper,
	acct account.Account,
	restorePointID string,
	targets []string,
) (RestoreOperation, error) {
	op := RestoreOperation{
		operation:      newOperation(opts, kw),
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
	dc, err := op.kopia.RestoreSingleItem(ctx, op.RestorePointID, op.Targets)
	if err != nil {
		return errors.Wrap(err, "retrieving service data")
	}

	gc, err := connector.NewGraphConnector(op.account)
	if err != nil {
		return errors.Wrap(err, "connecting to graph api")
	}

	if err := gc.RestoreMessages(ctx, dc); err != nil {
		return errors.Wrap(err, "restoring service data")
	}

	op.Status = Successful
	return nil
}
