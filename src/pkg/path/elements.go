package path

import (
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/pii"
)

var piiSafePathElems = pii.MapWithPlurals(
	// services
	UnknownService.String(),
	ExchangeService.String(),
	OneDriveService.String(),
	SharePointService.String(),
	ExchangeMetadataService.String(),
	OneDriveMetadataService.String(),
	SharePointMetadataService.String(),

	// categories
	UnknownCategory.String(),
	EmailCategory.String(),
	ContactsCategory.String(),
	EventsCategory.String(),
	FilesCategory.String(),
	ListsCategory.String(),
	LibrariesCategory.String(),
	PagesCategory.String(),
	DetailsCategory.String(),
)

var (
	// interface compliance required for handling PII
	_ clues.Concealer = &Elements{}
	_ fmt.Stringer    = &Elements{}

	// interface compliance for the observe package to display
	// values without concealing PII.
	_ clues.PlainStringer = &Elements{}
)

// Elements are a PII Concealer-compliant slice of elements within a path.
type Elements []string

// NewElements creates a new Elements slice by splitting the provided string.
func NewElements(p string) Elements {
	return Split(p)
}

// Conceal produces a concealed representation of the elements, suitable for
// logging, storing in errors, and other output.
func (el Elements) Conceal() string {
	escaped := make([]string, 0, len(el))

	for _, e := range el {
		escaped = append(escaped, escapeElement(e))
	}

	return join(pii.ConcealElements(escaped, piiSafePathElems))
}

// Format produces a concealed representation of the elements, even when
// used within a PrintF, suitable for logging, storing in errors,
// and other output.
func (el Elements) Format(fs fmt.State, _ rune) {
	fmt.Fprint(fs, el.Conceal())
}

// String returns a string that contains all path elements joined together.
// Elements that need escaping are escaped.  The result is not concealed, and
// is not suitable for logging or structured errors.
func (el Elements) String() string {
	escaped := make([]string, 0, len(el))

	for _, e := range el {
		escaped = append(escaped, escapeElement(e))
	}

	return join(escaped)
}

// PlainString returns an unescaped, unmodified string of the joined elements.
// The result is not concealed, and is not suitable for logging or structured
// errors.
func (el Elements) PlainString() string {
	return join(el)
}
