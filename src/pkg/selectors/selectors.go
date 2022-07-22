package selectors

import (
	"strings"

	"github.com/pkg/errors"
)

type service int

//go:generate stringer -type=service -linecomment
const (
	ServiceUnknown  service = iota // Unknown Service
	ServiceExchange                // Exchange
)

var ErrorBadSelectorCast = errors.New("wrong selector service type")

const (
	scopeKeyCategory    = "category"
	scopeKeyGranularity = "granularity"
	scopeKeyInfoFilter  = "info_filter"
)

// The granularity exprerssed by the scope.  Groups imply non-item granularity,
// such as a directory.  Items are individual files or objects.
// Filters are properties that search over service-specific info
const (
	Group  = "group"
	Item   = "item"
	Filter = "filter"
)

const (
	// AnyTgt is the target value used to select "any data of <type>"
	// Ex: {user: u1, events: AnyTgt) => all events for user u1.
	// In the event that "*" conflicts with a user value, such as a
	// folder named "*", calls to corso should escape the value with "\*"
	AnyTgt = "*"
	// NoneTgt is the target value used to select "no data of <type>"
	// This is primarily a fallback for empty values.  Adding NoneTgt or
	// None() to any selector will force all matches() checks on that
	// selector to fail.
	// Ex: {user: u1, events: NoneTgt} => matches nothing.
	NoneTgt = ""

	delimiter = ","
)

// ---------------------------------------------------------------------------
// Selector
// ---------------------------------------------------------------------------

// The core selector.  Has no api for setting or retrieving data.
// Is only used to pass along more specific selector instances.
type Selector struct {
	Service  service             `json:"service,omitempty"`    // The service scope of the data.  Exchange, Teams, Sharepoint, etc.
	Excludes []map[string]string `json:"exclusions,omitempty"` // A slice of exclusions.  Each exclusion applies to all inclusions.
	Includes []map[string]string `json:"scopes,omitempty"`     // A slice of inclusions.  Expected to get cast to a service wrapper within each service handler.
}

// helper for specific selector instance constructors.
func newSelector(s service) Selector {
	return Selector{
		Service:  s,
		Excludes: []map[string]string{},
		Includes: []map[string]string{},
	}
}

// Any returns the set matching any value.
func Any() []string {
	return []string{AnyTgt}
}

// None returns the set matching None of the values.
// This is primarily a fallback for empty values.  Adding None()
// to any selector will force all matches() checks on that selector
// to fail.
func None() []string {
	return []string{NoneTgt}
}

// ---------------------------------------------------------------------------
// Destination
// ---------------------------------------------------------------------------

type Destination map[string]string

var ErrorDestinationAlreadySet = errors.New("destination is already declared")

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func badCastErr(cast, is service) error {
	return errors.Wrapf(ErrorBadSelectorCast, "%s service is not %s", cast, is)
}

func existingDestinationErr(category, is string) error {
	return errors.Wrapf(ErrorDestinationAlreadySet, "%s destination already set to %s", category, is)
}

func join(s ...string) string {
	return strings.Join(s, delimiter)
}

func split(s string) []string {
	return strings.Split(s, delimiter)
}

// if the provided slice contains Any, returns [Any]
// if the slice contains None, returns [None]
// if the slice contains Any and None, returns the first
// if the slice is empty, returns [None]
// otherwise returns the input unchanged
func normalize(s []string) []string {
	if len(s) == 0 {
		return None()
	}
	for _, e := range s {
		if e == AnyTgt {
			return Any()
		}
		if e == NoneTgt {
			return None()
		}
	}
	return s
}
