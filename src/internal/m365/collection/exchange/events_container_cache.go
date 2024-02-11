package exchange

import (
	"context"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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
		return clues.WrapWC(ctx, err, "initializing calendar resolver")
	}

	return nil
}

func isSharedCalendar(defaultCalendarOwner string, c models.Calendarable) bool {
	// If we can't determine the owner, assume the calendar is owned by the
	// user.
	if len(defaultCalendarOwner) == 0 || c.GetOwner() == nil {
		return false
	}

	return !strings.EqualFold(defaultCalendarOwner, ptr.Val(c.GetOwner().GetAddress()))
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
	start := time.Now()

	logger.Ctx(ctx).Info("populating container cache")

	if err := ecc.init(ctx); err != nil {
		return clues.Wrap(err, "initializing")
	}

	el := errs.Local()

	containers, err := ecc.enumer.EnumerateContainers(
		ctx,
		ecc.userID,
		"")
	ctx = clues.Add(ctx, "num_enumerated_containers", len(containers))

	if err != nil {
		return clues.WrapWC(ctx, err, "enumerating containers")
	}

	var defaultCalendarOwner string

	// Determine the owner for the default calendar. We'll use this to detect and
	// skip shared calendars that are not owned by this user.
	for _, c := range containers {
		if ptr.Val(c.GetIsDefaultCalendar()) && c.GetOwner() != nil {
			defaultCalendarOwner = ptr.Val(c.GetOwner().GetAddress())
			ctx = clues.Add(ctx, "default_calendar_owner", defaultCalendarOwner)

			break
		}
	}

	for _, c := range containers {
		if el.Failure() != nil {
			return el.Failure()
		}

		// Skip shared calendars if we have enough information to determine the owner
		if isSharedCalendar(defaultCalendarOwner, c) {
			var ownerEmail string
			if c.GetOwner() != nil {
				ownerEmail = ptr.Val(c.GetOwner().GetAddress())
			}

			logger.Ctx(ctx).Infow(
				"skipping shared calendar",
				"name", ptr.Val(c.GetName()),
				"owner", ownerEmail)

			continue
		}

		cacheFolder := graph.NewCacheFolder(
			api.CalendarDisplayable{Calendarable: c},
			path.Builder{}.Append(ptr.Val(c.GetId())),
			path.Builder{}.Append(ptr.Val(c.GetName())))

		err := ecc.addFolder(&cacheFolder)
		if err != nil {
			err := clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation)
			errs.AddRecoverable(ctx, err)
		}
	}

	if err := ecc.populatePaths(ctx, errs); err != nil {
		return clues.Wrap(err, "populating paths")
	}

	logger.Ctx(ctx).Infow(
		"done populating container cache",
		"duration", time.Since(start))

	return el.Failure()
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (ecc *eventContainerCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkIDAndName(f); err != nil {
		return clues.WrapWC(ctx, err, "validating container")
	}

	temp := graph.NewCacheFolder(
		f,
		path.Builder{}.Append(ptr.Val(f.GetId())),          // storage path
		path.Builder{}.Append(ptr.Val(f.GetDisplayName()))) // display location

	if err := ecc.addFolder(&temp); err != nil {
		return clues.WrapWC(ctx, err, "adding container")
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, _, err := ecc.IDToPath(ctx, ptr.Val(f.GetId()))
	if err != nil {
		return clues.Wrap(err, "setting path to container id")
	}

	return nil
}
