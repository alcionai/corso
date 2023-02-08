package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/pkg/path"
)

func TestStateOf(t *testing.T) {
	fooP, err := path.Builder{}.
		Append("foo").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	require.NoError(t, err)
	barP, err := path.Builder{}.
		Append("bar").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	require.NoError(t, err)

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
		t.Run(test.name, func(t *testing.T) {
			state := StateOf(test.prev, test.curr)
			assert.Equal(t, test.expect, state)
		})
	}
}
