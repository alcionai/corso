package sharepoint

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	creds account.M365Config,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	b, err := bpc.Selector.ToSharePointBackup()
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "sharePointDataCollection: parsing selector")
	}

	var (
		el                   = errs.Local()
		collections          = []data.BackupCollection{}
		categories           = map[path.CategoryType]struct{}{}
		ssmb                 = prefixmatcher.NewStringSetBuilder()
		canUsePreviousBackup bool
	)

	ctx = clues.Add(
		ctx,
		"site_id", clues.Hide(bpc.ProtectedResource.ID()),
		"site_url", clues.Hide(bpc.ProtectedResource.Name()))

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		progressBar := observe.MessageWithCompletion(
			ctx,
			observe.Bulletf("%s", scope.Category().PathType()))
		defer close(progressBar)

		var spcs []data.BackupCollection

		switch scope.Category().PathType() {
		case path.ListsCategory:
			spcs, err = collectLists(
				ctx,
				bpc,
				ac,
				creds.AzureTenantID,
				su,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

			// Lists don't make use of previous metadata
			// TODO: Revisit when we add support of lists
			canUsePreviousBackup = true

		case path.LibrariesCategory:
			spcs, canUsePreviousBackup, err = collectLibraries(
				ctx,
				bpc,
				ac.Drives(),
				creds.AzureTenantID,
				ssmb,
				scope,
				su,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

		case path.PagesCategory:
			spcs, err = collectPages(
				ctx,
				bpc,
				creds,
				ac,
				su,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

			// Lists don't make use of previous metadata
			// TODO: Revisit when we add support of pages
			canUsePreviousBackup = true
		}

		collections = append(collections, spcs...)

		categories[scope.Category().PathType()] = struct{}{}
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			creds.AzureTenantID,
			bpc.ProtectedResource.ID(),
			path.SharePointService,
			categories,
			su,
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
	bpc inject.BackupProducerConfig,
	ac api.Client,
	tenantID string,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).Debug("Creating SharePoint List Collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
	)

	lists, err := preFetchLists(ctx, ac.Stable, bpc.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	for _, tuple := range lists {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			tenantID,
			bpc.ProtectedResource.ID(),
			path.SharePointService,
			path.ListsCategory,
			false,
			tuple.name)
		if err != nil {
			el.AddRecoverable(ctx, clues.Wrap(err, "creating list collection path").WithClues(ctx))
		}

		collection := NewCollection(
			dir,
			ac,
			List,
			su,
			bpc.Options)
		collection.AddJob(tuple.id)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

// collectLibraries constructs a onedrive Collections struct and Get()s
// all the drives associated with the site.
func collectLibraries(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ad api.Drives,
	tenantID string,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	scope selectors.SharePointScope,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Library collections")

	var (
		collections = []data.BackupCollection{}
		colls       = drive.NewCollections(
			&libraryBackupHandler{ad, scope},
			tenantID,
			bpc.ProtectedResource.ID(),
			su,
			bpc.Options)
	)

	odcs, canUsePreviousBackup, err := colls.Get(ctx, bpc.MetadataCollections, ssmb, errs)
	if err != nil {
		return nil, false, graph.Wrap(ctx, err, "getting library")
	}

	return append(collections, odcs...), canUsePreviousBackup, nil
}

// collectPages constructs a sharepoint Collections struct and Get()s the associated
// M365 IDs for the associated Pages.
func collectPages(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	creds account.M365Config,
	ac api.Client,
	su support.StatusUpdater,
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

	tuples, err := betaAPI.FetchPages(ctx, betaService, bpc.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			creds.AzureTenantID,
			bpc.ProtectedResource.ID(),
			path.SharePointService,
			path.PagesCategory,
			false,
			tuple.Name)
		if err != nil {
			el.AddRecoverable(ctx, clues.Wrap(err, "creating page collection path").WithClues(ctx))
		}

		collection := NewCollection(
			dir,
			ac,
			Pages,
			su,
			bpc.Options)
		collection.betaService = betaService
		collection.AddJob(tuple.ID)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}
