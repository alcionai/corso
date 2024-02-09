package kopia

import (
	"encoding/base64"
	"path"
	"strings"

	"github.com/alcionai/clues"
)

var encoder = base64.URLEncoding

// encodeElements takes a set of strings and returns a slice of the strings
// after encoding them to a file system-safe format. Elements are returned in
// the same order they were passed in.
func encodeElements(elements ...string) []string {
	encoded := make([]string, 0, len(elements))

	for _, e := range elements {
		encoded = append(encoded, encoder.EncodeToString([]byte(e)))
	}

	return encoded
}

// decodePath splits inputPath on the path separator and returns the base64
// decoding of each element in the path. If an error occurs then returns a mixed
// set of encoded and decoded elements and an error with information about each
// element that failed decoding.
func decodePath(inputPath string) ([]string, error) {
	res, err := decodeElements(strings.Split(inputPath, "/")...)
	return res, clues.Stack(err).OrNil()
}

// decodeElements returns the base64 decoding of each input element. If any
// element fails to decode it returns a mix of encoded (failed) decoded elements
// and an error.
func decodeElements(elements ...string) ([]string, error) {
	var (
		errs    *clues.Err
		decoded = make([]string, 0, len(elements))
	)

	for _, e := range elements {
		decodedBytes, err := encoder.DecodeString(e)
		// Make an additional string variable so we can just assign to it if there
		// was an error. This avoids a continue in the error check below.
		decodedElement := string(decodedBytes)

		if err != nil {
			errs = clues.Stack(
				errs,
				clues.Wrap(err, "decoding element").With("element", e))
			// Set bs to the input value so it gets returned in its encoded form.
			decodedElement = e
		}

		decoded = append(decoded, decodedElement)
	}

	return decoded, errs.OrNil()
}

// encodeAsPath takes a set of elements and returns the concatenated elements as
// if they were a path. The elements are joined with the separator in the golang
// path package.
func encodeAsPath(elements ...string) string {
	// Needs `/` to be used a separator here
	//nolint:forbidigo
	return path.Join(encodeElements(elements...)...)
}

// decodeElement takes an encoded element and decodes it if possible.
func decodeElement(element string) (string, error) {
	r, err := encoder.DecodeString(element)
	return string(r), err
}
