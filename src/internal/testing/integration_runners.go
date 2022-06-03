package testing

import (
	"fmt"
	"os"
	"strings"
)

const (
	CORSO_CI_TESTS              = "CORSO_CI_TESTS"
	CORSO_GRAPH_CONNECTOR_TESTS = "CORSO_GRAPH_CONNECTOR_TESTS"
	CORSO_REPOSITORY_TESTS      = "CORSO_REPOSITORY_TESTS"
)

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
