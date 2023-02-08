package exchange

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &eventCalendarCache{}

type eventCalendarCache struct {
	*containerResolver
	enumer containersEnumerator
	getter containerGetter
	userID string
}

// init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache]
func (ecc *eventCalendarCache) init(
	ctx context.Context,
) error {
	if ecc.containerResolver == nil {
		ecc.containerResolver = newContainerResolver()
	}

	return ecc.populateEventRoot(ctx)
}

// populateEventRoot manually fetches directories that are not returned during Graph for msgraph-sdk-go v. 40+
// DefaultCalendar is the traditional "Calendar".
// Action ensures that cache will stop at appropriate level.
// @error iff the struct is not properly instantiated
func (ecc *eventCalendarCache) populateEventRoot(ctx context.Context) error {
	container := DefaultCalendar

	f, err := ecc.getter.GetContainerByID(ctx, ecc.userID, container)
	if err != nil {
		return errors.Wrap(err, "fetching calendar "+support.ConnectorStackErrorTrace(err))
	}

	temp := graph.NewCacheFolder(
		f,
		path.Builder{}.Append(*f.GetId()), // storage path
		path.Builder{}.Append(*f.GetDisplayName())) // display location
	if err := ecc.addFolder(temp); err != nil {
		return errors.Wrap(err, "initializing calendar resolver")
	}

	return nil
}

// Populate utility function for populating eventCalendarCache.
// Executes 1 additional Graph Query
// @param baseID: ignored. Present to conform to interface
func (ecc *eventCalendarCache) Populate(
	ctx context.Context,
	baseID string,
	baseContainerPath ...string,
) error {
	if err := ecc.init(ctx); err != nil {
		return errors.Wrap(err, "initializing")
	}

	err := ecc.enumer.EnumerateContainers(
		ctx,
		ecc.userID,
		"",
		ecc.addFolder)
	if err != nil {
		return errors.Wrap(err, "enumerating containers")
	}

	if err := ecc.populatePaths(ctx, true); err != nil {
		return errors.Wrap(err, "establishing calendar paths")
	}

	return nil
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (ecc *eventCalendarCache) AddToCache(ctx context.Context, f graph.Container, useIDInPath bool) error {
	if err := checkIDAndName(f); err != nil {
		return errors.Wrap(err, "validating container")
	}

	temp := graph.NewCacheFolder(
		f,
		path.Builder{}.Append(*f.GetId()), // storage path
		path.Builder{}.Append(*f.GetDisplayName())) // display location

	if err := ecc.addFolder(temp); err != nil {
		return errors.Wrap(err, "adding container")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, _, err := ecc.IDToPath(ctx, *f.GetId(), true)
	if err != nil {
		return errors.Wrap(err, "setting path to container id")
	}

	return nil
}
