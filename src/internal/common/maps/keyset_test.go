package maps

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

func (suite *KeySetTestSuite) Test_HasKey() {
	tests := []struct {
		name   string
		keySet KeySet
		key    string
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "key exists in the set",
			keySet: KeySet{"key1": {}, "key2": {}},
			key:    "key1",
			expect: assert.True,
		},
		{
			name:   "key does not exist in the set",
			keySet: KeySet{"key1": {}, "key2": {}},
			key:    "nonexistent",
			expect: assert.False,
		},
		{
			name:   "empty set",
			keySet: KeySet{},
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

func (suite *KeySetTestSuite) Test_Keys() {
	tests := []struct {
		name   string
		keySet KeySet
		expect assert.ValueAssertionFunc
	}{
		{
			name:   "non-empty set",
			keySet: KeySet{"key1": {}, "key2": {}},
			expect: assert.NotEmpty,
		},
		{
			name:   "empty set",
			keySet: KeySet{},
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
