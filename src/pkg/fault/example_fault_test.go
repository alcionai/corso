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
	items = []string{}
)

type mockController struct {
	errors any
}

func connectClient() error              { return nil }
func dependencyCall() error             { return nil }
func getIthItem(i int) error            { return nil }
func getData() ([]string, error)        { return nil, nil }
func storeData([]string, *fault.Errors) {}

type mockOper struct {
	Errors *fault.Errors
}

func newOperation() mockOper          { return mockOper{fault.New(true)} }
func (m mockOper) Run() *fault.Errors { return m.Errors }

type mockDepenedency struct{}

func (md mockDepenedency) do() error {
	return errors.New("caught one")
}

var dependency = mockDepenedency{}

// ---------------------------------------------------------------------------
// examples
// ---------------------------------------------------------------------------

// ExampleNew highlights assumptions and best practices
// for generating fault.Errors structs.
func ExampleNew() {
	// fault.Errors should only be generated during the construction of
	// another controller, such as a new Backup or Restore Operations.
	// Configurations like failFast are set during construction.
	//
	// Generating new fault.Errors structs outside of an operation
	// controller is a smell, and should be avoided.  If you need
	// to aggregate errors, you should accept an interface and pass
	// an fault.Errors instance into it.
	ctrl = mockController{
		errors: fault.New(false),
	}
}

// ExampleErrors_Fail describes the assumptions and best practices
// for setting the Failure error.
func ExampleErrors_Fail() {
	errs := fault.New(false)

	// Fail() is used to record non-recoverable errors.
	//
	// Fail() should only get called in the last step before returning
	// a fault.Errors from a controller.  In all other cases, you
	// should simply return an error and expect the upstream controller
	// to call Fail() for you.
	topLevelHandler := func(errs *fault.Errors) *fault.Errors {
		if err := connectClient(); err != nil {
			return errs.Fail(err)
		}

		return errs
	}
	if errs := topLevelHandler(errs); errs.Err() != nil {
		fmt.Println(errs.Err())
	}

	// Only the topmost func in the stack should set the Fail() err.
	// IE: Fail() is not Wrap().  In lower levels, errors should get
	// wrapped and returned like normal, and only handled by fault
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

// ExampleErrors_Add describes the assumptions and best practices
// for aggregating iterable or recoverable errors.
func ExampleErrors_Add() {
	errs := fault.New(false)

	// Add() is used to record any recoverable error.
	//
	// Add() should only get called as the last error handling step
	// within a loop or stream.  In all other cases, you can return
	// an error like normal and expect the upstream point of iteration
	// to call Add() for you.
	for i := range items {
		clientBasedGetter := func(i int) error {
			if err := getIthItem(i); err != nil {
				// lower level calls don't Add to the fault.Errors.
				// they handl errors like normal.
				return errors.Wrap(err, "dependency")
			}

			return nil
		}

		if err := clientBasedGetter(i); err != nil {
			// Here at the top of the loop is the correct place
			// to Add an error using fault.
			errs.Add(err)
		}
	}

	// Iteration should exit anytime the primary error in fault is
	// non-nil.  fault.Errors does not expose the failFast flag
	// directly.  Instead, errors from Add() will automatically
	// promote to the Err() value.  Therefore, loops only ned to
	// check the errs.Err().  If it is non-nil, then the loop should break.
	for i := range items {
		if errs.Err() != nil {
			// if failFast == true errs.Add() was called,
			// we'll catch the error here.
			break
		}

		if err := getIthItem(i); err != nil {
			errs.Add(err)
		}
	}
}

// ExampleErrors_Err describes retrieving the non-recoverable error.
func ExampleErrors_Err() {
	errs := fault.New(false)
	errs.Fail(errors.New("catastrophe"))

	// Err() returns the primary failure.
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
	// Recoverable errors may still have been added using Add(err).
	//
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

// ExampleErrors_Errs describes retrieving individual errors.
func ExampleErrors_Errs() {
	errs := fault.New(false)
	errs.Add(errors.New("not catastrophic"))
	errs.Add(errors.New("something unwanted"))

	// Errs() gets the slice of all recoverable errors Add()ed during
	// the run, but which did not force the process to exit.
	//
	// Errs() only needs to be investigated by the end user at the
	// conclusion of an operation.  Checking Errs() within lower-
	// layer code is a smell.  Funcs should return a standard error,
	// or errs.Err(), if they need upstream handlers to handle the errors.
	errSl := errs.Errs()
	for _, err := range errSl {
		fmt.Println(err)
	}

	// One or more errors in errs.Errs() does not necessarily mean the
	// process failed.  You can have non-zero Errs() but a nil Err().
	if errs.Err() == nil {
		fmt.Println("Err() is nil")
	}

	// If Errs() is nil, then you can assume that no recoverable or
	// iteration-based errors occurred.  But that does not necessarily
	// mean the operation was able to complete.
	//
	// Even if Errs() contains zero items, Err() can be non-nil.
	// Make sure you check both.

	// Output: not catastrophic
	// something unwanted
	// Err() is nil
}

func ExampleTracker() {
	// It is common for Corso to run operations in parallel,
	// and for iterations to be nested within iterations.  To
	// avoid mistakenly returning an error that was sourced
	// from some other async iteration, recoverable instances
	// are aggrgated into a Tracker.
	errs := fault.New(false)
	trkr := errs.Tracker()

	err := func() error {
		for i := range items {
			// note that we check errs.Err() on every iteration,
			// not trkr.Err().  The loop should break if any
			// hard failure occurs, not just one within this loop.
			if errs.Err() != nil {
				break
			}

			if err := getIthItem(i); err != nil {
				// instead of calling errs.Add(err), we call the
				// trackers Add method.  The error will still get
				// added to the errs.Errs() set.  But if this err
				// causes the run to fail, only this tracker treats
				// it as the causal failure.
				trkr.Add(err)
			}
		}

		return trkr.Err()
	}()
	if err != nil {
		// handle the Err() that appeared in the tracker
		fmt.Println("err occurred", errs.Err())
	}

}

// ExampleErrorsE2e showcases a more complex integration.
func Example_errors_e2e() {
	oper := newOperation()

	// imagine that we're a user, calling into corso SDK.
	// (fake funcs used here to minimize example bloat)
	//
	// The operation is our controller, we expect it to
	// generate a new fault.Errors when constructed, and
	// to return that struct when we call Run()
	errs := oper.Run()

	// Let's investigate what went on inside.  Since we're at
	// the top of our controller, and returning a fault.Errors,
	// all the error handlers set the Fail() case.
	/* Run() */
	func() *fault.Errors {
		if err := connectClient(); err != nil {
			// Fail() here; we're top level in the controller
			// and this is a non-recoverable issue
			return oper.Errors.Fail(err)
		}

		data, err := getData()
		if err != nil {
			return oper.Errors.Fail(err)
		}

		// storeData will aggregate iterated errors into
		// oper.Errors.
		storeData(data, oper.Errors)

		// return oper.Errors here, in part to ensure it's
		// non-nil, and because we don't know if we've
		// aggregated any iterated errors yet.
		return oper.Errors
	}()

	// What about the lower level handling?  storeData didn't
	// return an error, so what's happening there?
	/* storeData */
	func(data []any, errs *fault.Errors) {
		// this is downstream in our code somewhere
		storer := func(a any) error {
			if err := dependencyCall(); err != nil {
				// we're not passing in or calling fault.Errors here,
				// because this isn't the iteration handler, it's just
				// a regular error.
				return errors.Wrap(err, "dependency")
			}

			return nil
		}

		for _, d := range data {
			if errs.Err() != nil {
				break
			}

			if err := storer(d); err != nil {
				// Since we're at the top of the iteration, we need
				// to add each error to the fault.Errors struct.
				errs.Add(err)
			}
		}
	}(nil, nil)

	// then at the end of the oper.Run, we investigate the results.
	if errs.Err() != nil {
		// handle the primary error
		fmt.Println("err occurred", errs.Err())
	}

	for _, err := range errs.Errs() {
		// handle each recoverable error
		fmt.Println("recoverable err occurred", err)
	}
}

// ExampleErrors_Err_return showcases when to return err or nil vs errs.Err()
func ExampleErrors_Err_return() {
	// The general rule of thumb is to always handle the error directly
	// by returning err, or nil, or any variety of extension (wrap,
	// stack, clues, etc).
	fn := func() error {
		if err := dependency.do(); err != nil {
			return errors.Wrap(err, "direct")
		}

		return nil
	}
	if err := fn(); err != nil {
		fmt.Println(err)
	}

	// The exception is if you're handling recoverable errors.  Those
	// funcs should always return errs.Err(), in case a recoverable
	// error happened on the last round of iteration.
	fn2 := func(todo []string, errs *fault.Errors) error {
		for range todo {
			if errs.Err() != nil {
				return errs.Err()
			}

			if err := dependency.do(); err != nil {
				errs.Add(errors.Wrap(err, "recoverable"))
			}
		}

		return errs.Err()
	}
	if err := fn2([]string{"a"}, fault.New(true)); err != nil {
		fmt.Println(err)
	}

	// Output: direct: caught one
	// recoverable: caught one
}
