package backup

import (
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

type OneDriveUnitSuite struct {
	tester.Suite
}

func TestOneDriveUnitSuite(t *testing.T) {
	suite.Run(t, &OneDriveUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveUnitSuite) TestAddOneDriveCommands() {
	expectUse := oneDriveServiceCommand

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			name:        "create onedrive",
			use:         createCommand,
			expectUse:   expectUse + " " + oneDriveServiceCommandCreateUseSuffix,
			expectShort: oneDriveCreateCmd().Short,
			expectRunE:  createOneDriveCmd,
		},
		{
			name:        "list onedrive",
			use:         listCommand,
			expectUse:   expectUse,
			expectShort: oneDriveListCmd().Short,
			expectRunE:  listOneDriveCmd,
		},
		{
			name:        "details onedrive",
			use:         detailsCommand,
			expectUse:   expectUse + " " + oneDriveServiceCommandDetailsUseSuffix,
			expectShort: oneDriveDetailsCmd().Short,
			expectRunE:  detailsOneDriveCmd,
		},
		{
			name:        "delete onedrive",
			use:         deleteCommand,
			expectUse:   expectUse + " " + oneDriveServiceCommandDeleteUseSuffix,
			expectShort: oneDriveDeleteCmd().Short,
			expectRunE:  deleteOneDriveCmd,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			c := addOneDriveCommands(cmd)
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

func (suite *OneDriveUnitSuite) TestBackupCreateFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: createCommand},
		addOneDriveCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			oneDriveServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.UserFN, flagsTD.FlgInputs(flagsTD.UsersInput),
			},
			flagsTD.PreparedGenericBackupFlags(),
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	opts := utils.MakeOneDriveOpts(cmd)
	co := utils.Control()
	backupOpts := utils.BackupOptions()

	// TODO(ashmrtn): Remove flag checks on control.Options to control.Backup once
	// restore flags are switched over too and we no longer parse flags beyond
	// connection info into control.Options.
	assert.Equal(t, control.FailFast, backupOpts.FailureHandling)
	assert.True(t, backupOpts.ToggleFeatures.DisableIncrementals)
	assert.True(t, backupOpts.ToggleFeatures.ForceItemDataDownload)

	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)

	assert.ElementsMatch(t, flagsTD.UsersInput, opts.Users)
	flagsTD.AssertGenericBackupFlags(t, cmd)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *OneDriveUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: listCommand},
		addOneDriveCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			oneDriveServiceCommand,
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

func (suite *OneDriveUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: detailsCommand},
		addOneDriveCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			oneDriveServiceCommand,
			[]string{
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, flagsTD.BackupInput,
				"--" + flags.SkipReduceFN,
			},
			flagsTD.PreparedProviderFlags(),
			flagsTD.PreparedStorageFlags()))

	co := utils.Control()

	assert.True(t, co.SkipReduce)
	assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
	flagsTD.AssertProviderFlags(t, cmd)
	flagsTD.AssertStorageFlags(t, cmd)
}

func (suite *OneDriveUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := cliTD.SetUpCmdHasFlags(
		t,
		&cobra.Command{Use: deleteCommand},
		addOneDriveCommands,
		[]cliTD.UseCobraCommandFn{
			flags.AddAllProviderFlags,
			flags.AddAllStorageFlags,
		},
		flagsTD.WithFlags(
			oneDriveServiceCommand,
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

func (suite *OneDriveUnitSuite) TestValidateOneDriveBackupCreateFlags() {
	table := []struct {
		name   string
		user   []string
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "no users",
			expect: assert.Error,
		},
		{
			name:   "users",
			user:   []string{"fnord"},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := validateOneDriveBackupCreateFlags(test.user)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}
