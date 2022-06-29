package selectors

import (
	"fmt"
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
	s := newSelector("tid", ServiceUnknown)
	assert.NotNil(t, s)
	assert.Equal(t, s.TenantID, "tid")
	assert.Equal(t, s.service, ServiceUnknown)
	assert.NotNil(t, s.scopes)
}

func (suite *SelectorSuite) TestSelector_Service() {
	table := []service{
		ServiceUnknown,
		ServiceExchange,
	}
	for _, test := range table {
		suite.T().Run(fmt.Sprintf("testing %d", test), func(t *testing.T) {
			s := newSelector("tid", test)
			assert.Equal(t, s.Service(), test)
		})
	}
}

func (suite *SelectorSuite) TestGetIota() {
	table := []struct {
		name   string
		val    string
		expect int
	}{
		{"zero", "0", 0},
		{"positive", "1", 1},
		{"negative", "-1", -1},
		{"empty", "", 0},
		{"NaN", "fnords", 0},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			m := map[string]string{"test": test.val}
			result := getIota(m, "test")
			assert.Equal(t, result, test.expect)
		})
	}
}

func (suite *SelectorSuite) TestBadCastErr() {
	err := badCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err)
}
