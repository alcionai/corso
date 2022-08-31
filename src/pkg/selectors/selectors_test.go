package selectors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SelectorSuite struct {
	suite.Suite
}

func TestSelectorSuite(t *testing.T) {
	suite.Run(t, new(SelectorSuite))
}

func (suite *SelectorSuite) TestNewSelector() {
	t := suite.T()
	s := newSelector(ServiceUnknown)
	assert.NotNil(t, s)
	assert.Equal(t, s.Service, ServiceUnknown)
	assert.NotNil(t, s.Includes)
}

func (suite *SelectorSuite) TestBadCastErr() {
	err := badCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err)
}

func (suite *SelectorSuite) TestExistingDestinationErr() {
	err := existingDestinationErr("foo", "bar")
	assert.Error(suite.T(), err)
}

func (suite *SelectorSuite) TestPrintable() {
	t := suite.T()

	sel := stubSelector()
	p := sel.Printable()

	assert.Equal(t, sel.Service.String(), p.Service)
	assert.Equal(t, 1, len(p.Excludes))
	assert.Equal(t, 1, len(p.Filters))
	assert.Equal(t, 1, len(p.Includes))
}

func (suite *SelectorSuite) TestPrintable_IncludedResources() {
	t := suite.T()

	sel := stubSelector()
	p := sel.Printable()
	res := p.Resources()

	assert.Equal(t, stubResource, res, "resource should state only the stub")

	sel.Includes = []scope{
		scope(stubScope("")),
		{scopeKeyResource: filterize("smarf"), scopeKeyDataType: filterize(unknownCatStub.String())},
		{scopeKeyResource: filterize("smurf"), scopeKeyDataType: filterize(unknownCatStub.String())},
	}
	p = sel.Printable()
	res = p.Resources()

	assert.True(t, strings.HasSuffix(res, "(2 more)"), "resource '"+res+"' should have (2 more) suffix")

	p.Includes = nil
	res = p.Resources()

	assert.Equal(t, stubResource, res, "resource on filters should state only the stub")

	p.Filters = nil
	res = p.Resources()

	assert.Equal(t, "All", res, "resource with no Includes or Filters should state All")
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
				stubResource: {rootCatStub.String()},
			},
		},
		{
			name: "disjoint resources",
			input: []scope{
				scope(stubScope("")),
				{
					scopeKeyResource: filterize("smarf"),
					scopeKeyDataType: filterize(unknownCatStub.String()),
				},
			},
			expect: map[string][]string{
				stubResource: {rootCatStub.String()},
				"smarf":      {unknownCatStub.String()},
			},
		},
		{
			name: "disjoint types",
			input: []scope{
				scope(stubScope("")),
				{
					scopeKeyResource: filterize(stubResource),
					scopeKeyDataType: filterize("other"),
				},
			},
			expect: map[string][]string{
				stubResource: {rootCatStub.String(), "other"},
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			rtm := toResourceTypeMap(test.input)
			assert.Equal(t, test.expect, rtm)
		})
	}
}

func (suite *SelectorSuite) TestContains() {
	t := suite.T()
	key := rootCatStub
	target := "fnords"
	does := stubScope("")
	does[key.String()] = filterize(target)
	doesNot := stubScope("")
	doesNot[key.String()] = filterize("smarf")

	assert.True(t, matches(does, key, target), "does contain")
	assert.False(t, matches(doesNot, key, target), "does not contain")
}
