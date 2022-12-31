package exchange

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &eventCalendarCache{}

type eventCalendarCache struct {
	*containerResolver
	enumer enumerateContainerser
	userID string
}

// Populate utility function for populating eventCalendarCache.
// Executes 1 additional Graph Query
// @param baseID: ignored. Present to conform to interface
func (ecc *eventCalendarCache) Populate(
	ctx context.Context,
	baseID string,
	baseContainerPath ...string,
) error {
	if ecc.containerResolver == nil {
		ecc.containerResolver = newContainerResolver()
	}

	err := ecc.enumer.EnumerateContainers(ctx, ecc.userID, "", ecc.addFolder)
	if err != nil {
		return err
	}

	return nil
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (ecc *eventCalendarCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkIDAndName(f); err != nil {
		return errors.Wrap(err, "adding cache folder")
	}

	temp := graph.NewCacheFolder(f, path.Builder{}.Append(*f.GetDisplayName()))

	if err := ecc.addFolder(temp); err != nil {
		return errors.Wrap(err, "adding cache folder")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, err := ecc.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}
