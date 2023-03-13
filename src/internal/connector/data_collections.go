package connector

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/discovery"
	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Data Collections
// ---------------------------------------------------------------------------

// DataCollections utility function to launch backup operations for exchange and
// onedrive. metadataCols contains any collections with metadata files that may
// be useful for the current backup. Metadata can include things like delta
// tokens or the previous backup's folder hierarchy. The absence of metadataCols
// results in all data being pulled.
func (gc *GraphConnector) DataCollections(
	ctx context.Context,
	sels selectors.Selector,
	metadata []data.RestoreCollection,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	ctx, end := D.Span(ctx, "gc:dataCollections", D.Index("service", sels.Service.String()))
	defer end()

	err := verifyBackupInputs(sels, gc.GetSiteIDs())
	if err != nil {
		return nil, nil, clues.Stack(err).WithClues(ctx)
	}

	serviceEnabled, err := checkServiceEnabled(
		ctx,
		gc.Owners.Users(),
		path.ServiceType(sels.Service),
		sels.DiscreteOwner)
	if err != nil {
		return nil, nil, err
	}

	if !serviceEnabled {
		return []data.BackupCollection{}, nil, nil
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		colls, excludes, err := exchange.DataCollections(
			ctx,
			sels,
			metadata,
			gc.credentials,
			gc.UpdateStatus,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, err
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

	case selectors.ServiceOneDrive:
		colls, excludes, err := onedrive.DataCollections(
			ctx,
			sels, metadata,
			gc.credentials.AzureTenantID,
			gc.itemClient,
			gc.Service,
			gc.UpdateStatus,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, err
		}

		for _, c := range colls {
			// kopia doesn't stream Items() from deleted collections.
			if c.State() != data.DeletedState {
				gc.incrementAwaitingMessages()
			}
		}

		return colls, excludes, nil

	case selectors.ServiceSharePoint:
		colls, excludes, err := sharepoint.DataCollections(
			ctx,
			gc.itemClient,
			sels,
			gc.credentials,
			gc.Service,
			gc,
			ctrlOpts,
			errs)
		if err != nil {
			return nil, nil, err
		}

		gc.incrementMessagesBy(len(colls))

		return colls, excludes, nil

	default:
		return nil, nil, clues.Wrap(clues.New(sels.Service.String()), "service not supported").WithClues(ctx)
	}
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

	var found bool

	for _, id := range ids {
		if strings.ToLower(id) == resourceOwner {
			found = true
			break
		}
	}

	if !found {
		return clues.New("resource owner not found within tenant").With("missing_resource_owner", sels.DiscreteOwner)
	}

	return nil
}

func checkServiceEnabled(
	ctx context.Context,
	au api.Users,
	service path.ServiceType,
	resource string,
) (bool, error) {
	if service == path.SharePointService {
		// No "enabled" check required for sharepoint
		return true, nil
	}

	_, info, err := discovery.User(ctx, au, resource)
	if err != nil {
		return false, err
	}

	if _, ok := info.DiscoveredServices[service]; !ok {
		return false, nil
	}

	return true, nil
}

// RestoreDataCollections restores data from the specified collections
// into M365 using the GraphAPI.
// SideEffect: gc.status is updated at the completion of operation
func (gc *GraphConnector) RestoreDataCollections(
	ctx context.Context,
	backupVersion int,
	acct account.Account,
	selector selectors.Selector,
	dest control.RestoreDestination,
	opts control.Options,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
) (*details.Details, error) {
	ctx, end := D.Span(ctx, "connector:restore")
	defer end()

	var (
		status *support.ConnectorOperationStatus
		deets  = &details.Builder{}
	)

	creds, err := acct.M365Config()
	if err != nil {
		return nil, errors.Wrap(err, "malformed azure credentials")
	}

	switch selector.Service {
	case selectors.ServiceExchange:
		status, err = exchange.RestoreExchangeDataCollections(ctx, creds, gc.Service, dest, dcs, deets, errs)
	case selectors.ServiceOneDrive:
		status, err = onedrive.RestoreCollections(ctx, backupVersion, gc.Service, dest, opts, dcs, deets, errs)
	case selectors.ServiceSharePoint:
		status, err = sharepoint.RestoreCollections(ctx, backupVersion, creds, gc.Service, dest, dcs, deets, errs)
	default:
		err = clues.Wrap(clues.New(selector.Service.String()), "service not supported")
	}

	gc.incrementAwaitingMessages()
	gc.UpdateStatus(status)

	return deets.Details(), err
}
