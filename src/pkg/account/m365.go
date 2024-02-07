package account

import (
	"encoding/json"
	"reflect"
	"slices"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/common/str"
	"github.com/alcionai/canario/src/pkg/credentials"
)

// config exported name consts
const (
	AzureTenantID = "AZURE_TENANT_ID"
)

var excludedM365ConfigFieldsForHashing = []string{"AzureClientSecret"}

type M365Config struct {
	credentials.M365 // requires: ClientID, ClientSecret
	AzureTenantID    string
}

// config key consts
const (
	keyAzureClientID     = "azure_clientid"
	keyAzureClientSecret = "azure_clientSecret"
	keyAzureTenantID     = "azure_tenantid"
)

// StringConfig transforms a m365Config struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c M365Config) StringConfig() (map[string]string, error) {
	cfg := map[string]string{
		keyAzureClientID:     c.AzureClientID,
		keyAzureClientSecret: c.AzureClientSecret,
		keyAzureTenantID:     c.AzureTenantID,
	}

	return cfg, c.validate()
}

// providerID returns the c.TenantID if ap is a ProviderM365.
func (c M365Config) providerID(ap accountProvider) string {
	if ap == ProviderM365 {
		return c.AzureTenantID
	}

	return ""
}

func (c M365Config) configHash() (string, error) {
	filteredM365Config := createFilteredM365ConfigForHashing(c)

	b, err := json.Marshal(filteredM365Config)
	if err != nil {
		return "", clues.Stack(err)
	}

	return str.GenerateHash(b), nil
}

// M365Config retrieves the M365Config details from the Account config.
func (a Account) M365Config() (M365Config, error) {
	c := M365Config{}
	if len(a.Config) > 0 {
		c.AzureClientID = a.Config[keyAzureClientID]
		c.AzureClientSecret = a.Config[keyAzureClientSecret]
		c.AzureTenantID = a.Config[keyAzureTenantID]
	}

	return c, c.validate()
}

func createFilteredM365ConfigForHashing(source M365Config) map[string]any {
	filteredM365Config := make(map[string]any)
	sourceValue := reflect.ValueOf(source)

	for i := 0; i < sourceValue.NumField(); i++ {
		fieldName := sourceValue.Type().Field(i).Name
		if !slices.Contains(excludedM365ConfigFieldsForHashing, fieldName) {
			filteredM365Config[fieldName] = sourceValue.Field(i).Interface()
		}
	}

	return filteredM365Config
}

func (c M365Config) validate() error {
	check := map[string]string{
		credentials.AzureClientID:     c.AzureClientID,
		credentials.AzureClientSecret: c.AzureClientSecret,
		AzureTenantID:                 c.AzureTenantID,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}
