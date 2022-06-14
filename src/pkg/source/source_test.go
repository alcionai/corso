package source

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SourceSuite struct {
	suite.Suite
}

func TestSourceSuite(t *testing.T) {
	suite.Run(t, new(SourceSuite))
}

func (suite *SourceSuite) TestNewSource() {
	t := suite.T()
	s := newSource("tid", ServiceUnknown)
	assert.NotNil(t, s)
	assert.Equal(t, s.TenantID, "tid")
	assert.Equal(t, s.service, ServiceUnknown)
	assert.NotNil(t, s.scopes)
}

func (suite *SourceSuite) TestSource_Service() {
	table := []service{
		ServiceUnknown,
		ServiceExchange,
	}
	for _, test := range table {
		suite.T().Run(fmt.Sprintf("testing %d", test), func(t *testing.T) {
			s := newSource("tid", test)
			assert.Equal(t, s.Service(), test)
		})
	}
}

func (suite *SourceSuite) TestValAtoI() {
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
			result := valAtoI(m, "test")
			assert.Equal(t, result, test.expect)
		})
	}
}

func (suite *SourceSuite) TestBadCastErr() {
	err := BadCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err)
}
