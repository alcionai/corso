package flags

import (
	"github.com/spf13/cobra"
)

func AddGenericBackupFlags(cmd *cobra.Command) {
	AddFailFastFlag(cmd)
	AddDisableIncrementalsFlag(cmd)
	AddForceItemDataDownloadFlag(cmd)
}
