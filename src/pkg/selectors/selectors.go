package selectors

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

type service int

//go:generate stringer -type=service -linecomment
const (
	ServiceUnknown    service = 0 // Unknown Service
	ServiceExchange   service = 1 // Exchange
	ServiceOneDrive   service = 2 // OneDrive
	ServiceSharePoint service = 3 // SharePoint
	ServiceGroups     service = 4 // Groups
)

var serviceToPathType = map[service]path.ServiceType{
	ServiceUnknown:    path.UnknownService,
	ServiceExchange:   path.ExchangeService,
	ServiceOneDrive:   path.OneDriveService,
	ServiceSharePoint: path.SharePointService,
	ServiceGroups:     path.GroupsService,
}

var (
	ErrorBadSelectorCast     = clues.New("wrong selector service type")
	ErrorNoMatchingItems     = clues.New("no items match the provided selectors")
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
	passAny = filters.Pass()
	failAny = filters.Fail()
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

type pathServicer interface {
	PathService() path.ServiceType
}

type reasoner interface {
	Reasons(tenantID string, useOwnerNameForID bool) []identity.Reasoner
}

// ---------------------------------------------------------------------------
// Selector
// ---------------------------------------------------------------------------

var _ idname.Provider = &Selector{}

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

	Cfg Config `json:"cfg,omitempty"`
}

// Config defines broad-scale selector behavior.
type Config struct {
	// OnlyMatchItemNames tells the reducer to ignore matching on itemRef values
	// and other item IDs in favor of matching the item name.  Normal behavior only
	// matches on itemRefs.
	OnlyMatchItemNames bool
}

// helper for specific selector instance constructors.
func newSelector(s service, resourceOwners []string) Selector {
	var owner string
	if len(resourceOwners) == 1 && resourceOwners[0] != AnyTgt {
		owner = resourceOwners[0]
	}

	return Selector{
		Service:        s,
		ResourceOwners: filterFor(scopeConfig{}, resourceOwners...),
		DiscreteOwner:  owner,
		Excludes:       []scope{},
		Includes:       []scope{},
	}
}

// Configure sets the selector configuration.
func (s *Selector) Configure(cfg Config) {
	s.Cfg = cfg
}

// ---------------------------------------------------------------------------
// protected resources & idname provider compliance
// ---------------------------------------------------------------------------

var _ idname.Provider = &Selector{}

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

// SetDiscreteOwnerIDName ensures the selector has the correct discrete owner
// id and name.  Assumes that these values are sourced using the current
// s.DiscreteOwner as input.  The reason for taking in both the id and name, and
// not just the name, is so that constructors can input owner aliases in place
// of  ids, with the expectation that the two will get sorted and re-written
// later on with this setter.
//
// If the id is empty, the original DiscreteOwner value is retained.
// If the name is empty, the id is duplicated as the name.
func (s Selector) SetDiscreteOwnerIDName(id, name string) Selector {
	r := s

	if len(id) == 0 {
		// assume a the discreteOwner is already set, and don't replace anything.
		id = s.DiscreteOwner
	}

	r.DiscreteOwner = id
	r.DiscreteOwnerName = name

	if len(name) == 0 {
		r.DiscreteOwnerName = id
	}

	return r
}

// isAnyProtectedResource returns true if the selector includes all resource owners.
func isAnyProtectedResource(s Selector) bool {
	return s.ResourceOwners.Comparator == filters.Passes
}

// isNoneProtectedResource returns true if the selector includes no resource owners.
func isNoneProtectedResource(s Selector) bool {
	return s.ResourceOwners.Comparator == filters.Fails
}

// splitByProtectedResource makes one shallow clone for each resourceOwner in the
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
func splitByProtectedResource[T scopeT, C categoryT](s Selector, allOwners []string, rootCat C) []Selector {
	if isNoneProtectedResource(s) {
		return []Selector{}
	}

	targets := allOwners

	if !isAnyProtectedResource(s) {
		targets = s.ResourceOwners.Targets
	}

	ss := make([]Selector, 0, len(targets))

	for _, ro := range targets {
		c := s
		c.DiscreteOwner = ro
		ss = append(ss, c)
	}

	return ss
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

// PathCategories returns the sets of path categories identified in each scope set.
func (s Selector) PathCategories() (selectorPathCategories, error) {
	ro, err := selectorAsIface[pathCategorier](s)
	if err != nil {
		return selectorPathCategories{}, err
	}

	return ro.PathCategories(), nil
}

// Reasons returns a deduplicated set of the backup reasons produced
// using the selector's discrete owner and each scopes' service and
// category types.
func (s Selector) Reasons(tenantID string, useOwnerNameForID bool) ([]identity.Reasoner, error) {
	ro, err := selectorAsIface[reasoner](s)
	if err != nil {
		return nil, err
	}

	return ro.Reasons(tenantID, useOwnerNameForID), nil
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
	case ServiceGroups:
		a, err = func() (any, error) { return s.ToGroupsRestore() }()
		t = a.(T)
	default:
		err = clues.Stack(ErrorUnrecognizedService, clues.New(s.Service.String()))
	}

	return t, err
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

var _ clues.Concealer = &Selector{}

type loggableSelector struct {
	Service        service             `json:"service,omitempty"`
	ResourceOwners string              `json:"resourceOwners,omitempty"`
	DiscreteOwner  string              `json:"discreteOwner,omitempty"`
	Excludes       []map[string]string `json:"exclusions,omitempty"`
	Filters        []map[string]string `json:"filters,omitempty"`
	Includes       []map[string]string `json:"includes,omitempty"`
}

func (s Selector) Conceal() string {
	ls := loggableSelector{
		Service:        s.Service,
		ResourceOwners: s.ResourceOwners.Conceal(),
		DiscreteOwner:  clues.Conceal(s.DiscreteOwner),
		Excludes:       toMSS(s.Excludes, false),
		Filters:        toMSS(s.Filters, false),
		Includes:       toMSS(s.Includes, false),
	}

	return ls.marshal()
}

func (s Selector) Format(fs fmt.State, _ rune) {
	fmt.Fprint(fs, s.Conceal())
}

func (s Selector) String() string {
	return s.Conceal()
}

func (s Selector) PlainString() string {
	ls := loggableSelector{
		Service:        s.Service,
		ResourceOwners: s.ResourceOwners.PlainString(),
		DiscreteOwner:  s.DiscreteOwner,
		Excludes:       toMSS(s.Excludes, true),
		Filters:        toMSS(s.Filters, true),
		Includes:       toMSS(s.Includes, true),
	}

	return ls.marshal()
}

func toMSS(scs []scope, plain bool) []map[string]string {
	mss := make([]map[string]string, 0, len(scs))

	for _, s := range scs {
		m := map[string]string{}

		for k, filt := range s {
			if plain {
				m[k] = filt.PlainString()
			} else {
				m[k] = filt.Conceal()
			}
		}

		mss = append(mss, m)
	}

	return mss
}

func (ls loggableSelector) marshal() string {
	bs, err := json.Marshal(ls)
	if err != nil {
		return "error-marshalling-selector"
	}

	return string(bs)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// produces the discrete set of path categories in the slice of scopes.
func pathCategoriesIn[T scopeT, C categoryT](ss []scope) []path.CategoryType {
	m := map[path.CategoryType]struct{}{}

	for _, s := range ss {
		t := T(s)

		lc := t.categorizer().leafCat()
		if lc == lc.unknownCat() {
			continue
		}

		m[lc.PathType()] = struct{}{}
	}

	return maps.Keys(m)
}
