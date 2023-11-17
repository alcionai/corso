package fault

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/pkg/logger"
)

// temporary hack identifier
// see: https://github.com/alcionai/corso/pull/2510#discussion_r1113532530
// TODO: https://github.com/alcionai/corso/issues/4003
const LabelForceNoBackupCreation = "label_forces_no_backup_creations"

type Bus struct {
	mu *sync.Mutex

	// When creating a local bus, the parent property retains a pointer
	// to the root Bus.  Even in the case of multiple chained creations of
	// local busses, the parent reference remains the original root bus,
	// and does not create a linked list of lineage.  Any errors and failures
	// created by a local instance will get fielded to the parent.  But only
	// local errors will returned by property getter funcs.
	parent *Bus

	// Failure probably identifies errors that were added to the bus
	// or a local Bus via AddRecoverable, but which were promoted
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

	// alerts contain purely informational messages and data.  They
	// represent situations where the end user should be aware of some
	// occurrence that is not an error, exception, skipped data, or
	// other runtime/persistence impacting issue.
	alerts []Alert

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

// Local constructs a new bus with a local reference to handle error aggregation
// in a constrained scope.  This allows the caller to review recoverable errors and
// failures within only the current codespace, as opposed to the global set of errors.
// The function that spawned the local bus should always return `bus.Failure()` to
// ensure that hard failures are propagated back upstream.
func (e *Bus) Local() *Bus {
	parent := e.parent

	// only use e if it is already the root instance
	if parent == nil {
		parent = e
	}

	return &Bus{
		mu:       &sync.Mutex{},
		parent:   parent,
		failFast: parent.failFast,
	}
}

// FailFast returns the failFast flag in the bus.
func (e *Bus) FailFast() bool {
	return e.failFast
}

// Fail sets the non-recoverable error (ie: bus.failure)
// in the bus.  If a failure error is already present,
// the error gets added to the recoverable slice for
// purposes of tracking.
func (e *Bus) Fail(err error) *Bus {
	if err == nil {
		return e
	}

	return e.setFailure(err)
}

// setErr handles setting bus.failure.  Sync locking gets
// handled upstream of this call.
func (e *Bus) setFailure(err error) *Bus {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.failure == nil {
		e.failure = err
	} else {
		// technically not a recoverable error: we're using the
		// recoverable slice as an overflow container here to
		// ensure everything is tracked.
		e.recoverable = append(e.recoverable, err)
	}

	if e.parent != nil {
		e.parent.setFailure(err)
	}

	return e
}

// AddRecoverable appends the error to the slice of recoverable
// errors (ie: bus.recoverable).  If failFast is true, the first
// added error will get copied to bus.failure, causing the bus
// to identify as non-recoverably failed.
func (e *Bus) AddRecoverable(ctx context.Context, err error) {
	if err == nil {
		return
	}

	e.logAndAddRecoverable(ctx, err, 1)
}

// logs the error and adds it to the bus.  If the error is a failure,
// it gets logged at an Error level.  Otherwise logs an Info.
func (e *Bus) logAndAddRecoverable(ctx context.Context, err error, skip int) {
	log := logger.CtxErrStack(ctx, err, skip+1)
	isFail := e.addRecoverableErr(err)

	if isFail {
		log.Errorf("failed on recoverable error: %v", err)
	} else {
		log.Infof("recoverable error: %v", err)
	}
}

// addErr handles adding errors to errors.errs.  Sync locking
// gets handled upstream of this call.  Returns true if the
// error is a failure, false otherwise.
func (e *Bus) addRecoverableErr(err error) bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	var isFail bool

	if e.failure == nil && e.failFast {
		if e.failure == nil {
			e.failure = err
		} else {
			// technically not a recoverable error: we're using the
			// recoverable slice as an overflow container here to
			// ensure everything is tracked.
			e.recoverable = append(e.recoverable, err)
		}

		if e.parent != nil {
			e.parent.setFailure(err)
		}

		isFail = true
	}

	e.recoverable = append(e.recoverable, err)

	// local bus instances must promote errors to the root bus.
	if e.parent != nil {
		e.parent.addRecoverableErr(err)
	}

	return isFail
}

// ---------------------------------------------------------------------------
// Non-error adders
// ---------------------------------------------------------------------------

// AddAlert appends a record of an Alert message to the fault bus.
// Importantly, alerts are not errors, exceptions, or skipped items.
// An alert should only be generated if no other fault functionality
// is in use, but that we still want the end user to clearly and
// plainly receive a notification about a runtime event.
func (e *Bus) AddAlert(ctx context.Context, a *Alert) {
	if a == nil {
		return
	}

	e.logAndAddAlert(ctx, a, 1)
}

func (e *Bus) logAndAddAlert(ctx context.Context, a *Alert, trace int) {
	logger.CtxStack(ctx, trace+1).
		With("alert", a).
		Info("alert: " + a.Message)
	e.addAlert(a)
}

func (e *Bus) addAlert(a *Alert) *Bus {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.alerts = append(e.alerts, *a)

	// local bus instances must promote alerts to the root bus.
	if e.parent != nil {
		e.parent.addAlert(a)
	}

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
func (e *Bus) AddSkip(ctx context.Context, s *Skipped) {
	if s == nil {
		return
	}

	e.logAndAddSkip(ctx, s, 1)
}

// logs the error and adds a skipped item.
func (e *Bus) logAndAddSkip(ctx context.Context, s *Skipped, trace int) {
	logger.CtxStack(ctx, trace+1).
		With("skipped", s).
		Info("skipped an item")
	e.addSkip(s)
}

func (e *Bus) addSkip(s *Skipped) *Bus {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.skipped = append(e.skipped, *s)

	// local bus instances must promote skipped items to the root bus.
	if e.parent != nil {
		e.parent.addSkip(s)
	}

	return e
}

// ---------------------------------------------------------------------------
// Results
// ---------------------------------------------------------------------------

// Errors returns the plain record of errors that were aggregated
// within a fult Bus.
func (e *Bus) Errors() *Errors {
	items, nonItems := itemsIn(e.failure, e.recoverable)

	return &Errors{
		Failure:   clues.ToCore(e.failure),
		Recovered: nonItems,
		Items:     items,
		Skipped:   slices.Clone(e.skipped),
		Alerts:    slices.Clone(e.alerts),
		FailFast:  e.failFast,
	}
}

// Failure returns the primary error.  If not nil, this
// indicates the operation exited prior to completion.
// If the bus is a local instance, this only returns the
// local failure, and will not return parent data.
func (e *Bus) Failure() error {
	return e.failure
}

// Recovered returns the slice of errors that occurred in
// recoverable points of processing.  This is often during
// iteration where a single failure (ex: retrieving an item),
// doesn't require the entire process to end.
// If the bus is a local instance, this only returns the
// local recovered errors, and will not return parent data.
func (e *Bus) Recovered() []error {
	return slices.Clone(e.recoverable)
}

// Skipped returns the slice of items that were permanently
// skipped during processing.
// If the bus is a local instance, this only returns the
// local skipped items, and will not return parent data.
func (e *Bus) Skipped() []Skipped {
	return slices.Clone(e.skipped)
}

// Alerts returns the slice of alerts generated during runtime.
// If the bus is a local alerts, this only returns the
// local failure, and will not return parent data.
func (e *Bus) Alerts() []Alert {
	return slices.Clone(e.alerts)
}

// ItemsAndRecovered returns the items that failed along with other
// recoverable errors
func (e *Bus) ItemsAndRecovered() ([]Item, []error) {
	var (
		is  = map[string]Item{}
		non = []error{}
	)

	for _, err := range e.recoverable {
		var ie *Item
		if !errors.As(err, &ie) {
			non = append(non, err)
			continue
		}

		is[ie.dedupeID()] = *ie
	}

	var ie *Item
	if errors.As(e.failure, &ie) {
		is[ie.dedupeID()] = *ie
	}

	return maps.Values(is), non
}

// Errors provides the errors data alone, without sync controls
// or adders/setters.  Expected to get called at the end of processing,
// as a way to aggregate results.
type Errors struct {
	// Failure identifies a non-recoverable error.  This includes
	// non-start cases (ex: cannot connect to client), hard-
	// stop issues (ex: credentials expired) or conscious exit
	// cases (ex: iteration error + failFast config).
	Failure *clues.ErrCore `json:"failure"`

	// Recovered is the set of NON-Item errors that accumulated
	// through a runtime under best-effort processing conditions.
	// They imply that an error occurred, but the process was able
	// to move on and complete afterwards.  Any error that can be
	// serialized to a fault.Item is found in the Items set instead.
	Recovered []*clues.ErrCore `json:"recovered"`

	// Items are the reduction of all errors (both the failure and the
	// recovered values) in the Errors struct into a slice of items,
	// deduplicated by their Namespace + ID.
	Items []Item `json:"items"`

	// Skipped is the accumulation of skipped items.  Skipped items
	// are not errors themselves, but instead represent some permanent
	// inability to process an item, due to a well-known cause.
	Skipped []Skipped `json:"skipped"`

	// Alerts contain purely informational messages and data.  They
	// represent situations where the end user should be aware of some
	// occurrence that is not an error, exception, skipped data, or
	// other runtime/persistence impacting issue.
	Alerts []Alert

	// If FailFast is true, then the first Recoverable error will
	// promote to the Failure spot, causing processing to exit.
	FailFast bool `json:"failFast"`
}

// itemsIn reduces all errors (both the failure and recovered values)
// in the Errors struct into a slice of items, deduplicated by their
// Namespace + ID.
// Any non-item error is serialized to a clues.ErrCore and returned in
// the second list.
func itemsIn(failure error, recovered []error) ([]Item, []*clues.ErrCore) {
	var (
		is  = map[string]Item{}
		non = []*clues.ErrCore{}
	)

	for _, err := range recovered {
		var ie *Item
		if !errors.As(err, &ie) {
			non = append(non, clues.ToCore(err))
			continue
		}

		is[ie.dedupeID()] = *ie
	}

	var ie *Item
	if errors.As(failure, &ie) {
		is[ie.dedupeID()] = *ie
	}

	return maps.Values(is), non
}

// Marshal runs json.Marshal on the errors.
func (e *Errors) Marshal() ([]byte, error) {
	bs, err := json.Marshal(e)
	return bs, err
}

// UnmarshalErrorsTo produces a func that complies with the unmarshaller
// type in streamStore.
func UnmarshalErrorsTo(e *Errors) func(io.ReadCloser) error {
	return func(rc io.ReadCloser) error {
		return json.NewDecoder(rc).Decode(e)
	}
}

// ---------------------------------------------------------------------------
// Print compatibility
// ---------------------------------------------------------------------------

// Print writes the DetailModel Entries to StdOut, in the format
// requested by the caller.
func (e *Errors) PrintItems(
	ctx context.Context,
	ignoreAlerts, ignoreErrors, ignoreSkips, ignoreRecovered bool,
) {
	if len(e.Alerts)+len(e.Items)+len(e.Skipped)+len(e.Recovered) == 0 ||
		(ignoreAlerts && ignoreErrors && ignoreSkips && ignoreRecovered) {
		return
	}

	sl := make([]print.Printable, 0)

	if !ignoreAlerts {
		for _, a := range e.Alerts {
			sl = append(sl, print.Printable(a))
		}
	}

	if !ignoreSkips {
		for _, s := range e.Skipped {
			sl = append(sl, print.Printable(s))
		}
	}

	if !ignoreErrors {
		for _, i := range e.Items {
			sl = append(sl, print.Printable(i))
		}
	}

	if !ignoreRecovered {
		for _, rcv := range e.Recovered {
			pec := errCoreToPrintable(rcv)
			sl = append(sl, print.Printable(&pec))
		}
	}

	print.All(ctx, sl...)
}

var _ print.Printable = &printableErrCore{}

type printableErrCore struct {
	*clues.ErrCore
}

func errCoreToPrintable(ec *clues.ErrCore) printableErrCore {
	if ec == nil {
		return printableErrCore{ErrCore: &clues.ErrCore{Msg: "<nil>"}}
	}

	return printableErrCore{ErrCore: ec}
}

func (pec printableErrCore) MinimumPrintable() any {
	return pec
}

func (pec printableErrCore) Headers(bool) []string {
	// NOTE: skipID does not make sense in this context
	return []string{"Error"}
}

func (pec printableErrCore) Values(bool) []string {
	if pec.ErrCore == nil {
		return []string{"<nil>"}
	}

	return []string{pec.Msg}
}
