package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// NewServiceHandler returns an instance of a struct capable of running various
// operations for a given service.
func (ctrl *Controller) NewServiceHandler(
	opts control.Options,
	service path.ServiceType,
) inject.ServiceHandler {
	switch service {
	case path.OneDriveService:
		return onedrive.NewOneDriveHandler(opts)

	case path.SharePointService:
		return sharepoint.NewSharePointHandler(opts)

	case path.GroupsService:
		return groups.NewGroupsHandler(opts)
	}

	return &inject.BaseServiceHandler{}
}

// ProduceExportCollections exports data from the specified collections
func (ctrl *Controller) ProduceExportCollections(
	ctx context.Context,
	backupVersion int,
	sels selectors.Selector,
	exportCfg control.ExportConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	stats *data.ExportStats,
	errs *fault.Bus,
) ([]export.Collectioner, error) {
	ctx, end := diagnostics.Span(ctx, "m365:export")
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: sels.PathService()})
	ctx = clues.Add(ctx, "export_config", exportCfg)

	var (
		expCollections []export.Collectioner
		status         *support.ControllerOperationStatus
		deets          = &details.Builder{}
		err            error
	)

	switch sels.Service {
	case selectors.ServiceOneDrive:
		expCollections, err = onedrive.ProduceExportCollections(
			ctx,
			backupVersion,
			exportCfg,
			opts,
			dcs,
			deets,
			stats,
			errs)
	case selectors.ServiceSharePoint:
		expCollections, err = sharepoint.ProduceExportCollections(
			ctx,
			backupVersion,
			exportCfg,
			opts,
			dcs,
			ctrl.backupDriveIDNames,
			deets,
			stats,
			errs)
	case selectors.ServiceGroups:
		expCollections, err = groups.ProduceExportCollections(
			ctx,
			backupVersion,
			exportCfg,
			opts,
			dcs,
			ctrl.backupDriveIDNames,
			ctrl.backupSiteIDWebURL,
			deets,
			stats,
			errs)

	default:
		err = clues.Wrap(clues.New(sels.Service.String()), "service not supported")
	}

	ctrl.incrementAwaitingMessages()
	ctrl.UpdateStatus(status)

	return expCollections, err
}
