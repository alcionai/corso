package data

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/path"
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
			state := StateOf(test.prev, test.curr, count.New())
			assert.Equal(suite.T(), test.expect, state)
		})
	}
}

func (suite *CollectionSuite) TestNewBaseCollection() {
	fooP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "bar")
	require.NoError(suite.T(), err, clues.ToCore(err))
	preP, err := path.Build("_t", "_u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))

	loc := path.Builder{}.Append("foo")

	table := []struct {
		name       string
		current    path.Path
		previous   path.Path
		doNotMerge bool

		expectCurrent    path.Path
		expectPrev       path.Path
		expectState      CollectionState
		expectDoNotMerge bool
	}{
		{
			name:             "NotMoved DoNotMerge",
			current:          fooP,
			previous:         fooP,
			doNotMerge:       true,
			expectCurrent:    fooP,
			expectPrev:       fooP,
			expectState:      NotMovedState,
			expectDoNotMerge: true,
		},
		{
			name:          "Moved",
			current:       fooP,
			previous:      barP,
			expectCurrent: fooP,
			expectPrev:    barP,
			expectState:   MovedState,
		},
		{
			name:          "PrefixMoved",
			current:       fooP,
			previous:      preP,
			expectCurrent: fooP,
			expectPrev:    preP,
			expectState:   MovedState,
		},
		{
			name:          "New",
			current:       fooP,
			expectCurrent: fooP,
			expectState:   NewState,
		},
		{
			name:        "Deleted",
			previous:    fooP,
			expectPrev:  fooP,
			expectState: DeletedState,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			b := NewBaseCollection(
				test.current,
				test.previous,
				loc,
				control.Options{},
				test.doNotMerge,
				count.New())

			assert.Equal(t, test.expectCurrent, b.FullPath(), "full path")
			assert.Equal(t, test.expectPrev, b.PreviousPath(), "previous path")
			assert.Equal(t, loc, b.LocationPath(), "location path")
			assert.Equal(t, test.expectState, b.State(), "state")
			assert.Equal(t, test.expectDoNotMerge, b.DoNotMergeItems(), "do not merge")
			assert.Equal(t, path.EmailCategory, b.Category(), "category")
		})
	}
}

func (suite *CollectionSuite) TestNewTombstoneCollection() {
	t := suite.T()

	fooP, err := path.Build("t", "u", path.ExchangeService, path.EmailCategory, false, "foo")
	require.NoError(t, err, clues.ToCore(err))

	c := NewTombstoneCollection(fooP, control.Options{}, count.New())
	assert.Nil(t, c.FullPath(), "full path")
	assert.Equal(t, fooP, c.PreviousPath(), "previous path")
	assert.Nil(t, c.LocationPath(), "location path")
	assert.Equal(t, DeletedState, c.State(), "state")
	assert.False(t, c.DoNotMergeItems(), "do not merge")
}
