package storage

import (
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/credentials"
)

// Move this to config
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

// storage parsing errors
var (
	errMissingRequired = clues.New("missing required storage configuration")
)

// ensures all required properties are present
func (c CommonConfig) validate() error {
	if len(c.CorsoPassphrase) == 0 {
		return clues.Stack(errMissingRequired, clues.New(credentials.CorsoPassphrase))
	}

	// kopiaCfgFilePath is not required
	return nil
}

// Helper for parsing the values in a config object.
// If the value is nil or not a string, returns an empty string.
func orEmptyString(v any) string {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("panic recovery casting %v to string\n", v)
		}
	}()

	if v == nil {
		return ""
	}

	return v.(string)
}
