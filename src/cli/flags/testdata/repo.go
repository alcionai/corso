package testdata

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/canario/src/cli/flags"
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

func PreparedGenericBackupFlags() []string {
	return []string{
		"--" + flags.FailFastFN,
		"--" + flags.DisableIncrementalsFN,
		"--" + flags.ForceItemDataDownloadFN,
	}
}

func AssertGenericBackupFlags(t *testing.T, cmd *cobra.Command) {
	assert.True(t, flags.FailFastFV, "fail fast flag")
	assert.True(t, flags.DisableIncrementalsFV, "disable incrementals flag")
	assert.True(t, flags.ForceItemDataDownloadFV, "force item data download flag")
}
