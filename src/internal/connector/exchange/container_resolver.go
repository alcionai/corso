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
}

func (cr *containerResolver) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, error) {
	return cr.idToPath(ctx, folderID, 0)
}

func (cr *containerResolver) idToPath(
	ctx context.Context,
	folderID string,
	depth int,
) (*path.Builder, error) {
	if depth >= maxIterations {
		return nil, errors.New("path contains cycle or is too tall")
	}

	c, ok := cr.cache[folderID]
	if !ok {
		return nil, errors.Errorf("folder %s not cached", folderID)
	}

	p := c.Path()
	if p != nil {
		return p, nil
	}

	parentPath, err := cr.idToPath(ctx, *c.GetParentFolderId(), depth+1)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving parent folder")
	}

	fullPath := parentPath.Append(*c.GetDisplayName())
	c.SetPath(fullPath)

	return fullPath, nil
}

// PathInCache utility function to return m365ID of folder if the path.Folders
// matches the directory of a container within the cache. A boolean result
// is provided to indicate whether the lookup was successful.
func (cr *containerResolver) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || cr == nil {
		return "", false
	}

	for _, contain := range cr.cache {
		if contain.Path() == nil {
			continue
		}

		if contain.Path().String() == pathString {
			return *contain.GetId(), true
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
func (cr *containerResolver) AddToCache(ctx context.Context, f graph.Container) error {
	temp := graph.CacheFolder{
		Container: f,
	}

	if err := cr.addFolder(temp); err != nil {
		return errors.Wrap(err, "adding cache folder")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, err := cr.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}

func (cr *containerResolver) populatePaths(ctx context.Context) error {
	var errs *multierror.Error

	// Populate all folder paths.
	for _, f := range cr.Items() {
		_, err := cr.IDToPath(ctx, *f.GetId())
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "populating path"))
		}
	}

	return errs.ErrorOrNil()
}
