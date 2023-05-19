package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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

type containersEnumerator interface {
	EnumerateContainers(
		ctx context.Context,
		userID, baseDirID string,
		fn func(graph.CachedContainer) error,
		errs *fault.Bus,
	) error
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
		return nil, nil, clues.New("folder not cached").WithClues(ctx)
	}

	p := c.Path()
	if p == nil {
		return nil, nil, clues.New("folder has no path").WithClues(ctx)
	}

	return p, c.Location(), nil
}

func (cr *containerResolver) refreshContainer(
	ctx context.Context,
	id string,
) (graph.CachedContainer, bool, error) {
	ctx = clues.Add(ctx, "refresh_container_id", id)
	logger.Ctx(ctx).Debug("refreshing container")

	if cr.refresher == nil {
		return nil, false, clues.New("nil refresher")
	}

	c, err := cr.refresher.refreshContainer(ctx, id)
	if err != nil && graph.IsErrDeletedInFlight(err) {
		logger.Ctx(ctx).Debug("container deleted")
		return nil, true, nil
	} else if err != nil {
		// This is some other error, just return it.
		return nil, false, clues.Wrap(err, "refreshing container").WithClues(ctx)
	}

	return c, false, nil
}

func (cr *containerResolver) idToPath(
	ctx context.Context,
	folderID string,
	depth int,
) (*path.Builder, *path.Builder, bool, bool, error) {
	ctx = clues.Add(ctx, "container_id", folderID)

	if depth >= maxIterations {
		return nil, nil, false, false, clues.New("path contains cycle or is too tall").WithClues(ctx)
	}

	c, ok := cr.cache[folderID]
	if !ok {
		c, shouldDelete, err := cr.refreshContainer(ctx, folderID)
		if err != nil {
			return nil, nil, false, false, clues.Wrap(err, "fetching uncached container")
		}

		if shouldDelete {
			logger.Ctx(ctx).Debug("fetching uncached folder showed it was deleted")
			return nil, nil, false, shouldDelete, err
		}

		if err := cr.addFolder(c); err != nil {
			return nil, nil, false, false, clues.Wrap(err, "adding new folder").WithClues(ctx)
		}

		// Retry populating this container's paths.
		// TODO(ashmrtn): May want to bump the depth here just so we don't get stuck
		// retrying too much if for some reason things keep moving around?
		pth, loc, _, shouldDelete, err := cr.idToPath(ctx, folderID, depth)
		if err != nil {
			err = clues.Wrap(err, "retry populating uncached folder")
		}

		return pth, loc, false, shouldDelete, err
	}

	p := c.Path()
	if p != nil {
		return p, c.Location(), true, false, nil
	}

	parentPath, parentLoc, parentCached, shouldDelete, err := cr.idToPath(
		ctx,
		ptr.Val(c.GetParentFolderId()),
		depth+1)
	if err != nil {
		return nil, nil, true, false, clues.Wrap(err, "retrieving parent folder")
	}

	if !parentCached {
		logger.Ctx(ctx).Debug("parent folder was refreshed")

		newContainer, currentShouldDelete, err := cr.refreshContainer(ctx, folderID)
		if err != nil {
			return nil, nil, true, false, clues.Wrap(err, "refreshing container").WithClues(ctx)
		}

		if currentShouldDelete {
			logger.Ctx(ctx).Debug("refreshing folder showed it was deleted")
			delete(cr.cache, folderID)

			return nil, nil, true, true, nil
		}

		// TODO(ashmrtn): May want to bump the depth here just so we don't get stuck
		// retrying too much if for some reason things keep moving around?
		if ptr.Val(newContainer.GetParentFolderId()) != ptr.Val(c.GetParentFolderId()) ||
			ptr.Val(newContainer.GetDisplayName()) != ptr.Val(c.GetDisplayName()) {
			delete(cr.cache, folderID)

			if err := cr.addFolder(newContainer); err != nil {
				return nil, nil, false, false, clues.Wrap(err, "updating cached folder").WithClues(ctx)
			}

			return cr.idToPath(ctx, folderID, depth)
		}
	}

	// If the parent wasn't found and refreshing the folder itself showed it
	// hadn't changed then just delete it.
	if shouldDelete {
		logger.Ctx(ctx).Debug("deleting folder since parent was deleted")
		delete(cr.cache, folderID)

		return nil, nil, true, true, nil
	}

	fullPath := parentPath.Append(ptr.Val(c.GetId()))
	c.SetPath(fullPath)

	locPath := parentLoc.Append(ptr.Val(c.GetDisplayName()))
	c.SetLocation(locPath)

	return fullPath, locPath, true, false, nil
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
	// Only require a non-nil non-empty parent if the path isn't already populated.
	if cf.Path() != nil {
		if err := checkIDAndName(cf); err != nil {
			return clues.Wrap(err, "adding item to cache")
		}
	} else {
		if err := checkRequiredValues(cf); err != nil {
			return clues.Wrap(err, "adding item to cache")
		}
	}

	if _, ok := cr.cache[ptr.Val(cf.GetId())]; ok {
		return nil
	}

	cr.cache[ptr.Val(cf.GetId())] = cf

	return nil
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
		return clues.Wrap(err, "adding cache folder").WithClues(ctx)
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, _, _, _, err := cr.idToPath(ctx, ptr.Val(f.GetId()), 0)
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

		_, _, _, _, err := cr.idToPath(ctx, ptr.Val(f.GetId()), 0)
		if err != nil {
			err = clues.Wrap(err, "populating path")
			el.AddRecoverable(err)
			lastErr = err
		}
	}

	return lastErr
}
