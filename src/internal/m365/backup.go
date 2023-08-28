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

	var (
		colls                []data.BackupCollection
		ssmb                 *prefixmatcher.StringSetMatcher
		canUsePreviousBackup bool
	)

	// All services except Exchange can make delta queries by default.
	// Exchange can only make delta queries if the mailbox is not over quota.
	canMakeDeltaQueries := true
	if service == path.ExchangeService {
		canMakeDeltaQueries, err = canExchangeMakeDeltaQueries(
			ctx,
			service,
			ctrl.AC.Users(),
			bpc.ProtectedResource.ID())
		if err != nil {
			return nil, nil, false, clues.Stack(err)
		}
	}

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

func (ctrl *Controller) IsServiceEnabled(
	ctx context.Context,
	service path.ServiceType,
	resourceOwner string,
) (bool, error) {
	switch service {
	case path.ExchangeService:
		return IsExchangeServiceEnabled(ctx, ctrl.AC.Users(), resourceOwner)
	case path.OneDriveService:
		return IsOneDriveServiceEnabled(ctx, ctrl.AC.Users(), resourceOwner)
	case path.SharePointService:
		return IsSharePointServiceEnabled(ctx, ctrl.AC.Users().Sites(), resourceOwner)
	case path.GroupsService:
		return true, nil
	}

	return false, clues.Wrap(clues.New(service.String()), "service not supported").WithClues(ctx)
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

func canExchangeMakeDeltaQueries(
	ctx context.Context,
	service path.ServiceType,
	gmi getMailboxer,
	resourceOwner string,
) (bool, error) {
	mi, err := GetMailboxInfo(ctx, gmi, resourceOwner)
	if err != nil {
		return false, clues.Stack(err)
	}

	return !mi.QuotaExceeded, nil
}
