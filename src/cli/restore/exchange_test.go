package restore

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/selectors"
)

type ExchangeSuite struct {
	suite.Suite
}

func TestExchangeSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSuite))
}

func (suite *ExchangeSuite) TestAddExchangeCommands() {
	expectUse := exchangeServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore exchange", restoreCommand, expectUse, exchangeRestoreCmd().Short, restoreExchangeCmd},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: test.use}

			c := addExchangeCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)
		})
	}
}

func (suite *ExchangeSuite) TestValidateExchangeRestoreFlags() {
	stub := []string{"id-stub"}

	table := []struct {
		name                                                                          string
		contacts, contactFolders, emails, emailFolders, events, eventCalendars, users []string
		backupID                                                                      string
		expect                                                                        assert.ErrorAssertionFunc
	}{
		{
			name:     "only backupid",
			backupID: "bid",
			expect:   assert.NoError,
		},
		{
			name:           "any values populated",
			backupID:       "bid",
			contacts:       stub,
			contactFolders: stub,
			emails:         stub,
			emailFolders:   stub,
			events:         stub,
			eventCalendars: stub,
			users:          stub,
			expect:         assert.NoError,
		},
		{
			name:   "nothing populated",
			expect: assert.Error,
		},
		{
			name:           "no backup id",
			contacts:       stub,
			contactFolders: stub,
			emails:         stub,
			emailFolders:   stub,
			events:         stub,
			eventCalendars: stub,
			users:          stub,
			expect:         assert.Error,
		},
		{
			name:           "no users",
			backupID:       "bid",
			contacts:       stub,
			contactFolders: stub,
			emails:         stub,
			emailFolders:   stub,
			events:         stub,
			eventCalendars: stub,
			expect:         assert.Error,
		},
		{
			name:           "no contact folders",
			backupID:       "bid",
			contacts:       stub,
			emails:         stub,
			emailFolders:   stub,
			events:         stub,
			users:          stub,
			eventCalendars: stub,
			expect:         assert.Error,
		},
		{
			name:           "no email folders",
			backupID:       "bid",
			contacts:       stub,
			contactFolders: stub,
			emails:         stub,
			events:         stub,
			eventCalendars: stub,
			users:          stub,
			expect:         assert.Error,
		},
		{
			name:           "no event calendars",
			backupID:       "bid",
			contacts:       stub,
			contactFolders: stub,
			emails:         stub,
			events:         stub,
			users:          stub,
			expect:         assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, validateExchangeRestoreFlags(
				test.contacts,
				test.contactFolders,
				test.emails,
				test.emailFolders,
				test.events,
				test.eventCalendars,
				test.users,
				test.backupID,
			))
		})
	}
}

func (suite *ExchangeSuite) TestIncludeExchangeRestoreDataSelectors() {
	stub := []string{"id-stub"}
	a := []string{utils.Wildcard}

	table := []struct {
		name                                                                          string
		contacts, contactFolders, emails, emailFolders, events, eventCalendars, users []string
		expectIncludeLen                                                              int
	}{
		{
			name:             "no selectors",
			expectIncludeLen: 0,
		},
		{
			name:             "any users",
			users:            a,
			expectIncludeLen: 3,
		},
		{
			name:             "single user",
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "multiple users",
			users:            []string{"fnord", "smarf"},
			expectIncludeLen: 3,
		},
		{
			name:             "any users, any data",
			contacts:         a,
			contactFolders:   a,
			emails:           a,
			emailFolders:     a,
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, any folders",
			contactFolders:   a,
			emailFolders:     a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 3,
		},
		{
			name:             "single user, single of each data",
			contacts:         stub,
			contactFolders:   stub,
			emails:           stub,
			emailFolders:     stub,
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "single user, single of each folder",
			contactFolders:   stub,
			emailFolders:     stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, contacts",
			contacts:         a,
			contactFolders:   stub,
			users:            a,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, contacts",
			contacts:         stub,
			contactFolders:   stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, emails",
			emails:           a,
			emailFolders:     stub,
			users:            a,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, emails",
			emails:           stub,
			emailFolders:     stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, events",
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, events",
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, contacts + email",
			contacts:         a,
			contactFolders:   a,
			emails:           a,
			emailFolders:     a,
			users:            a,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, contacts + email",
			contacts:         stub,
			contactFolders:   stub,
			emails:           stub,
			emailFolders:     stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "any users, email + event",
			emails:           a,
			emailFolders:     a,
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, email + event",
			emails:           stub,
			emailFolders:     stub,
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "any users, event + contact",
			contacts:         a,
			contactFolders:   a,
			events:           a,
			eventCalendars:   a,
			users:            a,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, event + contact",
			contacts:         stub,
			contactFolders:   stub,
			events:           stub,
			eventCalendars:   stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "many users, events",
			events:           []string{"foo", "bar"},
			eventCalendars:   []string{"baz", "qux"},
			users:            []string{"fnord", "smarf"},
			expectIncludeLen: 1,
		},
		{
			name:             "many users, events + contacts",
			contacts:         []string{"foo", "bar"},
			contactFolders:   []string{"foo", "bar"},
			events:           []string{"foo", "bar"},
			eventCalendars:   []string{"foo", "bar"},
			users:            []string{"fnord", "smarf"},
			expectIncludeLen: 2,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			includeExchangeRestoreDataSelectors(
				sel,
				test.contacts,
				test.contactFolders,
				test.emails,
				test.emailFolders,
				test.events,
				test.eventCalendars,
				test.users)
			assert.Equal(t, test.expectIncludeLen, len(sel.Includes))
		})
	}
}

func (suite *ExchangeSuite) TestFilterExchangeRestoreInfoSelectors() {
	stub := "id-stub"

	table := []struct {
		name                                                       string
		contactName                                                string
		after, before, sender, subject                             string
		organizer, recurs, startsAfter, startsBefore, eventSubject string
		expectFilterLen                                            int
	}{
		{
			name:            "no selectors",
			expectFilterLen: 0,
		},
		{
			name:            "contactName",
			contactName:     stub,
			expectFilterLen: 1,
		},
		{
			name:            "receivedAfter",
			after:           stub,
			expectFilterLen: 1,
		},
		{
			name:            "receivedAfter",
			after:           stub,
			expectFilterLen: 1,
		},
		{
			name:            "receivedBefore",
			before:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "sender",
			sender:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "subject",
			subject:         stub,
			expectFilterLen: 1,
		},
		{
			name:            "organizer",
			organizer:       stub,
			expectFilterLen: 1,
		},
		{
			name:            "recurs",
			recurs:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "startsAfter",
			startsAfter:     stub,
			expectFilterLen: 1,
		},
		{
			name:            "startsBefore",
			startsBefore:    stub,
			expectFilterLen: 1,
		},
		{
			name:            "eventSubject",
			eventSubject:    stub,
			expectFilterLen: 1,
		},
		{
			name:            "one of each",
			contactName:     stub,
			after:           stub,
			before:          stub,
			sender:          stub,
			subject:         stub,
			organizer:       stub,
			recurs:          stub,
			startsAfter:     stub,
			startsBefore:    stub,
			eventSubject:    stub,
			expectFilterLen: 10,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			filterExchangeRestoreInfoSelectors(
				sel,
				test.contactName,
				test.after,
				test.before,
				test.sender,
				test.subject,
				test.organizer,
				test.recurs,
				test.startsAfter,
				test.startsBefore,
				test.eventSubject)
			assert.Equal(t, test.expectFilterLen, len(sel.Filters))
		})
	}
}
