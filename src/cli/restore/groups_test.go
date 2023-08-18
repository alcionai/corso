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
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/cli/utils/testdata"
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
		{"restore groups", restoreCommand, expectUse, groupsRestoreCmd().Short, restoreGroupsCmd},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{Use: test.use}

			// normally a persistent flag from the root.
			// required to ensure a dry run.
			flags.AddRunModeFlag(cmd, true)

			c := addGroupsCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			tester.AreSameFunc(t, test.expectRunE, child.RunE)

			cmd.SetArgs([]string{
				"groups",
				"--" + flags.RunModeFN, flags.RunModeFlagTest,
				"--" + flags.BackupFN, testdata.BackupInput,

				"--" + flags.CollisionsFN, testdata.Collisions,
				"--" + flags.DestinationFN, testdata.Destination,
				"--" + flags.ToResourceFN, testdata.ToResource,

				"--" + flags.AWSAccessKeyFN, testdata.AWSAccessKeyID,
				"--" + flags.AWSSecretAccessKeyFN, testdata.AWSSecretAccessKey,
				"--" + flags.AWSSessionTokenFN, testdata.AWSSessionToken,

				"--" + flags.AzureClientIDFN, testdata.AzureClientID,
				"--" + flags.AzureClientTenantFN, testdata.AzureTenantID,
				"--" + flags.AzureClientSecretFN, testdata.AzureClientSecret,

				"--" + flags.CorsoPassphraseFN, testdata.CorsoPassphrase,

				// bool flags
				"--" + flags.RestorePermissionsFN,
			})

			cmd.SetOut(new(bytes.Buffer)) // drop output
			cmd.SetErr(new(bytes.Buffer)) // drop output
			err := cmd.Execute()
			// assert.NoError(t, err, clues.ToCore(err))
			assert.ErrorIs(t, err, utils.ErrNotYetImplemented, clues.ToCore(err))

			opts := utils.MakeGroupsOpts(cmd)
			assert.Equal(t, testdata.BackupInput, flags.BackupIDFV)

			assert.Equal(t, testdata.Collisions, opts.RestoreCfg.Collisions)
			assert.Equal(t, testdata.Destination, opts.RestoreCfg.Destination)
			assert.Equal(t, testdata.ToResource, opts.RestoreCfg.ProtectedResource)

			assert.Equal(t, testdata.AWSAccessKeyID, flags.AWSAccessKeyFV)
			assert.Equal(t, testdata.AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
			assert.Equal(t, testdata.AWSSessionToken, flags.AWSSessionTokenFV)

			assert.Equal(t, testdata.AzureClientID, flags.AzureClientIDFV)
			assert.Equal(t, testdata.AzureTenantID, flags.AzureClientTenantFV)
			assert.Equal(t, testdata.AzureClientSecret, flags.AzureClientSecretFV)

			assert.Equal(t, testdata.CorsoPassphrase, flags.CorsoPassphraseFV)
			assert.True(t, flags.RestorePermissionsFV)
		})
	}
}
