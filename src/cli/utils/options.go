package utils

import (
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/config"
	"github.com/alcionai/corso/src/pkg/control"
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
	opt.ToggleFeatures.DisableDelta = flags.DisableDeltaFV
	opt.ToggleFeatures.DisableSlidingWindowLimiter = flags.DisableSlidingWindowLimiterFV
	opt.ToggleFeatures.DisableLazyItemReader = flags.DisableLazyItemReaderFV
	opt.ToggleFeatures.ExchangeImmutableIDs = flags.EnableImmutableIDFV
	opt.ToggleFeatures.UseDeltaTree = flags.UseDeltaTreeFV
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

	opt.Incrementals.ForceFullEnumeration = flags.DisableIncrementalsFV
	opt.Incrementals.ForceItemDataRefresh = flags.ForceItemDataDownloadFV

	return opt
}
