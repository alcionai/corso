package sanitize

import (
	"bytes"
	"fmt"

	"golang.org/x/exp/slices"
)

// JSONString takes a []byte containing JSON as input and returns a []byte
// containing the same content but with any character codes < 0x20 that weren't
// escaped in the original escaped properly.
func JSONBytes(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	// Avoid most reallocations by just getting a buffer of the right size to
	// start with.
	// TODO(ashmrtn): We may actually want to overshoot this a little so we won't
	// cause a reallocation and possible doubling in size if we only need to
	// escape a few characters.
	buf := bytes.Buffer{}
	buf.Grow(len(input))

	for _, c := range input {
		switch {
		case c == '\n' || c == '\t' || c == '\r':
			// Whitespace characters also seem to be getting transformed inside JSON
			// strings already. We shouldn't further transform them because they could
			// just be formatting around the JSON fields so changing them will result
			// in invalid JSON.
			//
			// The set of whitespace characters was taken from RFC 8259 although space
			// is not included in this set as it's already > 0x20.
			buf.WriteByte(c)

		case c < 0x20:
			// Escape character ranges taken from RFC 8259. This case doesn't handle
			// escape characters (0x5c) or double quotes (0x22). We're assuming escape
			// characters don't require additional processing and that double quotes
			// are properly escaped by whatever handed us the JSON.
			//
			// We need to escape the character and transform it (e.x. linefeed -> \n).
			// We could use transforms like linefeed to \n, but it's actually easier,
			// if a little less space efficient, to just turn them into
			// multi-character sequences denoting a unicode character.
			buf.WriteString(fmt.Sprintf(`\u%04X`, c))

		default:
			buf.WriteByte(c)
		}
	}

	// Return a copy just so we don't hold a reference to internal bytes.Buffer
	// data.
	return slices.Clone(buf.Bytes())
}
