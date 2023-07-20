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

type SharePointUnitSuite struct {
	tester.Suite
}

func TestSharePointUnitSuite(t *testing.T) {
	suite.Run(t, &SharePointUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharePointUnitSuite) TestAddSharePointCommands() {
	expectUse := sharePointServiceCommand + " " + sharePointServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore onedrive", restoreCommand, expectUse, sharePointRestoreCmd().Short, restoreSharePointCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persistent flag from the root.
			// required to ensure a dry run.
			flags.AddRunModeFlag(cmd, true)

			c := addSharePointCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			cmd.SetArgs([]string{
				"sharepoint",
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, testdata.BackupInput,
				"--" + flags.LibraryFN, testdata.LibraryInput,
				"--" + flags.FileFN, testdata.FlgInputs(testdata.FileNameInput),
				"--" + flags.FolderFN, testdata.FlgInputs(testdata.FolderPathInput),
				"--" + flags.FileCreatedAfterFN, testdata.FileCreatedAfterInput,
				"--" + flags.FileCreatedBeforeFN, testdata.FileCreatedBeforeInput,
				"--" + flags.FileModifiedAfterFN, testdata.FileModifiedAfterInput,
				"--" + flags.FileModifiedBeforeFN, testdata.FileModifiedBeforeInput,
				"--" + flags.ListItemFN, testdata.FlgInputs(testdata.ListItemInput),
				"--" + flags.ListFolderFN, testdata.FlgInputs(testdata.ListFolderInput),
				"--" + flags.PageFN, testdata.FlgInputs(testdata.PageInput),
				"--" + flags.PageFolderFN, testdata.FlgInputs(testdata.PageFolderInput),

				"--" + flags.CollisionsFN, testdata.Collisions,
				"--" + flags.DestinationFN, testdata.Destination,
				"--" + flags.ToResourceFN, testdata.ToResource,

				"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
				"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
				"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

				"--" + flags.AzureClientIDFN, testdata.AzureClientID,
				"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
				"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

				"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

				// bool flags
				"--" + flags.RestorePermissionsFN,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeSharePointOpts(cmd)
			assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

			assert.Equal(t, testdata.LibraryInput, opts.Library)
			assert.ElementsMatch(t, testdata.FileNameInput, opts.FileName)
			assert.ElementsMatch(t, testdata.FolderPathInput, opts.FolderPath)
			assert.Equal(t, testdata.FileCreatedAfterInput, opts.FileCreatedAfter)
			assert.Equal(t, testdata.FileCreatedBeforeInput, opts.FileCreatedBefore)
			assert.Equal(t, testdata.FileModifiedAfterInput, opts.FileModifiedAfter)
			assert.Equal(t, testdata.FileModifiedBeforeInput, opts.FileModifiedBefore)

			assert.ElementsMatch(t, testdata.ListItemInput, opts.ListItem)
			assert.ElementsMatch(t, testdata.ListFolderInput, opts.ListFolder)

			assert.ElementsMatch(t, testdata.PageInput, opts.Page)
			assert.ElementsMatch(t, testdata.PageFolderInput, opts.PageFolder)

			assert.Equal(t, testdata.Collisions, opts.RestoreCfg.Collisions)
			assert.Equal(t, testdata.Destination, opts.RestoreCfg.Destination)
			assert.Equal(t, testdata.ToResource, opts.RestoreCfg.ProtectedResource)

			assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

			assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
			assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
			assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

			assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)

			// bool flags
			assert.True(t, flags.RestorePermissionsFV)
		})
	}
}
