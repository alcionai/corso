// Package path provides a set of functions for wrangling paths from the outside
// world into paths that corso can understand. Paths use the standard Unix path
// separator character '/'. If for some reason an individual element in a raw
// path contains the '/' character, it should be escaped with '\'. If the path
// contains '\' it should be escaped by turning it into '\\'.
//
// Paths can be split into elements by splitting on '/' if the '/' is not
// escaped. Additionally, corso may operate on segments in a path. Segments are
// made up of one or more path elements.
//
// Examples of paths splitting by elements and canonicalization with escaping:
// 1.
//   input path: `this/is/a/path`
//   elements of path: `this`, `is`, `a`, `path`
// 2.
//   input path: `this/is\/a/path`
//   elements of path: `this`, `is/a`, `path`
// 3.
//   input path: `this/is\\/a/path`
//   elements of path: `this`, `is\`, `a`, `path`
// 4.
//   input path: `this/is\\\/a/path`
//   elements of path: `this`, `is\/a`, `path`
// 5.
//   input path: `this/is//a/path`
//   elements of path: `this`, `is`, `a`, `path`
// 6.
//   input path: `this/is\//a/path`
//   elements of path: `this`, `is/`, `a`, `path`
// 7.
//   input path: `this/is/a/path/`
//   elements of path: `this`, `is`, `a`, `path`
// 8.
//   input path: `this/is/a/path\/`
//   elements of path: `this`, `is`, `a`, `path/`
package path

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	escapeCharacter = '\\'
	pathSeparator   = '/'
)

var charactersToEscape = map[rune]struct{}{
	pathSeparator:   {},
	escapeCharacter: {},
}

// TODO(ashmrtn): Getting the category should either be through type-switches or
// through a function, but if it's a function it should re-use existing enums
// for resource types.
// For now, adding generic functions to pull information from segments.
// Resources that don't have the requested information should return an empty
// string.
type Path interface {
	String() string
	Tenant() string
	User() string
	Folder() string
	Item() string
}

// Builder is a simple path representation that only tracks path elements. It
// can join, escape, and unescape elements. Higher-level packages are expected
// to wrap this struct to build resource-speicific contexts (e.x. an
// ExchangeMailPath).
// Resource-specific paths allow access to more information like segments in the
// path. Builders that are turned into resource paths later on do not need to
// manually add prefixes for items that normally appear in the data layer (ex.
// tenant ID, service, user ID, etc).
type Builder struct {
	// Unescaped version of elements.
	elements []string
}

// UnescapeAndAppend creates a copy of this Builder and adds one or more already
// escaped path elements to the end of the new Builder. Elements are added in
// the order they are passed.
func (pb Builder) UnescapeAndAppend(elements ...string) (*Builder, error) {
	res := &Builder{elements: make([]string, 0, len(pb.elements))}
	copy(res.elements, pb.elements)

	if err := res.appendElements(true, elements); err != nil {
		return nil, err
	}

	return res, nil
}

// Append creates a copy of this Builder and adds the given elements them to the
// end of the new Builder. Elements are added in the order they are passed.
func (pb Builder) Append(elements ...string) *Builder {
	res := &Builder{elements: make([]string, len(pb.elements))}
	copy(res.elements, pb.elements)

	// Unescaped elements can't fail validation.
	//nolint:errcheck
	res.appendElements(false, elements)

	return res
}

func (pb *Builder) appendElements(escaped bool, elements []string) error {
	for _, e := range elements {
		if len(e) == 0 {
			continue
		}

		tmp := e

		if escaped {
			tmp = trimTrailingSlash(tmp)
			// If tmp was just the path separator then it will be empty now.
			if len(tmp) == 0 {
				continue
			}

			if err := validateEscapedElement(tmp); err != nil {
				return err
			}

			tmp = unescape(tmp)
		}

		pb.elements = append(pb.elements, tmp)
	}

	return nil
}

// String returns a string that contains all path elements joined together.
// Elements of the path that need escaping are escaped.
func (pb Builder) String() string {
	escaped := make([]string, 0, len(pb.elements))

	for _, e := range pb.elements {
		escaped = append(escaped, escapeElement(e))
	}

	return join(escaped)
}

//nolint:unused
func (pb Builder) join(start, end int) string {
	return join(pb.elements[start:end])
}

// escapeElement takes a single path element and escapes all characters that
// require an escape sequence. If there are no characters that need escaping,
// the input is returned unchanged.
func escapeElement(element string) string {
	escapeIdx := make([]int, 0)

	for i, c := range element {
		if _, ok := charactersToEscape[c]; ok {
			escapeIdx = append(escapeIdx, i)
		}
	}

	if len(escapeIdx) == 0 {
		return element
	}

	startIdx := 0
	b := strings.Builder{}
	b.Grow(len(element) + len(escapeIdx))

	for _, idx := range escapeIdx {
		b.WriteString(element[startIdx:idx])
		b.WriteRune(escapeCharacter)

		startIdx = idx
	}

	// Add the end of the element after the last escape character.
	b.WriteString(element[startIdx:])

	return b.String()
}

// unescape returns the given element and converts it into a "raw"
// element that does not have escape characters before characters that need
// escaping. Using this function on segments that contain escaped path
// separators will result in an ambiguous or incorrect segment.
func unescape(element string) string {
	b := strings.Builder{}
	startIdx := 0
	prevWasEscape := false

	for i, c := range element {
		if c != escapeCharacter || prevWasEscape {
			prevWasEscape = false
			continue
		}

		// This is an escape character, remove it from the output.
		b.WriteString(element[startIdx:i])
		startIdx = i + 1
		prevWasEscape = true
	}

	b.WriteString(element[startIdx:])

	return b.String()
}

// validateEscapedElement takes an escaped element that has had trailing
// separators trimmed and ensures that no characters requiring escaping are
// unescaped and that no escape characters are combined with characters that
// don't need escaping.
func validateEscapedElement(element string) error {
	prevWasEscape := false

	for _, c := range element {
		switch prevWasEscape {
		case true:
			prevWasEscape = false

			if _, ok := charactersToEscape[c]; !ok {
				return errors.Errorf(
					"bad escape sequence in path: '%c%c'", escapeCharacter, c)
			}

		case false:
			if c == escapeCharacter {
				prevWasEscape = true
				continue
			}

			if _, ok := charactersToEscape[c]; ok {
				return errors.Errorf("unescaped '%c' in path", c)
			}
		}
	}

	if prevWasEscape {
		return errors.New("trailing escape character")
	}

	return nil
}

// trimTrailingSlash takes an escaped path element and returns an escaped path
// element with the trailing path separator character(s) removed if they were not
// escaped. If there were no trailing path separator character(s) or the separator(s)
// were escaped the input is returned unchanged.
func trimTrailingSlash(element string) string {
	for len(element) > 0 && element[len(element)-1] == pathSeparator {
		lastIdx := len(element) - 1
		numSlashes := 0

		for i := lastIdx - 1; i >= 0; i-- {
			if element[i] != escapeCharacter {
				break
			}

			numSlashes++
		}

		if numSlashes%2 != 0 {
			break
		}

		element = element[:lastIdx]
	}

	return element
}

// join returns a string containing the given elements joined by the path
// separator '/'.
func join(elements []string) string {
	// Have to use strings because path package does not handle escaped '/' and
	// '\' according to the escaping rules.
	return strings.Join(elements, string(pathSeparator))
}
