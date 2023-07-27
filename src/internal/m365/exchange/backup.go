package exchange

import (
	"context"
	"encoding/json"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// MetadataFileNames produces the category-specific set of filenames used to
// store graph metadata such as delta tokens and folderID->path references.
func MetadataFileNames(cat path.CategoryType) []string {
	switch cat {
	case path.EmailCategory, path.ContactsCategory:
		return []string{graph.DeltaURLsFileName, graph.PreviousPathFileName}
	default:
		return []string{graph.PreviousPathFileName}
	}
}

type CatDeltaPaths map[path.CategoryType]DeltaPaths

type DeltaPaths map[string]DeltaPath

func (dps DeltaPaths) AddDelta(k, d string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.Delta = d
	dps[k] = dp
}

func (dps DeltaPaths) AddPath(k, p string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.Path = p
	dps[k] = dp
}

type DeltaPath struct {
	Delta string
	Path  string
}

// ParseMetadataCollections produces a map of structs holding delta
// and path lookup maps.
func parseMetadataCollections(
	ctx context.Context,
	colls []data.RestoreCollection,
) (CatDeltaPaths, bool, error) {
	// cdp stores metadata
	cdp := CatDeltaPaths{
		path.ContactsCategory: {},
		path.EmailCategory:    {},
		path.EventsCategory:   {},
	}

	// found tracks the metadata we've loaded, to make sure we don't
	// fetch overlapping copies.
	found := map[path.CategoryType]map[string]struct{}{
		path.ContactsCategory: {},
		path.EmailCategory:    {},
		path.EventsCategory:   {},
	}

	// errors from metadata items should not stop the backup,
	// but it should prevent us from using previous backups
	errs := fault.New(true)

	for _, coll := range colls {
		var (
			breakLoop bool
			items     = coll.Items(ctx, errs)
			category  = coll.FullPath().Category()
		)

		for {
			select {
			case <-ctx.Done():
				return nil, false, clues.Wrap(ctx.Err(), "parsing collection metadata").WithClues(ctx)

			case item, ok := <-items:
				if !ok || errs.Failure() != nil {
					breakLoop = true
					break
				}

				var (
					m    = map[string]string{}
					cdps = cdp[category]
				)

				err := json.NewDecoder(item.ToReader()).Decode(&m)
				if err != nil {
					return nil, false, clues.New("decoding metadata json").WithClues(ctx)
				}

				switch item.UUID() {
				case graph.PreviousPathFileName:
					if _, ok := found[category]["path"]; ok {
						return nil, false, clues.Wrap(clues.New(category.String()), "multiple versions of path metadata").WithClues(ctx)
					}

					for k, p := range m {
						cdps.AddPath(k, p)
					}

					found[category]["path"] = struct{}{}

				case graph.DeltaURLsFileName:
					if _, ok := found[category]["delta"]; ok {
						return nil, false, clues.Wrap(clues.New(category.String()), "multiple versions of delta metadata").WithClues(ctx)
					}

					for k, d := range m {
						cdps.AddDelta(k, d)
					}

					found[category]["delta"] = struct{}{}
				}

				cdp[category] = cdps
			}

			if breakLoop {
				break
			}
		}
	}

	if errs.Failure() != nil {
		logger.CtxErr(ctx, errs.Failure()).Info("reading metadata collection items")

		return CatDeltaPaths{
			path.ContactsCategory: {},
			path.EmailCategory:    {},
			path.EventsCategory:   {},
		}, false, nil
	}

	// Remove any entries that contain a path or a delta, but not both.
	// That metadata is considered incomplete, and needs to incur a
	// complete backup on the next run.
	for _, dps := range cdp {
		for k, dp := range dps {
			if len(dp.Path) == 0 {
				delete(dps, k)
			}
		}
	}

	return cdp, true, nil
}

// ProduceBackupCollections returns a DataCollection which the caller can
// use to read mailbox data out for the specified user
func ProduceBackupCollections(
	ctx context.Context,
	ac api.Client,
	selector selectors.Selector,
	tenantID string,
	user idname.Provider,
	metadata []data.RestoreCollection,
	su support.StatusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "exchange dataCollection selector").WithClues(ctx)
	}

	var (
		collections = []data.BackupCollection{}
		el          = errs.Local()
		categories  = map[path.CategoryType]struct{}{}
		handlers    = BackupHandlers(ac)
	)

	// Turn on concurrency limiter middleware for exchange backups
	// unless explicitly disabled through DisableConcurrencyLimiterFN cli flag
	graph.InitializeConcurrencyLimiter(
		ctx,
		ctrlOpts.ToggleFeatures.DisableConcurrencyLimiter,
		ctrlOpts.Parallelism.ItemFetch)

	cdps, canUsePreviousBackup, err := parseMetadataCollections(ctx, metadata)
	if err != nil {
		return nil, nil, false, err
	}

	ctx = clues.Add(ctx, "can_use_previous_backup", canUsePreviousBackup)

	for _, scope := range eb.Scopes() {
		if el.Failure() != nil {
			break
		}

		dcs, err := createCollections(
			ctx,
			handlers,
			tenantID,
			user,
			scope,
			cdps[scope.Category().PathType()],
			ctrlOpts,
			su,
			errs)
		if err != nil {
			el.AddRecoverable(ctx, err)
			continue
		}

		categories[scope.Category().PathType()] = struct{}{}

		collections = append(collections, dcs...)
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			tenantID,
			user.ID(),
			path.ExchangeService,
			categories,
			su,
			errs)
		if err != nil {
			return nil, nil, false, err
		}

		collections = append(collections, baseCols...)
	}

	return collections, nil, canUsePreviousBackup, el.Failure()
}

// createCollections - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.ExchangeScope
// determines the type of collections that are retrieved.
func createCollections(
	ctx context.Context,
	handlers map[path.CategoryType]backupHandler,
	tenantID string,
	user idname.Provider,
	scope selectors.ExchangeScope,
	dps DeltaPaths,
	ctrlOpts control.Options,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	ctx = clues.Add(ctx, "category", scope.Category().PathType())

	var (
		allCollections = make([]data.BackupCollection, 0)
		category       = scope.Category().PathType()
		qp             = graph.QueryParams{
			Category:      category,
			ResourceOwner: user,
			TenantID:      tenantID,
		}
	)

	handler, ok := handlers[category]
	if !ok {
		return nil, clues.New("unsupported backup category type").WithClues(ctx)
	}

	progressBar := observe.MessageWithCompletion(
		ctx,
		observe.Bulletf("%s", qp.Category))
	defer close(progressBar)

	rootFolder, cc := handler.NewContainerCache(user.ID())

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
		ctrlOpts,
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
	dps DeltaPaths,
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
		delete(tombstones, cID)

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

		currPath, locPath, ok := includeContainer(ictx, qp, c, scope, category)
		// Only create a collection if the path matches the scope.
		if !ok {
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

		added, removed, newDelta, err := bh.itemEnumerator().
			GetAddedAndRemovedItemIDs(
				ictx,
				qp.ResourceOwner.ID(),
				cID,
				prevDelta,
				ctrlOpts.ToggleFeatures.ExchangeImmutableIDs,
				!ctrlOpts.ToggleFeatures.DisableDelta)
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
			newDelta = api.DeltaUpdate{Reset: true}
		}

		if len(newDelta.URL) > 0 {
			deltaURLs[cID] = newDelta.URL
		} else if !newDelta.Reset {
			logger.Ctx(ictx).Info("missing delta url")
		}

		edc := NewCollection(
			qp.ResourceOwner.ID(),
			currPath,
			prevPath,
			locPath,
			category,
			bh.itemHandler(),
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
			qp.ResourceOwner.ID(),
			nil, // marks the collection as deleted
			prevPath,
			nil, // tombstones don't need a location
			category,
			bh.itemHandler(),
			statusUpdater,
			ctrlOpts,
			false)
		collections[id] = &edc
	}

	logger.Ctx(ctx).Infow(
		"adding metadata collection entries",
		"num_paths_entries", len(currPaths),
		"num_deltas_entries", len(deltaURLs))

	col, err := graph.MakeMetadataCollection(
		qp.TenantID,
		qp.ResourceOwner.ID(),
		path.ExchangeService,
		qp.Category,
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(graph.PreviousPathFileName, currPaths),
			graph.NewMetadataEntry(graph.DeltaURLsFileName, deltaURLs),
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
func makeTombstones(dps DeltaPaths) map[string]string {
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
	if category == path.ContactsCategory && ptr.Val(c.GetDisplayName()) == api.DefaultContacts {
		loc = loc.Append(api.DefaultContacts)
	}

	dirPath, err := pb.ToDataLayerExchangePathForCategory(
		qp.TenantID,
		qp.ResourceOwner.ID(),
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
			qp.ResourceOwner.ID(),
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
		"matches_input", directory,
	).Debug("backup folder selection filter")

	return dirPath, loc, ok
}
