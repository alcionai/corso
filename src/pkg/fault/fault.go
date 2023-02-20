package fault

import (
	"sync"

	"golang.org/x/exp/slices"
)

type Bus struct {
	mu *sync.Mutex

	// failure identifies non-recoverable errors.  This includes
	// non-start cases (ex: cannot connect to client), hard-
	// stop issues (ex: credentials expired) or conscious exit
	// cases (ex: iteration error + failFast config).
	failure error

	// recoverable is the accumulation of recoverable errors
	// Eg: if a process is retrieving N items, and 1 of the
	// items fails to be retrieved, but the rest of them succeed,
	// we'd expect to see 1 error added to this slice.
	recoverable []error

	// if failFast is true, the first errs addition will
	// get promoted to the err value.  This signifies a
	// non-recoverable processing state, causing any running
	// processes to exit.
	failFast bool
}

// Errors provides the errors data alone, without sync
// controls, allowing the data to be persisted.
type Errors struct {
	Failure   error   `json:"failure"`
	Recovered []error `json:"recovered"`
	FailFast  bool    `json:"failFast"`

	// legacy support
	Err  error   `json:"err"`
	Errs []error `json:"errs"`
}

// New constructs a new error with default values in place.
func New(failFast bool) *Bus {
	return &Bus{
		mu:          &sync.Mutex{},
		recoverable: []error{},
		failFast:    failFast,
	}
}

// Failure returns the primary error.  If not nil, this
// indicates the operation exited prior to completion.
func (e *Bus) Failure() error {
	return e.failure
}

// Recovered returns the slice of errors that occurred in
// recoverable points of processing.  This is often during
// iteration where a single failure (ex: retrieving an item),
// doesn't require the entire process to end.
func (e *Bus) Recovered() []error {
	return e.recoverable
}

// Data returns the plain record of errors that were aggregated
// within a fult Bus.
func (e *Bus) Data() Errors {
	return Errors{
		Failure:   e.failure,
		Recovered: slices.Clone(e.recoverable),
		FailFast:  e.failFast,
	}
}

// Fail sets the non-recoverable error (ie: bus.failure)
// in thebus.  If a failure error is already present,
// the error gets added to the recoverable slice for
// purposes of tracking.
//
// TODO: Return Data, not Bus.  The consumers of a failure
// should care about the state of data, not the communication
// pattern.
func (e *Bus) Fail(err error) *Bus {
	if err == nil {
		return e
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	return e.setFailure(err)
}

// setErr handles setting bus.failure.  Sync locking gets
// handled upstream of this call.
func (e *Bus) setFailure(err error) *Bus {
	if e.failure == nil {
		e.failure = err
		return e
	}

	// technically not a recoverable error: we're using the
	// recoverable slice as an overflow container here to
	// ensure everything is tracked.
	e.recoverable = append(e.recoverable, err)

	return e
}

// AddRecoverable appends the error to the slice of recoverable
// errors (ie: bus.recoverable).  If failFast is true, the first
// added error will get copied to bus.failure, causing the bus
// to identify as non-recoverably failed.
//
// TODO: nil return, not Bus, since we don't want people to return
// from errors.AddRecoverable().
func (e *Bus) AddRecoverable(err error) *Bus {
	if err == nil {
		return e
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	return e.addRecoverableErr(err)
}

// addErr handles adding errors to errors.errs.  Sync locking
// gets handled upstream of this call.
func (e *Bus) addRecoverableErr(err error) *Bus {
	if e.failure == nil && e.failFast {
		e.setFailure(err)
	}

	e.recoverable = append(e.recoverable, err)

	return e
}

// ---------------------------------------------------------------------------
// Local aggregator
// ---------------------------------------------------------------------------

// Local constructs a new local bus to handle error aggregation in a
// constrained scope.  Local busses shouldn't be passed down  to other
// funcs, and the function that spawned the local bus should always
// return `local.Failure()` to ensure that hard failures are propagated
// back upstream.
func (e *Bus) Local() *localBus {
	return &localBus{
		mu:  &sync.Mutex{},
		bus: e,
	}
}

type localBus struct {
	mu      *sync.Mutex
	bus     *Bus
	current error
}

func (e *localBus) AddRecoverable(err error) {
	if err == nil {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.bus.Failure() == nil && e.bus.failFast {
		e.current = err
	}

	e.bus.AddRecoverable(err)
}

// Failure returns the failure that happened within the local bus.
// It does not return the underlying bus.Failure(), only the failure
// that was recorded within the local bus instance.  This error should
// get returned by any func which created a local bus.
func (e *localBus) Failure() error {
	return e.current
}
