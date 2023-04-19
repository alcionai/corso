package selectors

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type SelectorScopesSuite struct {
	tester.Suite
}

func TestSelectorScopesSuite(t *testing.T) {
	suite.Run(t, &SelectorScopesSuite{Suite: tester.NewUnitSuite(t)})
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
				stub[rootCatStub.String()] = filters.Equal([]string{""})
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.False,
		},
		{
			name: "blank target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filterFor(scopeConfig{}, "fnords")
				return stub
			},
			check:  "",
			expect: assert.False,
		},
		{
			name: "matching target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filterFor(scopeConfig{}, rootCatStub.String())
				return stub
			},
			check:  rootCatStub.String(),
			expect: assert.True,
		},
		{
			name: "non-matching target",
			scope: func() mockScope {
				stub := stubScope("")
				stub[rootCatStub.String()] = filterFor(scopeConfig{}, rootCatStub.String())
				return stub
			},
			check:  "smarf",
			expect: assert.False,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			test.expect(
				t,
				matches(test.scope(), rootCatStub, test.check))
		})
	}
}

func (suite *SelectorScopesSuite) TestGetCatValue() {
	t := suite.T()

	stub := stubScope("")
	stub[rootCatStub.String()] = filterFor(scopeConfig{}, rootCatStub.String())

	assert.Equal(t,
		[]string{rootCatStub.String()},
		getCatValue(stub, rootCatStub))
	assert.Equal(t,
		None(),
		getCatValue(stub, mockCategorizer("foo")))
}

func (suite *SelectorScopesSuite) TestIsAnyTarget() {
	t := suite.T()
	stub := stubScope("")
	assert.True(t, isAnyTarget(stub, rootCatStub))
	assert.True(t, isAnyTarget(stub, leafCatStub))
	assert.False(t, isAnyTarget(stub, mockCategorizer("smarf")))

	stub = stubScope("none")
	assert.False(t, isAnyTarget(stub, rootCatStub))
	assert.False(t, isAnyTarget(stub, leafCatStub))
	assert.False(t, isAnyTarget(stub, mockCategorizer("smarf")))
}

var reduceTestTable = []struct {
	name               string
	sel                func() mockSel
	expectLen          int
	expectPassesReduce assert.BoolAssertionFunc
	expectPasses       assert.BoolAssertionFunc
}{
	{
		name: "include all resource owners",
		sel: func() mockSel {
			sel := stubSelector(Any())
			sel.Filters = nil
			sel.Excludes = nil
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "include all scopes",
		sel: func() mockSel {
			sel := stubSelector(Any())
			sel.Filters = nil
			sel.Excludes = nil
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "include none resource owners",
		sel: func() mockSel {
			sel := stubSelector(None())
			sel.Includes[0] = scope(stubScope(AnyTgt))
			sel.Filters = nil
			sel.Excludes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.True, // passes() does not check owners
	},
	{
		name: "include none scopes",
		sel: func() mockSel {
			sel := stubSelector(Any())
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
			sel := stubSelector(Any())
			sel.Excludes = nil
			return sel
		},
		expectLen:    1,
		expectPasses: assert.True,
	},
	{
		name: "include all filter none",
		sel: func() mockSel {
			sel := stubSelector(Any())
			sel.Filters[0] = scope(stubInfoScope("none"))
			sel.Excludes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "include all exclude all",
		sel: func() mockSel {
			sel := stubSelector(Any())
			sel.Filters = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "include all exclude none",
		sel: func() mockSel {
			sel := stubSelector(Any())
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
			sel := stubSelector(Any())
			sel.Includes = nil
			return sel
		},
		expectLen:    0,
		expectPasses: assert.False,
	},
	{
		name: "filter all exclude none",
		sel: func() mockSel {
			sel := stubSelector(Any())
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
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			errs := fault.New(true)

			ds := deets()
			result := reduce[mockScope](
				ctx,
				&ds,
				test.sel().Selector,
				dataCats,
				errs)
			require.NotNil(t, result)
			require.NoError(t, errs.Failure(), "no recoverable errors", clues.ToCore(errs.Failure()))
			assert.Len(t, result.Entries, test.expectLen)
		})
	}
}

func (suite *SelectorScopesSuite) TestReduce_locationRef() {
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
						LocationRef: "a/b/c//defg",
					},
				},
			},
		}
	}
	dataCats := map[path.CategoryType]mockCategorizer{
		pathCatStub: rootCatStub,
	}

	for _, test := range reduceTestTable {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			ds := deets()
			result := reduce[mockScope](
				ctx,
				&ds,
				test.sel().Selector,
				dataCats,
				fault.New(true))
			require.NotNil(t, result)
			assert.Len(t, result.Entries, test.expectLen)
		})
	}
}

func (suite *SelectorScopesSuite) TestScopesByCategory() {
	t := suite.T()
	s1 := stubScope("")
	s2 := stubScope("")
	s2[scopeKeyCategory] = filterFor(scopeConfig{}, unknownCatStub.String())
	result := scopesByCategory[mockScope](
		[]scope{scope(s1), scope(s2)},
		map[path.CategoryType]mockCategorizer{
			path.UnknownCategory: rootCatStub,
		},
		false)
	assert.Len(t, result, 1)
	assert.Len(t, result[rootCatStub], 2)
	assert.Empty(t, result[leafCatStub])
}

func (suite *SelectorScopesSuite) TestPasses() {
	var (
		cat   = rootCatStub
		pth   = stubPath(suite.T(), "uid", []string{"fld"}, path.EventsCategory)
		entry = details.DetailsEntry{
			RepoRef: pth.String(),
		}
	)

	pvs, err := cat.pathValues(pth, entry)
	require.NoError(suite.T(), err)

	for _, test := range reduceTestTable {
		suite.Run(test.name, func() {
			t := suite.T()

			sel := test.sel()
			excl := toMockScope(sel.Excludes)
			filt := toMockScope(sel.Filters)
			incl := toMockScope(sel.Includes)
			result := passes(
				cat,
				pvs,
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
	short := "brunheelda"

	table := []struct {
		name     string
		cat      mockCategorizer
		rootVal  string
		leafVal  string
		shortRef string
		expect   assert.BoolAssertionFunc
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
		{
			name:     "leaf matches shortRef",
			rootVal:  rootCatStub.String(),
			leafVal:  short,
			shortRef: short,
			expect:   assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			pvs := stubPathValues()
			pvs[leafCatStub] = append(pvs[leafCatStub], test.shortRef)

			sc := stubScope("")
			sc[rootCatStub.String()] = filterFor(scopeConfig{}, test.rootVal)
			sc[leafCatStub.String()] = filterFor(scopeConfig{}, test.leafVal)

			test.expect(t, matchesPathValues(sc, cat, pvs))
		})
	}
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
		suite.Run(test.name, func() {
			t := suite.T()

			result := clean(test.input)
			assert.Equal(t, result, test.expect)
		})
	}
}

func (suite *SelectorScopesSuite) TestScopeConfig() {
	input := "input"

	table := []struct {
		name   string
		config scopeConfig
		expect int
	}{
		{
			name:   "no configs set",
			config: scopeConfig{},
			expect: int(filters.EqualTo),
		},
		{
			name:   "force prefix",
			config: scopeConfig{usePrefixFilter: true},
			expect: int(filters.TargetPrefixes),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := filterFor(test.config, input)
			assert.Equal(t, test.expect, int(result.Comparator))
		})
	}
}

var _ fmt.State = &mockFMTState{}

type mockFMTState struct {
	w io.Writer
}

func (ms mockFMTState) Write(bs []byte) (int, error) { return ms.w.Write(bs) }
func (ms mockFMTState) Width() (int, bool)           { return 0, false }
func (ms mockFMTState) Precision() (int, bool)       { return 0, false }
func (ms mockFMTState) Flag(int) bool                { return false }

func (suite *SelectorScopesSuite) TestScopesPII() {
	table := []struct {
		name          string
		s             mockScope
		contains      []string
		containsPlain []string
	}{
		{
			name:          "empty",
			s:             mockScope{},
			contains:      []string{`{}`},
			containsPlain: []string{`{}`},
		},
		{
			name: "multiple filters",
			s: mockScope{
				"pass": filterFor(scopeConfig{}, "*"),
				"fail": filterFor(scopeConfig{}, ""),
				"foo":  filterFor(scopeConfig{}, "bar"),
				"qux":  filterFor(scopeConfig{}, "fnords", "smarf"),
			},
			contains: []string{
				`"pass":"Pass"`,
				`"fail":"Fail"`,
				`"foo":"EQ:bar"`,
				`"qux":"EQ:fnords,smarf"`,
			},
			containsPlain: []string{
				`"pass":"Pass"`,
				`"fail":"Fail"`,
				`"foo":"EQ:bar"`,
				`"qux":"EQ:fnords,smarf"`,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := conceal(test.s)
			for _, c := range test.contains {
				assert.Contains(t, result, c, "conceal")
			}

			result = plainString(test.s)
			for _, c := range test.containsPlain {
				assert.Contains(t, result, c, "plainString")
			}

			sb := &strings.Builder{}
			fs := mockFMTState{sb}

			format(test.s, &fs, 0)
			for _, c := range test.contains {
				assert.Contains(t, sb.String(), c, "conceal")
			}
		})
	}
}
