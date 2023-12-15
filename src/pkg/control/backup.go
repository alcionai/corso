package control

import (
	"github.com/alcionai/corso/src/pkg/extensions"
)

// DefaultBackupOptions provides a Backup with the default values set.
func DefaultBackupConfig() BackupConfig {
	return BackupConfig{
		FailureHandling: FailAfterRecovery,
		Parallelism: Parallelism{
			CollectionBuffer: 4,
			ItemFetch:        4,
		},
		M365: BackupM365Config{
			DeltaPageSize: 500,
		},
	}
}

// Backup is the set of options used for backup operations. Each set of options
// is only applied to the backup operation it's passed to. To use the same set
// of options for multiple backup operations pass the struct to all operations.
type BackupConfig struct {
	FailureHandling      FailurePolicy                      `json:"failureHandling"`
	ItemExtensionFactory []extensions.CreateItemExtensioner `json:"-"`
	Parallelism          Parallelism                        `json:"parallelism"`
	ServiceRateLimiter   RateLimiter                        `json:"serviceRateLimiter"`
	Incrementals         IncrementalsConfig                 `json:"incrementalsConfig"`
	M365                 BackupM365Config                   `json:"m365Config"`

	// PreviewItemLimits defines the number of items and/or amount of data to
	// fetch on a best-effort basis for preview backups.
	//
	// Since this is not split out by service or data categories these limits
	// apply independently to all data categories that appear in a single backup
	// where they are set. For example, if doing a teams backup and there's both a
	// SharePoint site and Messages available, both data categories would try to
	// backup data until the set limits without paying attention to what the other
	// had already backed up.
	PreviewLimits PreviewItemLimits `json:"previewItemLimits"`
}

// BackupM365Config contains config options that are specific to backing up data
// from M365 or Corso features that are only available during M365 backups.
type BackupM365Config struct {
	// DeltaPageSize controls the quantity of items fetched in each page during
	// multi-page queries, such as graph api delta endpoints.
	DeltaPageSize int32 `json:"deltaPageSize"`

	// DisableDelta prevents backups from using delta based lookups,
	// forcing a backup by enumerating all items. This is different
	// from IncrementalsConfig.ForceFullEnumeration in that this does not even
	// make use of delta endpoints if a delta token is available. This is
	// necessary when the user has filled up the mailbox storage available to the
	// user as Microsoft prevents the API from being able to make calls
	// to delta endpoints.
	DisableDeltaEndpoint bool `json:"exchangeDeltaEndpoint,omitempty"`

	// ExchangeImmutableIDs denotes whether Corso should store items with
	// immutable Exchange IDs. This is only safe to set if the previous backup for
	// incremental backups used immutable IDs or if a full backup is being done.
	ExchangeImmutableIDs bool `json:"exchangeImmutableIDs,omitempty"`

	// see: https://github.com/alcionai/corso/issues/4688
	UseDriveDeltaTree bool `json:"useDriveDeltaTree"`
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

// IncrementalsConfig contains options specific to incremental backups and
// affects what data will be fetched from the external service being backed up.
type IncrementalsConfig struct {
	// ForceFullEnumeration prevents the use of a previous backup as the starting
	// point for the current backup. All data in the external service will be
	// discovered whether or not it's changed.
	ForceFullEnumeration bool `json:"forceFullEnumeration,omitempty"`

	// ForceItemDataRefresh causes data for all discovered items to be downloaded
	// from the external service instead of using unchanged data from previous
	// failed or successful backups where possible. Data dedupe will still occur
	// if the redownloaded data matches data previously backed up by corso.
	//
	// To control what items are discovered for the backup, see
	// ForceFullEnumeration.
	ForceItemDataRefresh bool `json:"forceItemDataRefresh,omitempty"`
}
