package fault_test

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// mock helpers
// ---------------------------------------------------------------------------

var (
	ctrl  any
	anErr error
	items = []string{}
)

type mockController struct {
	errors any
}

func connectClient() error      { return nil }
func dependencyCall() error     { return nil }
func getIthItem(i string) error { return nil }

// ---------------------------------------------------------------------------
// examples
// ---------------------------------------------------------------------------

// ExampleNewErrors highlights assumptions and best practices
// for generating Errors structs.
func Example_new() {
	// Errors should only be generated during the construction of
	// another controller, such as a new Backup or Restore Operations.
	// Configurations like failFast are handled during construction.
	ctrl = mockController{
		errors: fault.New(false),
	}
}

// ExampleErrorsFail describes the assumptions and best practices
// for setting the Failure error.
func Example_errors_Fail() {
	errs := fault.New(false)

	// Fail() should be used to record any error that highlights a
	// non-recoverable failure in a process.
	if err := connectClient(); err != nil {
		errs.Fail(err)
	}

	// Only the topmost handler of the error should set the Fail() err.
	// This will normally be the operation controller itself.
	// IE: Fail() is not Wrap().  In lower levels, errors should get
	// wrapped and returned like normal, and only handled by errors
	// at the end.
	lowLevelCall := func() error {
		if err := dependencyCall(); err != nil {
			// wrap here, deeper into the stack
			return errors.Wrap(err, "dependency")
		}

		return nil
	}

	if err := lowLevelCall(); err != nil {
		// fail here, at the top of the stack
		errs.Fail(err)
	}
}

// ExampleErrorsAdd describes the assumptions and best practices
// for aggregating iterable or recoverable errors.
func Example_errors_Add() {
	errs := fault.New(false)

	// Add() should be used to record any error in a recoverable
	// part of processing.
	for _, i := range items {
		if err := getIthItem(i); err != nil {
			errs.Add(err)
		}
	}

	// In case of failFast behavior, iteration should exit as soon
	// as an error occurs.  Errors does not expose the failFast flag
	// directly.  Instead, iterators should check the value of Err().
	// If it is non-nil, then the loop shold break.
	for _, i := range items {
		if errs.Err() != nil {
			break
		}

		errs.Add(getIthItem(i))
	}

	// Only the topmost handler of the error should Add() the err.
	// This will normally be the iteration loop itself.
	// IE: Add() is not Wrap().  In lower levels, errors should get
	// wrapped and returned like normally, and only added to the
	// errors at the end.
	clientBasedGetter := func(s string) error {
		if err := dependencyCall(); err != nil {
			// wrap here, deeper into the stack
			return errors.Wrap(err, "dependency")
		}

		return nil
	}

	for _, i := range items {
		if err := clientBasedGetter(i); err != nil {
			// add here, within the iteraton loop
			errs.Add(err)
		}
	}
}

// ExampleErrorsErr describes retrieving the non-recoverable error.
func Example_errors_Err() {
	errs := fault.New(false)
	errs.Fail(errors.New("catastrophe"))

	// Err() gets the primary failure error.
	err := errs.Err()
	fmt.Println(err)

	// if multiple Failures occur, each one after the first gets
	// added to the Errs slice.
	errs.Fail(errors.New("another catastrophe"))
	errSl := errs.Errs()

	for _, e := range errSl {
		fmt.Println(e)
	}

	// If Err() is nil, then you can assume the operation completed.
	// A complete operation is not necessarily an error-free operation.
	// Even if Err() is nil, Errs() can be non-empty.
	// Make sure you check both.

	errs = fault.New(true)

	// If failFast is set to true, then the first error Add()ed gets
	// promoted to the Err() position.

	errs.Add(errors.New("not catastrophic, but still becomes the Err()"))
	err = errs.Err()
	fmt.Println(err)

	// Output: catastrophe
	// another catastrophe
	// not catastrophic, but still becomes the Err()
}

// ExampleErrorsErrs describes retrieving individual errors.
func Example_errors_Errs() {
	errs := fault.New(false)
	errs.Add(errors.New("not catastrophic"))
	errs.Add(errors.New("something unwanted"))

	// Errs() gets the slice errors that were recorded, but were
	// considered recoverable.
	errSl := errs.Errs()
	for _, err := range errSl {
		fmt.Println(err)
	}

	// If Errs() is nil, then you can assume that no recoverable or
	// iteration-based errors occurred.  But that does not necessarily
	// mean the operation was able to complete.
	// Even if Errs() contains zero items, Err() can be non-nil.
	// Make sure you check both.

	// Output: not catastrophic
	// something unwanted
}
