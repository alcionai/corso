package utils

import (
	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/control"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.DefaultOptions()

	if flags.FailFastFV {
		opt.FailureHandling = control.FailFast
	}

	opt.DisableMetrics = flags.NoStatsFV
	opt.SkipReduce = flags.SkipReduceFV
	opt.ToggleFeatures.DisableIncrementals = flags.DisableIncrementalsFV
	opt.ToggleFeatures.DisableDelta = flags.DisableDeltaFV
	opt.ToggleFeatures.ExchangeImmutableIDs = flags.EnableImmutableIDFV
	opt.ToggleFeatures.DisableConcurrencyLimiter = flags.DisableConcurrencyLimiterFV
	opt.Parallelism.ItemFetch = flags.FetchParallelismFV

	return opt
}

func ControlWithConfig(cfg config.RepoDetails) control.Options {
	opt := Control()

	opt.Repo.User = cfg.RepoUser
	opt.Repo.Host = cfg.RepoHost

	return opt
}
