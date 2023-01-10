package selectors_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/selectors/testdata"
)

type SelectorReduceSuite struct {
	suite.Suite
}

func TestSelectorReduceSuite(t *testing.T) {
	suite.Run(t, new(SelectorReduceSuite))
}

func (suite *SelectorReduceSuite) TestReduce() {
	ctx, flush := tester.NewContext()
	defer flush()

	allDetails := testdata.GetDetailsSet()
	table := []struct {
		name     string
		selFunc  func() selectors.Reducer
		expected []details.DetailsEntry
	}{
		{
			name: "ExchangeAllMail",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Mails(selectors.Any(), selectors.Any()))

				return sel
			},
			expected: testdata.ExchangeEmailItems,
		},
		{
			name: "ExchangeMailFolderPrefixMatch",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailInboxPath.Folder()},
				))

				return sel
			},
			expected: testdata.ExchangeEmailItems,
		},
		{
			name: "ExchangeMailSubject",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailSubject("foo"))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeMailSubjectExcludeItem",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailSender("a-person"))
				sel.Exclude(sel.Mails(
					selectors.Any(),
					[]string{testdata.ExchangeEmailItemPath2.ShortRef()},
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeMailSender",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailSender("a-person"))

				return sel
			},
			expected: []details.DetailsEntry{
				testdata.ExchangeEmailItems[0],
				testdata.ExchangeEmailItems[1],
			},
		},
		{
			name: "ExchangeMailReceivedTime",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailReceivedBefore(
					common.FormatTime(testdata.Time1.Add(time.Second)),
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeMailID",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Mails(
					selectors.Any(),
					[]string{testdata.ExchangeEmailItemPath1.Item()},
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeMailShortRef",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Mails(
					selectors.Any(),
					[]string{testdata.ExchangeEmailItemPath1.ShortRef()},
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeAllEventsAndMailWithSubject",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Events(
					selectors.Any(),
					selectors.Any(),
				))
				sel.Filter(sel.MailSubject("foo"))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeEventsAndMailWithSubject",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.EventSubject("foo"))
				sel.Filter(sel.MailSubject("foo"))

				return sel
			},
			expected: []details.DetailsEntry{},
		},
		{
			name: "ExchangeAll",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.AllData())

				return sel
			},
			expected: append(
				append(
					append(
						[]details.DetailsEntry{},
						testdata.ExchangeEmailItems...),
					testdata.ExchangeContactsItems...),
				testdata.ExchangeEventsItems...,
			),
		},
		{
			name: "ExchangeMailByFolder",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailBasePath.Folder()},
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		// TODO (keepers): all folders are treated as prefix-matches at this time.
		// so this test actually does nothing different.  In the future, we'll
		// need to amend the non-prefix folder tests to expect non-prefix matches.
		{
			name: "ExchangeMailByFolderPrefix",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailBasePath.Folder()},
					selectors.PrefixMatch(), // force prefix matching
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
		},
		{
			name: "ExchangeMailByFolderRoot",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailInboxPath.Folder()},
				))

				return sel
			},
			expected: testdata.ExchangeEmailItems,
		},
		{
			name: "ExchangeContactByFolder",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.ContactFolders(
					[]string{testdata.ExchangeContactsBasePath.Folder()},
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeContactsItems[0]},
		},
		{
			name: "ExchangeContactByFolderRoot",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.ContactFolders(
					[]string{testdata.ExchangeContactsRootPath.Folder()},
				))

				return sel
			},
			expected: testdata.ExchangeContactsItems,
		},

		{
			name: "ExchangeEventsByFolder",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.EventCalendars(
					[]string{testdata.ExchangeEventsBasePath.Folder()},
				))

				return sel
			},
			expected: []details.DetailsEntry{testdata.ExchangeEventsItems[0]},
		},
		{
			name: "ExchangeEventsByFolderRoot",
			selFunc: func() selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.EventCalendars(
					[]string{testdata.ExchangeEventsRootPath.Folder()},
				))

				return sel
			},
			expected: testdata.ExchangeEventsItems,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			output := test.selFunc().Reduce(ctx, allDetails)
			assert.ElementsMatch(t, test.expected, output.Entries)
		})
	}
}
