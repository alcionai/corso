package tester

import (
	"os"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"
)

// Flags for declaring which scope of tests to run.
const (
	CorsoCITests        = "CORSO_CI_TESTS"
	CorsoE2ETests       = "CORSO_E2E_TESTS"
	CorsoLoadTests      = "CORSO_LOAD_TESTS"
	CorsoNightlyTests   = "CORSO_NIGHTLY_TESTS"
	CorsoRetentionTests = "CORSO_RETENTION_TESTS"
)

type Suite interface {
	suite.TestingSuite
	Run(name string, subtest func()) bool
}

// ---------------------------------------------------------------------------
// Unit
// ---------------------------------------------------------------------------

func NewUnitSuite(t *testing.T) *unitSuite {
	return new(unitSuite)
}

type unitSuite struct {
	suite.Suite
}

// ---------------------------------------------------------------------------
// Integration
// ---------------------------------------------------------------------------

func NewIntegrationSuite(
	t *testing.T,
	envSets [][]string,
	runOnAnyEnv ...string,
) *integrationSuite {
	// ensure clues does not obscure logging
	clues.SetHasher(clues.NoHash())

	RunOnAny(
		t,
		append(
			[]string{CorsoCITests},
			runOnAnyEnv...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(integrationSuite)
}

type integrationSuite struct {
	suite.Suite
}

// ---------------------------------------------------------------------------
// Smoke/e2e
// ---------------------------------------------------------------------------

func NewE2ESuite(
	t *testing.T,
	envSets [][]string,
	runOnAnyEnv ...string,
) *e2eSuite {
	// ensure clues does not obscure logging
	clues.SetHasher(clues.NoHash())

	RunOnAny(
		t,
		append(
			[]string{CorsoE2ETests},
			runOnAnyEnv...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(e2eSuite)
}

type e2eSuite struct {
	suite.Suite
}

// ---------------------------------------------------------------------------
// Load
// ---------------------------------------------------------------------------

func NewLoadSuite(
	t *testing.T,
	envSets [][]string,
	runOnAnyEnv ...string,
) *loadSuite {
	// ensure clues does not obscure logging
	clues.SetHasher(clues.NoHash())

	RunOnAny(
		t,
		append(
			[]string{CorsoLoadTests},
			runOnAnyEnv...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(loadSuite)
}

type loadSuite struct {
	suite.Suite
}

// ---------------------------------------------------------------------------
// Nightly
// ---------------------------------------------------------------------------

func NewNightlySuite(
	t *testing.T,
	envSets [][]string,
	runOnAnyEnv ...string,
) *nightlySuite {
	// ensure clues does not obscure logging
	clues.SetHasher(clues.NoHash())

	RunOnAny(
		t,
		append(
			[]string{CorsoNightlyTests},
			runOnAnyEnv...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(nightlySuite)
}

type nightlySuite struct {
	suite.Suite
}

// ---------------------------------------------------------------------------
// Retention; requires object locking enabled on the S3 bucket.
// ---------------------------------------------------------------------------

func NewRetentionSuite(
	t *testing.T,
	envSets [][]string,
	runOnAnyEnv ...string,
) *retentionSuite {
	// ensure clues does not obscure logging
	clues.SetHasher(clues.NoHash())

	RunOnAny(
		t,
		append(
			[]string{CorsoRetentionTests},
			runOnAnyEnv...,
		)...,
	)

	MustGetEnvSets(t, envSets...)

	return new(retentionSuite)
}

type retentionSuite struct {
	suite.Suite
}

// ---------------------------------------------------------------------------
// Run Condition Checkers
// ---------------------------------------------------------------------------

// RunOnAny takes in a list of env variable names and returns
// an error if all of them are zero valued.  Implication being:
// if any of those env vars are truthy, you should run the
// subsequent tests.
func RunOnAny(t *testing.T, tests ...string) {
	var l int
	for _, test := range tests {
		l += len(os.Getenv(test))
	}

	if l == 0 {
		t.Skipf(
			"one or more env vars must be flagged to run this test: %v",
			strings.Join(tests, ", "))
	}
}

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
