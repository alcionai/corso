package common

// UnionMaps produces a new map containing all the values of the other
// maps.  The last maps have the highes priority.  Key collisions with
// earlier maps will favor the last map with that key.
func UnionMaps[K comparable, V any](ms ...map[K]V) map[K]V {
	r := map[K]V{}

	for _, m := range ms {
		for k, v := range m {
			r[k] = v
		}
	}

	return r
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	r := map[K]V{}

	for k, v := range m {
		r[k] = v
	}

	return r
}
