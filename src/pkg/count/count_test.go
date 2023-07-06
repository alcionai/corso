package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type CountUnitSuite struct {
	tester.Suite
}

func TestCountUnitSuite(t *testing.T) {
	suite.Run(t, &CountUnitSuite{Suite: tester.NewUnitSuite(t)})
}

const testKey = key("just-for-testing")

func (suite *CountUnitSuite) TestBus_Inc() {
	newParent := func() *Bus {
		parent := New()
		parent.Inc(testKey)

		return parent
	}

	table := []struct {
		name        string
		skip        bool
		bus         *Bus
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
				test.bus.Inc(testKey)
			}

			result := test.bus.Get(testKey)
			assert.Equal(t, test.expect, result)

			resultTotal := test.bus.Total(testKey)
			assert.Equal(t, test.expectTotal, resultTotal)
		})
	}
}

func (suite *CountUnitSuite) TestBus_Add() {
	newParent := func() *Bus {
		parent := New()
		parent.Add(testKey, 2)

		return parent
	}

	table := []struct {
		name        string
		skip        bool
		bus         *Bus
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
				test.bus.Add(testKey, 4)
			}

			result := test.bus.Get(testKey)
			assert.Equal(t, test.expect, result)

			resultTotal := test.bus.Total(testKey)
			assert.Equal(t, test.expectTotal, resultTotal)
		})
	}
}

func (suite *CountUnitSuite) TestBus_Values() {
	table := []struct {
		name        string
		bus         func() *Bus
		expect      map[string]int64
		expectTotal map[string]int64
	}{
		{
			name:        "nil",
			bus:         func() *Bus { return nil },
			expect:      map[string]int64{},
			expectTotal: map[string]int64{},
		},
		{
			name: "none",
			bus: func() *Bus {
				parent := New()
				parent.Add(testKey, 2)

				l := parent.Local()

				return l
			},
			expect:      map[string]int64{},
			expectTotal: map[string]int64{string(testKey): 2},
		},
		{
			name: "some",
			bus: func() *Bus {
				parent := New()
				parent.Add(testKey, 2)

				l := parent.Local()
				l.Inc(testKey)

				return l
			},
			expect:      map[string]int64{string(testKey): 1},
			expectTotal: map[string]int64{string(testKey): 3},
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
