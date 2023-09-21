package count

type key string

const (
	// count of bucket-tokens consumed by api calls.
	APICallTokensConsumed key = "api-call-tokens-consumed"
	// count of times that items had collisions during restore,
	// and that collision was solved by replacing the item.
	CollisionReplace key = "collision-replace"
	// count of times that items had collisions during restore,
	// and that collision was solved by skipping the item.
	CollisionSkip key = "collision-skip"
	// NewItemCreated should be used for non-skip, non-replace,
	// non-meta item creation counting.  IE: use it specifically
	// for counting new items (no collision) or copied items.
	NewItemCreated key = "new-item-created"
	// count of api calls that resulted in failure due to throttling.
	ThrottledAPICalls key = "throttled-api-calls"
)
