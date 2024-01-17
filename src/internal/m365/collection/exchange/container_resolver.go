package exchange

import (
	"context"
	"errors"

	"github.com/alcionai/clues"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// common interfaces
// ---------------------------------------------------------------------------

type containerGetter interface {
	GetContainerByID(
		ctx context.Context,
		userID, dirID string,
	) (graph.Container, error)
}

type containersEnumerator[T any] interface {
	EnumerateContainers(
		ctx context.Context,
		userID, baseDirID string,
	) ([]T, error)
}

type containerRefresher interface {
	refreshContainer(
		ctx context.Context,
		dirID string,
	) (graph.CachedContainer, error)
}

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

// Exchange has a limit of 300 for folder depth. OneDrive has a limit on path
// length of 400 characters (including separators) which would be roughly 200
// folders if each folder is only a single character.
const maxIterations = 300

func newContainerResolver(refresher containerRefresher) *containerResolver {
	return &containerResolver{
		cache:     map[string]graph.CachedContainer{},
		refresher: refresher,
	}
}

type containerResolver struct {
	cache     map[string]graph.CachedContainer
	refresher containerRefresher
}

func (cr *containerResolver) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, *path.Builder, error) {
	ctx = clues.Add(ctx, "container_id", folderID)

	c, ok := cr.cache[folderID]
	if !ok {
		return nil, nil, clues.NewWC(ctx, "container not cached")
	}

	p := c.Path()
	if p == nil {
		return nil, nil, clues.NewWC(ctx, "cached container has no path")
	}

	return p, c.Location(), nil
}

// refreshContainer attempts to fetch the container with the given ID from Graph
// API. Returns a graph.CachedContainer if the container was found. If the
// container was deleted, returns nil, true, nil to note the container should
// be removed from the cache.
func (cr *containerResolver) refreshContainer(
	ctx context.Context,
	id string,
) (graph.CachedContainer, bool, error) {
	ctx = clues.Add(ctx, "refresh_container_id", id)
	logger.Ctx(ctx).Debug("refreshing container")

	if cr.refresher == nil {
		return nil, false, clues.NewWC(ctx, "nil refresher")
	}

	c, err := cr.refresher.refreshContainer(ctx, id)
	if err != nil && errors.Is(err, core.ErrNotFound) {
		logger.Ctx(ctx).Debug("container deleted")
		return nil, true, nil
	} else if err != nil {
		// This is some other error, just return it.
		return nil, false, clues.Wrap(err, "refreshing container")
	}

	return c, false, nil
}

// recoverContainer attempts to fetch a missing container from Graph API and
// populate the path for it. It returns
//   - the ID path for the folder
//   - the display name path for the folder
//   - if the folder was deleted
//   - any error that occurred
//
// If the folder is marked as deleted, child folders of this folder should be
// deleted if they haven't been moved to another folder.
func (cr *containerResolver) recoverContainer(
	ctx context.Context,
	folderID string,
	depth int,
) (*path.Builder, *path.Builder, bool, error) {
	c, deleted, err := cr.refreshContainer(ctx, folderID)
	if err != nil {
		return nil, nil, false, clues.Wrap(err, "fetching uncached container")
	}

	if deleted {
		logger.Ctx(ctx).Debug("fetching uncached container showed it was deleted")
		return nil, nil, deleted, err
	}

	if err := cr.addFolder(c); err != nil {
		return nil, nil, false, clues.WrapWC(ctx, err, "adding new container")
	}

	// Retry populating this container's paths.
	//
	// TODO(ashmrtn): May want to bump the depth here just so we don't get stuck
	// retrying too much if for some reason things keep moving around?
	resolved, err := cr.idToPath(ctx, folderID, depth)
	if err != nil {
		err = clues.Wrap(err, "repopulating uncached container")
	}

	return resolved.idPath, resolved.locPath, resolved.deleted, err
}

type resolvedPath struct {
	idPath  *path.Builder
	locPath *path.Builder
	cached  bool
	deleted bool
}

func (cr *containerResolver) idToPath(
	ctx context.Context,
	folderID string,
	depth int,
) (resolvedPath, error) {
	ctx = clues.Add(ctx, "container_id", folderID)

	if depth >= maxIterations {
		return resolvedPath{
				idPath:  nil,
				locPath: nil,
				cached:  false,
				deleted: false,
			},
			clues.NewWC(ctx, "path contains cycle or is too tall")
	}

	c, ok := cr.cache[folderID]
	if !ok {
		pth, loc, deleted, err := cr.recoverContainer(ctx, folderID, depth)
		if err != nil {
			err = clues.Stack(err)
		}

		return resolvedPath{
			idPath:  pth,
			locPath: loc,
			cached:  false,
			deleted: deleted,
		}, err
	}

	p := c.Path()
	if p != nil {
		return resolvedPath{
			idPath:  p,
			locPath: c.Location(),
			cached:  true,
			deleted: false,
		}, nil
	}

	resolved, err := cr.idToPath(
		ctx,
		ptr.Val(c.GetParentFolderId()),
		depth+1)
	if err != nil {
		return resolvedPath{
			idPath:  nil,
			locPath: nil,
			cached:  true,
			deleted: false,
		}, clues.Wrap(err, "retrieving parent container")
	}

	if !resolved.cached {
		logger.Ctx(ctx).Debug("parent container was refreshed")

		newContainer, shouldDelete, err := cr.refreshContainer(ctx, folderID)
		if err != nil {
			return resolvedPath{
				idPath:  nil,
				locPath: nil,
				cached:  true,
				deleted: false,
			}, clues.Wrap(err, "refreshing container")
		}

		if shouldDelete {
			logger.Ctx(ctx).Debug("refreshing container showed it was deleted")
			delete(cr.cache, folderID)

			return resolvedPath{
				idPath:  nil,
				locPath: nil,
				cached:  true,
				deleted: true,
			}, nil
		}

		// See if the newer version of the current container we got back has
		// changed. If it has then it could be that the container was moved prior to
		// deleting the parent and we just hit some eventual consistency case in
		// Graph.
		//
		// TODO(ashmrtn): May want to bump the depth here just so we don't get stuck
		// retrying too much if for some reason things keep moving around?
		if ptr.Val(newContainer.GetParentFolderId()) != ptr.Val(c.GetParentFolderId()) ||
			ptr.Val(newContainer.GetDisplayName()) != ptr.Val(c.GetDisplayName()) {
			delete(cr.cache, folderID)

			if err := cr.addFolder(newContainer); err != nil {
				return resolvedPath{
					idPath:  nil,
					locPath: nil,
					cached:  false,
					deleted: false,
				}, clues.WrapWC(ctx, err, "updating cached container")
			}

			return cr.idToPath(ctx, folderID, depth)
		}
	}

	// If the parent wasn't found and refreshing the current container produced no
	// diffs then delete the current container on the assumption that the parent
	// was deleted and the current container will later get deleted via eventual
	// consistency. If w're wrong then the container will get picked up again on
	// the next backup.
	if resolved.deleted {
		logger.Ctx(ctx).Debug("deleting container since parent was deleted")
		delete(cr.cache, folderID)

		return resolvedPath{
			idPath:  nil,
			locPath: nil,
			cached:  true,
			deleted: true,
		}, nil
	}

	fullPath := resolved.idPath.Append(ptr.Val(c.GetId()))
	c.SetPath(fullPath)

	locPath := resolved.locPath.Append(ptr.Val(c.GetDisplayName()))
	c.SetLocation(locPath)

	return resolvedPath{
		idPath:  fullPath,
		locPath: locPath,
		cached:  true,
		deleted: false,
	}, nil
}

// PathInCache is a utility function to return m365ID of a folder if the
// path.Folders matches the directory of a container within the cache. A boolean
// result is provided to indicate whether the lookup was successful.
func (cr *containerResolver) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || cr == nil {
		return "", false
	}

	for _, cc := range cr.cache {
		if cc.Path() == nil {
			continue
		}

		if cc.Path().String() == pathString {
			return ptr.Val(cc.GetId()), true
		}
	}

	return "", false
}

// LocationInCache is a utility function to return m365ID of a folder if the
// path.Folders matches the directory of a container within the cache. A boolean
// result is provided to indicate whether the lookup was successful.
func (cr *containerResolver) LocationInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || cr == nil {
		return "", false
	}

	for _, cc := range cr.cache {
		if cc.Location() == nil {
			continue
		}

		if cc.Location().String() == pathString {
			return ptr.Val(cc.GetId()), true
		}
	}

	return "", false
}

// addFolder adds a folder to the cache with the given ID. If the item is
// already in the cache does nothing. The path for the item is not modified.
func (cr *containerResolver) addFolder(cf graph.CachedContainer) error {
	var err error

	// Only require a non-nil non-empty parent if the path isn't already populated.
	if cf.Path() != nil {
		err = checkIDAndName(cf)
	} else {
		err = checkRequiredValues(cf)
	}

	if err != nil {
		return clues.Wrap(err, "validating container for cache")
	}

	if _, ok := cr.cache[ptr.Val(cf.GetId())]; ok {
		return nil
	}

	cr.cache[ptr.Val(cf.GetId())] = cf

	return nil
}

func (cr *containerResolver) ItemByID(id string) graph.CachedContainer {
	return cr.cache[id]
}

func (cr *containerResolver) Items() []graph.CachedContainer {
	res := make([]graph.CachedContainer, 0, len(cr.cache))

	for _, c := range cr.cache {
		res = append(res, c)
	}

	return res
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (cr *containerResolver) AddToCache(
	ctx context.Context,
	f graph.Container,
) error {
	temp := &graph.CacheFolder{
		Container: f,
	}
	if err := cr.addFolder(temp); err != nil {
		return clues.WrapWC(ctx, err, "adding cache folder")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, err := cr.idToPath(ctx, ptr.Val(f.GetId()), 0)
	if err != nil {
		return clues.Wrap(err, "adding cache entry")
	}

	return nil
}

func (cr *containerResolver) populatePaths(
	ctx context.Context,
	errs *fault.Bus,
) error {
	var (
		el      = errs.Local()
		lastErr error
	)

	// Populate all folder paths.
	for _, f := range cr.Items() {
		if el.Failure() != nil {
			return el.Failure()
		}

		_, err := cr.idToPath(ctx, ptr.Val(f.GetId()), 0)
		if err != nil {
			err = clues.Wrap(err, "populating path")
			el.AddRecoverable(ctx, err)
			lastErr = err
		}
	}

	return lastErr
}

// ---------------------------------------------------------------------------
// rankedContainerResolver
// ---------------------------------------------------------------------------

type rankedContainerResolver struct {
	graph.ContainerResolver
	// resolvedInclude is the ordered list of resolved container IDs to add to the
	// start of the Items result set.
	resolvedInclude []string
	// resolvedExclude is the set of items that shouldn't be included in the
	// result of Items or ItemByID. Uses actual container IDs instead of
	// well-known names.
	resolvedExclude map[string]struct{}
}

// newRankedContainerResolver creates a wrapper around base that returns results
// from Items in priority order. Priority is defined by includeRankedIDs. All
// items that don't appear in includeRankedIDs are considered to have equal
// priority but lower priority than those in includeRankedIDs.
//
// includeRankedIDs is the set of containers to place at the start of the result
// of Items in the order they should appear. IDs can either be actual
// container IDs or well-known container IDs like "inbox".
//
// excludeIDs is the set of IDs that shouldn't be in the results returned by
// Items. IDs can either be actual container IDs or well-known container IDs
// like "inbox".
//
// The include set takes priority over the exclude set, so container IDs
// appearing in both will be considered included and be returned by calls like
// Items and ItemByID.
func newRankedContainerResolver(
	ctx context.Context,
	base graph.ContainerResolver,
	getter containerGetter,
	userID string,
	includeRankedIDs []string,
	excludeIDs []string,
) (*rankedContainerResolver, error) {
	if base == nil {
		return nil, clues.New("nil base ContainerResolver")
	}

	cr := &rankedContainerResolver{
		resolvedInclude:   make([]string, 0, len(includeRankedIDs)),
		resolvedExclude:   make(map[string]struct{}, len(excludeIDs)),
		ContainerResolver: base,
	}

	// For both includes and excludes we need to get the container IDs from graph.
	// This is required because the user could hand us one of the "well-known"
	// IDs, which we don't use in the underlying container resolver. Resolving
	// these here will allow us to match by ID later on.
	for _, id := range includeRankedIDs {
		ictx := clues.Add(ctx, "container_id", id)

		c, err := getter.GetContainerByID(ctx, userID, id)
		if err != nil {
			return nil, clues.WrapWC(ictx, err, "getting ranked container")
		}

		gotID := ptr.Val(c.GetId())
		if len(gotID) == 0 {
			return nil, clues.NewWC(ictx, "ranked include container missing ID")
		}

		cr.resolvedInclude = append(cr.resolvedInclude, gotID)
	}

	for _, id := range excludeIDs {
		ictx := clues.Add(ctx, "container_id", id)

		c, err := getter.GetContainerByID(ctx, userID, id)
		if err != nil {
			return nil, clues.WrapWC(ictx, err, "getting exclude container")
		}

		gotID := ptr.Val(c.GetId())
		if len(gotID) == 0 {
			return nil, clues.NewWC(ictx, "exclude container missing ID")
		}

		cr.resolvedExclude[gotID] = struct{}{}
	}

	return cr, nil
}

func (cr *rankedContainerResolver) Items() []graph.CachedContainer {
	found := cr.ContainerResolver.Items()
	res := make([]graph.CachedContainer, 0, len(found))

	// Add the ranked items first.
	//
	// TODO(ashmrtn): If we need to handle a large number of ranked items we
	// should think about making a map of the ranked items for fast lookups later
	// in the function.
	for _, include := range cr.resolvedInclude {
		if c := cr.ContainerResolver.ItemByID(include); c != nil {
			res = append(res, c)
		}
	}

	// Add the remaining, filtering out any of the ones we need to exclude or that
	// we already added because they were ranked.
	for _, c := range found {
		if _, ok := cr.resolvedExclude[ptr.Val(c.GetId())]; ok {
			continue
		}

		if slices.Contains(cr.resolvedInclude, ptr.Val(c.GetId())) {
			continue
		}

		res = append(res, c)
	}

	return res
}

func (cr *rankedContainerResolver) ItemByID(id string) graph.CachedContainer {
	// Includes take priority over excludes so check those too.
	_, exclude := cr.resolvedExclude[id]
	includeIdx := slices.Index(cr.resolvedInclude, id)

	if exclude && includeIdx == -1 {
		return nil
	}

	return cr.ContainerResolver.ItemByID(id)
}
