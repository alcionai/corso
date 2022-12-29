package tester

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

// MustGetEnvVars retrieves the provided env vars from the os.
// Retrieved values are populated into the resulting map.
// If any of the env values are zero length, the test errors.
func MustGetEnvVars(t *testing.T, evs ...string) map[string]string {
	vals := map[string]string{}

	for _, ev := range evs {
		ge := os.Getenv(ev)
		require.NotEmpty(t, ev, ev+" env var required for test suite")

		vals[ev] = ge
	}

	return vals
}

// MustGetEnvSls retrieves the provided env vars from the os.
// Retrieved values are populated into the resulting map.
// If any of the env values are zero length, the test errors.
func MustGetEnvSets(t *testing.T, evs ...[]string) map[string]string {
	vals := map[string]string{}

	for _, ev := range evs {
		r := MustGetEnvVars(t, ev...)
		maps.Copy(vals, r)
	}

	return vals
}
