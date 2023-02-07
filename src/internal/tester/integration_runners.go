package tester

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

const (
	CorsoLoadTests                                = "CORSO_LOAD_TESTS"
	CorsoCITests                                  = "CORSO_CI_TESTS"
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

// File needs to be a single message .json
// Use: https://developer.microsoft.com/en-us/graph/graph-explorer for details
const CorsoGraphConnectorTestSupportFile = "CORSO_TEST_SUPPORT_FILE"

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

// LogTimeOfTest logs the test name and the time that it was run.
func LogTimeOfTest(t *testing.T) string {
	now := time.Now().UTC().Format(time.RFC3339Nano)
	name := t.Name()

	if name == "" {
		t.Logf("Test run at %s.", now)
		return now
	}

	t.Logf("%s run at %s", name, now)

	return now
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
