package common

import "strconv"

// parseBool returns the bool value represented by the string
// or false on error
func ParseBool(v string) bool {
	s, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}

	return s
}
