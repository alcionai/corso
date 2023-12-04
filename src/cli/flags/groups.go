package flags

import (
	"github.com/spf13/cobra"
)

const (
	DataMessages      = "messages"
	DataConversations = "conversations"
)

const (
	ChannelFN      = "channel"
	ConversationFN = "conversation"
	GroupFN        = "group"
	MessageFN      = "message"
	PostFN         = "post"

	MessageCreatedAfterFN    = "message-created-after"
	MessageCreatedBeforeFN   = "message-created-before"
	MessageLastReplyAfterFN  = "message-last-reply-after"
	MessageLastReplyBeforeFN = "message-last-reply-before"
)

var (
	ChannelFV      []string
	ConversationFV []string
	GroupFV        []string
	MessageFV      []string
	PostFV         []string

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

	fs.StringSliceVar(
		&ConversationFV,
		ConversationFN, nil,
		"Select data within a Group's Conversation.")

	fs.StringSliceVar(
		&PostFV,
		PostFN, nil,
		"Select Conversation Posts by reference.")
}

// AddGroupFlag adds the --group flag, which accepts either the id,
// the display name, or the mailbox address as its values.  Users are
// expected to supply the display name.  The ID is supported becase, well,
// IDs.  The mailbox address is supported as a lookup fallback for certain
// SDK cases, therefore it's also supported here, though that support
// isn't exposed to end users.
func AddGroupFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(
		&GroupFV,
		GroupFN, nil,
		"Backup data by group; accepts '"+Wildcard+"' to select all groups.")
}
