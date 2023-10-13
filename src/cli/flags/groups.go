package flags

import (
	"github.com/spf13/cobra"
)

const DataMessages = "messages"

const (
	ChannelFN   = "channel"
	GroupFN     = "group"
	MessageFN   = "message"
	GroupSiteFN = "site"

	MessageCreatedAfterFN    = "message-created-after"
	MessageCreatedBeforeFN   = "message-created-before"
	MessageLastReplyAfterFN  = "message-last-reply-after"
	MessageLastReplyBeforeFN = "message-last-reply-before"
)

var (
	ChannelFV []string
	GroupFV   []string
	MessageFV []string

	// Have to create a separate one for restore to make sure we
	// create on which accepts only one value
	GroupSiteFV string

	MessageCreatedAfterFV    string
	MessageCreatedBeforeFV   string
	MessageLastReplyAfterFV  string
	MessageLastReplyBeforeFV string
)

func AddSingleSiteIDFlag(cmd *cobra.Command, required bool) {
	cmd.Flags().StringVar(
		&GroupSiteFV,
		GroupSiteFN,
		"",
		"ID or URL of the site to restore.")

	// TODO: we can move this check to runtime once we support
	// restoring other resources from groups, ie chat messages and
	// only cause it to fail when the user is trying to restore site
	// data.
	if required {
		cobra.CheckErr(cmd.MarkFlagRequired(GroupSiteFN))
	}
}

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
func AddGroupFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&GroupFV,
		GroupFN, nil,
		"Backup data by group; accepts '"+Wildcard+"' to select all groups.")
}
