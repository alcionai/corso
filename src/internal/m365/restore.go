package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/exchange"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	"github.com/alcionai/corso/src/internal/m365/sharepoint"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ConsumeRestoreCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: status is updated at the completion of operation
func (ctrl *Controller) ConsumeRestoreCollections(
	ctx context.Context,
	backupVersion int,
	sels selectors.Selector,
	restoreCfg control.RestoreConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, error) {
	ctx, end := diagnostics.Span(ctx, "m365:restore")
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: sels.PathService()})
	ctx = clues.Add(ctx, "restore_config", restoreCfg) // TODO(rkeepers): needs PII control

	var (
		status *support.ControllerOperationStatus
		deets  = &details.Builder{}
		err    error
	)

	switch sels.Service {
	case selectors.ServiceExchange:
		status, err = exchange.ConsumeRestoreCollections(ctx, ctrl.AC, restoreCfg, dcs, deets, errs, ctr)
	case selectors.ServiceOneDrive:
		status, err = onedrive.ConsumeRestoreCollections(
			ctx,
			onedrive.NewRestoreHandler(ctrl.AC),
			backupVersion,
			restoreCfg,
			opts,
			ctrl.backupDriveIDNames,
			dcs,
			deets,
			errs,
			ctr)
	case selectors.ServiceSharePoint:
		status, err = sharepoint.ConsumeRestoreCollections(
			ctx,
			backupVersion,
			ctrl.AC,
			restoreCfg,
			opts,
			ctrl.backupDriveIDNames,
			dcs,
			deets,
			errs,
			ctr)
	default:
		err = clues.Wrap(clues.New(sels.Service.String()), "service not supported")
	}

	ctrl.incrementAwaitingMessages()
	ctrl.UpdateStatus(status)

	return deets.Details(), err
}
