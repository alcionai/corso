package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type SelectorScopesSuite struct {
	suite.Suite
}

func TestSelectorScopesSuite(t *testing.T) {
	suite.Run(t, new(SelectorScopesSuite))
}

func (suite *SelectorScopesSuite) TestContains() {
	t := suite.T()
	// any
	stub := stubScope("")
	assert.True(t, contains(stub, rootCatStub, rootCatStub.String()), "any")
	// none
	stub[rootCatStub.String()] = NoneTgt
	assert.False(t, contains(stub, rootCatStub, rootCatStub.String()), "none")
	// missing values
	assert.False(t, contains(stub, rootCatStub, ""), "missing target")
	stub[rootCatStub.String()] = ""
	assert.False(t, contains(stub, rootCatStub, rootCatStub.String()), "missing scope value")
	// specific values
	stub[rootCatStub.String()] = rootCatStub.String()
	assert.True(t, contains(stub, rootCatStub, rootCatStub.String()), "matching value")
	assert.False(t, contains(stub, rootCatStub, "smarf"), "non-matching value")
}

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
