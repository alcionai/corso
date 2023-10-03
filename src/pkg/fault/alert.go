package fault

import (
	"context"

	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/pkg/logger"
)

// Alerts returns the slice of alerts generated during runtime.
// If the bus is a local alerts, this only returns the
// local failure, and will not return parent data.
func (e *Bus) Alerts() []Alert {
	return slices.Clone(e.alerts)
}

// AddAlert appends a record of an Alert message to the fault bus.
// Importantly, alerts are not errors, exceptions, or skipped items.
// An alert should only be generated if no other fault functionality
// is in use, but that we still want the end user to clearly and
// plainly receive a notification about a runtime event.
func (e *Bus) AddAlert(ctx context.Context, a *Alert) {
	if a == nil {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.logAndAddAlert(ctx, a, 1)
}

// logs the error and adds an alert.
func (e *Bus) logAndAddAlert(ctx context.Context, a *Alert, trace int) {
	logger.CtxStack(ctx, trace+1).
		With("skipped", a).
		Info("recoverable error")
	e.addAlert(a)
}

func (e *Bus) addAlert(a *Alert) *Bus {
	e.alerts = append(e.alerts, *a)

	// local bus instances must promote alerts to the root bus.
	if e.parent != nil {
		e.parent.addAlert(a)
	}

	return e
}

var _ print.Printable = &Alert{}

// Alerts are informational-only notifications.  The purpose of alerts is to
// provide a means of end-user communication about important events without
// needing to generate runtime failures or recoverable errors. When generating
// an alert, no other fault feature (failure, recoverable, skip, etc) should
// be in use.  IE: Errors do not also get alerts, since the error itself is a
// form of end-user communication already.
type Alert struct {
	Item    Item   `json:"item"`
	Message string `json:"message"`
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
func (a Alert) Headers() []string {
	return []string{"Action", "Message", "Container", "Name", "ID"}
}

// Values populates the printable values matching the Headers list.
func (a Alert) Values() []string {
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

func NewAlert(message, namespace, itemID, name string, addtl map[string]any) *Alert {
	return &Alert{
		Message: message,
		Item: Item{
			Namespace:  namespace,
			ID:         itemID,
			Name:       name,
			Additional: addtl,
		},
	}
}
