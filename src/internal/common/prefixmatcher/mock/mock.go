package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
)

var _ prefixmatcher.MapReader = &PrefixMap{}

type PrefixMap struct {
	prefixmatcher.MapBuilder
}

func NewPrefixMap(m map[string]map[string]struct{}) *PrefixMap {
	r := PrefixMap{MapBuilder: prefixmatcher.NewBuilder[map[string]struct{}]()}

	for k, v := range m {
		r.Add(k, v)
	}

	return &r
}

func (pm PrefixMap) AssertEqual(t *testing.T, r prefixmatcher.MapReader) {
	if pm.Empty() {
		require.True(t, r.Empty(), "both prefix maps are empty")
		return
	}

	pks := pm.Keys()
	rks := r.Keys()

	assert.ElementsMatch(t, pks, rks, "prefix keys match")

	for _, pk := range pks {
		p, _ := pm.Get(pk)
		r, _ := r.Get(pk)
		assert.Equal(t, p, r, "values match")
	}
}
