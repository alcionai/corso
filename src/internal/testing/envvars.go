package testing

import (
	"errors"
	"os"
)

func RequireEnvVars(evs ...string) (map[string]string, error) {
	vals := map[string]string{}
	for _, ev := range evs {
		ge := os.Getenv(ev)
		if len(ge) == 0 {
			return nil, errors.New(ev + " env var required for test suite")
		}
		vals[ev] = ge
	}
	return vals, nil
}
