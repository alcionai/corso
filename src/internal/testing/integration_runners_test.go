package testing

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

const (
	CORSO_CI_TESTS              = "CORSO_CI_TESTS"
	CORSO_GRAPH_CONNECTOR_TESTS = "CORSO_GRAPH_CONNECTOR_TESTS"
	CORSO_REPOSITORY_TESTS      = "CORSO_REPOSITORY_TESTS"
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
	now := time.Now().UTC().Format("2016-01-02T15:04:05")
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if !ok || details != nil {
		t.Logf("Test run at %s.", now)
		return now
	}
	t.Logf("%s() run at %s", details.Name(), now)
	return now
}
