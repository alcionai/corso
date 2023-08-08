package count

type key string

const (
	// NewItemCreated should be used for non-skip, non-replace,
	// non-meta item creation counting.  IE: use it specifically
	// for counting new items (no collision) or copied items.
	NewItemCreated   key = "new-item-created"
	CollisionReplace key = "collision-replace"
	CollisionSkip    key = "collision-skip"
)
