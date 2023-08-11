package data

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type CollectionSuite struct {
	tester.Suite
}

func TestDataCollectionSuite(t *testing.T) {
	suite.Run(t, &CollectionSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionSuite) TestStateOf() {
	fooP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "bar")
	require.NoError(suite.T(), err, clues.ToCore(err))
	preP, err := path.Build("_t", "_u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))

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
			name:   "moved if prefix changes",
			prev:   fooP,
			curr:   preP,
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
