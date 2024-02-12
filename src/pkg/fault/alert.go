package fault

import (
	"github.com/alcionai/corso/src/cli/print"
)

const (
	AlertPreviousPathCollision = "previous_path_collision"
)

type alertType string

const (
	AlertTypeUnknown           = ""
	AlertTypeUserVisible       = "user_visible"
	AlertTypeInternalTelemetry = "internal_telemetry"
)

var _ print.Printable = &Alert{}

// Alerts are informational-only notifications.  The purpose of alerts is to
// provide a means of end-user communication about important events without
// needing to generate runtime failures or recoverable errors. When generating
// an alert, no other fault feature (failure, recoverable, skip, etc) should
// be in use.  IE: Errors do not also get alerts, since the error itself is a
// form of end-user communication already.
type Alert struct {
	Type    alertType `json:"type"`
	Item    Item      `json:"item"`
	Message string    `json:"message"`
}

// String complies with the stringer interface.
func (a *Alert) String() string {
	msg := "<nil>"

	if a != nil {
		msg = a.Message
	}

	if len(msg) == 0 {
		msg = "<missing>"
	}

	return "Alert: " + msg
}

func (a Alert) MinimumPrintable() any {
	return a
}

// Headers returns the human-readable names of properties of a skipped Item
// for printing out to a terminal.
func (a Alert) Headers(bool) []string {
	// NOTE: skipID does not make sense in this context and is skipped
	return []string{"Action", "Message", "Container", "Name", "ID"}
}

// Values populates the printable values matching the Headers list.
func (a Alert) Values(bool) []string {
	var cn string

	acn, ok := a.Item.Additional[AddtlContainerName]
	if ok {
		str, ok := acn.(string)
		if ok {
			cn = str
		}
	}

	return []string{"Alert", a.Message, cn, a.Item.Name, a.Item.ID}
}

// NewAlert produces a USER VISIBLE alert.  Use this if you want the end user
// to see the alert at the end of the run.
func NewAlert(
	message, namespace, entityID, name string,
	addtl map[string]any,
) *Alert {
	return &Alert{
		Type:    AlertTypeUserVisible,
		Message: message,
		Item: Item{
			Namespace:  namespace,
			ID:         entityID,
			Name:       name,
			Additional: addtl,
		},
	}
}

// NewTelemetryAlert produces a NON-USER VISIBLE alert.  Use this if you want
// to create an alert event for additional tracking and telemetry without the
// user seeing the details at the end of the run.  Note that this data is still
// accessible from the backup metadata, so it will be possible for users to
// see it, if wanted/necessary.  But they do have to go looking for it specifically.
func NewTelemetryAlert(
	message, namespace, entityID, name string,
	addtl map[string]any,
) *Alert {
	return &Alert{
		Type:    AlertTypeInternalTelemetry,
		Message: message,
		Item: Item{
			Namespace:  namespace,
			ID:         entityID,
			Name:       name,
			Additional: addtl,
		},
	}
}
