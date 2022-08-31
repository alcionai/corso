package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/filters"
)

// ---------------------------------------------------------------------------
// categorizers
// ---------------------------------------------------------------------------

// categorizer
type mockCategorizer string

const (
	unknownCatStub mockCategorizer = ""
	rootCatStub    mockCategorizer = "rootCatStub"
	leafCatStub    mockCategorizer = "leafCatStub"
)

var _ categorizer = unknownCatStub

func (mc mockCategorizer) String() string {
	return string(mc)
}

func (mc mockCategorizer) leafCat() categorizer {
	return mc
}

func (mc mockCategorizer) rootCat() categorizer {
	return rootCatStub
}

func (mc mockCategorizer) unknownCat() categorizer {
	return unknownCatStub
}

func (mc mockCategorizer) pathValues(path []string) map[categorizer]string {
	return map[categorizer]string{rootCatStub: "stub"}
}

func (mc mockCategorizer) pathKeys() []categorizer {
	return []categorizer{rootCatStub, leafCatStub}
}

func stubPathValues() map[categorizer]string {
	return map[categorizer]string{
		rootCatStub: rootCatStub.String(),
		leafCatStub: leafCatStub.String(),
	}
}

// ---------------------------------------------------------------------------
// scopers
// ---------------------------------------------------------------------------

// scoper
type mockScope scope

var _ scoper = &mockScope{}

func (ms mockScope) categorizer() categorizer {
	switch ms[scopeKeyCategory].Target {
	case rootCatStub.String():
		return rootCatStub
	case leafCatStub.String():
		return leafCatStub
	}

	return unknownCatStub
}

func (ms mockScope) matchesEntry(
	cat categorizer,
	pathValues map[categorizer]string,
	entry details.DetailsEntry,
) bool {
	return ms[shouldMatch].Target == "true"
}

func (ms mockScope) setDefaults() {}

const (
	shouldMatch  = "should-match-entry"
	stubResource = "stubResource"
)

// helper funcs
func stubScope(match string) mockScope {
	sm := "true"
	if len(match) > 0 {
		sm = match
	}

	return mockScope{
		rootCatStub.String(): passAny,
		scopeKeyCategory:     filters.Identity(rootCatStub.String()),
		scopeKeyGranularity:  filters.Identity(Item),
		scopeKeyResource:     filters.Identity(stubResource),
		scopeKeyDataType:     filters.Identity(rootCatStub.String()),
		shouldMatch:          filters.Identity(sm),
	}
}

// ---------------------------------------------------------------------------
// selectors
// ---------------------------------------------------------------------------

func stubSelector() Selector {
	return Selector{
		Service:  ServiceExchange,
		Excludes: []scope{scope(stubScope(""))},
		Filters:  []scope{scope(stubScope(""))},
		Includes: []scope{scope(stubScope(""))},
	}
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

func setScopesToDefault[T scopeT](ts []T) []T {
	for _, s := range ts {
		s.setDefaults()
	}

	return ts
}

// calls assert.Equal(t, getCatValue(sc, k)[0], v) on each k:v pair in the map
func scopeMustHave[T scopeT](t *testing.T, sc T, m map[categorizer]string) {
	for k, v := range m {
		t.Run(k.String(), func(t *testing.T) {
			assert.Equal(t, getCatValue(sc, k), split(v))
		})
	}
}
