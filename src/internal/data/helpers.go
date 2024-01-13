package data

import "sort"

// SortRestoreCollections performs an in-place sort on the provided collection.
func SortRestoreCollections(rcs []RestoreCollection) {
	sort.Slice(rcs, func(i, j int) bool {
		return rcs[i].FullPath().String() < rcs[j].FullPath().String()
	})
}
