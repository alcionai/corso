package count

type Key string

const (
	// count of bucket-tokens consumed by api calls.
	APICallTokensConsumed Key = "api-call-tokens-consumed"
	// count of api calls that resulted in failure due to throttling.
	ThrottledAPICalls Key = "throttled-api-calls"
)

// Tracked during backup
const (
	// amounts reported by kopia

	PersistedCachedFiles          Key = "persisted-cached-files"
	PersistedDirectories          Key = "persisted-directories"
	PersistedFiles                Key = "persisted-files"
	PersistedHashedBytes          Key = "persisted-hashed-bytes"
	PersistedNonCachedFiles       Key = "persisted-non-cached-files"
	PersistedNonMetaFiles         Key = "persisted-non-meta-files"
	PersistedNonMetaUploadedBytes Key = "persisted-non-meta-uploaded-bytes"
	PersistedUploadedBytes        Key = "persisted-uploaded-bytes"
	PersistenceErrors             Key = "persistence-errors"
	PersistenceExpectedErrors     Key = "persistence-expected-errors"
	PersistenceIgnoredErrors      Key = "persistence-ignored-errors"
	// amounts reported by data providers
	CollectionMoved               Key = "collection-moved"
	CollectionNew                 Key = "collection-state-new"
	CollectionNotMoved            Key = "collection-state-not-moved"
	CollectionTombstoned          Key = "collection-state-tombstoned"
	Collections                   Key = "collections"
	DeleteFolderMarker            Key = "delete-folder-marker"
	DeleteItemMarker              Key = "delete-item-marker"
	Drives                        Key = "drives"
	DriveTombstones               Key = "drive-tombstones"
	Files                         Key = "files"
	Folders                       Key = "folders"
	ItemsAdded                    Key = "items-added"
	ItemsRemoved                  Key = "items-removed"
	LazyDeletedInFlight           Key = "lazy-deleted-in-flight"
	Malware                       Key = "malware"
	MissingDelta                  Key = "missing-delta-token"
	NewDeltas                     Key = "new-delta-tokens"
	NewPrevPaths                  Key = "new-previous-paths"
	NoDeltaQueries                Key = "cannot-make-delta-queries"
	Packages                      Key = "packages"
	PagerResets                   Key = "pager-resets"
	PagesEnumerated               Key = "pages-enumerated"
	PrevDeltas                    Key = "previous-deltas"
	PrevPaths                     Key = "previous-paths"
	PreviousPathMetadataCollision Key = "previous-path-metadata-collision"
	ProviderItemsRead             Key = "provider-items-read"
	SkippedContainers             Key = "skipped-containers"
	StreamBytesAdded              Key = "stream-bytes-added"
	StreamDirsAdded               Key = "stream-dirs-added"
	StreamDirsFound               Key = "stream-dirs-found"
	StreamItemsAdded              Key = "stream-items-added"
	StreamItemsDeletedInFlight    Key = "stream-items-deleted-in-flight"
	StreamItemsErrored            Key = "stream-items-errored"
	StreamItemsFound              Key = "stream-items-found"
	StreamItemsRemoved            Key = "stream-items-removed"
	URLCacheMiss                  Key = "url-cache-miss"
	URLCacheRefresh               Key = "url-cache-refresh"
)

// Counted using clues error labels
const (
	BadCollPath                 = "bad_collection_path"
	BadPathPrefix               = "bad_path_prefix_creation"
	BadPrevPath                 = "unparsable_prev_path"
	CollectionTombstoneConflict = "collection_tombstone_conflicts_with_live_collection"
	ItemBeforeParent            = "item_before_parent"
	MissingParent               = "missing_parent"
	UnknownItemType             = "unknown_item_type"
)

// Tracked during restore
const (
	// count of times that items had collisions during restore,
	// and that collision was solved by replacing the item.
	CollisionReplace Key = "collision-replace"
	// count of times that items had collisions during restore,
	// and that collision was solved by skipping the item.
	CollisionSkip Key = "collision-skip"
	// NewItemCreated should be used for non-skip, non-replace,
	// non-meta item creation counting.  IE: use it specifically
	// for counting new items (no collision) or copied items.
	NewItemCreated Key = "new-item-created"
)
