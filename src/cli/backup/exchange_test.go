package backup

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/selectors"
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
		{"create exchange", createCommand, expectUse, exchangeCreateCmd().Short, createExchangeCmd},
		{"list exchange", listCommand, expectUse, exchangeListCmd().Short, listExchangeCmd},
		{"details exchange", detailsCommand, expectUse, exchangeDetailsCmd().Short, detailsExchangeCmd},
		{"delete exchange", deleteCommand, expectUse, exchangeDeleteCmd().Short, deleteExchangeCmd},
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

func (suite *ExchangeSuite) TestValidateBackupCreateFlags() {
	table := []struct {
		name       string
		a          bool
		user, data []string
		expect     assert.ErrorAssertionFunc
	}{
		{
			name:   "no users, not any",
			expect: assert.Error,
		},
		{
			name:   "any and data",
			a:      true,
			data:   []string{dataEmail},
			expect: assert.Error,
		},
		{
			name:   "unrecognized data",
			user:   []string{"fnord"},
			data:   []string{"smurfs"},
			expect: assert.Error,
		},
		{
			name:   "users, not any",
			user:   []string{"fnord"},
			expect: assert.NoError,
		},
		{
			name:   "no users, any",
			a:      true,
			expect: assert.NoError,
		},
		{
			name:   "users, any",
			a:      true,
			user:   []string{"fnord"},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, validateExchangeBackupCreateFlags(test.a, test.user, test.data))
		})
	}
}

func (suite *ExchangeSuite) TestExchangeBackupCreateSelectors() {
	table := []struct {
		name             string
		a                bool
		user, data       []string
		expectIncludeLen int
	}{
		{
			name:             "any",
			a:                true,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, no data",
			user:             []string{utils.Wildcard},
			expectIncludeLen: 3,
		},
		{
			name:             "single user, no data",
			user:             []string{"u1"},
			expectIncludeLen: 3,
		},
		{
			name:             "any users, contacts",
			user:             []string{utils.Wildcard},
			data:             []string{dataContacts},
			expectIncludeLen: 1,
		},
		{
			name:             "single user, contacts",
			user:             []string{"u1"},
			data:             []string{dataContacts},
			expectIncludeLen: 1,
		},
		{
			name:             "any users, email",
			user:             []string{utils.Wildcard},
			data:             []string{dataEmail},
			expectIncludeLen: 1,
		},
		{
			name:             "single user, email",
			user:             []string{"u1"},
			data:             []string{dataEmail},
			expectIncludeLen: 1,
		},
		{
			name:             "any users, events",
			user:             []string{utils.Wildcard},
			data:             []string{dataEvents},
			expectIncludeLen: 1,
		},
		{
			name:             "single user, events",
			user:             []string{"u1"},
			data:             []string{dataEvents},
			expectIncludeLen: 1,
		},
		{
			name:             "any users, contacts + email",
			user:             []string{utils.Wildcard},
			data:             []string{dataContacts, dataEmail},
			expectIncludeLen: 2,
		},
		{
			name:             "single user, contacts + email",
			user:             []string{"u1"},
			data:             []string{dataContacts, dataEmail},
			expectIncludeLen: 2,
		},
		{
			name:             "any users, email + events",
			user:             []string{utils.Wildcard},
			data:             []string{dataEmail, dataEvents},
			expectIncludeLen: 2,
		},
		{
			name:             "single user, email + events",
			user:             []string{"u1"},
			data:             []string{dataEmail, dataEvents},
			expectIncludeLen: 2,
		},
		{
			name:             "any users, events + contacts",
			user:             []string{utils.Wildcard},
			data:             []string{dataEvents, dataContacts},
			expectIncludeLen: 2,
		},
		{
			name:             "single user, events + contacts",
			user:             []string{"u1"},
			data:             []string{dataEvents, dataContacts},
			expectIncludeLen: 2,
		},
		{
			name:             "many users, events",
			user:             []string{"fnord", "smarf"},
			data:             []string{dataEvents},
			expectIncludeLen: 1,
		},
		{
			name:             "many users, events + contacts",
			user:             []string{"fnord", "smarf"},
			data:             []string{dataEvents, dataContacts},
			expectIncludeLen: 2,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := exchangeBackupCreateSelectors(test.a, test.user, test.data)
			assert.Equal(t, test.expectIncludeLen, len(sel.Includes))
		})
	}
}

func (suite *ExchangeSuite) TestValidateBackupDetailFlags() {
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
			eventCalendars: stub,
			users:          stub,
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
			test.expect(t, validateExchangeBackupDetailFlags(
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

func (suite *ExchangeSuite) TestIncludeExchangeBackupDetailDataSelectors() {
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
			includeExchangeBackupDetailDataSelectors(
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

func (suite *ExchangeSuite) TestFilterExchangeBackupDetailInfoSelectors() {
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
			filterExchangeBackupDetailInfoSelectors(
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
