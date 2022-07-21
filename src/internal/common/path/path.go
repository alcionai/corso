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

const (
	emailCategory = "email"
)

type Path struct {
}

// NewPath takes a path that is broken into segments and elements in the segment
// and returns a Path. Each element in the input is escaped.
func NewPath(segments [][]string) Path {
	return nil
}

// NewPathFromEscapedSegments takes already escaped segments of a path, verifies
// the segments are escaped properly, and returns a new Path struct. If there is
// an unescaped trailing '/' it is removed.
func NewPathFromEscapedSegments(segments []string) (Path, error) {
	return nil, errors.New("not implemented")
}

// String returns a string that contains all path segments joined
// together. Elements of the path that need escaping will be escaped.
func (p Path) String() string {
	return ""
}

// segment returns the nth segment of the path. Path segment indices are
// 0-based.
func (p Path) segment(n int) string {
	return ""
}

// TransformedSegments returns a slice of the Path segments where each segments
// has also been transformed such that it contains no characters outside the set
// of acceptable file system path characters.
func (p Path) TransformedSegments() []string {
	return nil
}

type ExchangeMailPath struct {
	*Path
}

// NewExchangeEmailPath creates and returns a new ExchangeEmailPath struct after
// verifying the path is properly escaped and contains information for the
// required segments.
func NewExchangeMailPath(
	tenant string,
	user string,
	folder []string,
	item string,
) (*ExchangeMailPath, error) {
	// TODO(ashmrtn): Verify required segments are not empty.

	p := NewPath([][]string{
		{tenant},
		{emailCategory},
		{user},
		folder,
		{item},
	})

	return &ExchangeMailPath{p}, nil
}

// Tenant returns the tenant ID for the referenced email resource.
func (emp ExchangeMailPath) Tenant() string {
	return emp.segment(0)
}

// Tenant returns an identifier noting this is a path for an email resource.
func (emp ExchangeMailPath) Category() string {
	return emp.segment(1)
}

// Tenant returns the user ID for the referenced email resource.
func (emp ExchangeMailPath) User() string {
	return emp.segment(2)
}

// Tenant returns the folder segment for the referenced email resource.
func (emp ExchangeMailPath) Folder() string {
	return emp.segment(3)
}

// Tenant returns the email ID for the referenced email resource.
func (emp ExchangeMailPath) Mail() string {
	return emp.segment(4)
}
