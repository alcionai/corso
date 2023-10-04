package backup

import (
	"strconv"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/tester"
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

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: createCommand},
		addExchangeCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			exchangeServiceCommand,
			[]string{
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
			},
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	opts := utils.MakeExchangeOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, flagsTD.MailboxInput, opts.Users)
	assert.Equal(t, flagsTD.FetchParallelism, strconv.Itoa(co.Parallelism.ItemFetch))
	assert.Equal(t, flagsTD.DeltaPageSize, strconv.Itoa(int(co.DeltaPageSize)))
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

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: listCommand},
		addExchangeCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			exchangeServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
			},
			flagsTD.PreparedBackupListFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	flagsTD.AssertBackupListFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *ExchangeUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: detailsCommand},
		addExchangeCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			exchangeServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
				"--" + flags.SkipReduceFN,
			},
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	co := utils.Control()

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	assert.True(t, co.SkipReduce)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *ExchangeUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: deleteCommand},
		addExchangeCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			exchangeServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
			},
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

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
