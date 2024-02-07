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
		{
			name:        "metadata-files onedrive",
			use:         metadataFilesCommand,
			expectUse:   expectUse,
			expectShort: oneDriveMetadataFilesCmd().Short,
			expectRunE:  metadataFilesOneDriveCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			parent := &cobra.Command{Use: metadataFilesCommand}

			cmd := cliTD.SetUpCmdHasFlags(
				t,
				parent,
				addOneDriveCommands,
				[]cliTD.UseCobraCommandFn{
					flags.AddAllProviderFlags,
					flags.AddAllStorageFlags,
				},
				flagsTD.WithFlags(
					oneDriveServiceCommand,
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
