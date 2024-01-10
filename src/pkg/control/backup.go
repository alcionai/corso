package control

// DefaultBackupOptions provides a Backup with the default values set.
func DefaultBackupConfig() BackupConfig {
	return BackupConfig{}
}

// BackupConfig is the set of options used for backup operations. Each set of
// options is only applied to the backup operation it's passed to. To use the
// same set of options for multiple backup operations pass the struct to all
// operations.
type BackupConfig struct {
	Incrementals IncrementalsConfig `json:"incrementalsConfig"`

	// PreviewLimits defines the number of items and/or amount of data to fetch on
	// a best-effort basis for preview backups.
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

// IncrementalsConfig contains options specific to incremental backups and
// affects what data will be fetched from the external service being backed up.
type IncrementalsConfig struct {
	// ForceFullEnumeration prevents the use of a previous backup as the starting
	// point for the current backup. All data in the external service will be
	// enumerated whether or not it's changed. Per-item storage will only get
	// updated if changes have occurred.
	ForceFullEnumeration bool `json:"forceFullEnumeration,omitempty"`

	// ForceItemDataRefresh causes the data for all enumerated items to replace
	// stored data, even if no changes have been detected. Storage-side data
	// deduplication still applies, but that's after item download, and items are
	// always downloaded when this flag is set.
	ForceItemDataRefresh bool `json:"forceItemDataRefresh,omitempty"`
}
