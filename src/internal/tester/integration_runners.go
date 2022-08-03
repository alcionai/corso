package tester

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	CorsoCITests             = "CORSO_CI_TESTS"
	CorsoCLIConfigTests      = "CORSO_CLI_CONFIG_TESTS"
	CorsoCLIRepoTests        = "CORSO_CLI_REPO_TESTS"
	CorsoGraphConnectorTests = "CORSO_GRAPH_CONNECTOR_TESTS"
	CorsoKopiaWrapperTests   = "CORSO_KOPIA_WRAPPER_TESTS"
	CorsoModelStoreTests     = "CORSO_MODEL_STORE_TESTS"
	CorsoOperationTests      = "CORSO_OPERATION_TESTS"
	CorsoRepositoryTests     = "CORSO_REPOSITORY_TESTS"
)

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
	now := time.Now().UTC().Format("2006-01-02T15:04:05.0000")
	name := t.Name()
	if name == "" {
		t.Logf("Test run at %s.", now)
		return now
	}
	t.Logf("%s run at %s", name, now)
	return now
}
