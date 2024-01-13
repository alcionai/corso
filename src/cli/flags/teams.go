package flags

import (
	"github.com/spf13/cobra"
)

const (
	TeamFN = "team"
)

var TeamFV []string

func AddTeamDetailsAndRestoreFlags(cmd *cobra.Command) {
	// TODO: implement flags
}

// AddTeamFlag adds the --team flag, which accepts id or name values.
// TODO: need to decide what the appropriate "name" to accept here is.
// keepers thinks its either DisplayName or MailNickname or Mail
// Mail is most accurate, MailNickame is accurate and shorter, but the end user
// may not see either one visibly.
// https://learn.microsoft.com/en-us/graph/api/team-list?view=graph-rest-1.0&tabs=http
func AddTeamFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&TeamFV,
		TeamFN, nil,
		"Backup data by team; accepts '"+Wildcard+"' to select all teams.")
}
