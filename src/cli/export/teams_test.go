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
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type TeamsUnitSuite struct {
	tester.Suite
}

func TestTeamsUnitSuite(t *testing.T) {
	suite.Run(t, &TeamsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TeamsUnitSuite) TestAddTeamsCommands() {
	expectUse := teamsServiceCommand + " " + teamsServiceCommandUseSuffix

	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"export teams", exportCommand, expectUse, teamsExportCmd().Short, exportTeamsCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persistent flag from the root.
			// required to ensure a dry run.
			flags.AddRunModeFlag(cmd, true)

			c := addTeamsCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			cmd.SetArgs([]string{
				"teams",
				testdata.RestoreDestination,
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, testdata.BackupInput,

				"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
				"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
				"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

				"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

				// bool flags
				"--" + flags.ArchiveFN,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			// assert.NoError(t, err, clues.ToCore(err))
			assert.ErrorIs(t, err, utils.ErrNotYetImplemented, clues.ToCore(err))

			opts := utils.MakeTeamsOpts(cmd)
			assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

			assert.Equal(t, testdata.Archive, opts.ExportCfg.Archive)

			assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

			assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)
		})
	}
}
