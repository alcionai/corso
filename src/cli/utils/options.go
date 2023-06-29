package utils

import (
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/control"
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

func InitConcurrencyControls(args []string) {
	// FIXME: normally we want to control the limit according to the
	// control options.  But we're racing for initialization here, so
	// we'll set it to the default max as a starting point.
	if slices.Contains(args, "exchange") {
		graph.InitializeConcurrencyLimiter(4)
	}
}
