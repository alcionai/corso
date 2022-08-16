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
	filterFunc := filters.NewEquals
	f := filterFunc(false, "", "foo")
	nf := filterFunc(true, "", "foo")

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
	filterFunc := filters.NewGreater
	f := filterFunc(false, "", "5")
	nf := filterFunc(true, "", "5")

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
	filterFunc := filters.NewLess
	f := filterFunc(false, "", "5")
	nf := filterFunc(true, "", "5")

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
	filterFunc := filters.NewBetween
	f := filterFunc(false, "", "abc", "def")
	nf := filterFunc(true, "", "abc", "def")

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
	filterFunc := filters.NewContains
	f := filterFunc(false, "", "smurfs")
	nf := filterFunc(true, "", "smurfs")

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

func (suite *FiltersSuite) TestIn() {
	filterFunc := filters.NewIn
	f := filterFunc(false, "", "murf")
	nf := filterFunc(true, "", "murf")

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
