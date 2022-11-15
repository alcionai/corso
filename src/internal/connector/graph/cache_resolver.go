package graph

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/path"
)

var _ ContainerResolver = &ContainerCache{}

type populatorFunc func(
	ctx context.Context,
	baseID string,
	baseContainerPath ...string,
)

type ContainerCache struct {
	cache map[string]CachedContainer
}

func NewContainerCache() *ContainerCache {
	return &ContainerCache{
		cache: map[string]CachedContainer{},
	}
}

func (cr *ContainerCache) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, error) {
	c, ok := cr.cache[folderID]
	if !ok {
		return nil, errors.Errorf("folder %s not cached", folderID)
	}

	p := c.Path()
	if p != nil {
		return p, nil
	}

	parentPath, err := cr.IDToPath(ctx, *c.GetParentFolderId())
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
func (cr *ContainerCache) PathInCache(pathString string) (string, bool) {
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

// AddFolder adds a folder to the cache with the given ID. If the item is
// already in the cache does nothing. The path for the item is not modified.
func (cr *ContainerCache) AddFolder(cf CacheFolder) error {
	// Only require a non-nil non-empty parent if the path isn't already
	// populated.
	if cf.p != nil {
		if err := CheckIDAndName(cf.Container); err != nil {
			return errors.Wrap(err, "adding item to cache")
		}
	} else {
		if err := CheckRequiredValues(cf.Container); err != nil {
			return errors.Wrap(err, "adding item to cache")
		}
	}

	if _, ok := cr.cache[*cf.GetId()]; ok {
		return nil
	}

	cr.cache[*cf.GetId()] = &cf

	return nil
}

// Items returns the list of Containers in the cache.
func (cr *ContainerCache) Items() []CachedContainer {
	res := make([]CachedContainer, 0, len(cr.cache))

	for _, c := range cr.cache {
		res = append(res, c)
	}

	return res
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (cr *ContainerCache) AddToCache(ctx context.Context, f Container) error {
	temp := CacheFolder{
		Container: f,
	}

	if err := cr.AddFolder(temp); err != nil {
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

// PopulatePaths ensures that all items in the cache can construct valid paths.
func (cr *ContainerCache) PopulatePaths(ctx context.Context) error {
	var errs *multierror.Error

	for _, f := range cr.Items() {
		_, err := cr.IDToPath(ctx, *f.GetId())
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "populating path"))
		}
	}

	return errs.ErrorOrNil()
}
