package selectors

import (
	"context"

	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/logger"
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

		// unknownType returns the unknown category value
		unknownCat() categorizer

		// isLeaf returns true if the category is one of the leaf categories.
		// eg: in a resourceOwner/folder/item structure, the item is the leaf.
		isLeaf() bool

		// pathValues should produce a map of category:string pairs populated by extracting
		// values out of the path.Path struct.
		//
		// Ex: given a path builder like ["tenant", "service", "resource", "dataType", "folder", "itemID"],
		// the func should use the path to construct a map similar to this:
		// {
		//   rootCat:   resource,
		//   folderCat: folder,
		//   itemCat:   itemID,
		// }
		pathValues(path.Path) map[categorizer]string

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
		// ItemInfo filter.  Unlike path filtering, the entry comparison requires service-specific
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

// makeScope produces a well formatted, typed scope that ensures all base values are populated.
func makeScope[T scopeT](
	cat categorizer,
	resources, vs []string,
	opts ...option,
) T {
	sc := &scopeConfig{}
	sc.populate(opts...)

	s := T{
		scopeKeyCategory:       filters.Identity(cat.String()),
		scopeKeyDataType:       filters.Identity(cat.leafCat().String()),
		cat.String():           filterize(*sc, vs...),
		cat.rootCat().String(): filterize(scopeConfig{}, resources...),
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
		scopeKeyCategory:   filters.Identity(cat.String()),
		scopeKeyDataType:   filters.Identity(cat.leafCat().String()),
		scopeKeyInfoFilter: filters.Identity(filterCat.String()),
		filterCat.String(): f(clean(vs)),
	}
}

// ---------------------------------------------------------------------------
// scope funcs
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

// getCategory returns the scope's category value.
// if s is a filter-type scope, returns the filter category.
func getCategory[T scopeT](s T) string {
	return s[scopeKeyCategory].Target
}

// getFilterCategory returns the scope's infoFilter category value.
func getFilterCategory[T scopeT](s T) string {
	return s[scopeKeyInfoFilter].Target
}

// getCatValue takes the value of s[cat], split it by the standard
// delimiter, and returns the slice.  If s[cat] is nil, returns
// None().
func getCatValue[T scopeT](s T, cat categorizer) []string {
	filt, ok := s[cat.String()]
	if !ok {
		return None()
	}

	if len(filt.Targets) > 0 {
		return filt.Targets
	}

	return split(filt.Target)
}

// set sets a value by category to the scope.  Only intended for internal
// use, not for exporting to callers.
func set[T scopeT](s T, cat categorizer, v []string, opts ...option) T {
	sc := &scopeConfig{}
	sc.populate(opts...)

	s[cat.String()] = filterize(*sc, v...)

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

// reduce filters the entries in the details to only those that match the
// inclusions, filters, and exclusions in the selector.
func reduce[T scopeT, C categoryT](
	ctx context.Context,
	deets *details.Details,
	s Selector,
	dataCategories map[path.CategoryType]C,
) *details.Details {
	ctx, end := D.Span(ctx, "selectors:reduce")
	defer end()

	if deets == nil {
		return nil
	}

	// aggregate each scope type by category for easier isolation in future processing.
	excls := scopesByCategory[T](s.Excludes, dataCategories, false)
	filts := scopesByCategory[T](s.Filters, dataCategories, true)
	incls := scopesByCategory[T](s.Includes, dataCategories, false)

	ents := []details.DetailsEntry{}

	// for each entry, compare that entry against the scopes of the same data type
	for _, ent := range deets.Items() {
		repoPath, err := path.FromDataLayerPath(ent.RepoRef, true)
		if err != nil {
			logger.Ctx(ctx).Debugw("transforming repoRef to path", "err", err)
			continue
		}

		dc, ok := dataCategories[repoPath.Category()]
		if !ok {
			continue
		}

		passed := passes(
			dc,
			dc.pathValues(repoPath),
			*ent,
			excls[dc],
			filts[dc],
			incls[dc],
		)
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
	pathValues map[categorizer]string,
	entry details.DetailsEntry,
) bool {
	// filterCategory requires matching against service-specific info values
	if len(getFilterCategory(sc)) > 0 {
		return sc.matchesInfo(entry.ItemInfo)
	}

	return matchesPathValues(sc, cat, pathValues, entry.ShortRef)
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
	shortRef string,
) bool {
	for _, c := range cat.pathKeys() {
		// the pathValues must have an entry for the given categorizer
		pathVal, ok := pathValues[c]
		if !ok {
			return false
		}

		cc := c.(C)

		if isNoneTarget(sc, cc) {
			return false
		}

		if isAnyTarget(sc, cc) {
			// continue, not return: all path keys must match the entry to succeed
			continue
		}

		var (
			match  bool
			isLeaf = c.isLeaf()
		)

		switch {
		// Leaf category - the scope can match either the path value (the item ID itself),
		// or the shortRef hash representing the item.
		case isLeaf && len(shortRef) > 0:
			match = matches(sc, cc, pathVal) || matches(sc, cc, shortRef)

		// all other categories (root, folder, etc) just need to pass the filter
		default:
			match = matches(sc, cc, pathVal)
		}

		if !match {
			return false
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
