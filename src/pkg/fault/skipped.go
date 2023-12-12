package fault

import (
	"context"

	"github.com/alcionai/corso/src/cli/print"
)

// AddSkipper presents an interface that allows callers to
// write additional skipped items to the complying struct.
type AddSkipper interface {
	AddSkip(ctx context.Context, s *Skipped)
}

// skipCause identifies the well-known conditions to Skip an item.  It is
// important that skip cause enumerations do not overlap with general error
// handling.  Skips must be well known, well documented, and consistent.
// Transient failures, undocumented or unknown conditions, and arbitrary
// handling should never produce a skipped item. Those cases should get
// handled as normal errors.
type skipCause string

const (
	// SkipMalware identifies a malware detection case.  Files that graph
	// api identifies as malware cannot be downloaded or uploaded, and will
	// permanently fail any attempts to backup or restore.
	SkipMalware skipCause = "malware_detected"

	// SkipOneNote identifies that a file was skipped because it
	// was a OneNote file that remains inaccessible (503 server response)
	// regardless of the number of retries.
	//nolint:lll
	// https://support.microsoft.com/en-us/office/restrictions-and-limitations-in-onedrive-and-sharepoint-64883a5d-228e-48f5-b3d2-eb39e07630fa#onenotenotebooks
	SkipOneNote skipCause = "inaccessible_one_note_file"

	// SkipInvalidRecipients identifies that an email was skipped because Exchange
	// believes it is not valid and fails any attempt to read it.
	SkipInvalidRecipients skipCause = "invalid_recipients_email"
)

var _ print.Printable = &Skipped{}

// Skipped items are permanently unprocessable due to well-known conditions.
// In order to skip an item, the following conditions should be met:
// 1. The conditions for skipping the item are well-known and
// well-documented.  End users need to be able to understand
// both the conditions and identifications of skips.
// 2. Skipping avoids a permanent and consistent failure.  If
// the underlying reason is transient or otherwise recoverable,
// the item should not be skipped.
//
// Skipped wraps Item primarily to minimize confusion when sharing the
// fault interface.  Skipped items are not errors, and Item{} errors are
// not the basis for a Skip.
type Skipped struct {
	Item Item `json:"item"`
}

// String complies with the stringer interface.
func (s *Skipped) String() string {
	if s == nil {
		return "<nil>"
	}

	return "skipped " + s.Item.Error() + ": " + s.Item.Cause
}

// HasCause compares the underlying cause against the parameter.
func (s *Skipped) HasCause(c skipCause) bool {
	if s == nil {
		return false
	}

	return s.Item.Cause == string(c)
}

func (s Skipped) MinimumPrintable() any {
	return s
}

// Headers returns the human-readable names of properties of a skipped Item
// for printing out to a terminal.
func (s Skipped) Headers(bool) []string {
	// NOTE: skipID does not make sense in this context
	return []string{"Action", "Type", "Name", "Container", "Cause"}
}

// Values populates the printable values matching the Headers list.
func (s Skipped) Values(bool) []string {
	var cn string

	acn, ok := s.Item.Additional[AddtlContainerName]
	if ok {
		str, ok := acn.(string)
		if ok {
			cn = str
		}
	}

	return []string{"Skip", s.Item.Type.Printable(), s.Item.Name, cn, s.Item.Cause}
}

// ContainerSkip produces a Container-kind Item for tracking skipped items.
func ContainerSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(ContainerType, cause, namespace, id, name, addtl)
}

// EmailSkip produces a Email-kind Item for tracking skipped items.
func EmailSkip(cause skipCause, user, id string, addtl map[string]any) *Skipped {
	return itemSkip(EmailType, cause, user, id, "", addtl)
}

// FileSkip produces a File-kind Item for tracking skipped items.
func FileSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(FileType, cause, namespace, id, name, addtl)
}

// OnwerSkip produces a ResourceOwner-kind Item for tracking skipped items.
func OwnerSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(ResourceOwnerType, cause, namespace, id, name, addtl)
}

// itemSkip produces a Item of the provided type for tracking skipped items.
func itemSkip(t ItemType, cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return &Skipped{
		Item: Item{
			Namespace:  namespace,
			ID:         id,
			Name:       name,
			Type:       t,
			Cause:      string(cause),
			Additional: addtl,
		},
	}
}
