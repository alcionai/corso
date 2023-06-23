package flags

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	UserFN    = "user"
	MailBoxFN = "mailbox"
)

var UserFV []string


// AddUserFlag adds the --user flag.
func AddUserFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&UserFV,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")
	cobra.CheckErr(cmd.MarkFlagRequired(UserFN))
}

// AddMailBoxFlag adds the --user and --mailbox flag.
func AddMailBoxFlag(cmd *cobra.Command) {
	flags := cmd.Flags()

	flags.StringSliceVar(
		&UserFV,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")

	cobra.CheckErr(flags.MarkDeprecated(UserFN, fmt.Sprintf("use --%s instead", MailBoxFN)))

	flags.StringSliceVar(
		&UserFV,
		MailBoxFN, nil,
		"Backup a specific mailbox's data; accepts '"+Wildcard+"' to select all mailbox.")
}