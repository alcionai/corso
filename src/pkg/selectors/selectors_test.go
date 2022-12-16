package selectors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

type SelectorSuite struct {
	suite.Suite
}

func TestSelectorSuite(t *testing.T) {
	suite.Run(t, new(SelectorSuite))
}

func (suite *SelectorSuite) TestNewSelector() {
	t := suite.T()
	s := newSelector(ServiceUnknown, Any())
	assert.NotNil(t, s)
	assert.Equal(t, s.Service, ServiceUnknown)
	assert.NotNil(t, s.Includes)
}

func (suite *SelectorSuite) TestBadCastErr() {
	err := badCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err)
}

func (suite *SelectorSuite) TestPrintable() {
	t := suite.T()

	sel := stubSelector(Any())
	p := sel.Printable()

	assert.Equal(t, sel.Service.String(), p.Service)
	assert.Equal(t, 1, len(p.Excludes))
	assert.Equal(t, 1, len(p.Filters))
	assert.Equal(t, 1, len(p.Includes))
}

func (suite *SelectorSuite) TestPrintable_IncludedResources() {
	table := []struct {
		name           string
		resourceOwners []string
		expect         func(string) bool
		reason         string
	}{
		{
			name:           "distinct",
			resourceOwners: []string{"foo", "smarf", "fnords"},
			expect: func(s string) bool {
				return strings.HasSuffix(s, "(2 more)")
			},
			reason: "should end with (2 more)",
		},
		{
			name:           "distinct",
			resourceOwners: nil,
			expect: func(s string) bool {
				return strings.HasSuffix(s, "None")
			},
			reason: "no resource owners should produce None",
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := stubSelector(test.resourceOwners)
			p := sel.Printable()
			res := p.Resources()

			assert.Equal(t, "All", res, "stub starts out as an all-pass")

			stubWithResource := func(resource string) scope {
				ss := stubScope("")
				ss[rootCatStub.String()] = filterize(scopeConfig{}, resource)

				return scope(ss)
			}

			sel.Includes = []scope{}
			sel.Filters = []scope{}

			for _, ro := range test.resourceOwners {
				sel.Includes = append(sel.Includes, stubWithResource(ro))
				sel.Filters = append(sel.Filters, stubWithResource(ro))
			}

			p = sel.Printable()
			res = p.Resources()

			assert.True(t, test.expect(res), test.reason)
		})
	}
}

func (suite *SelectorSuite) TestToResourceTypeMap() {
	table := []struct {
		name   string
		input  []scope
		expect map[string][]string
	}{
		{
			name:  "single scope",
			input: []scope{scope(stubScope(""))},
			expect: map[string][]string{
				"All": {rootCatStub.String()},
			},
		},
		{
			name: "disjoint resources",
			input: []scope{
				scope(stubScope("")),
				{
					rootCatStub.String(): filterize(scopeConfig{}, "smarf"),
					scopeKeyDataType:     filterize(scopeConfig{}, unknownCatStub.String()),
				},
			},
			expect: map[string][]string{
				"All":   {rootCatStub.String()},
				"smarf": {unknownCatStub.String()},
			},
		},
		{
			name: "multiple resources",
			input: []scope{
				scope(stubScope("")),
				{
					rootCatStub.String(): filterize(scopeConfig{}, join("smarf", "fnords")),
					scopeKeyDataType:     filterize(scopeConfig{}, unknownCatStub.String()),
				},
			},
			expect: map[string][]string{
				"All":    {rootCatStub.String()},
				"smarf":  {unknownCatStub.String()},
				"fnords": {unknownCatStub.String()},
			},
		},
		{
			name: "disjoint types",
			input: []scope{
				scope(stubScope("")),
				{
					rootCatStub.String(): filterize(scopeConfig{}, AnyTgt),
					scopeKeyDataType:     filterize(scopeConfig{}, "other"),
				},
			},
			expect: map[string][]string{
				"All": {rootCatStub.String(), "other"},
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			rtm := toResourceTypeMap[mockScope](test.input)
			assert.Equal(t, test.expect, rtm)
		})
	}
}

func (suite *SelectorSuite) TestResourceOwnersIn() {
	rootCat := rootCatStub.String()

	table := []struct {
		name   string
		input  []scope
		expect []string
	}{
		{
			name:   "nil",
			input:  nil,
			expect: []string{},
		},
		{
			name:   "empty",
			input:  []scope{},
			expect: []string{},
		},
		{
			name:   "single",
			input:  []scope{{rootCat: filters.Identity("foo")}},
			expect: []string{"foo"},
		},
		{
			name:   "multiple values",
			input:  []scope{{rootCat: filters.Identity(join("foo", "bar"))}},
			expect: []string{"foo", "bar"},
		},
		{
			name:   "with any",
			input:  []scope{{rootCat: filters.Identity(join("foo", "bar", AnyTgt))}},
			expect: []string{"foo", "bar"},
		},
		{
			name:   "with none",
			input:  []scope{{rootCat: filters.Identity(join("foo", "bar", NoneTgt))}},
			expect: []string{"foo", "bar"},
		},
		{
			name: "multiple scopes",
			input: []scope{
				{rootCat: filters.Identity(join("foo", "bar"))},
				{rootCat: filters.Identity(join("baz"))},
			},
			expect: []string{"foo", "bar", "baz"},
		},
		{
			name: "multiple scopes with duplicates",
			input: []scope{
				{rootCat: filters.Identity(join("foo", "bar"))},
				{rootCat: filters.Identity(join("baz", "foo"))},
			},
			expect: []string{"foo", "bar", "baz"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := resourceOwnersIn(test.input, rootCat)
			assert.ElementsMatch(t, test.expect, result)
		})
	}
}

func (suite *SelectorSuite) TestPathCategoriesIn() {
	leafCat := leafCatStub.String()
	f := filters.Identity(leafCat)

	table := []struct {
		name   string
		input  []scope
		expect []path.CategoryType
	}{
		{
			name:   "nil",
			input:  nil,
			expect: []path.CategoryType{},
		},
		{
			name:   "empty",
			input:  []scope{},
			expect: []path.CategoryType{},
		},
		{
			name:   "single",
			input:  []scope{{leafCat: f, scopeKeyCategory: f}},
			expect: []path.CategoryType{leafCatStub.PathType()},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := pathCategoriesIn[mockScope, mockCategorizer](test.input)
			assert.ElementsMatch(t, test.expect, result)
		})
	}
}

func (suite *SelectorSuite) TestContains() {
	t := suite.T()
	key := rootCatStub
	target := "fnords"
	does := stubScope("")
	does[key.String()] = filterize(scopeConfig{}, target)
	doesNot := stubScope("")
	doesNot[key.String()] = filterize(scopeConfig{}, "smarf")

	assert.True(t, matches(does, key, target), "does contain")
	assert.False(t, matches(doesNot, key, target), "does not contain")
}

func (suite *SelectorSuite) TestIsAnyResourceOwner() {
	t := suite.T()
	assert.False(t, isAnyResourceOwner(newSelector(ServiceUnknown, []string{"foo"})))
	assert.False(t, isAnyResourceOwner(newSelector(ServiceUnknown, []string{})))
	assert.False(t, isAnyResourceOwner(newSelector(ServiceUnknown, nil)))
	assert.True(t, isAnyResourceOwner(newSelector(ServiceUnknown, []string{AnyTgt})))
	assert.True(t, isAnyResourceOwner(newSelector(ServiceUnknown, Any())))
}

func (suite *SelectorSuite) TestIsNoneResourceOwner() {
	t := suite.T()
	assert.False(t, isNoneResourceOwner(newSelector(ServiceUnknown, []string{"foo"})))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, []string{})))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, nil)))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, []string{NoneTgt})))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, None())))
}

func (suite *SelectorSuite) TestSplitByResourceOnwer() {
	allOwners := []string{"foo", "bar", "baz", "qux"}

	table := []struct {
		name           string
		input          []string
		expectLen      int
		expectDiscrete []string
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []string{},
		},
		{
			name:  "noneTgt",
			input: []string{NoneTgt},
		},
		{
			name:  "none",
			input: None(),
		},
		{
			name:           "AnyTgt",
			input:          []string{AnyTgt},
			expectLen:      len(allOwners),
			expectDiscrete: allOwners,
		},
		{
			name:           "Any",
			input:          Any(),
			expectLen:      len(allOwners),
			expectDiscrete: allOwners,
		},
		{
			name:           "one owner",
			input:          []string{"fnord"},
			expectLen:      1,
			expectDiscrete: []string{"fnord"},
		},
		{
			name:           "two owners",
			input:          []string{"fnord", "smarf"},
			expectLen:      2,
			expectDiscrete: []string{"fnord", "smarf"},
		},
		{
			name:  "two owners and NoneTgt",
			input: []string{"fnord", "smarf", NoneTgt},
		},
		{
			name:           "two owners and AnyTgt",
			input:          []string{"fnord", "smarf", AnyTgt},
			expectLen:      len(allOwners),
			expectDiscrete: allOwners,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			s := newSelector(ServiceUnknown, test.input)
			result := splitByResourceOwner[mockScope](s, allOwners, rootCatStub)

			assert.Len(t, result, test.expectLen)

			for _, expect := range test.expectDiscrete {
				var found bool

				for _, sel := range result {
					if sel.DiscreteOwner == expect {
						found = true
						break
					}
				}

				assert.Truef(t, found, "%s in list of discrete owners", expect)
			}
		})
	}
}
