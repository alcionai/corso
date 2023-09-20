package count

type key string

const (
	APICallTokensConsumed key = "api-call-tokens-consumed"
	CollisionReplace      key = "collision-replace"
	CollisionSkip         key = "collision-skip"
	// NewItemCreated should be used for non-skip, non-replace,
	// non-meta item creation counting.  IE: use it specifically
	// for counting new items (no collision) or copied items.
	NewItemCreated    key = "new-item-created"
	ThrottledAPICalls key = "throttled-api-calls"
)
