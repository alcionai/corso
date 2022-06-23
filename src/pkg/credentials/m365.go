package credentials

import (
	"os"

	"github.com/pkg/errors"
)

// envvar consts
const (
	ClientID     = "CLIENT_ID"
	ClientSecret = "CLIENT_SECRET"
)

// M365 aggregates m365 credentials from flag and env_var values.
type M365 struct {
	ClientID     string
	ClientSecret string
}

// M365 is a helper for aggregating m365 secrets and credentials.
func GetM365() M365 {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	return M365{
		ClientID:     os.Getenv(ClientID),
		ClientSecret: os.Getenv(ClientSecret),
	}
}

func (c M365) Validate() error {
	check := map[string]string{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
	}
	for k, v := range check {
		if len(v) == 0 {
			return errors.Wrap(errMissingRequired, k)
		}
	}
	return nil
}
