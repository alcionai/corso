package account

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/credentials"
)

type M365Config struct {
	credentials.M365 // requires: ClientID, ClientSecret, TenantID

	// (todo) TenantID string
}

// config key consts
const (
	keyM365ClientID     = "m365_clientID"
	keyM365ClientSecret = "m365_clientSecret"
	keyM365TenantID     = "m365_tenantID"
)

// config exported name consts
const (
// (todo) TenantID     = "TENANT_ID"
)

func (c M365Config) Config() (config, error) {
	cfg := config{
		keyM365ClientID:     c.ClientID,
		keyM365ClientSecret: c.ClientSecret,
		keyM365TenantID:     c.TenantID,
	}
	return cfg, c.validate()
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
		credentials.TenantID:     c.TenantID,
	}
	for k, v := range check {
		if len(v) == 0 {
			return errors.Wrap(errMissingRequired, k)
		}
	}
	return nil
}
