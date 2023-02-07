package selectors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault/mock"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

type ExchangeSelectorSuite struct {
	suite.Suite
}

func TestExchangeSelectorSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSelectorSuite))
}

func (suite *ExchangeSelectorSuite) TestNewExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup(nil)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSelectorSuite) TestToExchangeBackup() {
	t := suite.T()
	eb := NewExchangeBackup(nil)
	s := eb.Selector
	eb, err := s.ToExchangeBackup()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSelectorSuite) TestNewExchangeRestore() {
	t := suite.T()
	er := NewExchangeRestore(nil)
	assert.Equal(t, er.Service, ServiceExchange)
	assert.NotZero(t, er.Scopes())
}

func (suite *ExchangeSelectorSuite) TestToExchangeRestore() {
	t := suite.T()
	eb := NewExchangeRestore(nil)
	s := eb.Selector
	eb, err := s.ToExchangeRestore()
	require.NoError(t, err)
	assert.Equal(t, eb.Service, ServiceExchange)
	assert.NotZero(t, eb.Scopes())
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Contacts() {
	t := suite.T()

	const (
		user   = "user"
		folder = AnyTgt
		c1     = "c1"
		c2     = "c2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Exclude(sel.Contacts([]string{folder}, []string{c1, c2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeContactFolder: folder,
			ExchangeContact:       join(c1, c2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Contacts() {
	t := suite.T()

	const (
		user   = "user"
		folder = AnyTgt
		c1     = "c1"
		c2     = "c2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Include(sel.Contacts([]string{folder}, []string{c1, c2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeContactFolder: folder,
			ExchangeContact:       join(c1, c2),
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContact)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_ContactFolders() {
	t := suite.T()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Exclude(sel.ContactFolders([]string{f1, f2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeContactFolder: join(f1, f2),
			ExchangeContact:       AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_ContactFolders() {
	t := suite.T()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Include(sel.ContactFolders([]string{f1, f2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeContactFolder: join(f1, f2),
			ExchangeContact:       AnyTgt,
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeContactFolder)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Events() {
	t := suite.T()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
		c1   = "c1"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Exclude(sel.Events([]string{c1}, []string{e1, e2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeEventCalendar: c1,
			ExchangeEvent:         join(e1, e2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_EventCalendars() {
	t := suite.T()

	const (
		user = "user"
		c1   = "c1"
		c2   = "c2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Exclude(sel.EventCalendars([]string{c1, c2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeEventCalendar: join(c1, c2),
			ExchangeEvent:         AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Events() {
	t := suite.T()

	const (
		user = "user"
		e1   = "e1"
		e2   = "e2"
		c1   = "c1"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Include(sel.Events([]string{c1}, []string{e1, e2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeEventCalendar: c1,
			ExchangeEvent:         join(e1, e2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_EventCalendars() {
	t := suite.T()

	const (
		user = "user"
		c1   = "c1"
		c2   = "c2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Include(sel.EventCalendars([]string{c1, c2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeEventCalendar: join(c1, c2),
			ExchangeEvent:         AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_Mails() {
	t := suite.T()

	const (
		user   = "user"
		folder = AnyTgt
		m1     = "m1"
		m2     = "m2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Exclude(sel.Mails([]string{folder}, []string{m1, m2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeMailFolder: folder,
			ExchangeMail:       join(m1, m2),
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_Mails() {
	t := suite.T()

	const (
		user   = "user"
		folder = AnyTgt
		m1     = "m1"
		m2     = "m2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Include(sel.Mails([]string{folder}, []string{m1, m2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeMailFolder: folder,
			ExchangeMail:       join(m1, m2),
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMail)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_MailFolders() {
	t := suite.T()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Exclude(sel.MailFolders([]string{f1, f2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeMailFolder: join(f1, f2),
			ExchangeMail:       AnyTgt,
		},
	)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_MailFolders() {
	t := suite.T()

	const (
		user = "user"
		f1   = "f1"
		f2   = "f2"
	)

	sel := NewExchangeBackup([]string{user})
	sel.Include(sel.MailFolders([]string{f1, f2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		ExchangeScope(scopes[0]),
		map[categorizer]string{
			ExchangeMailFolder: join(f1, f2),
			ExchangeMail:       AnyTgt,
		},
	)

	assert.Equal(t, sel.Scopes()[0].Category(), ExchangeMailFolder)
}

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Exclude_AllData() {
	t := suite.T()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel := NewExchangeBackup([]string{u1, u2})
	sel.Exclude(sel.AllData())
	scopes := sel.Excludes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
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

func (suite *ExchangeSelectorSuite) TestExchangeSelector_Include_AllData() {
	t := suite.T()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel := NewExchangeBackup([]string{u1, u2})
	sel.Include(sel.AllData())
	scopes := sel.Includes
	require.Len(t, scopes, 3)

	for _, sc := range scopes {
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

func (suite *ExchangeSelectorSuite) TestExchangeBackup_Scopes() {
	eb := NewExchangeBackup(Any())
	eb.Include(eb.AllData())

	scopes := eb.Scopes()
	assert.Len(suite.T(), scopes, 3)

	for _, sc := range scopes {
		cat := sc.Category()
		suite.T().Run(cat.String(), func(t *testing.T) {
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
		{ExchangeEvent, ExchangeEventCalendar, assert.NotEqual},
		{ExchangeEventCalendar, ExchangeEventCalendar, assert.Equal},
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
			eb := NewExchangeBackup(Any())
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
		{ExchangeCategoryUnknown, ExchangeUser, assert.True},
		{ExchangeContact, ExchangeContactFolder, assert.True},
		{ExchangeContact, ExchangeMailFolder, assert.False},
		{ExchangeContactFolder, ExchangeContact, assert.True},
		{ExchangeContactFolder, ExchangeMailFolder, assert.False},
		{ExchangeContactFolder, ExchangeEventCalendar, assert.False},
		{ExchangeEvent, ExchangeUser, assert.True},
		{ExchangeEvent, ExchangeContact, assert.False},
		{ExchangeEvent, ExchangeEventCalendar, assert.True},
		{ExchangeEvent, ExchangeContactFolder, assert.False},
		{ExchangeEventCalendar, ExchangeEvent, assert.True},
		{ExchangeEventCalendar, ExchangeEventCalendar, assert.True},
		{ExchangeEventCalendar, ExchangeUser, assert.True},
		{ExchangeEventCalendar, ExchangeCategoryUnknown, assert.False},
		{ExchangeMail, ExchangeMailFolder, assert.True},
		{ExchangeMail, ExchangeContact, assert.False},
		{ExchangeMailFolder, ExchangeMail, assert.True},
		{ExchangeMailFolder, ExchangeContactFolder, assert.False},
		{ExchangeMailFolder, ExchangeEventCalendar, assert.False},
		{ExchangeUser, ExchangeUser, assert.True},
		{ExchangeUser, ExchangeCategoryUnknown, assert.True},
		{ExchangeUser, ExchangeMail, assert.True},
		{ExchangeUser, ExchangeEventCalendar, assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.is.String()+test.expect.String(), func(t *testing.T) {
			eb := NewExchangeBackup(Any())
			eb.Includes = []scope{
				{scopeKeyCategory: filters.Identity(test.is.String())},
			}
			scope := eb.Scopes()[0]
			test.check(t, scope.IncludesCategory(test.expect))
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_Get() {
	eb := NewExchangeBackup(Any())
	eb.Include(eb.AllData())

	scopes := eb.Scopes()

	table := []exchangeCategory{
		ExchangeContact,
		ExchangeContactFolder,
		ExchangeEvent,
		ExchangeMail,
		ExchangeMailFolder,
	}
	for _, test := range table {
		suite.T().Run(test.String(), func(t *testing.T) {
			for _, sc := range scopes {
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
	es := NewExchangeRestore(Any())

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
	)

	infoWith := func(itype details.ItemType) details.ItemInfo {
		return details.ItemInfo{
			Exchange: &details.ExchangeInfo{
				ItemType:    itype,
				ContactName: name,
				EventRecurs: true,
				EventStart:  now,
				Organizer:   organizer,
				Sender:      sender,
				Subject:     subject,
				Received:    now,
			},
		}
	}

	table := []struct {
		name   string
		itype  details.ItemType
		scope  []ExchangeScope
		expect assert.BoolAssertionFunc
	}{
		{"any mail with a sender", details.ExchangeMail, es.MailSender(AnyTgt), assert.True},
		{"no mail, regardless of sender", details.ExchangeMail, es.MailSender(NoneTgt), assert.False},
		{"mail from a different sender", details.ExchangeMail, es.MailSender("magoo@ma.goo"), assert.False},
		{"mail from the matching sender", details.ExchangeMail, es.MailSender(sender), assert.True},
		{"mail with any subject", details.ExchangeMail, es.MailSubject(AnyTgt), assert.True},
		{"mail with none subject", details.ExchangeMail, es.MailSubject(NoneTgt), assert.False},
		{"mail with a different subject", details.ExchangeMail, es.MailSubject("fancy"), assert.False},
		{"mail with the matching subject", details.ExchangeMail, es.MailSubject(subject), assert.True},
		{"mail with a substring subject match", details.ExchangeMail, es.MailSubject(subject[5:9]), assert.True},
		{"mail received after the epoch", details.ExchangeMail, es.MailReceivedAfter(common.FormatTime(epoch)), assert.True},
		{"mail received after now", details.ExchangeMail, es.MailReceivedAfter(common.FormatTime(now)), assert.False},
		{
			"mail received after sometime later",
			details.ExchangeMail,
			es.MailReceivedAfter(common.FormatTime(future)),
			assert.False,
		},
		{
			"mail received before the epoch",
			details.ExchangeMail,
			es.MailReceivedBefore(common.FormatTime(epoch)),
			assert.False,
		},
		{"mail received before now", details.ExchangeMail, es.MailReceivedBefore(common.FormatTime(now)), assert.False},
		{
			"mail received before sometime later",
			details.ExchangeMail,
			es.MailReceivedBefore(common.FormatTime(future)),
			assert.True,
		},
		{"event with any organizer", details.ExchangeEvent, es.EventOrganizer(AnyTgt), assert.True},
		{"event with none organizer", details.ExchangeEvent, es.EventOrganizer(NoneTgt), assert.False},
		{"event with a different organizer", details.ExchangeEvent, es.EventOrganizer("fancy"), assert.False},
		{"event with the matching organizer", details.ExchangeEvent, es.EventOrganizer(organizer), assert.True},
		{"event that recurs", details.ExchangeEvent, es.EventRecurs("true"), assert.True},
		{"event that does not recur", details.ExchangeEvent, es.EventRecurs("false"), assert.False},
		{"event starting after the epoch", details.ExchangeEvent, es.EventStartsAfter(common.FormatTime(epoch)), assert.True},
		{"event starting after now", details.ExchangeEvent, es.EventStartsAfter(common.FormatTime(now)), assert.False},
		{
			"event starting after sometime later",
			details.ExchangeEvent,
			es.EventStartsAfter(common.FormatTime(future)),
			assert.False,
		},
		{
			"event starting before the epoch",
			details.ExchangeEvent,
			es.EventStartsBefore(common.FormatTime(epoch)),
			assert.False,
		},
		{"event starting before now", details.ExchangeEvent, es.EventStartsBefore(common.FormatTime(now)), assert.False},
		{
			"event starting before sometime later",
			details.ExchangeEvent,
			es.EventStartsBefore(common.FormatTime(future)),
			assert.True,
		},
		{"event with any subject", details.ExchangeEvent, es.EventSubject(AnyTgt), assert.True},
		{"event with none subject", details.ExchangeEvent, es.EventSubject(NoneTgt), assert.False},
		{"event with a different subject", details.ExchangeEvent, es.EventSubject("fancy"), assert.False},
		{"event with the matching subject", details.ExchangeEvent, es.EventSubject(subject), assert.True},
		{"contact with a different name", details.ExchangeContact, es.ContactName("blarps"), assert.False},
		{"contact with the same name", details.ExchangeContact, es.ContactName(name), assert.True},
		{"contact with a subname search", details.ExchangeContact, es.ContactName(name[2:5]), assert.True},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scopes := setScopesToDefault(test.scope)
			for _, scope := range scopes {
				test.expect(t, scope.matchesInfo(infoWith(test.itype)))
			}
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeScope_MatchesPath() {
	const (
		usr  = "userID"
		fld1 = "mailFolder"
		fld2 = "subFolder"
		mail = "mailID"
	)

	var (
		pth   = stubPath(suite.T(), usr, []string{fld1, fld2, mail}, path.EmailCategory)
		short = "thisisahashofsomekind"
		es    = NewExchangeRestore(Any())
	)

	table := []struct {
		name     string
		scope    []ExchangeScope
		shortRef string
		expect   assert.BoolAssertionFunc
	}{
		{"all items", es.AllData(), "", assert.True},
		{"all folders", es.MailFolders(Any()), "", assert.True},
		{"no folders", es.MailFolders(None()), "", assert.False},
		{"matching folder", es.MailFolders([]string{fld1}), "", assert.True},
		{"incomplete matching folder", es.MailFolders([]string{"mail"}), "", assert.False},
		{"non-matching folder", es.MailFolders([]string{"smarf"}), "", assert.False},
		{"non-matching folder substring", es.MailFolders([]string{fld1 + "_suffix"}), "", assert.False},
		{"matching folder prefix", es.MailFolders([]string{fld1}, PrefixMatch()), "", assert.True},
		{"incomplete folder prefix", es.MailFolders([]string{"mail"}, PrefixMatch()), "", assert.False},
		{"matching folder substring", es.MailFolders([]string{"Folder"}), "", assert.False},
		{"one of multiple folders", es.MailFolders([]string{"smarf", fld2}), "", assert.True},
		{"all mail", es.Mails(Any(), Any()), "", assert.True},
		{"no mail", es.Mails(Any(), None()), "", assert.False},
		{"matching mail", es.Mails(Any(), []string{mail}), "", assert.True},
		{"non-matching mail", es.Mails(Any(), []string{"smarf"}), "", assert.False},
		{"one of multiple mails", es.Mails(Any(), []string{"smarf", mail}), "", assert.True},
		{"mail short ref", es.Mails(Any(), []string{short}), short, assert.True},
		{"non-leaf short ref", es.Mails([]string{short}, []string{"foo"}), short, assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			scopes := setScopesToDefault(test.scope)
			var aMatch bool
			for _, scope := range scopes {
				pv := ExchangeMail.pathValues(pth)
				if matchesPathValues(scope, ExchangeMail, pv, short) {
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
		pth    path.Path
		expect map[exchangeCategory]string
	}{
		{
			ExchangeContact,
			stubPath(suite.T(), "uid", []string{"cFld", "cid"}, path.ContactsCategory),
			map[exchangeCategory]string{
				ExchangeContactFolder: "cFld",
				ExchangeContact:       "cid",
			},
		},
		{
			ExchangeEvent,
			stubPath(suite.T(), "uid", []string{"eCld", "eid"}, path.EventsCategory),
			map[exchangeCategory]string{
				ExchangeEventCalendar: "eCld",
				ExchangeEvent:         "eid",
			},
		},
		{
			ExchangeMail,
			stubPath(suite.T(), "uid", []string{"mFld", "mid"}, path.EmailCategory),
			map[exchangeCategory]string{
				ExchangeMailFolder: "mFld",
				ExchangeMail:       "mid",
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeRestore_Reduce() {
	var (
		contact            = stubRepoRef(path.ExchangeService, path.ContactsCategory, "uid", "cfld", "cid")
		event              = stubRepoRef(path.ExchangeService, path.EventsCategory, "uid", "ecld", "eid")
		mail               = stubRepoRef(path.ExchangeService, path.EmailCategory, "uid", "mfld", "mid")
		contactInSubFolder = stubRepoRef(path.ExchangeService, path.ContactsCategory, "uid", "cfld1/cfld2", "cid")
	)

	makeDeets := func(refs ...string) *details.Details {
		deets := &details.Details{
			DetailsModel: details.DetailsModel{
				Entries: []details.DetailsEntry{},
			},
		}

		for _, r := range refs {
			itype := details.UnknownType

			switch r {
			case contact:
				itype = details.ExchangeContact
			case event:
				itype = details.ExchangeEvent
			case mail:
				itype = details.ExchangeMail
			}

			deets.Entries = append(deets.Entries, details.DetailsEntry{
				RepoRef: r,
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: itype,
					},
				},
			})
		}

		return deets
	}

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
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			[]string{},
		},
		{
			"contact only",
			makeDeets(contact),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(contact),
		},
		{
			"event only",
			makeDeets(event),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(event),
		},
		{
			"mail only",
			makeDeets(mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(mail),
		},
		{
			"all",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(contact, event, mail),
		},
		{
			"only match contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.Contacts([]string{"cfld"}, []string{"cid"}))
				return er
			},
			arr(contact),
		},
		{
			"only match contactInSubFolder",
			makeDeets(contactInSubFolder, contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.ContactFolders([]string{"cfld1/cfld2"}))
				return er
			},
			arr(contactInSubFolder),
		},
		{
			"only match contactInSubFolder by prefix",
			makeDeets(contactInSubFolder, contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.ContactFolders([]string{"cfld1/cfld2"}, PrefixMatch()))
				return er
			},
			arr(contactInSubFolder),
		},
		{
			"only match contactInSubFolder by leaf folder",
			makeDeets(contactInSubFolder, contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.ContactFolders([]string{"cfld2"}))
				return er
			},
			arr(contactInSubFolder),
		},
		{
			"only match event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.Events([]string{"ecld"}, []string{"eid"}))
				return er
			},
			arr(event),
		},
		{
			"only match mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.Mails([]string{"mfld"}, []string{"mid"}))
				return er
			},
			arr(mail),
		},
		{
			"exclude contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Exclude(er.Contacts([]string{"cfld"}, []string{"cid"}))
				return er
			},
			arr(event, mail),
		},
		{
			"exclude event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Exclude(er.Events([]string{"ecld"}, []string{"eid"}))
				return er
			},
			arr(contact, mail),
		},
		{
			"exclude mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Exclude(er.Mails([]string{"mfld"}, []string{"mid"}))
				return er
			},
			arr(contact, event),
		},
		{
			"filter on mail subject",
			func() *details.Details {
				ds := makeDeets(mail)
				for i := range ds.Entries {
					ds.Entries[i].Exchange.Subject = "has a subject"
				}
				return ds
			}(),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Filter(er.MailSubject("subj"))
				return er
			},
			arr(mail),
		},
		{
			"filter on mail subject multiple input categories",
			func() *details.Details {
				mds := makeDeets(mail)
				for i := range mds.Entries {
					mds.Entries[i].Exchange.Subject = "has a subject"
				}

				ds := makeDeets(contact, event)
				ds.Entries = append(ds.Entries, mds.Entries...)

				return ds
			}(),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Filter(er.MailSubject("subj"))
				return er
			},
			arr(mail),
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			errs := mock.NewAdder()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets, errs)
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
			assert.Empty(t, errs.Errs)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeRestore_Reduce_locationRef() {
	var (
		contact         = stubRepoRef(path.ExchangeService, path.ContactsCategory, "uid", "id5/id6", "cid")
		contactLocation = "conts/my_cont"
		event           = stubRepoRef(path.ExchangeService, path.EventsCategory, "uid", "id1/id2", "eid")
		eventLocation   = "cal/my_cal"
		mail            = stubRepoRef(path.ExchangeService, path.EmailCategory, "uid", "id3/id4", "mid")
		mailLocation    = "inbx/my_mail"
	)

	makeDeets := func(refs ...string) *details.Details {
		deets := &details.Details{
			DetailsModel: details.DetailsModel{
				Entries: []details.DetailsEntry{},
			},
		}

		for _, r := range refs {
			var (
				location string
				itype    = details.UnknownType
			)

			switch r {
			case contact:
				itype = details.ExchangeContact
				location = contactLocation
			case event:
				itype = details.ExchangeEvent
				location = eventLocation
			case mail:
				itype = details.ExchangeMail
				location = mailLocation
			}

			deets.Entries = append(deets.Entries, details.DetailsEntry{
				RepoRef:     r,
				LocationRef: location,
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: itype,
					},
				},
			})
		}

		return deets
	}

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
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			[]string{},
		},
		{
			"contact only",
			makeDeets(contact),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(contact),
		},
		{
			"event only",
			makeDeets(event),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(event),
		},
		{
			"mail only",
			makeDeets(mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(mail),
		},
		{
			"all",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(contact, event, mail),
		},
		{
			"only match contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.Contacts([]string{contactLocation}, []string{"cid"}))
				return er
			},
			arr(contact),
		},
		{
			"only match event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.Events([]string{eventLocation}, []string{"eid"}))
				return er
			},
			arr(event),
		},
		{
			"only match mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore([]string{"uid"})
				er.Include(er.Mails([]string{mailLocation}, []string{"mid"}))
				return er
			},
			arr(mail),
		},
		{
			"exclude contact",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Exclude(er.Contacts([]string{contactLocation}, []string{"cid"}))
				return er
			},
			arr(event, mail),
		},
		{
			"exclude event",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Exclude(er.Events([]string{eventLocation}, []string{"eid"}))
				return er
			},
			arr(contact, mail),
		},
		{
			"exclude mail",
			makeDeets(contact, event, mail),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Exclude(er.Mails([]string{mailLocation}, []string{"mid"}))
				return er
			},
			arr(contact, event),
		},
		{
			"filter on mail subject",
			func() *details.Details {
				ds := makeDeets(mail)
				for i := range ds.Entries {
					ds.Entries[i].Exchange.Subject = "has a subject"
				}
				return ds
			}(),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Filter(er.MailSubject("subj"))
				return er
			},
			arr(mail),
		},
		{
			"filter on mail subject multiple input categories",
			func() *details.Details {
				mds := makeDeets(mail)
				for i := range mds.Entries {
					mds.Entries[i].Exchange.Subject = "has a subject"
				}

				ds := makeDeets(contact, event)
				ds.Entries = append(ds.Entries, mds.Entries...)

				return ds
			}(),
			func() *ExchangeRestore {
				er := NewExchangeRestore(Any())
				er.Include(er.AllData())
				er.Filter(er.MailSubject("subj"))
				return er
			},
			arr(mail),
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			errs := mock.NewAdder()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets, errs)
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
			assert.Empty(t, errs.Errs)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestScopesByCategory() {
	var (
		es       = NewExchangeRestore(Any())
		allData  = es.AllData()
		contacts = es.ContactFolders(Any())
		events   = es.EventCalendars(Any())
		mail     = es.MailFolders(Any())
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
	cats := map[path.CategoryType]exchangeCategory{
		path.ContactsCategory: ExchangeContact,
		path.EventsCategory:   ExchangeEvent,
		path.EmailCategory:    ExchangeMail,
	}

	table := []struct {
		name   string
		scopes input
		expect expect
	}{
		{"contacts only", makeInput(contacts), expect{1, 0, 0}},
		{"events only", makeInput(events), expect{0, 1, 0}},
		{"mail only", makeInput(mail), expect{0, 0, 1}},
		{"all", makeInput(allData, contacts, events, mail), expect{2, 2, 2}},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := scopesByCategory[ExchangeScope](test.scopes, cats, false)
			assert.Len(t, result[ExchangeContact], test.expect.contact)
			assert.Len(t, result[ExchangeEvent], test.expect.event)
			assert.Len(t, result[ExchangeMail], test.expect.mail)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestPasses() {
	short := "thisisahashofsomekind"
	entry := details.DetailsEntry{ShortRef: short}

	const (
		mid = "mailID"
		cat = ExchangeMail
	)

	var (
		es        = NewExchangeRestore(Any())
		otherMail = setScopesToDefault(es.Mails(Any(), []string{"smarf"}))
		mail      = setScopesToDefault(es.Mails(Any(), []string{mid}))
		noMail    = setScopesToDefault(es.Mails(Any(), None()))
		allMail   = setScopesToDefault(es.Mails(Any(), Any()))
		pth       = stubPath(suite.T(), "user", []string{"folder", mid}, path.EmailCategory)
	)

	table := []struct {
		name                        string
		excludes, filters, includes []ExchangeScope
		expect                      assert.BoolAssertionFunc
	}{
		{"empty", nil, nil, nil, assert.False},
		{"in Mail", nil, nil, mail, assert.True},
		{"in Other", nil, nil, otherMail, assert.False},
		{"in no Mail", nil, nil, noMail, assert.False},
		{"ex None filter mail", allMail, mail, nil, assert.False},
		{"ex Mail", mail, nil, allMail, assert.False},
		{"ex Other", otherMail, nil, allMail, assert.True},
		{"in and ex Mail", mail, nil, mail, assert.False},
		{"filter Mail", nil, mail, allMail, assert.True},
		{"filter Other", nil, otherMail, allMail, assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := passes(
				cat,
				cat.pathValues(pth),
				entry,
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
		es                  = NewExchangeRestore(Any())
		noMail              = setScopesToDefault(es.Mails(None(), None()))
		does                = setScopesToDefault(es.Mails(Any(), []string{target}))
		doesNot             = setScopesToDefault(es.Mails(Any(), []string{"smarf"}))
		wrongType           = setScopesToDefault(es.Contacts(Any(), Any()))
		wrongTypeGoodTarget = setScopesToDefault(es.Contacts(Any(), Any()))
	)

	table := []struct {
		name   string
		scopes []ExchangeScope
		expect assert.BoolAssertionFunc
	}{
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
		es           = NewExchangeRestore(Any())
		specificMail = setScopesToDefault(es.Mails(Any(), []string{"email"}))
		anyMail      = setScopesToDefault(es.Mails(Any(), Any()))
	)

	table := []struct {
		name   string
		scopes []ExchangeScope
		cat    exchangeCategory
		expect assert.BoolAssertionFunc
	}{
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
	t := suite.T()

	contactPath := stubPath(t, "user", []string{"cfolder", "contactitem"}, path.ContactsCategory)
	contactMap := map[categorizer]string{
		ExchangeContactFolder: contactPath.Folder(),
		ExchangeContact:       contactPath.Item(),
	}
	eventPath := stubPath(t, "user", []string{"ecalendar", "eventitem"}, path.EventsCategory)
	eventMap := map[categorizer]string{
		ExchangeEventCalendar: eventPath.Folder(),
		ExchangeEvent:         eventPath.Item(),
	}
	mailPath := stubPath(t, "user", []string{"mfolder", "mailitem"}, path.EmailCategory)
	mailMap := map[categorizer]string{
		ExchangeMailFolder: mailPath.Folder(),
		ExchangeMail:       mailPath.Item(),
	}

	table := []struct {
		cat    exchangeCategory
		path   path.Path
		expect map[categorizer]string
	}{
		{ExchangeContact, contactPath, contactMap},
		{ExchangeEvent, eventPath, eventMap},
		{ExchangeMail, mailPath, mailMap},
	}
	for _, test := range table {
		suite.T().Run(string(test.cat), func(t *testing.T) {
			assert.Equal(t, test.cat.pathValues(test.path), test.expect)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestExchangeCategory_PathKeys() {
	contact := []categorizer{ExchangeContactFolder, ExchangeContact}
	event := []categorizer{ExchangeEventCalendar, ExchangeEvent}
	mail := []categorizer{ExchangeMailFolder, ExchangeMail}
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
		suite.T().Run(string(test.cat), func(t *testing.T) {
			assert.Equal(t, test.cat.pathKeys(), test.expect)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestCategoryFromItemType() {
	table := []struct {
		name   string
		input  details.ItemType
		expect exchangeCategory
	}{
		{
			name:   "contact",
			input:  details.ExchangeContact,
			expect: ExchangeContact,
		},
		{
			name:   "event",
			input:  details.ExchangeEvent,
			expect: ExchangeEvent,
		},
		{
			name:   "mail",
			input:  details.ExchangeMail,
			expect: ExchangeMail,
		},
		{
			name:   "unknown",
			input:  details.UnknownType,
			expect: ExchangeCategoryUnknown,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result := categoryFromItemType(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func (suite *ExchangeSelectorSuite) TestCategory_PathType() {
	table := []struct {
		cat      exchangeCategory
		pathType path.CategoryType
	}{
		{ExchangeCategoryUnknown, path.UnknownCategory},
		{ExchangeContact, path.ContactsCategory},
		{ExchangeContactFolder, path.ContactsCategory},
		{ExchangeEvent, path.EventsCategory},
		{ExchangeEventCalendar, path.EventsCategory},
		{ExchangeMail, path.EmailCategory},
		{ExchangeMailFolder, path.EmailCategory},
		{ExchangeUser, path.UnknownCategory},
		{ExchangeFilterMailSender, path.EmailCategory},
		{ExchangeFilterMailSubject, path.EmailCategory},
		{ExchangeFilterMailReceivedAfter, path.EmailCategory},
		{ExchangeFilterMailReceivedBefore, path.EmailCategory},
		{ExchangeFilterContactName, path.ContactsCategory},
		{ExchangeFilterEventOrganizer, path.EventsCategory},
		{ExchangeFilterEventRecurs, path.EventsCategory},
		{ExchangeFilterEventStartsAfter, path.EventsCategory},
		{ExchangeFilterEventStartsBefore, path.EventsCategory},
		{ExchangeFilterEventSubject, path.EventsCategory},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.pathType, test.cat.PathType())
		})
	}
}
