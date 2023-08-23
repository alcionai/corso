package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ConsumeRestoreCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: status is updated at the completion of operation
func (ctrl *Controller) ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, error) {
	ctx, end := diagnostics.Span(ctx, "m365:restore")
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: rcc.Selector.PathService()})
	ctx = clues.Add(ctx, "restore_config", rcc.RestoreConfig)

	if len(dcs) == 0 {
		return nil, clues.New("no data collections to restore")
	}

	serviceEnabled, _, err := checkServiceEnabled(
		ctx,
		ctrl,
		rcc.Selector.PathService(),
		rcc.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	if !serviceEnabled {
		return nil, clues.Stack(graph.ErrServiceNotEnabled).WithClues(ctx)
	}

	var (
		service = rcc.Selector.PathService()
		status  *support.ControllerOperationStatus
		deets   = &details.Builder{}
	)

	switch service {
	case path.ExchangeService:
		status, err = exchange.ConsumeRestoreCollections(
			ctx,
			ctrl.AC,
			rcc,
			dcs,
			deets,
			errs,
			ctr)
	case path.OneDriveService:
		status, err = onedrive.ConsumeRestoreCollections(
			ctx,
			drive.NewRestoreHandler(ctrl.AC),
			rcc,
			ctrl.backupDriveIDNames,
			dcs,
			deets,
			errs,
			ctr)
	case path.SharePointService:
		status, err = sharepoint.ConsumeRestoreCollections(
			ctx,
			rcc,
			ctrl.AC,
			ctrl.backupDriveIDNames,
			dcs,
			deets,
			errs,
			ctr)
	default:
		err = clues.Wrap(clues.New(service.String()), "service not supported")
	}

	ctrl.incrementAwaitingMessages()
	ctrl.UpdateStatus(status)

	return deets.Details(), err
}
