package exchange

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
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
	//wantedOpts := []string{}
	return nil
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
// @see
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
