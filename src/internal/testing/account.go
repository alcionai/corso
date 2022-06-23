package testing

import (
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/pkg/errors"
)

// NewM365Account returns an account.Account object initialized with environment
// variables used for integration tests that use Graph Connector.
func NewM365Account() (account.Account, error) {
	cfg, err := readTestConfig()
	if err != nil {
		return account.Account{}, errors.Wrap(err, "configuring m365 account from test file")
	}

	return account.NewAccount(
		account.ProviderM365,
		account.M365Config{
			M365:     credentials.GetM365(),
			TenantID: cfg[testCfgTenantID],
		},
	)
}
