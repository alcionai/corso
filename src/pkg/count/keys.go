package count

type key string

const (
	// count of bucket-tokens consumed by api calls.
	APICallTokensConsumed key = "api-call-tokens-consumed"
	// count of api calls that resulted in failure due to throttling.
	ThrottledAPICalls key = "throttled-api-calls"
)

// Tracked during backup
const (
	// amounts reported by kopia
	PersistedCachedFiles          key = "persisted-cached-files"
	PersistedDirectories          key = "persisted-directories"
	PersistedFiles                key = "persisted-files"
	PersistedHashedBytes          key = "persisted-hashed-bytes"
	PersistedNonCachedFiles       key = "persisted-non-cached-files"
	PersistedNonMetaFiles         key = "persisted-non-meta-files"
	PersistedNonMetaUploadedBytes key = "persisted-non-meta-uploaded-bytes"
	PersistedUploadedBytes        key = "persisted-uploaded-bytes"
	PersistenceErrors             key = "persistence-errors"
	PersistenceExpectedErrors     key = "persistence-expected-errors"
	PersistenceIgnoredErrors      key = "persistence-ignored-errors"
	// amounts reported by data providers
	ProviderItemsRead key = "provider-items-read"
)

// Tracked during restore
const (
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
)
