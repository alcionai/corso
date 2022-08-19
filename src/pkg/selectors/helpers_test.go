package selectors

import "github.com/alcionai/corso/pkg/backup/details"

// ---------------------------------------------------------------------------
// categorizers
// ---------------------------------------------------------------------------

// categorizer
type mockCategorizer int

const (
	unknownCatStub mockCategorizer = iota
	rootCatStub
	leafCatStub
)

var _ categorizer = unknownCatStub

func (sc mockCategorizer) String() string {
	switch sc {
	case leafCatStub:
		return "leaf"
	case rootCatStub:
		return "root"
	}
	return "unknown"
}

func (sc mockCategorizer) includesType(cat categorizer) bool {
	switch sc {
	case rootCatStub:
		return cat == rootCatStub
	case leafCatStub:
		return true
	}
	return false
}

func (sc mockCategorizer) pathValues(path []string) map[categorizer]string {
	return map[categorizer]string{rootCatStub: "stub"}
}

func (sc mockCategorizer) pathKeys() []categorizer {
	return []categorizer{rootCatStub, leafCatStub}
}

// TODO: Uncomment when reducer func is added
// func stubPathValues() map[categorizer]string {
// 	return map[categorizer]string{
// 		rootCatStub: rootCatStub.String(),
// 		leafCatStub: leafCatStub.String(),
// 	}
// }

// ---------------------------------------------------------------------------
// scopers
// ---------------------------------------------------------------------------

// scoper
type mockScope scope

var _ scoper = &mockScope{}

func (ms mockScope) categorizer() categorizer {
	switch ms[scopeKeyCategory] {
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
	return ms[shouldMatch] == "true"
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
		rootCatStub.String(): AnyTgt,
		scopeKeyCategory:     rootCatStub.String(),
		scopeKeyGranularity:  Item,
		scopeKeyResource:     stubResource,
		scopeKeyDataType:     rootCatStub.String(),
		shouldMatch:          sm,
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
