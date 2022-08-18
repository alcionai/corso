package selectors

import (
	"github.com/alcionai/corso/pkg/backup/details"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

type (
	// categorizer recognizes service specific item categories.
	categorizer interface {
		// String should return the human readable name of the category.
		String() string

		// includesType should return true if the parameterized category is, contextually
		// within the service, a subset of the receiver category.  Ex: a Mail category
		// is a subset of a MailFolder category.
		includesType(categorizer) bool

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
	// TODO: Uncomment when reducer func is added
	// categoryT is the generic type interface of a categorizer
	// categoryT interface {
	// 	~int
	// 	categorizer
	// }
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
	scope map[string]string

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
	}
	// scopeT is the generic type interface of a scoper.
	scopeT interface {
		~map[string]string
		scoper
	}
)

// ---------------------------------------------------------------------------
// funcs
// ---------------------------------------------------------------------------

// TODO: Uncomment when selectors.go/contains() can be removed.
//
// contains returns true if the category is included in the scope's
// data type, and the target string is included in the scope.
// func contains[T scopeT](s T, cat categorizer, target string) bool {
// 	if !s.categorizer().includesType(cat) {
// 		return false
// 	}
// 	compare := s[cat.String()]
// 	if len(compare) == 0 {
// 		return false
// 	}
// 	if compare == NoneTgt {
// 		return false
// 	}
// 	if compare == AnyTgt {
// 		return true
// 	}
// 	return strings.Contains(compare, target)
// }

// getCatValue takes the value of s[cat], split it by the standard
// delimiter, and returns the slice.  If s[cat] is nil, returns
// None().
func getCatValue[T scopeT](s T, cat categorizer) []string {
	v, ok := s[cat.String()]
	if !ok {
		return None()
	}
	return split(v)
}

// granularity describes the granularity (directory || item)
// of the data in scope.
func granularity[T scopeT](s T) string {
	return s[scopeKeyGranularity]
}

// returns true if the category is included in the scope's category type,
// and the value is set to Any().
func isAnyTarget[T scopeT](s T, cat categorizer) bool {
	if !s.categorizer().includesType(cat) {
		return false
	}
	return s[cat.String()] == AnyTgt
}
