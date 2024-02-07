package tconfig

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/credentials"
)

var M365AcctCredEnvs = []string{
	credentials.AzureClientID,
	credentials.AzureClientSecret,
}

// NewM365Account returns an account.Account object initialized with environment
// variables used for integration tests that use the m365 Controller.
func NewM365Account(t *testing.T) account.Account {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "configuring m365 account from test configuration", clues.ToCore(err))

	acc, err := account.NewAccount(
		account.ProviderM365,
		account.M365Config{
			M365:          credentials.GetM365(),
			AzureTenantID: cfg[TestCfgAzureTenantID],
		})
	require.NoError(t, err, "initializing account", clues.ToCore(err))

	return acc
}

func NewFakeM365Account(t *testing.T) account.Account {
	acc, err := account.NewAccount(
		account.ProviderM365,
		account.M365Config{
			M365: credentials.M365{
				AzureClientID:     "client-12345",
				AzureClientSecret: "secret-abcde",
			},
			AzureTenantID: "tenant-09876",
		})
	require.NoError(t, err, "initializing mock account", clues.ToCore(err))

	return acc
}
