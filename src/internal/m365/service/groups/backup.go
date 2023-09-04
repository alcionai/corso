package groups

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	"github.com/alcionai/corso/src/internal/m365/collection/site"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
	b, err := bpc.Selector.ToGroupsBackup()
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "groupsDataCollection: parsing selector")
	}

	var (
		el                   = errs.Local()
		collections          = []data.BackupCollection{}
		categories           = map[path.CategoryType]struct{}{}
		ssmb                 = prefixmatcher.NewStringSetBuilder()
		canUsePreviousBackup bool
		sites                = map[string]string{}
	)

	ctx = clues.Add(
		ctx,
		"group_id", clues.Hide(bpc.ProtectedResource.ID()),
		"group_name", clues.Hide(bpc.ProtectedResource.Name()))

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		progressBar := observe.MessageWithCompletion(
			ctx,
			observe.Bulletf("%s", scope.Category().PathType()))
		defer close(progressBar)

		var dbcs []data.BackupCollection

		switch scope.Category().PathType() {
		case path.LibrariesCategory:
			// TODO(meain): Private channels get a separate SharePoint
			// site. We should also back those up and not just the
			// default one.
			resp, err := ac.Groups().GetRootSite(ctx, bpc.ProtectedResource.ID())
			if err != nil {
				return nil, nil, false, err
			}

			sites[ptr.Val(resp.GetId())] = ptr.Val(resp.GetName())

			pr := idname.NewProvider(ptr.Val(resp.GetId()), ptr.Val(resp.GetName()))
			sbpc := inject.BackupProducerConfig{
				LastBackupVersion: bpc.LastBackupVersion,
				Options:           bpc.Options,
				ProtectedResource: pr,
				Selector:          bpc.Selector,
			}

			dbcs, canUsePreviousBackup, err = site.CollectLibraries(
				ctx,
				sbpc,
				drive.NewGroupBackupHandler(
					bpc.ProtectedResource.ID(),
					ptr.Val(resp.GetId()),
					ac.Drives(),
					scope,
				),
				creds.AzureTenantID,
				ssmb,
				su,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}

		case path.ChannelMessagesCategory:
			dbcs, err = groups.CreateCollections(
				ctx,
				bpc,
				groups.NewChannelBackupHandler(bpc.ProtectedResource.ID(), ac.Channels()),
				creds.AzureTenantID,
				scope,
				su,
				errs)
			if err != nil {
				el.AddRecoverable(ctx, err)
				continue
			}
		}

		collections = append(collections, dbcs...)

		categories[scope.Category().PathType()] = struct{}{}
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			creds.AzureTenantID,
			bpc.ProtectedResource.ID(),
			path.GroupsService,
			categories,
			su,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	// Add metadata about sites
	p, err := path.Builder{}.ToServiceCategoryMetadataPath(
		creds.AzureTenantID,
		bpc.ProtectedResource.ID(),
		path.GroupsService,
		path.LibrariesCategory,
		false)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "making metadata path")
	}

	// TODO(meain): Should we store this one level above?
	p, err = p.Append(false, odConsts.SitesPathDir)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "appending sites to metadata path")
	}

	md, err := graph.MakeMetadataCollection(
		p,
		[]graph.MetadataCollectionEntry{
			// TODO(meain): Finalize on the name for file. If we are
			// storing it one level about, we will not be able to use
			// this name as it will be taken up by the folder.
			// TODO(meain): Or should this be previouspath file?
			graph.NewMetadataEntry(graph.SitesFileName, sites),
		},
		su)

	collections = append(collections, md)

	return collections, ssmb.ToReader(), canUsePreviousBackup, el.Failure()
}
