package filters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/filters"
)

type FiltersSuite struct {
	suite.Suite
}

func TestFiltersSuite(t *testing.T) {
	suite.Run(t, new(FiltersSuite))
}

func (suite *FiltersSuite) TestEquals() {
	makeFilt := filters.NewEquals
	f := makeFilt(false, "", "foo")
	nf := makeFilt(true, "", "foo")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestGreater() {
	makeFilt := filters.NewGreater
	f := makeFilt(false, "", "5")
	nf := makeFilt(true, "", "5")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestLess() {
	makeFilt := filters.NewLess
	f := makeFilt(false, "", "5")
	nf := makeFilt(true, "", "5")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestBetween() {
	makeFilt := filters.NewBetween
	f := makeFilt(false, "", "abc", "def")
	nf := makeFilt(true, "", "abc", "def")

	table := []struct {
		input    string
		expectF  assert.BoolAssertionFunc
		expectNF assert.BoolAssertionFunc
	}{
		{"cd", assert.True, assert.False},
		{"a", assert.False, assert.True},
		{"1", assert.False, assert.True},
		{"f", assert.False, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestContains() {
	makeFilt := filters.NewContains
	f := makeFilt(false, "", "smurfs")
	nf := makeFilt(true, "", "smurfs")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestContains_Joined() {
	makeFilt := filters.NewContains
	f := makeFilt(false, "", "smarf,userid")
	nf := makeFilt(true, "", "smarf,userid")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn() {
	makeFilt := filters.NewIn
	f := makeFilt(false, "", "murf")
	nf := makeFilt(true, "", "murf")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}

func (suite *FiltersSuite) TestIn_Joined() {
	makeFilt := filters.NewIn
	f := makeFilt(false, "", "userid")
	nf := makeFilt(true, "", "userid")

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
			test.expectF(t, f.Matches(test.input), "filter")
			test.expectNF(t, nf.Matches(test.input), "negated filter")
		})
	}
}
