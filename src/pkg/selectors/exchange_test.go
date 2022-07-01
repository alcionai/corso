package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExchangeSourceSuite struct {
	suite.Suite
}

func TestExchangeSourceSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSourceSuite))
}

func (suite *ExchangeSourceSuite) TestNewExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup()
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.Zero(t, eb.RestorePointID)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestToExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup()
	s := eb.Selector
	eb, err := s.ToExchangeBackup()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.Zero(t, eb.RestorePointID)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestNewExchangeRestore() {
	t := suite.T()
	er := NewExchangeRestore("rpid")
	assert.Equal(t, er.Service, ServiceExchange)
	assert.Equal(t, er.RestorePointID, "rpid")
	assert.NotZero(t, er.Scopes())
}

func (suite *ExchangeSourceSuite) TestToExchangeRestore() {
	t := suite.T()
	eb := NewExchangeRestore("rpid")
	s := eb.Selector
	eb, err := s.ToExchangeRestore()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.Equal(t, eb.RestorePointID, "rpid")
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestNewExchangeDestination() {
	t := suite.T()
	dest := NewExchangeDestination()
	assert.Len(t, dest, 0)
}

func (suite *ExchangeSourceSuite) TestExchangeDestination_Set() {
	dest := NewExchangeDestination()

	table := []struct {
		cat   exchangeCategory
		check assert.ErrorAssertionFunc
	}{
		{ExchangeCategoryUnknown, assert.NoError},
		{ExchangeCategoryUnknown, assert.Error},
		{ExchangeContact, assert.NoError},
		{ExchangeContact, assert.Error},
		{ExchangeContactFolder, assert.NoError},
		{ExchangeContactFolder, assert.Error},
		{ExchangeEvent, assert.NoError},
		{ExchangeEvent, assert.Error},
		{ExchangeMail, assert.NoError},
		{ExchangeMail, assert.Error},
		{ExchangeMailFolder, assert.NoError},
		{ExchangeMailFolder, assert.Error},
		{ExchangeUser, assert.NoError},
		{ExchangeUser, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			test.check(t, dest.Set(test.cat, "foo"))
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeDestination_Get() {
	dest := NewExchangeDestination()

	table := []struct {
		cat   exchangeCategory
		check assert.ErrorAssertionFunc
	}{
		{ExchangeCategoryUnknown, assert.NoError},
		{ExchangeContact, assert.NoError},
		{ExchangeContactFolder, assert.NoError},
		{ExchangeEvent, assert.NoError},
		{ExchangeMail, assert.NoError},
		{ExchangeMailFolder, assert.NoError},
		{ExchangeUser, assert.NoError},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, "bar", dest.Get(test.cat, "bar"))
			test.check(t, dest.Set(test.cat, "foo"))
			assert.Equal(t, "foo", dest.Get(test.cat, "bar"))
		})
	}
}

var allScopesExceptUnknown = map[string]string{
	ExchangeContact.String():       All,
	ExchangeContactFolder.String(): All,
	ExchangeEvent.String():         All,
	ExchangeMail.String():          All,
	ExchangeMailFolder.String():    All,
	ExchangeUser.String():          All,
}

func (suite *ExchangeSourceSuite) TestExchangeBackup_Scopes() {
	eb := NewExchangeBackup()
	eb.Includes = []map[string]string{allScopesExceptUnknown}
	// todo: swap the above for this
	// eb := NewExchangeBackup().IncludeUsers(All)

	scopes := eb.Scopes()
	assert.Len(suite.T(), scopes, 1)
	assert.Equal(
		suite.T(),
		allScopesExceptUnknown,
		map[string]string(scopes[0]))
}

func (suite *ExchangeSourceSuite) TestExchangeScope_Category() {
	table := []struct {
		is     exchangeCategory
		expect exchangeCategory
		check  assert.ComparisonAssertionFunc
	}{
		{ExchangeCategoryUnknown, ExchangeCategoryUnknown, assert.Equal},
		{ExchangeCategoryUnknown, ExchangeUser, assert.NotEqual},
		{ExchangeContact, ExchangeContact, assert.Equal},
		{ExchangeContact, ExchangeMailFolder, assert.NotEqual},
		{ExchangeContactFolder, ExchangeContactFolder, assert.Equal},
		{ExchangeContactFolder, ExchangeMailFolder, assert.NotEqual},
		{ExchangeEvent, ExchangeEvent, assert.Equal},
		{ExchangeEvent, ExchangeContact, assert.NotEqual},
		{ExchangeMail, ExchangeMail, assert.Equal},
		{ExchangeMail, ExchangeMailFolder, assert.NotEqual},
		{ExchangeMailFolder, ExchangeMailFolder, assert.Equal},
		{ExchangeMailFolder, ExchangeContactFolder, assert.NotEqual},
		{ExchangeUser, ExchangeUser, assert.Equal},
		{ExchangeUser, ExchangeCategoryUnknown, assert.NotEqual},
	}
	for _, test := range table {
		suite.T().Run(test.is.String()+test.expect.String(), func(t *testing.T) {
			eb := NewExchangeBackup()
			eb.Includes = []map[string]string{{scopeKeyCategory: test.is.String()}}
			scope := eb.Scopes()[0]
			test.check(t, test.expect, scope.Category())
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScope_IncludesCategory() {
	table := []struct {
		is     exchangeCategory
		expect exchangeCategory
		check  assert.BoolAssertionFunc
	}{
		{ExchangeCategoryUnknown, ExchangeCategoryUnknown, assert.False},
		{ExchangeCategoryUnknown, ExchangeUser, assert.False},
		{ExchangeContact, ExchangeContactFolder, assert.True},
		{ExchangeContact, ExchangeMailFolder, assert.False},
		{ExchangeContactFolder, ExchangeContact, assert.True},
		{ExchangeContactFolder, ExchangeMailFolder, assert.False},
		{ExchangeEvent, ExchangeUser, assert.True},
		{ExchangeEvent, ExchangeContact, assert.False},
		{ExchangeMail, ExchangeMailFolder, assert.True},
		{ExchangeMail, ExchangeContact, assert.False},
		{ExchangeMailFolder, ExchangeMail, assert.True},
		{ExchangeMailFolder, ExchangeContactFolder, assert.False},
		{ExchangeUser, ExchangeUser, assert.True},
		{ExchangeUser, ExchangeCategoryUnknown, assert.False},
		{ExchangeUser, ExchangeMail, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.is.String()+test.expect.String(), func(t *testing.T) {
			eb := NewExchangeBackup()
			eb.Includes = []map[string]string{{scopeKeyCategory: test.is.String()}}
			scope := eb.Scopes()[0]
			test.check(t, scope.IncludesCategory(test.expect))
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScope_Get() {
	eb := NewExchangeBackup()
	eb.Includes = []map[string]string{allScopesExceptUnknown}
	// todo: swap the above for this
	// eb := NewExchangeBackup().IncludeUsers(All)

	scope := eb.Scopes()[0]

	table := []exchangeCategory{
		ExchangeContact,
		ExchangeContactFolder,
		ExchangeEvent,
		ExchangeMail,
		ExchangeMailFolder,
		ExchangeUser,
	}

	assert.Equal(
		suite.T(),
		[]string{None},
		scope.Get(ExchangeCategoryUnknown))

	expect := []string{All}
	for _, test := range table {
		suite.T().Run(test.String(), func(t *testing.T) {
			assert.Equal(t, expect, scope.Get(test))
		})
	}
}
