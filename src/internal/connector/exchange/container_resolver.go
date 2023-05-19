package exchange

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
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

func newContainerResolver() *containerResolver {
	return &containerResolver{
		cache: map[string]graph.CachedContainer{},
	}
}

type containerResolver struct {
	cache map[string]graph.CachedContainer
}

func (cr *containerResolver) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, *path.Builder, error) {
	return cr.idToPath(ctx, folderID, 0)
}

func (cr *containerResolver) idToPath(
	ctx context.Context,
	folderID string,
	depth int,
) (*path.Builder, *path.Builder, error) {
	ctx = clues.Add(ctx, "container_id", folderID)

	if depth >= maxIterations {
		return nil, nil, clues.New("path contains cycle or is too tall").WithClues(ctx)
	}

	c, ok := cr.cache[folderID]
	if !ok {
		return nil, nil, clues.New("folder not cached").WithClues(ctx)
	}

	p := c.Path()
	if p != nil {
		return p, c.Location(), nil
	}

	parentPath, parentLoc, err := cr.idToPath(
		ctx,
		ptr.Val(c.GetParentFolderId()),
		depth+1)
	if err != nil {
		return nil, nil, clues.Wrap(err, "retrieving parent folder")
	}

	fullPath := parentPath.Append(ptr.Val(c.GetId()))
	c.SetPath(fullPath)

	locPath := parentLoc.Append(ptr.Val(c.GetDisplayName()))
	c.SetLocation(locPath)

	return fullPath, locPath, nil
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
	_, _, err := cr.IDToPath(ctx, ptr.Val(f.GetId()))
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

		_, _, err := cr.IDToPath(ctx, ptr.Val(f.GetId()))
		if err != nil {
			err = clues.Wrap(err, "populating path")
			el.AddRecoverable(err)
			lastErr = err
		}
	}

	return lastErr
}
