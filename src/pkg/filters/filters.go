package filters

import (
	"strings"
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
	Target     string     `json:"target"` // the value to compare against
	Negate     bool       `json:"negate"` // when true, negate the comparator result
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

// newFilter is the standard filter constructor.
func newFilter(c comparator, target string, negate bool) Filter {
	return Filter{c, target, negate}
}

// ----------------------------------------------------------------------------------------------------
// Comparisons
// ----------------------------------------------------------------------------------------------------

// Compare checks whether the input passes the filter.
func (f Filter) Compare(input string) bool {
	var cmp func(string, string) bool

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
	case Passes:
		return true
	case Fails:
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

// prefixString maps the comparators to string prefixes for printing.
var prefixString = map[comparator]string{
	EqualTo:        "eq:",
	GreaterThan:    "gt:",
	LessThan:       "lt:",
	TargetContains: "cont:",
	TargetIn:       "in:",
}

func (f Filter) String() string {
	switch f.Comparator {
	case Passes:
		return "pass"
	case Fails:
		return "fail"
	}

	return prefixString[f.Comparator] + f.Target
}
