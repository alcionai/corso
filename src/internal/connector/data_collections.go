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
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Data Collections
// ---------------------------------------------------------------------------

// DataCollections utility function to launch backup operations for exchange and onedrive
func (gc *GraphConnector) DataCollections(ctx context.Context, sels selectors.Selector) ([]data.Collection, error) {
	ctx, end := D.Span(ctx, "gc:dataCollections", D.Index("service", sels.Service.String()))
	defer end()

	err := verifyBackupInputs(sels, gc.Users)
	if err != nil {
		return nil, err
	}

	switch sels.Service {
	case selectors.ServiceExchange:
		return gc.ExchangeDataCollection(ctx, sels)
	case selectors.ServiceOneDrive:
		return gc.OneDriveDataCollections(ctx, sels)
	case selectors.ServiceSharePoint:
		colls, err := sharepoint.DataCollections(ctx, sels, gc.GetSiteIds(), gc.credentials.AzureTenantID, gc)
		if err != nil {
			return nil, err
		}

		for range colls {
			gc.incrementAwaitingMessages()
		}

		return colls, nil
	default:
		return nil, errors.Errorf("service %s not supported", sels)
	}
}

func verifyBackupInputs(sel selectors.Selector, mapOfUsers map[string]string) error {
	var personnel []string

	// retrieve users from selectors
	switch sel.Service {
	case selectors.ServiceExchange:
		backup, err := sel.ToExchangeBackup()
		if err != nil {
			return err
		}

		for _, scope := range backup.Scopes() {
			temp := scope.Get(selectors.ExchangeUser)
			personnel = append(personnel, temp...)
		}
	case selectors.ServiceOneDrive:
		backup, err := sel.ToOneDriveBackup()
		if err != nil {
			return err
		}

		for _, user := range backup.Scopes() {
			temp := user.Get(selectors.OneDriveUser)
			personnel = append(personnel, temp...)
		}

	default:
		return errors.New("service %s not supported")
	}

	// verify personnel
	normUsers := map[string]struct{}{}

	for k := range mapOfUsers {
		normUsers[strings.ToLower(k)] = struct{}{}
	}

	for _, user := range personnel {
		if _, ok := normUsers[strings.ToLower(user)]; !ok {
			return fmt.Errorf("%s user not found within tenant", user)
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
) ([]*exchange.Collection, error) {
	var (
		errs           *multierror.Error
		users          = scope.Get(selectors.ExchangeUser)
		allCollections = make([]*exchange.Collection, 0)
	)

	// Create collection of ExchangeDataCollection
	for _, user := range users {
		collections := make(map[string]*exchange.Collection)

		qp := graph.QueryParams{
			Category:      scope.Category().PathType(),
			ResourceOwner: user,
			FailFast:      gc.failFast,
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
			scope)

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

	for _, scope := range scopes {
		// Creates a map of collections based on scope
		dcs, err := gc.createExchangeCollections(ctx, scope)
		if err != nil {
			user := scope.Get(selectors.ExchangeUser)
			return nil, support.WrapAndAppend(user[0], err, errs)
		}

		for _, collection := range dcs {
			collections = append(collections, collection)
		}
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
				&gc.graphService,
				gc.UpdateStatus,
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
