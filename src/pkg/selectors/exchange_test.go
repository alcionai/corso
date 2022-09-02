package selectors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/filters"
)

type ExchangeSelectorSuite struct {
	suite.Suite
}

func TestExchangeSelectorSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSelectorSuite))
}

func (suite *ExchangeSelectorSuite) TestNewExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup()
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSelectorSuite) TestToExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup()
	s := eb.Selector
	eb, err := s.ToExchangeBackup()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSelectorSuite) TestNewExchangeRestore() {
	t := suite.T()
	er := NewExchangeRestore()
	assert.Equal(t, er.Service, ServiceExchange)
	assert.NotZero(t, er.Scopes())
}

func (suite *ExchangeSelectorSuite) TestToExchangeRestore() {
	t := suite.T()
	eb := NewExchangeRestore()
	s := eb.Selector
	eb, err := s.ToExchangeRestore()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Contacts() {
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
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeContactFolder: folder,
			ExchangeContact:       join(c1, c2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Contacts() {
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
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeContactFolder: folder,
			ExchangeContact:       join(c1, c2),
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContact)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_ContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Exclude(sel.ContactFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeContactFolder: join(f1, f2),
			ExchangeContact:       AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_ContactFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Include(sel.ContactFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeContactFolder: join(f1, f2),
			ExchangeContact:       AnyTgt,
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContactFolder)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Events() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
		c1   = "c1"
	)

	sel.Exclude(sel.Events([]string{user}, []string{c1}, []string{e1, e2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeEventCalendar: c1,
			ExchangeEvent:         join(e1, e2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_EventCalendars() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		c1   = "c1"
		c2   = "c2"
	)

	sel.Exclude(sel.EventCalendars([]string{user}, []string{c1, c2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeEventCalendar: join(c1, c2),
			ExchangeEvent:         AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Events() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
		c1   = "c1"
	)

	sel.Include(sel.Events([]string{user}, []string{c1}, []string{e1, e2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeEventCalendar: c1,
			ExchangeEvent:         join(e1, e2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_EventCalendars() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		c1   = "c1"
		c2   = "c2"
	)

	sel.Include(sel.EventCalendars([]string{user}, []string{c1, c2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:          user,
			ExchangeEventCalendar: join(c1, c2),
			ExchangeEvent:         AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Mails() {
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
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:       user,
			ExchangeMailFolder: folder,
			ExchangeMail:       join(m1, m2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Mails() {
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
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:       user,
			ExchangeMailFolder: folder,
			ExchangeMail:       join(m1, m2),
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMail)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_MailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Exclude(sel.MailFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:       user,
			ExchangeMailFolder: join(f1, f2),
			ExchangeMail:       AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_MailFolders() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel.Include(sel.MailFolders([]string{user}, []string{f1, f2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeUser:       user,
			ExchangeMailFolder: join(f1, f2),
			ExchangeMail:       AnyTgt,
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMailFolder)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Users() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Exclude(sel.Users([]string{u1, u2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			ExchangeScope(sc),
			map[categorizer]string{ExchangeUser: join(u1, u2)},
		)

		if sc[scopeKeyCategory].Compare(ExchangeContactFolder.String()) {
			scopeMustHave(
				t,
				ExchangeScope(sc),
				map[categorizer]string{
					ExchangeContact:       AnyTgt,
					ExchangeContactFolder: AnyTgt,
				},
			)
		}

		if sc[scopeKeyCategory].Compare(ExchangeEvent.String()) {
			scopeMustHave(
				t,
				ExchangeScope(sc),
				map[categorizer]string{
					ExchangeEvent: AnyTgt,
				},
			)
		}

		if sc[scopeKeyCategory].Compare(ExchangeMailFolder.String()) {
			scopeMustHave(
				t,
				ExchangeScope(sc),
				map[categorizer]string{
					ExchangeMail:       AnyTgt,
					ExchangeMailFolder: AnyTgt,
				},
			)
		}
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Users() {
	t := suite.T()
	sel := NewExchangeBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Include(sel.Users([]string{u1, u2}))
	scopes := sel.Includes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			ExchangeScope(sc),
			map[categorizer]string{ExchangeUser: join(u1, u2)},
		)

		if sc[scopeKeyCategory].Compare(ExchangeContactFolder.String()) {
			scopeMustHave(
				t,
				ExchangeScope(sc),
				map[categorizer]string{
					ExchangeContact:       AnyTgt,
					ExchangeContactFolder: AnyTgt,
				},
			)
		}

		if sc[scopeKeyCategory].Compare(ExchangeEvent.String()) {
			scopeMustHave(
				t,
				ExchangeScope(sc),
				map[categorizer]string{
					ExchangeEvent: AnyTgt,
				},
			)
		}

		if sc[scopeKeyCategory].Compare(ExchangeMailFolder.String()) {
			scopeMustHave(
				t,
				ExchangeScope(sc),
				map[categorizer]string{
					ExchangeMail:       AnyTgt,
					ExchangeMailFolder: AnyTgt,
				},
			)
		}
	}
}

func (suite *ExchangeSelectorSuite) TestNewExchangeDestination() {
	t := suite.T()
	dest := NewExchangeDestination()
	assert.Len(t, dest, 0)
}

func (suite *ExchangeSelectorSuite) TestExchangeDestination_Set() {
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

func (suite *ExchangeSelectorSuite) TestExchangeDestination_GetOrDefault() {
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

func (suite *ExchangeSelectorSuite) TestExchangeBackup_Scopes() {
	eb := NewExchangeBackup()
	eb.Include(eb.Users(Any()))

	scopes := eb.Scopes()
	assert.Len(suite.T(), scopes, 3)

	for _, sc := range scopes {
		cat := sc.Category()
		suite.T().Run(cat.String(), func(t *testing.T) {
			assert.True(t, sc.IsAny(ExchangeUser))
			switch sc.Category() {
			case ExchangeContactFolder:
				assert.True(t, sc.IsAny(ExchangeContact))
				assert.True(t, sc.IsAny(ExchangeContactFolder))
			case ExchangeEvent:
				assert.True(t, sc.IsAny(ExchangeEvent))
			case ExchangeMailFolder:
				assert.True(t, sc.IsAny(ExchangeMail))
				assert.True(t, sc.IsAny(ExchangeMailFolder))
			}
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeBackup_DiscreteScopes() {
	usrs := []string{"u1", "u2"}
	table := []struct {
		name     string
		include  []string
		discrete []string
		expect   []string
	}{
		{
			name:     "any user",
			include:  Any(),
			discrete: usrs,
			expect:   usrs,
		},
		{
			name:     "discrete user",
			include:  []string{"u3"},
			discrete: usrs,
			expect:   []string{"u3"},
		},
		{
			name:     "nil discrete slice",
			include:  Any(),
			discrete: nil,
			expect:   Any(),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			eb := NewExchangeBackup()
			eb.Include(eb.Users(test.include))

			scopes := eb.DiscreteScopes(test.discrete)
			for _, sc := range scopes {
				users := sc.Get(ExchangeUser)
				assert.Equal(t, test.expect, users)
			}
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_Category() {
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
			eb.Includes = []scope{
				{scopeKeyCategory: filters.Identity(test.is.String())},
			}
			scope := eb.Scopes()[0]
			test.check(t, test.expect, scope.Category())
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_IncludesCategory() {
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
			eb.Includes = []scope{
				{scopeKeyCategory: filters.Identity(test.is.String())},
			}
			scope := eb.Scopes()[0]
			test.check(t, scope.IncludesCategory(test.expect))
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_Get() {
	eb := NewExchangeBackup()
	eb.Include(eb.Users(Any()))

	scopes := eb.Scopes()

	table := []exchangeCategory{
		ExchangeContact,
		ExchangeContactFolder,
		ExchangeEvent,
		ExchangeMail,
		ExchangeMailFolder,
		ExchangeUser,
	}
	for _, test := range table {
		suite.T().Run(test.String(), func(t *testing.T) {
			for _, sc := range scopes {
				assert.Equal(t, Any(), sc.Get(ExchangeUser))
				switch sc.Category() {
				case ExchangeContactFolder:
					assert.Equal(t, Any(), sc.Get(ExchangeContact))
					assert.Equal(t, Any(), sc.Get(ExchangeContactFolder))
				case ExchangeEvent:
					assert.Equal(t, Any(), sc.Get(ExchangeEvent))
				case ExchangeMailFolder:
					assert.Equal(t, Any(), sc.Get(ExchangeMail))
					assert.Equal(t, Any(), sc.Get(ExchangeMailFolder))
				}
				assert.Equal(t, None(), sc.Get(ExchangeCategoryUnknown))
			}
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_MatchesInfo() {
	es := NewExchangeRestore()

	const (
		name      = "smarf mcfnords"
		organizer = "cooks@2many.smarf"
		sender    = "smarf@2many.cooks"
		subject   = "I have seen the fnords!"
	)

	var (
		epoch  = time.Time{}
		now    = time.Now()
		future = now.Add(1 * time.Minute)
		info   = &details.ExchangeInfo{
			ContactName: name,
			EventRecurs: true,
			EventStart:  now,
			Organizer:   organizer,
			Sender:      sender,
			Subject:     subject,
			Received:    now,
		}
	)

	table := []struct {
		name   string
		scope  []ExchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"any mail with a sender", es.MailSender(AnyTgt), assert.True},
		{"no mail, regardless of sender", es.MailSender(NoneTgt), assert.False},
		{"mail from a different sender", es.MailSender("magoo@ma.goo"), assert.False},
		{"mail from the matching sender", es.MailSender(sender), assert.True},
		{"mail with any subject", es.MailSubject(AnyTgt), assert.True},
		{"mail with none subject", es.MailSubject(NoneTgt), assert.False},
		{"mail with a different subject", es.MailSubject("fancy"), assert.False},
		{"mail with the matching subject", es.MailSubject(subject), assert.True},
		{"mail with a substring subject match", es.MailSubject(subject[5:9]), assert.True},
		{"mail received after the epoch", es.MailReceivedAfter(common.FormatTime(epoch)), assert.True},
		{"mail received after now", es.MailReceivedAfter(common.FormatTime(now)), assert.False},
		{"mail received after sometime later", es.MailReceivedAfter(common.FormatTime(future)), assert.False},
		{"mail received before the epoch", es.MailReceivedBefore(common.FormatTime(epoch)), assert.False},
		{"mail received before now", es.MailReceivedBefore(common.FormatTime(now)), assert.False},
		{"mail received before sometime later", es.MailReceivedBefore(common.FormatTime(future)), assert.True},
		{"event with any organizer", es.EventOrganizer(AnyTgt), assert.True},
		{"event with none organizer", es.EventOrganizer(NoneTgt), assert.False},
		{"event with a different organizer", es.EventOrganizer("fancy"), assert.False},
		{"event with the matching organizer", es.EventOrganizer(organizer), assert.True},
		{"event that recurs", es.EventRecurs("true"), assert.True},
		{"event that does not recur", es.EventRecurs("false"), assert.False},
		{"event starting after the epoch", es.EventStartsAfter(common.FormatTime(epoch)), assert.True},
		{"event starting after now", es.EventStartsAfter(common.FormatTime(now)), assert.False},
		{"event starting after sometime later", es.EventStartsAfter(common.FormatTime(future)), assert.False},
		{"event starting before the epoch", es.EventStartsBefore(common.FormatTime(epoch)), assert.False},
		{"event starting before now", es.EventStartsBefore(common.FormatTime(now)), assert.False},
		{"event starting before sometime later", es.EventStartsBefore(common.FormatTime(future)), assert.True},
		{"event with any subject", es.EventSubject(AnyTgt), assert.True},
		{"event with none subject", es.EventSubject(NoneTgt), assert.False},
		{"event with a different subject", es.EventSubject("fancy"), assert.False},
		{"event with the matching subject", es.EventSubject(subject), assert.True},
		{"contact with a different name", es.ContactName("blarps"), assert.False},
		{"contact with the same name", es.ContactName(name), assert.True},
		{"contact with a subname search", es.ContactName(name[2:5]), assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scopes := setScopesToDefault(test.scope)
			for _, scope := range scopes {
				test.expect(t, scope.matchesInfo(info))
			}
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_MatchesPath() {
	const (
		usr  = "userID"
		fld  = "mailFolder"
		mail = "mailID"
	)

	var (
		path = []string{"tid", usr, "email", fld, mail}
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
			scopes := setScopesToDefault(test.scope)
			var aMatch bool
			for _, scope := range scopes {
				pv := ExchangeMail.pathValues(path)
				if matchesPathValues(scope, ExchangeMail, pv) {
					aMatch = true
					break
				}
			}
			test.expect(t, aMatch)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestIdPath() {
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
			[]string{"tid", "uid", "email", "mFld", "mid"},
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

func (suite *ExchangeSelectorSuite) TestExchangeRestore_Reduce() {
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
		contact = "tid/uid/contacts/cfld/cid"
		event   = "tid/uid/events/ecld/eid"
		mail    = "tid/uid/email/mfld/mid"
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
				er.Include(er.Events([]string{"uid"}, []string{"ecld"}, []string{"eid"}))
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
				er.Exclude(er.Events([]string{"uid"}, []string{"ecld"}, []string{"eid"}))
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

func (suite *ExchangeSelectorSuite) TestScopesByCategory() {
	var (
		es       = NewExchangeRestore()
		users    = es.Users(Any())
		contacts = es.ContactFolders(Any(), Any())
		events   = es.Events(Any(), Any(), Any())
		mail     = es.MailFolders(Any(), Any())
	)

	type expect struct {
		contact int
		event   int
		mail    int
	}

	type input []scope

	makeInput := func(es ...[]ExchangeScope) []scope {
		mss := []scope{}

		for _, sl := range es {
			for _, s := range sl {
				mss = append(mss, scope(s))
			}
		}

		return mss
	}
	cats := map[pathType]exchangeCategory{
		exchangeContactPath: ExchangeContact,
		exchangeEventPath:   ExchangeEvent,
		exchangeMailPath:    ExchangeMail,
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
			result := scopesByCategory[ExchangeScope](test.scopes, cats)
			assert.Len(t, result[ExchangeContact], test.expect.contact)
			assert.Len(t, result[ExchangeEvent], test.expect.event)
			assert.Len(t, result[ExchangeMail], test.expect.mail)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestPasses() {
	deets := details.DetailsEntry{}

	const (
		mid = "mailID"
		cat = ExchangeMail
	)

	var (
		es        = NewExchangeRestore()
		anyUser   = setScopesToDefault(es.Users(Any()))
		noUser    = setScopesToDefault(es.Users(None()))
		mail      = setScopesToDefault(es.Mails(Any(), Any(), []string{mid}))
		otherMail = setScopesToDefault(es.Mails(Any(), Any(), []string{"smarf"}))
		noMail    = setScopesToDefault(es.Mails(Any(), Any(), None()))
		path      = []string{"tid", "user", "email", "folder", mid}
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
			result := passes(
				cat,
				cat.pathValues(path),
				deets,
				test.excludes,
				test.filters,
				test.includes)
			test.expect(t, result)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestContains() {
	target := "fnords"

	var (
		es                  = NewExchangeRestore()
		anyUser             = setScopesToDefault(es.Users(Any()))
		noMail              = setScopesToDefault(es.Mails(None(), None(), None()))
		does                = setScopesToDefault(es.Mails(Any(), Any(), []string{target}))
		doesNot             = setScopesToDefault(es.Mails(Any(), Any(), []string{"smarf"}))
		wrongType           = setScopesToDefault(es.Contacts(Any(), Any(), Any()))
		wrongTypeGoodTarget = setScopesToDefault(es.Contacts(Any(), Any(), Any()))
	)

	table := []struct {
		name   string
		scopes []ExchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"any user", anyUser, assert.True},
		{"no mail", noMail, assert.False},
		{"does contain", does, assert.True},
		{"does not contain", doesNot, assert.False},
		{"wrong type", wrongType, assert.False},
		{"wrong type but right target", wrongTypeGoodTarget, assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			var result bool
			for _, scope := range test.scopes {
				if scope.Matches(ExchangeMail, target) {
					result = true
					break
				}
			}
			test.expect(t, result)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestIsAny() {
	var (
		es           = NewExchangeRestore()
		anyUser      = setScopesToDefault(es.Users(Any()))
		noUser       = setScopesToDefault(es.Users(None()))
		specificMail = setScopesToDefault(es.Mails(Any(), Any(), []string{"email"}))
		anyMail      = setScopesToDefault(es.Mails(Any(), Any(), Any()))
	)

	table := []struct {
		name   string
		scopes []ExchangeScope
		cat    exchangeCategory
		expect assert.BoolAssertionFunc
	}{
		{"any user", anyUser, ExchangeUser, assert.True},
		{"no user", noUser, ExchangeUser, assert.False},
		{"specific mail", specificMail, ExchangeMail, assert.False},
		{"any mail", anyMail, ExchangeMail, assert.True},
		{"wrong category", anyMail, ExchangeEvent, assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			var result bool
			for _, scope := range test.scopes {
				if scope.IsAny(test.cat) {
					result = true
					break
				}
			}
			test.expect(t, result)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeCategory_leafCat() {
	table := []struct {
		cat    exchangeCategory
		expect exchangeCategory
	}{
		{exchangeCategory("foo"), exchangeCategory("foo")},
		{ExchangeCategoryUnknown, ExchangeCategoryUnknown},
		{ExchangeUser, ExchangeUser},
		{ExchangeMailFolder, ExchangeMail},
		{ExchangeMail, ExchangeMail},
		{ExchangeContactFolder, ExchangeContact},
		{ExchangeEvent, ExchangeEvent},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.expect, test.cat.leafCat(), test.cat.String())
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeCategory_PathValues() {
	contactPath := []string{"ten", "user", "contact", "cfolder", "contactitem"}
	contactMap := map[categorizer]string{
		ExchangeUser:          contactPath[1],
		ExchangeContactFolder: contactPath[3],
		ExchangeContact:       contactPath[4],
	}
	eventPath := []string{"ten", "user", "event", "ecalendar", "eventitem"}
	eventMap := map[categorizer]string{
		ExchangeUser:          eventPath[1],
		ExchangeEventCalendar: eventPath[3],
		ExchangeEvent:         eventPath[4],
	}
	mailPath := []string{"ten", "user", "mail", "mfolder", "mailitem"}
	mailMap := map[categorizer]string{
		ExchangeUser:       contactPath[1],
		ExchangeMailFolder: mailPath[3],
		ExchangeMail:       mailPath[4],
	}

	table := []struct {
		cat    exchangeCategory
		path   []string
		expect map[categorizer]string
	}{
		{ExchangeCategoryUnknown, nil, map[categorizer]string{}},
		{ExchangeCategoryUnknown, []string{"a"}, map[categorizer]string{}},
		{ExchangeContact, contactPath, contactMap},
		{ExchangeEvent, eventPath, eventMap},
		{ExchangeMail, mailPath, mailMap},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.cat.pathValues(test.path), test.expect)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeCategory_PathKeys() {
	contact := []categorizer{ExchangeUser, ExchangeContactFolder, ExchangeContact}
	event := []categorizer{ExchangeUser, ExchangeEventCalendar, ExchangeEvent}
	mail := []categorizer{ExchangeUser, ExchangeMailFolder, ExchangeMail}
	user := []categorizer{ExchangeUser}

	var empty []categorizer

	table := []struct {
		cat    exchangeCategory
		expect []categorizer
	}{
		{ExchangeCategoryUnknown, empty},
		{ExchangeContact, contact},
		{ExchangeEvent, event},
		{ExchangeMail, mail},
		{ExchangeUser, user},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.cat.pathKeys(), test.expect)
		})
	}
}
