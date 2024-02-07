package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/canario/src/internal/common/prefixmatcher"
)

var _ prefixmatcher.StringSetReader = &PrefixMap{}

type PrefixMap struct {
	prefixmatcher.StringSetBuilder
}

func NewPrefixMap(m map[string]map[string]struct{}) *PrefixMap {
	r := PrefixMap{StringSetBuilder: prefixmatcher.NewMatcher[map[string]struct{}]()}

	for k, v := range m {
		r.Add(k, v)
	}

	return &r
}

func (pm PrefixMap) AssertEqual(t *testing.T, r prefixmatcher.StringSetReader, description string) {
	if pm.Empty() {
		require.Truef(t, r.Empty(), "%s: result prefixMap should be empty but contains keys: %+v", description, r.Keys())
		return
	}

	pks := pm.Keys()
	rks := r.Keys()

	assert.ElementsMatchf(t, pks, rks, "%s: prefix keys match", description)

	for _, pk := range pks {
		p, _ := pm.Get(pk)
		r, _ := r.Get(pk)
		assert.Equal(t, p, r, description)
	}
}
