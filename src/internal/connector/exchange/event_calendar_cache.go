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
	cache          map[string]graph.CachedContainer
	gs             graph.Service
	userID, rootID string
}

// Populate utility function for populating eventCalendarCache.
// Executes 1 additional Graph Query
// @param baseID: M365ID of the base exchange.Calendar
func (ecc *eventCalendarCache) Populate(
	ctx context.Context,
	baseID string,
	baseContainerPath ...string,
) error {
	if ecc.cache == nil {
		ecc.cache = map[string]graph.CachedContainer{}
	}

	options, err := optionsForCalendars([]string{"name"})
	if err != nil {
		return err
	}

	directories := make(map[string]graph.Container)
	errUpdater := func(s string, e error) {
		err = support.WrapAndAppend(s, e, err)
	}

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
		ecc.rootID,
		errUpdater,
	)

	iterateErr := iter.Iterate(ctx, cb)
	if iterateErr != nil {
		return iterateErr
	}

	if err != nil {
		return err
	}

	for _, containerr := range directories {
		if err := ecc.AddToCache(ctx, containerr); err != nil {
			iterateErr = support.WrapAndAppend(
				"failure adding "+*containerr.GetDisplayName(),
				err,
				iterateErr)
		}
	}

	return iterateErr
}

func (ecc *eventCalendarCache) IDToPath(
	ctx context.Context,
	calendarID string,
) (*path.Builder, error) {
	c, ok := ecc.cache[calendarID]
	if !ok {
		return nil, errors.Errorf("calendar %s not cached", calendarID)
	}

	p := c.Path()
	if p != nil {
		return p, nil
	}

	parentPath, err := ecc.IDToPath(ctx, *c.GetParentFolderId())
	if err != nil {
		return nil, errors.Wrap(err, "retrieving parent calendar")
	}

	fullPath := parentPath.Append(*c.GetDisplayName())
	c.SetPath(fullPath)

	return fullPath, nil
}

// AddToCache places container into internal cache field. For EventCalendars
// this means that the object has to be transformed prior to calling
// this function.
func (ecc *eventCalendarCache) AddToCache(ctx context.Context, f graph.Container) error {
	ptr := f.GetDisplayName()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without display name", *f.GetId())
	}

	if _, ok := ecc.cache[*f.GetId()]; ok {
		return nil
	}

	ecc.cache[*f.GetId()] = &cacheFolder{
		Container: f,
		p:         path.Builder{}.Append(*f.GetDisplayName()),
	}

	_, err := ecc.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding event cache entry")
	}

	return nil
}

func (ecc *eventCalendarCache) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || ecc.cache == nil {
		return "", false
	}

	for _, containerr := range ecc.cache {
		if containerr.Path() == nil {
			continue
		}

		if containerr.Path().String() == pathString {
			return *containerr.GetId(), true
		}
	}

	return "", false
}

func (ecc *eventCalendarCache) Items() []graph.CachedContainer {
	res := make([]graph.CachedContainer, 0, len(ecc.cache))

	for _, c := range ecc.cache {
		res = append(res, c)
	}

	return res
}
