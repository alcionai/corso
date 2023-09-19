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
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	"github.com/alcionai/corso/src/cli/utils"
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
		{"restore sharepoint", restoreCommand, expectUse, sharePointRestoreCmd().Short, restoreSharePointCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// global flags not added by addCommands
			flags.AddRunModeFlag(cmd, true)
			flags.AddAllProviderFlags(cmd)
			flags.AddAllStorageFlags(cmd)

			c := addSharePointCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			flagsTD.WithFlags(
				cmd,
				sharePointServiceCommand,
				[]string{
					"--" + flags.RunModeFN, flags.RunModeFlagTest,
					"--" + flags.BackupFN, flagsTD.BackupInput,
					"--" + flags.LibraryFN, flagsTD.LibraryInput,
					"--" + flags.FileFN, flagsTD.FlgInputs(flagsTD.FileNameInput),
					"--" + flags.FolderFN, flagsTD.FlgInputs(flagsTD.FolderPathInput),
					"--" + flags.FileCreatedAfterFN, flagsTD.FileCreatedAfterInput,
					"--" + flags.FileCreatedBeforeFN, flagsTD.FileCreatedBeforeInput,
					"--" + flags.FileModifiedAfterFN, flagsTD.FileModifiedAfterInput,
					"--" + flags.FileModifiedBeforeFN, flagsTD.FileModifiedBeforeInput,
					"--" + flags.ListItemFN, flagsTD.FlgInputs(flagsTD.ListItemInput),
					"--" + flags.ListFolderFN, flagsTD.FlgInputs(flagsTD.ListFolderInput),
					"--" + flags.PageFN, flagsTD.FlgInputs(flagsTD.PageInput),
					"--" + flags.PageFolderFN, flagsTD.FlgInputs(flagsTD.PageFolderInput),

					"--" + flags.CollisionsFN, flagsTD.Collisions,
					"--" + flags.DestinationFN, flagsTD.Destination,
					"--" + flags.ToResourceFN, flagsTD.ToResource,

					// bool flags
					"--" + flags.RestorePermissionsFN,
				},
				flagsTD.PreparedProviderFlags(),
				flagsTD.PreparedStorageFlags())

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output

			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeSharePointOpts(cmd)
			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

			assert.Equal(t, flagsTD.LibraryInput, opts.Library)
			assert.ElementsMatch(t, flagsTD.FileNameInput, opts.FileName)
			assert.ElementsMatch(t, flagsTD.FolderPathInput, opts.FolderPath)
			assert.Equal(t, flagsTD.FileCreatedAfterInput, opts.FileCreatedAfter)
			assert.Equal(t, flagsTD.FileCreatedBeforeInput, opts.FileCreatedBefore)
			assert.Equal(t, flagsTD.FileModifiedAfterInput, opts.FileModifiedAfter)
			assert.Equal(t, flagsTD.FileModifiedBeforeInput, opts.FileModifiedBefore)

			assert.ElementsMatch(t, flagsTD.ListItemInput, opts.ListItem)
			assert.ElementsMatch(t, flagsTD.ListFolderInput, opts.ListFolder)

			assert.ElementsMatch(t, flagsTD.PageInput, opts.Page)
			assert.ElementsMatch(t, flagsTD.PageFolderInput, opts.PageFolder)

			assert.Equal(t, flagsTD.Collisions, opts.RestoreCfg.Collisions)
			assert.Equal(t, flagsTD.Destination, opts.RestoreCfg.Destination)
			assert.Equal(t, flagsTD.ToResource, opts.RestoreCfg.ProtectedResource)

			// bool flags
			assert.True(t, flags.RestorePermissionsFV)

			flagsTD.AssertProviderFlags(t, cmd)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
