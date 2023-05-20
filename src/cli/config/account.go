package config

import (
	"os"

	"github.com/alcionai/clues"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
)

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func m365ConfigsFromViper(vpr *viper.Viper) (account.M365Config, error) {
	var m365 account.M365Config

	providerType := vpr.GetString(AccountProviderTypeKey)
	if providerType != account.ProviderM365.String() {
		return m365, clues.New("unsupported account provider: " + providerType)
	}

	m365.AzureTenantID = vpr.GetString(AzureTenantIDKey)

	return m365, nil
}

func m365Overrides(in map[string]string) map[string]string {
	return map[string]string{
		account.AzureTenantID:  in[account.AzureTenantID],
		AccountProviderTypeKey: in[AccountProviderTypeKey],
	}
}

// configureAccount builds a complete account configuration from a mix of
// viper properties and manual overrides.
func configureAccount(
	vpr *viper.Viper,
	readConfigFromViper bool,
	overrides map[string]string,
) (account.Account, error) {
	var (
		m365Cfg account.M365Config
		acct    account.Account
		err     error
	)

	if readConfigFromViper {
		m365Cfg, err = m365ConfigsFromViper(vpr)
		if err != nil {
			return acct, clues.Wrap(err, "reading m365 configs from corso config file")
		}

		if err := mustMatchConfig(vpr, m365Overrides(overrides)); err != nil {
			return acct, clues.Wrap(err, "verifying m365 configs in corso config file")
		}
	}

	// compose the m365 config and credentials
	m365 := credentials.GetM365()
	if err := m365.Validate(); err != nil {
		return acct, clues.Wrap(err, "validating m365 credentials")
	}

	m365Cfg = account.M365Config{
		M365: m365,
		AzureTenantID: str.First(
			overrides[account.AzureTenantID],
			m365Cfg.AzureTenantID,
			os.Getenv(account.AzureTenantID)),
	}

	// ensure required properties are present
	if err := requireProps(map[string]string{
		credentials.AzureClientID:     m365Cfg.AzureClientID,
		credentials.AzureClientSecret: m365Cfg.AzureClientSecret,
		account.AzureTenantID:         m365Cfg.AzureTenantID,
	}); err != nil {
		return acct, err
	}

	// build the account
	acct, err = account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return acct, clues.Wrap(err, "retrieving m365 account configuration")
	}

	return acct, nil
}
