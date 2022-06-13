package source_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/source"
)

type ExchangeSourceSuite struct {
	suite.Suite
}

func TestExchangeSourceSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSourceSuite))
}

func (suite *ExchangeSourceSuite) TestNewExchangeSource() {
	t := suite.T()
	es := source.NewExchange("tid")
	assert.Equal(t, es.TenantID, "tid")
	assert.Equal(t, es.Service(), source.ServiceExchange)
	assert.NotZero(t, es.Scopes())
}
