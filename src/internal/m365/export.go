package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ExportRestoreCollections exports data from the specified collections
func (ctrl *Controller) ExportRestoreCollections(
	ctx context.Context,
	backupVersion int,
	sels selectors.Selector,
	exportCfg control.ExportConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
) ([]export.Collection, error) {
	ctx, end := diagnostics.Span(ctx, "m365:export")
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: sels.PathService()})
	ctx = clues.Add(ctx, "export_config", exportCfg) // TODO(meain): needs PII control

	var (
		expCollections []export.Collection
		status         *support.ControllerOperationStatus
		deets          = &details.Builder{}
		err            error
	)

	switch sels.Service {
	case selectors.ServiceOneDrive:
		expCollections, err = onedrive.ExportRestoreCollections(
			ctx,
			backupVersion,
			exportCfg,
			opts,
			dcs,
			deets,
			errs)
	default:
		err = clues.Wrap(clues.New(sels.Service.String()), "service not supported")
	}

	ctrl.incrementAwaitingMessages()
	ctrl.UpdateStatus(status)

	return expCollections, err
}
