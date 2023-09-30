package repository

import (
	"context"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type Restorer interface {
	NewRestore(
		ctx context.Context,
		backupID string,
		sel selectors.Selector,
		restoreCfg control.RestoreConfig,
	) (operations.RestoreOperation, error)
}

// NewRestore generates a restoreOperation runner.
func (r repository) NewRestore(
	ctx context.Context,
	backupID string,
	sel selectors.Selector,
	restoreCfg control.RestoreConfig,
) (operations.RestoreOperation, error) {
	return operations.NewRestoreOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		r.Provider,
		r.Account,
		model.StableID(backupID),
		sel,
		restoreCfg,
		r.Bus,
		count.New())
}
