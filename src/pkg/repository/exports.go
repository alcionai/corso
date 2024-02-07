package repository

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/model"
	"github.com/alcionai/canario/src/internal/operations"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/selectors"
	"github.com/alcionai/canario/src/pkg/store"
)

type Exporter interface {
	NewExport(
		ctx context.Context,
		backupID string,
		sel selectors.Selector,
		exportCfg control.ExportConfig,
	) (operations.ExportOperation, error)
}

// NewExport generates a exportOperation runner.
func (r repository) NewExport(
	ctx context.Context,
	backupID string,
	sel selectors.Selector,
	exportCfg control.ExportConfig,
) (operations.ExportOperation, error) {
	handler, err := r.Provider.NewServiceHandler(sel.PathService())
	if err != nil {
		return operations.ExportOperation{}, clues.Stack(err)
	}

	return operations.NewExportOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		handler,
		r.Account,
		model.StableID(backupID),
		sel,
		exportCfg,
		r.Bus)
}
