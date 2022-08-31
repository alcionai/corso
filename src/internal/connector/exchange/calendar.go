package exchange

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type calendarDisplayable struct {
	models.Calendarable
}

func (c calendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

func CreateCalendarDisplayable(entry any) *calendarDisplayable {
	calendar, ok := entry.(models.Calendarable)
	if !ok {
		return nil
	}

	return &calendarDisplayable{calendar}
}
