package selectors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
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

func (mc mockCategorizer) isUnion() bool {
	return mc == rootCatStub
}

func (mc mockCategorizer) isLeaf() bool {
	return mc == leafCatStub
}

func (mc mockCategorizer) pathValues(
	repo path.Path,
	ent details.Entry,
	cfg Config,
) (map[categorizer][]string, error) {
	return map[categorizer][]string{
		rootCatStub: {"root"},
		leafCatStub: {"leaf"},
	}, nil
}

func (mc mockCategorizer) pathKeys() []categorizer {
	return []categorizer{rootCatStub, leafCatStub}
}

func (mc mockCategorizer) PathType() path.CategoryType {
	switch mc {
	case leafCatStub:
		return path.EventsCategory
	default:
		return path.UnknownCategory
	}
}

func stubPathValues() map[categorizer][]string {
	return map[categorizer][]string{
		rootCatStub: {rootCatStub.String()},
		leafCatStub: {leafCatStub.String()},
	}
}

// ---------------------------------------------------------------------------
// scopers
// ---------------------------------------------------------------------------

// scoper
type mockScope scope

var _ scoper = &mockScope{}

func (s mockScope) categorizer() categorizer {
	switch s[scopeKeyCategory].Identity {
	case rootCatStub.String():
		return rootCatStub
	case leafCatStub.String():
		return leafCatStub
	}

	return unknownCatStub
}

func (s mockScope) matchesInfo(dii details.ItemInfo) bool {
	return s[shouldMatch].Target == "true"
}

func (s mockScope) setDefaults() {}

const (
	shouldMatch = "should-match-entry"
)

// helper funcs
func stubScope(match string) mockScope {
	sm := "true"
	if len(match) > 0 {
		sm = match
	}

	filt := passAny
	if match == "none" {
		filt = failAny
	}

	return mockScope{
		rootCatStub.String(): filt,
		leafCatStub.String(): filt,
		scopeKeyCategory:     filters.Identity(rootCatStub.String()),
		scopeKeyDataType:     filters.Identity(rootCatStub.String()),
		shouldMatch:          filters.Identity(sm),
	}
}

func stubInfoScope(match string) mockScope {
	sc := stubScope(match)
	sc[scopeKeyInfoCategory] = filters.Identity("true")

	return sc
}

func makeStubScope(cfg Config, match []string) mockScope {
	return makeScope[mockScope](leafCatStub, match, defaultItemOptions(cfg)...)
}

func (s mockScope) Matches(cat mockCategorizer, target string) bool {
	return matches(s, cat, target)
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

func (s mockScope) Conceal() string             { return conceal(s) }
func (s mockScope) Format(fs fmt.State, r rune) { format(s, fs, r) }
func (s mockScope) String() string              { return conceal(s) }
func (s mockScope) PlainString() string         { return plainString(s) }

// ---------------------------------------------------------------------------
// selectors
// ---------------------------------------------------------------------------

var _ servicerCategorizerProvider = &mockSel{}

type mockSel struct {
	Selector
}

func stubSelector(resourceOwners []string) mockSel {
	return mockSel{
		Selector: Selector{
			ResourceOwners: filterFor(scopeConfig{}, resourceOwners...),
			Service:        ServiceExchange,
			Excludes:       []scope{scope(stubScope(""))},
			Filters:        []scope{scope(stubScope(""))},
			Includes:       []scope{scope(stubScope(""))},
		},
	}
}

func (m mockSel) PathCategories() selectorPathCategories {
	return selectorPathCategories{
		Includes: []path.CategoryType{pathCatStub},
		Excludes: []path.CategoryType{pathCatStub},
		Filters:  []path.CategoryType{pathCatStub},
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

// calls assert.Equal(t, v, getCatValue(sc, k)[0]) on each k:v pair in the map
func scopeMustHave[T scopeT](t *testing.T, sc T, m map[categorizer][]string) {
	for k, vs := range m {
		t.Run(k.String(), func(t *testing.T) {
			assert.Equal(t, vs, getCatValue(sc, k), "Key: %s", k)
		})
	}
}

// stubPath ensures test path production matches that of fullPath design,
// stubbing out static values where necessary.
func stubPath(t *testing.T, user string, s []string, cat path.CategoryType) path.Path {
	pth, err := path.Build("tid", user, path.ExchangeService, cat, true, s...)
	require.NoError(t, err, clues.ToCore(err))

	return pth
}

// stubRepoRef ensures test path production matches that of repoRef design,
// stubbing out static values where necessary.
func stubRepoRef(service path.ServiceType, data path.CategoryType, resourceOwner, folders, item string) string {
	return strings.Join([]string{"tid", service.String(), resourceOwner, data.String(), folders, item}, "/")
}
