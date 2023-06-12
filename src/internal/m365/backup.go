package m365

import (
	"context"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/m365/discovery"
	"github.com/alcionai/corso/src/internal/m365/exchange"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	"github.com/alcionai/corso/src/internal/m365/sharepoint"
	"github.com/alcionai/corso/src/pkg/control"
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
func (gc *GraphConnector) ProduceBackupCollections(
	ctx context.Context,
	owner idname.Provider,
	sels selectors.Selector,
	metadata []data.RestoreCollection,
	lastBackupVersion int,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	ctx, end := diagnostics.Span(
		ctx,
		"gc:produceBackupCollections",
		diagnostics.Index("service", sels.Service.String()))
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: sels.PathService()})

	// Limit the max number of active requests to graph from this collection.
	ctrlOpts.Parallelism.ItemFetch = graph.Parallelism(sels.PathService()).
		ItemOverride(ctx, ctrlOpts.Parallelism.ItemFetch)

	err := verifyBackupInputs(sels, gc.IDNameLookup.IDs())
	if err != nil {
		return nil, nil, false, clues.Stack(err).WithClues(ctx)
	}

	serviceEnabled, canMakeDeltaQueries, err := checkServiceEnabled(
		ctx,
		gc.AC.Users(),
		path.ServiceType(sels.Service),
		sels.DiscreteOwner)
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

		ctrlOpts.ToggleFeatures.DisableDelta = true
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		colls, ssmb, canUsePreviousBackup, err = exchange.ProduceBackupCollections(
			ctx,
			gc.AC,
			sels,
			gc.credentials.AzureTenantID,
			owner,
			metadata,
			gc.UpdateStatus,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	case selectors.ServiceOneDrive:
		colls, ssmb, canUsePreviousBackup, err = onedrive.ProduceBackupCollections(
			ctx,
			gc.AC,
			sels,
			owner,
			metadata,
			lastBackupVersion,
			gc.credentials.AzureTenantID,
			gc.UpdateStatus,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	case selectors.ServiceSharePoint:
		colls, ssmb, canUsePreviousBackup, err = sharepoint.ProduceBackupCollections(
			ctx,
			gc.AC,
			sels,
			owner,
			metadata,
			gc.credentials,
			gc,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

	default:
		return nil, nil, false, clues.Wrap(clues.New(sels.Service.String()), "service not supported").WithClues(ctx)
	}

	for _, c := range colls {
		// kopia doesn't stream Items() from deleted collections,
		// and so they never end up calling the UpdateStatus closer.
		// This is a brittle workaround, since changes in consumer
		// behavior (such as calling Items()) could inadvertently
		// break the process state, putting us into deadlock or
		// panics.
		if c.State() != data.DeletedState {
			gc.incrementAwaitingMessages()
		}
	}

	return colls, ssmb, canUsePreviousBackup, nil
}

// IsBackupRunnable verifies that the users provided has the services enabled and
// data can be backed up. The canMakeDeltaQueries provides info if the mailbox is
// full and delta queries can be made on it.
func (gc *GraphConnector) IsBackupRunnable(
	ctx context.Context,
	service path.ServiceType,
	resourceOwner string,
) (bool, error) {
	if service == path.SharePointService {
		// No "enabled" check required for sharepoint
		return true, nil
	}

	info, err := gc.AC.Users().GetInfo(ctx, resourceOwner)
	if err != nil {
		return false, err
	}

	if !info.ServiceEnabled(service) {
		return false, clues.Wrap(graph.ErrServiceNotEnabled, "checking service access")
	}

	return true, nil
}

func verifyBackupInputs(sels selectors.Selector, siteIDs []string) error {
	var ids []string

	switch sels.Service {
	case selectors.ServiceExchange, selectors.ServiceOneDrive:
		// Exchange and OneDrive user existence now checked in checkServiceEnabled.
		return nil

	case selectors.ServiceSharePoint:
		ids = siteIDs
	}

	resourceOwner := strings.ToLower(sels.DiscreteOwner)

	if !filters.Equal(ids).Compare(resourceOwner) {
		return clues.Stack(graph.ErrResourceOwnerNotFound).With("missing_resource_owner", sels.DiscreteOwner)
	}

	return nil
}

func checkServiceEnabled(
	ctx context.Context,
	gi discovery.GetInfoer,
	service path.ServiceType,
	resource string,
) (bool, bool, error) {
	if service == path.SharePointService {
		// No "enabled" check required for sharepoint
		return true, true, nil
	}

	info, err := gi.GetInfo(ctx, resource)
	if err != nil {
		return false, false, err
	}

	if !info.ServiceEnabled(service) {
		return false, false, clues.Wrap(graph.ErrServiceNotEnabled, "checking service access")
	}

	canMakeDeltaQueries := true
	if service == path.ExchangeService {
		// we currently can only check quota exceeded for exchange
		canMakeDeltaQueries = info.CanMakeDeltaQueries()
	}

	return true, canMakeDeltaQueries, nil
}
