package groups

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// TODO: incremental support
// multiple lines in this file are commented out so that
// we can focus on v0 backups and re-integrate them later
// for v1 incrementals.
// since these lines represent otherwise standard boilerplate,
// it's simpler to comment them for tracking than to delete
// and re-discover them later.

func CreateCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	bh backupHandler,
	tenantID string,
	scope selectors.GroupsScope,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	ctx = clues.Add(ctx, "category", scope.Category().PathType())

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

	catProgress := observe.MessageWithCompletion(
		ctx,
		observe.Bulletf("%s", qp.Category))
	defer close(catProgress)

	channels, err := bh.getChannels(ctx)
	if err != nil {
		return nil, false, clues.Stack(err)
	}

	collections, err := populateCollections(
		ctx,
		qp,
		bh,
		su,
		channels,
		scope,
		cdps[scope.Category().PathType()],
		bpc.Options,
		errs)
	if err != nil {
		return nil, false, clues.Wrap(err, "filling collections")
	}

	for _, coll := range collections {
		allCollections = append(allCollections, coll)
	}

	return allCollections, canUsePreviousBackup, nil
}

func populateCollections(
	ctx context.Context,
	qp graph.QueryParams,
	bh backupHandler,
	statusUpdater support.StatusUpdater,
	channels []models.Channelable,
	scope selectors.GroupsScope,
	dps metadata.DeltaPaths,
	ctrlOpts control.Options,
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

	logger.Ctx(ctx).Info("filling collections", "len_deltapaths", len(dps))

	for _, c := range channels {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		var (
			cID         = ptr.Val(c.GetId())
			cName       = ptr.Val(c.GetDisplayName())
			err         error
			dp          = dps[cID]
			prevDelta   = dp.Delta
			prevPathStr = dp.Path // do not log: pii; log prevPath instead
			prevPath    path.Path
			ictx        = clues.Add(
				ctx,
				"channel_id", cID,
				"previous_delta", pii.SafeURL{
					URL:           prevDelta,
					SafePathElems: graph.SafeURLPathParams,
					SafeQueryKeys: graph.SafeURLQueryParams,
				})
		)

		delete(tombstones, cID)

		// Only create a collection if the path matches the scope.
		if !bh.includeContainer(ictx, qp, c, scope) {
			continue
		}

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				logger.CtxErr(ictx, err).Error("parsing prev path")
				// if the previous path is unusable, then the delta must be, too.
				prevDelta = ""
			}
		}

		ictx = clues.Add(ictx, "previous_path", prevPath)

		added, removed, du, err := bh.getChannelMessageIDsDelta(ctx, cID, prevDelta)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		if len(du.URL) > 0 {
			deltaURLs[cID] = du.URL
		} else if !du.Reset {
			logger.Ctx(ictx).Info("missing delta url")
		}

		currPath, err := bh.canonicalPath(path.Builder{}.Append(cID), qp.TenantID)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		// Remove any deleted IDs from the set of added IDs because items that are
		// deleted and then restored will have a different ID than they did
		// originally.
		for remove := range removed {
			delete(added, remove)
		}

		edc := NewCollection(
			bh,
			qp.ProtectedResource.ID(),
			currPath,
			prevPath,
			path.Builder{}.Append(cName),
			qp.Category,
			added,
			removed,
			statusUpdater,
			ctrlOpts,
			du.Reset)

		collections[cID] = &edc

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
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
			el.AddRecoverable(ctx, clues.Wrap(err, "conflict: tombstone exists for a live collection").WithClues(ictx))
			continue
		}

		// only occurs if it was a new folder that we picked up during the container
		// resolver phase that got deleted in flight by the time we hit this stage.
		if len(p) == 0 {
			continue
		}

		prevPath, err := pathFromPrevString(p)
		if err != nil {
			// technically shouldn't ever happen.  But just in case...
			logger.CtxErr(ictx, err).Error("parsing tombstone prev path")
			continue
		}

		edc := NewCollection(
			bh,
			qp.ProtectedResource.ID(),
			nil, // marks the collection as deleted
			prevPath,
			nil, // tombstones don't need a location
			qp.Category,
			nil, // no items added
			nil, // this deletes a directory, so no items deleted either
			statusUpdater,
			ctrlOpts,
			false)

		collections[id] = &edc
	}

	logger.Ctx(ctx).Infow(
		"adding metadata collection entries",
		"num_deltas_entries", len(deltaURLs),
		"num_paths_entries", len(collections))

	pathPrefix, err := path.Builder{}.ToServiceCategoryMetadataPath(
		qp.TenantID,
		qp.ProtectedResource.ID(),
		path.GroupsService,
		qp.Category,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path prefix")
	}

	col, err := graph.MakeMetadataCollection(
		pathPrefix,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(metadata.PreviousPathFileName, currPaths),
			graph.NewMetadataEntry(metadata.DeltaURLsFileName, deltaURLs),
		},
		statusUpdater)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata collection")
	}

	collections["metadata"] = col

	return collections, el.Failure()
}
