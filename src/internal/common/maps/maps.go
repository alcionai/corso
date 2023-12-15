package maps

func HasKeys(data map[string]any, keys ...string) bool {
	for _, k := range keys {
		if _, ok := data[k]; !ok {
			return false
		}
	}

	return true
}
