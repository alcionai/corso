package restore

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli/utils"
	ctesting "github.com/alcionai/corso/internal/testing"
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
		{"restore exchange", restoreCommand, expectUse, exchangeRestoreCmd.Short, restoreExchangeCmd},
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
			ctesting.AreSameFunc(t, test.expectRunE, child.RunE)
		})
	}
}

func (suite *ExchangeSuite) TestValidateExchangeRestoreFlags() {
	stub := []string{"id-stub"}
	table := []struct {
		name                                                          string
		contacts, contactFolders, emails, emailFolders, events, users []string
		backupID                                                      string
		expect                                                        assert.ErrorAssertionFunc
	}{
		{
			name:     "only backupid",
			backupID: "bid",
			expect:   assert.NoError,
		},
		{
			name:           "all values populated",
			backupID:       "bid",
			contacts:       stub,
			contactFolders: stub,
			emails:         stub,
			emailFolders:   stub,
			events:         stub,
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
			expect:         assert.Error,
		},
		{
			name:         "no contact folders",
			backupID:     "bid",
			contacts:     stub,
			emails:       stub,
			emailFolders: stub,
			events:       stub,
			users:        stub,
			expect:       assert.Error,
		},
		{
			name:           "no email folders",
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
				test.users,
				test.backupID,
			))
		})
	}
}

func (suite *ExchangeSuite) TestIncludeExchangeRestoreDataSelectors() {
	stub := []string{"id-stub"}
	all := []string{utils.Wildcard}
	table := []struct {
		name                                                          string
		contacts, contactFolders, emails, emailFolders, events, users []string
		expectIncludeLen                                              int
	}{
		{
			name:             "no selectors",
			expectIncludeLen: 0,
		},
		{
			name:             "all users",
			users:            all,
			expectIncludeLen: 1,
		},
		{
			name:             "single user",
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "multiple users",
			users:            []string{"fnord", "smarf"},
			expectIncludeLen: 1,
		},
		{
			name:             "all users, all data",
			contacts:         all,
			contactFolders:   all,
			emails:           all,
			emailFolders:     all,
			events:           all,
			users:            all,
			expectIncludeLen: 3,
		},
		{
			name:             "all users, all folders",
			contactFolders:   all,
			emailFolders:     all,
			users:            all,
			expectIncludeLen: 2,
		},
		{
			name:             "single user, single of each data",
			contacts:         stub,
			contactFolders:   stub,
			emails:           stub,
			emailFolders:     stub,
			events:           stub,
			users:            stub,
			expectIncludeLen: 3,
		},
		{
			name:             "single user, single of each folder",
			contactFolders:   stub,
			emailFolders:     stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "all users, contacts",
			contacts:         all,
			contactFolders:   stub,
			users:            all,
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
			name:             "all users, emails",
			emails:           all,
			emailFolders:     stub,
			users:            all,
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
			name:             "all users, events",
			events:           all,
			users:            all,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, events",
			events:           stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "all users, contacts + email",
			contacts:         all,
			contactFolders:   all,
			emails:           all,
			emailFolders:     all,
			users:            all,
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
			name:             "all users, email + event",
			emails:           all,
			emailFolders:     all,
			events:           all,
			users:            all,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, email + event",
			emails:           stub,
			emailFolders:     stub,
			events:           stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "all users, event + contact",
			contacts:         all,
			contactFolders:   all,
			events:           all,
			users:            all,
			expectIncludeLen: 2,
		},
		{
			name:             "single users, event + contact",
			contacts:         stub,
			contactFolders:   stub,
			events:           stub,
			users:            stub,
			expectIncludeLen: 2,
		},
		{
			name:             "many users, events",
			events:           []string{"foo", "bar"},
			users:            []string{"fnord", "smarf"},
			expectIncludeLen: 2,
		},
		{
			name:             "many users, events + contacts",
			contacts:         []string{"foo", "bar"},
			contactFolders:   []string{"foo", "bar"},
			events:           []string{"foo", "bar"},
			users:            []string{"fnord", "smarf"},
			expectIncludeLen: 6,
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
				test.users)
			assert.Equal(t, test.expectIncludeLen, len(sel.Includes))
		})
	}
}

func (suite *ExchangeSuite) TestIncludeExchangeRestoreInfoSelectors() {
	stub := []string{"id-stub"}
	stubs := []string{"a-stub", "b-stub"}
	all := []string{utils.Wildcard}
	table := []struct {
		name                           string
		after, before, sender, subject []string
		expectIncludeLen               int
	}{
		{
			name:             "no selectors",
			expectIncludeLen: 0,
		},
		{
			name:             "all receivedAfter",
			after:            all,
			expectIncludeLen: 1,
		},
		{
			name:             "single receivedAfter",
			after:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "multiple receivedAfter",
			after:            stubs,
			expectIncludeLen: 1,
		},
		{
			name:             "all receivedBefore",
			before:           all,
			expectIncludeLen: 1,
		},
		{
			name:             "single receivedBefore",
			before:           stub,
			expectIncludeLen: 1,
		},
		{
			name:             "multiple receivedBefore",
			before:           stubs,
			expectIncludeLen: 1,
		},
		{
			name:             "all senders",
			sender:           all,
			expectIncludeLen: 1,
		},
		{
			name:             "single sender",
			sender:           stub,
			expectIncludeLen: 1,
		},
		{
			name:             "multiple senders",
			sender:           stubs,
			expectIncludeLen: 1,
		},
		{
			name:             "all subjects",
			subject:          all,
			expectIncludeLen: 1,
		},
		{
			name:             "single subject",
			subject:          stub,
			expectIncludeLen: 1,
		},
		{
			name:             "multiple subjects",
			subject:          stubs,
			expectIncludeLen: 1,
		},
		{
			name:             "one of each",
			after:            stub,
			before:           stub,
			sender:           stub,
			subject:          stub,
			expectIncludeLen: 4,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			includeExchangeRestoreInfoSelectors(
				sel,
				test.after,
				test.before,
				test.sender,
				test.subject)
			assert.Equal(t, test.expectIncludeLen, len(sel.Includes))
		})
	}
}
