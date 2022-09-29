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
	input := "folderA"

	table := []struct {
		name     string
		target   string
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
			f := filters.Prefix(test.target)
			nf := filters.NotPrefix(test.target)
			test.expectF(t, f.Compare(input), "filter")
			test.expectNF(t, nf.Compare(input), "negated filter")
		})
	}
}
