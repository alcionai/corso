package credentials

import (
	"os"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/str"
)

// envvar consts
const (
	AzureClientID     = "AZURE_CLIENT_ID"
	AzureClientSecret = "AZURE_CLIENT_SECRET"
)

var (
	AzureClientTenantFV string
	AzureClientIDFV     string
	AzureClientSecretFV string
)

// M365 aggregates m365 credentials from flag and env_var values.
type M365 struct {
	AzureClientID     string
	AzureClientSecret string
}

// M365 is a helper for aggregating m365 secrets and credentials.
func GetM365() M365 {
	// check env and overide is flags found
	// var AzureClientID, AzureClientSecret string
	AzureClientID := str.First(AzureClientIDFV, os.Getenv(AzureClientID))
	AzureClientSecret := str.First(AzureClientSecretFV, os.Getenv(AzureClientSecret))

	return M365{
		AzureClientID:     AzureClientID,
		AzureClientSecret: AzureClientSecret,
	}
}

func (c M365) Validate() error {
	check := map[string]string{
		AzureClientID:     c.AzureClientID,
		AzureClientSecret: c.AzureClientSecret,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}
