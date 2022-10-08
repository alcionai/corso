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
	cache          map[string]cachedContainer
	gs             graph.Service
	userID, rootID string
}

func (ecc *eventCalendarCache) populateEventRoot(
	ctx context.Context,
	directoryID string,
	baseContainerPath []string,
) error {
	wantedOpts := []string{"name"}

	opts, err := optionsForIndividualCalendar(wantedOpts)
	if err != nil {
		return errors.Wrapf(err, "getting options for event cache %v", wantedOpts)
	}

	cal, err := ecc.gs.
		Client().
		UsersById(ecc.userID).
		CalendarsById(directoryID).
		Get(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "fetching default calendar "+support.ConnectorStackErrorTrace(err))
	}

	idPtr := cal.GetId()

	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("root calendar has no ID")
	}

	identifier := *idPtr
	transform := CreateCalendarDisplayable(cal, identifier)
	temp := eventCalendar{
		Container: transform,
		p:         path.Builder{}.Append(baseContainerPath...),
	}

	ecc.cache[identifier] = &temp
	ecc.rootID = identifier

	return nil
}

func (ecc *eventCalendarCache) Populate(
	ctx context.Context,
	baseID string,
	baseContainerPath ...string,
) error {
	if err := ecc.init(ctx, baseID, baseContainerPath); err != nil {
		return err
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
		if err := ecc.AddToCache(containerr); err != nil {
			iterateErr = support.WrapAndAppend(
				"failure adding "+*containerr.GetDisplayName(),
				err,
				iterateErr)
		}
	}

	return iterateErr
}

func (ecc *eventCalendarCache) init(
	ctx context.Context,
	baseNode string,
	baseContainerPath []string,
) error {
	if len(baseNode) == 0 {
		return errors.New("m365ID calendarID required for base")
	}

	if ecc.cache == nil {
		ecc.cache = map[string]cachedContainer{}
	}

	return ecc.populateEventRoot(ctx, baseNode, baseContainerPath)
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
func (ecc *eventCalendarCache) AddToCache(f graph.Container) error {
	if err := checkRequiredValues(f); err != nil {
		return err
	}

	if _, ok := ecc.cache[*f.GetId()]; ok {
		return nil
	}

	ecc.cache[*f.GetId()] = &eventCalendar{
		Container: f,
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
