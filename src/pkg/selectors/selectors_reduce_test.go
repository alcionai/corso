package selectors_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type SelectorReduceSuite struct {
	tester.Suite
}

func TestSelectorReduceSuite(t *testing.T) {
	suite.Run(t, &SelectorReduceSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SelectorReduceSuite) TestReduce() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name     string
		selFunc  func(t *testing.T, wantVersion int) selectors.Reducer
		expected func(t *testing.T, wantVersion int) []details.Entry
	}{
		{
			name: "ExchangeAllMail",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Mails(selectors.Any(), selectors.Any()))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					-1)
			},
		},
		{
			name: "ExchangeMailFolderPrefixMatch",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailInboxPath.FolderLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					-1)
			},
		},
		{
			name: "ExchangeMailSubject",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailSubject("foo"))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeMailSubjectExcludeItem",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				deets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion)

				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailSender("a-person"))
				sel.Exclude(sel.Mails(
					selectors.Any(),
					[]string{deets[1].ShortRef},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeMailSender",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailSender("a-person"))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0, 1)
			},
		},
		{
			name: "ExchangeMailReceivedTime",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.MailReceivedBefore(
					dttm.Format(testdata.Time1.Add(time.Second)),
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeMailID",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Mails(
					selectors.Any(),
					[]string{testdata.ExchangeEmailItemPath1.ItemLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeMailShortRef",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				deets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion)

				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Mails(
					selectors.Any(),
					[]string{deets[0].ShortRef},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeAllEventsAndMailWithSubject",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.Events(
					selectors.Any(),
					selectors.Any(),
				))
				sel.Filter(sel.MailSubject("foo"))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeEventsAndMailWithSubject",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Filter(sel.EventSubject("foo"))
				sel.Filter(sel.MailSubject("foo"))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return []details.Entry{}
			},
		},
		{
			name: "ExchangeAll",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.AllData())

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return append(
					append(
						testdata.GetItemsForVersion(
							t,
							path.ExchangeService,
							path.EmailCategory,
							wantVersion,
							-1),
						testdata.GetItemsForVersion(
							t,
							path.ExchangeService,
							path.EventsCategory,
							wantVersion,
							-1)...),
					testdata.GetItemsForVersion(
						t,
						path.ExchangeService,
						path.ContactsCategory,
						wantVersion,
						-1)...)
			},
		},
		{
			name: "ExchangeMailByFolder",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailBasePath.FolderLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		// TODO (keepers): all folders are treated as prefix-matches at this time.
		// so this test actually does nothing different.  In the future, we'll
		// need to amend the non-prefix folder tests to expect non-prefix matches.
		{
			name: "ExchangeMailByFolderPrefix",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailBasePath.FolderLocation()},
					selectors.PrefixMatch(), // force prefix matching
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeMailByFolderRoot",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.MailFolders(
					[]string{testdata.ExchangeEmailInboxPath.FolderLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantVersion,
					-1)
			},
		},
		{
			name: "ExchangeContactByFolder",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.ContactFolders(
					[]string{testdata.ExchangeContactsBasePath.FolderLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.ContactsCategory,
					wantVersion,
					0)
			},
		},
		{
			name: "ExchangeContactByFolderRoot",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.ContactFolders(
					[]string{testdata.ExchangeContactsRootPath.FolderLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.ContactsCategory,
					wantVersion,
					-1)
			},
		},

		{
			name: "ExchangeEventsByFolder",
			selFunc: func(t *testing.T, wantVersion int) selectors.Reducer {
				sel := selectors.NewExchangeRestore(selectors.Any())
				sel.Include(sel.EventCalendars(
					[]string{testdata.ExchangeEventsBasePath.FolderLocation()},
				))

				return sel
			},
			expected: func(t *testing.T, wantVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EventsCategory,
					wantVersion,
					0)
			},
		},
	}

	for v := 0; v <= version.Backup; v++ {
		suite.Run(fmt.Sprintf("version%d", v), func() {
			for _, test := range table {
				suite.Run(test.name, func() {
					t := suite.T()

					allDetails := testdata.GetDetailsSetForVersion(t, v)
					output := test.selFunc(t, v).Reduce(ctx, allDetails, fault.New(true))
					assert.ElementsMatch(t, test.expected(t, v), output.Entries)
				})
			}
		})
	}
}
