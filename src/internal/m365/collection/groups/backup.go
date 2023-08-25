package groups

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

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
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
	handler BackupHandler,
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

	// TODO(keepers): probably shouldn't call out channels here specifically.
	// This should be a generic container handler.  But we don't need
	// to worry about that until if/when we use this code to get email
	// conversations as well.
	// Also, this should be produced by the Handler.
	// chanPager := handler.NewChannelsPager(qp.ProtectedResource.ID())

	//  enumerating channels
	pager := handler.NewChannelsPager(qp.ProtectedResource.ID(), []string{})

	// Loop through all pages returned by Graph API.
	for {
		var (
			err  error
			page api.PageLinker
		)

		page, err = pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))
		if err != nil {
			return nil, graph.Wrap(ctx, err, "retrieving drives")
		}

		channels, err := pager.ValuesIn(page)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "extracting drives from response")
		}

		collections, err := populateCollections(
			ctx,
			qp,
			handler,
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

		nextLink := ptr.Val(page.GetOdataNextLink())
		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	return allCollections, nil
}

func populateCollections(
	ctx context.Context,
	qp graph.QueryParams,
	bh BackupHandler,
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
	// TODO(neha/keepers): figure out if deltas are stored per channel, or per group.
	// deltaURLs = map[string]string{}
	// currPaths = map[string]string{}
	// copy of previousPaths.  every channel present in the slice param
	// gets removed from this map; the remaining channels at the end of
	// the process have been deleted.
	// tombstones = makeTombstones(dps)

	logger.Ctx(ctx).Infow("filling collections")
	// , "len_deltapaths", len(dps))

	el := errs.Local()

	for _, c := range channels {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		cID := ptr.Val(c.GetId())
		// delete(tombstones, cID)

		var (
			err error
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

		// currPath, locPath
		// TODO(rkeepers): the handler should provide this functionality.
		// Only create a collection if the path matches the scope.
		if !includeContainer(ictx, qp, c, scope, qp.Category) {
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
		// TODO: Neha check this
		var fields []string
		// TODO: the handler should provide this implementation.
		// TODO: if we doing this messages are items for us.
		items, err := collectItems(
			ctx,
			bh.NewMessagePager(qp.ProtectedResource.ID(), ptr.Val(c.GetId()), fields))
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

		// TODO: retrieve from handler
		currPath, err := path.Builder{}.
			Append(ptr.Val(c.GetId())).
			ToDataLayerPath(
				qp.TenantID,
				qp.ProtectedResource.ID(),
				path.GroupsService,
				qp.Category,
				true)
		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		edc := NewCollection(
			qp.ProtectedResource.ID(),
			currPath,
			prevPath,
			path.Builder{}.Append(ptr.Val(c.GetDisplayName())),
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

		// FIXME: normally this goes before removal, but linters
		for _, item := range items {
			edc.added[ptr.Val(item.GetId())] = struct{}{}
		}
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

func collectItems(
	ctx context.Context,
	pager api.ChannelMessageDeltaEnumerator,
) ([]models.ChatMessageable, error) {
	items := []models.ChatMessageable{}

	for {
		// assume delta urls here, which allows single-token consumption
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))
		if err != nil {
			return nil, graph.Wrap(ctx, err, "getting page")
		}

		// if graph.IsErrInvalidDelta(err) {
		// 	logger.Ctx(ctx).Infow("Invalid previous delta link", "link", prevDelta)

		// 	invalidPrevDelta = true
		// 	newPaths = map[string]string{}

		// 	pager.Reset()

		// 	continue
		// }

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "getting items in page")
		}

		items = append(items, vals...)

		nextLink, _ := api.NextAndDeltaLink(page)

		// if len(deltaLink) > 0 {
		// 	newDeltaURL = deltaLink
		// }

		// Check if there are more items
		if len(nextLink) == 0 {
			break
		}

		logger.Ctx(ctx).Debugw("found nextLink", "next_link", nextLink)
		pager.SetNext(nextLink)
	}

	return items, nil
}

// Returns true if the container passes the scope comparison and should be included.
// Returns:
// - the path representing the directory as it should be stored in the repository.
// - the human-readable path using display names.
// - true if the path passes the scope comparison.
func includeContainer(
	ctx context.Context,
	qp graph.QueryParams,
	gd graph.Displayable,
	scope selectors.GroupsScope,
	category path.CategoryType,
) bool {
	// assume a single-level hierarchy
	directory := ptr.Val(gd.GetDisplayName())

	// TODO(keepers): awaiting parent branch to update to main
	ok := scope.Matches(selectors.GroupsChannelMessage, directory)

	logger.Ctx(ctx).With(
		"included", ok,
		"scope", scope,
		"match_target", directory,
	).Debug("backup folder selection filter")

	return ok
}
