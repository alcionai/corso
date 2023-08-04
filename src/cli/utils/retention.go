package utils

import (
	"time"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/control/repository"
)

type retentionCfgOpts struct {
	Mode     string
	Duration time.Duration
	Extend   bool

	Populated flags.PopulatedFlags
}

func makeRetentionCfgOpts(cmd *cobra.Command) retentionCfgOpts {
	return retentionCfgOpts{
		Mode:     flags.RetentionModeFV,
		Duration: flags.RetentionDurationFV,
		Extend:   flags.ExtendRetentionFV,

		// Populated contains the list of flags that appear in the command,
		// according to pflags. Use this to differentiate between an "empty" and a
		// "missing" value.
		Populated: flags.GetPopulatedFlags(cmd),
	}
}

// validate checks common restore flags for correctness.
func (opts retentionCfgOpts) validate() error {
	// Mode defaults to a valid value when pulled from flags. If coming from an
	// empty struct will return invalid error.
	if _, ok := repository.ValidRetentionModeNames()[opts.Mode]; !ok {
		return clues.New("invalid retention mode " + opts.Mode)
	}

	// TODO(ashmrtn): Add an upper bound check?
	if opts.Duration < 0 {
		return clues.New("negative retention duration")
	}

	return nil
}

// MakeRetentionConfig converts the current retentionCfgOpts into a
// repository.Retention struct for use in lower-layers of corso.
func (opts retentionCfgOpts) makeRetentionOpts() (repository.Retention, error) {
	retention := repository.Retention{}

	if err := opts.validate(); err != nil {
		return retention, clues.Stack(err)
	}

	// Only populate the fields that the user passed so that we don't accidentally
	// change retention values without meaning to. Even if the user passed the
	// same value as the default for the flag it gets marked as populated.
	if _, ok := opts.Populated[flags.RetentionModeFN]; ok {
		mode, ok := repository.ValidRetentionModeNames()[opts.Mode]
		if !ok {
			// Not sure how we'd get here since we validate above, but just in case.
			return retention, clues.New("invalid retention mode " + opts.Mode)
		}

		retention.Mode = ptr.To(mode)
	}

	if _, ok := opts.Populated[flags.RetentionDurationFN]; ok {
		retention.Duration = ptr.To(opts.Duration)
	}

	if _, ok := opts.Populated[flags.ExtendRetentionFN]; ok {
		retention.Extend = ptr.To(opts.Extend)
	}

	return retention, nil
}

func MakeRetentionOpts(cmd *cobra.Command) (repository.Retention, error) {
	opts, err := makeRetentionCfgOpts(cmd).makeRetentionOpts()
	return opts, clues.Stack(err).OrNil()
}
