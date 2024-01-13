package ptr

import (
	"time"

	"github.com/alcionai/corso/src/pkg/dttm"
)

// ptr package is a common package used for pointer
// access and deserialization.

// Val Generic function for dereferencing pointers.
// Microsoft Graph saves many variables as string pointers.
// Function will safely check if the point is nil prior to
// dereferencing the pointer. If the pointer is nil,
// an empty version of the object is returned.
// Operation does not work on Nested objects.
// For example:
// *evt.GetEnd().GetDateTime() will still cause a panic
// if evt is nil or GetEnd() is nil
func Val[T any](ptr *T) T {
	if ptr == nil {
		return *new(T)
	}

	return *ptr
}

// ValOK behaves the same as Val, except it also gives
// a boolean response for whether the pointer was nil
// (false) or non-nil (true).
func ValOK[T any](ptr *T) (T, bool) {
	if ptr == nil {
		return *new(T), false
	}

	return *ptr, true
}

// OrNow returns the value of the provided time, if the
// parameter is non-nil.  Otherwise it returns the current
// time in UTC.
func OrNow(t *time.Time) time.Time {
	if t == nil {
		t = &time.Time{}
	}

	return dttm.OrNow(*t)
}

// To generates a pointer from any value.  Primarily useful
// for generating pointers to strings and other primitives
// without needing to store a second variable.
func To[T any](t T) *T {
	return &t
}
