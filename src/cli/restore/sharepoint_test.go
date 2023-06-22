package restore

import (
	"bytes"
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/credentials"
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
			utils.AddRunModeFlag(cmd, true)

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
				"--" + utils.RunModeFN, utils.RunModeFlagTest,
				"--" + utils.BackupFN, testdata.BackupInput,

				"--" + utils.LibraryFN, testdata.LibraryInput,
				"--" + utils.FileFN, testdata.FlgInputs(testdata.FileNameInput),
				"--" + utils.FolderFN, testdata.FlgInputs(testdata.FolderPathInput),
				"--" + utils.FileCreatedAfterFN, testdata.FileCreatedAfterInput,
				"--" + utils.FileCreatedBeforeFN, testdata.FileCreatedBeforeInput,
				"--" + utils.FileModifiedAfterFN, testdata.FileModifiedAfterInput,
				"--" + utils.FileModifiedBeforeFN, testdata.FileModifiedBeforeInput,

				"--" + utils.ListItemFN, testdata.FlgInputs(testdata.ListItemInput),
				"--" + utils.ListFolderFN, testdata.FlgInputs(testdata.ListFolderInput),

				"--" + utils.PageFN, testdata.FlgInputs(testdata.PageInput),
				"--" + utils.PageFolderFN, testdata.FlgInputs(testdata.PageFolderInput),

				"--" + utils.AWSAccessKeyFN, testdata.AWSAccessKeyID,
				"--" + utils.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
				"--" + utils.AWSSessionTokenFN, testdata.AWSSessionToken,

				"--" + utils.AzureClientIDFN, testdata.AzureClientID,
				"--" + utils.AzureClientTenantFN, testdata.AzureTenantID,
				"--" + utils.AzureClientSecretFN, testdata.AzureClientSecret,

				"--" + credentials.CorsoPassphraseFN, testdata.CorsoPassphrase,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeSharePointOpts(cmd)
			assert.Equal(t, testdata.BackupInput, utils.BackupIDFV)

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

			assert.Equal(t, testdata.AWSAccessKeyID, utils.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, utils.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, utils.AWSSessionTokenFV)

			assert.Equal(t, testdata.AzureClientID, credentials.AzureClientIDFV)
			assert.Equal(t, testdata.AzureTenantID, credentials.AzureClientTenantFV)
			assert.Equal(t, testdata.AzureClientSecret, credentials.AzureClientSecretFV)

			assert.Equal(t, testdata.CorsoPassphrase, credentials.CorsoPassphraseFV)
		})
	}
}
