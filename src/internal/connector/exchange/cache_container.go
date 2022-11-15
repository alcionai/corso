package exchange

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CalendarDisplayable is a transformative struct that aligns
// models.Calendarable interface with the container interface.
// Calendars do not have a parentFolderID. Therefore,
// the call will always return nil
type CalendarDisplayable struct {
	models.Calendarable
}

// GetDisplayName returns the *string of the models.Calendable
// variant:  calendar.GetName()
func (c CalendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

// GetParentFolderId returns the default calendar name address
// EventCalendars have a flat hierarchy and Calendars are rooted
// at the default
//nolint:revive
func (c CalendarDisplayable) GetParentFolderId() *string {
	return nil
}

// CreateCalendarDisplayable helper function to create the
// calendarDisplayable during msgraph-sdk-go iterative process
// @param entry is the input supplied by pageIterator.Iterate()
// @param parentID of Calendar sets. Only populate when used with
// EventCalendarCache
func CreateCalendarDisplayable(entry any) *CalendarDisplayable {
	calendar, ok := entry.(models.Calendarable)
	if !ok {
		return nil
	}

	return &CalendarDisplayable{
		Calendarable: calendar,
	}
}
