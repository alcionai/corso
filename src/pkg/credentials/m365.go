package credentials

import (
	"os"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
)

// envvar consts
const (
	AzureClientID     = "AZURE_CLIENT_ID"
	AzureClientSecret = "AZURE_CLIENT_SECRET"
)

// M365 aggregates m365 credentials from flag and env_var values.
type M365 struct {
	AzureClientID     string
	AzureClientSecret string
}

// M365 is a helper for aggregating m365 secrets and credentials.
func GetM365() M365 {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	return M365{
		AzureClientID:     os.Getenv(AzureClientID),
		AzureClientSecret: os.Getenv(AzureClientSecret),
	}
}

func (c M365) Validate() error {
	check := map[string]string{
		AzureClientID:     c.AzureClientID,
		AzureClientSecret: c.AzureClientSecret,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, errors.New(k))
		}
	}

	return nil
}
