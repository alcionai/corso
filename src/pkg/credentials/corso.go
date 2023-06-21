package credentials

import (
	"github.com/alcionai/clues"
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
