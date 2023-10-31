package testdata

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/cli/flags"
)

func PreparedChannelFlags() []string {
	return []string{
		"--" + flags.ChannelFN, FlgInputs(ChannelInput),
		"--" + flags.MessageFN, FlgInputs(MessageInput),
		"--" + flags.MessageCreatedAfterFN, MessageCreatedAfterInput,
		"--" + flags.MessageCreatedBeforeFN, MessageCreatedBeforeInput,
		"--" + flags.MessageLastReplyAfterFN, MessageLastReplyAfterInput,
		"--" + flags.MessageLastReplyBeforeFN, MessageLastReplyBeforeInput,
	}
}

func AssertChannelFlags(t *testing.T, cmd *cobra.Command) {
	assert.ElementsMatch(t, ChannelInput, flags.ChannelFV)
	assert.ElementsMatch(t, MessageInput, flags.MessageFV)
	assert.Equal(t, MessageCreatedAfterInput, flags.MessageCreatedAfterFV)
	assert.Equal(t, MessageCreatedBeforeInput, flags.MessageCreatedBeforeFV)
	assert.Equal(t, MessageLastReplyAfterInput, flags.MessageLastReplyAfterFV)
	assert.Equal(t, MessageLastReplyBeforeInput, flags.MessageLastReplyBeforeFV)
}

func PreparedConversationFlags() []string {
	return []string{
		"--" + flags.ConversationFN, FlgInputs(ConversationInput),
		"--" + flags.PostFN, FlgInputs(PostInput),
	}
}

func AssertConversationFlags(t *testing.T, cmd *cobra.Command) {
	assert.Equal(t, ConversationInput, flags.ConversationFV)
	assert.Equal(t, PostInput, flags.PostFV)
}
