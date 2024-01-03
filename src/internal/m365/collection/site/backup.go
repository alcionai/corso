package site

import (
	"context"
	"fmt"
	stdpath "path"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Library collections")

	var (
		collections = []data.BackupCollection{}
		colls       = drive.NewCollections(
			bh,
			tenantID,
			bpc.ProtectedResource,
			su,
			bpc.Options,
			counter)
	)

	msg := fmt.Sprintf(
		"%s (%s)",
		path.LibrariesCategory.HumanString(),
		stdpath.Base(bpc.ProtectedResource.Name()))

	pcfg := observe.ProgressCfg{
		Indent:            1,
		CompletionMessage: func() string { return fmt.Sprintf("(found %d items)", colls.NumItems) },
	}
	progressBar := observe.MessageWithCompletion(ctx, pcfg, msg)
	close(progressBar)

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
	counter *count.Bus,
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
		creds.AzureClientSecret,
		counter)
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
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "creating page collection path"))
		}

		collection := NewPrefetchCollection(
			nil,
			dir,
			nil,
			nil,
			ac,
			scope,
			su,
			bpc.Options,
			counter.Local())
		collection.SetBetaService(betaService)
		collection.AddItem(tuple.ID, time.Now())

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

func CollectLists(
	ctx context.Context,
	bh backupHandler,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	tenantID string,
	scope selectors.SharePointScope,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	logger.Ctx(ctx).Debug("Creating SharePoint List Collections")

	var (
		collection data.BackupCollection
		el         = errs.Local()
		cl         = counter.Local()
		spcs       = make([]data.BackupCollection, 0)
		spcsMap    = make(map[string]data.BackupCollection)
		cfg        = api.CallConfig{Select: idAnd("list", "lastModifiedDateTime")}
		currPaths  = map[string]string{}
		prevPath   path.Path
	)

	dps, canUsePreviousBackup, err := parseListsMetadataCollections(ctx, path.ListsCategory, bpc.MetadataCollections)
	if err != nil {
		return nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePreviousBackup)

	tombstones := makeTombstones(dps)

	lists, err := bh.GetItems(ctx, cfg)
	if err != nil {
		return nil, false, err
	}

	for _, list := range lists {
		if el.Failure() != nil {
			break
		}

		if api.SkipListTemplates.HasKey(ptr.Val(list.GetList().GetTemplate())) {
			continue
		}

		var (
			listID      = ptr.Val(list.GetId())
			storageDir  = path.Elements{listID}
			dp          = dps[storageDir.String()]
			prevPathStr = dp.Path
		)

		delete(tombstones, listID)

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				err = clues.StackWC(ctx, err).Label(count.BadPrevPath)
				logger.CtxErr(ctx, err).Error("parsing prev path")

				return nil, false, err
			}
		}

		currPath, err := bh.canonicalPath(storageDir, tenantID)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "creating list collection path"))
		}

		modTime := ptr.Val(list.GetLastModifiedDateTime())

		lazyFetchCol := NewLazyFetchCollection(
			bh,
			currPath,
			prevPath,
			storageDir.Builder(),
			su,
			cl)

		lazyFetchCol.AddItem(
			ptr.Val(list.GetId()),
			modTime)

		collection = lazyFetchCol

		// Always use lazyFetchCol.
		// In case we receive zero mod time from graph fallback to prefetchCol.
		if modTime.IsZero() {
			prefetchCol := NewPrefetchCollection(
				bh,
				currPath,
				prevPath,
				storageDir.Builder(),
				ac,
				scope,
				su,
				bpc.Options,
				counter.Local())

			prefetchCol.AddItem(
				ptr.Val(list.GetId()),
				modTime)

			collection = prefetchCol
		}

		spcsMap[storageDir.String()] = collection

		currPaths[storageDir.String()] = currPath.String()
	}

	handleTombstones(ctx, bpc, tombstones, spcsMap, counter, el)

	pathPrefix, err := path.BuildMetadata(
		tenantID,
		bpc.ProtectedResource.ID(),
		path.SharePointService,
		path.ListsCategory,
		false)
	if err != nil {
		return nil, false, clues.WrapWC(ctx, err, "making metadata path prefix").
			Label(count.BadPathPrefix)
	}

	col, err := graph.MakeMetadataCollection(
		pathPrefix,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(metadata.PreviousPathFileName, currPaths),
		},
		su,
		counter.Local())
	if err != nil {
		return nil, false, clues.WrapWC(ctx, err, "making metadata collection")
	}

	spcsMap["metadata"] = col

	for _, spc := range spcsMap {
		spcs = append(spcs, spc)
	}

	return spcs, canUsePreviousBackup, el.Failure()
}

func handleTombstones(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	tombstones map[string]string,
	spcsMap map[string]data.BackupCollection,
	counter *count.Bus,
	el *fault.Bus,
) {
	for id, p := range tombstones {
		if el.Failure() != nil {
			return
		}

		ctx := clues.Add(ctx, "tombstone_id", id)

		if spcsMap[id] != nil {
			err := clues.NewWC(ctx, "conflict: tombstone exists for a live collection").Label(count.CollectionTombstoneConflict)
			el.AddRecoverable(ctx, err)

			continue
		}

		if len(p) == 0 {
			continue
		}

		prevPath, err := pathFromPrevString(p)
		if err != nil {
			err := clues.StackWC(ctx, err).Label(count.BadPrevPath)
			logger.CtxErr(ctx, err).Error("parsing tombstone prev path")

			continue
		}

		spcsMap[id] = data.NewTombstoneCollection(prevPath, bpc.Options, counter.Local())
	}
}

func idAnd(ss ...string) []string {
	id := []string{"id"}

	if len(ss) == 0 {
		return id
	}

	return append(id, ss...)
}
