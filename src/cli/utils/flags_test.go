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

type FlagUnitSuite struct {
	tester.Suite
}

func TestFlagUnitSuite(t *testing.T) {
	suite.Run(t, &FlagUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *FlagUnitSuite) TestAddAzureCredsFlags() {
	t := suite.T()

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.Equal(t, "tenantID", flags.AzureClientTenantFV, flags.AzureClientTenantFN)
			assert.Equal(t, "clientID", flags.AzureClientIDFV, flags.AzureClientIDFN)
			assert.Equal(t, "secret", flags.AzureClientSecretFV, flags.AzureClientSecretFN)
		},
	}

	flags.AddAzureCredsFlags(cmd)
	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + flags.AzureClientIDFN, "clientID",
		"--" + flags.AzureClientTenantFN, "tenantID",
		"--" + flags.AzureClientSecretFN, "secret",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *FlagUnitSuite) TestAddAWSCredsFlags() {
	t := suite.T()

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.Equal(t, "accesskey", flags.AWSAccessKeyFV, flags.AWSAccessKeyFN)
			assert.Equal(t, "secretkey", flags.AWSSecretAccessKeyFV, flags.AWSSecretAccessKeyFN)
			assert.Equal(t, "token", flags.AWSSessionTokenFV, flags.AWSSessionTokenFN)
		},
	}

	flags.AddAWSCredsFlags(cmd)
	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + flags.AWSAccessKeyFN, "accesskey",
		"--" + flags.AWSSecretAccessKeyFN, "secretkey",
		"--" + flags.AWSSessionTokenFN, "token",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *FlagUnitSuite) TestAddCorsoPassphraseFlags() {
	t := suite.T()

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.Equal(t, "passphrase", flags.CorsoPassphraseFV, flags.CorsoPassphraseFN)
		},
	}

	flags.AddCorsoPassphaseFlags(cmd)
	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + flags.CorsoPassphraseFN, "passphrase",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}
