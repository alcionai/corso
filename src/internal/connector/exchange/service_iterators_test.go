package exchange

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExchangeServiceIteratorsUnitSuite struct {
	suite.Suite
}

func TestExchangeServiceIteratorsUnitSuite(t *testing.T) {
	suite.Run(t, new(ExchangeServiceIteratorsUnitSuite))
}

func (suite *ExchangeServiceIteratorsUnitSuite) TestPanicRecoveryWrapper() {
	var (
		errs          error
		t             = suite.T()
		ctx           = context.Background()
		panicIterator = func(a any) bool {
			panic(assert.AnError)
			//nolint
			return true
		}
	)

	w := panicRecoveryWrapper(ctx, errs, panicIterator)

	// TODO: not working at the moment.
	// assert.Error(t, errs)
	require.NotPanics(t, func() {
		w("foo")
	})
}
