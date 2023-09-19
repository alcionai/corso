package backup

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	dtd "github.com/alcionai/corso/src/pkg/backup/details/testdata"
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

	cmd := &cobra.Command{Use: createCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addOneDriveCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		oneDriveServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,

		"--" + flags.UserFN, testdata.FlgInputs(testdata.UsersInput),

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		// bool flags
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.ForceItemDataDownloadFN,
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	opts := utils.MakeOneDriveOpts(cmd)
	co := utils.Control()

	assert.ElementsMatch(t, testdata.UsersInput, opts.Users)
	// no assertion for category data input

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	// bool flags
	assert.Equal(t, control.FailFast, co.FailureHandling)
	assert.True(t, co.ToggleFeatures.DisableIncrementals)
	assert.True(t, co.ToggleFeatures.ForceItemDataDownload)
}

func (suite *OneDriveUnitSuite) TestBackupListFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: listCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addOneDriveCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		oneDriveServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, testdata.BackupInput,

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		// bool flags
		"--" + flags.FailedItemsFN, "show",
		"--" + flags.SkippedItemsFN, "show",
		"--" + flags.RecoveredErrorsFN, "show",
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	assert.Equal(t, flags.ListFailedItemsFV, "show")
	assert.Equal(t, flags.ListSkippedItemsFV, "show")
	assert.Equal(t, flags.ListRecoveredErrorsFV, "show")
}

func (suite *OneDriveUnitSuite) TestBackupDetailsFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: detailsCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addOneDriveCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		oneDriveServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, testdata.BackupInput,

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

		// bool flags
		"--" + flags.SkipReduceFN,
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	co := utils.Control()

	assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

	assert.True(t, co.SkipReduce)
}

func (suite *OneDriveUnitSuite) TestBackupDeleteFlags() {
	t := suite.T()

	cmd := &cobra.Command{Use: deleteCommand}

	// normally a persistent flag from the root.
	// required to ensure a dry run.
	flags.AddRunModeFlag(cmd, true)

	c := addOneDriveCommands(cmd)
	require.NotNil(t, c)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		oneDriveServiceCommand,
		"--" + flags.RunModeFN, flags.RunModeFlagTest,
		"--" + flags.BackupFN, testdata.BackupInput,

		"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

		"--" + flags.AzureClientIDFN, testdata.AzureClientID,
		"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
		"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

		"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,
	})

	cmd.SetOut(new(bytes.Buffer)) // drop output
	cmd.SetErr(new(bytes.Buffer)) // drop output
	err := cmd.Execute()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

	assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

	assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)
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

func (suite *OneDriveUnitSuite) TestOneDriveBackupDetailsSelectors() {
	for v := 0; v <= version.Backup; v++ {
		suite.Run(fmt.Sprintf("version%d", v), func() {
			for _, test := range testdata.OneDriveOptionDetailLookups {
				suite.Run(test.Name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					bg := testdata.VersionedBackupGetter{
						Details: dtd.GetDetailsSetForVersion(t, v),
					}

					output, err := runDetailsOneDriveCmd(
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

func (suite *OneDriveUnitSuite) TestOneDriveBackupDetailsSelectorsBadFormats() {
	for _, test := range testdata.BadOneDriveOptionsFormats {
		suite.Run(test.Name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			output, err := runDetailsOneDriveCmd(
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
