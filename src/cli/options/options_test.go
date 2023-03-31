package options

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type OptionsUnitSuite struct {
	tester.Suite
}

func TestOptionsUnitSuite(t *testing.T) {
	suite.Run(t, &OptionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OptionsUnitSuite) TestAddExchangeCommands() {
	t := suite.T()

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.True(t, failFastFV, failFastFN)
			assert.True(t, disableIncrementalsFV, disableIncrementalsFN)
			assert.True(t, enablePermissionsBackupFV, enablePermissionsBackupFN)
			assert.True(t, noStatsFV, noStatsFN)
			assert.True(t, restorePermissionsFV, restorePermissionsFN)
			assert.True(t, skipReduceFV, skipReduceFN)
			assert.Equal(t, 2, fetchParallelismFV, fetchParallelismFN)
		},
	}

	// adds no-stats
	AddGlobalOperationFlags(cmd)

	AddFailFastFlag(cmd)
	AddDisableIncrementalsFlag(cmd)
	AddEnablePermissionsBackupFlag(cmd)
	AddRestorePermissionsFlag(cmd)
	AddSkipReduceFlag(cmd)

	AddFetchParallelismFlag(cmd)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + failFastFN,
		"--" + disableIncrementalsFN,
		"--" + enablePermissionsBackupFN,
		"--" + noStatsFN,
		"--" + restorePermissionsFN,
		"--" + skipReduceFN,

		"--" + fetchParallelismFN, "2",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}
