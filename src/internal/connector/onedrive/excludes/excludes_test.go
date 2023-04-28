package excludes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/onedrive/excludes"
	"github.com/alcionai/corso/src/internal/tester"
)

type ParentsItemsUnitSuite struct {
	tester.Suite
}

func TestParentsItemsUnitSuite(t *testing.T) {
	suite.Run(t, &ParentsItemsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ParentsItemsUnitSuite) TestEmpty() {
	pm := excludes.NewParentsItems()
	assert.True(suite.T(), pm.Empty())
}

func (suite *ParentsItemsUnitSuite) TestAdd_Get() {
	t := suite.T()
	pm := excludes.NewParentsItems()
	kvs := map[string]map[string]struct{}{
		"hello": {"world": {}},
		"hola":  {"mundo": {}},
		"foo":   {"bar": {}},
	}

	for k, v := range kvs {
		pm.Add(k, v)
	}

	for k, v := range kvs {
		val, ok := pm.Get(k)
		assert.True(t, ok, "searching for key", k)
		assert.Equal(t, v, val, "returned value")
	}

	assert.ElementsMatch(t, maps.Keys(kvs), pm.Keys())
}

func (suite *ParentsItemsUnitSuite) TestAdd_Union() {
	t := suite.T()
	pm := excludes.NewParentsItems()
	pm.Add("hello", map[string]struct{}{
		"world": {},
		"mundo": {},
	})
	pm.Add("hello", map[string]struct{}{
		"goodbye": {},
		"aideu":   {},
	})

	expect := map[string]struct{}{
		"world":   {},
		"mundo":   {},
		"goodbye": {},
		"aideu":   {},
	}

	result, _ := pm.Get("hello")
	assert.Equal(t, expect, result)
	assert.ElementsMatch(t, []string{"hello"}, pm.Keys())
}

func (suite *ParentsItemsUnitSuite) TestLongestPrefix() {
	key := "hello"
	value := "world"

	table := []struct {
		name          string
		inputKVs      map[string]map[string]struct{}
		searchKey     string
		expectedKey   string
		expectedValue map[string]struct{}
		expectedFound assert.BoolAssertionFunc
	}{
		{
			name: "Empty Prefix",
			inputKVs: map[string]map[string]struct{}{
				"": {value: {}},
			},
			searchKey:     key,
			expectedKey:   "",
			expectedValue: map[string]struct{}{value: {}},
			expectedFound: assert.True,
		},
		{
			name: "Exact Match",
			inputKVs: map[string]map[string]struct{}{
				key: {value: {}},
			},
			searchKey:     key,
			expectedKey:   key,
			expectedValue: map[string]struct{}{value: {}},
			expectedFound: assert.True,
		},
		{
			name: "Prefix Match",
			inputKVs: map[string]map[string]struct{}{
				key[:len(key)-2]: {value: {}},
			},
			searchKey:     key,
			expectedKey:   key[:len(key)-2],
			expectedValue: map[string]struct{}{value: {}},
			expectedFound: assert.True,
		},
		{
			name: "Longest Prefix Match",
			inputKVs: map[string]map[string]struct{}{
				key[:len(key)-2]: {value: {}},
				"":               {value + "2": {}},
				key[:len(key)-4]: {value + "3": {}},
			},
			searchKey:     key,
			expectedKey:   key[:len(key)-2],
			expectedValue: map[string]struct{}{value: {}},
			expectedFound: assert.True,
		},
		{
			name: "No Match",
			inputKVs: map[string]map[string]struct{}{
				"foo": {value: {}},
			},
			searchKey:     key,
			expectedKey:   "",
			expectedValue: nil,
			expectedFound: assert.False,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			pm := excludes.NewParentsItems()

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
