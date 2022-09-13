package selectors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/filters"
)

// ---------------------------------------------------------------------------
// categorizers
// ---------------------------------------------------------------------------

// categorizer
type mockCategorizer string

const (
	unknownCatStub mockCategorizer = ""
	// wrap Exchange data here to get around path pkg assertions about path content.
	rootCatStub mockCategorizer = mockCategorizer(ExchangeUser)
	leafCatStub mockCategorizer = mockCategorizer(ExchangeEvent)

	pathServiceStub = path.ExchangeService
	pathCatStub     = path.EmailCategory
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

func (mc mockCategorizer) pathValues(pth path.Path) map[categorizer]string {
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
	shouldMatch = "should-match-entry"
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
		scopeKeyDataType:     filters.Identity(rootCatStub.String()),
		shouldMatch:          filters.Identity(sm),
	}
}

// ---------------------------------------------------------------------------
// selectors
// ---------------------------------------------------------------------------

type mockSel struct {
	Selector
}

func stubSelector() mockSel {
	return mockSel{
		Selector: Selector{
			Service:  ServiceExchange,
			Excludes: []scope{scope(stubScope(""))},
			Filters:  []scope{scope(stubScope(""))},
			Includes: []scope{scope(stubScope(""))},
		},
	}
}

func (s mockSel) Printable() Printable {
	return toPrintable[mockScope](s.Selector)
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
			assert.Equal(t, getCatValue(sc, k), split(v), "Key: %s", k)
		})
	}
}

// stubPath ensures test path production matches that of fullPath design,
// stubbing out static values where necessary.
func stubPath(t *testing.T, user string, s []string, cat path.CategoryType) path.Path {
	pth, err := path.Builder{}.
		Append(s...).
		ToDataLayerExchangePathForCategory("tid", user, cat, true)
	require.NoError(t, err)

	return pth
}

// stubRepoRef ensures test path production matches that of repoRef design,
// stubbing out static values where necessary.
func stubRepoRef(service path.ServiceType, data path.CategoryType, resourceOwner, folders, item string) string {
	return strings.Join([]string{"tid", service.String(), resourceOwner, data.String(), folders, item}, "/")
}
