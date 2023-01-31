package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/discovery"
	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
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
	metadata []data.Collection,
	ctrlOpts control.Options,
) ([]data.Collection, map[string]struct{}, error) {
	ctx, end := D.Span(ctx, "gc:dataCollections", D.Index("service", sels.Service.String()))
	defer end()

	err := verifyBackupInputs(sels, gc.GetUsers(), gc.GetSiteIDs())
	if err != nil {
		return nil, nil, err
	}

	serviceEnabled, err := checkServiceEnabled(ctx, gc.Owners.Users(), path.ServiceType(sels.Service), sels.DiscreteOwner)
	if err != nil {
		return nil, nil, err
	}

	if !serviceEnabled {
		return []data.Collection{}, nil, nil
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		colls, excludes, err := exchange.DataCollections(
			ctx,
			sels,
			metadata,
			gc.credentials,
			// gc.Service,
			gc.UpdateStatus,
			ctrlOpts)
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
		return gc.OneDriveDataCollections(ctx, sels, ctrlOpts)

	case selectors.ServiceSharePoint:
		colls, excludes, err := sharepoint.DataCollections(
			ctx,
			gc.itemClient,
			sels,
			gc.credentials,
			gc.Service,
			gc,
			ctrlOpts)
		if err != nil {
			return nil, nil, err
		}

		for range colls {
			gc.incrementAwaitingMessages()
		}

		return colls, excludes, nil

	default:
		return nil, nil, errors.Errorf("service %s not supported", sels.Service.String())
	}
}

func verifyBackupInputs(sels selectors.Selector, userPNs, siteIDs []string) error {
	var ids []string

	switch sels.Service {
	case selectors.ServiceExchange, selectors.ServiceOneDrive:
		ids = userPNs

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
		return fmt.Errorf("resource owner [%s] not found within tenant", sels.DiscreteOwner)
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

// ---------------------------------------------------------------------------
// OneDrive
// ---------------------------------------------------------------------------

type odFolderMatcher struct {
	scope selectors.OneDriveScope
}

func (fm odFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.OneDriveFolder)
}

func (fm odFolderMatcher) Matches(dir string) bool {
	return fm.scope.Matches(selectors.OneDriveFolder, dir)
}

// OneDriveDataCollections returns a set of DataCollection which represents the OneDrive data
// for the specified user
func (gc *GraphConnector) OneDriveDataCollections(
	ctx context.Context,
	selector selectors.Selector,
	ctrlOpts control.Options,
) ([]data.Collection, map[string]struct{}, error) {
	odb, err := selector.ToOneDriveBackup()
	if err != nil {
		return nil, nil, errors.Wrap(err, "oneDriveDataCollection: parsing selector")
	}

	var (
		user        = selector.DiscreteOwner
		collections = []data.Collection{}
		allExcludes = map[string]struct{}{}
		errs        error
	)

	// for each scope that includes oneDrive items, get all
	for _, scope := range odb.Scopes() {
		logger.Ctx(ctx).With("user", user).Debug("Creating OneDrive collections")

		odcs, excludes, err := onedrive.NewCollections(
			gc.itemClient,
			gc.credentials.AzureTenantID,
			user,
			onedrive.OneDriveSource,
			odFolderMatcher{scope},
			gc.Service,
			gc.UpdateStatus,
			ctrlOpts,
		).Get(ctx)
		if err != nil {
			return nil, nil, support.WrapAndAppend(user, err, errs)
		}

		collections = append(collections, odcs...)

		maps.Copy(allExcludes, excludes)
	}

	for range collections {
		gc.incrementAwaitingMessages()
	}

	return collections, allExcludes, errs
}
