package flags

import (
	"github.com/spf13/cobra"
)

const DataMessages = "messages"

const (
	ChannelFN = "channel"
	GroupFN   = "group"
	MessageFN = "message"

	MessageCreatedAfterFN    = "message-created-after"
	MessageCreatedBeforeFN   = "message-created-before"
	MessageLastReplyAfterFN  = "message-last-reply-after"
	MessageLastReplyBeforeFN = "message-last-reply-before"
)

var (
	ChannelFV []string
	GroupFV   []string
	MessageFV []string

	MessageCreatedAfterFV    string
	MessageCreatedBeforeFV   string
	MessageLastReplyAfterFV  string
	MessageLastReplyBeforeFV string
)

func AddGroupDetailsAndRestoreFlags(cmd *cobra.Command) {
	fs := cmd.Flags()

	fs.StringSliceVar(
		&ChannelFV,
		ChannelFN, nil,
		"Select data within a Team's Channel.")

	fs.StringSliceVar(
		&MessageFV,
		MessageFN, nil,
		"Select messages by reference.")

	fs.StringVar(
		&MessageCreatedAfterFV,
		MessageCreatedAfterFN, "",
		"Select messages created after this datetime.")

	fs.StringVar(
		&MessageCreatedBeforeFV,
		MessageCreatedBeforeFN, "",
		"Select messages created before this datetime.")

	fs.StringVar(
		&MessageLastReplyAfterFV,
		MessageLastReplyAfterFN, "",
		"Select messages with replies after this datetime.")

	fs.StringVar(
		&MessageLastReplyBeforeFV,
		MessageLastReplyBeforeFN, "",
		"Select messages with replies before this datetime.")
}

// AddGroupFlag adds the --group flag, which accepts id or name values.
// TODO: need to decide what the appropriate "name" to accept here is.
// keepers thinks its either DisplayName or MailNickname or Mail
// Mail is most accurate, MailNickame is accurate and shorter, but the end user
// may not see either one visibly.
// https://learn.microsoft.com/en-us/graph/api/group-list?view=graph-rest-1.0&tabs=http
func AddGroupFlag(
	cmd *cobra.Command,
	completionFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective),
) {
	cmd.Flags().StringSliceVar(
		&GroupFV,
		GroupFN, nil,
		"Backup data by group; accepts '"+Wildcard+"' to select all groups.")

	cobra.CheckErr(cmd.RegisterFlagCompletionFunc(GroupFN, completionFunc))
}
