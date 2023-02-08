package credentials

import (
	"os"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
)

// envvar consts
const (
	CorsoPassphrase = "CORSO_PASSPHRASE"
)

// Corso aggregates corso credentials from flag and env_var values.
type Corso struct {
	CorsoPassphrase string // required
}

// GetCorso is a helper for aggregating Corso secrets and credentials.
func GetCorso() Corso {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	corsoPassph := os.Getenv(CorsoPassphrase)

	return Corso{
		CorsoPassphrase: corsoPassph,
	}
}

func (c Corso) Validate() error {
	check := map[string]string{
		CorsoPassphrase: c.CorsoPassphrase,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, errors.New(k))
		}
	}

	return nil
}
