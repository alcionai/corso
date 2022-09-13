package selectors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/filters"
)

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
				stub[rootCatStub.String()] = failAny
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.False,
		},
		{
			name: "blank value",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filters.Equal("")
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.False,
		},
		{
			name: "blank target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filterize("fnords")
				return stub
			},
			check:  "",
			expect: assert.False,
		},
		{
			name: "matching target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filterize(rootCatStub.String())
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.True,
		},
		{
			name: "non-matching target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filterize(rootCatStub.String())
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
				matches(test.scope(), rootCatStub, test.check))
		})
	}
}

func (suite *SelectorScopesSuite) TestGetCatValue() {
	t := suite.T()
	stub := stubScope("")
	stub[rootCatStub.String()] = filterize(rootCatStub.String())
	assert.Equal(t,
		[]string{rootCatStub.String()},
		getCatValue(stub, rootCatStub))
	assert.Equal(t, None(), getCatValue(stub, leafCatStub))
}

func (suite *SelectorScopesSuite) TestIsAnyTarget() {
	t := suite.T()
	stub := stubScope("")
	assert.True(t, isAnyTarget(stub, rootCatStub))
	assert.False(t, isAnyTarget(stub, leafCatStub))
}

var reduceTestTable = []struct {
	name         string
	sel          func() mockSel
	expectLen    int
	expectPasses assert.BoolAssertionFunc
}{
	{
		name: "include all",
		sel: func() mockSel {
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
		sel: func() mockSel {
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
		sel: func() mockSel {
			sel := stubSelector()
			sel.Excludes = nil
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "include all filter none",
		sel: func() mockSel {
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
		sel: func() mockSel {
			sel := stubSelector()
			sel.Filters = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "include all exclude none",
		sel: func() mockSel {
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
		sel: func() mockSel {
			sel := stubSelector()
			sel.Includes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "filter all exclude none",
		sel: func() mockSel {
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
					{
						RepoRef: stubRepoRef(
							pathServiceStub,
							pathCatStub,
							rootCatStub.String(),
							"stub",
							leafCatStub.String(),
						),
					},
				},
			},
		}
	}
	dataCats := map[path.CategoryType]mockCategorizer{
		pathCatStub: rootCatStub,
	}

	for _, test := range reduceTestTable {
		suite.T().Run(test.name, func(t *testing.T) {
			ds := deets()
			result := reduce[mockScope](
				context.Background(),
				&ds,
				test.sel().Selector,
				dataCats)
			require.NotNil(t, result)
			assert.Len(t, result.Entries, test.expectLen)
		})
	}
}

func (suite *SelectorScopesSuite) TestScopesByCategory() {
	t := suite.T()
	s1 := stubScope("")
	s2 := stubScope("")
	s2[scopeKeyCategory] = filterize(unknownCatStub.String())
	result := scopesByCategory[mockScope](
		[]scope{scope(s1), scope(s2)},
		map[path.CategoryType]mockCategorizer{
			path.UnknownCategory: rootCatStub,
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
	cat := rootCatStub
	pvs := stubPathValues()

	table := []struct {
		name    string
		rootVal string
		leafVal string
		expect  assert.BoolAssertionFunc
	}{
		{
			name:    "matching values",
			rootVal: rootCatStub.String(),
			leafVal: leafCatStub.String(),
			expect:  assert.True,
		},
		{
			name:    "any",
			rootVal: AnyTgt,
			leafVal: AnyTgt,
			expect:  assert.True,
		},
		{
			name:    "none",
			rootVal: NoneTgt,
			leafVal: NoneTgt,
			expect:  assert.False,
		},
		{
			name:    "mismatched values",
			rootVal: "fnords",
			leafVal: "smarf",
			expect:  assert.False,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sc := stubScope("")
			sc[rootCatStub.String()] = filterize(test.rootVal)
			sc[leafCatStub.String()] = filterize(test.leafVal)

			test.expect(t, matchesPathValues(sc, cat, pvs))
		})
	}
}

func (suite *SelectorScopesSuite) TestAddToSet() {
	t := suite.T()
	set := []string{}

	set = addToSet(set, []string{})
	assert.Len(t, set, 0)

	set = addToSet(set, []string{"a"})
	assert.Len(t, set, 1)
	assert.Equal(t, set[0], "a")

	set = addToSet(set, []string{"a"})
	assert.Len(t, set, 1)

	set = addToSet(set, []string{"a", "b"})
	assert.Len(t, set, 2)
	assert.Equal(t, set[0], "a")
	assert.Equal(t, set[1], "b")

	set = addToSet(set, []string{"c", "d"})
	assert.Len(t, set, 4)
	assert.Equal(t, set[0], "a")
	assert.Equal(t, set[1], "b")
	assert.Equal(t, set[2], "c")
	assert.Equal(t, set[3], "d")
}

func (suite *SelectorScopesSuite) TestClean() {
	table := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "nil",
			input:  nil,
			expect: None(),
		},
		{
			name:   "has anyTgt",
			input:  []string{"a", AnyTgt},
			expect: Any(),
		},
		{
			name:   "has noneTgt",
			input:  []string{"a", NoneTgt},
			expect: None(),
		},
		{
			name:   "has anyTgt and noneTgt, any first",
			input:  []string{"a", AnyTgt, NoneTgt},
			expect: Any(),
		},
		{
			name:   "has noneTgt and anyTgt, none first",
			input:  []string{"a", NoneTgt, AnyTgt},
			expect: None(),
		},
		{
			name:   "already clean",
			input:  []string{"a", "b"},
			expect: []string{"a", "b"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := clean(test.input)
			assert.Equal(t, result, test.expect)
		})
	}
}

func (suite *SelectorScopesSuite) TestWrapFilter() {
	table := []struct {
		name       string
		filter     filterFunc
		input      []string
		comparator int
		target     string
	}{
		{
			name:       "any",
			filter:     filters.Contains,
			input:      Any(),
			comparator: int(filters.Passes),
			target:     AnyTgt,
		},
		{
			name:       "none",
			filter:     filters.In,
			input:      None(),
			comparator: int(filters.Fails),
			target:     NoneTgt,
		},
		{
			name:       "something",
			filter:     filters.Equal,
			input:      []string{"userid"},
			comparator: int(filters.EqualTo),
			target:     "userid",
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ff := wrapFilter(test.filter)(test.input)
			assert.Equal(t, int(ff.Comparator), test.comparator)
			assert.Equal(t, ff.Target, test.target)
		})
	}
}
