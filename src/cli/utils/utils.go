package utils

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/alcionai/corso/pkg/repository"
)

// RequireProps validates the existence of the properties
//  in the map.  Expects the format map[propName]propVal.
func RequireProps(props map[string]string) error {
	for name, val := range props {
		if len(val) == 0 {
			return errors.New(name + " is required to perform this command")
		}
	}
	return nil
}

// aggregates m365 details from flag and env_var values.
type m365Vars struct {
	ClientID     string
	ClientSecret string
	TenantID     string
}

// GetM365Vars is a helper for aggregating m365 connection details.
func GetM365Vars() m365Vars {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	return m365Vars{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TenantID:     os.Getenv("TENANT_ID"),
	}
}

// CloseRepo handles closing a repo.
func CloseRepo(ctx context.Context, r *repository.Repository) {
	if err := r.Close(ctx); err != nil {
		fmt.Print("Error closing repository:", err)
	}
}
