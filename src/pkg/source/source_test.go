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

func (suite *SourceSuite) TestBadCastErr() {
	err := BadCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err)
}
