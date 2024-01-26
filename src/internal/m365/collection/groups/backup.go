package groups

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
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

// TODO: incremental support for channels

func CreateCollections[C graph.GetIDer, I groupsItemer](
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	bh backupHandler[C, I],
	tenantID string,
	scope selectors.GroupsScope,
	su support.StatusUpdater,
	useLazyReader bool,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	var (
		allCollections = make([]data.BackupCollection, 0)
		category       = scope.Category().PathType()
		qp             = graph.QueryParams{
			Category:          category,
			ProtectedResource: bpc.ProtectedResource,
			TenantID:          tenantID,
		}
	)

	cdps, canUsePreviousBackup, err := parseMetadataCollections(ctx, bpc.MetadataCollections)
	if err != nil {
		return nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePreviousBackup)

	cc := api.CallConfig{
		CanMakeDeltaQueries: bh.canMakeDeltaQueries(),
	}

	containers, err := bh.getContainers(ctx, cc)
	if err != nil {
		return nil, false, clues.Stack(err)
	}

	counter.Add(count.Channels, int64(len(containers)))

	collections, err := populateCollections[C, I](
		ctx,
		qp,
		bh,
		su,
		containers,
		scope,
		cdps[scope.Category().PathType()],
		useLazyReader,
		bpc.Options,
		counter,
		errs)
	if err != nil {
		return nil, false, clues.Wrap(err, "filling collections")
	}

	for _, coll := range collections {
		allCollections = append(allCollections, coll)
	}

	return allCollections, canUsePreviousBackup, nil
}

func populateCollections[C graph.GetIDer, I groupsItemer](
	ctx context.Context,
	qp graph.QueryParams,
	bh backupHandler[C, I],
	statusUpdater support.StatusUpdater,
	containers []container[C],
	scope selectors.GroupsScope,
	dps metadata.DeltaPaths,
	useLazyReader bool,
	ctrlOpts control.Options,
	counter *count.Bus,
	errs *fault.Bus,
) (map[string]data.BackupCollection, error) {
	var (
		// channel ID -> BackupCollection.
		collections = map[string]data.BackupCollection{}
		// channel ID -> delta url or folder path lookups
		deltaURLs = map[string]string{}
		currPaths = map[string]string{}
		// copy of previousPaths.  every channel present in the slice param
		// gets removed from this map; the remaining channels at the end of
		// the process have been deleted.
		tombstones = makeTombstones(dps)
		el         = errs.Local()
	)

	logger.Ctx(ctx).Infow("filling collections", "len_deltapaths", len(dps))

	for _, c := range containers {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		var (
			cl          = counter.Local()
			cID         = ptr.Val(c.container.GetId())
			err         error
			dp          = dps[c.storageDirFolders.String()]
			prevDelta   = dp.Delta
			prevPathStr = dp.Path // do not log: pii; log prevPath instead
			prevPath    path.Path
			ictx        = clues.Add(
				ctx,
				"collection_path", c,
				"previous_delta", pii.SafeURL{
					URL:           prevDelta,
					SafePathElems: graph.SafeURLPathParams,
					SafeQueryKeys: graph.SafeURLQueryParams,
				})
		)

		ictx = clues.AddLabelCounter(ictx, cl.PlainAdder())

		delete(tombstones, cID)

		// Only create a collection if the path matches the scope.
		if !bh.includeContainer(c.container, scope) {
			cl.Inc(count.SkippedContainers)
			continue
		}

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				err = clues.StackWC(ctx, err).Label(count.BadPrevPath)
				logger.CtxErr(ictx, err).Error("parsing prev path")
				// if the previous path is unusable, then the delta must be, too.
				prevDelta = ""
			}
		}

		ictx = clues.Add(ictx, "previous_path", prevPath)

		// if the channel has no email property, it is unable to process delta tokens
		// and will return an error if a delta token is queried.
		cc := api.CallConfig{
			CanMakeDeltaQueries: bh.canMakeDeltaQueries() && c.canMakeDeltaQueries,
		}

		addAndRem, err := bh.getContainerItemIDs(ctx, c.storageDirFolders, prevDelta, cc)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		removed := str.SliceToMap(addAndRem.Removed)

		cl.Add(count.ItemsAdded, int64(len(addAndRem.Added)))
		cl.Add(count.ItemsRemoved, int64(len(removed)))

		if len(addAndRem.DU.URL) > 0 {
			deltaURLs[c.storageDirFolders.String()] = addAndRem.DU.URL
		} else if !addAndRem.DU.Reset {
			logger.Ctx(ictx).Info("missing delta url")
		}

		currPath, err := bh.canonicalPath(c.storageDirFolders)
		if err != nil {
			err = clues.StackWC(ctx, err).Label(count.BadCollPath)
			el.AddRecoverable(ctx, err)

			continue
		}

		// Remove any deleted IDs from the set of added IDs because items that are
		// deleted and then restored will have a different ID than they did
		// originally.
		for remove := range removed {
			delete(addAndRem.Added, remove)
		}

		edc := NewCollection(
			data.NewBaseCollection(
				currPath,
				prevPath,
				c.humanLocation.Builder(),
				ctrlOpts,
				addAndRem.DU.Reset,
				cl),
			bh,
			qp.ProtectedResource.ID(),
			addAndRem.Added,
			removed,
			c,
			statusUpdater,
			useLazyReader)

		collections[c.storageDirFolders.String()] = edc

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[c.storageDirFolders.String()] = currPath.String()
	}

	// A tombstone is a channel that needs to be marked for deletion.
	// The only situation where a tombstone should appear is if the channel exists
	// in the `previousPath` set, but does not exist in the enumeration.
	for id, p := range tombstones {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		var (
			err  error
			ictx = clues.Add(ctx, "tombstone_id", id)
		)

		if collections[id] != nil {
			err := clues.NewWC(ictx, "conflict: tombstone exists for a live collection").Label(count.CollectionTombstoneConflict)
			el.AddRecoverable(ctx, err)

			continue
		}

		// only occurs if it was a new folder that we picked up during the container
		// resolver phase that got deleted in flight by the time we hit this stage.
		if len(p) == 0 {
			continue
		}

		prevPath, err := pathFromPrevString(p)
		if err != nil {
			err := clues.StackWC(ctx, err).Label(count.BadPrevPath)
			// technically shouldn't ever happen.  But just in case...
			logger.CtxErr(ictx, err).Error("parsing tombstone prev path")

			continue
		}

		collections[id] = data.NewTombstoneCollection(prevPath, ctrlOpts, counter.Local())
	}

	logger.Ctx(ctx).Infow(
		"adding metadata collection entries",
		"num_deltas_entries", len(deltaURLs),
		"num_paths_entries", len(collections))

	pathPrefix, err := path.BuildMetadata(
		qp.TenantID,
		qp.ProtectedResource.ID(),
		path.GroupsService,
		qp.Category,
		false)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "making metadata path prefix").
			Label(count.BadPathPrefix)
	}

	col, err := graph.MakeMetadataCollection(
		pathPrefix,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(metadata.PreviousPathFileName, currPaths),
			graph.NewMetadataEntry(metadata.DeltaURLsFileName, deltaURLs),
		},
		statusUpdater,
		counter.Local())
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "making metadata collection")
	}

	collections["metadata"] = col

	return collections, el.Failure()
}
