package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ConstsUnitSuite struct {
	tester.Suite
}

func TestConstsUnitSuite(t *testing.T) {
	suite.Run(t, &ConstsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ConstsUnitSuite) TestIsWithin() {
	table := []struct {
		name         string
		low, high, v int
		expect       assert.BoolAssertionFunc
	}{
		{"1 < 3 < 5", 1, 5, 3, assert.True},
		{"1 < 3, no high", 1, 0, 3, assert.True},
		{"1 <= 1 <= 1", 1, 1, 1, assert.True},
		{"1 <= 1 <= 5", 1, 5, 1, assert.True},
		{"1 <= 5 <= 5", 1, 5, 5, assert.True},
		{"1 <= 0 <= 2", 1, 1, 0, assert.False},
		{"1 <= 3 <= 2", 1, 1, 3, assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			test.expect(t, isWithin(test.low, test.high, test.v))
		})
	}
}
