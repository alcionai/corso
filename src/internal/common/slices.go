package common

func ContainsString(super []string, sub string) bool {
	for _, s := range super {
		if s == sub {
			return true
		}
	}
	return false
}
