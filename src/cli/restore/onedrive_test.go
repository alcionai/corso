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

			// global flags not added by addCommands
			flags.AddRunModeFlag(cmd, true)
			flags.AddAllProviderFlags(cmd)
			flags.AddAllStorageFlags(cmd)

			c := addOneDriveCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			flagsTD.WithFlags(
				cmd,
				oneDriveServiceCommand,
				[]string{
					"--" + flags.RunModeFN, flags.RunModeFlagTest,
					"--" + flags.BackupFN, flagsTD.BackupInput,
					"--" + flags.FileFN, flagsTD.FlgInputs(flagsTD.FileNameInput),
					"--" + flags.FolderFN, flagsTD.FlgInputs(flagsTD.FolderPathInput),
					"--" + flags.FileCreatedAfterFN, flagsTD.FileCreatedAfterInput,
					"--" + flags.FileCreatedBeforeFN, flagsTD.FileCreatedBeforeInput,
					"--" + flags.FileModifiedAfterFN, flagsTD.FileModifiedAfterInput,
					"--" + flags.FileModifiedBeforeFN, flagsTD.FileModifiedBeforeInput,

					"--" + flags.CollisionsFN, flagsTD.Collisions,
					"--" + flags.DestinationFN, flagsTD.Destination,
					"--" + flags.ToResourceFN, flagsTD.ToResource,

					// bool flags
					"--" + flags.NoPermissionsFN,
				},
				flagsTD.PreparedProviderFlags(),
				flagsTD.PreparedStorageFlags())

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output

			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeOneDriveOpts(cmd)
			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

			assert.ElementsMatch(t, flagsTD.FileNameInput, opts.FileName)
			assert.ElementsMatch(t, flagsTD.FolderPathInput, opts.FolderPath)
			assert.Equal(t, flagsTD.FileCreatedAfterInput, opts.FileCreatedAfter)
			assert.Equal(t, flagsTD.FileCreatedBeforeInput, opts.FileCreatedBefore)
			assert.Equal(t, flagsTD.FileModifiedAfterInput, opts.FileModifiedAfter)
			assert.Equal(t, flagsTD.FileModifiedBeforeInput, opts.FileModifiedBefore)

			assert.Equal(t, flagsTD.Collisions, opts.RestoreCfg.Collisions)
			assert.Equal(t, flagsTD.Destination, opts.RestoreCfg.Destination)
			assert.Equal(t, flagsTD.ToResource, opts.RestoreCfg.ProtectedResource)

			// bool flags
			assert.True(t, flags.NoPermissionsFV)

			flagsTD.AssertProviderFlags(t, cmd)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
