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
			assert.True(t, failFastFV, FailFastFN)
			assert.True(t, disableIncrementalsFV, DisableIncrementalsFN)
			assert.True(t, noStatsFV, NoStatsFN)
			assert.True(t, restorePermissionsFV, RestorePermissionsFN)
			assert.True(t, skipReduceFV, SkipReduceFN)
			assert.Equal(t, 2, fetchParallelismFV, FetchParallelismFN)
		},
	}

	// adds no-stats
	AddGlobalOperationFlags(cmd)

	AddFailFastFlag(cmd)
	AddDisableIncrementalsFlag(cmd)
	AddRestorePermissionsFlag(cmd)
	AddSkipReduceFlag(cmd)

	AddFetchParallelismFlag(cmd)

	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + FailFastFN,
		"--" + DisableIncrementalsFN,
		"--" + NoStatsFN,
		"--" + RestorePermissionsFN,
		"--" + SkipReduceFN,

		"--" + FetchParallelismFN, "2",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *OptionsUnitSuite) TestDisableConcurrencyLimiterFlag() {
	tests := []struct {
		name     string
		args     []string
		assertFn func(*testing.T, bool, string)
		addFlag  bool
	}{
		{
			name: "--disable-concurrency-limiter not set",
			args: []string{"test"},
			assertFn: func(t *testing.T, f bool, s string) {
				assert.False(t, f, s)
			},
			addFlag: false,
		},
		{
			name: "--disable-concurrency-limiter set",
			args: []string{"test", "--disable-concurrency-limiter"},
			assertFn: func(t *testing.T, f bool, s string) {
				assert.True(t, f, s)
			},
			addFlag: true,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{
				Use: "test",
				Run: func(cmd *cobra.Command, args []string) {
					test.assertFn(t, disableConcurrencyLimiterFV, DisableConcurrencyLimiterFN)
				},
			}

			if test.addFlag {
				AddDisableConcurrencyLimiterFlag(cmd)
			}

			cmd.SetArgs(test.args)
			err := cmd.Execute()
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}
