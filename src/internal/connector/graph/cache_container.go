package graph

import (
	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/path"
)

// CachedContainer is used for local unit tests but also makes it so that this
// code can be broken into generic- and service-specific chunks later on to
// reuse logic in IDToPath.
type CachedContainer interface {
	Container
	Path() *path.Builder
	SetPath(*path.Builder)
}

// checkRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func CheckRequiredValues(c Container) error {
	idPtr := c.GetId()
	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("folder without ID")
	}

	ptr := c.GetDisplayName()
	if ptr == nil || len(*ptr) == 0 {
		return clues.New("folder missing display name").With("container_id", *idPtr)
	}

	ptr = c.GetParentFolderId()
	if ptr == nil || len(*ptr) == 0 {
		return clues.New("folder missing parent ID").With("container_parent_id", *idPtr)
	}

	return nil
}

// ======================================
// cachedContainer Implementations
// ======================================

var _ CachedContainer = &CacheFolder{}

type CacheFolder struct {
	Container
	p *path.Builder
}

// NewCacheFolder public constructor for struct
func NewCacheFolder(c Container, pb *path.Builder) CacheFolder {
	cf := CacheFolder{
		Container: c,
		p:         pb,
	}

	return cf
}

// =========================================
// Required Functions to satisfy interfaces
// =========================================

func (cf CacheFolder) Path() *path.Builder {
	return cf.p
}

func (cf *CacheFolder) SetPath(newPath *path.Builder) {
	cf.p = newPath
}

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
//
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
