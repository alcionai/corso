package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Data Collections
// ---------------------------------------------------------------------------

// ProduceBackupCollections generates a slice of data.BackupCollections for the service
// specified in the selectors.
// The metadata field can include things like delta tokens or the previous backup's
// folder hierarchy. The absence of metadata causes the collection creation to ignore
// prior history (ie, incrementals) and run a full backup.
func (ctrl *Controller) ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	errs *fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	service := bpc.Selector.PathService()

	ctx, end := diagnostics.Span(
		ctx,
		"m365:produceBackupCollections",
		diagnostics.Index("service", bpc.Selector.PathService().String()))
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: service})

	// Limit the max number of active requests to graph from this collection.
	bpc.Options.Parallelism.ItemFetch = graph.Parallelism(service).
		ItemOverride(ctx, bpc.Options.Parallelism.ItemFetch)

	err := verifyBackupInputs(bpc.Selector, ctrl.IDNameLookup.IDs())
	if err != nil {
		return nil, nil, false, clues.Stack(err).WithClues(ctx)
	}

	serviceEnabled, canMakeDeltaQueries, err := checkServiceEnabled(
		ctx,
		ctrl,
		service,
		bpc.ProtectedResource.ID())
	if err != nil {
		return nil, nil, false, err
	}

	if !serviceEnabled {
		return []data.BackupCollection{}, nil, false, nil
	}

	var (
		colls                []data.BackupCollection
		ssmb                 *prefixmatcher.StringSetMatcher
		canUsePreviousBackup bool
	)

	if !canMakeDeltaQueries {
		logger.Ctx(ctx).Info("delta requests not available")

		bpc.Options.ToggleFeatures.DisableDelta = true
	}

	switch service {
	case path.ExchangeService:
		colls, ssmb, canUsePreviousBackup, err = exchange.ProduceBackupCollections(
			ctx,
			bpc,
			ctrl.AC,
			ctrl.credentials.AzureTenantID,
			ctrl.UpdateStatus,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	case path.OneDriveService:
		colls, ssmb, canUsePreviousBackup, err = onedrive.ProduceBackupCollections(
			ctx,
			bpc,
			ctrl.AC,
			ctrl.credentials.AzureTenantID,
			ctrl.UpdateStatus,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	case path.SharePointService:
		colls, ssmb, canUsePreviousBackup, err = sharepoint.ProduceBackupCollections(
			ctx,
			bpc,
			ctrl.AC,
			ctrl.credentials,
			ctrl.UpdateStatus,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	case path.GroupsService:
		colls, ssmb, canUsePreviousBackup, err = groups.ProduceBackupCollections(
			ctx,
			bpc,
			ctrl.AC,
			ctrl.credentials,
			ctrl.UpdateStatus,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	default:
		return nil, nil, false, clues.Wrap(clues.New(service.String()), "service not supported").WithClues(ctx)
	}

	for _, c := range colls {
		// kopia doesn't stream Items() from deleted collections,
		// and so they never end up calling the UpdateStatus closer.
		// This is a brittle workaround, since changes in consumer
		// behavior (such as calling Items()) could inadvertently
		// break the process state, putting us into deadlock or
		// panics.
		if c.State() != data.DeletedState {
			ctrl.incrementAwaitingMessages()
		}
	}

	return colls, ssmb, canUsePreviousBackup, nil
}

// IsBackupRunnable verifies that the users provided has the services enabled and
// data can be backed up. The canMakeDeltaQueries provides info if the mailbox is
// full and delta queries can be made on it.
func (ctrl *Controller) IsBackupRunnable(
	ctx context.Context,
	service path.ServiceType,
	resourceOwner string,
) (bool, error) {
	if service == path.GroupsService {
		_, err := ctrl.AC.Groups().GetByID(ctx, resourceOwner)
		if err != nil {
			// TODO(meain): check for error message in case groups are
			// not enabled at all similar to sharepoint
			return false, err
		}

		return true, nil
	}

	if service == path.SharePointService {
		_, err := ctrl.AC.Sites().GetRoot(ctx)
		if err != nil {
			if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
				return false, clues.Stack(graph.ErrServiceNotEnabled, err)
			}

			return false, err
		}

		return true, nil
	}

	info, err := ctrl.AC.Users().GetInfo(ctx, resourceOwner)
	if err != nil {
		return false, clues.Stack(err)
	}

	if !info.ServiceEnabled(service) {
		return false, clues.Wrap(graph.ErrServiceNotEnabled, "checking service access")
	}

	return true, nil
}

func verifyBackupInputs(sels selectors.Selector, cachedIDs []string) error {
	var ids []string

	switch sels.Service {
	case selectors.ServiceExchange, selectors.ServiceOneDrive:
		// Exchange and OneDrive user existence now checked in checkServiceEnabled.
		return nil

	case selectors.ServiceSharePoint, selectors.ServiceGroups:
		ids = cachedIDs
	}

	if !filters.Contains(ids).Compare(sels.ID()) {
		return clues.Stack(graph.ErrResourceOwnerNotFound).
			With("selector_protected_resource", sels.DiscreteOwner)
	}

	return nil
}

func checkServiceEnabled(
	ctx context.Context,
	ctrl *Controller,
	service path.ServiceType,
	resource string,
) (bool, bool, error) {
	switch service {
	case path.ExchangeService:
		return checkExchangeServiceEnabled(ctx, ctrl, resource)
	case path.OneDriveService:
		return checkOneDriveServiceEnabled(ctx, ctrl, resource)
	case path.SharePointService:
		return checkSharepointServiceEnabled(ctx, ctrl, resource)
	case path.GroupsService:
		return true, true, nil
	}

	return false, false, clues.Wrap(clues.New(service.String()), "service not supported").WithClues(ctx)
}

func checkOneDriveServiceEnabled(
	ctx context.Context,
	ctrl *Controller,
	resource string,
) (bool, bool, error) {
	if _, err := ctrl.AC.Users().GetDefaultDrive(ctx, resource); err != nil {
		if !clues.HasLabel(err, graph.LabelsMysiteNotFound) && !clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			logger.CtxErr(ctx, err).Error("getting user's default drive")

			return false, false, graph.Wrap(ctx, err, "getting user's default drive info")
		}

		logger.Ctx(ctx).Info("resource owner does not have a drive")

		return false, false, nil
	}

	return true, true, nil
}

func checkExchangeServiceEnabled(
	ctx context.Context,
	ctrl *Controller,
	resource string,
) (bool, bool, error) {
	enabled, err := ctrl.AC.Users().IsExchangeServiceEnabled(ctx, resource)
	if err != nil {
		return false, false, clues.Stack(err)
	}

	if !enabled {
		return false, false, nil
	}

	mi, err := ctrl.AC.Users().GetMailboxInfo(ctx, resource)
	if err != nil {
		return false, false, clues.Stack(err)
	}

	return true, !mi.QuotaExceeded, nil
}

func checkSharepointServiceEnabled(
	ctx context.Context,
	ctrl *Controller,
	resource string,
) (bool, bool, error) {
	_, err := ctrl.AC.Sites().GetRoot(ctx)
	if err != nil {
		if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, false, nil
		}

		return false, false, err
	}

	return true, true, nil
}
