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
	scopeKeyCategory = "category"
)

const (
	// All is the wildcard value used to express "all data of <type>"
	// Ex: {user: u1, events: All) => all events for user u1.
	All = "ß∂ƒ∑´®≈ç√¬˜"
	// None is usesd to express "no data of <type>"
	// Ex: {user: u1, events: None} => no events for user u1.
	None = "√ç≈œ´∆¬˚¨π"

	delimiter = ","
)

// ---------------------------------------------------------------------------
// Selector
// ---------------------------------------------------------------------------

// The core selector.  Has no api for setting or retrieving data.
// Is only used to pass along more specific selector instances.
type Selector struct {
	RestorePointID string              `json:"restorePointID,omitempty"` // A restore point id, used only by restore operations.
	Service        service             `json:"service,omitempty"`        // The service scope of the data.  Exchange, Teams, Sharepoint, etc.
	Excludes       []map[string]string `json:"exclusions,omitempty"`     // A slice of exclusions.  Each exclusion applies to all inclusions.
	Includes       []map[string]string `json:"scopes,omitempty"`         // A slice of inclusions.  Expected to get cast to a service wrapper within each service handler.
}

// helper for specific selector instance constructors.
func newSelector(s service, restorePointID string) Selector {
	return Selector{
		RestorePointID: restorePointID,
		Service:        s,
		Excludes:       []map[string]string{},
		Includes:       []map[string]string{},
	}
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
