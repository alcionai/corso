package kopia

import (
	"encoding/base64"
	"path"
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

func decodeElements(elements ...string) []string {
	decoded := make([]string, 0, len(elements))

	for _, e := range elements {
		bs, err := encoder.DecodeString(e)
		if err != nil {
			decoded = append(decoded, "error decoding: "+e)
			continue
		}

		decoded = append(decoded, string(bs))
	}

	return decoded
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
