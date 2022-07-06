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

func (suite *ExchangeSourceSuite) TestExchangeSelector_ExcludeContacts() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		c1     = "c1"
		c2     = "c2"
	)

	sel.ExcludeContacts(user, folder, c1, c2)
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], folder)
	assert.Equal(t, scope[ExchangeContact.String()], join(c1, c2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_IncludeContacts() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		c1     = "c1"
		c2     = "c2"
	)

	sel.IncludeContacts(user, folder, c1, c2)
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], folder)
	assert.Equal(t, scope[ExchangeContact.String()], join(c1, c2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContact)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_ExcludeContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.ExcludeContactFolders(user, f1, f2)
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], join(f1, f2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_IncludeContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.IncludeContactFolders(user, f1, f2)
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], join(f1, f2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContactFolder)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_ExcludeEvents() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
	)

	sel.ExcludeEvents(user, e1, e2)
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeEvent.String()], join(e1, e2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_IncludeEvents() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
	)

	sel.IncludeEvents(user, e1, e2)
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeEvent.String()], join(e1, e2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeEvent)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_ExcludeMail() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		m1     = "m1"
		m2     = "m2"
	)

	sel.ExcludeMail(user, folder, m1, m2)
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], folder)
	assert.Equal(t, scope[ExchangeMail.String()], join(m1, m2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_IncludeMail() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		m1     = "m1"
		m2     = "m2"
	)

	sel.IncludeMail(user, folder, m1, m2)
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], folder)
	assert.Equal(t, scope[ExchangeMail.String()], join(m1, m2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMail)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_ExcludeMailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.ExcludeMailFolders(user, f1, f2)
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], join(f1, f2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_IncludeMailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.IncludeMailFolders(user, f1, f2)
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], join(f1, f2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMailFolder)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_ExcludeUsers() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.ExcludeUsers(u1, u2)
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], join(u1, u2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_IncludeUsers() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.IncludeUsers(u1, u2)
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], join(u1, u2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeUser)
}

func (suite *ExchangeSourceSuite) TestNewExchangeDestination() {
	t := suite.T()
	dest := NewExchangeDestination()
	assert.Len(t, dest, 0)
}

func (suite *ExchangeSourceSuite) TestExchangeDestination_Set() {
	dest := NewExchangeDestination()

	table := []exchangeCategory{
		ExchangeCategoryUnknown,
		ExchangeContact,
		ExchangeContactFolder,
		ExchangeEvent,
		ExchangeMail,
		ExchangeMailFolder,
		ExchangeUser,
	}
	for _, test := range table {
		suite.T().Run(test.String(), func(t *testing.T) {
			assert.NoError(t, dest.Set(test, "foo"))
			assert.Error(t, dest.Set(test, "foo"))
		})
	}

	assert.NoError(suite.T(), dest.Set(ExchangeUser, ""))
}

func (suite *ExchangeSourceSuite) TestExchangeDestination_GetOrDefault() {
	dest := NewExchangeDestination()

	table := []exchangeCategory{
		ExchangeCategoryUnknown,
		ExchangeContact,
		ExchangeContactFolder,
		ExchangeEvent,
		ExchangeMail,
		ExchangeMailFolder,
		ExchangeUser,
	}
	for _, test := range table {
		suite.T().Run(test.String(), func(t *testing.T) {
			assert.Equal(t, "bar", dest.GetOrDefault(test, "bar"))
			assert.NoError(t, dest.Set(test, "foo"))
			assert.Equal(t, "foo", dest.GetOrDefault(test, "bar"))
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
