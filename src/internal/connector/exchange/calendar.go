package exchange

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CalendarDisplayable is a transformative struct that aligns
// models.Calendarable interface with the Displayable interface.
type CalendarDisplayable struct {
	models.Calendarable
}

// GetDisplayName returns the *string of the calendar name
func (c CalendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

// CreateCalendarDisplayable helper function to create the
// calendarDisplayable during msgraph-sdk-go iterative process
// @param entry is the input supplied by pageIterator.Iterate()
func CreateCalendarDisplayable(entry any) *CalendarDisplayable {
	calendar, ok := entry.(models.Calendarable)
	if !ok {
		return nil
	}

	return &CalendarDisplayable{calendar}
}
