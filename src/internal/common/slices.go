package common

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

// Equal returns true if both slices contain the same
// elements in the same order
func AreSameSlice[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true
}
