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
	"errors"
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
}

// newPath takes a path that is broken into segments and elements in the segment
// and returns a Base. Each element in the input is escaped.
func newPath(segments [][]string) Base {
	return Base{}
}

// NewPathFromEscapedSegments takes already escaped segments of a path, verifies
// the segments are escaped properly, and returns a new Base struct. If there is
// an unescaped trailing '/' it is removed.
func newPathFromEscapedSegments(segments []string) (Base, error) {
	return Base{}, errors.New("not implemented")
}

// String returns a string that contains all path segments joined
// together. Elements of the path that need escaping will be escaped.
func (p Base) String() string {
	return ""
}

// segment returns the nth segment of the path. Path segment indices are
// 0-based.
func (p Base) segment(n int) string {
	return ""
}

// unescapedSegmentElements returns the unescaped version of the elements that
// comprise the requested segment.
func (p Base) unescapedSegmentElements(n int) []string {
	return nil
}

// TransformedSegments returns a slice of the path segments where each segments
// has also been transformed such that it contains no characters outside the set
// of acceptable file system path characters.
func (p Base) TransformedSegments() []string {
	return nil
}
