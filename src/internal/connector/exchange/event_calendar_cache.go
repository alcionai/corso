package exchange

import (
	"context"

	mscal "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &eventCalendarCache{}

type eventCalendarCache struct {
	*containerResolver
	gs     graph.Service
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

	options, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return err
	}

	var (
		errs        error
		directories = make([]graph.Container, 0)
	)

	builder := ecc.gs.Client().UsersById(ecc.userID).Calendars()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, cal := range resp.GetValue() {
			temp := CreateCalendarDisplayable(cal)
			if err := checkIDAndName(temp); err != nil {
				errs = support.WrapAndAppend(
					"adding folder to cache",
					err,
					errs,
				)

				continue
			}

			directories = append(directories, temp)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = mscal.NewCalendarsRequestBuilder(*resp.GetOdataNextLink(), ecc.gs.Adapter())
	}

	for _, container := range directories {
		temp := cacheFolder{
			Container: container,
			p:         path.Builder{}.Append(*container.GetDisplayName()),
		}

		if err := ecc.addFolder(temp); err != nil {
			errs = support.WrapAndAppend(
				"failure adding "+*container.GetDisplayName(),
				err,
				errs)
		}
	}

	return errs
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (ecc *eventCalendarCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkIDAndName(f); err != nil {
		return errors.Wrap(err, "adding cache folder")
	}

	temp := cacheFolder{
		Container: f,
		p:         path.Builder{}.Append(*f.GetDisplayName()),
	}

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
