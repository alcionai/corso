package tester

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// AreSameFunc asserts whether the two funcs are the same func.
func AreSameFunc(t *testing.T, expect, have any) {
	assert.Equal(
		t,
		runtime.FuncForPC(
			reflect.
				ValueOf(expect).
				Pointer(),
		).Name(),
		runtime.FuncForPC(
			reflect.
				ValueOf(have).
				Pointer(),
		).Name(),
	)
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
