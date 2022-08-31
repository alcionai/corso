package exchange

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/path"
)

// For now this is rather generic, but in the future could be expanded to
// provide other functions that would help us backup metadata about a folder.
type container interface {
	ParentID() string
	DisplayName() string
	ID() string
}

type containerResolver interface {
	Fetch(ctx context.Context, id string) (container, error)
	FetchRoot(context.Context) (container, error)
}

type cachedContainer struct {
	container
	*path.Builder
}

func NewCachingContainerResolver(
	prefix *path.Builder,
	cr containerResolver,
) (*cachingContainerResolver, error) {
	if prefix == nil {
		return nil, errors.New("nil prefix")
	}

	if cr == nil {
		return nil, errors.New("nil resolver")
	}

	return &cachingContainerResolver{
		cached:   map[string]cachedContainer{},
		resolver: cr,
		prefix:   prefix,
	}, nil
}

// cachingContainerResolver is a container (i.e. "folder") path resolver that
// also caches results as it goes along. Results are cached for the lifetime of
// the resolver instance. The resolver takes a containerID for the final path
// element and recursively resolves the parent containers to form a full path
// (e.x. resolve on `this/is/a/path` would first fetch information about `path`,
// then `a` and so on until it reached `this`).
type cachingContainerResolver struct {
	// cached maps from container ID -> cached information about the container.
	cached map[string]cachedContainer
	// Function to lookup directories with an external API. Must contain all
	// information necessary for a lookup except for the ID of the object being
	// looked up.
	resolver containerResolver
	// Prefix path for the root container in the hierarchy. Completely replaces
	// the path of the root container, allowing callers to easily modify that to
	// fit their needs. For example, if a root container had a path of
	// "/Top of Information Store" but prefix was path.Builder{} then all returned
	// paths would start with one of the immediate children of the root container.
	prefix *path.Builder
}

func (ccr *cachingContainerResolver) Initialize(ctx context.Context) error {
	c, err := ccr.resolver.FetchRoot(ctx)
	if err != nil {
		return errors.Wrap(err, "fetching root container")
	}

	ccr.cached[c.ID()] = cachedContainer{
		container: c,
		Builder:   ccr.prefix,
	}

	return nil
}

// Lookup takes a folder ID and recursively resolves parent IDs to get the full
// path of the item. It is the caller's responsibility to convert the returned
// Builder into a resource-specific path as needed.
func (ccr *cachingContainerResolver) Lookup(
	ctx context.Context,
	containerID string,
) (*path.Builder, error) {
	p, ok := ccr.cached[containerID]
	if ok {
		return p.Builder, nil
	}

	c, err := ccr.resolver.Fetch(ctx, containerID)
	if err != nil {
		return nil, errors.Wrapf(err, "fetching folder with ID %s", containerID)
	}

	parentPath, err := ccr.Lookup(ctx, c.ParentID())
	if err != nil {
		return nil, err
	}

	fullPath := parentPath.Append(c.DisplayName())

	ccr.cached[c.ID()] = cachedContainer{
		container: c,
		Builder:   fullPath,
	}

	return fullPath, nil
}
