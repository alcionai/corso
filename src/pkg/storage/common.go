package storage

import (
	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/credentials"
)

type CommonConfig struct {
	credentials.Corso // requires: CorsoPassphrase

	KopiaCfgDir string
}

// config key consts
const (
	keyCommonCorsoPassphrase = "common_corsoPassphrase"
	keyCommonKopiaCfgDir     = "common_kopiaCfgDir"
)

// StringConfig transforms a commonConfig struct into a plain
// map[string]string.  All values in the original struct which
// serialize into the map are expected to be strings.
func (c CommonConfig) StringConfig() (map[string]string, error) {
	cfg := map[string]string{
		keyCommonCorsoPassphrase: c.CorsoPassphrase,
		keyCommonKopiaCfgDir:     c.KopiaCfgDir,
	}

	return cfg, c.validate()
}

// CommonConfig retrieves the CommonConfig details from the Storage config.
func (s Storage) CommonConfig() (CommonConfig, error) {
	c := CommonConfig{}

	if len(s.Config) > 0 {
		c.CorsoPassphrase = orEmptyString(s.Config[keyCommonCorsoPassphrase])
		c.KopiaCfgDir = orEmptyString(s.Config[keyCommonKopiaCfgDir])
	}

	return c, c.validate()
}

// ensures all required properties are present
func (c CommonConfig) validate() error {
	if len(c.CorsoPassphrase) == 0 {
		return clues.Stack(errMissingRequired, errors.New(credentials.CorsoPassphrase))
	}

	// kopiaCfgFilePath is not required
	return nil
}
