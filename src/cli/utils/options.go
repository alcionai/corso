package utils

import (
	"github.com/alcionai/canario/src/cli/flags"
	"github.com/alcionai/canario/src/pkg/config"
	"github.com/alcionai/canario/src/pkg/control"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.DefaultOptions()

	if flags.FailFastFV {
		opt.FailureHandling = control.FailFast
	}

	dps := int32(flags.DeltaPageSizeFV)
	if dps > 500 || dps < 1 {
		dps = 500
	}

	opt.DeltaPageSize = dps
	opt.DisableMetrics = flags.NoStatsFV
	opt.SkipReduce = flags.SkipReduceFV
	opt.ToggleFeatures.DisableIncrementals = flags.DisableIncrementalsFV
	opt.ToggleFeatures.ForceItemDataDownload = flags.ForceItemDataDownloadFV
	opt.ToggleFeatures.DisableDelta = flags.DisableDeltaFV
	opt.ToggleFeatures.DisableSlidingWindowLimiter = flags.DisableSlidingWindowLimiterFV
	opt.ToggleFeatures.DisableLazyItemReader = flags.DisableLazyItemReaderFV
	opt.ToggleFeatures.ExchangeImmutableIDs = flags.EnableImmutableIDFV
	opt.ToggleFeatures.UseOldDeltaProcess = flags.UseOldDeltaProcessFV
	opt.Parallelism.ItemFetch = flags.FetchParallelismFV

	return opt
}

func ControlWithConfig(cfg config.RepoDetails) control.Options {
	opt := Control()

	opt.Repo.User = cfg.RepoUser
	opt.Repo.Host = cfg.RepoHost

	return opt
}

func ParseBackupOptions() control.BackupConfig {
	opt := control.DefaultBackupConfig()

	if flags.FailFastFV {
		opt.FailureHandling = control.FailFast
	}

	dps := int32(flags.DeltaPageSizeFV)
	if dps > 500 || dps < 1 {
		dps = 500
	}

	opt.M365.DeltaPageSize = dps
	opt.M365.DisableDeltaEndpoint = flags.DisableDeltaFV
	opt.M365.ExchangeImmutableIDs = flags.EnableImmutableIDFV
	opt.M365.UseOldDriveDeltaProcess = flags.UseOldDeltaProcessFV
	opt.ServiceRateLimiter.DisableSlidingWindowLimiter = flags.DisableSlidingWindowLimiterFV
	opt.Parallelism.ItemFetch = flags.FetchParallelismFV
	opt.Incrementals.ForceFullEnumeration = flags.DisableIncrementalsFV
	opt.Incrementals.ForceItemDataRefresh = flags.ForceItemDataDownloadFV

	return opt
}
