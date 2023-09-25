package backup

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	"github.com/alcionai/corso/src/cli/utils"
	utilsTD "github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	dtd "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
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
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			name:        "create exchange",
			use:         createCommand,
			expectUse:   expectUse + " " + exchangeServiceCommandCreateUseSuffix,
			expectShort: exchangeCreateCmd().Short,
			expectRunE:  createExchangeCmd,
		},
		{
			name:        "list exchange",
			use:         listCommand,
			expectUse:   expectUse,
			expectShort: exchangeListCmd().Short,
			expectRunE:  listExchangeCmd,
		},
		{
			name:        "details exchange",
			use:         detailsCommand,
			expectUse:   expectUse + " " + exchangeServiceCommandDetailsUseSuffix,
			expectShort: exchangeDetailsCmd().Short,
			expectRunE:  detailsExchangeCmd,
		},
		{
			name:        "delete exchange",
			use:         deleteCommand,
			expectUse:   expectUse + " " + exchangeServiceCommandDeleteUseSuffix,
			expectShort: exchangeDeleteCmd().Short,
			expectRunE:  deleteExchangeCmd,
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

func (suite *ExchangeUnitSuite) TestBackupCreateFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: createCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addExchangeCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		exchangeServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	// Test arg parsing for few args
	args := []string{
		exchangeServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,

		"--" + flags.MailBoxFN, flagsTD.FlgInputs(flagsTD.MailboxInput),
		"--" + flags.CategoryDataFN, flagsTD.FlgInputs(flagsTD.ExchangeCategoryDataInput),

		"--" + flags.FetchParallelismFN, flagsTD.FetchParallelism,
		"--" + flags.DeltaPageSizeFN, flagsTD.DeltaPageSize,

		// bool flags
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.ForceItemDataDownloadFN,
		"--" + flags.DisableDeltaFN,
		"--" + flags.EnableImmutableIDFN,
		"--" + flags.DisableConcurrencyLimiterFN,
	}

	args = append(args, flagsTD.PreparedProviderFlags()...)
	args = append(args, flagsTD.PreparedStorageFlags()...)

	cmd.SetArgs(args)

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	opts := utils.MakeExchangeOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, flagsTD.MailboxInput, opts.Users)
	// no assertion for category data input

	assert.Equal(t, flagsTD.FetchParallelism, strconv.Itoa(co.Parallelism.ItemFetch))
	assert.Equal(t, flagsTD.DeltaPageSize, strconv.Itoa(int(co.DeltaPageSize)))

	// bool flags
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)
	assert.True(t, co.ToggleFeatures.DisableDelta)
	assert.True(t, co.ToggleFeatures.ExchangeImmutableIDs)
	assert.True(t, co.ToggleFeatures.DisableConcurrencyLimiter)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *ExchangeUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: listCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addExchangeCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		exchangeServiceCommand, []string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedBackupListFlags(),
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	flagsTD.AssertBackupListFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *ExchangeUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: detailsCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addExchangeCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		exchangeServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
			"--" + flags.SkipReduceFN,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	co := utils.Control()

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	assert.True(t, co.SkipReduce)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *ExchangeUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: deleteCommand}

	// global flags not added by addCommands
	flags.AddRunModeFlag(cmd, true)
	flags.AddAllProviderFlags(cmd)
	flags.AddAllStorageFlags(cmd)

	c := addExchangeCommands(cmd)
	require.NotNil(t, c)

	flagsTD.WithFlags(
		cmd,
		exchangeServiceCommand,
		[]string{
			"--" + flags.RunModeFN, flags.RunModeFlagTest,
			"--" + flags.BackupFN, flagsTD.BackupInput,
		},
		flagsTD.PreparedProviderFlags(),
		flagsTD.PreparedStorageFlags())

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output

	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
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
			user:             []string{flags.Wildcard},
			expectIncludeLen: 3,
		},
		{
			name:             "single user, no data",
			user:             []string{"u1"},
			expectIncludeLen: 3,
		},
		{
			name:             "any users, contacts",
			user:             []string{flags.Wildcard},
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
			user:             []string{flags.Wildcard},
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
			user:             []string{flags.Wildcard},
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
			user:             []string{flags.Wildcard},
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
			user:             []string{flags.Wildcard},
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
			user:             []string{flags.Wildcard},
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
	for v := 0; v <= version.Backup; v++ {
		suite.Run(fmt.Sprintf("version%d", v), func() {
			for _, test := range utilsTD.ExchangeOptionDetailLookups {
				suite.Run(test.Name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					bg := utilsTD.VersionedBackupGetter{
						Details: dtd.GetDetailsSetForVersion(t, v),
					}

					output, err := runDetailsExchangeCmd(
						ctx,
						bg,
						"backup-ID",
						test.Opts(t, v),
						false)
					assert.NoError(t, err, clues.ToCore(err))
					assert.ElementsMatch(t, test.Expected(t, v), output.Entries)
				})
			}
		})
	}
}

func (suite *ExchangeUnitSuite) TestExchangeBackupDetailsSelectorsBadFormats() {
	for _, test := range utilsTD.BadExchangeOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			output, err := runDetailsExchangeCmd(
				ctx,
				test.BackupGetter,
				"backup-ID",
				test.Opts(t, version.Backup),
				false)
			assert.Error(t, err, clues.ToCore(err))
			assert.Empty(t, output)
		})
	}
}
