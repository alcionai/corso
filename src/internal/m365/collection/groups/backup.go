package groups

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
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
	// dps DeltaPaths,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
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

	catProgress := observe.MessageWithCompletion(
		ctx,
		observe.Bulletf("%s", qp.Category))
	defer close(catProgress)

	channels, err := bh.getChannels(ctx)
	if err != nil {
		return nil, clues.Stack(err)
	}

	collections, err := populateCollections(
		ctx,
		qp,
		bh,
		su,
		channels,
		scope,
		// dps,
		bpc.Options,
		errs)
	if err != nil {
		return nil, clues.Wrap(err, "filling collections")
	}

	for _, coll := range collections {
		allCollections = append(allCollections, coll)
	}

	return allCollections, nil
}

func populateCollections(
	ctx context.Context,
	qp graph.QueryParams,
	bh backupHandler,
	statusUpdater support.StatusUpdater,
	channels []models.Channelable,
	scope selectors.GroupsScope,
	// dps DeltaPaths,
	ctrlOpts control.Options,
	errs *fault.Bus,
) (map[string]data.BackupCollection, error) {
	// channel ID -> BackupCollection.
	channelCollections := map[string]data.BackupCollection{}

	// channel ID -> delta url or folder path lookups
	// deltaURLs = map[string]string{}
	// currPaths = map[string]string{}
	// copy of previousPaths.  every channel present in the slice param
	// gets removed from this map; the remaining channels at the end of
	// the process have been deleted.
	// tombstones = makeTombstones(dps)

	logger.Ctx(ctx).Info("filling collections")
	// , "len_deltapaths", len(dps))

	el := errs.Local()

	for _, c := range channels {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		// delete(tombstones, cID)

		var (
			cID   = ptr.Val(c.GetId())
			cName = ptr.Val(c.GetDisplayName())
			err   error
			// dp          = dps[cID]
			// prevDelta   = dp.Delta
			// prevPathStr = dp.Path // do not log: pii; log prevPath instead
			// prevPath    path.Path
			ictx = clues.Add(
				ctx,
				"channel_id", cID)
			// "previous_delta", pii.SafeURL{
			// 	URL:           prevDelta,
			// 	SafePathElems: graph.SafeURLPathParams,
			// 	SafeQueryKeys: graph.SafeURLQueryParams,
			// })
		)

		// Only create a collection if the path matches the scope.
		if !bh.includeContainer(ictx, qp, c, scope) {
			continue
		}

		// if len(prevPathStr) > 0 {
		// 	if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
		// 		logger.CtxErr(ictx, err).Error("parsing prev path")
		// 		// if the previous path is unusable, then the delta must be, too.
		// 		prevDelta = ""
		// 	}
		// }

		// ictx = clues.Add(ictx, "previous_path", prevPath)

		items, _, err := bh.getChannelMessageIDsDelta(ctx, cID, "")
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		// if len(newDelta.URL) > 0 {
		// 	deltaURLs[cID] = newDelta.URL
		// } else if !newDelta.Reset {
		// 	logger.Ctx(ictx).Info("missing delta url")
		// }

		var prevPath path.Path

		currPath, err := bh.canonicalPath(path.Builder{}.Append(cID), qp.TenantID)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		edc := NewCollection(
			bh,
			qp.ProtectedResource.ID(),
			currPath,
			prevPath,
			path.Builder{}.Append(cName),
			qp.Category,
			statusUpdater,
			ctrlOpts)

		channelCollections[cID] = &edc

		// TODO: handle deleted items for v1 backup.
		// // Remove any deleted IDs from the set of added IDs because items that are
		// // deleted and then restored will have a different ID than they did
		// // originally.
		// for _, remove := range removed {
		// 	delete(edc.added, remove)
		// 	edc.removed[remove] = struct{}{}
		// }

		// // add the current path for the container ID to be used in the next backup
		// // as the "previous path", for reference in case of a rename or relocation.
		// currPaths[cID] = currPath.String()

		// FIXME: normally this goes before removal, but the linters require no bottom comments
		maps.Copy(edc.added, items)
	}

	// TODO: handle tombstones here

	logger.Ctx(ctx).Infow(
		"adding metadata collection entries",
		// "num_deltas_entries", len(deltaURLs),
		"num_paths_entries", len(channelCollections))

	// col, err := graph.MakeMetadataCollection(
	// 	qp.TenantID,
	// 	qp.ProtectedResource.ID(),
	// 	path.ExchangeService,
	// 	qp.Category,
	// 	[]graph.MetadataCollectionEntry{
	// 		graph.NewMetadataEntry(graph.PreviousPathFileName, currPaths),
	// 		graph.NewMetadataEntry(graph.DeltaURLsFileName, deltaURLs),
	// 	},
	// 	statusUpdater)
	// if err != nil {
	// 	return nil, clues.Wrap(err, "making metadata collection")
	// }

	// channelCollections["metadata"] = col

	return channelCollections, el.Failure()
}
