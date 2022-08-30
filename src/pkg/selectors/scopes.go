package selectors

import (
	"strings"

	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/filters"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

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

		// unknownType returns the unknown category value
		unknownCat() categorizer

		// pathValues should produce a map of category:string pairs populated by extracting
		// values out of the path that match the given categorizer.
		//
		// Ex: given a path like "tenant/service/root/dataType/folder/itemID", the func should
		// autodetect the data type using 'service' and 'dataType', and use the remaining
		// details to construct a map similar to this:
		// {
		//   rootCat:   root,
		//   folderCat: folder,
		//   itemCat:   itemID,
		// }
		pathValues([]string) map[categorizer]string

		// pathKeys produces a list of categorizers that can be used as keys in the pathValues
		// map.  The combination of the two funcs generically interprets the context of the
		// ids in a path with the same keys that it uses to retrieve those values from a scope,
		// so that the two can be compared.
		pathKeys() []categorizer
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
	// comply with.
	scoper interface {
		// Every scope is expected to contain a reference to its category.  This allows users
		// to evaluate structs with a call to myscope.Category().  Category() is expected to
		// return the service-specific type of the categorizer, since the end user is expected
		// to be operating within that context.
		// This func returns the same value as the categorizer interface so that the funcs
		// internal to scopes.go can utilize the scope's category without the service context.
		categorizer() categorizer

		// matchesEntry is used to determine if the scope values match with either the pathValues,
		// or the DetailsEntry for the given category.
		// The path comparison (using cat and pathValues) can be handled generically within
		// scopes.go.  However, the entry comparison requires service-specific context in order
		// for the scope to extract the correct serviceInfo in the entry.
		//
		// Params:
		// cat - the category type expressed in the Path.  Not the category of the Scope.  If the
		//   scope does not align with this parameter, the result is automatically false.
		// pathValues - the result of categorizer.pathValues() for the Path being checked.
		// entry - the details entry containing extended service info for the item that a filter may
		//   compare.  Identification of the correct entry Info service is left up to the scope.
		matchesEntry(cat categorizer, pathValues map[categorizer]string, entry details.DetailsEntry) bool

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

// makeScope produces a well formatted, typed scope that ensures all base values are populated.
func makeScope[T scopeT](
	granularity string,
	cat categorizer,
	resources, vs []string,
) T {
	s := T{
		scopeKeyCategory:       filters.NewIdentity(cat.String()),
		scopeKeyDataType:       filters.NewIdentity(cat.leafCat().String()),
		scopeKeyGranularity:    filters.NewIdentity(granularity),
		scopeKeyResource:       filters.NewIdentity(join(resources...)),
		cat.String():           filterize(vs...),
		cat.rootCat().String(): filterize(resources...),
	}

	return s
}

// makeFilterScope produces a well formatted, typed scope, with properties specifically oriented
// towards identifying filter-type scopes, that ensures all base values are populated.
func makeFilterScope[T scopeT](
	cat, filterCat categorizer,
	vs []string,
	f func([]string) filters.Filter,
) T {
	return T{
		scopeKeyCategory:    filters.NewIdentity(cat.String()),
		scopeKeyDataType:    filters.NewIdentity(cat.leafCat().String()),
		scopeKeyGranularity: filters.NewIdentity(Filter),
		scopeKeyInfoFilter:  filters.NewIdentity(filterCat.String()),
		scopeKeyResource:    filters.NewIdentity(Filter),
		filterCat.String():  f(clean(vs)),
	}
}

// ---------------------------------------------------------------------------
// scope funcs
// ---------------------------------------------------------------------------

// matches returns true if the category is included in the scope's
// data type, and the target string is included in the scope.
func matches[T scopeT, C categoryT](s T, cat C, target string) bool {
	if !typeAndCategoryMatches(cat, s.categorizer()) {
		return false
	}

	if len(target) == 0 {
		return false
	}

	return s[cat.String()].Matches(target)
}

// getCategory returns the scope's category value.
// if s is a filter-type scope, returns the filter category.
func getCategory[T scopeT](s T) string {
	return s[scopeKeyCategory].Target
}

// getFilterCategory returns the scope's infoFilter category value.
func getFilterCategory[T scopeT](s T) string {
	return s[scopeKeyInfoFilter].Target
}

// getGranularity returns the scope's granularity value.
func getGranularity[T scopeT](s T) string {
	return s[scopeKeyGranularity].Target
}

// getCatValue takes the value of s[cat], split it by the standard
// delimiter, and returns the slice.  If s[cat] is nil, returns
// None().
func getCatValue[T scopeT](s T, cat categorizer) []string {
	v, ok := s[cat.String()]
	if !ok {
		return None()
	}

	return split(v.Target)
}

// set sets a value by category to the scope.  Only intended for internal
// use, not for exporting to callers.
func set[T scopeT](s T, cat categorizer, v []string) T {
	s[cat.String()] = filterize(v...)
	return s
}

// granularity describes the granularity (directory || item)
// of the data in scope.
func granularity[T scopeT](s T) string {
	return s[scopeKeyGranularity].Target
}

// returns true if the category is included in the scope's category type,
// and the value is set to Any().
func isAnyTarget[T scopeT, C categoryT](s T, cat C) bool {
	if !typeAndCategoryMatches(cat, s.categorizer()) {
		return false
	}

	return s[cat.String()].Target == AnyTgt
}

// reduce filters the entries in the details to only those that match the
// inclusions, filters, and exclusions in the selector.
//
func reduce[T scopeT, C categoryT](
	deets *details.Details,
	s Selector,
	dataCategories map[pathType]C,
) *details.Details {
	if deets == nil {
		return nil
	}

	// aggregate each scope type by category for easier isolation in future processing.
	excls := scopesByCategory[T](s.Excludes, dataCategories)
	filts := scopesByCategory[T](s.Filters, dataCategories)
	incls := scopesByCategory[T](s.Includes, dataCategories)

	ents := []details.DetailsEntry{}

	// for each entry, compare that entry against the scopes of the same data type
	for _, ent := range deets.Entries {
		// todo: use Path pkg for this
		path := strings.Split(ent.RepoRef, "/")

		dc, ok := dataCategories[pathTypeIn(path)]
		if !ok {
			continue
		}

		passed := passes(
			dc,
			dc.pathValues(path),
			ent,
			excls[dc],
			filts[dc],
			incls[dc],
		)
		if passed {
			ents = append(ents, ent)
		}
	}

	reduced := &details.Details{DetailsModel: deets.DetailsModel}
	reduced.Entries = ents

	return reduced
}

// TODO: this is a hack.  We don't want these values declared here- it will get
// unwieldy to have all of them for all services.  They should be declared in
// paths, since that's where service- and data-type-specific assertions are owned.
type pathType int

const (
	unknownPathType pathType = iota
	exchangeEventPath
	exchangeContactPath
	exchangeMailPath
)

// return the service data type of the path.
// TODO: this is a hack.  We don't want this identification to occur in this
// package.  It should get handled in paths, since that's where service- and
// data-type-specific assertions are owned.
// Ideally, we'd use something like path.DataType() instead of this func.
func pathTypeIn(path []string) pathType {
	// not all paths will be len=3.  Most should be longer.
	// This just protects us from panicing below.
	if len(path) < 3 {
		return unknownPathType
	}

	switch path[2] {
	case "mail":
		return exchangeMailPath
	case "contact":
		return exchangeContactPath
	case "event":
		return exchangeEventPath
	}

	return unknownPathType
}

// groups each scope by its category of data (specified by the service-selector).
// ex: a slice containing the scopes [mail1, mail2, event1]
// would produce a map like { mail: [1, 2], event: [1] }
// so long as "mail" and "event" are contained in cats.
func scopesByCategory[T scopeT, C categoryT](
	scopes []scope,
	cats map[pathType]C,
) map[C][]T {
	m := map[C][]T{}
	for _, cat := range cats {
		m[cat] = []T{}
	}

	for _, sc := range scopes {
		for _, cat := range cats {
			t := T(sc)
			if typeAndCategoryMatches(cat, t.categorizer()) {
				m[cat] = append(m[cat], t)
			}
		}
	}

	return m
}

// passes compares each path to the included and excluded exchange scopes.  Returns true
// if the path is included, passes filters, and not excluded.
func passes[T scopeT](
	cat categorizer,
	pathValues map[categorizer]string,
	entry details.DetailsEntry,
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
			if inc.matchesEntry(cat, pathValues, entry) {
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
		if !filt.matchesEntry(cat, pathValues, entry) {
			return false
		}
	}

	// any matching exclusion means failure
	for _, exc := range excs {
		if exc.matchesEntry(cat, pathValues, entry) {
			return false
		}
	}

	return true
}

// matchesPathValues will check whether the pathValues have matching entries
// in the scope.  The keys of the values to match against are identified by
// the categorizer.
// Standard expectations apply: None() or missing values always fail, Any()
// always succeeds.
func matchesPathValues[T scopeT, C categoryT](
	sc T,
	cat C,
	pathValues map[categorizer]string,
) bool {
	// if scope specifies a filter category,
	// path checking is automatically skipped.
	if len(getFilterCategory(sc)) > 0 {
		return false
	}

	for _, c := range cat.pathKeys() {
		scopeVals := getCatValue(sc, c)
		// the scope must define the targets to match on
		if len(scopeVals) == 0 {
			return false
		}
		// None() fails all matches
		if scopeVals[0] == NoneTgt {
			return false
		}
		// the path must contain a value to match against
		pathVal, ok := pathValues[c]
		if !ok {
			return false
		}
		// all parts of the scope must match
		cc := c.(C)
		if !isAnyTarget(sc, cc) {
			f := filters.NewContains(false, cc, join(scopeVals...))
			if !f.Matches(pathVal) {
				return false
			}
		}
	}

	return true
}

// ---------------------------------------------------------------------------
// categorizer funcs
// ---------------------------------------------------------------------------

// categoryMatches returns true if:
// - neither type is 'unknown'
// - either type is the root type
// - the leaf types match
func categoryMatches[C categoryT](a, b C) bool {
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
