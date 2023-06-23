package utils

import (
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/cli/flags"
)

// Control produces the control options based on the user's flags.
func Control() control.Options {
	opt := control.Defaults()

	if flags.FailFastFV {
		opt.FailureHandling = control.FailFast
	}

	opt.DisableMetrics = flags.NoStatsFV
	opt.RestorePermissions = flags.RestorePermissionsFV
	opt.SkipReduce = flags.SkipReduceFV
	opt.ToggleFeatures.DisableIncrementals = flags.DisableIncrementalsFV
	opt.ToggleFeatures.DisableDelta = flags.DisableDeltaFV
	opt.ToggleFeatures.ExchangeImmutableIDs = flags.EnableImmutableIDFV
	opt.ToggleFeatures.DisableConcurrencyLimiter = flags.DisableConcurrencyLimiterFV
	opt.Parallelism.ItemFetch = flags.FetchParallelismFV

	return opt
}