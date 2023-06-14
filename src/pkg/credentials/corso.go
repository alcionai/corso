package credentials

import (
	"os"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/str"
)

// envvar consts
const (
	CorsoPassphrase = "CORSO_PASSPHRASE"
	// Corso Flags
	CorsoPassphraseFN = "passphrase"
)

var CorsoPassphraseFV string

// Corso aggregates corso credentials from flag and env_var values.
type Corso struct {
	CorsoPassphrase string // required
}

// GetCorso is a helper for aggregating Corso secrets and credentials.
func GetCorso() Corso {
	// todo (rkeeprs): read from either corso config file or env vars.
	// https://github.com/alcionai/corso/issues/120
	corsoPassph := str.First(CorsoPassphraseFV, os.Getenv(CorsoPassphrase))

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
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}
