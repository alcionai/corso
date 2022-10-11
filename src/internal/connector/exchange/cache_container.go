package exchange

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

// cachedContainer is used for local unit tests but also makes it so that this
// code can be broken into generic- and service-specific chunks later on to
// reuse logic in IDToPath.
type cachedContainer interface {
	graph.Container
	Path() *path.Builder
	SetPath(*path.Builder)
}

// checkRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func checkRequiredValues(c graph.Container) error {
	idPtr := c.GetId()
	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("folder without ID")
	}

	ptr := c.GetDisplayName()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without display name", *idPtr)
	}

	ptr = c.GetParentFolderId()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without parent ID", *idPtr)
	}

	return nil
}

//======================================
// cachedContainer Implementations
//======================

var (
	_ cachedContainer = &eventCalendar{}
	_ cachedContainer = &contactFolder{}
	_ cachedContainer = &mailFolder{}
)

type contactFolder struct {
	graph.Container
	p *path.Builder
}

func (cf contactFolder) Path() *path.Builder {
	return cf.p
}

func (cf *contactFolder) SetPath(newPath *path.Builder) {
	cf.p = newPath
}

type eventCalendar struct {
	graph.Container
	p *path.Builder
}

func (ev eventCalendar) Path() *path.Builder {
	return ev.p
}

func (ev *eventCalendar) SetPath(newPath *path.Builder) {
	ev.p = newPath
}

// mailFolder structure that implements the cachedContainer interface
type mailFolder struct {
	graph.Container
	p *path.Builder
}

//=========================================
// Required Functions to satisfy interfaces
//=====================================

func (mf mailFolder) Path() *path.Builder {
	return mf.p
}

func (mf *mailFolder) SetPath(newPath *path.Builder) {
	mf.p = newPath
}

//
// CalendarDisplayable is a transformative struct that aligns
// models.Calendarable interface with the container interface.
// Calendars do not have the 2 of the
type CalendarDisplayable struct {
	models.Calendarable
	parentID string
}

// GetDisplayName returns the *string of the calendar name
func (c CalendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

// GetParentFolderId returns the default calendar name address
// EventCalendars have a flat hierarchy and Calendars are rooted
// at the default
//nolint:revive
func (c CalendarDisplayable) GetParentFolderId() *string {
	return &c.parentID
}

// CreateCalendarDisplayable helper function to create the
// calendarDisplayable during msgraph-sdk-go iterative process
// @param entry is the input supplied by pageIterator.Iterate()
// @param parentID of Calendar sets. Only populate when used with
// EventCalendarCache
func CreateCalendarDisplayable(entry any, parentID string) *CalendarDisplayable {
	calendar, ok := entry.(models.Calendarable)
	if !ok {
		return nil
	}

	return &CalendarDisplayable{
		Calendarable: calendar,
		parentID:     parentID,
	}
}
