package credentials

import (
	"os"

	"github.com/pkg/errors"
)

// envvar consts
const (
	CorsoPassword = "CORSO_PASSWORD"
)

// Corso aggregates corso credentials from flag and env_var values.
type Corso struct {
	CorsoPassword string // required
}

// GetCorso is a helper for aggregating Corso secrets and credentials.
func GetCorso() Corso {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	corsoPasswd := os.Getenv(CorsoPassword)
	return Corso{
		CorsoPassword: corsoPasswd,
	}
}

func (c Corso) Validate() error {
	check := map[string]string{
		CorsoPassword: c.CorsoPassword,
	}
	for k, v := range check {
		if len(v) == 0 {
			return errors.Wrap(errMissingRequired, k)
		}
	}
	return nil
}
