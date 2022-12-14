package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
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
		return gc.ExchangeDataCollection(ctx, sels, metadata, ctrlOpts)
	case selectors.ServiceOneDrive:
		return gc.OneDriveDataCollections(ctx, sels, ctrlOpts)
	case selectors.ServiceSharePoint:
		colls, err := sharepoint.DataCollections(
			ctx,
			sels,
			gc.GetSiteIDs(),
			gc.credentials.AzureTenantID,
			gc,
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

	resourceOwners, err := sels.ResourceOwners()
	if err != nil {
		return errors.Wrap(err, "invalid backup inputs")
	}

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

	for _, ro := range resourceOwners.Includes {
		if _, ok := normROs[strings.ToLower(ro)]; !ok {
			return fmt.Errorf("included resource owner %s not found within tenant", ro)
		}
	}

	for _, ro := range resourceOwners.Excludes {
		if _, ok := normROs[strings.ToLower(ro)]; !ok {
			return fmt.Errorf("excluded resource owner %s not found within tenant", ro)
		}
	}

	for _, ro := range resourceOwners.Filters {
		if _, ok := normROs[strings.ToLower(ro)]; !ok {
			return fmt.Errorf("filtered resource owner %s not found within tenant", ro)
		}
	}

	return nil
}

// ---------------------------------------------------------------------------
// Exchange
// ---------------------------------------------------------------------------

// createExchangeCollections - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.ExchangeScope
// determines the type of collections that are retrieved.
func (gc *GraphConnector) createExchangeCollections(
	ctx context.Context,
	scope selectors.ExchangeScope,
	deltas map[string]string,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	var (
		errs           *multierror.Error
		users          = scope.Get(selectors.ExchangeUser)
		allCollections = make([]data.Collection, 0)
	)

	// Create collection of ExchangeDataCollection
	for _, user := range users {
		collections := make(map[string]data.Collection)

		qp := graph.QueryParams{
			Category:      scope.Category().PathType(),
			ResourceOwner: user,
			Credentials:   gc.credentials,
		}

		foldersComplete, closer := observe.MessageWithCompletion(fmt.Sprintf("∙ %s - %s:", qp.Category, user))
		defer closer()
		defer close(foldersComplete)

		resolver, err := exchange.PopulateExchangeContainerResolver(ctx, qp)
		if err != nil {
			return nil, errors.Wrap(err, "getting folder cache")
		}

		err = exchange.FilterContainersAndFillCollections(
			ctx,
			qp,
			collections,
			gc.UpdateStatus,
			resolver,
			scope,
			deltas,
			ctrlOpts)

		if err != nil {
			return nil, errors.Wrap(err, "filling collections")
		}

		foldersComplete <- struct{}{}

		for _, collection := range collections {
			gc.incrementAwaitingMessages()

			allCollections = append(allCollections, collection)
		}
	}

	return allCollections, errs.ErrorOrNil()
}

// ExchangeDataCollections returns a DataCollection which the caller can
// use to read mailbox data out for the specified user
// Assumption: User exists
//
//	Add iota to this call -> mail, contacts, calendar,  etc.
func (gc *GraphConnector) ExchangeDataCollection(
	ctx context.Context,
	selector selectors.Selector,
	metadata []data.Collection,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, errors.Wrap(err, "exchangeDataCollection: parsing selector")
	}

	var (
		scopes      = eb.DiscreteScopes(gc.GetUsers())
		collections = []data.Collection{}
		errs        error
	)

	_, deltas, err := exchange.ParseMetadataCollections(ctx, metadata)
	if err != nil {
		return nil, err
	}

	for _, scope := range scopes {
		// Creates a map of collections based on scope
		dcs, err := gc.createExchangeCollections(ctx, scope, deltas, control.Options{})
		if err != nil {
			user := scope.Get(selectors.ExchangeUser)
			return nil, support.WrapAndAppend(user[0], err, errs)
		}

		collections = append(collections, dcs...)
	}

	return collections, errs
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
		scopes      = odb.DiscreteScopes(gc.GetUsers())
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
				gc,
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
