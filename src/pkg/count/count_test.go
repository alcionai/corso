package count_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/count"
)

type CountUnitSuite struct {
	tester.Suite
}

func TestCountUnitSuite(t *testing.T) {
	suite.Run(t, &CountUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CountUnitSuite) TestBus_Inc() {
	newParent := func() *count.Bus {
		parent := count.New()
		parent.Inc(count.Test)

		return parent
	}

	table := []struct {
		name        string
		skip        bool
		bus         *count.Bus
		expect      int64
		expectTotal int64
	}{
		{
			name:        "nil",
			bus:         nil,
			expect:      -1,
			expectTotal: -1,
		},
		{
			name:        "none",
			skip:        true,
			bus:         newParent().Local(),
			expect:      0,
			expectTotal: 1,
		},
		{
			name:        "one",
			bus:         newParent().Local(),
			expect:      1,
			expectTotal: 2,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			if !test.skip {
				test.bus.Inc(count.Test)
			}

			result := test.bus.Get(count.Test)
			assert.Equal(t, test.expect, result)

			resultTotal := test.bus.Total(count.Test)
			assert.Equal(t, test.expectTotal, resultTotal)
		})
	}
}

func (suite *CountUnitSuite) TestBus_Add() {
	newParent := func() *count.Bus {
		parent := count.New()
		parent.Add(count.Test, 2)

		return parent
	}

	table := []struct {
		name        string
		skip        bool
		bus         *count.Bus
		expect      int64
		expectTotal int64
	}{
		{
			name:        "nil",
			bus:         nil,
			expect:      -1,
			expectTotal: -1,
		},
		{
			name:        "none",
			skip:        true,
			bus:         newParent().Local(),
			expect:      0,
			expectTotal: 2,
		},
		{
			name:        "some",
			bus:         newParent().Local(),
			expect:      4,
			expectTotal: 6,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			if !test.skip {
				test.bus.Add(count.Test, 4)
			}

			result := test.bus.Get(count.Test)
			assert.Equal(t, test.expect, result)

			resultTotal := test.bus.Total(count.Test)
			assert.Equal(t, test.expectTotal, resultTotal)
		})
	}
}

func (suite *CountUnitSuite) TestBus_Values() {
	table := []struct {
		name        string
		bus         func() *count.Bus
		expect      map[string]int64
		expectTotal map[string]int64
	}{
		{
			name:        "nil",
			bus:         func() *count.Bus { return nil },
			expect:      map[string]int64{},
			expectTotal: map[string]int64{},
		},
		{
			name: "none",
			bus: func() *count.Bus {
				parent := count.New()
				parent.Add(count.Test, 2)

				l := parent.Local()

				return l
			},
			expect:      map[string]int64{},
			expectTotal: map[string]int64{string(count.Test): 2},
		},
		{
			name: "some",
			bus: func() *count.Bus {
				parent := count.New()
				parent.Add(count.Test, 2)

				l := parent.Local()
				l.Inc(count.Test)

				return l
			},
			expect:      map[string]int64{string(count.Test): 1},
			expectTotal: map[string]int64{string(count.Test): 3},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			b := test.bus()

			result := b.Values()
			assert.Equal(t, test.expect, result)

			resultTotal := b.TotalValues()
			assert.Equal(t, test.expectTotal, resultTotal)
		})
	}
}
