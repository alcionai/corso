package exchange

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// calendarDisplayable is a transformative struct that aligns
// models.Calendarable interface with the displayable interface.
type calendarDisplayable struct {
	models.Calendarable
}

// GetDisplayName returns the *string of the calendar name
func (c calendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

// CreateCalendarDisplayable helper function to create the
// calendarDisplayable during msgraph-sdk-go iterative process
// @param entry is the input supplied by pageIterator.Iterate()
func CreateCalendarDisplayable(entry any) *calendarDisplayable {
	calendar, ok := entry.(models.Calendarable)
	if !ok {
		return nil
	}

	return &calendarDisplayable{calendar}
}
