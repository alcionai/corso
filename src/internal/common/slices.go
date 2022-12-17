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
