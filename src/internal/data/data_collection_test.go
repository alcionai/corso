package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/path"
)

type mockColl struct {
	p path.Path
}

func (mc mockColl) Items() <-chan Stream {
	return nil
}

func (mc mockColl) FullPath() path.Path {
	return mc.p
}

type CollectionSuite struct {
	suite.Suite
}

// ------------------------------------------------------------------------------------------------
// tests
// ------------------------------------------------------------------------------------------------

func TestCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionSuite))
}

func (suite *CollectionSuite) TestResourceOwnerSet() {
	t := suite.T()
	toColl := func(t *testing.T, resource string) Collection {
		p, err := path.Builder{}.
			Append("foo").
			ToDataLayerExchangePathForCategory("tid", resource, path.EventsCategory, false)
		require.NoError(t, err)

		return mockColl{p}
	}

	table := []struct {
		name   string
		input  []Collection
		expect []string
	}{
		{
			name:   "empty",
			input:  []Collection{},
			expect: []string{},
		},
		{
			name:   "nil",
			input:  nil,
			expect: []string{},
		},
		{
			name:   "single resource",
			input:  []Collection{toColl(t, "fnords")},
			expect: []string{"fnords"},
		},
		{
			name:   "multiple resource",
			input:  []Collection{toColl(t, "fnords"), toColl(t, "smarfs")},
			expect: []string{"fnords", "smarfs"},
		},
		{
			name:   "duplciate resources",
			input:  []Collection{toColl(t, "fnords"), toColl(t, "smarfs"), toColl(t, "fnords")},
			expect: []string{"fnords", "smarfs"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			rs := ResourceOwnerSet(test.input)
			assert.ElementsMatch(t, test.expect, rs)
		})
	}
}
