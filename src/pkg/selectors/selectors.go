package selectors

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

type service int

//go:generate stringer -type=service -linecomment
const (
	ServiceUnknown    service = iota // Unknown Service
	ServiceExchange                  // Exchange
	ServiceOneDrive                  // OneDrive
	ServiceSharePoint                // SharePoint
)

var serviceToPathType = map[service]path.ServiceType{
	ServiceUnknown:    path.UnknownService,
	ServiceExchange:   path.ExchangeService,
	ServiceOneDrive:   path.OneDriveService,
	ServiceSharePoint: path.SharePointService,
}

var (
	ErrorBadSelectorCast = errors.New("wrong selector service type")
	ErrorNoMatchingItems = errors.New("no items match the specified selectors")
)

const (
	scopeKeyCategory   = "category"
	scopeKeyInfoFilter = "info_filter"
	scopeKeyDataType   = "type"
)

// The granularity exprerssed by the scope.  Groups imply non-item granularity,
// such as a directory.  Items are individual files or objects.
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
)

var (
	delimiter = fmt.Sprint(0x1F)
	passAny   = filters.Pass()
	failAny   = filters.Fail()
)

// All is the resource name that gets output when the resource is AnyTgt.
// It is not used aside from printing resources.
const All = "All"

type Reducer interface {
	Reduce(context.Context, *details.Details) *details.Details
}

// selectorResourceOwners aggregates all discrete resource owner ids described
// in the selector.  Any and None values are ignored.  ResourceOwner sets are
// grouped by their scope type (includes, excludes, filters).
type selectorResourceOwners struct {
	Includes []string
	Excludes []string
	Filters  []string
}

type resourceOwnerer interface {
	ResourceOwners() selectorResourceOwners
}

// ---------------------------------------------------------------------------
// Selector
// ---------------------------------------------------------------------------

// The core selector.  Has no api for setting or retrieving data.
// Is only used to pass along more specific selector instances.
type Selector struct {
	// The service scope of the data.  Exchange, Teams, SharePoint, etc.
	Service service `json:"service,omitempty"`
	// A slice of exclusion scopes.  Exclusions apply globally to all
	// inclusions/filters, with any-match behavior.
	Excludes []scope `json:"exclusions,omitempty"`
	// A slice of filter scopes.  All inclusions must also match ALL filters.
	Filters []scope `json:"filters,omitempty"`
	// A slice of inclusion scopes.  Comparators must match either one of these,
	// or all filters, to be included.
	Includes []scope `json:"includes,omitempty"`
}

// helper for specific selector instance constructors.
func newSelector(s service) Selector {
	return Selector{
		Service:  s,
		Excludes: []scope{},
		Includes: []scope{},
	}
}

func (s Selector) String() string {
	bs, err := json.Marshal(s)
	if err != nil {
		return "error"
	}

	return string(bs)
}

// appendScopes iterates through each scope in the list of scope slices,
// calling setDefaults() to ensure it is completely populated, and appends
// those scopes to the `to` slice.
func appendScopes[T scopeT](to []scope, scopes ...[]T) []scope {
	if len(to) == 0 {
		to = []scope{}
	}

	for _, scopeSl := range scopes {
		for _, s := range scopeSl {
			s.setDefaults()
			to = append(to, scope(s))
		}
	}

	return to
}

// scopes retrieves the list of scopes in the selector.
// future TODO: if Inclues is nil, return filters.
func scopes[T scopeT](s Selector) []T {
	scopes := []T{}

	for _, v := range s.Includes {
		scopes = append(scopes, T(v))
	}

	return scopes
}

// discreteScopes retrieves the list of scopes in the selector.
// for any scope in the `Includes` set, if scope.IsAny(rootCat),
// then that category's value is replaced with the provided set of
// discrete identifiers.
// If discreteIDs is an empty slice, returns the normal scopes(s).
// future TODO: if Includes is nil, return filters.
func discreteScopes[T scopeT, C categoryT](
	s Selector,
	rootCat C,
	discreteIDs []string,
) []T {
	sl := []T{}

	if len(discreteIDs) == 0 {
		return scopes[T](s)
	}

	for _, v := range s.Includes {
		t := T(v)

		if isAnyTarget(t, rootCat) {
			w := T{}
			for k, v := range t {
				w[k] = v
			}

			set(w, rootCat, discreteIDs)
			t = w
		}

		sl = append(sl, t)
	}

	return sl
}

// Returns the path.ServiceType matching the selector service.
func (s Selector) PathService() path.ServiceType {
	return serviceToPathType[s.Service]
}

// Reduce is a quality-of-life interpreter that allows Reduce to be called
// from the generic selector by interpreting the selector service type rather
// than have the caller make that interpretation.  Returns an error if the
// service is unsupported.
func (s Selector) Reduce(ctx context.Context, deets *details.Details) (*details.Details, error) {
	r, err := selectorAsIface[Reducer](s)
	if err != nil {
		return nil, err
	}

	return r.Reduce(ctx, deets), nil
}

func (s Selector) ResourceOwners() (selectorResourceOwners, error) {
	ro, err := selectorAsIface[resourceOwnerer](s)
	if err != nil {
		return selectorResourceOwners{}, err
	}

	return ro.ResourceOwners(), nil
}

// transformer for arbitrary selector interfaces
func selectorAsIface[T any](s Selector) (T, error) {
	var (
		a   any
		t   T
		err error
	)

	switch s.Service {
	case ServiceExchange:
		a, err = func() (any, error) { return s.ToExchangeRestore() }()
		t = a.(T)
	case ServiceOneDrive:
		a, err = func() (any, error) { return s.ToOneDriveRestore() }()
		t = a.(T)
	case ServiceSharePoint:
		a, err = func() (any, error) { return s.ToSharePointRestore() }()
		t = a.(T)
	default:
		err = errors.New("service not supported: " + s.Service.String())
	}

	return t, err
}

// ---------------------------------------------------------------------------
// Printing Selectors for Human Reading
// ---------------------------------------------------------------------------

type Printable struct {
	Service  string              `json:"service"`
	Excludes map[string][]string `json:"excludes,omitempty"`
	Filters  map[string][]string `json:"filters,omitempty"`
	Includes map[string][]string `json:"includes,omitempty"`
}

type printabler interface {
	Printable() Printable
}

// ToPrintable creates the minimized display of a selector, formatted for human readability.
func (s Selector) ToPrintable() Printable {
	p, err := selectorAsIface[printabler](s)
	if err != nil {
		return Printable{}
	}

	return p.Printable()
}

// toPrintable creates the minimized display of a selector, formatted for human readability.
func toPrintable[T scopeT](s Selector) Printable {
	return Printable{
		Service:  s.Service.String(),
		Excludes: toResourceTypeMap[T](s.Excludes),
		Filters:  toResourceTypeMap[T](s.Filters),
		Includes: toResourceTypeMap[T](s.Includes),
	}
}

// Resources generates a tabular-readable output of the resources in Printable.
// Only the first (arbitrarily picked) resource is displayed.  All others are
// simply counted.  If no inclusions exist, uses Filters.  If no filters exist,
// defaults to "None".
// Resource refers to the top-level entity in the service. User for Exchange,
// Site for sharepoint, etc.
func (p Printable) Resources() string {
	s := resourcesShortFormat(p.Includes)
	if len(s) == 0 {
		s = resourcesShortFormat(p.Filters)
	}

	if len(s) == 0 {
		s = "None"
	}

	return s
}

// returns a string with the resources in the map.  Shortened to the first resource key,
// plus, if more exist, " (len-1 more)"
func resourcesShortFormat(m map[string][]string) string {
	var s string

	for k := range m {
		s = k
		break
	}

	if len(s) > 0 && len(m) > 1 {
		s = fmt.Sprintf("%s (%d more)", s, len(m)-1)
	}

	return s
}

// Transforms the slice to a single map.
// Keys are each service's rootCat value.
// Values are the set of all scopeKeyDataTypes for the resource.
func toResourceTypeMap[T scopeT](s []scope) map[string][]string {
	if len(s) == 0 {
		return nil
	}

	r := make(map[string][]string)

	for _, sc := range s {
		t := T(sc)
		res := sc[t.categorizer().rootCat().String()]
		k := res.Target

		if res.Target == AnyTgt {
			k = All
		}

		for _, sk := range split(k) {
			r[sk] = addToSet(r[sk], split(sc[scopeKeyDataType].Target))
		}
	}

	return r
}

// returns v if set is empty,
// unions v with set, otherwise.
func addToSet(set []string, v []string) []string {
	if len(set) == 0 {
		return v
	}

	for _, vv := range v {
		var matched bool

		for _, s := range set {
			if vv == s {
				matched = true
				break
			}
		}

		if !matched {
			set = append(set, vv)
		}
	}

	return set
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// produces the discrete set of resource owners in the slice of scopes.
// Any and None values are discarded.
func resourceOwnersIn(s []scope, rootCat string) []string {
	rm := map[string]struct{}{}

	for _, sc := range s {
		for _, v := range split(sc[rootCat].Target) {
			rm[v] = struct{}{}
		}
	}

	rs := []string{}

	for k := range rm {
		if k != AnyTgt && k != NoneTgt {
			rs = append(rs, k)
		}
	}

	return rs
}

// ---------------------------------------------------------------------------
// scope helpers
// ---------------------------------------------------------------------------

type scopeConfig struct {
	usePathFilter   bool
	usePrefixFilter bool
}

type option func(*scopeConfig)

func (sc *scopeConfig) populate(opts ...option) {
	for _, opt := range opts {
		opt(sc)
	}
}

// PrefixMatch ensures the selector uses a Prefix comparator, instead
// of contains or equals.  Will not override a default Any() or None()
// comparator.
func PrefixMatch() option {
	return func(sc *scopeConfig) {
		sc.usePrefixFilter = true
	}
}

// pathType is an internal-facing option.  It is assumed that scope
// constructors will provide the pathType option whenever a folder-
// level scope (ie, a scope that compares path hierarchies) is created.
func pathType() option {
	return func(sc *scopeConfig) {
		sc.usePathFilter = true
	}
}

func badCastErr(cast, is service) error {
	return errors.Wrapf(ErrorBadSelectorCast, "%s service is not %s", cast, is)
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
// otherwise returns the input
func clean(s []string) []string {
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

// filterize turns the slice into a filter.
// if the input is Any(), returns a passAny filter.
// if the input is None(), returns a failAny filter.
// if the scopeConfig specifies a filter, use that filter.
// if the input is len(1), returns an Equals filter.
// otherwise returns a Contains filter.
func filterize(sc scopeConfig, s ...string) filters.Filter {
	s = clean(s)

	if len(s) == 0 || s[0] == NoneTgt {
		return failAny
	}

	if s[0] == AnyTgt {
		return passAny
	}

	if sc.usePathFilter {
		if sc.usePrefixFilter {
			return filters.PathPrefix(s)
		}

		return filters.PathContains(s)
	}

	if sc.usePrefixFilter {
		return filters.Prefix(join(s...))
	}

	if len(s) == 1 {
		return filters.Equal(s[0])
	}

	return filters.Contains(join(s...))
}

type filterFunc func(string) filters.Filter

// wrapFilter produces a func that filterizes the input by:
// - cleans the input string
// - normalizes the cleaned input (returns anyFail if empty, allFail if *)
// - joins the string
// - and generates a filter with the joined input.
func wrapFilter(ff filterFunc) func([]string) filters.Filter {
	return func(s []string) filters.Filter {
		s = clean(s)

		if len(s) == 1 {
			if s[0] == AnyTgt {
				return passAny
			}

			if s[0] == NoneTgt {
				return failAny
			}
		}

		ss := join(s...)

		return ff(ss)
	}
}
