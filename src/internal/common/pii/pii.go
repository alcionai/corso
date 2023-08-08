package pii

import (
	"strings"

	"github.com/alcionai/clues"
	"golang.org/x/exp/slices"
)

// MapWithPlurls places the toLower value of each string
// into a map[string]struct{}, along with a copy of the that
// string as a plural (ex: FoO => foo, foos).
func MapWithPlurals(ss ...string) map[string]struct{} {
	mss := make(map[string]struct{}, len(ss)*2)

	for _, s := range ss {
		tl := strings.ToLower(s)
		mss[tl] = struct{}{}
		mss[tl+"s"] = struct{}{}
	}

	return mss
}

// ConcealElements conceals each element in elems that does not appear in
// the safe map.  A copy of elems containing the changes is returned.
func ConcealElements(elems []string, safe map[string]struct{}) []string {
	if len(elems) == 0 {
		return []string{}
	}

	ces := slices.Clone(elems)

	for i := range ces {
		ce := ces[i]

		if _, ok := safe[strings.ToLower(ce)]; !ok {
			ces[i] = clues.Conceal(ce)
		}
	}

	return ces
}
