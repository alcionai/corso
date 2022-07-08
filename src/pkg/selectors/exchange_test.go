package selectors

import (
	"testing"

	"github.com/alcionai/corso/pkg/restorepoint"
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
	assert.Zero(t, eb.BackupID)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestToExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup()
	s := eb.Selector
	eb, err := s.ToExchangeBackup()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.Zero(t, eb.BackupID)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestNewExchangeRestore() {
	t := suite.T()
	er := NewExchangeRestore("backupID")
	assert.Equal(t, er.Service, ServiceExchange)
	assert.Equal(t, er.BackupID, "backupID")
	assert.NotZero(t, er.Scopes())
}

func (suite *ExchangeSourceSuite) TestToExchangeRestore() {
	t := suite.T()
	eb := NewExchangeRestore("rpid")
	s := eb.Selector
	eb, err := s.ToExchangeRestore()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.Equal(t, eb.BackupID, "rpid")
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_Contacts() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		c1     = "c1"
		c2     = "c2"
	)

	sel.Exclude(sel.Contacts(user, folder, c1, c2))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], folder)
	assert.Equal(t, scope[ExchangeContact.String()], join(c1, c2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_Contacts() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		c1     = "c1"
		c2     = "c2"
	)

	sel.Include(sel.Contacts(user, folder, c1, c2))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], folder)
	assert.Equal(t, scope[ExchangeContact.String()], join(c1, c2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContact)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_ContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Exclude(sel.ContactFolders(user, f1, f2))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeContact.String()], None)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_ContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Include(sel.ContactFolders(user, f1, f2))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeContact.String()], All)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContactFolder)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_Events() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
	)

	sel.Exclude(sel.Events(user, e1, e2))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeEvent.String()], join(e1, e2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_Events() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
	)

	sel.Include(sel.Events(user, e1, e2))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeEvent.String()], join(e1, e2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeEvent)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_Mails() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		m1     = "m1"
		m2     = "m2"
	)

	sel.Exclude(sel.Mails(user, folder, m1, m2))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], folder)
	assert.Equal(t, scope[ExchangeMail.String()], join(m1, m2))
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_Mails() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = All
		m1     = "m1"
		m2     = "m2"
	)

	sel.Include(sel.Mails(user, folder, m1, m2))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], folder)
	assert.Equal(t, scope[ExchangeMail.String()], join(m1, m2))

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMail)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_MailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Exclude(sel.MailFolders(user, f1, f2))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeMail.String()], None)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_MailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Include(sel.MailFolders(user, f1, f2))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeMail.String()], All)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMailFolder)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_Users() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Exclude(sel.Users(u1, u2))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], join(u1, u2))
	assert.Equal(t, scope[ExchangeContact.String()], None)
	assert.Equal(t, scope[ExchangeContactFolder.String()], None)
	assert.Equal(t, scope[ExchangeEvent.String()], None)
	assert.Equal(t, scope[ExchangeMail.String()], None)
	assert.Equal(t, scope[ExchangeMailFolder.String()], None)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_Users() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Include(sel.Users(u1, u2))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], join(u1, u2))
	assert.Equal(t, scope[ExchangeContact.String()], All)
	assert.Equal(t, scope[ExchangeContactFolder.String()], All)
	assert.Equal(t, scope[ExchangeEvent.String()], All)
	assert.Equal(t, scope[ExchangeMail.String()], All)
	assert.Equal(t, scope[ExchangeMailFolder.String()], All)

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

func (suite *ExchangeSourceSuite) TestExchangeScope_IncludesPath() {
	const (
		usr  = "userID"
		fld  = "mailFolder"
		mail = "mailID"
	)
	var (
		path = []string{"tid", usr, "mail", fld, mail}
		es   = NewExchangeRestore("rpid")
	)

	table := []struct {
		name   string
		scope  exchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"all user's items", es.Users(All), assert.True},
		{"no user's items", es.Users(None), assert.False},
		{"matching user", es.Users(usr), assert.True},
		{"non-maching user", es.Users("smarf"), assert.False},
		{"one of multiple users", es.Users("smarf", usr), assert.True},
		{"all folders", es.MailFolders(All, All), assert.True},
		{"no folders", es.MailFolders(All, None), assert.False},
		{"matching folder", es.MailFolders(All, fld), assert.True},
		{"non-matching folder", es.MailFolders(All, "smarf"), assert.False},
		{"one of multiple folders", es.MailFolders(All, "smarf", fld), assert.True},
		{"all mail", es.Mails(All, All, All), assert.True},
		{"no mail", es.Mails(All, All, None), assert.False},
		{"matching mail", es.Mails(All, All, mail), assert.True},
		{"non-matching mail", es.Mails(All, All, "smarf"), assert.False},
		{"one of multiple mails", es.Mails(All, All, "smarf", mail), assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scope := extendExchangeScopeValues(All, test.scope)
			test.expect(t, scope.includesPath(ExchangeMail, path))
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScope_ExcludesPath() {
	const (
		usr  = "userID"
		fld  = "mailFolder"
		mail = "mailID"
	)
	var (
		path = []string{"tid", usr, "mail", fld, mail}
		es   = NewExchangeRestore("rpid")
	)

	table := []struct {
		name   string
		scope  exchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"all user's items", es.Users(All), assert.True},
		{"no user's items", es.Users(None), assert.False},
		{"matching user", es.Users(usr), assert.True},
		{"non-maching user", es.Users("smarf"), assert.False},
		{"one of multiple users", es.Users("smarf", usr), assert.True},
		{"all folders", es.MailFolders(None, All), assert.True},
		{"no folders", es.MailFolders(None, None), assert.False},
		{"matching folder", es.MailFolders(None, fld), assert.True},
		{"non-matching folder", es.MailFolders(None, "smarf"), assert.False},
		{"one of multiple folders", es.MailFolders(None, "smarf", fld), assert.True},
		{"all mail", es.Mails(None, None, All), assert.True},
		{"no mail", es.Mails(None, None, None), assert.False},
		{"matching mail", es.Mails(None, None, mail), assert.True},
		{"non-matching mail", es.Mails(None, None, "smarf"), assert.False},
		{"one of multiple mails", es.Mails(None, None, "smarf", mail), assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scope := extendExchangeScopeValues(None, test.scope)
			test.expect(t, scope.excludesPath(ExchangeMail, path))
		})
	}
}

func (suite *ExchangeSourceSuite) TestIdPath() {
	table := []struct {
		cat    exchangeCategory
		path   []string
		expect map[exchangeCategory]string
	}{
		{
			ExchangeContact,
			[]string{"tid", "uid", "contact", "cFld", "cid"},
			map[exchangeCategory]string{
				ExchangeUser:          "uid",
				ExchangeContactFolder: "cFld",
				ExchangeContact:       "cid",
			},
		},
		{
			ExchangeEvent,
			[]string{"tid", "uid", "event", "eid"},
			map[exchangeCategory]string{
				ExchangeUser:  "uid",
				ExchangeEvent: "eid",
			},
		},
		{
			ExchangeMail,
			[]string{"tid", "uid", "mail", "mFld", "mid"},
			map[exchangeCategory]string{
				ExchangeUser:       "uid",
				ExchangeMailFolder: "mFld",
				ExchangeMail:       "mid",
			},
		},
		{
			ExchangeCategoryUnknown,
			[]string{"tid", "uid", "contact", "cFld", "cid"},
			map[exchangeCategory]string{
				ExchangeUser: "uid",
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeRestore_FilterDetails() {
	makeDeets := func(refs ...string) *restorepoint.Details {
		deets := &restorepoint.Details{
			Entries: []restorepoint.DetailsEntry{},
		}
		for _, r := range refs {
			deets.Entries = append(deets.Entries, restorepoint.DetailsEntry{
				RepoRef: r,
			})
		}
		return deets
	}
	const (
		contact = "tid/uid/contact/cfld/cid"
		event   = "tid/uid/event/eid"
		mail    = "tid/uid/mail/mfld/mid"
	)
	table := []struct {
		name         string
		deets        *restorepoint.Details
		makeSelector func() *ExchangeRestore
		expect       []string
	}{
		{
			"no refs",
			makeDeets(),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				return er
			},
			[]string{},
		},
		{
			"contact only",
			makeDeets(contact),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				return er
			},
			[]string{contact},
		},
		{
			"event only",
			makeDeets(event),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				return er
			},
			[]string{event},
		},
		{
			"mail only",
			makeDeets(mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				return er
			},
			[]string{mail},
		},
		{
			"all",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				return er
			},
			[]string{contact, event, mail},
		},
		{
			"only match contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Contacts("uid", "cfld", "cid"))
				return er
			},
			[]string{contact},
		},
		{
			"only match event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Events("uid", "eid"))
				return er
			},
			[]string{event},
		},
		{
			"only match mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Mails("uid", "mfld", "mid"))
				return er
			},
			[]string{mail},
		},
		{
			"exclude contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				er.Exclude(er.Contacts("uid", "cfld", "cid"))
				return er
			},
			[]string{event, mail},
		},
		{
			"exclude event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				er.Exclude(er.Events("uid", "eid"))
				return er
			},
			[]string{contact, mail},
		},
		{
			"exclude mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore("rpid")
				er.Include(er.Users(All))
				er.Exclude(er.Mails("uid", "mfld", "mid"))
				return er
			},
			[]string{contact, event},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := test.makeSelector()
			results := sel.FilterDetails(test.deets)
			assert.Equal(t, test.expect, results)
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScopesByCategory() {
	var (
		es       = NewExchangeRestore("rpid")
		users    = es.Users(All)
		contacts = es.ContactFolders(All, All)
		events   = es.Events(All, All)
		mail     = es.MailFolders(All, All)
	)
	type expect struct {
		contact int
		event   int
		mail    int
	}
	type input []map[string]string
	table := []struct {
		name   string
		scopes input
		expect expect
	}{
		{"users: one of each", input{users}, expect{1, 1, 1}},
		{"contacts only", input{contacts}, expect{1, 0, 0}},
		{"events only", input{events}, expect{0, 1, 0}},
		{"mail only", input{mail}, expect{0, 0, 1}},
		{"all", input{users, contacts, events, mail}, expect{2, 2, 2}},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := exchangeScopesByCategory(test.scopes)
			assert.Equal(t, test.expect.contact, len(result[ExchangeContact.String()]))
			assert.Equal(t, test.expect.event, len(result[ExchangeEvent.String()]))
			assert.Equal(t, test.expect.mail, len(result[ExchangeMail.String()]))
		})
	}
}

func (suite *ExchangeSourceSuite) TestMatchExchangeEntry() {
	const (
		mail = "mailID"
		cat  = ExchangeMail
	)
	include := func(s map[string]string) exchangeScope {
		return extendExchangeScopeValues(All, exchangeScope(s))
	}
	exclude := func(s map[string]string) exchangeScope {
		return extendExchangeScopeValues(None, exchangeScope(s))
	}
	var (
		es          = NewExchangeRestore("rpid")
		inAll       = include(es.Users(All))
		inNone      = include(es.Users(None))
		inMail      = include(es.Mails(All, All, mail))
		inOtherMail = include(es.Mails(All, All, "smarf"))
		exAll       = exclude(es.Users(All))
		exNone      = exclude(es.Users(None))
		exMail      = exclude(es.Mails(None, None, mail))
		exOtherMail = exclude(es.Mails(None, None, "smarf"))
		path        = []string{"tid", "user", "mail", "folder", mail}
	)

	table := []struct {
		name     string
		includes []exchangeScope
		excludes []exchangeScope
		expect   assert.BoolAssertionFunc
	}{
		{"empty", []exchangeScope{}, []exchangeScope{}, assert.False},
		{"in all", []exchangeScope{inAll}, []exchangeScope{}, assert.True},
		{"in None", []exchangeScope{inNone}, []exchangeScope{}, assert.False},
		{"in Mail", []exchangeScope{inMail}, []exchangeScope{}, assert.True},
		{"in Other", []exchangeScope{inOtherMail}, []exchangeScope{}, assert.False},
		{"ex all", []exchangeScope{inAll}, []exchangeScope{exAll}, assert.False},
		{"ex None", []exchangeScope{inAll}, []exchangeScope{exNone}, assert.True},
		{"in Mail", []exchangeScope{inAll}, []exchangeScope{exMail}, assert.False},
		{"in Other", []exchangeScope{inAll}, []exchangeScope{exOtherMail}, assert.True},
		{"in and ex mail", []exchangeScope{inMail}, []exchangeScope{exMail}, assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, matchExchangeEntry(cat, path, test.includes, test.excludes))
		})
	}
}
