package keys

type Set map[string]struct{}

func (ks Set) HasKey(key string) bool {
	if _, ok := ks[key]; ok {
		return true
	}

	return false
}

func (ks Set) Keys() []string {
	sliceKeys := make([]string, 0)

	for k := range ks {
		sliceKeys = append(sliceKeys, k)
	}

	return sliceKeys
}

func HasKeys(data map[string]any, keys ...string) bool {
	for _, k := range keys {
		if _, ok := data[k]; !ok {
			return false
		}
	}

	return true
}
