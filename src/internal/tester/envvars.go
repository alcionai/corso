package tester

import (
	"errors"
	"os"

	"golang.org/x/exp/maps"
)

// GetRequiredEnvVars retrieves the provided env vars from the os.
// Retrieved values are populated into the resulting map.
// If any of the env values are zero length, returns an error.
func GetRequiredEnvVars(evs ...string) (map[string]string, error) {
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

// GetRequiredEnvSls retrieves the provided env vars from the os.
// Retrieved values are populated into the resulting map.
// If any of the env values are zero length, returns an error.
func GetRequiredEnvSls(evs ...[]string) (map[string]string, error) {
	vals := map[string]string{}

	for _, ev := range evs {
		r, err := GetRequiredEnvVars(ev...)
		if err != nil {
			return nil, err
		}

		maps.Copy(vals, r)
	}

	return vals, nil
}
