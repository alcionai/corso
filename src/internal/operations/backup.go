package operations

import (
	"context"
	"sync"

	"github.com/alcionai/corso/internal/kopia"
)

// BackupOperation wraps an operation with backup-specific props.
type BackupOperation struct {
	operation
	Version string
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
	targets []string,
) (BackupOperation, error) {
	// todo - initialize a graphConnector
	// gc, err := graphConnector.Connect(bo.account)

	bo := BackupOperation{
		operation: newOperation(opts, kw),
		Version:   "v0",
		Targets:   targets,
		Work:      []string{},
	}
	if err := bo.validate(); err != nil {
		return BackupOperation{}, err
	}

	return bo, nil
}

func (bo BackupOperation) validate() error {
	return bo.operation.validate()
}

// Run begins a synchronous backup operation.
func (bo BackupOperation) Run(ctx context.Context) error {
	// todo - use the graphConnector to create datastreams
	// dStreams, err := bo.gc.BackupOp(bo.Targets)

	prog := newOpProgress()
	go recordProgress(ctx, bo, prog)

	var wg sync.WaitGroup
	// todo - send backup write request to BackupWriter
	// wg.Add(1)
	// err = kopia.BackupWriter(ctx, bo.gc.TenantID, wg, prog, dStreams...)
	wg.Wait()

	bo.Status = Successful
	return nil
}

// updates the BackupOperation.Work and BackupOperation.Errors with the
// stream of ongoing progress from the backupWriter
func recordProgress(ctx context.Context, bo BackupOperation, op *opProgress) {
	errs := op.errorChan
	prog := op.progressChan
	for {
		select {

		case err, ok := <-errs:
			if !ok {
				errs = nil
				break
			}
			bo.Errors = append(bo.Errors, err)

		case work, ok := <-prog:
			if !ok {
				prog = nil
				break
			}
			bo.Work = append(bo.Work, work)

		// exit if the context is canceled or terminated.
		case <-ctx.Done():
			return
		}

		// exit when both channels are closed
		if errs == nil && prog == nil {
			return
		}
	}
}
