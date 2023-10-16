package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

func (ctrl *Controller) GetRestoreResource(
	ctx context.Context,
	service path.ServiceType,
	rc control.RestoreConfig,
	orig idname.Provider,
) (path.ServiceType, idname.Provider, error) {
	var (
		svc path.ServiceType
		pr  idname.Provider
		err error
	)

	switch service {
	case path.ExchangeService:
		svc, pr, err = exchange.GetRestoreResource(ctx, ctrl.AC, rc, ctrl.IDNameLookup, orig)
	case path.OneDriveService:
		svc, pr, err = onedrive.GetRestoreResource(ctx, ctrl.AC, rc, ctrl.IDNameLookup, orig)
	case path.SharePointService:
		svc, pr, err = sharepoint.GetRestoreResource(ctx, ctrl.AC, rc, ctrl.IDNameLookup, orig)
	case path.GroupsService:
		svc, pr, err = groups.GetRestoreResource(ctx, ctrl.AC, rc, ctrl.IDNameLookup, orig)
	default:
		err = clues.New("unknown service").With("service", service)
	}

	if err != nil {
		return path.UnknownService, nil, err
	}

	ctrl.IDNameLookup = idname.NewCache(map[string]string{pr.ID(): pr.Name()})

	return svc, pr, nil
}

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

	var (
		service = rcc.Selector.PathService()
		status  *support.ControllerOperationStatus
		deets   = &details.Builder{}
		err     error
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
			drive.NewUserDriveRestoreHandler(ctrl.AC),
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
	case path.GroupsService:
		status, err = groups.ConsumeRestoreCollections(
			ctx,
			rcc,
			ctrl.AC,
			ctrl.backupDriveIDNames,
			ctrl.backupSiteIDWebURL,
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
