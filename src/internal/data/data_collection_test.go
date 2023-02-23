package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/path"
)

type DataCollectionSuite struct {
	suite.Suite
}

func TestDataCollectionSuite(t *testing.T) {
	suite.Run(t, new(DataCollectionSuite))
}

func (suite *DataCollectionSuite) TestStateOf() {
	fooP, err := path.Builder{}.
		Append("foo").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	aw.MustNoErr(suite.T(), err)
	barP, err := path.Builder{}.
		Append("bar").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	aw.MustNoErr(suite.T(), err)

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
		suite.T().Run(test.name, func(t *testing.T) {
			state := StateOf(test.prev, test.curr)
			assert.Equal(t, test.expect, state)
		})
	}
}
