package exchange

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
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
		fn func(graph.CacheFolder) error,
	) error
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
	// newAdditions is a map of displayName: ID tuples.
	// Since we don't know the ID of newly added containers until
	// they've been generated via graph api, we don't have an easy
	// way of mapping destination names onto CachedContainers.
	// This map should get updated on every call to AddToCache, and
	// can be accessed via the AddedIDs func, so that we avoid
	// FolderAlreadyExists errors for types that store IDs instead
	// of display names in their path.
	newAdditions map[string]string
}

func (cr *containerResolver) IDToPath(
	ctx context.Context,
	folderID string,
	useIDInPath bool,
) (*path.Builder, *path.Builder, error) {
	return cr.idToPath(ctx, folderID, 0, useIDInPath)
}

func (cr *containerResolver) idToPath(
	ctx context.Context,
	folderID string,
	depth int,
	useIDInPath bool,
) (*path.Builder, *path.Builder, error) {
	if depth >= maxIterations {
		return nil, nil, errors.New("path contains cycle or is too tall")
	}

	c, ok := cr.cache[folderID]
	if !ok {
		return nil, nil, errors.Errorf("folder %s not cached", folderID)
	}

	p := c.Path()
	if p != nil {
		return p, c.Location(), nil
	}

	parentPath, parentLoc, err := cr.idToPath(ctx, *c.GetParentFolderId(), depth+1, useIDInPath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "retrieving parent folder")
	}

	toAppend := *c.GetDisplayName()
	if useIDInPath {
		toAppend = *c.GetId()
	}

	fullPath := parentPath.Append(toAppend)
	c.SetPath(fullPath)

	var locPath *path.Builder

	if parentLoc != nil {
		locPath = parentLoc.Append(*c.GetDisplayName())
		c.SetLocation(locPath)
	}

	return fullPath, locPath, nil
}

// PathInCache utility function to return m365ID of folder if the path.Folders
// matches the directory of a container within the cache. A boolean result
// is provided to indicate whether the lookup was successful.
func (cr *containerResolver) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || cr == nil {
		return "", false
	}

	for _, cc := range cr.cache {
		if cc.Path() == nil {
			continue
		}

		if cc.Path().String() == pathString {
			return *cc.GetId(), true
		}
	}

	return "", false
}

// addFolder adds a folder to the cache with the given ID. If the item is
// already in the cache does nothing. The path for the item is not modified.
func (cr *containerResolver) addFolder(cf graph.CacheFolder) error {
	// Only require a non-nil non-empty parent if the path isn't already populated.
	if cf.Path() != nil {
		if err := checkIDAndName(cf.Container); err != nil {
			return errors.Wrap(err, "adding item to cache")
		}
	} else {
		if err := checkRequiredValues(cf.Container); err != nil {
			return errors.Wrap(err, "adding item to cache")
		}
	}

	if _, ok := cr.cache[*cf.GetId()]; ok {
		return nil
	}

	cr.cache[*cf.GetId()] = &cf

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
	useIDInPath bool,
) error {
	temp := graph.CacheFolder{
		Container: f,
	}

	if len(cr.newAdditions) == 0 {
		cr.newAdditions = map[string]string{}
	}

	cr.newAdditions[*f.GetDisplayName()] = *f.GetId()

	if err := cr.addFolder(temp); err != nil {
		delete(cr.newAdditions, *f.GetDisplayName())
		return errors.Wrap(err, "adding cache folder")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, _, err := cr.IDToPath(ctx, *f.GetId(), useIDInPath)
	if err != nil {
		delete(cr.newAdditions, *f.GetDisplayName())
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}

// DestinationNameToID returns the ID of the destination container.  Dest is
// assumed to be a display name.  The ID is only populated if the destination
// was added using `AddToCache()`.  Returns an empty string if not found.
func (cr *containerResolver) DestinationNameToID(dest string) string {
	return cr.newAdditions[dest]
}

func (cr *containerResolver) populatePaths(ctx context.Context, useIDInPath bool) error {
	var errs *multierror.Error

	// Populate all folder paths.
	for _, f := range cr.Items() {
		_, _, err := cr.IDToPath(ctx, *f.GetId(), useIDInPath)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "populating path"))
		}
	}

	return errs.ErrorOrNil()
}
