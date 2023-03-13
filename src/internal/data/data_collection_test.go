package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type DataCollectionSuite struct {
	tester.Suite
}

func TestDataCollectionSuite(t *testing.T) {
	suite.Run(t, &DataCollectionSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DataCollectionSuite) TestStateOf() {
	fooP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(suite.T(), err)
	barP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "bar")
	require.NoError(suite.T(), err)

	table := []struct {
		name   string
		prev   path.Path
		curr   path.Path
		expect CollectionState
	}{
		{
			name:   "new",
			curr:   fooP,
			expect: NewState,
		},
		{
			name:   "not moved",
			prev:   fooP,
			curr:   fooP,
			expect: NotMovedState,
		},
		{
			name:   "moved",
			prev:   fooP,
			curr:   barP,
			expect: MovedState,
		},
		{
			name:   "deleted",
			prev:   fooP,
			expect: DeletedState,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			state := StateOf(test.prev, test.curr)
			assert.Equal(suite.T(), test.expect, state)
		})
	}
}
