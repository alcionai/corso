package utils

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/credentials"
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
			assert.Equal(t, "tenantID", credentials.AzureClientTenantFV, AzureClientTenantFN)
			assert.Equal(t, "clientID", credentials.AzureClientIDFV, AzureClientIDFN)
			assert.Equal(t, "secret", credentials.AzureClientSecretFV, AzureClientSecretFN)
		},
	}

	AddAzureCredsFlags(cmd)
	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + AzureClientIDFN, "clientID",
		"--" + AzureClientTenantFN, "tenantID",
		"--" + AzureClientSecretFN, "secret",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *FlagUnitSuite) TestAddAWSCredsFlags() {
	t := suite.T()

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.Equal(t, "accesskey", AccessKeyFV, AccessKeyFN)
			assert.Equal(t, "secretkey", SecretAccessKeyFV, SecretAccessKeyFN)
			assert.Equal(t, "token", SessionTokenFV, SessionTokenFN)
		},
	}

	AddAWSCredsFlags(cmd)
	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + AccessKeyFN, "accesskey",
		"--" + SecretAccessKeyFN, "secretkey",
		"--" + SessionTokenFN, "token",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *FlagUnitSuite) TestAddCorsoPassphraseFlags() {
	t := suite.T()

	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			assert.Equal(t, "passphrase", credentials.CorsoPassphraseFV, credentials.CorsoPassphraseFN)
		},
	}

	AddCorsoPassphaseFlags(cmd)
	// Test arg parsing for few args
	cmd.SetArgs([]string{
		"test",
		"--" + credentials.CorsoPassphraseFN, "passphrase",
	})

	err := cmd.Execute()
	require.NoError(t, err, clues.ToCore(err))
}
