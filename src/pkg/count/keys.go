package count

type Key string

// ---------------------------------------------------------------------------
// General Rules:
// 1. Avoid the word "error", prefer "err".  That minimizes log result
//    contamination when using common filters like "logs with 'error'".
// 2. When some key (ex: Foo) can be counted with both an in-process
//    count, and also an end-of-process count, and the two may not be
//    equal, use Foo for the end of process count, and TotalFooProcessed
//    for the in-process count.
// ---------------------------------------------------------------------------

const (
	// count of bucket-tokens consumed by api calls.
	APICallTokensConsumed Key = "api-call-tokens-consumed"
	// count of api calls that resulted in failure due to throttling.
	ThrottledAPICalls Key = "throttled-api-calls"
)

// backup amounts reported by kopia
const (
	PersistedCachedFiles          Key = "persisted-cached-files"
	PersistedDirectories          Key = "persisted-directories"
	PersistedFiles                Key = "persisted-files"
	PersistedHashedBytes          Key = "persisted-hashed-bytes"
	PersistedNonCachedFiles       Key = "persisted-non-cached-files"
	PersistedNonMetaFiles         Key = "persisted-non-meta-files"
	PersistedNonMetaUploadedBytes Key = "persisted-non-meta-uploaded-bytes"
	PersistedUploadedBytes        Key = "persisted-uploaded-bytes"
	PersistenceErrs               Key = "persistence-errs"
	PersistenceExpectedErrs       Key = "persistence-expected-errs"
	PersistenceIgnoredErrs        Key = "persistence-ignored-errs"
)

// backup amounts reported by data providers
const (
	Channels                      Key = "channels"
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
	MetadataItems                 Key = "metadata-items"
	MetadataBytes                 Key = "metadata-bytes"
	MissingDelta                  Key = "missing-delta-token"
	NewDeltas                     Key = "new-delta-tokens"
	NewPrevPaths                  Key = "new-previous-paths"
	NoDeltaQueries                Key = "cannot-make-delta-queries"
	Packages                      Key = "packages"
	PagerResets                   Key = "pager-resets"
	PrevDeltas                    Key = "previous-deltas"
	PrevPaths                     Key = "previous-paths"
	PreviousPathMetadataCollision Key = "previous-path-metadata-collision"
	Sites                         Key = "sites"
	SkippedContainers             Key = "skipped-containers"
	StreamBytesAdded              Key = "stream-bytes-added"
	StreamDirsAdded               Key = "stream-dirs-added"
	StreamDirsFound               Key = "stream-dirs-found"
	StreamItemsAdded              Key = "stream-items-added"
	StreamItemsDeletedInFlight    Key = "stream-items-deleted-in-flight"
	StreamItemsErred              Key = "stream-items-erred"
	StreamItemsFound              Key = "stream-items-found"
	StreamItemsRemoved            Key = "stream-items-removed"
	TotalContainersSkipped        Key = "total-containers-skipped"
	URLCacheMiss                  Key = "url-cache-miss"
	URLCacheRefresh               Key = "url-cache-refresh"
	URLCacheItemNotFound          Key = "url-cache-item-not-found"
)

// Total___Processed counts are used to track raw processing numbers
// for values that may have a similar, but different, end result count.
// For example: a delta query may add the same folder to many different pages.
// instead of adding logic to catch folder duplications and only count new
// entries, we can increment TotalFoldersProcessed for every duplication,
// and use a separate Key (Folders) for the end count of folders produced
// at the end of the delta enumeration.
const (
	TotalDeleteFilesProcessed   Key = "total-delete-files-processed"
	TotalDeleteFoldersProcessed Key = "total-delete-folders-processed"
	TotalDeltasProcessed        Key = "total-deltas-processed"
	TotalFilesProcessed         Key = "total-files-processed"
	TotalFoldersProcessed       Key = "total-folders-processed"
	TotalMalwareProcessed       Key = "total-malware-processed"
	TotalPackagesProcessed      Key = "total-packages-processed"
	TotalPagesEnumerated        Key = "total-pages-enumerated"
)

// miscellaneous
const (
	RequiresUserPnToIDMigration Key = "requires-user-pn-to-id-migration"
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
