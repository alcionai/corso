package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
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
) ([]data.Collection, error) {
	ctx, end := D.Span(ctx, "gc:dataCollections", D.Index("service", sels.Service.String()))
	defer end()

	err := verifyBackupInputs(sels, gc.GetUsers(), gc.GetSiteIDs())
	if err != nil {
		return nil, err
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		colls, err := exchange.DataCollections(
			ctx,
			sels,
			metadata,
			gc.GetUsers(),
			gc.credentials,
			// gc.Service,
			gc.UpdateStatus,
			ctrlOpts)
		if err != nil {
			return nil, err
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

		return colls, nil

	case selectors.ServiceOneDrive:
		return gc.OneDriveDataCollections(ctx, sels, ctrlOpts)

	case selectors.ServiceSharePoint:
		colls, err := sharepoint.DataCollections(
			ctx,
			sels,
			gc.GetSiteIDs(),
			gc.credentials.AzureTenantID,
			gc.Service,
			gc,
			ctrlOpts)
		if err != nil {
			return nil, err
		}

		for range colls {
			gc.incrementAwaitingMessages()
		}

		return colls, nil

	default:
		return nil, errors.Errorf("service %s not supported", sels.Service.String())
	}
}

func verifyBackupInputs(sels selectors.Selector, userPNs, siteIDs []string) error {
	var ids []string

	resourceOwners := sels.DiscreteResourceOwners()

	switch sels.Service {
	case selectors.ServiceExchange, selectors.ServiceOneDrive:
		ids = userPNs

	case selectors.ServiceSharePoint:
		ids = siteIDs
	}

	// verify resourceOwners
	normROs := map[string]struct{}{}

	for _, id := range ids {
		normROs[strings.ToLower(id)] = struct{}{}
	}

	for _, ro := range resourceOwners {
		if _, ok := normROs[strings.ToLower(ro)]; !ok {
			return fmt.Errorf("included resource owner %s not found within tenant", ro)
		}
	}

	return nil
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
) ([]data.Collection, error) {
	odb, err := selector.ToOneDriveBackup()
	if err != nil {
		return nil, errors.Wrap(err, "oneDriveDataCollection: parsing selector")
	}

	var (
		scopes      = odb.DiscreteScopes([]string{selector.DiscreteOwner})
		collections = []data.Collection{}
		errs        error
	)

	// for each scope that includes oneDrive items, get all
	for _, scope := range scopes {
		for _, user := range scope.Get(selectors.OneDriveUser) {
			logger.Ctx(ctx).With("user", user).Debug("Creating OneDrive collections")

			odcs, err := onedrive.NewCollections(
				gc.credentials.AzureTenantID,
				user,
				onedrive.OneDriveSource,
				odFolderMatcher{scope},
				gc.Service,
				gc.UpdateStatus,
				ctrlOpts,
			).Get(ctx)
			if err != nil {
				return nil, support.WrapAndAppend(user, err, errs)
			}

			collections = append(collections, odcs...)
		}
	}

	for range collections {
		gc.incrementAwaitingMessages()
	}

	return collections, errs
}
