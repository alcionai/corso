package support

import (
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type attendee struct {
	name     string
	email    string
	response string
}

// FormatAttendeses
// - First Name <email@example.com>, Accepted | Declined | Tentative | No Response
func FormatAttendees(event models.Eventable) string {
	var (
		failed   int
		response = event.GetAttendees()
		required = make([]attendee, 0)
		optional = make([]attendee, 0)
		resource = make([]attendee, 0)
	)

	for _, entry := range response {
		if guardCheckForAttendee(entry) {
			failed++
			continue
		}

		temp := attendee{
			name:     *entry.GetEmailAddress().GetName(),
			email:    *entry.GetEmailAddress().GetAddress(),
			response: entry.GetStatus().GetResponse().String(),
		}

		switch *entry.GetType() {
		case models.REQUIRED_ATTENDEETYPE:
			required = append(required, temp)

		case models.OPTIONAL_ATTENDEETYPE:
			optional = append(optional, temp)

		case models.RESOURCE_ATTENDEETYPE:
			resource = append(resource, temp)
		}
	}

	req := attendeeListToString(required, "Required")
	opt := attendeeListToString(optional, "Optional")
	res := attendeeListToString(resource, "Resource")

	return req + opt + res
}

func attendeeListToString(attendList []attendee, heading string) string {
	var message string
	if len(attendList) > 0 {
		message = heading + ":\n"
		for _, resource := range attendList {
			message += "- " + resource.simplePrint() + "\n"
		}

		message += "\n\n"
	}

	return message
}

func guardCheckForAttendee(attendee models.Attendeeable) bool {
	if attendee.GetType() == nil {
		return true
	}

	if attendee.GetStatus() == nil {
		return true
	}

	if attendee.GetStatus().GetResponse() == nil {
		return true
	}

	if attendee.GetEmailAddress() == nil {
		return true
	}

	if attendee.GetEmailAddress().GetName() == nil ||
		attendee.GetEmailAddress().GetAddress() == nil {
		return true
	}

	return false
}

func (at *attendee) simplePrint() string {
	contents := fmt.Sprintf("%s <%s>, %s", at.name, at.email, at.response)
	return contents
}
