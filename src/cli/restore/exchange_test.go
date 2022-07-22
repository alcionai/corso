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
			name:           "any values populated",
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
	any := []string{utils.Wildcard}
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
			name:             "any users",
			users:            any,
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
			expectIncludeLen: 6,
		},
		{
			name:             "any users, any data",
			contacts:         any,
			contactFolders:   any,
			emails:           any,
			emailFolders:     any,
			events:           any,
			users:            any,
			expectIncludeLen: 3,
		},
		{
			name:             "any users, any folders",
			contactFolders:   any,
			emailFolders:     any,
			users:            any,
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
			name:             "any users, contacts",
			contacts:         any,
			contactFolders:   stub,
			users:            any,
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
			emails:           any,
			emailFolders:     stub,
			users:            any,
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
			events:           any,
			users:            any,
			expectIncludeLen: 1,
		},
		{
			name:             "single user, events",
			events:           stub,
			users:            stub,
			expectIncludeLen: 1,
		},
		{
			name:             "any users, contacts + email",
			contacts:         any,
			contactFolders:   any,
			emails:           any,
			emailFolders:     any,
			users:            any,
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
			emails:           any,
			emailFolders:     any,
			events:           any,
			users:            any,
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
			name:             "any users, event + contact",
			contacts:         any,
			contactFolders:   any,
			events:           any,
			users:            any,
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

func (suite *ExchangeSuite) TestFilterExchangeRestoreInfoSelectors() {
	stub := []string{"id-stub"}
	twoStubs := []string{"a-stub", "b-stub"}
	any := []string{utils.Wildcard}
	table := []struct {
		name                           string
		after, before, sender, subject string
		expectFilterLen                int
	}{
		{
			name:            "no selectors",
			expectFilterLen: 0,
		},
		{
			name:            "any receivedAfter",
			after:           any,
			expectFilterLen: 1,
		},
		{
			name:            "single receivedAfter",
			after:           stub,
			expectFilterLen: 1,
		},
		{
			name:            "multiple receivedAfter",
			after:           twoStubs,
			expectFilterLen: 1,
		},
		{
			name:            "any receivedBefore",
			before:          any,
			expectFilterLen: 1,
		},
		{
			name:            "single receivedBefore",
			before:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "multiple receivedBefore",
			before:          twoStubs,
			expectFilterLen: 1,
		},
		{
			name:            "any senders",
			sender:          any,
			expectFilterLen: 1,
		},
		{
			name:            "single sender",
			sender:          stub,
			expectFilterLen: 1,
		},
		{
			name:            "multiple senders",
			sender:          twoStubs,
			expectFilterLen: 1,
		},
		{
			name:            "any subjects",
			subject:         any,
			expectFilterLen: 1,
		},
		{
			name:            "single subject",
			subject:         stub,
			expectFilterLen: 1,
		},
		{
			name:            "multiple subjects",
			subject:         twoStubs,
			expectFilterLen: 1,
		},
		{
			name:            "one of each",
			after:           stub,
			before:          stub,
			sender:          stub,
			subject:         stub,
			expectFilterLen: 4,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			sel := selectors.NewExchangeRestore()
			filterExchangeRestoreInfoSelectors(
				sel,
				test.after,
				test.before,
				test.sender,
				test.subject)
			assert.Equal(t, test.expectFilterLen, len(sel.Filters))
		})
	}
}
