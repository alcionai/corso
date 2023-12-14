package control

import (
	"github.com/alcionai/corso/src/pkg/extensions"
)

// DefaultBackupOptions provides a Backup with the default values set.
func DefaultBackupOptions() Backup {
	return Backup{
		FailureHandling: FailAfterRecovery,
		DeltaPageSize:   500,
		Parallelism: Parallelism{
			CollectionBuffer: 4,
			ItemFetch:        4,
		},
	}
}

// Backup is the set of options used for backup operations. Each set of options
// is only applied to the backup operation it's passed to. To use the same set
// of options for multiple backup operations pass the struct to all operations.
type Backup struct {
	// DeltaPageSize controls the quantity of items fetched in each page
	// during multi-page queries, such as graph api delta endpoints.
	DeltaPageSize        int32                              `json:"deltaPageSize"`
	FailureHandling      FailurePolicy                      `json:"failureHandling"`
	ItemExtensionFactory []extensions.CreateItemExtensioner `json:"-"`
	Parallelism          Parallelism                        `json:"parallelism"`
	ToggleFeatures       BackupToggles                      `json:"toggleFeatures"`
	ServiceRateLimiter   RateLimiter                        `json:"serviceRateLimiter"`

	// PreviewItemLimits defines the number of items and/or amount of data to
	// fetch on a best-effort basis. Right now it's used for preview backups.
	//
	// Since this is not split out by service or data categories these limits
	// apply independently to all data categories that appear in a single backup
	// where they are set. For example, if doing a teams backup and there's both a
	// SharePoint site and Messages available, both data categories would try to
	// backup data until the set limits without paying attention to what the other
	// had already backed up.
	PreviewLimits PreviewItemLimits `json:"previewItemLimits"`
}

type Parallelism struct {
	// CollectionBuffer sets the number of items in a collection to buffer before
	// blocking.
	CollectionBuffer int

	// ItemFetch sets the number of items to fetch in parallel when populating
	// items within a collection.
	ItemFetch int
}

// PreviewItemLimits describes best-effort maximum values to attempt to reach in
// this backup. Preview backups are used to demonstrate value by being quick to
// create.
type PreviewItemLimits struct {
	MaxItems             int
	MaxItemsPerContainer int
	MaxContainers        int
	MaxBytes             int64
	MaxPages             int
	Enabled              bool
}

type BackupToggles struct {
	// DisableIncrementals prevents backups from using incremental lookups,
	// forcing a new, complete backup of all data regardless of prior state.
	DisableIncrementals bool `json:"exchangeIncrementals,omitempty"`

	// ForceItemDataDownload disables finding cached items in previous failed
	// backups (i.e. kopia-assisted incrementals). Data dedupe will still occur
	// since that is based on content hashes. Items that have not changed since
	// the previous backup (i.e. in the merge base) will not be redownloaded. Use
	// DisableIncrementals to control that behavior.
	ForceItemDataDownload bool `json:"forceItemDataDownload,omitempty"`

	// DisableDelta prevents backups from using delta based lookups,
	// forcing a backup by enumerating all items. This is different
	// from DisableIncrementals in that this does not even makes use of
	// delta endpoints with or without a delta token. This is necessary
	// when the user has filled up the mailbox storage available to the
	// user as Microsoft prevents the API from being able to make calls
	// to delta endpoints.
	DisableDelta bool `json:"exchangeDelta,omitempty"`

	// ExchangeImmutableIDs denotes whether Corso should store items with
	// immutable Exchange IDs. This is only safe to set if the previous backup for
	// incremental backups used immutable IDs or if a full backup is being done.
	ExchangeImmutableIDs bool `json:"exchangeImmutableIDs,omitempty"`

	RunMigrations bool `json:"runMigrations"`

	// see: https://github.com/alcionai/corso/issues/4688
	UseDeltaTree bool `json:"useDeltaTree"`
}
