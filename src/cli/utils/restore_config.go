package utils

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/dttm"
)

type RestoreCfgOpts struct {
	Collisions  string
	Destination string
	// DTTMFormat is the timestamp format appended
	// to the default folder name.  Defaults to
	// dttm.HumanReadable.
	DTTMFormat        dttm.TimeFormat
	ProtectedResource string
	SkipPermissions   bool

	Populated flags.PopulatedFlags
}

func makeRestoreCfgOpts(cmd *cobra.Command) RestoreCfgOpts {
	return RestoreCfgOpts{
		Collisions:        flags.CollisionsFV,
		Destination:       flags.DestinationFV,
		DTTMFormat:        dttm.HumanReadable,
		ProtectedResource: flags.ToResourceFV,
		SkipPermissions:   flags.NoPermissionsFV,

		// populated contains the list of flags that appear in the
		// command, according to pflags.  Use this to differentiate
		// between an "empty" and a "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// ValidateRestoreConfigFlags checks common restore flags for correctness and interdependencies.
func ValidateRestoreConfigFlags(opts RestoreCfgOpts) error {
	_, populated := opts.Populated[flags.CollisionsFN]
	isValid := control.IsValidCollisionPolicy(control.CollisionPolicy(opts.Collisions))

	if populated && !isValid {
		return clues.New(fmt.Sprintf("invalid collision policy: %s", flags.CollisionsFN))
	}

	return nil
}

func MakeRestoreConfig(
	ctx context.Context,
	opts RestoreCfgOpts,
) control.RestoreConfig {
	if len(opts.DTTMFormat) == 0 {
		opts.DTTMFormat = dttm.HumanReadable
	}

	restoreCfg := control.DefaultRestoreConfig(opts.DTTMFormat)

	if _, ok := opts.Populated[flags.CollisionsFN]; ok {
		restoreCfg.OnCollision = control.CollisionPolicy(opts.Collisions)
	}

	if _, ok := opts.Populated[flags.DestinationFN]; ok {
		restoreCfg.Location = opts.Destination
	}

	restoreCfg.ProtectedResource = opts.ProtectedResource
	restoreCfg.IncludePermissions = !opts.SkipPermissions

	Infof(ctx, "Restoring to folder %s", restoreCfg.Location)

	return restoreCfg
}
