package restore

import (
	"bytes"
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
)

type OneDriveUnitSuite struct {
	tester.Suite
}

func TestOneDriveUnitSuite(t *testing.T) {
	suite.Run(t, &OneDriveUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveUnitSuite) TestAddOneDriveCommands() {
	expectUse := oneDriveServiceCommand + " " + oneDriveServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore onedrive", restoreCommand, expectUse, oneDriveRestoreCmd().Short, restoreOneDriveCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persistent flag from the root.
			// required to ensure a dry run.
			flags.AddRunModeFlag(cmd, true)

			c := addOneDriveCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			cmd.SetArgs([]string{
				"onedrive",
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, testdata.BackupInput,
				"--" + flags.FileFN, testdata.FlgInputs(testdata.FileNameInput),
				"--" + flags.FolderFN, testdata.FlgInputs(testdata.FolderPathInput),
				"--" + flags.FileCreatedAfterFN, testdata.FileCreatedAfterInput,
				"--" + flags.FileCreatedBeforeFN, testdata.FileCreatedBeforeInput,
				"--" + flags.FileModifiedAfterFN, testdata.FileModifiedAfterInput,
				"--" + flags.FileModifiedBeforeFN, testdata.FileModifiedBeforeInput,
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

			opts := utils.MakeOneDriveOpts(cmd)
			assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

			assert.ElementsMatch(t, testdata.FileNameInput, opts.FileName)
			assert.ElementsMatch(t, testdata.FolderPathInput, opts.FolderPath)
			assert.Equal(t, testdata.FileCreatedAfterInput, opts.FileCreatedAfter)
			assert.Equal(t, testdata.FileCreatedBeforeInput, opts.FileCreatedBefore)
			assert.Equal(t, testdata.FileModifiedAfterInput, opts.FileModifiedAfter)
			assert.Equal(t, testdata.FileModifiedBeforeInput, opts.FileModifiedBefore)

			assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

			assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
			assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
			assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

			assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)
		})
	}
}
