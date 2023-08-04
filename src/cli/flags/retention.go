package flags

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control/repository"
)

const (
	RetentionModeFN     = "retention-mode"
	RetentionDurationFN = "retention-duration"
	ExtendRetentionFN   = "extend-retention"
)

var (
	RetentionModeFV     string
	RetentionDurationFV time.Duration
	ExtendRetentionFV   bool
)

// AddRetentionConfigFlags adds the retention config flag set.
func AddRetentionConfigFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&RetentionModeFV,
		RetentionModeFN,
		repository.NoRetention.String(),
		"Sets object locking mode (if any) to use in remote storage: "+
			repository.NoRetention.String()+", "+
			repository.GovernanceRetention.String()+", or "+
			repository.ComplianceRetention.String())
	cobra.CheckErr(fs.MarkHidden(RetentionModeFN))

	fs.DurationVar(
		&RetentionDurationFV,
		RetentionDurationFN,
		time.Duration(0),
		"Set the amount of time to lock individual objects in remote storage")
	cobra.CheckErr(fs.MarkHidden(RetentionDurationFN))

	fs.BoolVar(
		&ExtendRetentionFV,
		ExtendRetentionFN,
		false,
		"Extends object locks during maintenance. "+
			"Extends locks by the most recently set value of "+RetentionDurationFN)
	cobra.CheckErr(fs.MarkHidden(ExtendRetentionFN))
}
