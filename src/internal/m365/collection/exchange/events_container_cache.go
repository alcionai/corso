package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ graph.ContainerResolver = &eventContainerCache{}

type eventContainerCache struct {
	*containerResolver
	enumer containersEnumerator[models.Calendarable]
	getter containerGetter
	userID string
}

// init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache]
func (ecc *eventContainerCache) init(
	ctx context.Context,
) error {
	if ecc.containerResolver == nil {
		ecc.containerResolver = newContainerResolver(nil)
	}

	return ecc.populateEventRoot(ctx)
}

// populateEventRoot manually fetches directories that are not returned during Graph for msgraph-sdk-go v. 40+
// DefaultCalendar is the traditional "Calendar".
// Action ensures that cache will stop at appropriate level.
// @error iff the struct is not properly instantiated
func (ecc *eventContainerCache) populateEventRoot(ctx context.Context) error {
	container := api.DefaultCalendar

	f, err := ecc.getter.GetContainerByID(ctx, ecc.userID, container)
	if err != nil {
		return clues.Wrap(err, "fetching calendar")
	}

	temp := graph.NewCacheFolder(
		f,
		path.Builder{}.Append(ptr.Val(f.GetId())),          // storage path
		path.Builder{}.Append(ptr.Val(f.GetDisplayName()))) // display location
	if err := ecc.addFolder(&temp); err != nil {
		return clues.Wrap(err, "initializing calendar resolver").WithClues(ctx)
	}

	return nil
}

// Populate utility function for populating eventCalendarCache.
// Executes 1 additional Graph Query
// @param baseID: ignored. Present to conform to interface
func (ecc *eventContainerCache) Populate(
	ctx context.Context,
	errs *fault.Bus,
	baseID string,
	baseContainerPath ...string,
) error {
	if err := ecc.init(ctx); err != nil {
		return clues.Wrap(err, "initializing")
	}

	el := errs.Local()

	containers, err := ecc.enumer.EnumerateContainers(
		ctx,
		ecc.userID,
		"",
		false)
	if err != nil {
		return clues.Wrap(err, "enumerating containers")
	}

	for _, c := range containers {
		if el.Failure() != nil {
			return el.Failure()
		}

		cacheFolder := graph.NewCacheFolder(
			api.CalendarDisplayable{Calendarable: c},
			path.Builder{}.Append(ptr.Val(c.GetId())),
			path.Builder{}.Append(ptr.Val(c.GetName())))

		err := ecc.addFolder(&cacheFolder)
		if err != nil {
			errs.AddRecoverable(
				ctx,
				graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
		}
	}

	if err := ecc.populatePaths(ctx, errs); err != nil {
		return clues.Wrap(err, "populating paths")
	}

	return el.Failure()
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (ecc *eventContainerCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkIDAndName(f); err != nil {
		return clues.Wrap(err, "validating container").WithClues(ctx)
	}

	temp := graph.NewCacheFolder(
		f,
		path.Builder{}.Append(ptr.Val(f.GetId())),          // storage path
		path.Builder{}.Append(ptr.Val(f.GetDisplayName()))) // display location

	if err := ecc.addFolder(&temp); err != nil {
		return clues.Wrap(err, "adding container").WithClues(ctx)
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, _, err := ecc.IDToPath(ctx, ptr.Val(f.GetId()))
	if err != nil {
		return clues.Wrap(err, "setting path to container id")
	}

	return nil
}
