package debug

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/cli/flags"
	flagsTD "github.com/alcionai/canario/src/cli/flags/testdata"
	cliTD "github.com/alcionai/canario/src/cli/testdata"
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
		{
			name:        "metdata-files groups",
			use:         metadataFilesCommand,
			expectUse:   expectUse,
			expectShort: groupsMetadataFilesCmd().Short,
			expectRunE:  metadataFilesGroupsCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			parent := &cobra.Command{Use: metadataFilesCommand}

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
						"--" + flags.RunModeFN, flags.RunModeFlagTest,
						"--" + flags.BackupFN, flagsTD.BackupInput,
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

			assert.Equal(t, flagsTD.BackupInput, flags.BackupIDFV)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
