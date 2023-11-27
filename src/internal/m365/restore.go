package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ConsumeRestoreCollections restores data from the specified collections
// into M365 using the GraphAPI.
func (ctrl *Controller) ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, *data.CollectionStats, error) {
	ctx, end := diagnostics.Span(ctx, "m365:restore")
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: rcc.Selector.PathService()})
	ctx = clues.Add(ctx, "restore_config", rcc.RestoreConfig)

	if len(dcs) == 0 {
		return nil, nil, clues.New("no data collections to restore")
	}

	var (
		service = rcc.Selector.PathService()
		stats   *data.CollectionStats
		deets   *details.Details
		err     error
	)

	switch service {
	case path.ExchangeService:
		deets, stats, err = exchange.ConsumeRestoreCollections(
			ctx,
			ctrl.AC,
			rcc,
			dcs,
			errs,
			ctr)
	case path.OneDriveService:
		deets, stats, err = onedrive.ConsumeRestoreCollections(
			ctx,
			drive.NewUserDriveRestoreHandler(ctrl.AC),
			rcc,
			ctrl.backupDriveIDNames,
			dcs,
			errs,
			ctr)
	case path.SharePointService:
		deets, stats, err = sharepoint.ConsumeRestoreCollections(
			ctx,
			rcc,
			ctrl.AC,
			ctrl.backupDriveIDNames,
			dcs,
			errs,
			ctr)
	case path.GroupsService:
		deets, stats, err = groups.ConsumeRestoreCollections(
			ctx,
			rcc,
			ctrl.AC,
			ctrl.backupDriveIDNames,
			ctrl.backupSiteIDWebURL,
			dcs,
			errs,
			ctr)
	default:
		err = clues.Wrap(clues.New(service.String()), "service not supported")
	}

	return deets, stats, err
}
