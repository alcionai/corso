package credentials

import "os"

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
