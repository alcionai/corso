package common

import "strconv"

func ContainsString(super []string, sub string) bool {
	for _, s := range super {
		if s == sub {
			return true
		}
	}

	return false
}

// First returns the first non-zero valued string
func First(vs ...string) string {
	for _, v := range vs {
		if len(v) > 0 {
			return v
		}
	}

	return ""
}

// parseBool returns the bool value represented by the string
// or false on error
func ParseBool(v string) bool {
	s, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}

	return s
}
