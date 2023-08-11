package selectors

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

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
// types & interfaces
// ---------------------------------------------------------------------------

// leafProperty describes metadata associated with a leaf categorizer
type leafProperty struct {
	// pathKeys describes the categorizer keys used to map scope type to a value
	// extracted from a path.Path.
	// The order of the slice is important, and should match the order in which
	// these types appear in the path.Path for each type.
	// Ex: given: exchangeMail
	//	categoryPath => [ExchangeUser, ExchangeMailFolder, ExchangeMail]
	//	suggests that scopes involving exchange mail will need to match a user,
	//	mailFolder, and mail; appearing in the path in that order.
	pathKeys []categorizer

	// pathType produces the path.CategoryType representing this leafType.
	// This allows the scope to type to be compared using the more commonly recognized
	// path category consts.
	// Ex: given: exchangeMail
	//	pathType => path.EmailCategory
	pathType path.CategoryType
}

type (
	// categorizer recognizes service specific item categories.
	categorizer interface {
		// String should return the human readable name of the category.
		String() string

		// leafCat should return the lowest level type matching the category.  If the type
		// has multiple leaf types (ex: the root category) or no leaves (ex: unknown values),
		// the same value is returned.  Otherwise, if the receiver is an intermediary type,
		// such as a folder, then the child value should be returned.
		// Ex: fooFolder.leafCat() => foo.
		leafCat() categorizer

		// rootCat returns the root category for the categorizer
		rootCat() categorizer

		// unknownCat returns the unknown category value
		unknownCat() categorizer

		// isUnion returns true if the category can be used to match against any leaf category.
		// This can occur when an itemInfo property is used as an alternative resourceOwner id.
		isUnion() bool

		// isLeaf returns true if the category is one of the leaf categories.
		// eg: in a resourceOwner/folder/item structure, the item is the leaf.
		isLeaf() bool

		// pathValues takes in two paths, both variants of the repoRef, one containing the standard
		// repoRef, and the other amended to include the locationRef directories (if available).  It
		// should produce two maps of category:string pairs populated by extracting the values out of
		// each path.Path.
		//
		// Ex: given a path builder like ["tenant", "service", "resource", "dataType", "folder", "itemID"],
		// the func should use the path to construct a map similar to this:
		// {
		//   folderCat: folder,
		//   itemCat:   itemID,
		// }
		pathValues(path.Path, details.Entry, Config) (map[categorizer][]string, error)

		// pathKeys produces a list of categorizers that can be used as keys in the pathValues
		// map.  The combination of the two funcs generically interprets the context of the
		// ids in a path with the same keys that it uses to retrieve those values from a scope,
		// so that the two can be compared.
		pathKeys() []categorizer

		// PathType converts the category's leaf type into the matching path.CategoryType.
		// Exported due to common use by consuming packages.
		PathType() path.CategoryType
	}
	// categoryT is the generic type interface of a categorizer
	categoryT interface {
		~string
		categorizer
	}
)

type (
	// scopes are generic containers that hold comparable values and other metadata expressing
	// "the data to match on".  The matching behaviors that utilize scopes are: Inclusion (any-
	// match), Filter (all-match), and Exclusion (any-match).
	//
	// The values in a scope fall into one of two categories: comparables and metadata.
	//
	// Comparable values should be keyed by a categorizer.String() value, where that categorizer
	// is identified by the category set for the given service.  These values will be used in
	// path value comparisons (where the categorizer.pathValues() of the same key must match the
	// scope values), and details.Entry comparisons (where some entry.ServiceInfo is related to
	// the scope value).  Comparable values can also express a wildcard match (AnyTgt) or a no-
	// match (NoneTgt).
	//
	// Metadata values express details that are common across all service instances: data
	// granularity (group or item), resource (id of the root path resource), core data type
	// (human readable), or whether the scope is a filter-type or an inclusion-/exclusion-type.
	// Metadata values can be used in either logical processing of scopes, and/or for presentation
	// to end users.
	scope map[string]filters.Filter

	// scoper describes the minimum necessary interface that a soundly built scope should
	// comply with to be usable by selector generics.
	scoper interface {
		// Every scope is expected to contain a reference to its category.  This allows users
		// to evaluate structs with a call to myscope.Category().  Category() is expected to
		// return the service-specific type of the categorizer, since the end user is expected
		// to be operating within that context.
		// This func returns the same value as the categorizer interface so that the funcs
		// internal to scopes.go can utilize the scope's category without the service context.
		categorizer() categorizer

		// matchesInfo is used to determine if the scope values match a specific DetailsEntry
		// ItemInfo value.  Unlike path comparison, the entry comparison requires service-specific
		// context in order for the scope to extract the correct serviceInfo in the entry.
		//
		// Params:
		// info - the details entry itemInfo containing extended service info that a filter may
		//   compare.  Identification of the correct entry Info service is left up to the fulfiller.
		matchesInfo(info details.ItemInfo) bool

		// setDefaults populates default values for certain scope categories.
		// Primarily to ensure that root- or mid-tier scopes (such as folders)
		// cascade 'Any' matching to more granular categories.
		setDefaults()
	}
	// scopeT is the generic type interface of a scoper.
	scopeT interface {
		~map[string]filters.Filter
		scoper
	}
)

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

// ---------------------------------------------------------------------------
// scope config & constructors
// ---------------------------------------------------------------------------

// constructs the default item-scope comparator options according
// to the selector configuration.
//   - if cfg.OnlyMatchItemNames == false, then comparison assumes item IDs,
//     which are case sensitive, resulting in StrictEqualsMatch
func defaultItemOptions(cfg Config) []option {
	opts := []option{}

	if !cfg.OnlyMatchItemNames {
		opts = append(opts, StrictEqualMatch())
	}

	return opts
}

type scopeConfig struct {
	usePathFilter         bool
	usePrefixFilter       bool
	useSuffixFilter       bool
	useEqualsFilter       bool
	useStrictEqualsFilter bool
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

// StrictEqualsMatch ensures the selector uses a StrictEquals comparator, instead
// of contains.  Will not override a default Any() or None() comparator.
func StrictEqualMatch() option {
	return func(sc *scopeConfig) {
		sc.useStrictEqualsFilter = true
	}
}

// ExactMatch ensures the selector uses an Equals comparator, instead
// of contains.  Will not override a default Any() or None() comparator.
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

type filterFunc func([]string) filters.Filter

// filterize turns the slice into a filter.
// if the input is Any(), returns a passAny filter.
// if the input is None(), returns a failAny filter.
// if the scopeConfig specifies a filter, use that filter.
// if the input is len(1), returns an Equals filter.
// otherwise returns a Contains filter.
func filterFor(sc scopeConfig, targets ...string) filters.Filter {
	return filterize(sc, nil, targets...)
}

// filterize turns the slice into a filter.
// if the input is Any(), returns a passAny filter.
// if the input is None(), returns a failAny filter.
// if the scopeConfig specifies a filter, use that filter.
// if defaultFilter is non-nil, returns that filter.
// if the input is len(1), returns an Equals filter.
// otherwise returns a Contains filter.
func filterize(
	sc scopeConfig,
	defaultFilter filterFunc,
	targets ...string,
) filters.Filter {
	targets = clean(targets)

	if len(targets) == 0 || targets[0] == NoneTgt {
		return failAny
	}

	if targets[0] == AnyTgt {
		return passAny
	}

	if sc.usePathFilter {
		if sc.useEqualsFilter {
			return filters.PathEquals(targets)
		}

		if sc.usePrefixFilter {
			return filters.PathPrefix(targets)
		}

		if sc.useSuffixFilter {
			return filters.PathSuffix(targets)
		}

		return filters.PathContains(targets)
	}

	if sc.usePrefixFilter {
		return filters.Prefix(targets)
	}

	if sc.useSuffixFilter {
		return filters.Suffix(targets)
	}

	if sc.useStrictEqualsFilter {
		return filters.StrictEqual(targets)
	}

	if defaultFilter != nil {
		return defaultFilter(targets)
	}

	return filters.Equal(targets)
}

// pathFilterFactory returns the appropriate path filter
// (contains, prefix, or suffix) for the provided options.
// If multiple options are flagged, Prefix takes priority.
// If no options are provided, returns PathContains.
func pathFilterFactory(opts ...option) filterFunc {
	sc := &scopeConfig{}
	sc.populate(opts...)

	var ff filterFunc

	switch true {
	case sc.usePrefixFilter:
		ff = filters.PathPrefix
	case sc.useSuffixFilter:
		ff = filters.PathSuffix
	case sc.useEqualsFilter:
		ff = filters.PathEquals
	default:
		ff = filters.PathContains
	}

	return wrapSliceFilter(ff)
}

func wrapSliceFilter(ff filterFunc) filterFunc {
	return func(s []string) filters.Filter {
		s = clean(s)

		if f, ok := isAnyOrNone(s); ok {
			return f
		}

		return ff(s)
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

// makeScope produces a well formatted, typed scope that ensures all base values are populated.
func makeScope[T scopeT](
	cat categorizer,
	tgts []string,
	opts ...option,
) T {
	sc := &scopeConfig{}
	sc.populate(opts...)

	s := T{
		scopeKeyCategory: filters.Identity(cat.String()),
		scopeKeyDataType: filters.Identity(cat.leafCat().String()),
		cat.String():     filterFor(*sc, tgts...),
	}

	return s
}

// makeInfoScope produces a well formatted, typed scope, with properties specifically oriented
// towards identifying filter-type scopes, that ensures all base values are populated.
func makeInfoScope[T scopeT](
	cat, infoCat categorizer,
	tgts []string,
	ff filterFunc,
	opts ...option,
) T {
	sc := &scopeConfig{}
	sc.populate(opts...)

	return T{
		scopeKeyCategory:     filters.Identity(cat.String()),
		scopeKeyDataType:     filters.Identity(cat.leafCat().String()),
		scopeKeyInfoCategory: filters.Identity(infoCat.String()),
		infoCat.String():     filterize(*sc, ff, tgts...),
	}
}

// ---------------------------------------------------------------------------
// Stringers and Concealers
// ---------------------------------------------------------------------------

// loggableMSS transforms the scope into a map by stringifying each filter.
func loggableMSS[T scopeT](s T, plain bool) map[string]string {
	m := map[string]string{}

	for k, filt := range s {
		if plain {
			m[k] = filt.PlainString()
		} else {
			m[k] = filt.Conceal()
		}
	}

	return m
}

func conceal[T scopeT](s T) string {
	return marshalScope(loggableMSS(s, false))
}

func format[T scopeT](s T, fs fmt.State, _ rune) {
	fmt.Fprint(fs, conceal(s))
}

func plainString[T scopeT](s T) string {
	return marshalScope(loggableMSS(s, true))
}

func marshalScope(mss map[string]string) string {
	bs, err := json.Marshal(mss)
	if err != nil {
		return "error-marshalling-selector"
	}

	return string(bs)
}

// ---------------------------------------------------------------------------
// reducer & filtering
// ---------------------------------------------------------------------------

// reduce filters the entries in the details to only those that match the
// inclusions, filters, and exclusions in the selector.
func reduce[T scopeT, C categoryT](
	ctx context.Context,
	deets *details.Details,
	s Selector,
	dataCategories map[path.CategoryType]C,
	errs *fault.Bus,
) *details.Details {
	ctx, end := diagnostics.Span(ctx, "selectors:reduce")
	defer end()

	if deets == nil {
		return nil
	}

	el := errs.Local()

	// if a DiscreteOwner is specified, only match details for that owner.
	matchesResourceOwner := s.ResourceOwners
	if len(s.DiscreteOwner) > 0 {
		matchesResourceOwner = filterFor(scopeConfig{}, s.DiscreteOwner)
	}

	// aggregate each scope type by category for easier isolation in future processing.
	excls := scopesByCategory[T](s.Excludes, dataCategories, false)
	filts := scopesByCategory[T](s.Filters, dataCategories, true)
	incls := scopesByCategory[T](s.Includes, dataCategories, false)

	ents := []details.Entry{}

	// for each entry, compare that entry against the scopes of the same data type
	for _, ent := range deets.Items() {
		ictx := clues.Add(ctx, "short_ref", ent.ShortRef)

		repoPath, err := path.FromDataLayerPath(ent.RepoRef, true)
		if err != nil {
			el.AddRecoverable(ctx, clues.Wrap(err, "transforming repoRef to path").WithClues(ictx))
			continue
		}

		// first check, every entry needs to have at least one protected resource
		// that matches the selector's protected resources.
		if !matchesResourceOwner.CompareAny(
			path.ServiceResourcesToResources(
				repoPath.ServiceResources())...) {
			continue
		}

		dc, ok := dataCategories[repoPath.Category()]
		if !ok {
			continue
		}

		e, f, i := excls[dc], filts[dc], incls[dc]

		// at least one filter or inclusion must be presentt
		if len(f)+len(i) == 0 {
			continue
		}

		pv, err := dc.pathValues(repoPath, *ent, s.Cfg)
		if err != nil {
			el.AddRecoverable(ctx, clues.Wrap(err, "getting path values").WithClues(ictx))
			continue
		}

		passed := passes(dc, pv, *ent, e, f, i)
		if passed {
			ents = append(ents, *ent)
		}
	}

	reduced := &details.Details{DetailsModel: deets.DetailsModel}
	reduced.Entries = ents

	return reduced
}

// groups each scope by its category of data (specified by the service-selector).
// ex: a slice containing the scopes [mail1, mail2, event1]
// would produce a map like { mail: [1, 2], event: [1] }
// so long as "mail" and "event" are contained in cats.
// For ALL-mach requirements, scopes used as filters should force inclusion using
// includeAll=true, independent of the category.
func scopesByCategory[T scopeT, C categoryT](
	scopes []scope,
	cats map[path.CategoryType]C,
	includeAll bool,
) map[C][]T {
	m := map[C][]T{}
	for _, cat := range cats {
		m[cat] = []T{}
	}

	for _, sc := range scopes {
		for _, cat := range cats {
			t := T(sc)
			// include a scope if the data category matches, or the caller forces inclusion.
			if includeAll || typeAndCategoryMatches(cat, t.categorizer()) {
				m[cat] = append(m[cat], t)
			}
		}
	}

	return m
}

// passes compares each path to the included and excluded exchange scopes.  Returns true
// if the path is included, passes filters, and not excluded.
func passes[T scopeT, C categoryT](
	cat C,
	pathValues map[categorizer][]string,
	entry details.Entry,
	excs, filts, incs []T,
) bool {
	// a passing match requires either a filter or an inclusion
	if len(incs)+len(filts) == 0 {
		return false
	}

	// skip this check if 0 inclusions were populated
	// since filters act as the inclusion check in that case
	if len(incs) > 0 {
		// at least one inclusion must apply.
		var included bool

		for _, inc := range incs {
			if matchesEntry(inc, cat, pathValues, entry) {
				included = true
				break
			}
		}

		if !included {
			return false
		}
	}

	// all filters must pass
	for _, filt := range filts {
		if !matchesEntry(filt, cat, pathValues, entry) {
			return false
		}
	}

	// any matching exclusion means failure
	for _, exc := range excs {
		if matchesEntry(exc, cat, pathValues, entry) {
			return false
		}
	}

	return true
}

// matchesEntry determines whether the category and scope require a path
// comparison or an entry info comparison.
func matchesEntry[T scopeT, C categoryT](
	sc T,
	cat C,
	pathValues map[categorizer][]string,
	entry details.Entry,
) bool {
	// InfoCategory requires matching against service-specific info values
	if len(getInfoCategory(sc)) > 0 {
		return sc.matchesInfo(entry.ItemInfo)
	}

	return matchesPathValues(sc, cat, pathValues)
}

// matchesPathValues will check whether the pathValues have matching entries
// in the scope.  The keys of the values to match against are identified by
// the categorizer.
// Standard expectations apply: None() or missing values always fail, Any()
// always succeeds.
func matchesPathValues[T scopeT, C categoryT](
	sc T,
	cat C,
	pathValues map[categorizer][]string,
) bool {
	for _, c := range cat.pathKeys() {
		// resourceOwners are now checked at the beginning of the reduction.
		if c == c.rootCat() {
			continue
		}

		cc := c.(C)

		if isNoneTarget(sc, cc) {
			return false
		}

		if isAnyTarget(sc, cc) {
			// continue, not return: all path keys must match the entry to succeed
			continue
		}

		// the pathValues must have an entry for the given categorizer
		pathVals, ok := pathValues[c]
		if !ok || len(pathVals) == 0 {
			return false
		}

		if !matchesAny(sc, cc, pathVals) {
			return false
		}
	}

	return true
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// matches returns true if the category is included in the scope's
// data type, and the input string passes the scope's filter for
// that category.
func matches[T scopeT, C categoryT](s T, cat C, inpt string) bool {
	if !typeAndCategoryMatches(cat, s.categorizer()) {
		return false
	}

	if len(inpt) == 0 {
		return false
	}

	return s[cat.String()].Compare(inpt)
}

// matchesAny returns true if the category is included in the scope's
// data type, and any one of the input strings passes the scope's filter.
func matchesAny[T scopeT, C categoryT](s T, cat C, inpts []string) bool {
	if !typeAndCategoryMatches(cat, s.categorizer()) {
		return false
	}

	if len(inpts) == 0 {
		return false
	}

	return s[cat.String()].CompareAny(inpts...)
}

// getCategory returns the scope's category value.
// if s is an info-type scope, returns the info category.
func getCategory[T scopeT](s T) string {
	return s[scopeKeyCategory].Identity
}

// getInfoCategory returns the scope's infoFilter category value.
func getInfoCategory[T scopeT](s T) string {
	return s[scopeKeyInfoCategory].Identity
}

// getCatValue takes the value of s[cat] and returns the slice.
// If s[cat] is nil, returns None().
func getCatValue[T scopeT](s T, cat categorizer) []string {
	filt, ok := s[cat.String()]
	if !ok {
		return None()
	}

	if len(filt.Targets) > 0 {
		return filt.Targets
	}

	return filt.Targets
}

// set sets a value by category to the scope.  Only intended for internal
// use, not for exporting to callers.
func set[T scopeT](s T, cat categorizer, v []string, opts ...option) T {
	sc := &scopeConfig{}
	sc.populate(opts...)

	s[cat.String()] = filterFor(*sc, v...)

	return s
}

// returns true if the category is included in the scope's category type,
// and the value is set to None().
func isNoneTarget[T scopeT, C categoryT](s T, cat C) bool {
	if !typeAndCategoryMatches(cat, s.categorizer()) {
		return false
	}

	return s[cat.String()].Comparator == filters.Fails
}

// returns true if the category is included in the scope's category type,
// and the value is set to Any().
func isAnyTarget[T scopeT, C categoryT](s T, cat C) bool {
	if !typeAndCategoryMatches(cat, s.categorizer()) {
		return false
	}

	return s[cat.String()].Comparator == filters.Passes
}

// categoryMatches returns true if:
// - neither type is 'unknown'
// - either type is the root type
// - the leaf types match
func categoryMatches[C categoryT](a, b C) bool {
	if a.isUnion() || b.isUnion() {
		return true
	}

	u := a.unknownCat()
	if a == u || b == u {
		return false
	}

	r := a.rootCat()
	if a == r || b == r {
		return true
	}

	return a.leafCat() == b.leafCat()
}

// typeAndCategoryMatches returns true if:
// - both parameters are the same categoryT type
// - the category matches for both types
func typeAndCategoryMatches[C categoryT](a C, b categorizer) bool {
	bb, ok := b.(C)
	if !ok {
		return false
	}

	return categoryMatches(a, bb)
}
