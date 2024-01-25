package export

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/tester"
)

type TeamsChatsUnitSuite struct {
	tester.Suite
}

func TestTeamsChatsUnitSuite(t *testing.T) {
	suite.Run(t, &TeamsChatsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TeamsChatsUnitSuite) TestAddTeamsChatsCommands() {
	expectUse := teamschatsServiceCommand + " " + teamschatsServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"export teamschats", exportCommand, expectUse, teamschatsExportCmd().Short, exportTeamsChatsCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			parent := &cobra.Command{Use: exportCommand}

			cmd := cliTD.SetUpCmdHasFlags(
				t,
				parent,
				addTeamsChatsCommands,
				[]cliTD.UseCobraCommandFn{
					flags.AddAllProviderFlags,
					flags.AddAllStorageFlags,
				},
				flagsTD.WithFlags(
					teamschatsServiceCommand,
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

			opts := utils.MakeTeamsChatsOpts(cmd)

			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
			assert.Equal(t, flagsTD.Archive, opts.ExportCfg.Archive)
			assert.Equal(t, flagsTD.FormatType, opts.ExportCfg.Format)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
