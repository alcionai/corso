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
	return Base{}, errors.New("not implemented")
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

// join returns a string containing the given elements joined by the path
// separator '/'.
func join(elements []string) string {
	// Have to use strings because path package does not handle escaped '/' and
	// '\' according to the escaping rules.
	return strings.Join(elements, string(pathSeparator))
}
