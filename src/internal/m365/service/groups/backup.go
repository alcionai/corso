package groups

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	kinject "github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	"github.com/alcionai/corso/src/internal/m365/collection/site"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

func ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	creds account.M365Config,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, error) {
	b, err := bpc.Selector.ToGroupsBackup()
	if err != nil {
		return nil, nil, clues.Wrap(err, "groupsDataCollection: parsing selector")
	}

	var (
		el                 = errs.Local()
		collections        = []data.BackupCollection{}
		categories         = map[path.CategoryType]struct{}{}
		ssmb               = prefixmatcher.NewStringSetBuilder()
		sitesPreviousPaths = map[string]string{}
	)

	ctx = clues.Add(
		ctx,
		"group_id", clues.Hide(bpc.ProtectedResource.ID()),
		"group_name", clues.Hide(bpc.ProtectedResource.Name()))

	group, err := ac.Groups().GetByID(
		ctx,
		bpc.ProtectedResource.ID(),
		api.CallConfig{})
	if err != nil {
		return nil, nil, clues.WrapWC(ctx, err, "getting group")
	}

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		cl := counter.Local()
		ictx := clues.AddLabelCounter(ctx, cl.PlainAdder())
		ictx = clues.Add(ictx, "category", scope.Category().PathType())

		var dbcs []data.BackupCollection

		switch scope.Category().PathType() {
		case path.LibrariesCategory:
			sites, err := ac.Groups().GetAllSites(
				ictx,
				bpc.ProtectedResource.ID(),
				errs)
			if err != nil {
				return nil, nil, err
			}

			cl.Add(count.Sites, int64(len(sites)))

			siteMetadataCollection := map[string][]data.RestoreCollection{}

			// Once we have metadata collections for chat as well, we will have to filter those out
			for _, c := range bpc.MetadataCollections {
				siteID := c.FullPath().Elements().Last()
				siteMetadataCollection[siteID] = append(siteMetadataCollection[siteID], c)
			}

			for _, s := range sites {
				var (
					scl  = cl.Local()
					pr   = idname.NewProvider(ptr.Val(s.GetId()), ptr.Val(s.GetWebUrl()))
					sbpc = inject.BackupProducerConfig{
						LastBackupVersion:   bpc.LastBackupVersion,
						Options:             bpc.Options,
						ProtectedResource:   pr,
						Selector:            bpc.Selector,
						MetadataCollections: siteMetadataCollection[ptr.Val(s.GetId())],
					}
					bh = drive.NewGroupBackupHandler(
						bpc.ProtectedResource.ID(),
						ptr.Val(s.GetId()),
						ac.Drives(),
						scope)
				)

				ictx = clues.Add(
					ictx,
					"site_id", ptr.Val(s.GetId()),
					"site_weburl", graph.LoggableURL(ptr.Val(s.GetWebUrl())))

				sp, err := bh.SitePathPrefix(creds.AzureTenantID)
				if err != nil {
					return nil, nil, clues.WrapWC(ictx, err, "getting site path").Label(count.BadPathPrefix)
				}

				sitesPreviousPaths[ptr.Val(s.GetId())] = sp.String()

				cs, canUsePreviousBackup, err := site.CollectLibraries(
					ictx,
					sbpc,
					bh,
					creds.AzureTenantID,
					ssmb,
					su,
					scl,
					errs)
				if err != nil {
					el.AddRecoverable(ictx, err)
					continue
				}

				if !canUsePreviousBackup {
					dbcs = append(dbcs, data.NewTombstoneCollection(sp, control.Options{}, scl))
				}

				dbcs = append(dbcs, cs...)
			}
		case path.ChannelMessagesCategory:
			var (
				cs                   []data.BackupCollection
				canUsePreviousBackup bool
				err                  error
			)

			pcfg := observe.ProgressCfg{
				Indent: 1,
				// TODO(meain): Use number of messages and not channels
				CompletionMessage: func() string { return fmt.Sprintf("(found %d channels)", len(cs)) },
			}
			progressBar := observe.MessageWithCompletion(ictx, pcfg, scope.Category().PathType().HumanString())

			if !api.IsTeam(ictx, group) {
				continue
			}

			bh := groups.NewChannelBackupHandler(bpc.ProtectedResource.ID(), ac.Channels())

			cs, canUsePreviousBackup, err = groups.CreateCollections(
				ictx,
				bpc,
				bh,
				creds.AzureTenantID,
				scope,
				su,
				cl,
				errs)
			if err != nil {
				el.AddRecoverable(ictx, err)
				continue
			}

			if !canUsePreviousBackup {
				tp, err := bh.PathPrefix(creds.AzureTenantID)
				if err != nil {
					return nil, nil, clues.WrapWC(ictx, err, "getting message path").Label(count.BadPathPrefix)
				}

				dbcs = append(dbcs, data.NewTombstoneCollection(tp, control.Options{}, cl))
			}

			dbcs = append(dbcs, cs...)

			close(progressBar)
		case path.ConversationPostsCategory:
			var (
				bh  = groups.NewConversationBackupHandler(bpc.ProtectedResource.ID(), ac.Conversations())
				cs  []data.BackupCollection
				err error
			)

			pcfg := observe.ProgressCfg{
				Indent:            1,
				CompletionMessage: func() string { return fmt.Sprintf("(found %d conversations)", len(cs)) },
			}
			progressBar := observe.MessageWithCompletion(ictx, pcfg, scope.Category().PathType().HumanString())

			cs, canUsePreviousBackup, err := groups.CreateCollections(
				ictx,
				bpc,
				bh,
				creds.AzureTenantID,
				scope,
				su,
				counter,
				errs)
			if err != nil {
				el.AddRecoverable(ictx, err)
				continue
			}

			if !canUsePreviousBackup {
				tp, err := bh.PathPrefix(creds.AzureTenantID)
				if err != nil {
					return nil, nil, clues.Wrap(err, "getting conversations path")
				}

				dbcs = append(dbcs, data.NewTombstoneCollection(tp, control.Options{}, counter))
			}

			dbcs = append(dbcs, cs...)

			close(progressBar)
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
			counter,
			errs)
		if err != nil {
			return nil, nil, err
		}

		collections = append(collections, baseCols...)
	}

	// Add metadata about sites
	md, err := getSitesMetadataCollection(
		creds.AzureTenantID,
		bpc.ProtectedResource.ID(),
		sitesPreviousPaths,
		su,
		counter)
	if err != nil {
		return nil, nil, err
	}

	collections = append(collections, md)

	counter.Add(count.Collections, int64(len(collections)))

	logger.Ctx(ctx).Infow("produced collections", "stats", counter.Values())

	return collections, ssmb.ToReader(), el.Failure()
}

// ---------------------------------------------------------------------------
// metadata
// ---------------------------------------------------------------------------

func getSitesMetadataCollection(
	tenantID, groupID string,
	sites map[string]string,
	su support.StatusUpdater,
	counter *count.Bus,
) (data.BackupCollection, error) {
	p, err := path.BuildMetadata(
		tenantID,
		groupID,
		path.GroupsService,
		path.LibrariesCategory,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path")
	}

	p, err = p.Append(false, odConsts.SitesPathDir)
	if err != nil {
		return nil, clues.Wrap(err, "appending sites to metadata path")
	}

	md, err := graph.MakeMetadataCollection(
		p,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(metadata.PreviousPathFileName, sites),
		},
		su,
		counter.Local())

	return md, err
}

func MetadataFiles(
	ctx context.Context,
	reason identity.Reasoner,
	r kinject.RestoreProducer,
	manID manifest.ID,
	errs *fault.Bus,
) ([][]string, error) {
	pth, err := path.BuildMetadata(
		reason.Tenant(),
		reason.ProtectedResource(),
		reason.Service(),
		reason.Category(),
		true,
		odConsts.SitesPathDir,
		metadata.PreviousPathFileName)
	if err != nil {
		return nil, err
	}

	dir, err := pth.Dir()
	if err != nil {
		return nil, clues.Wrap(err, "building metadata collection path")
	}

	dcs, err := r.ProduceRestoreCollections(
		ctx,
		string(manID),
		[]path.RestorePaths{{StoragePath: pth, RestorePath: dir}},
		nil,
		errs)
	if err != nil {
		return nil, err
	}

	sites, err := deserializeSiteMetadata(ctx, dcs)
	if err != nil {
		return nil, err
	}

	filePaths := [][]string{}

	for k := range sites {
		for _, fn := range metadata.AllMetadataFileNames() {
			filePaths = append(filePaths, []string{odConsts.SitesPathDir, k, fn})
		}
	}

	return filePaths, nil
}

func deserializeSiteMetadata(
	ctx context.Context,
	cols []data.RestoreCollection,
) (map[string]string, error) {
	logger.Ctx(ctx).Infow(
		"deserializing previous sites metadata",
		"num_collections", len(cols))

	var (
		prevFolders = map[string]string{}
		errs        = fault.New(true) // metadata item reads should not fail backup
	)

	for _, col := range cols {
		if errs.Failure() != nil {
			break
		}

		items := col.Items(ctx, errs)

		for breakLoop := false; !breakLoop; {
			select {
			case <-ctx.Done():
				return nil, clues.WrapWC(ctx, ctx.Err(), "deserializing previous sites metadata")

			case item, ok := <-items:
				if !ok {
					breakLoop = true
					break
				}

				var (
					err  error
					ictx = clues.Add(ctx, "item_uuid", item.ID())
				)

				switch item.ID() {
				case metadata.PreviousPathFileName:
					err = drive.DeserializeMap(item.ToReader(), prevFolders)

				default:
					logger.Ctx(ictx).Infow(
						"skipping unknown metadata file",
						"file_name", item.ID())

					continue
				}

				if err == nil {
					// Successful decode.
					continue
				}

				if err != nil {
					return nil, clues.StackWC(ictx, err)
				}
			}
		}
	}

	// if reads from items failed, return empty but no error
	if errs.Failure() != nil {
		logger.CtxErr(ctx, errs.Failure()).Info("reading metadata collection items")

		return map[string]string{}, nil
	}

	return prevFolders, nil
}
