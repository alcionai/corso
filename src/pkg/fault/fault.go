package fault

import (
	"sync"

	"golang.org/x/exp/slices"
)

type errors struct {
	mu *sync.Mutex

	// err identifies non-recoverable errors.  This includes
	// non-start cases (ex: cannot connect to client), hard-
	// stop issues (ex: credentials expired) or conscious exit
	// cases (ex: iteration error + failFast config).
	err error

	// errs is the accumulation of recoverable or iterated
	// errors.  Eg: if a process is retrieving N items, and
	// 1 of the items fails to be retrieved, but the rest of
	// them succeed, we'd expect to see 1 error added to this
	// slice.
	errs []error

	// if failFast is true, the first errs addition will
	// get promoted to the err value.  This signifies a
	// non-recoverable processing state, causing any running
	// processes to exit.
	failFast bool
}

// ErrorsData provides the errors data alone, without sync
// controls, allowing the data to be persisted.
type ErrorsData struct {
	Err      error   `json:"err"`
	Errs     []error `json:"errs"`
	FailFast bool    `json:"failFast"`
}

// New constructs a new error with default values in place.
func New(failFast bool) *errors {
	return &errors{
		mu:       &sync.Mutex{},
		errs:     []error{},
		failFast: failFast,
	}
}

// Err returns the primary error.  If not nil, this
// indicates the operation exited prior to completion.
func (e *errors) Err() error {
	return e.err
}

// Errs returns the slice of recoverable and
// iterated errors.
func (e *errors) Errs() []error {
	return e.errs
}

// Data returns the plain set of error data
// without any sync properties.
func (e *errors) Data() ErrorsData {
	return ErrorsData{
		Err:      e.err,
		Errs:     slices.Clone(e.errs),
		FailFast: e.failFast,
	}
}

// TODO: introduce Failer interface

// Fail sets the non-recoverable error (ie: errors.err)
// in the errors struct.  If a non-recoverable error is
// already present, the error gets added to the errs slice.
func (e *errors) Fail(err error) *errors {
	if err == nil {
		return e
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	return e.setErr(err)
}

// setErr handles setting errors.err.  Sync locking gets
// handled upstream of this call.
func (e *errors) setErr(err error) *errors {
	if e.err != nil {
		return e.addErr(err)
	}

	e.err = err

	return e
}

// TODO: introduce Adder interface

// Add appends the error to the slice of recoverable and
// iterated errors (ie: errors.errs).  If failFast is true,
// the first Added error will get copied to errors.err,
// causing the errors struct to identify as non-recoverably
// failed.
func (e *errors) Add(err error) *errors {
	if err == nil {
		return e
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	return e.addErr(err)
}

// addErr handles adding errors to errors.errs.  Sync locking
// gets handled upstream of this call.
func (e *errors) addErr(err error) *errors {
	if e.err == nil && e.failFast {
		e.setErr(err)
	}

	e.errs = append(e.errs, err)

	return e
}
