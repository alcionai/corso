package export

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/cli/flags"
	flagsTD "github.com/alcionai/canario/src/cli/flags/testdata"
	cliTD "github.com/alcionai/canario/src/cli/testdata"
	"github.com/alcionai/canario/src/cli/utils"
	"github.com/alcionai/canario/src/internal/tester"
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
			parent := &cobra.Command{Use: exportCommand}

			cmd := cliTD.SetUpCmdHasFlags(
				t,
				parent,
				addGroupsCommands,
				[]cliTD.UseCobraCommandFn{
					flags.AddAllProviderFlags,
					flags.AddAllStorageFlags,
				},
				flagsTD.WithFlags(
					groupsServiceCommand,
					[]string{
						flagsTD.RestoreDestination,
						"--" + flags.RunModeFN, flags.RunModeFlagTest,
						"--" + flags.BackupFN, flagsTD.BackupInput,
						"--" + flags.FormatFN, flagsTD.FormatType,
						"--" + flags.ArchiveFN,
					},
					flagsTD.PreparedProviderFlags(),
					flagsTD.PreparedStorageFlags()))

			cliTD.CheckCmdChild(
				t,
				parent,
				3,
				test.expectUse,
				test.expectShort,
				test.expectRunE)

			opts := utils.MakeGroupsOpts(cmd)

			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
			assert.Equal(t, flagsTD.Archive, opts.ExportCfg.Archive)
			assert.Equal(t, flagsTD.FormatType, opts.ExportCfg.Format)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
