package tester

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
)

func checkPopulated(v reflect.Value) error {
	if v.IsZero() {
		return clues.New("zero-valued field")
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	var errs *clues.Err

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if err := checkPopulated(f); err != nil {
			errs = clues.Stack(errs, clues.Wrap(err, fmt.Sprintf("field at index %d", i)))
		}
	}

	return errs.OrNil()
}

func isEmptyContainer(v reflect.Value) bool {
	// Handle pointers to things.
	deref := v

	for k := deref.Kind(); k == reflect.Pointer; k = deref.Kind() {
		deref = deref.Elem()
	}

	// Check for empty maps, slices, or arrays.
	if (deref.Kind() == reflect.Slice && deref.Len() == 0) ||
		(deref.Kind() == reflect.Map && deref.Len() == 0) ||
		(deref.Kind() == reflect.Array && deref.Len() == 0) {
		return true
	}

	return false
}

// CheckPopulated returns an error if input is not fully populated. To be
// considered fully populated it must be non-zero-valued. For basic types this
// just means it isn't the zero value. For structs this means that every field
// is not zero-valued. This check is recursive for structs.
func CheckPopulated(input any) error {
	return checkPopulated(reflect.ValueOf(input))
}

func checkNotPopulated(v reflect.Value) error {
	if isEmptyContainer(v) {
		return nil
	}

	if !v.IsZero() {
		return clues.New("non-zero-valued field")
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	var errs *clues.Err

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if err := checkNotPopulated(f); err != nil {
			errs = clues.Stack(errs, clues.Wrap(err, fmt.Sprintf("field at index %d", i)))
		}
	}

	return errs.OrNil()
}

// NilOrZero return true if the input is nil or if it's the zero value for the
// type. If the input is a struct then all fields are recursively checked to
// see if they're the zero value for their type.
func NilOrZero(input any) bool {
	if input == nil {
		return true
	}

	if isEmptyContainer(reflect.ValueOf(input)) {
		return true
	}

	err := checkNotPopulated(reflect.ValueOf(input))

	return err != nil
}

// AssertEmptyOrEqual checks either:
//   - got is nil or the zero value if expected is nil
//   - expected and got are equal if expected is not nil
func AssertEmptyOrEqual(
	t *testing.T,
	expect any,
	got any,
	msgAndArgs ...any,
) bool {
	if expect == nil {
		return assert.True(t, NilOrZero(got), "empty got value: %+v", got)
	}

	if isEmptyContainer(reflect.ValueOf(expect)) {
		return assert.True(t, NilOrZero(got), "empty got value: %+v", got)
	}

	return assert.Equal(t, expect, got, msgAndArgs...)
}
