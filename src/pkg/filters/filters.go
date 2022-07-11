package filters

import "strings"

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
	// "foo" contains "f"
	Contains
	// "f" is found in "foo"
	In
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

// NewEquals creates a filter which Matches(v) is true if
// target == v
func NewEquals(negate bool, category any, target string) Filter {
	return Filter{Equal, category, norm(target), negate}
}

// NewGreater creates a filter which Matches(v) is true if
// target > v
func NewGreater(negate bool, category any, target string) Filter {
	return Filter{Greater, category, norm(target), negate}
}

// NewLess creates a filter which Matches(v) is true if
// target < v
func NewLess(negate bool, category any, target string) Filter {
	return Filter{Less, category, norm(target), negate}
}

// NewBetween creates a filter which Matches(v) is true if
// lesser < v && v < greater
func NewBetween(negate bool, category any, lesser, greater string) Filter {
	return Filter{Between, category, norm(join(lesser, greater)), negate}
}

// NewContains creates a filter which Matches(v) is true if
// super.Contains(v)
func NewContains(negate bool, category any, super string) Filter {
	return Filter{Contains, category, norm(super), negate}
}

// NewIn creates a filter which Matches(v) is true if
// v.Contains(substr)
func NewIn(negate bool, category any, substr string) Filter {
	return Filter{In, category, norm(substr), negate}
}

// Checks whether the filter matches the input
func (f Filter) Matches(input string) bool {
	var cmp func(string, string) bool
	switch f.Comparator {
	case Equal:
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
	}
	result := cmp(f.Target, norm(input))
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
