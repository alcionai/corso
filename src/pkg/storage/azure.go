package storage

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common"
)

type AzureConfig struct {
	Container      string // required
	Prefix         string
	StorageAccount string
	StorageKey     string
}

// config key consts
const (
	keyAzContainer = "az_container"
	keyAzPrefix    = "az_prefix"
)

// config exported name consts
const (
	Container = "container"
)

func (c AzureConfig) Normalize() AzureConfig {
	return AzureConfig{
		Container: c.Container,
		Prefix:    common.NormalizePrefix(c.Prefix),
	}
}

// StringConfig transforms a s3Config struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c AzureConfig) StringConfig() (map[string]string, error) {
	cn := c.Normalize()
	cfg := map[string]string{
		keyAzContainer: cn.Container,
		keyAzPrefix:    cn.Prefix,
	}

	return cfg, c.validate()
}

// S3Config retrieves the S3Config details from the Storage config.
func (s Storage) AzureConfig() (AzureConfig, error) {
	c := AzureConfig{}

	if len(s.Config) > 0 {
		c.Container = orEmptyString(s.Config[keyAzContainer])
		c.Prefix = orEmptyString(s.Config[keyAzPrefix])
	}

	return c, c.validate()
}

func (c AzureConfig) validate() error {
	check := map[string]string{
		Container: c.Container,
	}
	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}
