package sharepoint

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type statusUpdater interface {
	UpdateStatus(status *support.ConnectorOperationStatus)
}

type connector interface {
	statusUpdater

	Service() graph.Service
	IncrementAwaitingMessages()
}

// DataCollections returns a set of DataCollection which represents the SharePoint data
// for the specified user
func DataCollections(
	ctx context.Context,
	selector selectors.Selector,
	siteIDs []string,
	tenantID string,
	con connector,
) ([]data.Collection, error) {
	b, err := selector.ToSharePointBackup()
	if err != nil {
		return nil, errors.Wrap(err, "sharePointDataCollection: parsing selector")
	}

	var (
		scopes      = b.DiscreteScopes(siteIDs)
		collections = []data.Collection{}
		serv        = con.Service()
		errs        error
	)

	// for each scope that includes oneDrive items, get all
	for _, scope := range scopes {
		// Creates a slice of collections based on scope
		dcs, err := createSharePointCollections(ctx, serv, scope, tenantID, con)
		if err != nil {
			return nil, support.WrapAndAppend(scope.Get(selectors.SharePointSite)[0], err, errs)
		}

		for _, collection := range dcs {
			collections = append(collections, collection)
		}
	}

	for range collections {
		con.IncrementAwaitingMessages()
	}

	return collections, errs
}

// createSharePointCollections - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.SharePointScope
// determines the type of collections that are retrieved.
func createSharePointCollections(
	ctx context.Context,
	serv graph.Service,
	scope selectors.SharePointScope,
	tenantID string,
	updater statusUpdater,
) ([]data.Collection, error) {
	var (
		errs        *multierror.Error
		sites       = scope.Get(selectors.SharePointSite)
		category    = scope.Category().PathType()
		collections = make([]data.Collection, 0)
	)

	// Create collection of sharePoint data
	for _, site := range sites {
		foldersComplete, closer := observe.MessageWithCompletion(fmt.Sprintf("∙ %s - %s:", category, site))
		defer closer()
		defer close(foldersComplete)

		switch category {
		case path.FilesCategory: // TODO: better category for drives
			spcs, err := collectLibraries(
				ctx,
				serv,
				tenantID,
				site,
				scope,
				updater)
			if err != nil {
				return nil, support.WrapAndAppend(site, err, errs)
			}

			collections = append(collections, spcs...)
		}

		foldersComplete <- struct{}{}
	}

	return collections, errs.ErrorOrNil()
}

func collectLibraries(
	ctx context.Context,
	serv graph.Service,
	tenantID string,
	siteID string,
	scope selectors.SharePointScope,
	updater statusUpdater,
) ([]data.Collection, error) {
	var (
		collections = []data.Collection{}
		errs        error
	)

	logger.Ctx(ctx).With("site", siteID).Debug("Creating SharePoint Library collections")

	colls := onedrive.NewCollections(
		tenantID,
		siteID,
		onedrive.SharePointSource,
		folderMatcher{scope},
		serv,
		updater.UpdateStatus)

	odcs, err := colls.Get(ctx)
	if err != nil {
		return nil, support.WrapAndAppend(siteID, err, errs)
	}

	return append(collections, odcs...), errs
}

type folderMatcher struct {
	scope selectors.SharePointScope
}

func (fm folderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.SharePointFolder)
}

func (fm folderMatcher) Matches(dir string) bool {
	return fm.scope.Matches(selectors.SharePointFolder, dir)
}
