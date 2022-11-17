package tester

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	CorsoLoadTests                                = "CORSO_LOAD_TESTS"
	CorsoCITests                                  = "CORSO_CI_TESTS"
	CorsoCLIBackupTests                           = "CORSO_COMMAND_LINE_BACKUP_TESTS"
	CorsoCLIConfigTests                           = "CORSO_COMMAND_LINE_CONFIG_TESTS"
	CorsoCLIRepoTests                             = "CORSO_COMMAND_LINE_REPO_TESTS"
	CorsoCLIRestoreTests                          = "CORSO_COMMAND_LINE_RESTORE_TESTS"
	CorsoCLITests                                 = "CORSO_COMMAND_LINE_TESTS"
	CorsoConnectorCreateExchangeCollectionTests   = "CORSO_CONNECTOR_CREATE_EXCHANGE_COLLECTION_TESTS"
	CorsoConnectorCreateSharePointCollectionTests = "CORSO_CONNECTOR_CREATE_SHAREPOINT_COLLECTION_TESTS"
	CorsoConnectorDataCollectionTests             = "CORSO_CONNECTOR_DATA_COLLECTION_TESTS"
	CorsoGraphConnectorTests                      = "CORSO_GRAPH_CONNECTOR_TESTS"
	CorsoGraphConnectorExchangeTests              = "CORSO_GRAPH_CONNECTOR_EXCHANGE_TESTS"
	CorsoGraphConnectorOneDriveTests              = "CORSO_GRAPH_CONNECTOR_ONE_DRIVE_TESTS"
	CorsoKopiaWrapperTests                        = "CORSO_KOPIA_WRAPPER_TESTS"
	CorsoModelStoreTests                          = "CORSO_MODEL_STORE_TESTS"
	CorsoOneDriveTests                            = "CORSO_ONE_DRIVE_TESTS"
	CorsoOperationTests                           = "CORSO_OPERATION_TESTS"
	CorsoRepositoryTests                          = "CORSO_REPOSITORY_TESTS"
)

// File needs to be a single message .json
// Use: https://developer.microsoft.com/en-us/graph/graph-explorer for details
const CorsoGraphConnectorTestSupportFile = "CORSO_TEST_SUPPORT_FILE"

// RunOnAny takes in a list of env variable names and returns
// an error if all of them are zero valued.  Implication being:
// if any of those env vars are truthy, you should run the
// subsequent tests.
func RunOnAny(tests ...string) error {
	var l int
	for _, test := range tests {
		l += len(os.Getenv(test))
	}

	if l == 0 {
		return fmt.Errorf(
			"%s env vars are not flagged for testing",
			strings.Join(tests, ", "))
	}

	return nil
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
