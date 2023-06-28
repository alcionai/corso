package utils

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/control"
)

type RestoreCfgOpts struct {
	Collisions  string
	Destination string

	Populated flags.PopulatedFlags
}

func makeRestoreCfgOpts(cmd *cobra.Command) RestoreCfgOpts {
	return RestoreCfgOpts{
		Collisions:  flags.CollisionsFV,
		Destination: flags.DestinationFV,

		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// validateRestoreConfigFlags checks common flags for correctness and interdependencies
func validateRestoreConfigFlags(opts RestoreCfgOpts) error {
	_, populated := opts.Populated[flags.CollisionsFN]
	_, valid := control.ValidCollisionPolicies()[control.CollisionPolicy(flags.CollisionsFV)]

	if populated && valid {
		return clues.New("invalid entry for " + flags.CollisionsFN)
	}

	return nil
}

func MakeRestoreConfig(
	ctx context.Context,
	opts RestoreCfgOpts,
	locationTimeFormat dttm.TimeFormat,
) control.RestoreConfig {
	restoreCfg := control.DefaultRestoreConfig(locationTimeFormat)

	if _, ok := opts.Populated[flags.CollisionsFN]; ok {
		restoreCfg.OnCollision = control.CollisionPolicy(opts.Collisions)
	}

	if _, ok := opts.Populated[flags.DestinationFN]; ok {
		restoreCfg.Location = opts.Destination
	}

	Infof(ctx, "Restoring to folder %s", restoreCfg.Location)

	return restoreCfg
}
