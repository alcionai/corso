package sharepoint

import (
	"context"
	"net/http"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/sharepoint/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	m365api "github.com/alcionai/corso/src/pkg/services/m365/api"
)

type statusUpdater interface {
	UpdateStatus(status *support.ConnectorOperationStatus)
}

// DataCollections returns a set of DataCollection which represents the SharePoint data
// for the specified user
func DataCollections(
	ctx context.Context,
	itemClient *http.Client,
	selector selectors.Selector,
	creds account.M365Config,
	serv graph.Servicer,
	su statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	b, err := selector.ToSharePointBackup()
	if err != nil {
		return nil, nil, clues.Wrap(err, "sharePointDataCollection: parsing selector")
	}

	var (
		el          = errs.Local()
		site        = b.DiscreteOwner
		collections = []data.BackupCollection{}
		categories  = map[path.CategoryType]struct{}{}
	)

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		foldersComplete, closer := observe.MessageWithCompletion(
			ctx,
			observe.Bulletf("%s", scope.Category().PathType()))
		defer closer()
		defer close(foldersComplete)

		var spcs []data.BackupCollection

		switch scope.Category().PathType() {
		case path.ListsCategory:
			spcs, err = collectLists(
				ctx,
				serv,
				creds.AzureTenantID,
				site,
				su,
				ctrlOpts,
				errs)
			if err != nil {
				el.AddRecoverable(err)
				continue
			}

		case path.LibrariesCategory:
			spcs, _, err = collectLibraries(
				ctx,
				itemClient,
				serv,
				creds.AzureTenantID,
				site,
				scope,
				su,
				ctrlOpts,
				errs)
			if err != nil {
				el.AddRecoverable(err)
				continue
			}

		case path.PagesCategory:
			spcs, err = collectPages(
				ctx,
				creds,
				serv,
				site,
				su,
				ctrlOpts,
				errs)
			if err != nil {
				el.AddRecoverable(err)
				continue
			}
		}

		collections = append(collections, spcs...)
		foldersComplete <- struct{}{}

		categories[scope.Category().PathType()] = struct{}{}
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			nil,
			creds.AzureTenantID,
			site,
			path.SharePointService,
			categories,
			su.UpdateStatus,
			errs)
		if err != nil {
			return nil, nil, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, nil, el.Failure()
}

func collectLists(
	ctx context.Context,
	serv graph.Servicer,
	tenantID, siteID string,
	updater statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).With("site", siteID).Debug("Creating SharePoint List Collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
	)

	lists, err := preFetchLists(ctx, serv, siteID)
	if err != nil {
		return nil, err
	}

	for _, tuple := range lists {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			tenantID,
			siteID,
			path.SharePointService,
			path.ListsCategory,
			false,
			tuple.name)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "creating list collection path").WithClues(ctx))
		}

		collection := NewCollection(dir, serv, List, updater.UpdateStatus, ctrlOpts)
		collection.AddJob(tuple.id)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

// collectLibraries constructs a onedrive Collections struct and Get()s
// all the drives associated with the site.
func collectLibraries(
	ctx context.Context,
	itemClient *http.Client,
	serv graph.Servicer,
	tenantID, siteID string,
	scope selectors.SharePointScope,
	updater statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]map[string]struct{}, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Library collections")

	var (
		collections = []data.BackupCollection{}
		colls       = onedrive.NewCollections(
			itemClient,
			tenantID,
			siteID,
			onedrive.SharePointSource,
			folderMatcher{scope},
			serv,
			updater.UpdateStatus,
			ctrlOpts)
	)

	// TODO(ashmrtn): Pass previous backup metadata when SharePoint supports delta
	// token-based incrementals.
	odcs, excludes, err := colls.Get(ctx, nil, errs)
	if err != nil {
		return nil, nil, graph.Wrap(ctx, err, "getting library")
	}

	return append(collections, odcs...), excludes, nil
}

// collectPages constructs a sharepoint Collections struct and Get()s the associated
// M365 IDs for the associated Pages.
func collectPages(
	ctx context.Context,
	creds account.M365Config,
	serv graph.Servicer,
	siteID string,
	updater statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Pages collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
	)

	// make the betaClient
	// Need to receive From DataCollection Call
	adpt, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret)
	if err != nil {
		return nil, clues.Wrap(err, "creating azure client adapter")
	}

	betaService := m365api.NewBetaService(adpt)

	tuples, err := api.FetchPages(ctx, betaService, siteID)
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			creds.AzureTenantID,
			siteID,
			path.SharePointService,
			path.PagesCategory,
			false,
			tuple.Name)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "creating page collection path").WithClues(ctx))
		}

		collection := NewCollection(dir, serv, Pages, updater.UpdateStatus, ctrlOpts)
		collection.betaService = betaService
		collection.AddJob(tuple.ID)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

type folderMatcher struct {
	scope selectors.SharePointScope
}

func (fm folderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.SharePointLibraryFolder)
}

func (fm folderMatcher) Matches(dir string) bool {
	return fm.scope.Matches(selectors.SharePointLibraryFolder, dir)
}
