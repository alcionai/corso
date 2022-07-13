// Package path provides a set of functions for wrangling paths from the outside
// world into paths that corso can understand. Paths use the standard Unix path
// separator character '/'. If for some reason an individual element in a raw
// path contains the '/' character, it should be escaped with '\'. If the path
// contains '\' it should be escaped by turning it into '\\'.
//
// Examples:
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
//
// Internally, corso segments paths into distinct segments:
// * tenant and service
// * user, if applicable
// * subdirectories, if applicable
// * individual item if not a directory, if applicable
//
// Segments that are not required for a particular path are not exposed in any
// of the functions that return path information.

package path

import (
	"errors"
)

type Path struct {
}

// NewPath takes a directory path and an individual item name, validates escape
// characters are used properly, and returns a Path struct. Any segments of the
// path that are empty are ignored in output functions.
func NewPath(tenantAndService, user, dirString, item string) (*Path, error) {
	return nil, errors.New("not implemented")
}

// String returns a string that contains all path segments joined
// together. Elements of the path that need escaping will be escaped.
func (p *Path) String() string {
	return ""
}

// HashedSegments returns a slice of path segments that aligns with the path
// segmentation corso uses internally. This means that number of segments
// returned by this function may be less than the number of segments that would
// be returned by `strings.Split(path, '/')`. Each segment is a hash of the
// underlying data of the segment. This function can be used to obtain path
// segments that are safe for external libraries that may not accept special
// characters or non-ascii characters.
func (p *Path) HashedSegments() []string {
	return nil
}

// Elements returns all elements in the path according to path splitting rules
// that account for escaped characters. See the examples above.
func (p *Path) Elements() []string {
	return nil
}
