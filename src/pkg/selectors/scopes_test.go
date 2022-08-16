package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/backup/details"
)

// ---------------------------------------------------------------------------
// consts and mocks
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

func stubPathValues() map[categorizer]string {
	return map[categorizer]string{
		rootCatStub: rootCatStub.String(),
		leafCatStub: leafCatStub.String(),
	}
}

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
// tests
// ---------------------------------------------------------------------------

type SelectorScopesSuite struct {
	suite.Suite
}

func TestSelectorScopesSuite(t *testing.T) {
	suite.Run(t, new(SelectorScopesSuite))
}

// TODO: Uncomment when contains() is switched for the scopes.go version
//
// func (suite *SelectorScopesSuite) TestContains() {
// 	t := suite.T()
// 	// any
// 	stub := stubScope("")
// 	assert.True(t, contains(stub, rootCatStub, rootCatStub.String()), "any")
// 	// none
// 	stub[rootCatStub.String()] = NoneTgt
// 	assert.False(t, contains(stub, rootCatStub, rootCatStub.String()), "none")
// 	// missing values
// 	assert.False(t, contains(stub, rootCatStub, ""), "missing target")
// 	stub[rootCatStub.String()] = ""
// 	assert.False(t, contains(stub, rootCatStub, rootCatStub.String()), "missing scope value")
// 	// specific values
// 	stub[rootCatStub.String()] = rootCatStub.String()
// 	assert.True(t, contains(stub, rootCatStub, rootCatStub.String()), "matching value")
// 	assert.False(t, contains(stub, rootCatStub, "smarf"), "non-matching value")
// }

func (suite *SelectorScopesSuite) TestGetCatValue() {
	t := suite.T()
	stub := stubScope("")
	stub[rootCatStub.String()] = rootCatStub.String()
	assert.Equal(t, []string{rootCatStub.String()}, getCatValue(stub, rootCatStub))
	assert.Equal(t, None(), getCatValue(stub, leafCatStub))
}

func (suite *SelectorScopesSuite) TestGranularity() {
	t := suite.T()
	stub := stubScope("")
	assert.Equal(t, Item, granularity(stub))
}

func (suite *SelectorScopesSuite) TestIsAnyTarget() {
	t := suite.T()
	stub := stubScope("")
	assert.True(t, isAnyTarget(stub, rootCatStub))
	assert.False(t, isAnyTarget(stub, leafCatStub))
}
