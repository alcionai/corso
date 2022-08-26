package tester

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
)

var M365AcctCredEnvs = []string{
	credentials.ClientID,
	credentials.ClientSecret,
}

// NewM365Account returns an account.Account object initialized with environment
// variables used for integration tests that use Graph Connector.
func NewM365Account(t *testing.T) account.Account {
	cfg, err := readTestConfig()
	require.NoError(t, err, "configuring m365 account from test configuration")

	acc, err := account.NewAccount(
		account.ProviderM365,
		account.M365Config{
			M365:     credentials.GetM365(),
			TenantID: cfg[TestCfgTenantID],
		},
	)
	require.NoError(t, err, "initializing account")

	return acc
}
