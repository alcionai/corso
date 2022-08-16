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
		pathValues([]string) map[categorizer]string
		// pathKeys produces a list of categorizers that can be used as keys in the pathValues
		// map.  The combination of the two funcs generically interprets the context of the
		// ids in a path with the same keys that it uses to retrieve those values from a scope,
		// so that the two can be compared.
		pathKeys() []categorizer
	}
	// categoryT is the generic type interface of a categorizer
	categoryT interface {
		~int
		categorizer
	}
)

type (
	scope map[string]string
	// scoper describes the minimum necessary interface that a soundly built scope should
	// comply with.
	scoper interface {
		categorizer() categorizer
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
