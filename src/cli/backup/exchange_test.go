package backup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type ExchangeUnitSuite struct {
	tester.Suite
}

func TestExchangeUnitSuite(t *testing.T) {
	suite.Run(t, &ExchangeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExchangeUnitSuite) TestAddExchangeCommands() {
	expectUse := exchangeServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		flags       []string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			"create exchange",
			createCommand,
			expectUse + " " + exchangeServiceCommandCreateUseSuffix,
			exchangeCreateCmd().Short,
			[]string{"user", "data", "disable-incrementals", "fail-fast", "fetch-parallelism", "skip-reduce", "no-stats"},
			createExchangeCmd,
		},
		{
			"list exchange",
			listCommand,
			expectUse,
			exchangeListCmd().Short,
			[]string{"backup", "failed-items", "skipped-items", "recovered-errors"},
			listExchangeCmd,
		},
		{
			"details exchange",
			detailsCommand,
			expectUse + " " + exchangeServiceCommandDetailsUseSuffix,
			exchangeDetailsCmd().Short,
			[]string{
				"backup",
				"email",
				"email-folder",
				"email-subject",
				"email-sender",
				"email-received-after",
				"email-received-before",
				"event",
				"event-calendar",
				"event-subject",
				"event-organizer",
				"event-recurs",
				"event-starts-after",
				"event-starts-before",
				"contact",
				"contact-folder",
				"contact-name",
			},
			detailsExchangeCmd,
		},
		{
			"delete exchange",
			deleteCommand,
			expectUse + " " + exchangeServiceCommandDeleteUseSuffix,
			exchangeDeleteCmd().Short,
			[]string{"backup"},
			deleteExchangeCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

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

func (suite *ExchangeUnitSuite) TestValidateBackupCreateFlags() {
	table := []struct {
		name       string
		user, data []string
		expect     assert.ErrorAssertionFunc
	}{
		{
			name:   "no users or data",
			expect: assert.Error,
		},
		{
			name:   "no users only data",
			data:   []string{dataEmail},
			expect: assert.Error,
		},
		{
			name:   "unrecognized data category",
			user:   []string{"fnord"},
			data:   []string{"smurfs"},
			expect: assert.Error,
		},
		{
			name:   "only users no data",
			user:   []string{"fnord"},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			err := validateExchangeBackupCreateFlags(test.user, test.data)
			test.expect(t, err, clues.ToCore(err))
		})
	}
}

func (suite *ExchangeUnitSuite) TestExchangeBackupCreateSelectors() {
	table := []struct {
		name             string
		user, data       []string
		expectIncludeLen int
	}{
		{
			name:             "default: one of each category, all None() matchers",
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
		suite.Run(test.name, func() {
			t := suite.T()

			sel := exchangeBackupCreateSelectors(test.user, test.data)
			assert.Equal(t, test.expectIncludeLen, len(sel.Includes))
		})
	}
}

func (suite *ExchangeUnitSuite) TestExchangeBackupDetailsSelectors() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.ExchangeOptionDetailLookups {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsExchangeCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
				false)
			assert.NoError(t, err, clues.ToCore(err))
			assert.ElementsMatch(t, test.Expected, output.Entries)
		})
	}
}

func (suite *ExchangeUnitSuite) TestExchangeBackupDetailsSelectorsBadFormats() {
	ctx, flush := tester.NewContext()
	defer flush()

	for _, test := range testdata.BadExchangeOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			output, err := runDetailsExchangeCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts,
				false)
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, output)
		})
	}
}
