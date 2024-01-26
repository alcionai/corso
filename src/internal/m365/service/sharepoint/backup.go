package sharepoint

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/site"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type sharePointBackup struct{}

// NewBackup provides a struct that matches standard apis
// across m365/service handlers.
func NewBackup() *sharePointBackup {
	return &sharePointBackup{}
}

func (sharePointBackup) ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	creds account.M365Config,
	su support.StatusUpdater,
	counter *count.Bus,
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

		var spcs []data.BackupCollection

		qp := graph.QueryParams{
			ProtectedResource: bpc.ProtectedResource,
			TenantID:          creds.AzureClientID,
		}

		switch scope.Category().PathType() {
		case path.ListsCategory:
			bh := site.NewListsBackupHandler(qp, ac.Lists())

			spcs, canUsePreviousBackup, err = site.CollectLists(
				ctx,
				bh,
				bpc,
				ac,
				creds.AzureTenantID,
				scope,
				su,
				counter,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

		case path.LibrariesCategory:
			spcs, canUsePreviousBackup, err = site.CollectLibraries(
				ctx,
				bpc,
				drive.NewSiteBackupHandler(
					qp,
					ac.Drives(),
					scope,
					bpc.Selector.PathService()),
				creds.AzureTenantID,
				ssmb,
				su,
				counter,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

		case path.PagesCategory:
			spcs, err = site.CollectPages(
				ctx,
				bpc,
				creds,
				ac,
				scope,
				su,
				counter,
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
			counter,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, ssmb.ToReader(), canUsePreviousBackup, el.Failure()
}

// ListsMetadataFileNames only contains PreviousPathFileName
// and not DeltaURLsFileName because graph apis do not have delta support
// for Sharepoint Lists
func ListsMetadataFileNames() []string {
	return []string{metadata.PreviousPathFileName}
}
