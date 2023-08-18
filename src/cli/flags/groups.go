package flags

import (
	"github.com/spf13/cobra"
)

const (
	GroupFN = "group"
)

var GroupFV []string

func AddGroupDetailsAndRestoreFlags(cmd *cobra.Command) {
	// TODO: implement flags
}

// AddGroupFlag adds the --group flag, which accepts id or name values.
// TODO: need to decide what the appropriate "name" to accept here is.
// keepers thinks its either DisplayName or MailNickname or Mail
// Mail is most accurate, MailNickame is accurate and shorter, but the end user
// may not see either one visibly.
// https://learn.microsoft.com/en-us/graph/api/group-list?view=graph-rest-1.0&tabs=http
func AddGroupFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&GroupFV,
		GroupFN, nil,
		"Backup data by group; accepts '"+Wildcard+"' to select all groups.")
}
