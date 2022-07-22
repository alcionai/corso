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

var (
	charactersToEscape = map[rune]struct{}{
		pathSeparator:   {},
		escapeCharacter: {},
	}
)

var errMissingSegment = errors.New("missing required path segment")

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

type Base struct {
	// Escaped path elements.
	elements []string
	// Contains starting index in elements of each segment.
	segmentIdx []int
}

// newPath takes a path that is broken into segments and elements in the segment
// and returns a Base. Each element in the input is escaped.
func newPath(segments [][]string) Base {
	if len(segments) == 0 {
		return Base{}
	}

	res := Base{segmentIdx: make([]int, 0, len(segments))}
	idx := 0
	for _, s := range segments {
		res.segmentIdx = append(res.segmentIdx, idx)

		for _, e := range s {
			if len(e) == 0 {
				continue
			}

			res.elements = append(res.elements, escapeElement(e))
			idx++
		}
	}

	return res
}

// NewPathFromEscapedSegments takes already escaped segments of a path, verifies
// the segments are escaped properly, and returns a new Base struct. If there is
// an unescaped trailing '/' it is removed.
func newPathFromEscapedSegments(segments []string) (Base, error) {
	b := Base{}

	if err := validateSegments(segments); err != nil {
		return b, errors.Wrap(err, "validating escaped path")
	}

	// Make a copy of the input so we don't modify the original slice.
	tmpSegments := make([]string, len(segments))
	copy(tmpSegments, segments)
	tmpSegments[len(tmpSegments)-1] = trimTrailingSlash(tmpSegments[len(tmpSegments)-1])

	for _, s := range tmpSegments {
		newElems := split(s)

		if len(newElems) == 0 {
			continue
		}

		b.segmentIdx = append(b.segmentIdx, len(b.elements))
		b.elements = append(b.elements, newElems...)
	}

	return b, nil
}

// String returns a string that contains all path segments joined
// together. Elements of the path that need escaping will be escaped.
func (b Base) String() string {
	return join(b.elements)
}

// segment returns the nth segment of the path. Path segment indices are
// 0-based. As this function is used exclusively by wrappers of path, it does no
// bounds checking. Callers are expected to have validated the number of
// segments when making the path.
func (b Base) segment(n int) string {
	if n == len(b.segmentIdx)-1 {
		return join(b.elements[b.segmentIdx[n]:])
	}

	return join(b.elements[b.segmentIdx[n]:b.segmentIdx[n+1]])
}

// unescapedSegmentElements returns the unescaped version of the elements that
// comprise the requested segment.
func (p Base) unescapedSegmentElements(n int) []string {
	return nil
}

// TransformedSegments returns a slice of the path segments where each segments
// has also been transformed such that it contains no characters outside the set
// of acceptable file system path characters.
func (b Base) TransformedSegments() []string {
	return nil
}

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

	b := strings.Builder{}
	b.Grow(len(element) + len(escapeIdx))
	startIdx := 0

	for _, idx := range escapeIdx {
		b.WriteString(element[startIdx:idx])
		b.WriteRune(escapeCharacter)
		startIdx = idx
	}

	// Add the end of the element after the last escape character.
	b.WriteString(element[startIdx:])

	return b.String()
}

func validateSegments(segments []string) error {
	for _, segment := range segments {
		prevWasEscape := false

		for _, c := range segment {
			if c == escapeCharacter {
				// Either the character before this was also an escape character in which
				// case we're no longer escaping things (so prevWasEscape should become
				// false) or the previous character was not an escape character and now
				// we're escaping things. In either case we should invert prevWasEscape.
				prevWasEscape = !prevWasEscape
			} else {
				if prevWasEscape {
					if _, ok := charactersToEscape[c]; !ok {
						return errors.Errorf(
							"bad escape sequence in path: '%c%c'", escapeCharacter, c)
					}
				}

				prevWasEscape = false
			}
		}

		if prevWasEscape {
			return errors.New("trailing escape character")
		}
	}

	return nil
}

// trimTrailingSlash takes an escaped path element and returns an escaped path
// element with the trailing path separator character removed if it was not
// escaped. If there was no trailing path separator character or the separator
// was escaped the input is returned unchanged.
func trimTrailingSlash(element string) string {
	lastIdx := len(element) - 1

	if element[lastIdx] != pathSeparator {
		return element
	}

	numSlashes := 0
	for i := lastIdx - 1; i >= 0; i-- {
		if element[i] != escapeCharacter {
			break
		}

		numSlashes++
	}

	if numSlashes%2 != 0 {
		return element
	}

	return element[:lastIdx]
}

// join returns a string containing the given elements joined by the path
// separator '/'.
func join(elements []string) string {
	// Have to use strings because path package does not handle escaped '/' and
	// '\' according to the escaping rules.
	return strings.Join(elements, string(pathSeparator))
}

// split returns a slice of path elements for the given segment when the segment
// is split on the path separator according to the escaping rules.
func split(segment string) []string {
	res := make([]string, 0)
	numEscapes := 0
	startIdx := 0
	// Start with true to ignore leading separator.
	prevWasSeparator := true

	for i, c := range segment {
		if c == escapeCharacter {
			numEscapes++
			prevWasSeparator = false
			continue
		}

		if c != pathSeparator {
			prevWasSeparator = false
			numEscapes = 0
			continue
		}

		// Remaining is just path separator handling.
		if numEscapes%2 != 0 {
			// This is an escaped separator.
			prevWasSeparator = false
			numEscapes = 0
			continue
		}

		// Ignore leading separator characters and don't add elements that would
		// be empty.
		if !prevWasSeparator {
			res = append(res, segment[startIdx:i])
		}

		// We don't want to include the path separator in the result.
		startIdx = i + 1
		prevWasSeparator = true
		numEscapes = 0
	}

	// Add the final segment because the loop above won't catch it. There should
	// be no trailing separator character, but do a bounds check to be safe.
	res = append(res, segment[startIdx:])

	return res
}
