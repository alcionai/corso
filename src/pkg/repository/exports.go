package repository

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
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
	handler, err := r.Provider.NewServiceHandler(r.Opts, sel.PathService())
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
