package filters_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/filters"
)

type FiltersSuite struct {
	tester.Suite
}

func TestFiltersSuite(t *testing.T) {
	suite.Run(t, &FiltersSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *FiltersSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *FiltersSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func sl(s ...string) []string {
	return append([]string{}, s...)
}

var (
	foo    = sl("foo")
	five   = sl("5")
	smurfs = sl("smurfs")
)

func (suite *FiltersSuite) TestEquals() {
	f := filters.Equal(foo)
	nf := filters.NotEqual(foo)

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"foo", assert.True, assert.False},
		{"FOO", assert.True, assert.False},
		{" foo ", assert.True, assert.False},
		{"bar", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestEquals_any() {
	f := filters.Equal(foo)
	nf := filters.NotEqual(foo)

	table := []struct {
		name     string
		input    []string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"includes target", []string{"foo", "bar"}, assert.True, assert.True},
		{"not includes target", []string{"baz", "qux"}, assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			test.expectF(t, f.CompareAny(test.input...), "filter")
			test.expectNF(t, nf.CompareAny(test.input...), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestStrictEquals() {
	f := filters.StrictEqual(foo)
	nf := filters.NotStrictEqual(foo)

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"foo", assert.True, assert.False},
		{"FOO", assert.False, assert.True},
		{" foo ", assert.False, assert.True},
		{"bar", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestGreater() {
	f := filters.Greater(five)
	nf := filters.NotGreater(five)

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"4", assert.True, assert.False},
		{"5", assert.False, assert.True},
		{"6", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestLess() {
	f := filters.Less(five)
	nf := filters.NotLess(five)

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"6", assert.True, assert.False},
		{"5", assert.False, assert.True},
		{"4", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestContains() {
	f := filters.Contains(smurfs)
	nf := filters.NotContains(smurfs)

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"murf", assert.True, assert.False},
		{"frum", assert.False, assert.True},
		{"ssmurfss", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn() {
	f := filters.In(sl("murf"))
	nf := filters.NotIn(sl("murf"))

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"smurfs", assert.True, assert.False},
		{"sfrums", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn_MultipleTargets() {
	f := filters.In(sl("murf", "foo"))
	nf := filters.NotIn(sl("murf", "foo"))

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"smurfs", assert.True, assert.False},
		{"foo", assert.True, assert.False},
		{"sfrums", assert.False, assert.True},
		{"oof", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn_MultipleTargets_Joined() {
	f := filters.In(sl("userid", "foo"))
	nf := filters.NotIn(sl("userid", "foo"))

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"smarf,userid", assert.True, assert.False},
		{"smarf,foo", assert.True, assert.False},
		{"arf,user", assert.False, assert.True},
		{"arf,oof", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn_Joined() {
	f := filters.In(sl("userid"))
	nf := filters.NotIn(sl("userid"))

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"smarf,userid", assert.True, assert.False},
		{"arf,user", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestPrefixes() {
	target := sl("folderA")
	f := filters.Prefix(target)
	nf := filters.NotPrefix(target)

	table := []struct {
		name     string
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"Exact match - same case", "folderA", assert.True, assert.False},
		{"Exact match - different case", "Foldera", assert.True, assert.False},
		{"Prefix match - same case", "folderA/folderB", assert.True, assert.False},
		{"Prefix match - different case", "Foldera/folderB", assert.True, assert.False},
		{"Should not match substring", "folder1/folderA", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestSuffixes() {
	target := sl("folderB")
	f := filters.Suffix(target)
	nf := filters.NotSuffix(target)

	table := []struct {
		name     string
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"Exact match - same case", "folderB", assert.True, assert.False},
		{"Exact match - different case", "Folderb", assert.True, assert.False},
		{"Suffix match - same case", "folderA/folderB", assert.True, assert.False},
		{"Suffix match - different case", "Foldera/folderb", assert.True, assert.False},
		{"Should not match substring", "folderB/folder1", assert.False, assert.True},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

// ---------------------------------------------------------------------------
// path comparators
// ---------------------------------------------------------------------------

var pathElemNormalizationTable = []struct {
	name    string
	targets []string
	expect  []string
}{
	{"Single - no slash", []string{"fA"}, []string{"/fa/"}},
	{"Single - pre slash", []string{"/fA"}, []string{"/fa/"}},
	{"Single - suff slash", []string{"fA/"}, []string{"/fa/"}},
	{"Single - both slashes", []string{"/fA/"}, []string{"/fa/"}},
	{"Multipath - no slash", []string{"fA/fB"}, []string{"/fa/fb/"}},
	{"Multipath - pre slash", []string{"/fA/fB"}, []string{"/fa/fb/"}},
	{"Multipath - suff slash", []string{"fA/fB/"}, []string{"/fa/fb/"}},
	{"Multipath - both slashes", []string{"/fA/fB/"}, []string{"/fa/fb/"}},
	{"Multi input - no slash", []string{"fA", "fB"}, []string{"/fa/", "/fb/"}},
	{"Multi input - pre slash", []string{"/fA", "/fB"}, []string{"/fa/", "/fb/"}},
	{"Multi input - suff slash", []string{"fA/", "fB/"}, []string{"/fa/", "/fb/"}},
	{"Multi input - both slashes", []string{"/fA/", "/fB/"}, []string{"/fa/", "/fb/"}},
}

type baf struct {
	fn  assert.BoolAssertionFunc
	yes bool
}

var (
	yes = baf{
		fn:  assert.True,
		yes: true,
	}
	no = baf{
		fn:  assert.False,
		yes: false,
	}
)

var pathComparisonsTable = []struct {
	name           string
	targets        []string
	input          string
	expectContains baf
	expectEquals   baf
	expectPrefix   baf
	expectSuffix   baf
}{
	{"single folder partial", []string{"f"}, "/fA", no, no, no, no},
	{"single folder target partial", []string{"f"}, "/fA/fB", no, no, no, no},
	{"multi folder input partial", []string{"A/f"}, "/fA/fB", no, no, no, no},
	{"longer target - single folder", []string{"fA"}, "/f", no, no, no, no},
	{"longer target - multi folder", []string{"fA/fB"}, "/fA/f", no, no, no, no},
	{"non-matching - single folder", []string{"fA"}, "/af", no, no, no, no},
	{"non-matching - multi folder", []string{"fA/fB"}, "/fA/bf", no, no, no, no},

	{"Exact - same case", []string{"fA"}, "/fA", yes, yes, yes, yes},
	{"Exact - different case", []string{"fa"}, "/fA", yes, yes, yes, yes},
	{"Exact - multiple folders", []string{"fA/fB"}, "/fA/fB", yes, yes, yes, yes},
	{"Exact - target slash variations - prefix", []string{"/fA"}, "/fA", yes, yes, yes, yes},
	{"Exact - target slash variations - suffix", []string{"fA/"}, "/fA", yes, yes, yes, yes},
	{"Exact - target slash variations - both", []string{"/fA/"}, "/fA", yes, yes, yes, yes},
	{"Exact - input slash variations - none", []string{"fA"}, "fA", yes, yes, yes, yes},
	{"Exact - input slash variations - prefix", []string{"fA"}, "/fA", yes, yes, yes, yes},
	{"Exact - input slash variations - suffix", []string{"fA"}, "fA/", yes, yes, yes, yes},
	{"Exact - input slash variations - both", []string{"fA"}, "/fA/", yes, yes, yes, yes},

	{"Prefix - same case", []string{"fA"}, "/fA/fB", yes, no, yes, no},
	{"Prefix - different case", []string{"fa"}, "/fA/fB", yes, no, yes, no},
	{"Prefix - multiple folders", []string{"fa/fb"}, "/fA/fB/fC", yes, no, yes, no},
	{"Prefix - target slash variations - none", []string{"fA"}, "/fA/fb", yes, no, yes, no},
	{"Prefix - target slash variations - prefix", []string{"/fA"}, "/fA/fb", yes, no, yes, no},
	{"Prefix - target slash variations - suffix", []string{"fA/"}, "/fA/fb", yes, no, yes, no},
	{"Prefix - target slash variations - both", []string{"/fA/"}, "/fA/fb", yes, no, yes, no},
	{"Prefix - input slash variations - none", []string{"fA"}, "fA/fb", yes, no, yes, no},
	{"Prefix - input slash variations - prefix", []string{"fA"}, "/fA/fb", yes, no, yes, no},
	{"Prefix - input slash variations - suffix", []string{"fA"}, "fA/fb/", yes, no, yes, no},
	{"Prefix - input slash variations - both", []string{"fA"}, "/fA/fb/", yes, no, yes, no},

	{"Suffix - same case", []string{"fB"}, "/fA/fB", yes, no, no, yes},
	{"Suffix - different case", []string{"fb"}, "/fA/fB", yes, no, no, yes},
	{"Suffix - multiple folders", []string{"fb/fc"}, "/fA/fB/fC", yes, no, no, yes},
	{"Suffix - target slash variations - none", []string{"fB"}, "/fA/fb", yes, no, no, yes},
	{"Suffix - target slash variations - prefix", []string{"/fB"}, "/fA/fb", yes, no, no, yes},
	{"Suffix - target slash variations - suffix", []string{"fB/"}, "/fA/fb", yes, no, no, yes},
	{"Suffix - target slash variations - both", []string{"/fB/"}, "/fA/fb", yes, no, no, yes},
	{"Suffix - input slash variations - none", []string{"fB"}, "fA/fb", yes, no, no, yes},
	{"Suffix - input slash variations - prefix", []string{"fB"}, "/fA/fb", yes, no, no, yes},
	{"Suffix - input slash variations - suffix", []string{"fB"}, "fA/fb/", yes, no, no, yes},
	{"Suffix - input slash variations - both", []string{"fB"}, "/fA/fb/", yes, no, no, yes},

	{"Contains - same case", []string{"fB"}, "/fA/fB/fC", yes, no, no, no},
	{"Contains - different case", []string{"fb"}, "/fA/fB/fC", yes, no, no, no},
	{"Contains - multiple folders", []string{"fb/fc"}, "/fA/fB/fC/fD", yes, no, no, no},
	{"Contains - target slash variations - none", []string{"fB"}, "/fA/fb/fc", yes, no, no, no},
	{"Contains - target slash variations - prefix", []string{"/fB"}, "/fA/fb/fc", yes, no, no, no},
	{"Contains - target slash variations - suffix", []string{"fB/"}, "/fA/fb/fc", yes, no, no, no},
	{"Contains - target slash variations - both", []string{"/fB/"}, "/fA/fb/fc", yes, no, no, no},
	{"Contains - input slash variations - none", []string{"fB"}, "fA/fb/fc", yes, no, no, no},
	{"Contains - input slash variations - prefix", []string{"fB"}, "/fA/fb/fc/", yes, no, no, no},
	{"Contains - input slash variations - suffix", []string{"fB"}, "fA/fb/fc/", yes, no, no, no},
	{"Contains - input slash variations - both", []string{"fB"}, "/fA/fb/fc/", yes, no, no, no},

	{"Slice - one exact matches", []string{"foo", "fa/f", "fA"}, "/fA", yes, yes, yes, yes},
	{"Slice - none match", []string{"foo", "fa/f", "f"}, "/fA", no, no, no, no},
}

func (suite *FiltersSuite) TestPathPrefix() {
	for _, test := range pathComparisonsTable {
		suite.Run(test.name, func() {
			var (
				t  = suite.T()
				f  = filters.PathPrefix(test.targets)
				nf = filters.NotPathPrefix(test.targets)
			)

			test.expectPrefix.fn(t, f.Compare(test.input), "filter")
			if test.expectPrefix.yes {
				no.fn(t, nf.Compare(test.input), "negated filter")
			} else {
				yes.fn(t, nf.Compare(test.input), "negated filter")
			}
		})
	}
}

func (suite *FiltersSuite) TestPathPrefix_NormalizedTargets() {
	for _, test := range pathElemNormalizationTable {
		suite.Run(test.name, func() {
			t := suite.T()

			f := filters.PathPrefix(test.targets)
			assert.Equal(t, test.expect, f.NormalizedTargets)
		})
	}
}

func (suite *FiltersSuite) TestPathContains() {
	for _, test := range pathComparisonsTable {
		suite.Run(test.name, func() {
			var (
				t  = suite.T()
				f  = filters.PathContains(test.targets)
				nf = filters.NotPathContains(test.targets)
			)

			test.expectContains.fn(t, f.Compare(test.input), "filter")
			if test.expectContains.yes {
				no.fn(t, nf.Compare(test.input), "negated filter")
			} else {
				yes.fn(t, nf.Compare(test.input), "negated filter")
			}
		})
	}
}

func (suite *FiltersSuite) TestPathContains_NormalizedTargets() {
	for _, test := range pathElemNormalizationTable {
		suite.Run(test.name, func() {
			t := suite.T()

			f := filters.PathContains(test.targets)
			assert.Equal(t, test.expect, f.NormalizedTargets)
		})
	}
}

func (suite *FiltersSuite) TestPathSuffix() {
	for _, test := range pathComparisonsTable {
		suite.Run(test.name, func() {
			var (
				t  = suite.T()
				f  = filters.PathSuffix(test.targets)
				nf = filters.NotPathSuffix(test.targets)
			)

			test.expectSuffix.fn(t, f.Compare(test.input), "filter")
			if test.expectSuffix.yes {
				no.fn(t, nf.Compare(test.input), "negated filter")
			} else {
				yes.fn(t, nf.Compare(test.input), "negated filter")
			}
		})
	}
}

func (suite *FiltersSuite) TestPathSuffix_NormalizedTargets() {
	for _, test := range pathElemNormalizationTable {
		suite.Run(test.name, func() {
			t := suite.T()

			f := filters.PathSuffix(test.targets)
			assert.Equal(t, test.expect, f.NormalizedTargets)
		})
	}
}

func (suite *FiltersSuite) TestPathEquals() {
	for _, test := range pathComparisonsTable {
		suite.Run(test.name, func() {
			var (
				t  = suite.T()
				f  = filters.PathEquals(test.targets)
				nf = filters.NotPathEquals(test.targets)
			)

			test.expectEquals.fn(t, f.Compare(test.input), "filter")
			if test.expectEquals.yes {
				no.fn(t, nf.Compare(test.input), "negated filter")
			} else {
				yes.fn(t, nf.Compare(test.input), "negated filter")
			}
		})
	}
}

func (suite *FiltersSuite) TestPathEquals_NormalizedTargets() {
	for _, test := range pathElemNormalizationTable {
		suite.Run(test.name, func() {
			t := suite.T()

			f := filters.PathEquals(test.targets)
			assert.Equal(t, test.expect, f.NormalizedTargets)
		})
	}
}

// ---------------------------------------------------------------------------
// pii handling
// ---------------------------------------------------------------------------

func (suite *FiltersSuite) TestFilter_pii() {
	targets := []string{"fnords", "smarf", "*"}

	table := []struct {
		name string
		f    filters.Filter
	}{
		{"equal", filters.Equal(targets)},
		{"contains", filters.Contains(targets)},
		{"greater", filters.Greater(targets)},
		{"less", filters.Less(targets)},
		{"prefix", filters.Prefix(targets)},
		{"suffix", filters.Suffix(targets)},
		{"pathprefix", filters.PathPrefix(targets)},
		{"pathsuffix", filters.PathSuffix(targets)},
		{"pathcontains", filters.PathContains(targets)},
		{"pathequals", filters.PathEquals(targets)},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t           = suite.T()
				expect      = string(test.f.Comparator) + ":***,***,*"
				expectPlain = string(test.f.Comparator) + ":" + strings.Join(targets, ",")
			)

			result := test.f.Conceal()
			assert.Equal(t, expect, result, "conceal")

			result = test.f.String()
			assert.Equal(t, expect, result, "string")

			result = test.f.PlainString()
			assert.Equal(t, expectPlain, result, "plainString")

			result = fmt.Sprintf("%s", test.f)
			assert.Equal(t, expect, result, "fmt %%s")

			result = fmt.Sprintf("%v", test.f)
			assert.Equal(t, expect, result, "fmt %%v")

			result = fmt.Sprintf("%+v", test.f)
			assert.Equal(t, expect, result, "fmt %%+v")
		})
	}

	table2 := []struct {
		name        string
		f           filters.Filter
		expect      string
		expectPlain string
	}{
		{"pass", filters.Pass(), "Pass", "Pass"},
		{"fail", filters.Fail(), "Fail", "Fail"},
		{
			"identity",
			filters.Identity("id"),
			filters.IdentityValue + ":***",
			filters.IdentityValue + ":id",
		},
		{
			"identity",
			filters.Identity("*"),
			filters.IdentityValue + ":*",
			filters.IdentityValue + ":*",
		},
	}
	for _, test := range table2 {
		suite.Run(test.name, func() {
			t := suite.T()

			result := test.f.Conceal()
			assert.Equal(t, test.expect, result, "conceal")

			result = test.f.String()
			assert.Equal(t, test.expect, result, "string")

			result = test.f.PlainString()
			assert.Equal(t, test.expectPlain, result, "plainString")

			result = fmt.Sprintf("%s", test.f)
			assert.Equal(t, test.expect, result, "fmt %%s")

			result = fmt.Sprintf("%v", test.f)
			assert.Equal(t, test.expect, result, "fmt %%v")

			result = fmt.Sprintf("%+v", test.f)
			assert.Equal(t, test.expect, result, "fmt %%+v")
		})
	}
}
