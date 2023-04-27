package connector

import (
	"context"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	ctx, end := diagnostics.Span(
		ctx,
		"gc:produceBackupCollections",
		diagnostics.Index("service", sels.Service.String()))
	defer end()

	// Limit the max number of active requests to graph from this collection.
	ctrlOpts.Parallelism.ItemFetch = graph.Parallelism(sels.PathService()).
		ItemOverride(ctx, ctrlOpts.Parallelism.ItemFetch)

	err := verifyBackupInputs(sels, gc.IDNameLookup.IDs())
	if err != nil {
		return nil, nil, clues.Stack(err).WithClues(ctx)
	}

	info, err := gc.Discovery.Users().GetInfo(ctx, sels.DiscreteOwner)
	if err != nil {
		return nil, nil, err
	}

	serviceEnabled, err := checkServiceEnabled(info, path.ServiceType(sels.Service))
	if err != nil {
		return nil, nil, err
	}

	if !serviceEnabled {
		return []data.BackupCollection{}, nil, nil
	}

	var (
		colls    []data.BackupCollection
		excludes map[string]map[string]struct{}
	)

	deltaAvailable := info.CanMakeDeltaQueries()
	if !deltaAvailable {
		logger.Ctx(ctx).Info("delta requests not available")
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		colls, excludes, err = exchange.DataCollections(
			ctx,
			sels,
			owner,
			metadata,
			gc.credentials,
			gc.UpdateStatus,
			ctrlOpts,
			deltaAvailable,
			errs)
		if err != nil {
			return nil, nil, err
		}

	case selectors.ServiceOneDrive:
		colls, excludes, err = onedrive.DataCollections(
			ctx,
			sels,
			owner,
			metadata,
			lastBackupVersion,
			gc.credentials.AzureTenantID,
			gc.itemClient,
			gc.Service,
			gc.UpdateStatus,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, err
		}

	case selectors.ServiceSharePoint:
		colls, excludes, err = sharepoint.DataCollections(
			ctx,
			gc.itemClient,
			sels,
			owner,
			metadata,
			gc.credentials,
			gc.Service,
			gc,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, err
		}

	default:
		return nil, nil, clues.Wrap(clues.New(sels.Service.String()), "service not supported").WithClues(ctx)
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

	return colls, excludes, nil
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
	info *api.UserInfo,
	service path.ServiceType,
) (bool, error) {
	if service == path.SharePointService {
		// No "enabled" check required for sharepoint
		return true, nil
	}

	if !info.ServiceEnabled(service) {
		return false, clues.Wrap(graph.ErrServiceNotEnabled, "checking service access")
	}

	return true, nil
}

// ConsumeRestoreCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: gc.status is updated at the completion of operation
func (gc *GraphConnector) ConsumeRestoreCollections(
	ctx context.Context,
	backupVersion int,
	acct account.Account,
	selector selectors.Selector,
	dest control.RestoreDestination,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
) (*details.Details, error) {
	ctx, end := diagnostics.Span(ctx, "connector:restore")
	defer end()

	var (
		status *support.ConnectorOperationStatus
		deets  = &details.Builder{}
	)

	creds, err := acct.M365Config()
	if err != nil {
		return nil, clues.Wrap(err, "malformed azure credentials")
	}

	switch selector.Service {
	case selectors.ServiceExchange:
		status, err = exchange.RestoreExchangeDataCollections(ctx, creds, gc.Service, dest, dcs, deets, errs)
	case selectors.ServiceOneDrive:
		status, err = onedrive.RestoreCollections(ctx, creds, backupVersion, gc.Service, dest, opts, dcs, deets, errs)
	case selectors.ServiceSharePoint:
		status, err = sharepoint.RestoreCollections(ctx, backupVersion, creds, gc.Service, dest, dcs, deets, errs)
	default:
		err = clues.Wrap(clues.New(selector.Service.String()), "service not supported")
	}

	gc.incrementAwaitingMessages()
	gc.UpdateStatus(status)

	return deets.Details(), err
}
