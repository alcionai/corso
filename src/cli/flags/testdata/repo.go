package testdata

import (
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/v3/assert"

	"github.com/alcionai/corso/src/cli/flags"
)

func PreparedStorageFlags() []string {
	return []string{
		"--" + flags.AWSAccessKeyFN, AWSAccessKeyID,
		"--" + flags.AWSSecretAccessKeyFN, AWSSecretAccessKey,
		"--" + flags.AWSSessionTokenFN, AWSSessionToken,

		"--" + flags.PassphraseFN, CorsoPassphrase,
	}
}

func AssertStorageFlags(t *testing.T, cmd *cobra.Command) {
	assert.Equal(t, AWSAccessKeyID, flags.AWSAccessKeyFV)
	assert.Equal(t, AWSSecretAccessKey, flags.AWSSecretAccessKeyFV)
	assert.Equal(t, AWSSessionToken, flags.AWSSessionTokenFV)

	assert.Equal(t, CorsoPassphrase, flags.PassphraseFV)
}

func PreparedProviderFlags() []string {
	return []string{
		"--" + flags.AzureClientIDFN, AzureClientID,
		"--" + flags.AzureClientTenantFN, AzureTenantID,
		"--" + flags.AzureClientSecretFN, AzureClientSecret,
	}
}

func AssertProviderFlags(t *testing.T, cmd *cobra.Command) {
	assert.Equal(t, AzureClientID, flags.AzureClientIDFV)
	assert.Equal(t, AzureTenantID, flags.AzureClientTenantFV)
	assert.Equal(t, AzureClientSecret, flags.AzureClientSecretFV)
}
