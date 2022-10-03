package backup

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
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
		{"create exchange", createCommand, expectUse + exchangeServiceCommandCreateUseSuffix, exchangeCreateCmd().Short, createExchangeCmd},
		{"list exchange", listCommand, expectUse, exchangeListCmd().Short, listExchangeCmd},
		{"details exchange", detailsCommand, expectUse + exchangeServiceCommandDetailsUseSuffix, exchangeDetailsCmd().Short, detailsExchangeCmd},
		{"delete exchange", deleteCommand, expectUse + exchangeServiceCommandDeleteUseSuffix, exchangeDeleteCmd().Short, deleteExchangeCmd},
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

func (suite *ExchangeSuite) TestExchangeBackupDetailsSelectors() {
	ctx := context.Background()

	for _, test := range testdata.ExchangeOptionDetailLookups {
		suite.T().Run(test.Name, func(t *testing.T) {
			output, err := runDetailsExchangeCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
			)
			assert.NoError(t, err)

			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *ExchangeSuite) TestExchangeBackupDetailsSelectorsBadBackupID() {
	t := suite.T()
	ctx := context.Background()
	backupGetter := &testdata.MockBackupGetter{}

	output, err := runDetailsExchangeCmd(
		ctx,
		backupGetter,
		"backup-ID",
		utils.ExchangeOpts{},
	)
	assert.Error(t, err)

	assert.Empty(t, output)
}

// TODO(ashmrtn): Uncomment these when the CLI validates flag input values.
//func (suite *ExchangeSuite) TestExchangeBackupDetailsSelectorsBadFormats() {
//	ctx := context.Background()
//
//	for _, test := range testdata.BadExchangeOptionsFormats {
//		suite.T().Run(test.Name, func(t *testing.T) {
//			output, err := runDetailsExchangeCmd(
//				ctx,
//				test.BackupGetter,
//				"backup-ID",
//				test.Opts,
//			)
//			assert.Error(t, err)
//
//			assert.Empty(t, output)
//		})
//	}
//}
