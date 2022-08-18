package storage

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/credentials"
)

type CommonConfig struct {
	credentials.Corso // requires: CorsoPassword

	KopiaCfgDir string
}

// config key consts
const (
	keyCommonCorsoPassword = "common_corsoPassword"
	keyCommonKopiaCfgDir   = "common_kopiaCfgDir"
)

// StringConfig transforms a commonConfig struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c CommonConfig) StringConfig() (map[string]string, error) {
	cfg := map[string]string{
		keyCommonCorsoPassword: c.CorsoPassword,
		keyCommonKopiaCfgDir:   c.KopiaCfgDir,
	}
	return cfg, c.validate()
}

// CommonConfig retrieves the CommonConfig details from the Storage config.
func (s Storage) CommonConfig() (CommonConfig, error) {
	c := CommonConfig{}
	if len(s.Config) > 0 {
		c.CorsoPassword = orEmptyString(s.Config[keyCommonCorsoPassword])
		c.KopiaCfgDir = orEmptyString(s.Config[keyCommonKopiaCfgDir])
	}
	return c, c.validate()
}

// ensures all required properties are present
func (c CommonConfig) validate() error {
	if len(c.CorsoPassword) == 0 {
		return errors.Wrap(errMissingRequired, credentials.CorsoPassword)
	}
	// kopiaCfgFilePath is not required
	return nil
}
