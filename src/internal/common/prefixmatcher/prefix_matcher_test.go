package prefixmatcher_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/tester"
)

type PrefixMatcherUnitSuite struct {
	tester.Suite
}

func TestPrefixMatcherUnitSuite(t *testing.T) {
	suite.Run(t, &PrefixMatcherUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PrefixMatcherUnitSuite) TestEmpty() {
	pm := prefixmatcher.NewMatcher[string]()
	assert.True(suite.T(), pm.Empty())
}

func (suite *PrefixMatcherUnitSuite) TestAdd_Get() {
	t := suite.T()
	pm := prefixmatcher.NewMatcher[string]()
	kvs := map[string]string{
		"hello": "world",
		"hola":  "mundo",
		"foo":   "bar",
	}

	for k, v := range kvs {
		pm.Add(k, v)
	}

	for k, v := range kvs {
		val, ok := pm.Get(k)
		assert.True(t, ok, "searching for key", k)
		assert.Equal(t, v, val, "returned value")
	}
}

func (suite *PrefixMatcherUnitSuite) TestLongestPrefix() {
	key := "hello"
	value := "world"

	table := []struct {
		name          string
		inputKVs      map[string]string
		searchKey     string
		expectedKey   string
		expectedValue string
		expectedFound assert.BoolAssertionFunc
	}{
		{
			name: "Empty Prefix",
			inputKVs: map[string]string{
				"": value,
			},
			searchKey:     key,
			expectedKey:   "",
			expectedValue: value,
			expectedFound: assert.True,
		},
		{
			name: "Exact Match",
			inputKVs: map[string]string{
				key: value,
			},
			searchKey:     key,
			expectedKey:   key,
			expectedValue: value,
			expectedFound: assert.True,
		},
		{
			name: "Prefix Match",
			inputKVs: map[string]string{
				key[:len(key)-2]: value,
			},
			searchKey:     key,
			expectedKey:   key[:len(key)-2],
			expectedValue: value,
			expectedFound: assert.True,
		},
		{
			name: "Longest Prefix Match",
			inputKVs: map[string]string{
				key[:len(key)-2]: value,
				"":               value + "2",
				key[:len(key)-4]: value + "3",
			},
			searchKey:     key,
			expectedKey:   key[:len(key)-2],
			expectedValue: value,
			expectedFound: assert.True,
		},
		{
			name: "No Match",
			inputKVs: map[string]string{
				"foo": value,
			},
			searchKey:     key,
			expectedKey:   "",
			expectedValue: "",
			expectedFound: assert.False,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			pm := prefixmatcher.NewMatcher[string]()

			for k, v := range test.inputKVs {
				pm.Add(k, v)
			}

			k, v, ok := pm.LongestPrefix(test.searchKey)
			assert.Equal(t, test.expectedKey, k, "key")
			assert.Equal(t, test.expectedValue, v, "value")
			test.expectedFound(t, ok, "found")
		})
	}
}
