package filters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/filters"
)

type FiltersSuite struct {
	suite.Suite
}

func TestFiltersSuite(t *testing.T) {
	suite.Run(t, new(FiltersSuite))
}

func (suite *FiltersSuite) TestEquals() {
	f := filters.Equal("foo")
	nf := filters.NotEqual("foo")

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"foo", assert.True, assert.False},
		{"bar", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestEquals_any() {
	f := filters.Equal("foo")
	nf := filters.NotEqual("foo")

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
		suite.T().Run(test.name, func(t *testing.T) {
			test.expectF(t, f.CompareAny(test.input...), "filter")
			test.expectNF(t, nf.CompareAny(test.input...), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestGreater() {
	f := filters.Greater("5")
	nf := filters.NotGreater("5")

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
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestLess() {
	f := filters.Less("5")
	nf := filters.NotLess("5")

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
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestContains() {
	f := filters.Contains("smurfs")
	nf := filters.NotContains("smurfs")

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"murf", assert.True, assert.False},
		{"frum", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestContains_Joined() {
	f := filters.Contains("smarf,userid")
	nf := filters.NotContains("smarf,userid")

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"userid", assert.True, assert.False},
		{"f,userid", assert.True, assert.False},
		{"fnords", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn() {
	f := filters.In("murf")
	nf := filters.NotIn("murf")

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"smurfs", assert.True, assert.False},
		{"sfrums", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn_Joined() {
	f := filters.In("userid")
	nf := filters.NotIn("userid")

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"smarf,userid", assert.True, assert.False},
		{"arf,user", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestPrefixes() {
	target := "folderA"
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
		suite.T().Run(test.name, func(t *testing.T) {
			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestPathPrefix() {
	table := []struct {
		name     string
		targets  []string
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"Exact - same case", []string{"fA"}, "/fA", assert.True, assert.False},
		{"Exact - different case", []string{"fa"}, "/fA", assert.True, assert.False},
		{"Prefix - same case", []string{"fA"}, "/fA/fB", assert.True, assert.False},
		{"Prefix - different case", []string{"fa"}, "/fA/fB", assert.True, assert.False},
		{"Exact - multiple folders", []string{"fA/fB"}, "/fA/fB", assert.True, assert.False},
		{"Prefix - single folder partial", []string{"f"}, "/fA/fB", assert.False, assert.True},
		{"Prefix - multi folder partial", []string{"fA/f"}, "/fA/fB", assert.False, assert.True},
		{"Target Longer - single folder", []string{"fA"}, "/f", assert.False, assert.True},
		{"Target Longer - multi folder", []string{"fA/fB"}, "/fA/f", assert.False, assert.True},
		{"Not prefix - single folder", []string{"fA"}, "/af", assert.False, assert.True},
		{"Not prefix - multi folder", []string{"fA/fB"}, "/fA/bf", assert.False, assert.True},
		{"Exact - target variations - none", []string{"fA"}, "/fA", assert.True, assert.False},
		{"Exact - target variations - prefix", []string{"/fA"}, "/fA", assert.True, assert.False},
		{"Exact - target variations - suffix", []string{"fA/"}, "/fA", assert.True, assert.False},
		{"Exact - target variations - both", []string{"/fA/"}, "/fA", assert.True, assert.False},
		{"Exact - input variations - none", []string{"fA"}, "fA", assert.True, assert.False},
		{"Exact - input variations - prefix", []string{"fA"}, "/fA", assert.True, assert.False},
		{"Exact - input variations - suffix", []string{"fA"}, "fA/", assert.True, assert.False},
		{"Exact - input variations - both", []string{"fA"}, "/fA/", assert.True, assert.False},
		{"Prefix - target variations - none", []string{"fA"}, "/fA/fb", assert.True, assert.False},
		{"Prefix - target variations - prefix", []string{"/fA"}, "/fA/fb", assert.True, assert.False},
		{"Prefix - target variations - suffix", []string{"fA/"}, "/fA/fb", assert.True, assert.False},
		{"Prefix - target variations - both", []string{"/fA/"}, "/fA/fb", assert.True, assert.False},
		{"Prefix - input variations - none", []string{"fA"}, "fA/fb", assert.True, assert.False},
		{"Prefix - input variations - prefix", []string{"fA"}, "/fA/fb", assert.True, assert.False},
		{"Prefix - input variations - suffix", []string{"fA"}, "fA/fb/", assert.True, assert.False},
		{"Prefix - input variations - both", []string{"fA"}, "/fA/fb/", assert.True, assert.False},
		{"Slice - one matches", []string{"foo", "fa/f", "fA"}, "/fA/fb", assert.True, assert.True},
		{"Slice - none match", []string{"foo", "fa/f", "f"}, "/fA/fb", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			f := filters.PathPrefix(test.targets)
			nf := filters.NotPathPrefix(test.targets)

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestPathContains() {
	table := []struct {
		name     string
		targets  []string
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"Exact - same case", []string{"fA"}, "/fA", assert.True, assert.False},
		{"Exact - different case", []string{"fa"}, "/fA", assert.True, assert.False},
		{"Cont - same case single target", []string{"fA"}, "/Z/fA/B", assert.True, assert.False},
		{"Cont - different case single target", []string{"fA"}, "/z/fa/b", assert.True, assert.False},
		{"Cont - same case multi target", []string{"Z/fA"}, "/Z/fA/B", assert.True, assert.False},
		{"Cont - different case multi target", []string{"fA/B"}, "/z/fa/b", assert.True, assert.False},
		{"Exact - multiple folders", []string{"Z/fA/B"}, "/Z/fA/B", assert.True, assert.False},
		{"Cont - single folder partial", []string{"folder"}, "/Z/fA/fB", assert.False, assert.True},
		{"Cont - multi folder partial", []string{"fA/fold"}, "/Z/fA/fB", assert.False, assert.True},
		{"Target Longer - single folder", []string{"fA"}, "/folder", assert.False, assert.True},
		{"Target Longer - multi folder", []string{"fA/fB"}, "/fA/fold", assert.False, assert.True},
		{"Not cont - single folder", []string{"fA"}, "/afolder", assert.False, assert.True},
		{"Not cont - single target", []string{"fA"}, "/z/afolder/bfolder", assert.False, assert.True},
		{"Not cont - multi folder", []string{"fA/fB"}, "/z/fA/bfolder", assert.False, assert.True},
		{"Exact - target variations - none", []string{"fA"}, "/fA", assert.True, assert.False},
		{"Exact - target variations - prefix", []string{"/fA"}, "/fA", assert.True, assert.False},
		{"Exact - target variations - suffix", []string{"fA/"}, "/fA", assert.True, assert.False},
		{"Exact - target variations - both", []string{"/fA/"}, "/fA", assert.True, assert.False},
		{"Exact - input variations - none", []string{"fA"}, "fA", assert.True, assert.False},
		{"Exact - input variations - prefix", []string{"fA"}, "/fA", assert.True, assert.False},
		{"Exact - input variations - suffix", []string{"fA"}, "fA/", assert.True, assert.False},
		{"Exact - input variations - both", []string{"fA"}, "/fA/", assert.True, assert.False},
		{"Cont - target variations - none", []string{"fA"}, "/fA/fb", assert.True, assert.False},
		{"Cont - target variations - prefix", []string{"/fA"}, "/fA/fb", assert.True, assert.False},
		{"Cont - target variations - suffix", []string{"fA/"}, "/fA/fb", assert.True, assert.False},
		{"Cont - target variations - both", []string{"/fA/"}, "/fA/fb", assert.True, assert.False},
		{"Cont - input variations - none", []string{"fA"}, "fA/fb", assert.True, assert.False},
		{"Cont - input variations - prefix", []string{"fA"}, "/fA/fb", assert.True, assert.False},
		{"Cont - input variations - suffix", []string{"fA"}, "fA/fb/", assert.True, assert.False},
		{"Cont - input variations - both", []string{"fA"}, "/fA/fb/", assert.True, assert.False},
		{"Slice - one matches", []string{"foo", "fa/f", "fA"}, "/fA/fb", assert.True, assert.True},
		{"Slice - none match", []string{"foo", "fa/f", "f"}, "/fA/fb", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			f := filters.PathContains(test.targets)
			nf := filters.NotPathContains(test.targets)

			test.expectF(t, f.Compare(test.input), "filter")
			test.expectNF(t, nf.Compare(test.input), "negated filter")
		})
	}
}
