package exchange

import (
	"context"

	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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
		asyncError  error
		directories = make(map[string]graph.Container)
		errUpdater  = func(s string, e error) {
			asyncError = support.WrapAndAppend(s, e, err)
		}
	)

	query, err := ecc.gs.Client().UsersById(ecc.userID).Calendars().Get(ctx, options)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	iter, err := msgraphgocore.NewPageIterator(
		query,
		ecc.gs.Adapter(),
		models.CreateCalendarCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return err
	}

	cb := IterativeCollectCalendarContainers(
		directories,
		"",
		errUpdater,
	)

	iterateErr := iter.Iterate(ctx, cb)
	if iterateErr != nil {
		return errors.Wrap(iterateErr, support.ConnectorStackErrorTrace(iterateErr))
	}

	// check for errors created during iteration
	if asyncError != nil {
		return err
	}

	for _, container := range directories {
		if err := checkIDAndName(container); err != nil {
			iterateErr = support.WrapAndAppend(
				"adding folder to cache",
				err,
				iterateErr,
			)

			continue
		}

		temp := cacheFolder{
			Container: container,
			p:         path.Builder{}.Append(*container.GetDisplayName()),
		}

		if err := ecc.addFolder(temp); err != nil {
			iterateErr = support.WrapAndAppend(
				"failure adding "+*container.GetDisplayName(),
				err,
				iterateErr)
		}
	}

	return iterateErr
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
