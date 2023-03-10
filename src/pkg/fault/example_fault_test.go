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

func connectClient() error           { return nil }
func dependencyCall() error          { return nil }
func getIthItem(i int) error         { return nil }
func getData() ([]string, error)     { return nil, nil }
func storeData([]string, *fault.Bus) {}

type mockOper struct {
	Errors *fault.Bus
}

func newOperation() mockOper       { return mockOper{fault.New(true)} }
func (m mockOper) Run() *fault.Bus { return m.Errors }

type mockDepenedency struct{}

func (md mockDepenedency) do() error {
	return errors.New("caught one")
}

var dependency = mockDepenedency{}

// ---------------------------------------------------------------------------
// examples
// ---------------------------------------------------------------------------

// ExampleNew highlights assumptions and best practices
// for generating fault.Bus structs.
func ExampleNew() {
	// New fault.Bus instances should only get generated during initialization.
	// Such as when starting up a new Backup or Restore Operation.
	// Configuration (eg: failFast) is set during construction and cannot
	// be updated.
	ctrl = mockController{
		errors: fault.New(false),
	}
}

// ExampleBus_Fail describes the assumptions and best practices
// for setting the Failure error.
func ExampleBus_Fail() {
	errs := fault.New(false)

	// Fail() is used to record non-recoverable errors.
	//
	// Fail() should only get called in the last step before returning
	// a fault.Bus from a controller.  In all other cases, you
	// can stick to standard golang error handling and expect some upstream
	// controller to call Fail() for you (if necessary).
	topLevelHandler := func(errs *fault.Bus) *fault.Bus {
		if err := connectClient(); err != nil {
			return errs.Fail(err)
		}

		return errs
	}
	if errs := topLevelHandler(errs); errs.Failure() != nil {
		fmt.Println(errs.Failure())
	}

	// Only the top-most func in the stack should set the failure.
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

// ExampleBus_AddRecoverable describes the assumptions and best practices
// for aggregating iterable or recoverable errors.
func ExampleBus_AddRecoverable() {
	errs := fault.New(false)

	// AddRecoverable() is used to record any recoverable error.
	//
	// What counts as a recoverable error?  That's up to the given
	// implementation.  Normally, it's an inability to process one
	// of many items within an iteration (ex: couldn't download 1 of
	// 1000 emails).  But just because an error occurred during a loop
	// doesn't mean it's recoverable, ex: a failure to retrieve the next
	// page when accumulating a batch of resources isn't usually
	// recoverable.  The choice is always up to the function at hand.
	//
	// AddRecoverable() should only get called as the top-most location
	// of error handling within the recoverable process.  Child functions
	// should stick to normal golang error handling and expect the upstream
	// controller to call AddRecoverable() for you.
	for i := range items {
		clientBasedGetter := func(i int) error {
			if err := getIthItem(i); err != nil {
				// lower level calls don't AddRecoverable to the fault.Bus.
				// they stick to normal golang error handling.
				return errors.Wrap(err, "dependency")
			}

			return nil
		}

		if err := clientBasedGetter(i); err != nil {
			// Here at the top of the loop is the correct place
			// to aggregate the error using fault.
			// Side note: technically, you should use a local bus
			// here (see below) instead of errs.
			errs.AddRecoverable(err)
		}
	}

	// Iteration should exit anytime the fault failure is non-nil.
	// fault.Bus does not expose the failFast flag directly.  Instead,
	// when failFast is true, errors from AddRecoverable() automatically
	// promote to the Failure() spot.  Recoverable handling only needs to
	// check the errs.Failure().  If it is non-nil, then the loop should break.
	for i := range items {
		if errs.Failure() != nil {
			// if failFast == true errs.AddRecoverable() was called,
			// we'll catch the error here.
			break
		}

		if err := getIthItem(i); err != nil {
			errs.AddRecoverable(err)
		}
	}
}

// ExampleBus_Failure describes retrieving the non-recoverable error.
func ExampleBus_Failure() {
	errs := fault.New(false)
	errs.Fail(errors.New("catastrophe"))

	// Failure() returns the primary failure.
	err := errs.Failure()
	fmt.Println(err)

	// if multiple Failures occur, each one after the first gets
	// added to the Recoverable slice as an overflow measure.
	errs.Fail(errors.New("another catastrophe"))
	errSl := errs.Recovered()

	for _, e := range errSl {
		fmt.Println(e)
	}

	// If Failure() is nil, then you can assume the operation completed.
	// A complete operation is not necessarily an error-free operation.
	// Recoverable errors may still have been added using AddRecoverable(err).
	// Make sure you check both.

	// If failFast is set to true, then the first recoerable error Added gets
	// promoted to the Err() position.
	errs = fault.New(true)
	errs.AddRecoverable(errors.New("not catastrophic, but still becomes the Failure()"))
	err = errs.Failure()
	fmt.Println(err)

	// Output: catastrophe
	// another catastrophe
	// not catastrophic, but still becomes the Failure()
}

// ExampleBus_Recovered describes the errors that processing was able to
// recover from and continue.
func ExampleErrors_Recovered() {
	errs := fault.New(false)
	errs.AddRecoverable(errors.New("not catastrophic"))
	errs.AddRecoverable(errors.New("something unwanted"))

	// Recovered() gets the slice of all recoverable errors added during
	// the run, but which did not cause a failure.
	//
	// Recovered() should never be investigated during lower level processing.
	// Implementation only ever needs to check Failure().  If an error didn't
	// promote to the Failure slot, then it should be ignored.
	//
	// The end user, at the conclusion of an operation, is the intended recipient
	// of the Recovered error slice.  After returning to the interface layer
	// (the CLI or SDK), it's the job of the end user at that location to
	// iterate through those errors and record them as wanted.
	errSl := errs.Recovered()
	for _, err := range errSl {
		fmt.Println(err)
	}

	// One or more errors in errs.Recovered() does not necessarily mean the
	// process failed.  You can have non-zero Recovered() but a nil Failure().
	if errs.Failure() == nil {
		fmt.Println("Failure() is nil")
	}

	// Inversely, if Recovered() is nil, then you can assume that no recoverable
	// or iteration-based errors occurred.  But that does not necessarily
	// mean the operation was able to complete.
	//
	// Even if Recovered() contains zero items, Err() can be non-nil.
	// Make sure you check both.

	// Output: not catastrophic
	// something unwanted
	// Failure() is nil
}

func ExampleBus_Local() {
	// It is common for Corso to run operations in parallel,
	// and for iterations to be nested within iterations.  To
	// avoid mistakenly returning an error that was sourced
	// from some other async iteration, recoverable instances
	// are aggrgated into a Local.
	errs := fault.New(false)
	el := errs.Local()

	err := func() error {
		for i := range items {
			if el.Failure() != nil {
				break
			}

			if err := getIthItem(i); err != nil {
				// instead of calling errs.AddRecoverable(err), we call the
				// local bus's Add method.  The error will still get
				// added to the errs.Recovered() set.  But if this err
				// causes the run to fail, only this local bus treats
				// it as the causal failure.
				el.AddRecoverable(err)
			}
		}

		return el.Failure()
	}()
	if err != nil {
		// handle the Failure() that appeared in the local bus.
		fmt.Println("failure occurred", errs.Failure())
	}
}

// ExampleE2e showcases a more complex integration.
func Example_e2e() {
	oper := newOperation()

	// imagine that we're a user, calling into corso SDK.
	// (fake funcs used here to minimize example bloat)
	//
	// The operation is our controller, we expect it to
	// generate a new fault.Bus when constructed, and
	// to return that struct when we call Run()
	errs := oper.Run()

	// Let's investigate what went on inside.  Since we're at
	// the top of our controller, and returning a fault.Bus,
	// all the error handlers set the Fail() case.

	/* Run() */
	func() *fault.Bus {
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
	err := func(data []any, errs *fault.Bus) error {
		// this is downstream in our code somewhere
		storer := func(a any) error {
			if err := dependencyCall(); err != nil {
				// we're not passing in or calling fault.Bus here,
				// because this isn't the iteration handler, it's just
				// a regular error.
				return errors.Wrap(err, "dependency")
			}

			return nil
		}

		el := errs.Local()

		for _, d := range data {
			if el.Failure() != nil {
				break
			}

			if err := storer(d); err != nil {
				// Since we're at the top of the iteration, we need
				// to add each error to the fault.localBus struct.
				el.AddRecoverable(err)
			}
		}

		// at the end of the func, we need to return local.Failure()
		// just in case the local bus promoted an error to the failure
		// position.  If we don't return it like normal error handling,
		// then we'll lose scope of that error.
		return el.Failure()
	}(nil, nil)
	if err != nil {
		fmt.Println("errored", err)
	}

	// At the end of the oper.Run, when returning to the interface
	// layer, we investigate the results.
	if errs.Failure() != nil {
		// handle the primary error
		fmt.Println("err occurred", errs.Failure())
	}

	for _, err := range errs.Recovered() {
		// handle each recoverable error
		fmt.Println("recoverable err occurred", err)
	}
}

// ExampleBus_Failure_return showcases when to return an error or
// nil vs errs.Failure() vs *fault.Bus
func ExampleErrors_Failure_return() {
	// The general rule of thumb is stick to standard golang error
	// handling whenever possible.
	fn := func() error {
		if err := dependency.do(); err != nil {
			return errors.Wrap(err, "direct")
		}

		return nil
	}
	if err := fn(); err != nil {
		fmt.Println(err)
	}

	// The first exception is if you're handling recoverable errors.  Recoverable
	// error handling should create a local bus instance, and return localBus.Failure()
	// so that the immediate upstream caller can be made aware of the current failure.
	fn2 := func(todo []string, errs *fault.Bus) error {
		for range todo {
			if errs.Failure() != nil {
				return errs.Failure()
			}

			if err := dependency.do(); err != nil {
				errs.AddRecoverable(errors.Wrap(err, "recoverable"))
			}
		}

		return errs.Failure()
	}
	if err := fn2([]string{"a"}, fault.New(true)); err != nil {
		fmt.Println(err)
	}

	// The second exception is if you're returning at the interface layer.
	// In that case, you're expected to return the fault.Bus itself, so that
	// callers can review the fault data.
	operationFn := func(errs *fault.Bus) *fault.Bus {
		if _, err := getData(); err != nil {
			return errs.Fail(err)
		}

		return errs
	}

	fbus := operationFn(fault.New(true))

	if fbus.Failure() != nil {
		fmt.Println("failure", fbus.Failure())
	}

	for _, err := range fbus.Recovered() {
		fmt.Println("recovered", err)
	}

	// Output: direct: caught one
	// recoverable: caught one
}

// ExampleBus_AddSkip showcases when to use AddSkip instead of an error.
func ExampleBus_AddSkip() {
	errs := fault.New(false)

	// Some conditions cause well-known problems that we want Corso to skip
	// over, instead of error out.  An initial case is when Graph API identifies
	// a file as containing malware.  We can't download the file: it'll always
	// error.  Our only option is to skip it.
	errs.AddSkip(fault.FileSkip(
		fault.SkipMalware,
		"file-id",
		"file-name",
		map[string]any{"foo": "bar"},
	))

	// later on, after processing, end users can scrutinize the skipped items.
	fmt.Println(errs.Skipped()[0].String())

	// Output: skipped processing file: malware_detected
}
