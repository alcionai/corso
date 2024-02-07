package repository

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/model"
	"github.com/alcionai/canario/src/internal/operations"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/selectors"
	"github.com/alcionai/canario/src/pkg/store"
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
	handler, err := r.Provider.NewServiceHandler(sel.PathService())
	if err != nil {
		return operations.RestoreOperation{}, clues.Stack(err)
	}

	return operations.NewRestoreOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		handler,
		r.Account,
		model.StableID(backupID),
		sel,
		restoreCfg,
		r.Bus,
		count.New())
}
