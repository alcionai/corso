package account

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/credentials"
)

// config exported name consts
const (
	TenantID = "TENANT_ID"
)

type M365Config struct {
	credentials.M365 // requires: ClientID, ClientSecret
	TenantID         string
}

// config key consts
const (
	keyM365ClientID     = "m365_clientID"
	keyM365ClientSecret = "m365_clientSecret"
	keyM365TenantID     = "m365_tenantID"
)

// StringConfig transforms a m365Config struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c M365Config) StringConfig() (map[string]string, error) {
	cfg := map[string]string{
		keyM365ClientID:     c.ClientID,
		keyM365ClientSecret: c.ClientSecret,
		keyM365TenantID:     c.TenantID,
	}

	return cfg, c.validate()
}

// providerID returns the c.TenantID if ap is a ProviderM365.
func (c M365Config) providerID(ap accountProvider) string {
	if ap == ProviderM365 {
		return c.TenantID
	}

	return ""
}

// M365Config retrieves the M365Config details from the Account config.
func (a Account) M365Config() (M365Config, error) {
	c := M365Config{}
	if len(a.Config) > 0 {
		c.ClientID = a.Config[keyM365ClientID]
		c.ClientSecret = a.Config[keyM365ClientSecret]
		c.TenantID = a.Config[keyM365TenantID]
	}

	return c, c.validate()
}

func (c M365Config) validate() error {
	check := map[string]string{
		credentials.ClientID:     c.ClientID,
		credentials.ClientSecret: c.ClientSecret,
		TenantID:                 c.TenantID,
	}

	for k, v := range check {
		if len(v) == 0 {
			return errors.Wrap(errMissingRequired, k)
		}
	}

	return nil
}
