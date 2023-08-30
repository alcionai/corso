package site

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/graph"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// CollectLibraries constructs a onedrive Collections struct and Get()s
// all the drives associated with the site.
func CollectLibraries(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	bh drive.BackupHandler,
	tenantID string,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Library collections")

	var (
		collections = []data.BackupCollection{}
		colls       = drive.NewCollections(
			bh,
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

// CollectPages constructs a sharepoint Collections struct and Get()s the associated
// M365 IDs for the associated Pages.
func CollectPages(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	creds account.M365Config,
	ac api.Client,
	scope selectors.SharePointScope,
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
			scope,
			su,
			bpc.Options)
		collection.SetBetaService(betaService)
		collection.AddJob(tuple.ID)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

func CollectLists(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	tenantID string,
	scope selectors.SharePointScope,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).Debug("Creating SharePoint List Collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
	)

	lists, err := PreFetchLists(ctx, ac.Stable, bpc.ProtectedResource.ID())
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
			tuple.Name)
		if err != nil {
			el.AddRecoverable(ctx, clues.Wrap(err, "creating list collection path").WithClues(ctx))
		}

		collection := NewCollection(
			dir,
			ac,
			scope,
			su,
			bpc.Options)
		collection.AddJob(tuple.ID)

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}
