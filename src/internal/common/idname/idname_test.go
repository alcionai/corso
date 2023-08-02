package idname

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type IDNameUnitSuite struct {
	tester.Suite
}

func TestIDNameUnitSuite(t *testing.T) {
	suite.Run(t, &IDNameUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *IDNameUnitSuite) TestAdd() {
	table := []struct {
		name       string
		inID       string
		inName     string
		searchID   string
		searchName string
	}{
		{
			name:       "basic",
			inID:       "foo",
			inName:     "bar",
			searchID:   "foo",
			searchName: "bar",
		},
		{
			name:       "change casing",
			inID:       "FNORDS",
			inName:     "SMARF",
			searchID:   "fnords",
			searchName: "smarf",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cache := NewCache(nil)

			cache.Add(test.inID, test.inName)

			id, found := cache.IDOf(test.searchName)
			assert.True(t, found)
			assert.Equal(t, test.inID, id)

			name, found := cache.NameOf(test.searchID)
			assert.True(t, found)
			assert.Equal(t, test.inName, name)
		})
	}
}
