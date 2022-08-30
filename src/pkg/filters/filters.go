package filters

import (
	"strings"
)

type comparator int

const (
	UnknownComparator comparator = iota
	// a == b
	Equal
	// a > b
	Greater
	// a < b
	Less
	// a < b < c
	Between
	// "foo,bar,baz" contains "foo"
	Contains
	// "foo" is found in "foo,bar,baz"
	In
	// always passes
	Pass
	// always fails
	Fail
	// passthrough for the target
	Identity
)

const delimiter = ","

func join(s ...string) string {
	return strings.Join(s, delimiter)
}

func split(s string) []string {
	return strings.Split(s, delimiter)
}

func norm(s string) string {
	return strings.ToLower(s)
}

// Filter contains a comparator func and the target to
// compare values against.  Filter.Matches(v) returns
// true if Filter.Comparer(filter.target, v) is true.
type Filter struct {
	Comparator comparator `json:"comparator"`
	Category   any        `json:"category"` // a caller-provided identifier.  Probably an iota or string const.
	Target     string     `json:"target"`   // the value to compare against
	Negate     bool       `json:"negate"`   // when true, negate the comparator result
}

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

// NewEquals creates a filter which Matches(v) is true if
// target == v
func NewEquals(negate bool, category any, target string) Filter {
	return Filter{Equal, category, target, negate}
}

// NewGreater creates a filter which Matches(v) is true if
// target > v
func NewGreater(negate bool, category any, target string) Filter {
	return Filter{Greater, category, target, negate}
}

// NewLess creates a filter which Matches(v) is true if
// target < v
func NewLess(negate bool, category any, target string) Filter {
	return Filter{Less, category, target, negate}
}

// NewBetween creates a filter which Matches(v) is true if
// lesser < v && v < greater
func NewBetween(negate bool, category any, lesser, greater string) Filter {
	return Filter{Between, category, join(lesser, greater), negate}
}

// NewContains creates a filter which Matches(v) is true if
// super.Contains(v)
func NewContains(negate bool, category any, super string) Filter {
	return Filter{Contains, category, super, negate}
}

// NewIn creates a filter which Matches(v) is true if
// v.Contains(substr)
func NewIn(negate bool, category any, substr string) Filter {
	return Filter{In, category, substr, negate}
}

// NewPass creates a filter where Matches(v) always returns true
func NewPass() Filter {
	return Filter{Pass, nil, "*", false}
}

// NewFail creates a filter where Matches(v) always returns false
func NewFail() Filter {
	return Filter{Fail, nil, "", false}
}

// NewIdentity creates a filter intended to hold values, rather than
// compare them.  Functionally, it'll behave the same as Equals.
func NewIdentity(id string) Filter {
	return Filter{Identity, nil, id, false}
}

// ----------------------------------------------------------------------------------------------------
// Comparisons
// ----------------------------------------------------------------------------------------------------

// Checks whether the filter matches the input
func (f Filter) Matches(input string) bool {
	var cmp func(string, string) bool

	switch f.Comparator {
	case Equal, Identity:
		cmp = equals
	case Greater:
		cmp = greater
	case Less:
		cmp = less
	case Between:
		cmp = between
	case Contains:
		cmp = contains
	case In:
		cmp = in
	case Pass:
		return true
	case Fail:
		return false
	}

	result := cmp(norm(f.Target), norm(input))
	if f.Negate {
		result = !result
	}

	return result
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

// assumes target is a delimited string.
// true if both:
//   - less(target[0], input)
//   - greater(target[1], input)
func between(target, input string) bool {
	parts := split(target)
	return less(parts[0], input) && greater(parts[1], input)
}

// true if target contains input as a substring.
func contains(target, input string) bool {
	return strings.Contains(target, input)
}

// true if input contains target as a substring.
func in(target, input string) bool {
	return strings.Contains(input, target)
}

// ----------------------------------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------------------------------

// Targets returns the Target value split into a slice.
func (f Filter) Targets() []string {
	return split(f.Target)
}

func (f Filter) String() string {
	var prefix string

	switch f.Comparator {
	case Equal:
		prefix = "eq:"
	case Greater:
		prefix = "gt:"
	case Less:
		prefix = "lt:"
	case Between:
		prefix = "btwn:"
	case Contains:
		prefix = "cont:"
	case In:
		prefix = "in:"
	case Pass:
		return "pass"
	case Fail:
		return "fail"
	case Identity:
	default: // no prefix
	}

	return prefix + f.Target
}
