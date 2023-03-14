package fault

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/print"
)

type Bus struct {
	mu *sync.Mutex

	// Failure probably identifies errors that were added to the bus
	// or localBus via AddRecoverable, but which were promoted
	// to the failure position due to failFast=true configuration.
	// Alternatively, the process controller might have set failure
	// by calling Fail(err).
	failure error

	// recoverable is the accumulation of recoverable errors.
	// Eg: if a process is retrieving N items, and 1 of the
	// items fails to be retrieved, but the rest of them succeed,
	// we'd expect to see 1 error added to this slice.
	recoverable []error

	// skipped is the accumulation of skipped items.  Skipped items
	// are not errors themselves, but instead represent some permanent
	// inability to process an item, due to a well-known cause.
	skipped []Skipped

	// if failFast is true, the first errs addition will
	// get promoted to the err value.  This signifies a
	// non-recoverable processing state, causing any running
	// processes to exit.
	failFast bool
}

// New constructs a new error with default values in place.
func New(failFast bool) *Bus {
	return &Bus{
		mu:          &sync.Mutex{},
		recoverable: []error{},
		failFast:    failFast,
	}
}

// FailFast returs the failFast flag in the bus.
func (e *Bus) FailFast() bool {
	return e.failFast
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
	return slices.Clone(e.recoverable)
}

// Skipped returns the slice of items that were permanently
// skipped during processing.
func (e *Bus) Skipped() []Skipped {
	return slices.Clone(e.skipped)
}

// Fail sets the non-recoverable error (ie: bus.failure)
// in the bus.  If a failure error is already present,
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

// AddSkip appends a record of a Skipped item to the fault bus.
// Importantly, skipped items are not the same as recoverable
// errors.  An item should only be skipped under the following
// conditions.  All other cases should be handled as errors.
// 1. The conditions for skipping the item are well-known and
// well-documented.  End users need to be able to understand
// both the conditions and identifications of skips.
// 2. Skipping avoids a permanent and consistent failure.  If
// the underlying reason is transient or otherwise recoverable,
// the item should not be skipped.
func (e *Bus) AddSkip(s *Skipped) *Bus {
	if s == nil {
		return e
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	return e.addSkip(s)
}

func (e *Bus) addSkip(s *Skipped) *Bus {
	e.skipped = append(e.skipped, *s)
	return e
}

// Errors returns the plain record of errors that were aggregated
// within a fult Bus.
func (e *Bus) Errors() Errors {
	return Errors{
		Failure:   e.failure,
		Recovered: slices.Clone(e.recoverable),
		Items:     itemsIn(e.failure, e.recoverable),
		Skipped:   slices.Clone(e.skipped),
		FailFast:  e.failFast,
	}
}

// ---------------------------------------------------------------------------
// Errors Data
// ---------------------------------------------------------------------------

// Errors provides the errors data alone, without sync controls
// or adders/setters.  Expected to get called at the end of processing,
// as a way to aggregate results.
type Errors struct {
	// Failure identifies a non-recoverable error.  This includes
	// non-start cases (ex: cannot connect to client), hard-
	// stop issues (ex: credentials expired) or conscious exit
	// cases (ex: iteration error + failFast config).
	Failure error `json:"failure"`

	// Recovered errors accumulate through a runtime under
	// best-effort processing conditions.  They imply that an
	// error occurred, but the process was able to move on and
	// complete afterwards.
	// Eg: if a process is retrieving N items, and 1 of the
	// items fails to be retrieved, but the rest of them succeed,
	// we'd expect to see 1 error added to this slice.
	Recovered []error `json:"-"`

	// Items are the reduction of all errors (both the failure and the
	// recovered values) in the Errors struct into a slice of items,
	// deduplicated by their ID.
	Items []Item `json:"items"`

	// Skipped is the accumulation of skipped items.  Skipped items
	// are not errors themselves, but instead represent some permanent
	// inability to process an item, due to a well-known cause.
	Skipped []Skipped `json:"skipped"`

	// If FailFast is true, then the first Recoverable error will
	// promote to the Failure spot, causing processing to exit.
	FailFast bool `json:"failFast"`
}

// itemsIn reduces all errors (both the failure and recovered values)
// in the Errors struct into a slice of items, deduplicated by their
// ID.
func itemsIn(failure error, recovered []error) []Item {
	is := map[string]Item{}

	for _, err := range recovered {
		var ie *Item
		if !errors.As(err, &ie) {
			continue
		}

		is[ie.ID] = *ie
	}

	var ie *Item
	if errors.As(failure, &ie) {
		is[ie.ID] = *ie
	}

	return maps.Values(is)
}

// Marshal runs json.Marshal on the errors.
func (e Errors) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalErrorsTo produces a func that complies with the unmarshaller
// type in streamStore.
func UnmarshalErrorsTo(e *Errors) func(io.ReadCloser) error {
	return func(rc io.ReadCloser) error {
		return json.NewDecoder(rc).Decode(e)
	}
}

// Print writes the DetailModel Entries to StdOut, in the format
// requested by the caller.
func (e Errors) PrintItems(ctx context.Context) {
	count := len(e.Items) + len(e.Skipped)
	if count == 0 {
		return
	}

	sl := make([]print.Printable, 0, count)

	for _, s := range e.Skipped {
		sl = append(sl, print.Printable(s))
	}

	for _, i := range e.Items {
		sl = append(sl, print.Printable(i))
	}

	print.All(ctx, sl...)
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

	if e.current == nil && e.bus.failFast {
		e.current = err
	}

	e.bus.AddRecoverable(err)
}

// AddSkip appends a record of a Skipped item to the local bus.
// Importantly, skipped items are not the same as recoverable
// errors.  An item should only be skipped under the following
// conditions.  All other cases should be handled as errors.
// 1. The conditions for skipping the item are well-known and
// well-documented.  End users need to be able to understand
// both the conditions and identifications of skips.
// 2. Skipping avoids a permanent and consistent failure.  If
// the underlying reason is transient or otherwise recoverable,
// the item should not be skipped.
func (e *localBus) AddSkip(s *Skipped) {
	if s == nil {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.bus.AddSkip(s)
}

// Failure returns the failure that happened within the local bus.
// It does not return the underlying bus.Failure(), only the failure
// that was recorded within the local bus instance.  This error should
// get returned by any func which created a local bus.
func (e *localBus) Failure() error {
	return e.current
}

// temporary hack identifier
// see: https://github.com/alcionai/corso/pull/2510#discussion_r1113532530
const LabelForceNoBackupCreation = "label_forces_no_backup_creations"
