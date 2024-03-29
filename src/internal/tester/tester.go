package tester

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AreSameFunc asserts whether the two funcs are the same func.
func AreSameFunc(t *testing.T, expect, have any) {
	assert.Equal(
		t,
		runtime.FuncForPC(
			reflect.
				ValueOf(expect).
				Pointer()).Name(),
		runtime.FuncForPC(
			reflect.
				ValueOf(have).
				Pointer()).Name())
}

type TestT interface {
	Logf(format string, args ...any)
	Name() string
	TempDir() string
	require.TestingT
}

// LogTimeOfTest logs the test name and the time that it was run.
func LogTimeOfTest(t TestT) string {
	now := time.Now().UTC().Format(time.RFC3339Nano)
	name := t.Name()

	if len(name) == 0 {
		t.Logf("Test run at %s.", now)
		return now
	}

	t.Logf("%s run at %s", name, now)

	return now
}
