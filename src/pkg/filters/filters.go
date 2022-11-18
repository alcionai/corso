package filters

import (
	"strings"

	"github.com/alcionai/corso/src/pkg/path"
)

type comparator int

const (
	UnknownComparator comparator = iota
	// a == b
	EqualTo
	// a > b
	GreaterThan
	// a < b
	LessThan
	// "foo,bar,baz" contains "foo"
	TargetContains
	// "foo" is found in "foo,bar,baz"
	TargetIn
	// always passes
	Passes
	// always fails
	Fails
	// passthrough for the target
	IdentityValue
	// "foo" is a prefix of "foobarbaz"
	TargetPrefixes
	// "foo" equals any complete element prefix of "foo/bar/baz"
	TargetPathPrefix
	// "foo" equals any complete element in "foo/bar/baz"
	TargetPathContains
)

func norm(s string) string {
	return strings.ToLower(s)
}

// normPathElem ensures the string is:
// 1. prefixed with a single path.pathSeparator (ex: `/`)
// 2. suffixed with a single path.pathSeparator (ex: `/`)
// This is done to facilitate future regex comparisons
// without re-running the prefix-suffix addition multiple
// times per target.
func normPathElem(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[0] != path.PathSeparator {
		s = string(path.PathSeparator) + s
	}

	s = path.TrimTrailingSlash(s) + string(path.PathSeparator)

	return s
}

// Filter contains a comparator func and the target to
// compare values against.  Filter.Matches(v) returns
// true if Filter.Comparer(filter.target, v) is true.
type Filter struct {
	Comparator        comparator `json:"comparator"`
	Target            string     `json:"target"`            // the value to compare against
	Targets           []string   `json:"targets"`           // the set of values to compare
	NormalizedTargets []string   `json:"normalizedTargets"` // the set of comparable values post normalization
	Negate            bool       `json:"negate"`            // when true, negate the comparator result
}

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

// Equal creates a filter where Compare(v) is true if
// target == v
func Equal(target string) Filter {
	return newFilter(EqualTo, target, false)
}

// NotEqual creates a filter where Compare(v) is true if
// target != v
func NotEqual(target string) Filter {
	return newFilter(EqualTo, target, true)
}

// Greater creates a filter where Compare(v) is true if
// target > v
func Greater(target string) Filter {
	return newFilter(GreaterThan, target, false)
}

// NotGreater creates a filter where Compare(v) is true if
// target <= v
func NotGreater(target string) Filter {
	return newFilter(GreaterThan, target, true)
}

// Less creates a filter where Compare(v) is true if
// target < v
func Less(target string) Filter {
	return newFilter(LessThan, target, false)
}

// NotLess creates a filter where Compare(v) is true if
// target >= v
func NotLess(target string) Filter {
	return newFilter(LessThan, target, true)
}

// Contains creates a filter where Compare(v) is true if
// target.Contains(v)
func Contains(target string) Filter {
	return newFilter(TargetContains, target, false)
}

// NotContains creates a filter where Compare(v) is true if
// !target.Contains(v)
func NotContains(target string) Filter {
	return newFilter(TargetContains, target, true)
}

// In creates a filter where Compare(v) is true if
// v.Contains(target)
func In(target string) Filter {
	return newFilter(TargetIn, target, false)
}

// NotIn creates a filter where Compare(v) is true if
// !v.Contains(target)
func NotIn(target string) Filter {
	return newFilter(TargetIn, target, true)
}

// Pass creates a filter where Compare(v) always returns true
func Pass() Filter {
	return newFilter(Passes, "*", false)
}

// Fail creates a filter where Compare(v) always returns false
func Fail() Filter {
	return newFilter(Fails, "", false)
}

// Identity creates a filter intended to hold values, rather than
// compare them.  Comparatively, it'll behave the same as Equals.
func Identity(id string) Filter {
	return newFilter(IdentityValue, id, false)
}

// Prefix creates a filter where Compare(v) is true if
// target.Prefix(v)
func Prefix(target string) Filter {
	return newFilter(TargetPrefixes, target, false)
}

// NotPrefix creates a filter where Compare(v) is true if
// !target.Prefix(v)
func NotPrefix(target string) Filter {
	return newFilter(TargetPrefixes, target, true)
}

// PathPrefix creates a filter where Compare(v) is true if
// target.Prefix(v) &&
// split(target)[i].Equals(split(v)[i]) for _all_ i in 0..len(target)-1
// ex: target "/foo/bar" returns true for input "/foo/bar/baz",
// but false for "/foo/barbaz"
//
// Unlike single-target filters, this filter accepts a
// slice of targets, will compare an input against each target
// independently, and returns true if one or more of the
// comparisons succeed.
func PathPrefix(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newSliceFilter(TargetPathPrefix, targets, tgts, false)
}

// NotPathPrefix creates a filter where Compare(v) is true if
// !target.Prefix(v) ||
// !split(target)[i].Equals(split(v)[i]) for _any_ i in 0..len(target)-1
// ex: target "/foo/bar" returns false for input "/foo/bar/baz",
// but true for "/foo/barbaz"
//
// Unlike single-target filters, this filter accepts a
// slice of targets, will compare an input against each target
// independently, and returns true if one or more of the
// comparisons succeed.
func NotPathPrefix(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newSliceFilter(TargetPathPrefix, targets, tgts, true)
}

// PathContains creates a filter where Compare(v) is true if
// for _any_ elem e in split(v), target.Equals(e) ||
// for _any_ sequence of elems in split(v), target.Equals(path.Join(e[n:m]))
// ex: target "foo" returns true for input "/baz/foo/bar",
// but false for "/baz/foobar"
// ex: target "baz/foo" returns true for input "/baz/foo/bar",
// but false for "/baz/foobar"
//
// Unlike single-target filters, this filter accepts a
// slice of targets, will compare an input against each target
// independently, and returns true if one or more of the
// comparisons succeed.
func PathContains(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newSliceFilter(TargetPathContains, targets, tgts, false)
}

// NotPathContains creates a filter where Compare(v) is true if
// for _every_ elem e in split(v), !target.Equals(e) ||
// for _every_ sequence of elems in split(v), !target.Equals(path.Join(e[n:m]))
// ex: target "foo" returns false for input "/baz/foo/bar",
// but true for "/baz/foobar"
// ex: target "baz/foo" returns false for input "/baz/foo/bar",
// but true for "/baz/foobar"
//
// Unlike single-target filters, this filter accepts a
// slice of targets, will compare an input against each target
// independently, and returns true if one or more of the
// comparisons succeed.
func NotPathContains(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newSliceFilter(TargetPathContains, targets, tgts, true)
}

// newFilter is the standard filter constructor.
func newFilter(c comparator, target string, negate bool) Filter {
	return Filter{
		Comparator: c,
		Target:     target,
		Negate:     negate,
	}
}

// newSliceFilter constructs filters that contain multiple targets
func newSliceFilter(c comparator, targets, normTargets []string, negate bool) Filter {
	return Filter{
		Comparator:        c,
		Targets:           targets,
		NormalizedTargets: normTargets,
		Negate:            negate,
	}
}

// ----------------------------------------------------------------------------------------------------
// Comparisons
// ----------------------------------------------------------------------------------------------------

// CompareAny checks whether any one of all the provided
// inputs passes the filter.
//
// Note that, as a gotcha, CompareAny can resolve truthily
// for both the standard and negated versions of a filter.
// Ex: consider the input CompareAny(true, false), which
// will return true for both Equals(true) and NotEquals(true),
// because at least one element matches for both filters.
func (f Filter) CompareAny(inputs ...string) bool {
	for _, in := range inputs {
		if f.Compare(in) {
			return true
		}
	}

	return false
}

// Compare checks whether the input passes the filter.
func (f Filter) Compare(input string) bool {
	var (
		cmp      func(string, string) bool
		hasSlice bool
	)

	switch f.Comparator {
	case EqualTo, IdentityValue:
		cmp = equals
	case GreaterThan:
		cmp = greater
	case LessThan:
		cmp = less
	case TargetContains:
		cmp = contains
	case TargetIn:
		cmp = in
	case TargetPrefixes:
		cmp = prefixed
	case TargetPathPrefix:
		cmp = pathPrefix
		hasSlice = true
	case TargetPathContains:
		cmp = pathContains
		hasSlice = true
	case Passes:
		return true
	case Fails:
		return false
	}

	targets := []string{f.Target}
	if hasSlice {
		targets = f.NormalizedTargets
	}

	for _, tgt := range targets {
		success := cmp(norm(tgt), norm(input))
		if f.Negate {
			success = !success
		}

		// any-match
		if success {
			return true
		}
	}

	return false
}

// true if t == i
func equals(target, input string) bool {
	return target == input
}

// true if t > i
func greater(target, input string) bool {
	return target > input
}

// true if t < i
func less(target, input string) bool {
	return target < input
}

// true if target contains input as a substring.
func contains(target, input string) bool {
	return strings.Contains(target, input)
}

// true if input contains target as a substring.
func in(target, input string) bool {
	return strings.Contains(input, target)
}

// true if target has input as a prefix.
func prefixed(target, input string) bool {
	return strings.HasPrefix(input, target)
}

// true if target is an _element complete_ prefix match
// on the input.  Element complete means we do not
// succeed on partial element matches (ex: "/foo" does
// not match "/foobar").
//
// As a precondition, assumes the target value has been
// passed through normPathElem().
//
// The input is assumed to be the complete path that may
// have the target as a prefix.
func pathPrefix(target, input string) bool {
	return strings.HasPrefix(normPathElem(input), target)
}

// true if target has an _element complete_ equality
// with any element, or any sequence of elements, from
// the input.  Element complete means we do not succeed
// on partial element matches (ex: foo does not match
// /foobar, and foo/bar does not match foo/barbaz).
//
// As a precondition, assumes the target value has been
// passed through normPathElem().
//
// Input is assumed to be the complete path that may
// contain the target as an element or sequence of elems.
func pathContains(target, input string) bool {
	return strings.Contains(normPathElem(input), target)
}

// ----------------------------------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------------------------------

// prefixString maps the comparators to string prefixes for printing.
var prefixString = map[comparator]string{
	EqualTo:            "eq:",
	GreaterThan:        "gt:",
	LessThan:           "lt:",
	TargetContains:     "cont:",
	TargetIn:           "in:",
	TargetPrefixes:     "pfx:",
	TargetPathPrefix:   "pathPfx:",
	TargetPathContains: "pathCont:",
}

func (f Filter) String() string {
	switch f.Comparator {
	case Passes:
		return "pass"
	case Fails:
		return "fail"
	}

	if len(f.Targets) > 0 {
		return prefixString[f.Comparator] + strings.Join(f.Targets, ",")
	}

	return prefixString[f.Comparator] + f.Target
}
