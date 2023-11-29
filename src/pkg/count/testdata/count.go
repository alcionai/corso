package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/pkg/count"
)

type Expected map[count.Key]int64

func (e Expected) Compare(
	t *testing.T,
	bus *count.Bus,
) {
	vs := bus.Values()
	results := map[count.Key]int64{}

	for k := range e {
		results[k] = bus.Get(k)
		delete(vs, string(k))
	}

	for k, v := range vs {
		t.Logf("unchecked count %q: %d", k, v)
	}

	assert.Equal(t, e, Expected(results))
}
