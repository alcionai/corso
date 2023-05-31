package connector

import (
	"context"
	"strings"
	"sync"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/connector/discovery"
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
)

const (
	// copyBufferSize is used for chunked upload
	// Microsoft recommends 5-10MB buffers
	// https://docs.microsoft.com/en-us/graph/api/driveitem-createuploadsession?view=graph-rest-1.0#best-practices
	copyBufferSize = 5 * 1024 * 1024
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
) ([]data.BackupCollection, prefixmatcher.StringSetReader, error) {
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
		return nil, nil, clues.Stack(err).WithClues(ctx)
	}

	serviceEnabled, canMakeDeltaQueries, err := checkServiceEnabled(
		ctx,
		gc.Discovery.Users(),
		path.ServiceType(sels.Service),
		sels.DiscreteOwner)
	if err != nil {
		return nil, nil, err
	}

	if !serviceEnabled {
		return []data.BackupCollection{}, nil, nil
	}

	var (
		colls []data.BackupCollection
		ssmb  *prefixmatcher.StringSetMatcher
	)

	if !canMakeDeltaQueries {
		logger.Ctx(ctx).Info("delta requests not available")

		ctrlOpts.ToggleFeatures.DisableDelta = true
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		colls, ssmb, err = exchange.DataCollections(
			ctx,
			gc.Discovery,
			sels,
			gc.credentials.AzureTenantID,
			owner,
			metadata,
			gc.UpdateStatus,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, err
		}

	case selectors.ServiceOneDrive:
		colls, ssmb, err = onedrive.DataCollections(
			ctx,
			gc.Discovery,
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
		colls, ssmb, err = sharepoint.DataCollections(
			ctx,
			gc.Discovery,
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

	return colls, ssmb, nil
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

	info, err := gc.Discovery.Users().GetInfo(ctx, resourceOwner)
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

// ConsumeRestoreCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: gc.status is updated at the completion of operation
func (gc *GraphConnector) ConsumeRestoreCollections(
	ctx context.Context,
	backupVersion int,
	acct account.Account,
	sels selectors.Selector,
	dest control.RestoreDestination,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
) (*details.Details, error) {
	ctx, end := diagnostics.Span(ctx, "connector:restore")
	defer end()

	ctx = graph.BindRateLimiterConfig(ctx, graph.LimiterCfg{Service: sels.PathService()})

	var (
		status *support.ConnectorOperationStatus
		deets  = &details.Builder{}
	)

	creds, err := acct.M365Config()
	if err != nil {
		return nil, clues.Wrap(err, "malformed azure credentials")
	}

	// Buffer pool for uploads
	pool := sync.Pool{
		New: func() interface{} {
			b := make([]byte, copyBufferSize)
			return &b
		},
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		status, err = exchange.RestoreCollections(ctx,
			creds,
			gc.Discovery,
			gc.Service,
			dest,
			dcs,
			deets,
			errs)
	case selectors.ServiceOneDrive:
		status, err = onedrive.RestoreCollections(
			ctx,
			onedrive.NewRestoreHandler(gc.Discovery),
			backupVersion,
			dest,
			opts,
			dcs,
			deets,
			&pool,
			errs)
	case selectors.ServiceSharePoint:
		status, err = sharepoint.RestoreCollections(
			ctx,
			backupVersion,
			gc.Discovery,
			creds,
			dest,
			opts,
			dcs,
			deets,
			&pool,
			errs)
	default:
		err = clues.Wrap(clues.New(sels.Service.String()), "service not supported")
	}

	gc.incrementAwaitingMessages()
	gc.UpdateStatus(status)

	return deets.Details(), err
}
