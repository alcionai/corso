package export

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

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestAddGroupsCommands() {
	expectUse := groupsServiceCommand + " " + groupsServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"export groups", exportCommand, expectUse, groupsExportCmd().Short, exportGroupsCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// persistent flags not added by addCommands
			flags.AddRunModeFlag(cmd, true)

			c := addGroupsCommands(cmd)
			require.NotNil(t, c)

			// non-persistent flags not added by addCommands
			flags.AddAllProviderFlags(c)
			flags.AddAllStorageFlags(c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			flagsTD.WithFlags(
				cmd,
				groupsServiceCommand,
				[]string{
					flagsTD.RestoreDestination,
					"--" + flags.RunModeFN, flags.RunModeFlagTest,
					"--" + flags.BackupFN, flagsTD.BackupInput,

					"--" + flags.FormatFN, flagsTD.FormatType,

					// bool flags
					"--" + flags.ArchiveFN,
				},
				flagsTD.PreparedProviderFlags(),
				flagsTD.PreparedStorageFlags())

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output

			err := cmd.Execute()
			assert.NoError(t, err, clues.ToCore(err))

			opts := utils.MakeGroupsOpts(cmd)
			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)

			assert.Equal(t, flagsTD.Archive, opts.ExportCfg.Archive)
			assert.Equal(t, flagsTD.FormatType, opts.ExportCfg.Format)

			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
