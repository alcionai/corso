package sharepoint

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type statusUpdater interface {
	UpdateStatus(status *support.ConnectorOperationStatus)
}

// DataCollections returns a set of DataCollection which represents the SharePoint data
// for the specified user
func DataCollections(
	ctx context.Context,
	selector selectors.Selector,
	tenantID string,
	serv graph.Servicer,
	su statusUpdater,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	b, err := selector.ToSharePointBackup()
	if err != nil {
		return nil, errors.Wrap(err, "sharePointDataCollection: parsing selector")
	}

	var (
		site        = b.DiscreteOwner
		collections = []data.Collection{}
		errs        error
	)

	for _, scope := range b.Scopes() {
		foldersComplete, closer := observe.MessageWithCompletion(ctx, observe.Bulletf(
			"%s - %s",
			scope.Category().PathType(), site))
		defer closer()
		defer close(foldersComplete)

		var spcs []data.Collection

		switch scope.Category().PathType() {
		case path.ListsCategory:
			spcs, err = collectLists(
				ctx,
				serv,
				tenantID,
				site,
				su,
				ctrlOpts)
			if err != nil {
				return nil, support.WrapAndAppend(site, err, errs)
			}

		case path.LibrariesCategory:
			spcs, err = collectLibraries(
				ctx,
				serv,
				tenantID,
				site,
				scope,
				su,
				ctrlOpts)
			if err != nil {
				return nil, support.WrapAndAppend(site, err, errs)
			}
		case path.PagesCategory:
			spcs, err = collectPages(
				ctx,
				serv,
				tenantID,
				site,
				scope,
				su,
				ctrlOpts)
			if err != nil {
				return nil, support.WrapAndAppend(site, err, errs)
			}
		}

		collections = append(collections, spcs...)
		foldersComplete <- struct{}{}
	}

	return collections, errs
}

func collectLists(
	ctx context.Context,
	serv graph.Servicer,
	tenantID, siteID string,
	updater statusUpdater,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	logger.Ctx(ctx).With("site", siteID).Debug("Creating SharePoint List Collections")

	spcs := make([]data.Collection, 0)

	tuples, err := preFetchLists(ctx, serv, siteID)
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		dir, err := path.Builder{}.Append(tuple.name).
			ToDataLayerSharePointPath(
				tenantID,
				siteID,
				path.ListsCategory,
				false)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create collection path for site: %s", siteID)
		}

		collection := NewCollection(dir, serv, List, updater.UpdateStatus, ctrlOpts)
		collection.AddJob(tuple.id)

		spcs = append(spcs, collection)
	}

	return spcs, nil
}

// collectLibraries constructs a onedrive Collections struct and Get()s
// all the drives associated with the site.
func collectLibraries(
	ctx context.Context,
	serv graph.Servicer,
	tenantID, siteID string,
	scope selectors.SharePointScope,
	updater statusUpdater,
	ctrlOpts control.Options,
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
		updater.UpdateStatus,
		ctrlOpts)

	odcs, err := colls.Get(ctx)
	if err != nil {
		return nil, support.WrapAndAppend(siteID, err, errs)
	}

	return append(collections, odcs...), errs
}

// collectPages constructs a sharepoint Collections struct and Get()s the associated
// M365 IDs for the associated Pages.
// TODO: Credentials necessary to create separate HTTP client
func collectPages(
	ctx context.Context,
	serv graph.Servicer,
	tenantID, siteID string,
	scope selectors.SharePointScope,
	updater statusUpdater,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	logger.Ctx(ctx).With("site", siteID).Debug("Creating SharePoint Pages collections")

	spcs := make([]data.Collection, 0)

	tuples, err := fetchPages(ctx, serv, siteID)
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		dir, err := path.Builder{}.Append(tuple.name).
			ToDataLayerSharePointPath(
				tenantID,
				siteID,
				path.PagesCategory,
				false)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create collection path for site: %s", siteID)
		}

		collection := NewCollection(dir, serv, Pages, updater.UpdateStatus, ctrlOpts)
		collection.AddJob(tuple.id)

		spcs = append(spcs, collection)
	}

	return spcs, nil
}

type folderMatcher struct {
	scope selectors.SharePointScope
}

func (fm folderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.SharePointLibrary)
}

func (fm folderMatcher) Matches(dir string) bool {
	return fm.scope.Matches(selectors.SharePointLibrary, dir)
}
