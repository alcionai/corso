package repository

import (
	"context"

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
	return operations.NewExportOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		r.Provider.NewServiceHandler(r.Opts, sel.PathService()),
		r.Account,
		model.StableID(backupID),
		sel,
		exportCfg,
		r.Bus)
}
