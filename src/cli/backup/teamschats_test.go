package backup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/cli/flags"
	flagsTD "github.com/alcionai/canario/src/cli/flags/testdata"
	cliTD "github.com/alcionai/canario/src/cli/testdata"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/control"
)

type TeamsChatsUnitSuite struct {
	tester.Suite
}

func TestTeamsChatsUnitSuite(t *testing.T) {
	suite.Run(t, &TeamsChatsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TeamsChatsUnitSuite) TestAddTeamsChatsCommands() {
	expectUse := teamschatsServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			name:        "create teamschats",
			use:         createCommand,
			expectUse:   expectUse + " " + teamschatsServiceCommandCreateUseSuffix,
			expectShort: teamschatsCreateCmd().Short,
			expectRunE:  createTeamsChatsCmd,
		},
		{
			name:        "list teamschats",
			use:         listCommand,
			expectUse:   expectUse,
			expectShort: teamschatsListCmd().Short,
			expectRunE:  listTeamsChatsCmd,
		},
		{
			name:        "details teamschats",
			use:         detailsCommand,
			expectUse:   expectUse + " " + teamschatsServiceCommandDetailsUseSuffix,
			expectShort: teamschatsDetailsCmd().Short,
			expectRunE:  detailsTeamsChatsCmd,
		},
		{
			name:        "delete teamschats",
			use:         deleteCommand,
			expectUse:   expectUse + " " + teamschatsServiceCommandDeleteUseSuffix,
			expectShort: teamschatsDeleteCmd().Short,
			expectRunE:  deleteTeamsChatsCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addTeamsChatsCommands(cmd)
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

func (suite *TeamsChatsUnitSuite) TestValidateTeamsChatsBackupCreateFlags() {
	table := []struct {
		name   string
		cats   []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "none",
			cats:   []string{},
			expect: assert.NoError,
		},
		{
			name:   "chats",
			cats:   []string{flags.DataChats},
			expect: assert.NoError,
		},
		{
			name: "all allowed",
			cats: []string{
				flags.DataChats,
			},
			expect: assert.NoError,
		},
		{
			name:   "bad inputs",
			cats:   []string{"foo"},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := validateTeamsChatsBackupCreateFlags([]string{"*"}, test.cats)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *TeamsChatsUnitSuite) TestBackupCreateFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: createCommand},
		addTeamsChatsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			teamschatsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.UserFN, flagsTD.FlgInputs(flagsTD.UsersInput),
				"--" + flags.CategoryDataFN, flagsTD.FlgInputs(flagsTD.TeamsChatsCategoryDataInput),
			},
			flagsTD.PreparedGenericBackupFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	opts := utils.MakeTeamsChatsOpts(cmd)
	co := utils.Control()
	backupOpts := utils.ParseBackupOptions()

	// TODO(ashmrtn): Remove flag checks on control.Options to control.Backup once
	// restore flags are switched over too and we no longer parse flags beyond
	// connection info into control.Options.
	assert.Equal(t, control.FailFast, backupOpts.FailureHandling)
	assert.True(t, backupOpts.Incrementals.ForceFullEnumeration)
	assert.True(t, backupOpts.Incrementals.ForceItemDataRefresh)

	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)

	assert.ElementsMatch(t, flagsTD.UsersInput, opts.Users)
	flagsTD.AssertGenericBackupFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *TeamsChatsUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: listCommand},
		addTeamsChatsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			teamschatsServiceCommand,
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

func (suite *TeamsChatsUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: detailsCommand},
		addTeamsChatsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			teamschatsServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
				"--" + flags.SkipReduceFN,
			},
			flagsTD.PreparedTeamsChatsFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	co := utils.Control()

	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	assert.True(t, co.SkipReduce)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
	flagsTD.AssertTeamsChatsFlags(t, cmd)
}

func (suite *TeamsChatsUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: deleteCommand},
		addTeamsChatsCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			teamschatsServiceCommand,
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
