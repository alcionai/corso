package selectors

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
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
	ErrorBadSelectorCast     = clues.New("wrong selector service type")
	ErrorNoMatchingItems     = clues.New("no items match the specified selectors")
	ErrorUnrecognizedService = clues.New("unrecognized service")
)

const (
	scopeKeyCategory     = "category"
	scopeKeyInfoCategory = "details_info_category"
	scopeKeyDataType     = "type"
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
	delimiter = string('\x1F')
	passAny   = filters.Pass()
	failAny   = filters.Fail()
)

// All is the resource name that gets output when the resource is AnyTgt.
// It is not used aside from printing resources.
const All = "All"

type Reducer interface {
	Reduce(context.Context, *details.Details, *fault.Bus) *details.Details
}

// selectorResourceOwners aggregates all discrete path category types described
// in the selector.  Category sets are grouped by their scope type (includes,
// excludes, filters).
type selectorPathCategories struct {
	Includes []path.CategoryType
	Excludes []path.CategoryType
	Filters  []path.CategoryType
}

type pathCategorier interface {
	PathCategories() selectorPathCategories
}

// ---------------------------------------------------------------------------
// Selector
// ---------------------------------------------------------------------------

// The core selector.  Has no api for setting or retrieving data.
// Is only used to pass along more specific selector instances.
type Selector struct {
	// The service scope of the data.  Exchange, Teams, SharePoint, etc.
	Service service `json:"service,omitempty"`

	// A record of the resource owners matched by this selector.
	ResourceOwners filters.Filter `json:"resourceOwners,omitempty"`

	// The single resource owner being observed by the selector.
	// Selectors are constructed by passing in a list of ResourceOwners,
	// and those owners represent the "total" data that should be operated
	// across all corso operations.  But any single operation (backup,restore,
	// etc) will only observe a single user at a time, and that user is
	// represented by this value.
	//
	// If the constructor is passed a len=1 list of owners, this value is
	// automatically matched to that entry.  For lists with more than one
	// owner, the user is expected to call SplitByResourceOwner(), and
	// iterate over the results, where each one will populate this field
	// with a different owner.
	DiscreteOwner string `json:"discreteOwner,omitempty"`
	// display name for the DiscreteOwner.
	DiscreteOwnerName string `json:"discreteOwnerName,omitempty"`

	// A slice of exclusion scopes.  Exclusions apply globally to all
	// inclusions/filters, with any-match behavior.
	Excludes []scope `json:"exclusions,omitempty"`
	// A slice of info scopes.  All inclusions must also match ALL filters.
	Filters []scope `json:"filters,omitempty"`
	// A slice of inclusion scopes.  Comparators must match either one of these,
	// or all filters, to be included.
	Includes []scope `json:"includes,omitempty"`
}

// helper for specific selector instance constructors.
func newSelector(s service, resourceOwners []string) Selector {
	var owner string
	if len(resourceOwners) == 1 && resourceOwners[0] != AnyTgt {
		owner = resourceOwners[0]
	}

	return Selector{
		Service:        s,
		ResourceOwners: filterize(scopeConfig{}, resourceOwners...),
		DiscreteOwner:  owner,
		Excludes:       []scope{},
		Includes:       []scope{},
	}
}

// DiscreteResourceOwners returns the list of individual resourceOwners used
// in the selector.
// TODO(rkeepers): remove in favor of split and s.DiscreteOwner
func (s Selector) DiscreteResourceOwners() []string {
	return split(s.ResourceOwners.Target)
}

// SetDiscreteOwnerIDName ensures the selector has the correct discrete owner
// id and name.  It is assumed that these values are sourced using the current
// s.DiscreteOwner as input.  The reason for taking in both the id and name, and
// not just the name, is so that constructors can input owner aliases in place
// of  ids, with the expectation that the two will get sorted and re-written
// later on with this setter.
func (s Selector) SetDiscreteOwnerIDName(id, name string) Selector {
	r := s

	if len(id) == 0 {
		// assume a the discreteOwner is already set, and don't replace anything.
		r.DiscreteOwnerName = s.DiscreteOwner
		return r
	}

	r.DiscreteOwner = id
	r.DiscreteOwnerName = name

	if len(name) == 0 {
		r.DiscreteOwnerName = id
	}

	return r
}

// ID returns s.discreteOwner, which is assumed to be a stable ID.
func (s Selector) ID() string {
	return s.DiscreteOwner
}

// Name returns s.discreteOwnerName.  If that value is empty, it returns
// s.DiscreteOwner instead.
func (s Selector) Name() string {
	if len(s.DiscreteOwnerName) == 0 {
		return s.DiscreteOwner
	}

	return s.DiscreteOwnerName
}

// isAnyResourceOwner returns true if the selector includes all resource owners.
func isAnyResourceOwner(s Selector) bool {
	return s.ResourceOwners.Comparator == filters.Passes
}

// isNoneResourceOwner returns true if the selector includes no resource owners.
func isNoneResourceOwner(s Selector) bool {
	return s.ResourceOwners.Comparator == filters.Fails
}

// SplitByResourceOwner makes one shallow clone for each resourceOwner in the
// selector, specifying a new DiscreteOwner for each one.
// If the original selector already specified a discrete slice of resource owners,
// only those owners are used in the result.
// If the original selector allowed Any() resource owner, the allOwners parameter
// is used to populate the slice.  allOwners is assumed to be the complete slice of
// resourceOwners in the tenant for the given service.
// If the original selector specified None(), thus failing all resource owners,
// an empty slice is returned.
//
// temporarily, clones all scopes in each selector and replaces the owners with
// the discrete owner.
func splitByResourceOwner[T scopeT, C categoryT](s Selector, allOwners []string, rootCat C) []Selector {
	if isNoneResourceOwner(s) {
		return []Selector{}
	}

	targets := allOwners

	if !isAnyResourceOwner(s) {
		targets = split(s.ResourceOwners.Target)
	}

	ss := make([]Selector, 0, len(targets))

	for _, ro := range targets {
		c := s
		c.DiscreteOwner = ro
		ss = append(ss, c)
	}

	return ss
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
func scopes[T scopeT](s Selector) []T {
	scopes := []T{}

	for _, v := range s.Includes {
		scopes = append(scopes, T(v))
	}

	return scopes
}

// Returns the path.ServiceType matching the selector service.
func (s Selector) PathService() path.ServiceType {
	return serviceToPathType[s.Service]
}

// Reduce is a quality-of-life interpreter that allows Reduce to be called
// from the generic selector by interpreting the selector service type rather
// than have the caller make that interpretation.  Returns an error if the
// service is unsupported.
func (s Selector) Reduce(
	ctx context.Context,
	deets *details.Details,
	errs *fault.Bus,
) (*details.Details, error) {
	r, err := selectorAsIface[Reducer](s)
	if err != nil {
		return nil, err
	}

	return r.Reduce(ctx, deets, errs), nil
}

// returns the sets of path categories identified in each scope set.
func (s Selector) PathCategories() (selectorPathCategories, error) {
	ro, err := selectorAsIface[pathCategorier](s)
	if err != nil {
		return selectorPathCategories{}, err
	}

	return ro.PathCategories(), nil
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
		err = clues.Stack(ErrorUnrecognizedService, clues.New(s.Service.String()))
	}

	return t, err
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

// produces the discrete set of path categories in the slice of scopes.
func pathCategoriesIn[T scopeT, C categoryT](ss []scope) []path.CategoryType {
	rm := map[path.CategoryType]struct{}{}

	for _, s := range ss {
		t := T(s)

		lc := t.categorizer().leafCat()
		if lc == lc.unknownCat() {
			continue
		}

		rm[lc.PathType()] = struct{}{}
	}

	rs := []path.CategoryType{}

	for k := range rm {
		rs = append(rs, k)
	}

	return rs
}

// ---------------------------------------------------------------------------
// scope constructors
// ---------------------------------------------------------------------------

type scopeConfig struct {
	usePathFilter   bool
	usePrefixFilter bool
	useSuffixFilter bool
	useEqualsFilter bool
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

// SuffixMatch ensures the selector uses a Suffix comparator, instead
// of contains or equals.  Will not override a default Any() or None()
// comparator.
func SuffixMatch() option {
	return func(sc *scopeConfig) {
		sc.useSuffixFilter = true
	}
}

// ExactMatch ensures the selector uses an Equals comparator, instead
// of contains.  Will not override a default Any() or None()
// comparator.
func ExactMatch() option {
	return func(sc *scopeConfig) {
		sc.useEqualsFilter = true
	}
}

// pathComparator is an internal-facing option.  It is assumed that scope
// constructors will provide the pathComparator option whenever a folder-
// level scope (ie, a scope that compares path hierarchies) is created.
func pathComparator() option {
	return func(sc *scopeConfig) {
		sc.usePathFilter = true
	}
}

func badCastErr(cast, is service) error {
	return clues.Stack(ErrorBadSelectorCast, clues.New(fmt.Sprintf("%s is not %s", cast, is)))
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
		if sc.useEqualsFilter {
			return filters.PathEquals(s)
		}

		if sc.usePrefixFilter {
			return filters.PathPrefix(s)
		}

		if sc.useSuffixFilter {
			return filters.PathSuffix(s)
		}

		return filters.PathContains(s)
	}

	if sc.usePrefixFilter {
		return filters.Prefix(join(s...))
	}

	if sc.useSuffixFilter {
		return filters.Suffix(join(s...))
	}

	if len(s) == 1 {
		return filters.Equal(s[0])
	}

	return filters.Contains(join(s...))
}

type (
	filterFunc      func(string) filters.Filter
	sliceFilterFunc func([]string) filters.Filter
)

// pathFilterFactory returns the appropriate path filter
// (contains, prefix, or suffix) for the provided options.
// If multiple options are flagged, Prefix takes priority.
// If no options are provided, returns PathContains.
func pathFilterFactory(opts ...option) sliceFilterFunc {
	sc := &scopeConfig{}
	sc.populate(opts...)

	var ff sliceFilterFunc

	switch true {
	case sc.usePrefixFilter:
		ff = filters.PathPrefix
	case sc.useSuffixFilter:
		ff = filters.PathSuffix
	default:
		ff = filters.PathContains
	}

	return wrapSliceFilter(ff)
}

func wrapSliceFilter(ff sliceFilterFunc) sliceFilterFunc {
	return func(s []string) filters.Filter {
		s = clean(s)

		if f, ok := isAnyOrNone(s); ok {
			return f
		}

		return ff(s)
	}
}

// wrapFilter produces a func that filterizes the input by:
// - cleans the input string
// - normalizes the cleaned input (returns anyFail if empty, allFail if *)
// - joins the string
// - and generates a filter with the joined input.
func wrapFilter(ff filterFunc) sliceFilterFunc {
	return func(s []string) filters.Filter {
		s = clean(s)

		if f, ok := isAnyOrNone(s); ok {
			return f
		}

		return ff(join(s...))
	}
}

// returns (<filter>, true) if s is len==1 and s[0] is
// anyTgt or noneTgt, implying that the caller should use
// the returned filter.  On (<filter>, false), the caller
// can ignore the returned filter.
// a special case exists for len(s)==0, interpreted as
// "noneTgt"
func isAnyOrNone(s []string) (filters.Filter, bool) {
	switch len(s) {
	case 0:
		return failAny, true

	case 1:
		switch s[0] {
		case AnyTgt:
			return passAny, true
		case NoneTgt:
			return failAny, true
		}
	}

	return failAny, false
}
