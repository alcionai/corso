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
func FormatAttendees(event models.Eventable, isHTML bool) string {
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

	req := attendeeListToString(required, "Required", isHTML)
	opt := attendeeListToString(optional, "Optional", isHTML)
	res := attendeeListToString(resource, "Resource", isHTML)

	return req + opt + res
}

func attendeeListToString(attendList []attendee, heading string, isHTML bool) string {
	var (
		message        string
		carriageReturn string
	)

	if isHTML {
		carriageReturn = "<br>"
	} else {
		carriageReturn = "\n"
	}

	if len(attendList) > 0 {
		message = heading + ":" + carriageReturn
		for _, resource := range attendList {
			message += "- " + resource.simplePrint(isHTML) + " " + carriageReturn
		}

		message += carriageReturn + carriageReturn
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

func (at *attendee) simplePrint(isHTML bool) string {
	var contents string
	if isHTML {
		contents = fmt.Sprintf("%s &lt;%s&gt;, %s", at.name, at.email, at.response)
	} else {
		contents = fmt.Sprintf("%s <%s>, %s", at.name, at.email, at.response)
	}

	return contents
}
