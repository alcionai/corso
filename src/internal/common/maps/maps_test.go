package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type MapsUnitTestSuite struct {
	tester.Suite
}

func TestMapsUnitTestSuite(t *testing.T) {
	suite.Run(t, &MapsUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MapsUnitTestSuite) Test_HasKeys() {
	tests := []struct {
		name   string
		data   map[string]any
		keys   []string
		expect assert.BoolAssertionFunc
	}{
		{
			name: "has all keys",
			data: map[string]any{
				"key1": "data1",
				"key2": 2,
				"key3": struct{}{},
			},
			keys:   []string{"key1", "key2", "key3"},
			expect: assert.True,
		},
		{
			name: "has some keys",
			data: map[string]any{
				"key1": "data1",
				"key2": 2,
			},
			keys:   []string{"key1", "key2", "key3"},
			expect: assert.False,
		},
		{
			name: "has no key",
			data: map[string]any{
				"key1": "data1",
				"key2": 2,
			},
			keys:   []string{"key4", "key5", "key6"},
			expect: assert.False,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			test.expect(suite.T(), HasKeys(test.data, test.keys...))
		})
	}
}
