package selectors

import (
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
