package utils

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
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
			assert.True(t, flags.FailFastFV, flags.FailFastFN)
			assert.True(t, flags.DisableIncrementalsFV, flags.DisableIncrementalsFN)
			assert.True(t, flags.DisableDeltaFV, flags.DisableDeltaFN)
			assert.True(t, flags.NoStatsFV, flags.NoStatsFN)
			assert.True(t, flags.RestorePermissionsFV, flags.RestorePermissionsFN)
			assert.True(t, flags.SkipReduceFV, flags.SkipReduceFN)
			assert.Equal(t, 2, flags.FetchParallelismFV, flags.FetchParallelismFN)
			assert.True(t, flags.DisableConcurrencyLimiterFV, flags.DisableConcurrencyLimiterFN)
			assert.Equal(t, 499, flags.DeltaPageSizeFV, flags.DeltaPageSizeFN)
		},
	}

	// adds no-stats
	flags.AddGlobalOperationFlags(cmd)

	flags.AddFailFastFlag(cmd)
	flags.AddDisableIncrementalsFlag(cmd)
	flags.AddDisableDeltaFlag(cmd)
	flags.AddRestorePermissionsFlag(cmd)
	flags.AddSkipReduceFlag(cmd)
	flags.AddFetchParallelismFlag(cmd)
	flags.AddDisableConcurrencyLimiterFlag(cmd)
	flags.AddDeltaPageSizeFlag(cmd)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.DisableDeltaFN,
		"--" + flags.NoStatsFN,
		"--" + flags.RestorePermissionsFN,
		"--" + flags.SkipReduceFN,
		"--" + flags.FetchParallelismFN, "2",
		"--" + flags.DisableConcurrencyLimiterFN,
		"--" + flags.DeltaPageSizeFN, "499",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}
