package exchange

import (
	"context"

	"github.com/alcionai/clues"

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
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

func CreateCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	handlers map[path.CategoryType]backupHandler,
	tenantID string,
	scope selectors.ExchangeScope,
	dps metadata.DeltaPaths,
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

	handler, ok := handlers[category]
	if !ok {
		return nil, clues.New("unsupported backup category type").WithClues(ctx)
	}

	foldersComplete := observe.MessageWithCompletion(
		ctx,
		observe.Bulletf("%s", qp.Category.HumanString()))
	defer close(foldersComplete)

	rootFolder, cc := handler.NewContainerCache(bpc.ProtectedResource.ID())

	if err := cc.Populate(ctx, errs, rootFolder); err != nil {
		return nil, clues.Wrap(err, "populating container cache")
	}

	collections, err := populateCollections(
		ctx,
		qp,
		handler,
		su,
		cc,
		scope,
		dps,
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

// populateCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a BackupCollection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
//
// TODO(ashmrtn): This should really return []data.BackupCollection but
// unfortunately some of our tests rely on being able to lookup returned
// collections by ID and it would be non-trivial to change them.
func populateCollections(
	ctx context.Context,
	qp graph.QueryParams,
	bh backupHandler,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.ExchangeScope,
	dps metadata.DeltaPaths,
	ctrlOpts control.Options,
	errs *fault.Bus,
) (map[string]data.BackupCollection, error) {
	var (
		// folder ID -> BackupCollection.
		collections = map[string]data.BackupCollection{}
		// folder ID -> delta url or folder path lookups
		deltaURLs = map[string]string{}
		currPaths = map[string]string{}
		// copy of previousPaths.  any folder found in the resolver get
		// deleted from this map, leaving only the deleted folders behind
		tombstones = makeTombstones(dps)
		category   = qp.Category
	)

	logger.Ctx(ctx).Infow("filling collections", "len_deltapaths", len(dps))

	el := errs.Local()

	for _, c := range resolver.Items() {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		cID := ptr.Val(c.GetId())

		var (
			err         error
			dp          = dps[cID]
			prevDelta   = dp.Delta
			prevPathStr = dp.Path // do not log: pii; log prevPath instead
			prevPath    path.Path
			ictx        = clues.Add(
				ctx,
				"container_id", cID,
				"previous_delta", pii.SafeURL{
					URL:           prevDelta,
					SafePathElems: graph.SafeURLPathParams,
					SafeQueryKeys: graph.SafeURLQueryParams,
				})
		)

		// Only create a collection if the path matches the scope.
		currPath, locPath, ok := includeContainer(ictx, qp, c, scope, category)
		if !ok {
			continue
		}

		delete(tombstones, cID)

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				logger.CtxErr(ictx, err).Error("parsing prev path")
				// if the previous path is unusable, then the delta must be, too.
				prevDelta = ""
			}
		}

		ictx = clues.Add(ictx, "previous_path", prevPath)

		cc := api.CallConfig{
			CanMakeDeltaQueries: !ctrlOpts.ToggleFeatures.DisableDelta,
			UseImmutableIDs:     ctrlOpts.ToggleFeatures.ExchangeImmutableIDs,
		}

		addAndRem, err := bh.itemEnumerator().
			GetAddedAndRemovedItemIDs(
				ictx,
				qp.ProtectedResource.ID(),
				cID,
				prevDelta,
				cc)
		if err != nil {
			if !graph.IsErrDeletedInFlight(err) {
				el.AddRecoverable(ctx, clues.Stack(err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			// race conditions happen, containers might get deleted while
			// this process is in flight.  If that happens, force the collection
			// to reset. This prevents any old items from being retained in
			// storage.  If the container (or its children) are sill missing
			// on the next backup, they'll get tombstoned.
			addAndRem.DU = pagers.DeltaUpdate{Reset: true}
		}

		if len(addAndRem.DU.URL) > 0 {
			deltaURLs[cID] = addAndRem.DU.URL
		} else if !addAndRem.DU.Reset {
			logger.Ctx(ictx).Info("missing delta url")
		}

		edc := NewCollection(
			data.NewBaseCollection(
				currPath,
				prevPath,
				locPath,
				ctrlOpts,
				addAndRem.DU.Reset),
			qp.ProtectedResource.ID(),
			bh.itemHandler(),
			addAndRem.Added,
			addAndRem.Removed,
			// TODO: produce a feature flag that allows selective
			// enabling of valid modTimes.  This currently produces
			// rare-case failures with incorrect details merging.
			// Root cause is not yet known.
			false,
			statusUpdater)

		collections[cID] = edc

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
	}

	// A tombstone is a folder that needs to be marked for deletion.
	// The only situation where a tombstone should appear is if the folder exists
	// in the `previousPath` set, but does not exist in the current container
	// resolver (which contains all the resource owners' current containers).
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

		collections[id] = data.NewTombstoneCollection(prevPath, ctrlOpts)
	}

	logger.Ctx(ctx).Infow(
		"adding metadata collection entries",
		"num_paths_entries", len(currPaths),
		"num_deltas_entries", len(deltaURLs))

	pathPrefix, err := path.BuildMetadata(
		qp.TenantID,
		qp.ProtectedResource.ID(),
		path.ExchangeService,
		qp.Category,
		false)
	if err != nil {
		return nil, clues.Wrap(err, "making metadata path")
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

// produces a set of id:path pairs from the deltapaths map.
// Each entry in the set will, if not removed, produce a collection
// that will delete the tombstone by path.
func makeTombstones(dps metadata.DeltaPaths) map[string]string {
	r := make(map[string]string, len(dps))

	for id, v := range dps {
		r[id] = v.Path
	}

	return r
}

func pathFromPrevString(ps string) (path.Path, error) {
	p, err := path.FromDataLayerPath(ps, false)
	if err != nil {
		return nil, clues.Wrap(err, "parsing previous path string")
	}

	return p, nil
}

// Returns true if the container passes the scope comparison and should be included.
// Returns:
// - the path representing the directory as it should be stored in the repository.
// - the human-readable path using display names.
// - true if the path passes the scope comparison.
func includeContainer(
	ctx context.Context,
	qp graph.QueryParams,
	c graph.CachedContainer,
	scope selectors.ExchangeScope,
	category path.CategoryType,
) (path.Path, *path.Builder, bool) {
	var (
		directory string
		locPath   path.Path
		pb        = c.Path()
		loc       = c.Location()
	)

	// Clause ensures that DefaultContactFolder is inspected properly
	if category == path.ContactsCategory && len(loc.Elements()) == 0 {
		loc = loc.Append(ptr.Val(c.GetDisplayName()))
	}

	dirPath, err := pb.ToDataLayerExchangePathForCategory(
		qp.TenantID,
		qp.ProtectedResource.ID(),
		category,
		false)
	// Containers without a path (e.g. Root mail folder) always err here.
	if err != nil {
		return nil, nil, false
	}

	directory = dirPath.Folder(false)

	if loc != nil {
		locPath, err = loc.ToDataLayerExchangePathForCategory(
			qp.TenantID,
			qp.ProtectedResource.ID(),
			category,
			false)
		// Containers without a path (e.g. Root mail folder) always err here.
		if err != nil {
			return nil, nil, false
		}

		directory = locPath.Folder(false)
	}

	var ok bool

	switch category {
	case path.EmailCategory:
		ok = scope.Matches(selectors.ExchangeMailFolder, directory)
	case path.ContactsCategory:
		ok = scope.Matches(selectors.ExchangeContactFolder, directory)
	case path.EventsCategory:
		ok = scope.Matches(selectors.ExchangeEventCalendar, directory)
	default:
		return nil, nil, false
	}

	logger.Ctx(ctx).With(
		"included", ok,
		"scope", scope,
		"matches_input", directory).Debug("backup folder selection filter")

	return dirPath, loc, ok
}
