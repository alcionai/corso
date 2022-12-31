package exchange

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
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
// into a Collection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func filterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	getter addedAndRemovedItemIDsGetter,
	collections map[string]data.Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.ExchangeScope,
	dps DeltaPaths,
	ctrlOpts control.Options,
) error {
	var (
		errs error
		// folder ID -> delta url or folder path lookups
		deltaURLs = map[string]string{}
		currPaths = map[string]string{}
		// copy of previousPaths.  any folder found in the resolver get
		// deleted from this map, leaving only the deleted folders behind
		tombstones = makeTombstones(dps)
	)

	// TODO(rkeepers): this should be passed in from the caller, probably
	// as an interface that satisfies the NewCollection requirements.
	// But this will work for the short term.
	ac, err := api.NewClient(qp.Credentials)
	if err != nil {
		return err
	}

	for _, c := range resolver.Items() {
		if ctrlOpts.FailFast && errs != nil {
			return errs
		}

		// cannot be moved out of the loop,
		// else we run into state issues.
		service, err := createService(qp.Credentials)
		if err != nil {
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			continue
		}

		cID := *c.GetId()
		delete(tombstones, cID)

		currPath, ok := includeContainer(qp, c, scope)
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
				logger.Ctx(ctx).Error(err)
				// if the previous path is unusable, then the delta must be, too.
				prevDelta = ""
			}
		}

		added, removed, newDelta, err := getter.GetAddedAndRemovedItemIDs(ctx, qp.ResourceOwner, cID, prevDelta)
		if err != nil {
			if graph.IsErrDeletedInFlight(err) == nil {
				errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			} else {
				// race conditions happen, containers might get deleted while
				// this process is in flight.  If that happens, force the collection
				// to reset which will prevent any old items from being retained in
				// storage.  If the container (or its children) are sill missing
				// on the next backup, they'll get tombstoned.
				newDelta = api.DeltaUpdate{Reset: true}
			}

			continue
		}

		if len(newDelta.URL) > 0 {
			deltaURLs[cID] = newDelta.URL
		}

		edc := NewCollection(
			qp.ResourceOwner,
			currPath,
			prevPath,
			scope.Category().PathType(),
			ac,
			service,
			statusUpdater,
			ctrlOpts,
			newDelta.Reset)

		collections[cID] = &edc
		edc.added = append(edc.added, added...)
		edc.removed = append(edc.removed, removed...)

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
	}

	// A tombstone is a folder that needs to be marked for deletion.
	// The only situation where a tombstone should appear is if the folder exists
	// in the `previousPath` set, but does not exist in the current container
	// resolver (which contains all the resource owners' current containers).
	for id, p := range tombstones {
		service, err := createService(qp.Credentials)
		if err != nil {
			errs = support.WrapAndAppend(p, err, errs)
			continue
		}

		if collections[id] != nil {
			errs = support.WrapAndAppend(p, errors.New("conflict: tombstone exists for a live collection"), errs)
			continue
		}

		// only occurs if it was a new folder that we picked up during the container
		// resolver phase that got deleted in flight by the time we hit this stage.
		if len(p) == 0 {
			continue
		}

		prevPath, err := pathFromPrevString(p)
		if err != nil {
			// technically shouldn't ever happen.  But just in case, we need to catch
			// it for protection.
			logger.Ctx(ctx).Errorw("parsing tombstone path", "err", err)
			continue
		}

		edc := NewCollection(
			qp.ResourceOwner,
			nil, // marks the collection as deleted
			prevPath,
			scope.Category().PathType(),
			ac,
			service,
			statusUpdater,
			ctrlOpts,
			false)
		collections[id] = &edc
	}

	entries := []graph.MetadataCollectionEntry{
		graph.NewMetadataEntry(graph.PreviousPathFileName, currPaths),
	}

	if len(deltaURLs) > 0 {
		entries = append(entries, graph.NewMetadataEntry(graph.DeltaURLsFileName, deltaURLs))
	}

	if col, err := graph.MakeMetadataCollection(
		qp.Credentials.AzureTenantID,
		qp.ResourceOwner,
		path.ExchangeService,
		qp.Category,
		entries,
		statusUpdater,
	); err != nil {
		errs = support.WrapAndAppend("making metadata collection", err, errs)
	} else if col != nil {
		collections["metadata"] = col
	}

	return errs
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
