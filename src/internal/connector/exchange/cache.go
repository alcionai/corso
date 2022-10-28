package exchange

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

func newFolderCache() *folderCache {
	return &folderCache{
		cache: map[string]graph.CachedContainer{},
	}
}

type folderCache struct {
	cache map[string]graph.CachedContainer
}

func (fc *folderCache) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, error) {
	c, ok := fc.cache[folderID]
	if !ok {
		return nil, errors.Errorf("folder %s not cached", folderID)
	}

	p := c.Path()
	if p != nil {
		return p, nil
	}

	parentPath, err := fc.IDToPath(ctx, *c.GetParentFolderId())
	if err != nil {
		return nil, errors.Wrap(err, "retrieving parent folder")
	}

	fullPath := parentPath.Append(*c.GetDisplayName())
	c.SetPath(fullPath)

	return fullPath, nil
}

// PathInCache utility function to return m365ID of folder if the pathString
// matches the path of a container within the cache. A boolean function
// accompanies the call to indicate whether the lookup was successful.
func (fc *folderCache) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || fc.cache == nil {
		return "", false
	}

	for _, contain := range fc.cache {
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
func (fc *folderCache) addFolder(cf cacheFolder) error {
	// Only require a non-nil non-empty parent if the path isn't already
	// populated.
	if cf.p != nil {
		if err := checkIDAndName(cf.Container); err != nil {
			return errors.Wrap(err, "adding item to cache")
		}
	} else {
		if err := checkRequiredValues(cf.Container); err != nil {
			return errors.Wrap(err, "adding item to cache")
		}
	}

	if _, ok := fc.cache[*cf.GetId()]; ok {
		return nil
	}

	fc.cache[*cf.GetId()] = &cf

	return nil
}

func (fc *folderCache) Items() []graph.CachedContainer {
	res := make([]graph.CachedContainer, 0, len(fc.cache))

	for _, c := range fc.cache {
		res = append(res, c)
	}

	return res
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (fc *folderCache) AddToCache(ctx context.Context, f graph.Container) error {
	temp := cacheFolder{
		Container: f,
	}

	if err := fc.addFolder(temp); err != nil {
		return errors.Wrap(err, "adding cache folder")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, err := fc.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}

func (fc *folderCache) populatePaths(ctx context.Context) error {
	var errs *multierror.Error

	// Populate all folder paths.
	for _, f := range fc.Items() {
		_, err := fc.IDToPath(ctx, *f.GetId())
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "populating path"))
		}
	}

	return errs.ErrorOrNil()
}
