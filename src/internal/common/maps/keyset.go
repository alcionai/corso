package maps

type KeySet map[string]struct{}

func (ks KeySet) HasKey(key string) bool {
	if _, ok := ks[key]; ok {
		return true
	}

	return false
}

func (ks KeySet) Keys() []string {
	sliceKeys := make([]string, 0)

	for k := range ks {
		sliceKeys = append(sliceKeys, k)
	}

	return sliceKeys
}
