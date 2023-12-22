package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type KeySetTestSuite struct {
	tester.Suite
}

func TestKeySetTestSuite(t *testing.T) {
	suite.Run(t, &KeySetTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *KeySetTestSuite) TestHasKey() {
	tests := []struct {
		name   string
		keySet Set
		key    string
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "key exists in the set",
			keySet: Set{"key1": {}, "key2": {}},
			key:    "key1",
			expect: assert.True,
		},
		{
			name:   "key does not exist in the set",
			keySet: Set{"key1": {}, "key2": {}},
			key:    "nonexistent",
			expect: assert.False,
		},
		{
			name:   "empty set",
			keySet: Set{},
			key:    "key",
			expect: assert.False,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			test.expect(suite.T(), test.keySet.HasKey(test.key))
		})
	}
}

func (suite *KeySetTestSuite) TestKeys() {
	tests := []struct {
		name   string
		keySet Set
		expect assert.ValueAssertionFunc
	}{
		{
			name:   "non-empty set",
			keySet: Set{"key1": {}, "key2": {}},
			expect: assert.NotEmpty,
		},
		{
			name:   "empty set",
			keySet: Set{},
			expect: assert.Empty,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			keys := test.keySet.Keys()
			test.expect(suite.T(), keys, []string{"key1", "key2"})
		})
	}
}

func (suite *KeySetTestSuite) TestHasKeys() {
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
