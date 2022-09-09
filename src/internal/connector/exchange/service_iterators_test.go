package exchange

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExchangeServiceIteratorsUnitSuite struct {
	suite.Suite
}

func TestExchangeServiceIteratorsUnitSuite(t *testing.T) {
	suite.Run(t, new(ExchangeServiceIteratorsUnitSuite))
}

func (suite *ExchangeServiceIteratorsUnitSuite) TestPanicRecoveryWrapper() {
	ctx := context.Background()
	recoverPanic := func() {
		defer iteratorPanicRecovery(ctx)
		panic(assert.AnError)
	}

	// this test shouldn't panic.
	// unfortunately, assert.NotPanics() will fail if a panic occurs,
	// even if we recover from it.
	recoverPanic()
}
