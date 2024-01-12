package exchange

import (
	"context"
	"errors"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

const (
	defaultPreviewMaxContainers        = 5
	defaultPreviewMaxItemsPerContainer = 10
	defaultPreviewMaxItems             = defaultPreviewMaxContainers * defaultPreviewMaxItemsPerContainer
)

func CreateCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	handlers map[path.CategoryType]backupHandler,
	tenantID string,
	scope selectors.ExchangeScope,
	dps metadata.DeltaPaths,
	su support.StatusUpdater,
	counter *count.Bus,
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
		collections map[string]data.BackupCollection
		err         error
	)

	handler, ok := handlers[category]
	if !ok {
		return nil, clues.NewWC(ctx, "unsupported backup category type")
	}

	progressMessage := observe.MessageWithCompletion(
		ctx,
		observe.ProgressCfg{
			Indent:            1,
			CompletionMessage: func() string { return fmt.Sprintf("(found %d folders)", len(collections)) },
		},
		qp.Category.HumanString())
	defer close(progressMessage)

	rootFolder, cc := handler.NewContainerCache(bpc.ProtectedResource.ID())

	if err := cc.Populate(ctx, errs, rootFolder); err != nil {
		return nil, clues.Wrap(err, "populating container cache")
	}

	collections, err = populateCollections(
		ctx,
		qp,
		handler,
		su,
		cc,
		scope,
		dps,
		bpc.Options,
		counter,
		errs)
	if err != nil {
		return nil, clues.Wrap(err, "filling collections")
	}

	counter.Add(count.Collections, int64(len(collections)))

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
	counter *count.Bus,
	errs *fault.Bus,
) (map[string]data.BackupCollection, error) {
	var (
		err error
		// folder ID -> BackupCollection.
		collections = map[string]data.BackupCollection{}
		// folder ID -> delta url or folder path lookups
		deltaURLs = map[string]string{}
		currPaths = map[string]string{}
		// copy of previousPaths.  any folder found in the resolver get
		// deleted from this map, leaving only the deleted folders behind
		tombstones = makeTombstones(dps)
		category   = qp.Category

		// Limits and counters below are currently only used for preview backups
		// since they only act on a subset of items. Make a copy of the passed in
		// limits so we can log both the passed in options and what they were set to
		// if we used default values for some things.
		effectiveLimits = ctrlOpts.PreviewLimits

		addedItems      int
		addedContainers int
	)

	el := errs.Local()

	// Preview backups select a reduced set of data. This is managed by ordering
	// the set of results from the container resolver and reducing the number of
	// items selected from each container.
	if effectiveLimits.Enabled {
		resolver, err = newRankedContainerResolver(
			ctx,
			resolver,
			bh.folderGetter(),
			qp.ProtectedResource.ID(),
			// TODO(ashmrtn): Includes and excludes should really be associated with
			// the service not the data category. This is because a single data
			// handler may be used for multiple services (e.x. drive handler is used
			// for OneDrive, SharePoint, and Groups/Teams).
			bh.previewIncludeContainers(),
			bh.previewExcludeContainers())
		if err != nil {
			return nil, clues.Wrap(err, "creating ranked container resolver")
		}

		// Configure limits with reasonable defaults if they're not set.
		if effectiveLimits.MaxContainers == 0 {
			effectiveLimits.MaxContainers = defaultPreviewMaxContainers
		}

		if effectiveLimits.MaxItemsPerContainer == 0 {
			effectiveLimits.MaxItemsPerContainer = defaultPreviewMaxItemsPerContainer
		}

		if effectiveLimits.MaxItems == 0 {
			effectiveLimits.MaxItems = defaultPreviewMaxItems
		}
	}

	logger.Ctx(ctx).Infow(
		"filling collections",
		"len_deltapaths", len(dps),
		"limits", ctrlOpts.PreviewLimits,
		"effective_limits", effectiveLimits)
	counter.Add(count.PrevDeltas, int64(len(dps)))

	for _, c := range resolver.Items() {
		if el.Failure() != nil {
			return nil, el.Failure()
		}

		var (
			err        error
			cl         = counter.Local()
			itemConfig = api.CallConfig{
				CanMakeDeltaQueries: !ctrlOpts.ToggleFeatures.DisableDelta,
			}
			cID         = ptr.Val(c.GetId())
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

		ictx = clues.AddLabelCounter(ictx, cl.PlainAdder())

		// Only create a collection if the path matches the scope.
		currPath, locPath, ok := includeContainer(ictx, qp, c, scope, category)
		if !ok {
			cl.Inc(count.SkippedContainers)
			continue
		}

		ictx = clues.Add(
			ictx,
			"current_path", currPath,
			"current_location", locPath)

		delete(tombstones, cID)

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				err = clues.Stack(err).Label(count.BadPrevPath)
				logger.CtxErr(ictx, err).Error("parsing prev path")
				// if the previous path is unusable, then the delta must be, too.
				prevDelta = ""
			}
		}

		ictx = clues.Add(ictx, "previous_path", prevPath)

		// Since part of this is about figuring out how many items to get for this
		// particular container we need to reconfigure for every container we see.
		if effectiveLimits.Enabled {
			toAdd := effectiveLimits.MaxItems - addedItems

			if addedContainers >= effectiveLimits.MaxContainers || toAdd <= 0 {
				break
			}

			if toAdd > effectiveLimits.MaxItemsPerContainer {
				toAdd = effectiveLimits.MaxItemsPerContainer
			}

			// Delta tokens generated with this CallConfig shouldn't be used for
			// regular backups. They may have different query parameters which will
			// cause incorrect output for regular backups.
			itemConfig.LimitResults = toAdd
		}

		addAndRem, err := bh.itemEnumerator().
			GetAddedAndRemovedItemIDs(
				ictx,
				qp.ProtectedResource.ID(),
				cID,
				prevDelta,
				itemConfig)
		if err != nil {
			if !errors.Is(err, core.ErrNotFound) {
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
			cl.Inc(count.MissingDelta)
		}

		edc := NewCollection(
			data.NewBaseCollection(
				currPath,
				prevPath,
				locPath,
				ctrlOpts,
				addAndRem.DU.Reset,
				cl),
			qp.ProtectedResource.ID(),
			bh.itemHandler(),
			addAndRem.Added,
			addAndRem.Removed,
			// TODO: produce a feature flag that allows selective
			// enabling of valid modTimes.  This currently produces
			// rare failures with incorrect details merging.
			// Root cause is not yet known.
			false,
			statusUpdater,
			cl)

		collections[cID] = edc

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
		addedItems += len(addAndRem.Added)
		addedContainers++
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
			err := clues.WrapWC(ictx, err, "conflict: tombstone exists for a live collection").
				Label(count.CollectionTombstoneConflict)
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
			err = clues.StackWC(ctx, err).Label(count.BadPrevPath)
			// technically shouldn't ever happen.  But just in case...
			logger.CtxErr(ictx, err).Error("parsing tombstone prev path")

			continue
		}

		collections[id] = data.NewTombstoneCollection(prevPath, ctrlOpts, counter)
	}

	counter.Add(count.NewDeltas, int64(len(deltaURLs)))
	counter.Add(count.NewPrevPaths, int64(len(currPaths)))

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
		statusUpdater,
		count.New())
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
