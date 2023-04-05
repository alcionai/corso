package pii

import "strings"

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
