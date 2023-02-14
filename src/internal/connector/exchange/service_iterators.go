package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type addedAndRemovedItemIDsGetter interface {
	GetAddedAndRemovedItemIDs(
		ctx context.Context,
		user, containerID, oldDeltaToken string,
	) ([]string, []string, api.DeltaUpdate, error)
}

// filterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a BackupCollection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func filterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	getter addedAndRemovedItemIDsGetter,
	collections map[string]data.BackupCollection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.ExchangeScope,
	dps DeltaPaths,
	ctrlOpts control.Options,
	errs *fault.Errors,
) error {
	var (
		// folder ID -> delta url or folder path lookups
		deltaURLs = map[string]string{}
		currPaths = map[string]string{}
		// copy of previousPaths.  any folder found in the resolver get
		// deleted from this map, leaving only the deleted folders behind
		tombstones = makeTombstones(dps)
	)

	logger.Ctx(ctx).Infow(
		"filling collections",
		"metadata_count",
		len(dps))

	// TODO(rkeepers): this should be passed in from the caller, probably
	// as an interface that satisfies the NewCollection requirements.
	// But this will work for the short term.
	ac, err := api.NewClient(qp.Credentials)
	if err != nil {
		return err
	}

	ibt, err := itemerByType(ac, scope.Category().PathType())
	if err != nil {
		return err
	}

	et := errs.Tracker()

	for _, c := range resolver.Items() {
		if errs.Err() != nil {
			return et.Err()
		}

		cID := *c.GetId()
		delete(tombstones, cID)

		currPath, locPath, ok := includeContainer(qp, c, scope)
		// Only create a collection if the path matches the scope.
		if !ok {
			continue
		}

		var (
			dp          = dps[cID]
			prevDelta   = dp.delta
			prevPathStr = dp.path
			prevPath    path.Path
		)

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				logger.Ctx(ctx).With("err", err).Errorw("parsing prev path", clues.InErr(err).Slice()...)
				// if the previous path is unusable, then the delta must be, too.
				prevDelta = ""
			}
		}

		added, removed, newDelta, err := getter.GetAddedAndRemovedItemIDs(ctx, qp.ResourceOwner, cID, prevDelta)
		if err != nil {
			if !graph.IsErrDeletedInFlight(err) {
				et.Add(err)
				continue
			}

			// race conditions happen, containers might get deleted while
			// this process is in flight.  If that happens, force the collection
			// to reset. This prevents any old items from being retained in
			// storage.  If the container (or its children) are sill missing
			// on the next backup, they'll get tombstoned.
			newDelta = api.DeltaUpdate{Reset: true}
		}

		if len(newDelta.URL) > 0 {
			deltaURLs[cID] = newDelta.URL
		}

		if qp.Category != path.EventsCategory {
			locPath = nil
		}

		edc := NewCollection(
			qp.ResourceOwner,
			currPath,
			prevPath,
			locPath,
			scope.Category().PathType(),
			ibt,
			statusUpdater,
			ctrlOpts,
			newDelta.Reset)

		collections[cID] = &edc

		for _, add := range added {
			edc.added[add] = struct{}{}
		}

		// Remove any deleted IDs from the set of added IDs because items that are
		// deleted and then restored will have a different ID than they did
		// originally.
		for _, remove := range removed {
			delete(edc.added, remove)
			edc.removed[remove] = struct{}{}
		}

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
	}

	// A tombstone is a folder that needs to be marked for deletion.
	// The only situation where a tombstone should appear is if the folder exists
	// in the `previousPath` set, but does not exist in the current container
	// resolver (which contains all the resource owners' current containers).
	for id, p := range tombstones {
		if errs.Err() != nil {
			return et.Err()
		}

		if collections[id] != nil {
			et.Add(clues.Wrap(err, "conflict: tombstone exists for a live collection").WithClues(ctx))
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
			logger.Ctx(ctx).With("err", err).Errorw("parsing tombstone prev path", clues.InErr(err).Slice()...)
			continue
		}

		edc := NewCollection(
			qp.ResourceOwner,
			nil, // marks the collection as deleted
			prevPath,
			nil, // tombstones don't need a location
			scope.Category().PathType(),
			ibt,
			statusUpdater,
			ctrlOpts,
			false)
		collections[id] = &edc
	}

	entries := []graph.MetadataCollectionEntry{
		graph.NewMetadataEntry(graph.PreviousPathFileName, currPaths),
	}

	logger.Ctx(ctx).Infow(
		"adding metadata collection entries",
		"num_paths_entries", len(currPaths),
		"num_deltas_entries", len(deltaURLs))

	if len(deltaURLs) > 0 {
		entries = append(entries, graph.NewMetadataEntry(graph.DeltaURLsFileName, deltaURLs))
	}

	col, err := graph.MakeMetadataCollection(
		qp.Credentials.AzureTenantID,
		qp.ResourceOwner,
		path.ExchangeService,
		qp.Category,
		entries,
		statusUpdater)
	if err != nil {
		return clues.Wrap(err, "making metadata collection")
	}

	collections["metadata"] = col

	return et.Err()
}

// produces a set of id:path pairs from the deltapaths map.
// Each entry in the set will, if not removed, produce a collection
// that will delete the tombstone by path.
func makeTombstones(dps DeltaPaths) map[string]string {
	r := make(map[string]string, len(dps))

	for id, v := range dps {
		r[id] = v.path
	}

	return r
}

func pathFromPrevString(ps string) (path.Path, error) {
	p, err := path.FromDataLayerPath(ps, false)
	if err != nil {
		return nil, errors.Wrap(err, "parsing previous path string")
	}

	return p, nil
}

func itemerByType(ac api.Client, category path.CategoryType) (itemer, error) {
	switch category {
	case path.EmailCategory:
		return ac.Mail(), nil
	case path.EventsCategory:
		return ac.Events(), nil
	case path.ContactsCategory:
		return ac.Contacts(), nil
	default:
		return nil, clues.New("category not registered in getFetchIDFunc")
	}
}
