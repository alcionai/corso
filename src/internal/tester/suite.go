package tester

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"
)

// Flags for declaring which scope of tests to run.
const (
	CorsoCITests   = "CORSO_CI_TESTS"
	CorsoE2ETests  = "CORSO_E2E_TESTS"
	CorsoLoadTests = "CORSO_LOAD_TESTS"

	CorsoCLIBackupTests                           = "CORSO_COMMAND_LINE_BACKUP_TESTS"
	CorsoCLIConfigTests                           = "CORSO_COMMAND_LINE_CONFIG_TESTS"
	CorsoCLIRepoTests                             = "CORSO_COMMAND_LINE_REPO_TESTS"
	CorsoCLIRestoreTests                          = "CORSO_COMMAND_LINE_RESTORE_TESTS"
	CorsoCLITests                                 = "CORSO_COMMAND_LINE_TESTS"
	CorsoConnectorCreateCollectionTests           = "CORSO_CONNECTOR_CREATE_COLLECTION_TESTS"
	CorsoConnectorCreateExchangeCollectionTests   = "CORSO_CONNECTOR_CREATE_EXCHANGE_COLLECTION_TESTS"
	CorsoConnectorCreateSharePointCollectionTests = "CORSO_CONNECTOR_CREATE_SHAREPOINT_COLLECTION_TESTS"
	CorsoConnectorDataCollectionTests             = "CORSO_CONNECTOR_DATA_COLLECTION_TESTS"
	CorsoConnectorExchangeFolderCacheTests        = "CORSO_CONNECTOR_EXCHANGE_FOLDER_CACHE_TESTS"
	CorsoConnectorRestoreExchangeCollectionTests  = "CORSO_CONNECTOR_RESTORE_EXCHANGE_COLLECTION_TESTS"
	CorsoGraphConnectorTests                      = "CORSO_GRAPH_CONNECTOR_TESTS"
	CorsoGraphConnectorExchangeTests              = "CORSO_GRAPH_CONNECTOR_EXCHANGE_TESTS"
	CorsoGraphConnectorOneDriveTests              = "CORSO_GRAPH_CONNECTOR_ONE_DRIVE_TESTS"
	CorsoGraphConnectorSharePointTests            = "CORSO_GRAPH_CONNECTOR_SHAREPOINT_TESTS"
	CorsoKopiaWrapperTests                        = "CORSO_KOPIA_WRAPPER_TESTS"
	CorsoModelStoreTests                          = "CORSO_MODEL_STORE_TESTS"
	CorsoOneDriveTests                            = "CORSO_ONE_DRIVE_TESTS"
	CorsoOperationTests                           = "CORSO_OPERATION_TESTS"
	CorsoOperationBackupTests                     = "CORSO_OPERATION_BACKUP_TESTS"
	CorsoRepositoryTests                          = "CORSO_REPOSITORY_TESTS"
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
	//nolint:forbidigo
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
	//nolint:forbidigo
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
	//nolint:forbidigo
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
	//nolint:forbidigo
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
