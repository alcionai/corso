package options

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/credentials"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.Defaults()

	if failFastFV {
		opt.FailureHandling = control.FailFast
	}

	opt.DisableMetrics = noStatsFV
	opt.RestorePermissions = restorePermissionsFV
	opt.SkipReduce = skipReduceFV
	opt.ToggleFeatures.DisableIncrementals = disableIncrementalsFV
	opt.ToggleFeatures.DisableDelta = disableDeltaFV
	opt.ToggleFeatures.ExchangeImmutableIDs = enableImmutableID
	opt.ToggleFeatures.DisableConcurrencyLimiter = disableConcurrencyLimiterFV
	opt.Parallelism.ItemFetch = fetchParallelismFV

	return opt
}

// ---------------------------------------------------------------------------
// Operations Flags
// ---------------------------------------------------------------------------

const (
	FailFastFN                  = "fail-fast"
	FetchParallelismFN          = "fetch-parallelism"
	NoStatsFN                   = "no-stats"
	RestorePermissionsFN        = "restore-permissions"
	SkipReduceFN                = "skip-reduce"
	DisableDeltaFN              = "disable-delta"
	DisableIncrementalsFN       = "disable-incrementals"
	EnableImmutableIDFN         = "enable-immutable-id"
	DisableConcurrencyLimiterFN = "disable-concurrency-limiter"
)

var (
	failFastFV           bool
	fetchParallelismFV   int
	noStatsFV            bool
	restorePermissionsFV bool
	skipReduceFV         bool
)

// s3 bucket info from flags
var (
	Bucket          string
	Endpoint        string
	Prefix          string
	DoNotUseTLS     bool
	DoNotVerifyTLS  bool
	AccessKey       string
	SecretAccessKey string
	SessionToken    string
)

// AddGlobalOperationFlags adds the global operations flag set.
func AddGlobalOperationFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.BoolVar(&noStatsFV, NoStatsFN, false, "disable anonymous usage statistics gathering")

	// Flags addition ordering should follow the order we want them to appear in help and docs:
	// More generic and more frequently used flags take precedence.

	// S3 flags
	fs.StringVar(&Bucket, "bucket", "", "Name of S3 bucket for repo. (required)")
	fs.StringVar(&Prefix, "prefix", "", "Repo prefix within bucket.")
	fs.StringVar(&Endpoint, "endpoint", "s3.amazonaws.com", "S3 service endpoint.")
	fs.BoolVar(&DoNotUseTLS, "disable-tls", false, "Disable TLS (HTTPS)")
	fs.BoolVar(&DoNotVerifyTLS, "disable-tls-verification", false, "Disable TLS (HTTPS) certificate verification.")
	fs.StringVar(&AccessKey, "access-key", "", "S3 access key")
	fs.StringVar(&SecretAccessKey, "secret-access-key", "", "S3 access key")
	fs.StringVar(&SessionToken, "session-token", "", "S3 session token")

	// M365 flags
	fs.StringVar(&credentials.AClientID, "client-id", "", "Azure app client ID")
	fs.StringVar(&credentials.AClientSecret, "client-secret", "", "Azure app client secret")
}

// AddFailFastFlag adds a flag to toggle fail-fast error handling behavior.
func AddFailFastFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&failFastFV, FailFastFN, false, "stop processing immediately if any error occurs")
	// TODO: reveal this flag when fail-fast support is implemented
	cobra.CheckErr(fs.MarkHidden(FailFastFN))
}

// AddRestorePermissionsFlag adds OneDrive flag for restoring permissions
func AddRestorePermissionsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&restorePermissionsFV, RestorePermissionsFN, false, "Restore permissions for files and folders")
}

// AddSkipReduceFlag adds a hidden flag that allows callers to skip the selector
// reduction step.  Currently only intended for details commands, not restore.
func AddSkipReduceFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&skipReduceFV, SkipReduceFN, false, "Skip the selector reduce filtering")
	cobra.CheckErr(fs.MarkHidden(SkipReduceFN))
}

// AddFetchParallelismFlag adds a hidden flag that allows callers to reduce call
// paralellism (ie, the corso worker pool size) from 4 to as low as 1.
func AddFetchParallelismFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.IntVar(
		&fetchParallelismFV,
		FetchParallelismFN,
		4,
		"Control the number of concurrent data fetches for Exchange. Valid range is [1-4]. Default: 4")
	cobra.CheckErr(fs.MarkHidden(FetchParallelismFN))
}

// ---------------------------------------------------------------------------
// Feature Flags
// ---------------------------------------------------------------------------

var (
	disableIncrementalsFV bool
	disableDeltaFV        bool
)

// Adds the hidden '--disable-incrementals' cli flag which, when set, disables
// incremental backups.
func AddDisableIncrementalsFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableIncrementalsFV,
		DisableIncrementalsFN,
		false,
		"Disable incremental data retrieval in backups.")
	cobra.CheckErr(fs.MarkHidden(DisableIncrementalsFN))
}

// Adds the hidden '--disable-delta' cli flag which, when set, disables
// delta based backups.
func AddDisableDeltaFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableDeltaFV,
		DisableDeltaFN,
		false,
		"Disable delta based data retrieval in backups.")
	cobra.CheckErr(fs.MarkHidden(DisableDeltaFN))
}

var enableImmutableID bool

// Adds the hidden '--enable-immutable-id' cli flag which, when set, enables
// immutable IDs for Exchange
func AddEnableImmutableIDFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&enableImmutableID,
		EnableImmutableIDFN,
		false,
		"Enable exchange immutable ID.")
	cobra.CheckErr(fs.MarkHidden(EnableImmutableIDFN))
}

var disableConcurrencyLimiterFV bool

// AddDisableConcurrencyLimiterFlag adds a hidden cli flag which, when set,
// removes concurrency limits when communicating with graph API. This
// flag is only relevant for exchange backups for now
func AddDisableConcurrencyLimiterFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&disableConcurrencyLimiterFV,
		DisableConcurrencyLimiterFN,
		false,
		"Disable concurrency limiter middleware. Default: false")
	cobra.CheckErr(fs.MarkHidden(DisableConcurrencyLimiterFN))
}
