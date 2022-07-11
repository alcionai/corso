package filters_test

import (
	"testing"

	"github.com/alcionai/corso/pkg/filters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FiltersSuite struct {
	suite.Suite
}

func TestFiltersSuite(t *testing.T) {
	suite.Run(t, new(FiltersSuite))
}

func (suite *FiltersSuite) TestEquals() {
	make := filters.NewEquals
	f := make(false, "", "foo")
	nf := make(true, "", "foo")

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
	make := filters.NewGreater
	f := make(false, "", "5")
	nf := make(true, "", "5")

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
	make := filters.NewLess
	f := make(false, "", "5")
	nf := make(true, "", "5")

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
	make := filters.NewBetween
	f := make(false, "", "abc", "def")
	nf := make(true, "", "abc", "def")

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
	make := filters.NewContains
	f := make(false, "", "smurfs")
	nf := make(true, "", "smurfs")

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
	make := filters.NewIn
	f := make(false, "", "murf")
	nf := make(true, "", "murf")

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
