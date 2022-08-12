package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
)

// prerequisite: readRepoConfig must have been run prior to this to populate the global viper values.
func m365ConfigsFromViper(vpr *viper.Viper) (account.M365Config, error) {
	var m365 account.M365Config

	providerType := vpr.GetString(AccountProviderTypeKey)
	if providerType != account.ProviderM365.String() {
		return m365, errors.New("unsupported account provider: " + providerType)
	}

	m365.TenantID = vpr.GetString(TenantIDKey)

	return m365, nil
}

func m365Overrides(in map[string]string) map[string]string {
	return map[string]string{
		account.TenantID:       in[account.TenantID],
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
			return acct, errors.Wrap(err, "reading m365 configs from corso config file")
		}

		if err := mustMatchConfig(vpr, m365Overrides(overrides)); err != nil {
			return acct, errors.Wrap(err, "verifying m365 configs in corso config file")
		}
	}

	// compose the m365 config and credentials
	m365 := credentials.GetM365()
	if err := m365.Validate(); err != nil {
		return acct, errors.Wrap(err, "validating m365 credentials")
	}

	m365Cfg = account.M365Config{
		M365:     m365,
		TenantID: common.First(overrides[account.TenantID], m365Cfg.TenantID, os.Getenv(account.TenantID)),
	}

	// ensure required properties are present
	if err := utils.RequireProps(map[string]string{
		credentials.ClientID:     m365Cfg.ClientID,
		credentials.ClientSecret: m365Cfg.ClientSecret,
		account.TenantID:         m365Cfg.TenantID,
	}); err != nil {
		return acct, err
	}

	// build the account
	acct, err = account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return acct, errors.Wrap(err, "retrieving m365 account configuration")
	}

	return acct, nil
}
