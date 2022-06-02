package storage

import "github.com/pkg/errors"

type CommonConfig struct {
	CorsoPassword string // required
}

// envvar consts
const (
	CORSO_PASSWORD = "CORSO_PASSWORD"
)

// config key consts
const (
	keyCommonCorsoPassword = "common_corsoPassword"
)

func (c CommonConfig) Config() (config, error) {
	cfg := config{
		keyCommonCorsoPassword: c.CorsoPassword,
	}
	return cfg, c.validate()
}

// CommonConfig retrieves the CommonConfig details from the Storage config.
func (s Storage) CommonConfig() (CommonConfig, error) {
	c := CommonConfig{}
	if len(s.Config) > 0 {
		c.CorsoPassword = orEmptyString(s.Config[keyCommonCorsoPassword])
	}
	return c, c.validate()
}

// ensures all required properties are present
func (c CommonConfig) validate() error {
	if len(c.CorsoPassword) == 0 {
		return errors.Wrap(errMissingRequired, CORSO_PASSWORD)
	}
	return nil
}
