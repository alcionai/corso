package filters

import (
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/pkg/path"
)

type comparator string

//go:generate stringer -type=comparator -linecomment
const (
	UnknownComparator comparator = ""
	// norm(a) == norm(b)
	EqualTo = "EQ"
	// a === b
	StrictEqualTo = "StrictEQ"
	// a > b
	GreaterThan = "GT"
	// a < b
	LessThan = "LT"
	// "foo,bar,baz" contains "foo"
	TargetContains = "Cont"
	// "foo" is found in "foo,bar,baz"
	TargetIn = "IN"
	// always passes
	Passes = "Pass"
	// always fails
	Fails = "Fail"
	// passthrough for the target
	IdentityValue = "Identity"
	// "foo" is a prefix of "foobarbaz"
	TargetPrefixes = "Pfx"
	// "baz" is a suffix of "foobarbaz"
	TargetSuffixes = "Sfx"
	// "foo" equals any complete element prefix of "foo/bar/baz"
	TargetPathPrefix = "PathPfx"
	// "foo" equals any complete element in "foo/bar/baz"
	TargetPathContains = "PathCont"
	// "baz" equals any complete element suffix of "foo/bar/baz"
	TargetPathSuffix = "PathSfx"
	// "foo/bar/baz" equals the complete path "foo/bar/baz"
	TargetPathEquals = "PathEQ"
)

func (c comparator) String() string {
	return string(c)
}

func normAll(ss []string) []string {
	r := slices.Clone(ss)
	for i := range r {
		r[i] = norm(r[i])
	}

	return r
}

func norm(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// normPathElem ensures the string is:
// 1. prefixed with a single path.pathSeparator (ex: `/`)
// 2. suffixed with a single path.pathSeparator (ex: `/`)
// This is done to facilitate future regex comparisons
// without re-running the prefix-suffix addition multiple
// times per target.
func normPathElem(s string) string {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return s
	}

	if s[0] != path.PathSeparator {
		s = string(path.PathSeparator) + s
	}

	s = path.TrimTrailingSlash(s)
	s = strings.ToLower(s)
	s += string(path.PathSeparator)

	return s
}

// Filter contains a comparator func and the target to
// compare values against.  Filter.Matches(v) returns
// true if Filter.Comparer(filter.target, v) is true.
type Filter struct {
	Comparator        comparator `json:"comparator_type"`   // the type of comparison
	Targets           []string   `json:"targets"`           // the set of values to compare
	NormalizedTargets []string   `json:"normalizedTargets"` // the set of comparable values post normalization
	Negate            bool       `json:"negate"`            // when true, negate the comparator result

	// only used when the filter's purpose is to hold a value without intent for comparison
	Identity string `json:"identity"`

	// deprecated, kept around for deserialization
	Target        string `json:"target"` // the value to compare against
	ComparatorInt int    `json:"comparator"`
}

// ----------------------------------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------------------------------

// Identity creates a filter intended to hold values, rather than
// compare them.  Comparatively, it'll behave the same as Equals.
func Identity(id string) Filter {
	return Filter{
		Comparator:        IdentityValue,
		Identity:          id,
		Targets:           []string{id},
		NormalizedTargets: normAll([]string{id}),
	}
}

// Equal creates a filter where Compare(v) is true if, for any target string,
// norm(target) == norm(v)
func Equal(target []string) Filter {
	return newFilter(EqualTo, target, normAll(target), false)
}

// NotEqual creates a filter where Compare(v) is true if, for any target string,
// target != v
func NotEqual(target []string) Filter {
	return newFilter(EqualTo, target, normAll(target), true)
}

// StrictEqual creates a filter where Compare(v) is true if, for any target string,
// target === v.  Target and v are not normalized for this comparison.  The comparison
// is case sensitive and ignores character folding.
func StrictEqual(target []string) Filter {
	return newFilter(StrictEqualTo, target, normAll(target), false)
}

// NotStrictEqual creates a filter where Compare(v) is true if, for any target string,
// target != v
func NotStrictEqual(target []string) Filter {
	return newFilter(StrictEqualTo, target, normAll(target), true)
}

// Greater creates a filter where Compare(v) is true if, for any target string,
// target > v
func Greater(target []string) Filter {
	return newFilter(GreaterThan, target, normAll(target), false)
}

// NotGreater creates a filter where Compare(v) is true if, for any target string,
// target <= v
func NotGreater(target []string) Filter {
	return newFilter(GreaterThan, target, normAll(target), true)
}

// Less creates a filter where Compare(v) is true if, for any target string,
// target < v
func Less(target []string) Filter {
	return newFilter(LessThan, target, normAll(target), false)
}

// NotLess creates a filter where Compare(v) is true if, for any target string,
// target >= v
func NotLess(target []string) Filter {
	return newFilter(LessThan, target, normAll(target), true)
}

// Contains creates a filter where Compare(v) is true if, for any target string,
// target.Contains(v)
func Contains(target []string) Filter {
	return newFilter(TargetContains, target, normAll(target), false)
}

// NotContains creates a filter where Compare(v) is true if, for any target string,
// !target.Contains(v)
func NotContains(target []string) Filter {
	return newFilter(TargetContains, target, normAll(target), true)
}

// In creates a filter where Compare(v) is true if, for any target string,
// v.Contains(target)
func In(target []string) Filter {
	return newFilter(TargetIn, target, normAll(target), false)
}

// NotIn creates a filter where Compare(v) is true if, for any target string,
// !v.Contains(target)
func NotIn(target []string) Filter {
	return newFilter(TargetIn, target, normAll(target), true)
}

// Pass creates a filter where Compare(v) always returns true
func Pass() Filter {
	return newFilter(Passes, []string{"*"}, nil, false)
}

// Fail creates a filter where Compare(v) always returns false
func Fail() Filter {
	return newFilter(Fails, []string{""}, nil, false)
}

// Prefix creates a filter where Compare(v) is true if, for any target string,
// target.Prefix(v)
func Prefix(target []string) Filter {
	return newFilter(TargetPrefixes, target, normAll(target), false)
}

// NotPrefix creates a filter where Compare(v) is true if, for any target string,
// !target.Prefix(v)
func NotPrefix(target []string) Filter {
	return newFilter(TargetPrefixes, target, normAll(target), true)
}

// Suffix creates a filter where Compare(v) is true if, for any target string,
// target.Suffix(v)
func Suffix(target []string) Filter {
	return newFilter(TargetSuffixes, target, normAll(target), false)
}

// NotSuffix creates a filter where Compare(v) is true if, for any target string,
// !target.Suffix(v)
func NotSuffix(target []string) Filter {
	return newFilter(TargetSuffixes, target, normAll(target), true)
}

// PathPrefix creates a filter where Compare(v) is true if, for any target string,
// target.Prefix(v) &&
// split(target)[i].Equals(split(v)[i]) for _all_ i in 0..len(target)-1
// ex: target "/foo/bar" returns true for input "/foo/bar/baz",
// but false for "/foo/barbaz"
func PathPrefix(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathPrefix, targets, tgts, false)
}

// NotPathPrefix creates a filter where Compare(v) is true if, for any target string,
// !target.Prefix(v) ||
// !split(target)[i].Equals(split(v)[i]) for _any_ i in 0..len(target)-1
// ex: target "/foo/bar" returns false for input "/foo/bar/baz",
// but true for "/foo/barbaz"
func NotPathPrefix(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathPrefix, targets, tgts, true)
}

// PathContains creates a filter where Compare(v) is true if, for any target string,
// for _any_ elem e in split(v), target.Equals(e) ||
// for _any_ sequence of elems in split(v), target.Equals(path.Join(e[n:m]))
// ex: target "foo" returns true for input "/baz/foo/bar",
// but false for "/baz/foobar"
// ex: target "baz/foo" returns true for input "/baz/foo/bar",
// but false for "/baz/foobar"
func PathContains(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathContains, targets, tgts, false)
}

// NotPathContains creates a filter where Compare(v) is true if, for any target string,
// for _every_ elem e in split(v), !target.Equals(e) ||
// for _every_ sequence of elems in split(v), !target.Equals(path.Join(e[n:m]))
// ex: target "foo" returns false for input "/baz/foo/bar",
// but true for "/baz/foobar"
// ex: target "baz/foo" returns false for input "/baz/foo/bar",
// but true for "/baz/foobar"
func NotPathContains(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathContains, targets, tgts, true)
}

// PathSuffix creates a filter where Compare(v) is true if, for any target string,
// target.Suffix(v) &&
// split(target)[i].Equals(split(v)[i]) for _all_ i in 0..len(target)-1
// ex: target "/bar/baz" returns true for input "/foo/bar/baz",
// but false for "/foobar/baz"
func PathSuffix(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathSuffix, targets, tgts, false)
}

// NotPathSuffix creates a filter where Compare(v) is true if
// !target.Suffix(v) ||
// !split(target)[i].Equals(split(v)[i]) for _any_ i in 0..len(target)-1
// ex: target "/bar/baz" returns false for input "/foo/bar/baz",
// but true for "/foobar/baz"
func NotPathSuffix(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathSuffix, targets, tgts, true)
}

// PathEquals creates a filter where Compare(v) is true if, for any target string,
// target.Equals(v) &&
// split(target)[i].Equals(split(v)[i]) for _all_ i in 0..len(target)-1
// ex: target "foo" returns true for inputs "/foo/", "/foo", and "foo/"
// but false for "/foo/bar", "bar/foo/", and "/foobar/"
func PathEquals(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathEquals, targets, tgts, false)
}

// NotPathEquals creates a filter where Compare(v) is true if, for any target string,
// !target.Equals(v) ||
// !split(target)[i].Equals(split(v)[i]) for _all_ i in 0..len(target)-1
// ex: target "foo" returns true "/foo/bar", "bar/foo/", and "/foobar/"
// but false for for inputs "/foo/", "/foo", and "foo/"
func NotPathEquals(targets []string) Filter {
	tgts := make([]string, len(targets))
	for i := range targets {
		tgts[i] = normPathElem(targets[i])
	}

	return newFilter(TargetPathEquals, targets, tgts, true)
}

// newFilter constructs filters that contain multiple targets
func newFilter(c comparator, targets, normTargets []string, negate bool) Filter {
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
	var cmp func(string, string) bool

	// select comparison func
	switch f.Comparator {
	case EqualTo, IdentityValue, TargetPathEquals:
		cmp = equals
	case StrictEqualTo:
		cmp = strictEquals
	case GreaterThan:
		cmp = greater
	case LessThan:
		cmp = less
	case TargetContains, TargetPathContains:
		cmp = contains
	case TargetIn:
		cmp = in
	case TargetPrefixes, TargetPathPrefix:
		cmp = prefixed
	case TargetSuffixes, TargetPathSuffix:
		cmp = suffixed
	case Passes:
		return true
	case Fails:
		return false
	}

	var (
		res     bool
		targets = f.NormalizedTargets
		_input  = norm(input)
		// most comparators expect cmp(target, input)
		// path comparators expect cmp(input, target)
		swapParams bool
	)

	// set conditional behavior
	switch f.Comparator {
	case TargetContains:
		// legacy case handling for contains, which checks for
		// strings.Contains(target, input) instead of (input, target)
		swapParams = true
	case StrictEqualTo:
		targets = f.Targets
		_input = input
	case TargetPathPrefix, TargetPathContains, TargetPathSuffix, TargetPathEquals:
		// As a precondition, assumes each entry in the NormalizedTargets
		// list has been passed through normPathElem().
		_input = normPathElem(input)
	}

	if len(targets) == 0 {
		targets = f.Targets
	}

	for _, tgt := range targets {
		t, i := tgt, _input

		if swapParams {
			t, i = _input, tgt
		}

		res = cmp(t, i)

		// any-match
		if res {
			break
		}
	}

	if f.Negate {
		res = !res
	}

	return res
}

// true if t == i, case insensitive and folded
func equals(target, input string) bool {
	return strings.EqualFold(target, input)
}

// true if t == i, case sensitive and not folded
func strictEquals(target, input string) bool {
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

// true if input contains target as a substring.
func contains(target, input string) bool {
	return strings.Contains(input, target)
}

// true if input contains target as a substring.
func in(target, input string) bool {
	return strings.Contains(input, target)
}

// true if target has input as a prefix.
func prefixed(target, input string) bool {
	return strings.HasPrefix(input, target)
}

// true if target has input as a suffix.
func suffixed(target, input string) bool {
	return strings.HasSuffix(input, target)
}

// ----------------------------------------------------------------------------------------------------
// Printers and PII control
// ----------------------------------------------------------------------------------------------------

var _ clues.PlainConcealer = &Filter{}

var safeFilterValues = map[string]struct{}{"*": {}}

func (f Filter) Conceal() string {
	fcs := f.Comparator

	switch f.Comparator {
	case Passes, Fails:
		return string(fcs)
	}

	concealed := pii.ConcealElements(f.Targets, safeFilterValues)

	return string(fcs) + ":" + strings.Join(concealed, ",")
}

func (f Filter) Format(fs fmt.State, _ rune) {
	fmt.Fprint(fs, f.Conceal())
}

func (f Filter) String() string {
	return f.Conceal()
}

func (f Filter) PlainString() string {
	fcs := f.Comparator

	switch f.Comparator {
	case Passes, Fails:
		return string(fcs)
	}

	return string(fcs) + ":" + strings.Join(f.Targets, ",")
}
