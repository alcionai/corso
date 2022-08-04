package selectors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	user = "me@my.onmicrosoft.com"
)

var (
	dataType = ExchangeEvent.String()
)

func stubScope() map[string]string {
	return map[string]string{
		ExchangeEvent.String(): AnyTgt,
		ExchangeUser.String():  user,
		scopeKeyCategory:       dataType,
		scopeKeyGranularity:    Group,
		scopeKeyResource:       user,
		scopeKeyDataType:       dataType,
	}
}

func stubSelector() Selector {
	return Selector{
		Service:  ServiceExchange,
		Excludes: []map[string]string{stubScope()},
		Filters:  []map[string]string{stubScope()},
		Includes: []map[string]string{stubScope()},
	}
}

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

	assert.Equal(t, user, res, "resource should state only the user")

	sel.Includes = []map[string]string{
		stubScope(),
		{scopeKeyResource: "smarf", scopeKeyDataType: dataType},
		{scopeKeyResource: "smurf", scopeKeyDataType: dataType}}
	p = sel.Printable()
	res = p.Resources()

	assert.True(t, strings.HasSuffix(res, "(2 more)"), "resource '"+res+"' should have (2 more) suffix")

	p.Includes = nil
	res = p.Resources()

	assert.Equal(t, user, res, "resource on filters should state only the user")

	p.Filters = nil
	res = p.Resources()

	assert.Equal(t, "All", res, "resource with no Includes or Filters should state All")
}

func (suite *SelectorSuite) TestToResourceTypeMap() {
	table := []struct {
		name   string
		input  []map[string]string
		expect map[string][]string
	}{
		{
			name:  "single scope",
			input: []map[string]string{stubScope()},
			expect: map[string][]string{
				user: {dataType},
			},
		},
		{
			name: "disjoint resources",
			input: []map[string]string{
				stubScope(),
				{
					scopeKeyResource: "smarf",
					scopeKeyDataType: dataType,
				},
			},
			expect: map[string][]string{
				user:    {dataType},
				"smarf": {dataType},
			},
		},
		{
			name: "disjoint types",
			input: []map[string]string{
				stubScope(),
				{
					scopeKeyResource: user,
					scopeKeyDataType: "other",
				},
			},
			expect: map[string][]string{
				user: {dataType, "other"},
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
	key := "key"
	target := "fnords"
	does := map[string]string{key: target}
	doesNot := map[string]string{key: "smarf"}
	assert.True(t, contains(does, key, target))
	assert.False(t, contains(doesNot, key, target))
}
