package credentials

import "os"

// M365 aggregates m365 credentials from flag and env_var values.
type M365 struct {
	ClientID     string
	ClientSecret string
	TenantID     string
}

// M365 is a helper for aggregating m365 secrets and credentials.
func GetM365() M365 {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	return M365{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TenantID:     os.Getenv("TENANT_ID"),
	}
}
