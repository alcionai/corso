package selectors

import (
	"testing"

	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func stubPathValues() map[categorizer]string {
	return map[categorizer]string{
		rootCatStub: rootCatStub.String(),
		leafCatStub: leafCatStub.String(),
	}
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type SelectorScopesSuite struct {
	suite.Suite
}

func TestSelectorScopesSuite(t *testing.T) {
	suite.Run(t, new(SelectorScopesSuite))
}

func (suite *SelectorScopesSuite) TestContains() {
	table := []struct {
		name   string
		scope  func() mockScope
		check  string
		expect assert.BoolAssertionFunc
	}{
		{
			name: "any",
			scope: func() mockScope {
				stub := stubScope("")
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.True,
		},
		{
			name: "none",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = NoneTgt
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.False,
		},
		{
			name: "blank value",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = ""
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.False,
		},
		{
			name: "blank target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = "fnords"
				return stub
			},
			check:  "",
			expect: assert.False,
		},
		{
			name: "matching target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = rootCatStub.String()
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.True,
		},
		{
			name: "non-matching target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = rootCatStub.String()
				return stub
			},
			check:  "smarf",
			expect: assert.False,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(
				t,
				contains(test.scope(), rootCatStub, test.check))
		})
	}
}

func (suite *SelectorScopesSuite) TestGetCatValue() {
	t := suite.T()
	stub := stubScope("")
	stub[rootCatStub.String()] = rootCatStub.String()
	assert.Equal(t, []string{rootCatStub.String()}, getCatValue(stub, rootCatStub))
	assert.Equal(t, None(), getCatValue(stub, leafCatStub))
}

func (suite *SelectorScopesSuite) TestGranularity() {
	t := suite.T()
	stub := stubScope("")
	assert.Equal(t, Item, granularity(stub))
}

func (suite *SelectorScopesSuite) TestIsAnyTarget() {
	t := suite.T()
	stub := stubScope("")
	assert.True(t, isAnyTarget(stub, rootCatStub))
	assert.False(t, isAnyTarget(stub, leafCatStub))
}

var reduceTestTable = []struct {
	name         string
	sel          func() Selector
	expectLen    int
	expectPasses assert.BoolAssertionFunc
}{
	{
		name: "include all",
		sel: func() Selector {
			sel := stubSelector()
			sel.Filters = nil
			sel.Excludes = nil
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "include none",
		sel: func() Selector {
			sel := stubSelector()
			sel.Includes[0] = scope(stubScope("none"))
			sel.Filters = nil
			sel.Excludes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "filter and include all",
		sel: func() Selector {
			sel := stubSelector()
			sel.Excludes = nil
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "include all filter none",
		sel: func() Selector {
			sel := stubSelector()
			sel.Filters[0] = scope(stubScope("none"))
			sel.Excludes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "include all exclude all",
		sel: func() Selector {
			sel := stubSelector()
			sel.Filters = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "include all exclude none",
		sel: func() Selector {
			sel := stubSelector()
			sel.Filters = nil
			sel.Excludes[0] = scope(stubScope("none"))
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "filter all exclude all",
		sel: func() Selector {
			sel := stubSelector()
			sel.Includes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "filter all exclude none",
		sel: func() Selector {
			sel := stubSelector()
			sel.Includes = nil
			sel.Excludes[0] = scope(stubScope("none"))
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
}

func (suite *SelectorScopesSuite) TestReduce() {
	deets := func() details.Details {
		return details.Details{
			DetailsModel: details.DetailsModel{
				Entries: []details.DetailsEntry{
					{RepoRef: rootCatStub.String() + "/stub/" + leafCatStub.String()},
				},
			},
		}
	}
	dataCats := map[pathType]mockCategorizer{
		unknownPathType: rootCatStub,
	}
	for _, test := range reduceTestTable {
		suite.T().Run(test.name, func(t *testing.T) {
			ds := deets()
			result := reduce[mockScope](&ds, test.sel(), dataCats)
			require.NotNil(t, result)
			require.Len(t, result.Entries, test.expectLen)
		})
	}
}

func (suite *SelectorScopesSuite) TestPathTypeIn() {
	t := suite.T()
	assert.Equal(t, unknownPathType, pathTypeIn([]string{}), "empty")
	assert.Equal(t, exchangeMailPath, pathTypeIn([]string{"", "", "mail"}), "mail")
	assert.Equal(t, exchangeContactPath, pathTypeIn([]string{"", "", "contact"}), "contact")
	assert.Equal(t, exchangeEventPath, pathTypeIn([]string{"", "", "event"}), "event")
	assert.Equal(t, unknownPathType, pathTypeIn([]string{"", "", "fnords"}), "bogus")
}

func (suite *SelectorScopesSuite) TestScopesByCategory() {
	t := suite.T()
	s1 := stubScope("")
	s2 := stubScope("")
	s2[scopeKeyCategory] = unknownCatStub.String()
	result := scopesByCategory[mockScope](
		[]scope{scope(s1), scope(s2)},
		map[pathType]mockCategorizer{
			unknownPathType: rootCatStub,
		})
	assert.Len(t, result, 1)
	assert.Len(t, result[rootCatStub], 1)
	assert.Empty(t, result[leafCatStub])
}

func (suite *SelectorScopesSuite) TestPasses() {
	cat := rootCatStub
	pathVals := map[categorizer]string{}
	entry := details.DetailsEntry{}
	for _, test := range reduceTestTable {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := test.sel()
			excl := toMockScope(sel.Excludes)
			filt := toMockScope(sel.Filters)
			incl := toMockScope(sel.Includes)
			result := passes(
				cat,
				pathVals,
				entry,
				excl, filt, incl)
			test.expectPasses(t, result)
		})
	}
}

func toMockScope(sc []scope) []mockScope {
	if len(sc) == 0 {
		return nil
	}
	ms := []mockScope{}
	for _, s := range sc {
		ms = append(ms, mockScope(s))
	}
	return ms
}

func (suite *SelectorScopesSuite) TestMatchesPathValues() {
	t := suite.T()
	cat := rootCatStub
	sc := stubScope("")
	sc[rootCatStub.String()] = rootCatStub.String()
	sc[leafCatStub.String()] = leafCatStub.String()
	pvs := stubPathValues()
	assert.True(t, matchesPathValues(sc, cat, pvs), "matching values")
	// "any" seems like it should pass, but this is the path value,
	// not the scope value, so unless the scope is also "any", it fails.
	pvs[rootCatStub] = AnyTgt
	pvs[leafCatStub] = AnyTgt
	assert.False(t, matchesPathValues(sc, cat, pvs), "any")
	pvs[rootCatStub] = NoneTgt
	pvs[leafCatStub] = NoneTgt
	assert.False(t, matchesPathValues(sc, cat, pvs), "none")
	pvs[rootCatStub] = "foo"
	pvs[leafCatStub] = "bar"
	assert.False(t, matchesPathValues(sc, cat, pvs), "mismatched values")
}
