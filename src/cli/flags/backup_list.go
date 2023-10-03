package flags

import "github.com/spf13/cobra"

const Show = "show"

func AddAllBackupListFlags(cmd *cobra.Command) {
	AddFailedItemsFN(cmd)
	AddSkippedItemsFN(cmd)
	AddRecoveredErrorsFN(cmd)
	AddAlertsFN(cmd)
}

func AddFailedItemsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&ListFailedItemsFV, FailedItemsFN, Show,
		"Toggles showing or hiding the list of items that failed.")
}

func AddSkippedItemsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&ListSkippedItemsFV, SkippedItemsFN, Show,
		"Toggles showing or hiding the list of items that were skipped.")
}

func AddRecoveredErrorsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&ListRecoveredErrorsFV, RecoveredErrorsFN, Show,
		"Toggles showing or hiding the list of errors which Corso recovered from.")
}

func AddAlertsFN(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&ListAlertsFV, AlertsFN, Show,
		"Toggles showing or hiding the list of Alerts produced during the operation.")
}
