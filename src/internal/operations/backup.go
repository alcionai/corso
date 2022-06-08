package operations

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/pkg/credentials"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation
	Version string

	creds credentials.M365

	Targets []string // something for targets/filter/source/app&users/etc
	Work    []string // something to reference the artifacts created, or at least their count

	// todo - graphConnector data streams
	// dataStreams  []*DataStream
}

// NewBackupOperation constructs and validates a backup operation.
func NewBackupOperation(
	ctx context.Context,
	opts OperationOpts,
	kw *kopia.KopiaWrapper,
	creds credentials.M365,
	targets []string,
) (BackupOperation, error) {
	bo := BackupOperation{
		operation: newOperation(opts, kw),
		Version:   "v0",
		creds:     creds,
		Targets:   targets,
		Work:      []string{},
	}
	if err := bo.validate(); err != nil {
		return BackupOperation{}, err
	}

	return bo, nil
}

func (bo BackupOperation) validate() error {
	if err := bo.creds.Validate(); err != nil {
		return errors.Wrap(err, "invalid credentials")
	}
	return bo.operation.validate()
}

// Run begins a synchronous backup operation.
func (bo BackupOperation) Run(ctx context.Context) error {
	_, err := connector.NewGraphConnector(bo.creds.TenantID, bo.creds.ClientID, bo.creds.ClientSecret)
	if err != nil {
		return errors.Wrap(err, "connecting to graph api")
	}

	// todo - use the graphConnector to create datastreams
	// dStreams, err := bo.gc.BackupOp(bo.Targets)

	// todo - send backup write request to BackupWriter
	// err = kopia.BackupWriter(ctx, bo.gc.TenantID, wg, prog, dStreams...)

	bo.Status = Successful
	return nil
}
