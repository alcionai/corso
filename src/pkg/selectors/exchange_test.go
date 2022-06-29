package selectors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/selectors"
)

type ExchangeSourceSuite struct {
	suite.Suite
}

func TestExchangeSourceSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSourceSuite))
}

func (suite *ExchangeSourceSuite) TestNewExchangeSource() {
	t := suite.T()
	es := selectors.NewExchange("tid")
	assert.Equal(t, es.TenantID, "tid")
	assert.Equal(t, es.Service(), selectors.ServiceExchange)
	assert.NotZero(t, es.Scopes())
}
