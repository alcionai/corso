package sharepoint

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	betaAPI "github.com/alcionai/corso/src/internal/m365/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type statusUpdater interface {
	UpdateStatus(status *support.ConnectorOperationStatus)
}

// DataCollections returns a set of DataCollection which represents the SharePoint data
// for the specified user
func DataCollections(
	ctx context.Context,
	ac api.Client,
	selector selectors.Selector,
	site idname.Provider,
	metadata []data.RestoreCollection,
	creds account.M365Config,
	su statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	b, err := selector.ToSharePointBackup()
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "sharePointDataCollection: parsing selector")
	}

	ctx = clues.Add(
		ctx,
		"site_id", clues.Hide(site.ID()),
		"site_url", clues.Hide(site.Name()))

	var (
		el                   = errs.Local()
		collections          = []data.BackupCollection{}
		categories           = map[path.CategoryType]struct{}{}
		ssmb                 = prefixmatcher.NewStringSetBuilder()
		canUsePreviousBackup bool
	)

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		foldersComplete := observe.MessageWithCompletion(
			ctx,
			observe.Bulletf("%s", scope.Category().PathType()))
		defer close(foldersComplete)

		var spcs []data.BackupCollection

		switch scope.Category().PathType() {
		case path.ListsCategory:
			spcs, err = collectLists(
				ctx,
				ac,
				creds.AzureTenantID,
				site,
				su,
				ctrlOpts,
				errs)
			if err != nil {
				el.AddRecoverable(err)
				continue
			}

			// Lists don't make use of previous metadata
			// TODO: Revisit when we add support of lists
			canUsePreviousBackup = true

		case path.LibrariesCategory:
			spcs, canUsePreviousBackup, err = collectLibraries(
				ctx,
				ac.Drives(),
				creds.AzureTenantID,
				site,
				metadata,
				ssmb,
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
				ac,
				site,
				su,
				ctrlOpts,
				errs)
			if err != nil {
				el.AddRecoverable(err)
				continue
			}

			// Lists don't make use of previous metadata
			// TODO: Revisit when we add support of pages
			canUsePreviousBackup = true
		}

		collections = append(collections, spcs...)
		foldersComplete <- struct{}{}

		categories[scope.Category().PathType()] = struct{}{}
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			creds.AzureTenantID,
			site.ID(),
			path.SharePointService,
			categories,
			su.UpdateStatus,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, ssmb.ToReader(), canUsePreviousBackup, el.Failure()
}

func collectLists(
	ctx context.Context,
	ac api.Client,
	tenantID string,
	site idname.Provider,
	updater statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).Debug("Creating SharePoint List Collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
	)

	lists, err := preFetchLists(ctx, ac.Stable, site.ID())
	if err != nil {
		return nil, err
	}

	for _, tuple := range lists {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			tenantID,
			site.ID(),
			path.SharePointService,
			path.ListsCategory,
			false,
			tuple.name)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "creating list collection path").WithClues(ctx))
		}

		collection := NewCollection(
			dir,
			ac,
			List,
			updater.UpdateStatus,
			ctrlOpts)
		collection.AddJob(tuple.id)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

// collectLibraries constructs a onedrive Collections struct and Get()s
// all the drives associated with the site.
func collectLibraries(
	ctx context.Context,
	ad api.Drives,
	tenantID string,
	site idname.Provider,
	metadata []data.RestoreCollection,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	scope selectors.SharePointScope,
	updater statusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Library collections")

	var (
		collections = []data.BackupCollection{}
		colls       = onedrive.NewCollections(
			&libraryBackupHandler{ad},
			tenantID,
			site.ID(),
			folderMatcher{scope},
			updater.UpdateStatus,
			ctrlOpts)
	)

	odcs, canUsePreviousBackup, err := colls.Get(ctx, metadata, ssmb, errs)
	if err != nil {
		return nil, false, graph.Wrap(ctx, err, "getting library")
	}

	return append(collections, odcs...), canUsePreviousBackup, nil
}

// collectPages constructs a sharepoint Collections struct and Get()s the associated
// M365 IDs for the associated Pages.
func collectPages(
	ctx context.Context,
	creds account.M365Config,
	ac api.Client,
	site idname.Provider,
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

	betaService := betaAPI.NewBetaService(adpt)

	tuples, err := betaAPI.FetchPages(ctx, betaService, site.ID())
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			creds.AzureTenantID,
			site.ID(),
			path.SharePointService,
			path.PagesCategory,
			false,
			tuple.Name)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "creating page collection path").WithClues(ctx))
		}

		collection := NewCollection(
			dir,
			ac,
			Pages,
			updater.UpdateStatus,
			ctrlOpts)
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
