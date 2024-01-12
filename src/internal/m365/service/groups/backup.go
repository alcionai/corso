package groups

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

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
	"github.com/alcionai/corso/src/pkg/selectors"
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
		el                     = errs.Local()
		collections            = []data.BackupCollection{}
		categories             = map[path.CategoryType]struct{}{}
		globalItemIDExclusions = prefixmatcher.NewStringSetBuilder()
		sitesPreviousPaths     = map[string]string{}
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

	bc := backupCommon{ac, bpc, creds, group, sitesPreviousPaths, su}

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		cl := counter.Local()
		ictx := clues.AddLabelCounter(ctx, cl.PlainAdder())
		ictx = clues.Add(ictx, "category", scope.Category().PathType())

		var colls []data.BackupCollection

		switch scope.Category().PathType() {
		case path.LibrariesCategory:
			colls, err = backupLibraries(
				ictx,
				bc,
				scope,
				globalItemIDExclusions,
				cl,
				el)
		case path.ChannelMessagesCategory:
			colls, err = backupChannels(
				ictx,
				bc,
				scope,
				cl,
				el)
		case path.ConversationPostsCategory:
			colls, err = backupConversations(
				ictx,
				bc,
				scope,
				cl,
				el)
		}

		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		collections = append(collections, colls...)

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

	return collections, globalItemIDExclusions.ToReader(), el.Failure()
}

type backupCommon struct {
	apiCli             api.Client
	producerConfig     inject.BackupProducerConfig
	creds              account.M365Config
	group              models.Groupable
	sitesPreviousPaths map[string]string
	statusUpdater      support.StatusUpdater
}

func backupLibraries(
	ctx context.Context,
	bc backupCommon,
	scope selectors.GroupsScope,
	globalItemIDExclusions *prefixmatcher.StringSetMatchBuilder,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	var (
		colls = []data.BackupCollection{}
		el    = errs.Local()
	)

	sites, err := bc.apiCli.Groups().GetAllSites(
		ctx,
		bc.producerConfig.ProtectedResource.ID(),
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	counter.Add(count.Sites, int64(len(sites)))

	siteMetadataCollection := map[string][]data.RestoreCollection{}

	// Once we have metadata collections for chat as well, we will have to filter those out
	for _, c := range bc.producerConfig.MetadataCollections {
		siteID := c.FullPath().Elements().Last()
		siteMetadataCollection[siteID] = append(siteMetadataCollection[siteID], c)
	}

	for _, s := range sites {
		if el.Failure() != nil {
			break
		}

		var (
			cl   = counter.Local()
			pr   = idname.NewProvider(ptr.Val(s.GetId()), ptr.Val(s.GetWebUrl()))
			sbpc = inject.BackupProducerConfig{
				LastBackupVersion:   bc.producerConfig.LastBackupVersion,
				Options:             bc.producerConfig.Options,
				ProtectedResource:   pr,
				Selector:            bc.producerConfig.Selector,
				MetadataCollections: siteMetadataCollection[ptr.Val(s.GetId())],
			}
			bh = drive.NewGroupBackupHandler(
				bc.producerConfig.ProtectedResource.ID(),
				ptr.Val(s.GetId()),
				bc.apiCli.Drives(),
				scope)
		)

		ictx := clues.Add(
			ctx,
			"site_id", ptr.Val(s.GetId()),
			"site_weburl", graph.LoggableURL(ptr.Val(s.GetWebUrl())))

		sp, err := bh.SitePathPrefix(bc.creds.AzureTenantID)
		if err != nil {
			return nil, clues.WrapWC(ictx, err, "getting site path").Label(count.BadPathPrefix)
		}

		bc.sitesPreviousPaths[ptr.Val(s.GetId())] = sp.String()

		cs, canUsePreviousBackup, err := site.CollectLibraries(
			ictx,
			sbpc,
			bh,
			bc.creds.AzureTenantID,
			globalItemIDExclusions,
			bc.statusUpdater,
			cl,
			errs)
		if err != nil {
			el.AddRecoverable(ictx, err)
			continue
		}

		if !canUsePreviousBackup {
			colls = append(colls, data.NewTombstoneCollection(sp, control.Options{}, cl))
		}

		colls = append(colls, cs...)
	}

	return colls, el.Failure()
}

func backupChannels(
	ctx context.Context,
	bc backupCommon,
	scope selectors.GroupsScope,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	var (
		colls                []data.BackupCollection
		canUsePreviousBackup bool
	)

	progressMessage := observe.MessageWithCompletion(
		ctx,
		observe.ProgressCfg{
			Indent: 1,
			// TODO(meain): Use number of messages and not channels
			CompletionMessage: func() string { return fmt.Sprintf("(found %d channels)", len(colls)) },
		},
		scope.Category().PathType().HumanString())
	defer close(progressMessage)

	if !api.IsTeam(ctx, bc.group) {
		return colls, nil
	}

	bh := groups.NewChannelBackupHandler(
		bc.producerConfig.ProtectedResource.ID(),
		bc.apiCli.Channels())

	// Always disable lazy reader for channels until #4321 support is added
	useLazyReader := false

	colls, canUsePreviousBackup, err := groups.CreateCollections(
		ctx,
		bc.producerConfig,
		bh,
		bc.creds.AzureTenantID,
		scope,
		bc.statusUpdater,
		useLazyReader,
		counter,
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	if !canUsePreviousBackup {
		tp, err := bh.PathPrefix(bc.creds.AzureTenantID)
		if err != nil {
			err = clues.WrapWC(ctx, err, "getting message path").Label(count.BadPathPrefix)
			return nil, err
		}

		colls = append(colls, data.NewTombstoneCollection(tp, control.Options{}, counter))
	}

	return colls, nil
}

func backupConversations(
	ctx context.Context,
	bc backupCommon,
	scope selectors.GroupsScope,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	groupEmail := strings.Clone(ptr.Val(bc.group.GetMail()))
	// This is unlikely, but if it does happen in the wild, we should investigate it.
	if len(groupEmail) == 0 {
		return nil, clues.New("group has no mail address")
	}

	var (
		bh = groups.NewConversationBackupHandler(
			bc.producerConfig.ProtectedResource.ID(),
			bc.apiCli.Conversations(),
			groupEmail)
		colls []data.BackupCollection
	)

	progressMessage := observe.MessageWithCompletion(
		ctx,
		observe.ProgressCfg{
			Indent:            1,
			CompletionMessage: func() string { return fmt.Sprintf("(found %d conversations)", len(colls)) },
		},
		scope.Category().PathType().HumanString())
	defer close(progressMessage)

	useLazyReader := !bc.producerConfig.Options.ToggleFeatures.DisableLazyItemReader

	colls, canUsePreviousBackup, err := groups.CreateCollections(
		ctx,
		bc.producerConfig,
		bh,
		bc.creds.AzureTenantID,
		scope,
		bc.statusUpdater,
		useLazyReader,
		counter,
		errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	if !canUsePreviousBackup {
		tp, err := bh.PathPrefix(bc.creds.AzureTenantID)
		if err != nil {
			err = clues.WrapWC(ctx, err, "getting conversation path").Label(count.BadPathPrefix)
			return nil, err
		}

		colls = append(colls, data.NewTombstoneCollection(tp, control.Options{}, counter))
	}

	return colls, nil
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
