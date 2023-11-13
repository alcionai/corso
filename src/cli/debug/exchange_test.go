package debug

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	flagsTD "github.com/alcionai/corso/src/cli/flags/testdata"
	cliTD "github.com/alcionai/corso/src/cli/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type ExchangeUnitSuite struct {
	tester.Suite
}

func TestExchangeUnitSuite(t *testing.T) {
	suite.Run(t, &ExchangeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExchangeUnitSuite) TestExchangeCommands() {
	expectUse := exchangeServiceCommand + " " + exchangeServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{
			name:        "metdata-files exchange",
			use:         metadataFilesCommand,
			expectUse:   expectUse,
			expectShort: exchangeMetadataFilesCmd().Short,
			expectRunE:  metadataFilesExchangeCmd,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			parent := &cobra.Command{Use: metadataFilesCommand}

			cmd := cliTD.SetUpCmdHasFlags(
				t,
				parent,
				addExchangeCommands,
				[]cliTD.UseCobraCommandFn{
					flags.AddAllProviderFlags,
					flags.AddAllStorageFlags,
				},
				flagsTD.WithFlags(
					exchangeServiceCommand,
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
			flagsTD.AssertProviderFlags(t, cmd)
			flagsTD.AssertStorageFlags(t, cmd)
		})
	}
}
