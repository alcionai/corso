package testing

import (
	"reflect"
	"runtime"
	"testing"

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
