package selectors

import (
	"testing"
	"time"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/backup/details"
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
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestToExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup()
	s := eb.Selector
	eb, err := s.ToExchangeBackup()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestNewExchangeRestore() {
	t := suite.T()
	er := NewExchangeRestore()
	assert.Equal(t, er.Service, ServiceExchange)
	assert.NotZero(t, er.Scopes())
}

func (suite *ExchangeSourceSuite) TestToExchangeRestore() {
	t := suite.T()
	eb := NewExchangeRestore()
	s := eb.Selector
	eb, err := s.ToExchangeRestore()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_Contacts() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user   = "user"
		folder = AnyTgt
		c1     = "c1"
		c2     = "c2"
	)

	sel.Exclude(sel.Contacts([]string{user}, []string{folder}, []string{c1, c2}))
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
		folder = AnyTgt
		c1     = "c1"
		c2     = "c2"
	)

	sel.Include(sel.Contacts([]string{user}, []string{folder}, []string{c1, c2}))
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

	sel.Exclude(sel.ContactFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeContact.String()], AnyTgt)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_ContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Include(sel.ContactFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeContactFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeContact.String()], AnyTgt)

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

	sel.Exclude(sel.Events([]string{user}, []string{e1, e2}))
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

	sel.Include(sel.Events([]string{user}, []string{e1, e2}))
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
		folder = AnyTgt
		m1     = "m1"
		m2     = "m2"
	)

	sel.Exclude(sel.Mails([]string{user}, []string{folder}, []string{m1, m2}))
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
		folder = AnyTgt
		m1     = "m1"
		m2     = "m2"
	)

	sel.Include(sel.Mails([]string{user}, []string{folder}, []string{m1, m2}))
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

	sel.Exclude(sel.MailFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Excludes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeMail.String()], AnyTgt)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_MailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Include(sel.MailFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Includes
	require.Equal(t, 1, len(scopes))

	scope := scopes[0]
	assert.Equal(t, scope[ExchangeUser.String()], user)
	assert.Equal(t, scope[ExchangeMailFolder.String()], join(f1, f2))
	assert.Equal(t, scope[ExchangeMail.String()], AnyTgt)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMailFolder)
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Exclude_Users() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Exclude(sel.Users([]string{u1, u2}))
	scopes := sel.Excludes
	require.Equal(t, 6, len(scopes))

	for _, scope := range scopes {
		assert.Contains(t, join(u1, u2), scope[ExchangeUser.String()])
		if scope[scopeKeyCategory] == ExchangeContactFolder.String() {
			assert.Equal(t, scope[ExchangeContact.String()], AnyTgt)
			assert.Equal(t, scope[ExchangeContactFolder.String()], AnyTgt)
		}
		if scope[scopeKeyCategory] == ExchangeEvent.String() {
			assert.Equal(t, scope[ExchangeEvent.String()], AnyTgt)
		}
		if scope[scopeKeyCategory] == ExchangeMailFolder.String() {
			assert.Equal(t, scope[ExchangeMail.String()], AnyTgt)
			assert.Equal(t, scope[ExchangeMailFolder.String()], AnyTgt)
		}
	}
}

func (suite *ExchangeSourceSuite) TestExchangeSelector_Include_Users() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Include(sel.Users([]string{u1, u2}))
	scopes := sel.Includes
	require.Equal(t, 6, len(scopes))

	for _, scope := range scopes {
		assert.Contains(t, join(u1, u2), scope[ExchangeUser.String()])
		if scope[scopeKeyCategory] == ExchangeContactFolder.String() {
			assert.Equal(t, scope[ExchangeContact.String()], AnyTgt)
			assert.Equal(t, scope[ExchangeContactFolder.String()], AnyTgt)
		}
		if scope[scopeKeyCategory] == ExchangeEvent.String() {
			assert.Equal(t, scope[ExchangeEvent.String()], AnyTgt)
		}
		if scope[scopeKeyCategory] == ExchangeMailFolder.String() {
			assert.Equal(t, scope[ExchangeMail.String()], AnyTgt)
			assert.Equal(t, scope[ExchangeMailFolder.String()], AnyTgt)
		}
	}
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
	ExchangeContact.String():       AnyTgt,
	ExchangeContactFolder.String(): AnyTgt,
	ExchangeEvent.String():         AnyTgt,
	ExchangeMail.String():          AnyTgt,
	ExchangeMailFolder.String():    AnyTgt,
	ExchangeUser.String():          AnyTgt,
}

func (suite *ExchangeSourceSuite) TestExchangeBackup_Scopes() {
	eb := NewExchangeBackup()
	eb.Includes = []map[string]string{allScopesExceptUnknown}
	// todo: swap the above for this
	// eb := NewExchangeBackup().IncludeUsers(AnyTgt)

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
	// eb := NewExchangeBackup().IncludeUsers(AnyTgt)

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
		None(),
		scope.Get(ExchangeCategoryUnknown))

	expect := Any()
	for _, test := range table {
		suite.T().Run(test.String(), func(t *testing.T) {
			assert.Equal(t, expect, scope.Get(test))
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScope_MatchesInfo() {
	es := NewExchangeRestore()
	const (
		sender  = "smarf@2many.cooks"
		subject = "I have seen the fnords!"
	)
	var (
		epoch = time.Time{}
		now   = time.Now()
		then  = now.Add(1 * time.Minute)
		info  = &details.ExchangeInfo{
			Sender:   sender,
			Subject:  subject,
			Received: now,
		}
	)

	table := []struct {
		name   string
		scope  []ExchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"any mail with a sender", es.MailSender(Any()), assert.True},
		{"no mail, regardless of sender", es.MailSender(None()), assert.False},
		{"mail from a different sender", es.MailSender([]string{"magoo@ma.goo"}), assert.False},
		{"mail from the matching sender", es.MailSender([]string{sender}), assert.True},
		{"mail with any subject", es.MailSubject(Any()), assert.True},
		{"no mail, regardless of subject", es.MailSubject(None()), assert.False},
		{"mail with a different subject", es.MailSubject([]string{"fancy"}), assert.False},
		{"mail with the matching subject", es.MailSubject([]string{subject}), assert.True},
		{"mail with a substring subject match", es.MailSubject([]string{subject[5:9]}), assert.True},
		{"mail received after the epoch", es.MailReceivedAfter(common.FormatTime(epoch)), assert.True},
		{"mail received after now", es.MailReceivedAfter(common.FormatTime(now)), assert.False},
		{"mail received after sometime later", es.MailReceivedAfter(common.FormatTime(then)), assert.False},
		{"mail received before the epoch", es.MailReceivedBefore(common.FormatTime(epoch)), assert.False},
		{"mail received before now", es.MailReceivedBefore(common.FormatTime(now)), assert.False},
		{"mail received before sometime later", es.MailReceivedBefore(common.FormatTime(then)), assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scopes := extendExchangeScopeValues(test.scope)
			for _, scope := range scopes {
				test.expect(t, scope.matchesInfo(scope.Category(), info))
			}
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScope_MatchesPath() {
	const (
		usr  = "userID"
		fld  = "mailFolder"
		mail = "mailID"
	)
	var (
		path = []string{"tid", usr, "mail", fld, mail}
		es   = NewExchangeRestore()
	)

	table := []struct {
		name   string
		scope  []ExchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"all user's items", es.Users(Any()), assert.True},
		{"no user's items", es.Users(None()), assert.False},
		{"matching user", es.Users([]string{usr}), assert.True},
		{"non-matching user", es.Users([]string{"smarf"}), assert.False},
		{"one of multiple users", es.Users([]string{"smarf", usr}), assert.True},
		{"all folders", es.MailFolders(Any(), Any()), assert.True},
		{"no folders", es.MailFolders(Any(), None()), assert.False},
		{"matching folder", es.MailFolders(Any(), []string{fld}), assert.True},
		{"non-matching folder", es.MailFolders(Any(), []string{"smarf"}), assert.False},
		{"one of multiple folders", es.MailFolders(Any(), []string{"smarf", fld}), assert.True},
		{"all mail", es.Mails(Any(), Any(), Any()), assert.True},
		{"no mail", es.Mails(Any(), Any(), None()), assert.False},
		{"matching mail", es.Mails(Any(), Any(), []string{mail}), assert.True},
		{"non-matching mail", es.Mails(Any(), Any(), []string{"smarf"}), assert.False},
		{"one of multiple mails", es.Mails(Any(), Any(), []string{"smarf", mail}), assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scopes := extendExchangeScopeValues(test.scope)
			var aMatch bool
			for _, scope := range scopes {
				if scope.matchesPath(ExchangeMail, path) {
					aMatch = true
					break
				}
			}
			test.expect(t, aMatch)
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

func (suite *ExchangeSourceSuite) TestExchangeRestore_Reduce() {
	makeDeets := func(refs ...string) *details.Details {
		deets := &details.Details{
			DetailsModel: details.DetailsModel{
				Entries: []details.DetailsEntry{},
			},
		}
		for _, r := range refs {
			deets.Entries = append(deets.Entries, details.DetailsEntry{
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
	arr := func(s ...string) []string {
		return s
	}
	table := []struct {
		name         string
		deets        *details.Details
		makeSelector func() *ExchangeRestore
		expect       []string
	}{
		{
			"no refs",
			makeDeets(),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				return er
			},
			[]string{},
		},
		{
			"contact only",
			makeDeets(contact),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				return er
			},
			arr(contact),
		},
		{
			"event only",
			makeDeets(event),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				return er
			},
			arr(event),
		},
		{
			"mail only",
			makeDeets(mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				return er
			},
			arr(mail),
		},
		{
			"all",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				return er
			},
			arr(contact, event, mail),
		},
		{
			"only match contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Contacts([]string{"uid"}, []string{"cfld"}, []string{"cid"}))
				return er
			},
			arr(contact),
		},
		{
			"only match event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Events([]string{"uid"}, []string{"eid"}))
				return er
			},
			arr(event),
		},
		{
			"only match mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Mails([]string{"uid"}, []string{"mfld"}, []string{"mid"}))
				return er
			},
			arr(mail),
		},
		{
			"exclude contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				er.Exclude(er.Contacts([]string{"uid"}, []string{"cfld"}, []string{"cid"}))
				return er
			},
			arr(event, mail),
		},
		{
			"exclude event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				er.Exclude(er.Events([]string{"uid"}, []string{"eid"}))
				return er
			},
			arr(contact, mail),
		},
		{
			"exclude mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore()
				er.Include(er.Users(Any()))
				er.Exclude(er.Mails([]string{"uid"}, []string{"mfld"}, []string{"mid"}))
				return er
			},
			arr(contact, event),
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := test.makeSelector()
			results := sel.Reduce(test.deets)
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *ExchangeSourceSuite) TestExchangeScopesByCategory() {
	var (
		es       = NewExchangeRestore()
		users    = es.Users(Any())
		contacts = es.ContactFolders(Any(), Any())
		events   = es.Events(Any(), Any())
		mail     = es.MailFolders(Any(), Any())
	)
	type expect struct {
		contact int
		event   int
		mail    int
	}
	type input []map[string]string
	makeInput := func(es ...[]ExchangeScope) []map[string]string {
		mss := []map[string]string{}
		for _, sl := range es {
			for _, s := range sl {
				mss = append(mss, map[string]string(s))
			}
		}
		return mss
	}
	table := []struct {
		name   string
		scopes input
		expect expect
	}{
		{"users: one of each", makeInput(users), expect{1, 1, 1}},
		{"contacts only", makeInput(contacts), expect{1, 0, 0}},
		{"events only", makeInput(events), expect{0, 1, 0}},
		{"mail only", makeInput(mail), expect{0, 0, 1}},
		{"all", makeInput(users, contacts, events, mail), expect{2, 2, 2}},
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
	var exchangeInfo *details.ExchangeInfo
	const (
		mid = "mailID"
		cat = ExchangeMail
	)
	var (
		es        = NewExchangeRestore()
		anyUser   = extendExchangeScopeValues(es.Users(Any()))
		noUser    = extendExchangeScopeValues(es.Users(None()))
		mail      = extendExchangeScopeValues(es.Mails(Any(), Any(), []string{mid}))
		otherMail = extendExchangeScopeValues(es.Mails(Any(), Any(), []string{"smarf"}))
		noMail    = extendExchangeScopeValues(es.Mails(Any(), Any(), None()))
		path      = []string{"tid", "user", "mail", "folder", mid}
	)

	table := []struct {
		name                        string
		excludes, filters, includes []ExchangeScope
		expect                      assert.BoolAssertionFunc
	}{
		{"empty", nil, nil, nil, assert.False},
		{"in Any", nil, nil, anyUser, assert.True},
		{"in None", nil, nil, noUser, assert.False},
		{"in Mail", nil, nil, mail, assert.True},
		{"in Other", nil, nil, otherMail, assert.False},
		{"in no Mail", nil, nil, noMail, assert.False},
		{"ex Any", anyUser, nil, anyUser, assert.False},
		{"ex Any filter", anyUser, anyUser, nil, assert.False},
		{"ex None", noUser, nil, anyUser, assert.True},
		{"ex None filter mail", noUser, mail, nil, assert.True},
		{"ex None filter any user", noUser, anyUser, nil, assert.False},
		{"ex Mail", mail, nil, anyUser, assert.False},
		{"ex Other", otherMail, nil, anyUser, assert.True},
		{"in and ex Mail", mail, nil, mail, assert.False},
		{"filter Any", nil, anyUser, nil, assert.False},
		{"filter None", nil, noUser, anyUser, assert.False},
		{"filter Mail", nil, mail, anyUser, assert.True},
		{"filter Other", nil, otherMail, anyUser, assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, matchExchangeEntry(cat, path, exchangeInfo, test.excludes, test.filters, test.includes))
		})
	}
}
